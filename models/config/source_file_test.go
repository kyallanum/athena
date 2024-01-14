package models

import (
	"reflect"
	"testing"
)

func TestNewConfigFile(t *testing.T) {
	testTable := []struct {
		name           string
		source         string
		expectedOutput IConfigurationSource
		expectedError  error
	}{
		{
			name:   "Test_New_File_Config",
			source: "../../examples/apt-term-config.json",
			expectedOutput: &FileSource{
				ConfigurationSource: ConfigurationSource{
					source_type: "file",
					source:      "",
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			fileSource := NewFileSource(test.source)

			if fileSource.SourceType() != test.expectedOutput.SourceType() {
				t.Errorf("Source Type was not set properly. \n\tExpected: %s\n\tReceived: %s", test.expectedOutput.SourceType(), fileSource.SourceType())
			}

			if reflect.TypeOf(fileSource).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Configuration source not of the right type. \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(fileSource).String())
			}
		})
	}
}

func TestLoadFileConfig(t *testing.T) {
	testTable := []struct {
		name           string
		source         string
		expectedOutput []byte
		expectedError  error
	}{
		{
			name:           "Test_Bad_String",
			source:         "../../examples/apt-term-config.json",
			expectedOutput: []byte{},
			expectedError:  nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			fileSource := NewFileSource(test.source)

			fileConfig, err := fileSource.Config()

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Configuration Load created an unexpected error: \n\tExpected: %s\n\tReceived: %s", test.expectedError.Error(), err.Error())
			}

			if reflect.TypeOf(fileConfig).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Configuration Load returned the wrong object type: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(fileConfig).String())
			}
		})
	}
}
