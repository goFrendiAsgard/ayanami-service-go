package integrationtest

import (
	"github.com/state-alchemists/ayanami-service-go/config"
	"github.com/state-alchemists/ayanami-service-go/msgbroker"
	"github.com/state-alchemists/ayanami-service-go/service"
	"log"
)

// MainServiceCmd emulating cmd's main function
func MainServiceCmd() {
	serviceName := "cmd"
	// define broker
	broker, err := msgbroker.NewNats(config.GetNatsURL())
	if err != nil {
		log.Fatal(err)
	}
	// define services
	services := service.Services{
		service.NewCmd(serviceName, "cowsay",
			[]string{"/bin/sh", "-c", "echo $input | cowsay -n"},
		),
		service.NewCmd(serviceName, "figlet",
			[]string{"figlet", "$input"},
		),
	}
	// consume and publish forever
	forever := make(chan bool)
	services.ConsumeAndPublish(broker, serviceName)
	<-forever
}
