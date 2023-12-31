package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSourceFile(t *testing.T) {
	source := "../../examples/apt-term-config.json"

	configSource, err := Source(source)
	fmt.Println(reflect.TypeOf(configSource).String())
	if err != nil {
		t.Errorf("Error returned from getSource when it shouldn't have: \n\t%s", err.Error())
	}

	if reflect.TypeOf(configSource).String() != "*models.FileSource" {
		t.Errorf("Correct data type not returned from GetSource")
	}
}

func TestGetSourceFileBadFile(t *testing.T) {
	source := "../../examples/apt-term.json"

	_, err := Source(source)

	if err == nil {
		t.Errorf("No error was returned when one should have been.")
	}

	if err.Error() != "unable to get file information for file: ../../examples/apt-term.json. error: stat ../../examples/apt-term.json: no such file or directory" {
		t.Errorf("Incorrect error returned when attempting to stat a non-existant file.")
	}
}

func TestGetSourceWeb(t *testing.T) {
	source := "http://example.org"

	configWeb, err := Source(source)
	if err != nil {
		t.Errorf("Error returned from getSource when it shouldn't have: \n\t%s", err.Error())
	}

	if reflect.TypeOf(configWeb).String() != "*models.WebSource" {
		t.Errorf("Correct data type not returned from GetSource: \n\t%s", reflect.TypeOf(configWeb).String())
	}
}

func TestGetSourceWebBadAddress(t *testing.T) {
	source := "http://192.168.0.1/index.html"

	_, err := Source(source)

	if err == nil {
		t.Errorf("Error not returned from getSource when one should have")
	}

	if err.Error() != "url provided is not reachable. please check the URL and try again" {
		t.Errorf("Error not returned properly from getSource when one should have: \n\t%s", err.Error())
	}
}
