package service

import (
	"reflect"
	"testing"
)

func TestIsStringInArray(t *testing.T) {
	actual := IsStringInArray("shinji", []string{"rei", "shinji", "asuka"})
	if true != actual {
		t.Errorf("Expected `%t`, get `%t`", true, actual)
	}

	actual = IsStringInArray("gandalf", []string{"rei", "shinji", "asuka"})
	if false != actual {
		t.Errorf("Expected `%t`, get `%t`", false, actual)
	}
}

func TestAppendUniqueArray(t *testing.T) {
	var expected, actual []string

	expected = []string{"rei", "shinji", "asuka"}
	actual = AppendUniqueString("shinji", []string{"rei", "shinji", "asuka"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected `%#v`, get `%#v`", expected, actual)
	}

	expected = []string{"rei", "shinji", "asuka", "gandalf"}
	actual = AppendUniqueString("gandalf", []string{"rei", "shinji", "asuka"})
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected `%#v`, get `%#v`", expected, actual)
	}

}
