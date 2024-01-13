package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSource(t *testing.T) {
	testTable := []struct {
		name           string
		source         string
		expectedOutput IConfigurationSource
		expectedError  error
	}{
		{
			name:           "Test_Good_Source_File",
			source:         "../../examples/apt-term-config.json",
			expectedOutput: &FileSource{},
			expectedError:  nil,
		},
		{
			name:           "Test_Bad_Source_File",
			source:         "../../examples/apt-term-bad.json",
			expectedOutput: nil,
			expectedError:  fmt.Errorf("unable to get file information for file: ../../examples/apt-term-bad.json. error: stat ../../examples/apt-term-bad.json: no such file or directory"),
		},
		{
			name:           "Test_Good_Web_Source",
			source:         "https://raw.githubusercontent.com/kyallanum/athena/main/examples/apt-term-config.json",
			expectedOutput: &WebSource{},
			expectedError:  nil,
		},
		{
			name:           "Test_Bad_Web_Source",
			source:         "http://192.168.0.1",
			expectedOutput: nil,
			expectedError:  fmt.Errorf("url provided is not reachable. please check the URL and try again"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			configSource, err := Source(test.source)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Error immproperly returned from Source function: \n\tExpected: %s\n\tReceived: %s", test.expectedError.Error(), err.Error())
			}

			if configSource != test.expectedOutput {
				if reflect.TypeOf(configSource).String() != reflect.TypeOf(test.expectedOutput).String() {
					t.Errorf("Object returned is not of the right type: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(test.expectedOutput).String())
				}
			}
		})
	}
}
