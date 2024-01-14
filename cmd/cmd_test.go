package cmd

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	config "github.com/kyallanum/athena/models/config"
	models "github.com/kyallanum/athena/models/library"
	logs "github.com/kyallanum/athena/models/logs"
	"github.com/sirupsen/logrus"
)

func checkExpectedError(actualError, expectedError any) bool {
	if actualError == nil && expectedError == nil {
		return true
	}

	if actualError != nil {
		if expectedError != nil {
			if strings.Contains(actualError.(error).Error(), expectedError.(error).Error()) {
				return true
			}
		}
	}
	return false
}

func TestResolveFile(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	testTable := []struct {
		name                    string
		fileName                string
		configFile              string
		expectedReturn          *models.Library
		expectedLogFileError    error
		expectedConfigFileError error
		expectedResolveError    error
	}{
		{
			name:                    "Test_Good_Log_And_Configuration_File",
			fileName:                "../examples/apt-term.log",
			configFile:              "../examples/apt-term-config.json",
			expectedReturn:          &models.Library{},
			expectedLogFileError:    nil,
			expectedConfigFileError: nil,
			expectedResolveError:    nil,
		},
		{
			name:                    "Test_Bad_Log_File",
			fileName:                "../examples/apt-term-bad.log",
			configFile:              "../examples/apt-term-config.json",
			expectedReturn:          nil,
			expectedLogFileError:    fmt.Errorf("../examples/apt-term-bad.log: no such file or directory"),
			expectedConfigFileError: nil,
			expectedResolveError:    fmt.Errorf("log file contains no contents"),
		},
		{
			name:                    "Test_Bad_Config_File",
			fileName:                "../examples/apt-term.log",
			configFile:              "../examples/apt-term-config-bad.json",
			expectedReturn:          nil,
			expectedLogFileError:    nil,
			expectedConfigFileError: fmt.Errorf("../examples/apt-term-config-bad.json: no such file or directory"),
			expectedResolveError:    fmt.Errorf("configuration file has no contents"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			logFile, err := logs.LoadLogFile(test.fileName)
			if !checkExpectedError(err, test.expectedLogFileError) {
				t.Errorf("Error was incorrectly thrown while loading log file when it shouldn't have: \n\tExpected: %v\n\tReceived: %v", test.expectedLogFileError, err)
			}
			configFile, err := config.CreateConfiguration(test.configFile)
			if !checkExpectedError(err, test.expectedConfigFileError) {
				t.Errorf("Error was incorrectly thrown while loading configuration file when it shouldn't have: \n\tExpected: %v\n\tReceived: %v", test.expectedConfigFileError, err)
			}

			library, err := resolveLogFile(logFile, configFile, logger)
			if !checkExpectedError(err, test.expectedResolveError) {
				t.Errorf("Error was incorrectly thrown while resolving log file: \n\tExpected: %v\n\tReceived: %v", test.expectedResolveError, err)
			}

			if reflect.TypeOf(library).String() != reflect.TypeOf(test.expectedReturn).String() {
				t.Errorf("Log file resolution did not return the proper type: Expected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedReturn).String(), reflect.TypeOf(library).String())
			}
		})
	}
}

func TestConfigErrors(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	testTable := []struct {
		name          string
		configuration *config.Configuration
		expectedError error
	}{
		{
			name: "Test_No_Config_Name",
			configuration: &config.Configuration{
				Rules: make([]config.Rule, 1),
			},
			expectedError: fmt.Errorf("configuration file contains no log name"),
		},
		{
			name: "Test_No_Config_Rules",
			configuration: &config.Configuration{
				Name: "test",
			},
			expectedError: fmt.Errorf("configuration does not have any rules"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			logFile, _ := logs.LoadLogFile("../examples/apt-term.log")

			_, err := resolveLogFile(logFile, test.configuration, logger)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Error was not properly returned when checking configuration: \n\tExpected: %v\n\tReceived: %v", test.expectedError, err)
			}
		})
	}
}
