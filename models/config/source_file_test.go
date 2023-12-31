package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileConfig(t *testing.T) {
	source := "./examples/apt-term-config.json"

	fileSource := NewFileSource(source)

	if fileSource.SourceType() != "file" {
		t.Errorf("Configuration source type not set properly")
	}

	if reflect.TypeOf(fileSource).String() != "*models.ConfigFileSource" {
		t.Errorf("Configuration source not of the right type")
	}
}

func TestLoadFileConfig(t *testing.T) {
	source := "../../examples/apt-term-config.json"
	fileSource := NewFileSource(source)

	fileConfig, err := fileSource.Config()
	if err != nil {
		t.Errorf("Configuration Load created an error:\n\t%v", err)
	}

	fmt.Printf("t: %v", reflect.TypeOf(fileConfig))

	if reflect.TypeOf(fileConfig).String() != "[]uint8" {
		t.Errorf("LoadConfig did not return the correct type")
	}
}
