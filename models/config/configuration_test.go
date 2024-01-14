package models

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

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

func TestTranslateRegex(t *testing.T) {
	testTable := []struct {
		name           string
		regex          string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Test_Good_Regex",
			regex:          `(?<test_name>)`,
			expectedOutput: `(?P<test_name>)`,
			expectedError:  nil,
		}, {
			name:           "Test_Empty_Regex",
			regex:          "",
			expectedOutput: "",
			expectedError:  fmt.Errorf("empty search terms are not allowed"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			regexToTranslate := test.regex

			err := translateConfigurationNamedGroups(&regexToTranslate)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("An error occurred while attempting to translate regex: \n\tExpected: %v\n\tReceived: %v\n\t", test.expectedError, err)
			}

			if regexToTranslate != test.expectedOutput {
				t.Errorf("Regex was not translated properly \n\tExpected: %v\n\tReceived: %v", test.expectedOutput, regexToTranslate)
			}
		})
	}
}

func TestTranslateConfiguration(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("Configuration not translated properly.")
		}
	}()

	configFile, _ := os.Open("./examples/apt-term-config.json")
	defer configFile.Close()

	scanner := bufio.NewScanner(configFile)

	configFileLines := make([]byte, 0)
	for scanner.Scan() {
		configFileLines = append(configFileLines, scanner.Bytes()...)
	}

	configObject := &Configuration{}
	json.Unmarshal(configFileLines, &configObject)

	configObject.TranslateConfiguration()
}

func TestCreateConfiguration(t *testing.T) {
	testTable := []struct {
		name           string
		source         string
		expectedOutput *Configuration
		expectedError  error
	}{
		{
			name:           "Test_From_Good_File",
			source:         "../../examples/apt-term-config.json",
			expectedOutput: &Configuration{},
			expectedError:  nil,
		},
		{
			name:           "Test_From_Bad_File",
			source:         "bad_file_name",
			expectedOutput: nil,
			expectedError:  fmt.Errorf("unable to create configuration object: \n\tunable to get file information for file: bad_file_name. error: stat bad_file_name: no such file or directory"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			config, err := CreateConfiguration(test.source)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Error was improperly returned from CreateConfiguration: \n\tExpected: %v\n\tReceived: %v", test.expectedError.Error(), err.Error())
			}

			if reflect.TypeOf(config).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Improper type returned from CreateConfiguration: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(config).String())
			}
		})
	}
}

func TestCreateConfigurationFromWeb(t *testing.T) {
	file, _ := os.Open("../../examples/apt-term-config.json")
	scanner := bufio.NewScanner(file)

	configFileBytes := make([]byte, 0)

	for scanner.Scan() {
		configFileBytes = append(configFileBytes, scanner.Bytes()...)
	}

	testTable := []struct {
		name           string
		server         *httptest.Server
		expectedOutput *Configuration
		expectedError  error
	}{
		{
			name: "good-test-create-web-config",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(configFileBytes)
			})),
			expectedOutput: &Configuration{},
			expectedError:  nil,
		},
		{
			name: "bad-test-create-web-config-404",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			})),
			expectedOutput: &Configuration{},
			expectedError:  fmt.Errorf("received status code 404 when attempting to get file"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()
			testWebSource, err := CreateConfiguration(test.server.URL)

			if reflect.TypeOf(testWebSource).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Create Configuration returned the wrong output type for web configuration: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(testWebSource).String())
			}

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Create Configuration returned an improper error when using incorrect web URL: \n\tExpected: %v\n\tReceived: %v", test.expectedError.Error(), err.Error())
			}
		})
	}
}

func TestResolveRule(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	logFile, _ := logs.LoadLogFile("../../examples/apt-term.log")
	currentConfig, _ := CreateConfiguration("../../examples/apt-term-config.json")

	testTable := []struct {
		name           string
		currentRule    *Rule
		expectedOutput *models.RuleData
		expectedError  error
	}{
		{
			name:           "Test_Good_Rule",
			currentRule:    &currentConfig.Rules[0],
			expectedOutput: &models.RuleData{},
			expectedError:  nil,
		},
		{
			name:           "Test_Bad_Rule",
			currentRule:    &Rule{},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("unable to resolve search terms for rule : \n\truntime error: index out of range [0] with length 0"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); !checkExpectedError(err, test.expectedError) {
					t.Errorf("The expected error was not returned: \n\tExpected: %v\n\tReceived: %v", test.expectedError.Error(), err.(error).Error())
				}
			}()

			ruleData, _ := ResolveRule(logFile, test.currentRule, logger)

			if reflect.TypeOf(ruleData).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("The incorrect datatype was not returned: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.currentRule).String(), reflect.TypeOf(ruleData).String())
			}
		})
	}
}
