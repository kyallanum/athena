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
			name:                    "Test good log and configuration file",
			fileName:                "../examples/apt-term.log",
			configFile:              "../examples/apt-term-config.json",
			expectedReturn:          &models.Library{},
			expectedLogFileError:    nil,
			expectedConfigFileError: nil,
			expectedResolveError:    nil,
		},
		{
			name:                    "Test bad log file",
			fileName:                "../examples/apt-term-bad.log",
			configFile:              "../examples/apt-term-config.json",
			expectedReturn:          nil,
			expectedLogFileError:    fmt.Errorf("../examples/apt-term-bad.log: no such file or directory"),
			expectedConfigFileError: nil,
			expectedResolveError:    nil,
		},
		{
			name:                    "Test bad config file",
			fileName:                "../examples/apt-term.log",
			configFile:              "../examples/apt-term-config-bad.json",
			expectedReturn:          nil,
			expectedLogFileError:    nil,
			expectedConfigFileError: fmt.Errorf("../examples/apt-term-config-bad.json: no such file or directory"),
			expectedResolveError:    nil,
		},
	}

	for _, currentTest := range testTable {
		t.Run(currentTest.name, func(t *testing.T) {
			logFile, err := logs.LoadLogFile(currentTest.fileName)
			if err != currentTest.expectedLogFileError {
				if !strings.Contains(err.Error(), currentTest.expectedLogFileError.Error()) {
					t.Errorf("Error was incorrectly thrown while loading log file when it shouldn't have: \n\t%s", err.Error())
				}
				return
			}
			configFile, err := config.CreateConfiguration(currentTest.configFile)
			if err != currentTest.expectedConfigFileError {
				if !strings.Contains(err.Error(), currentTest.expectedConfigFileError.Error()) {
					t.Errorf("Error was incorrectly thrown while loading configuration file when it shouldn't have: \n\t%s", err.Error())
				}
				return
			}

			library, err := resolveLogFile(logFile, configFile, logger)
			if err != currentTest.expectedResolveError {
				if !strings.Contains(err.Error(), currentTest.expectedResolveError.Error()) {
					t.Errorf("Error was incorrectly thrown while resolving log file: \n\t%s", err.Error())
				}
				return
			}

			if reflect.TypeOf(library).String() != reflect.TypeOf(currentTest.expectedReturn).String() {
				t.Errorf("Log file resolution did not return the proper type: %s", reflect.TypeOf(library).String())
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

			if err.Error() != test.expectedError.Error() {
				t.Errorf("Error was not properly returned when checking configuration: %s", err.Error())
			}
		})
	}
}
