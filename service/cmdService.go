package service

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// NewCmd create new cmd
func NewCmd(serviceName string, methodName string, command []string) CommonService {
	inputs := getCmdInputs(command)
	outputs := []string{"output"}
	wrappedFunction := createCmdWrapper(serviceName, methodName, command, inputs, outputs)
	return NewService(serviceName, methodName, inputs, outputs, wrappedFunction)
}

func getCmdInputs(command []string) []string {
	var inputs []string
	re1, err := regexp.Compile(`\$([a-zA-Z0-9]+)`)
	if err != nil {
		return inputs
	}
	re2, err := regexp.Compile(`\${([a-zA-Z0-9]+)}`)
	if err != nil {
		return inputs
	}
	for _, part := range command {
		matches1 := re1.FindAllStringSubmatch(part, -1)
		matches2 := re2.FindAllStringSubmatch(part, -1)
		inputs = addMatchesToCmdInputs(inputs, matches1)
		inputs = addMatchesToCmdInputs(inputs, matches2)
	}
	return inputs
}

func addMatchesToCmdInputs(inputs []string, matches [][]string) []string {
	for _, match := range matches {
		inputs = AppendUniqueString(match[1], inputs)
	}
	return inputs
}

func createCmdWrapper(serviceName, methodName string, templateCmd []string, inputVarNames, outputVarNames []string) WrappedFunction {
	return func(inputs Dictionary) (Dictionary, error) {
		outputs := make(Dictionary)
		// preprocess cmd
		cmd := make([]string, len(templateCmd))
		for cmdIndex := range templateCmd {
			cmd[cmdIndex] = templateCmd[cmdIndex]
			for _, varName := range inputVarNames {
				varValue := fmt.Sprintf("%s", inputs.Get(varName))
				pattern1 := fmt.Sprintf("$%s", varName)
				pattern2 := fmt.Sprintf("${%s}", varName)
				// if varValue doesn't started and ended with double quote, add double quote to it. Otherwise, let it be
				if cmd[cmdIndex] != pattern1 && cmd[cmdIndex] != pattern2 {
					varValue = escapeQuoteCmdParam(varValue)
				}
				cmd[cmdIndex] = strings.Replace(cmd[cmdIndex], pattern1, varValue, -1)
				cmd[cmdIndex] = strings.Replace(cmd[cmdIndex], pattern2, varValue, -1)
			}
		}
		// run the command
		log.Printf("[INFO: %s.%s] Executing `%#v`", serviceName, methodName, cmd)
		outByte, err := exec.Command(cmd[0], cmd[1:]...).Output()
		if err != nil {
			log.Printf("[ERROR: %s.%s] Error while executing `%#v`: %s", serviceName, methodName, cmd, err)
			return outputs, err
		}
		// assemble outputs
		outputVal := string(outByte)
		for _, outputVarName := range outputVarNames {
			err := outputs.Set(outputVarName, outputVal)
			if err != nil {
				log.Printf("[ERROR: %s.%s] Error while assigning output `%s`: %s", serviceName, methodName, outputVarName, err)
				return outputs, err
			}
		}
		return outputs, err
	}
}

func escapeQuoteCmdParam(str string) string {
	if len(str) == 0 {
		return str
	}
	runic := []rune(str)
	if string(runic[0]) != `"` || string(runic[len(str)-1]) != `"` {
		escapedStr := strings.Replace(str, `"`, `\"`, -1)
		return fmt.Sprintf(`"%s"`, escapedStr)
	}
	return str
}
