package models

import (
	"bufio"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	logs "github.com/kyallanum/athena/models/logs"
)

func TestTranslateRegex(t *testing.T) {
	regexToTranslate := "(?<test_name>)"
	err := translateConfigurationNamedGroups(&regexToTranslate)

	if err != nil {
		t.Errorf("An error occurred while attempting to translate regex: \n\t%v", err)
	}

	if regexToTranslate != "(?P<test_name>)" {
		t.Errorf("Regex was not translated properly.")
	}
}

func TestTranslateRegexEmptyString(t *testing.T) {
	regexToTranslate := ""
	err := translateConfigurationNamedGroups(&regexToTranslate)

	if err.Error() != "empty search terms are not allowed" {
		t.Errorf("Error was not properly returned when checking for empty string.")
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

func TestCreateConfigurationFromFile(t *testing.T) {
	config, err := CreateConfiguration("../../examples/apt-term-config.json")
	if err != nil {
		t.Errorf("Error returned during CreateConfiguration when there shouldn't have.")
	}

	if config.Name != "Apt Terminal" {
		t.Errorf("Name improperly returned from CreateConfiguration")
	}

	if reflect.TypeOf(config).String() != "*models.Configuration" {
		t.Errorf("Improper type returned when running CreateConfiguration")
	}
}

func TestCreateConfigurationFromFileBadFile(t *testing.T) {
	_, err := CreateConfiguration("bad_file_name")
	if err == nil {
		t.Errorf("Error not returned when it should have after calling CreateConfiguration")
	}

	if err.Error() != "unable to create configuration object: \n\tunable to get file information for file: bad_file_name. error: stat bad_file_name: no such file or directory" {
		t.Errorf("Error improperly returned when it should have. \nError: %s", err.Error())
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
		name               string
		server             *httptest.Server
		expectedOutputType string
		expectedErr        string
	}{
		{
			name: "good-test-create-web-config",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(configFileBytes)
			})),
			expectedOutputType: "*models.Configuration",
			expectedErr:        "",
		},
		{
			name: "bad-test-create-web-config-404",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			})),
			expectedOutputType: "*models.Configuration",
			expectedErr:        "received status code 404 when attempting to get file",
		},
	}

	for _, testServer := range testTable {
		t.Run(testServer.name, func(t *testing.T) {
			defer testServer.server.Close()
			testWebSource, err := CreateConfiguration(testServer.server.URL)

			if reflect.TypeOf(testWebSource).String() != testServer.expectedOutputType {
				t.Errorf("Create Configuration returned the wrong output type for web configuration")
			}

			if (err != nil) && (testServer.expectedErr != "") {
				if !strings.Contains(err.Error(), testServer.expectedErr) {
					t.Errorf("Create Configuration returned an improper error when using incorrect web URL")
				}
			}

		})
	}
}

func TestResolveRule(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("../../examples/apt-term.log")

	currentConfig, _ := CreateConfiguration("../../examples/apt-term-config.json")

	currentRule := currentConfig.Rules[0]

	ruleData, err := ResolveRule(logFile, &currentRule)
	if err != nil {
		t.Errorf("An error was returned when one should not have been: \n\t%s", err.Error())
	}

	if reflect.TypeOf(ruleData).String() != "*models.RuleData" {
		t.Errorf("The incorrect datatype was not returned: \n\t%s", reflect.TypeOf(ruleData).String())
	}
}

func TestResolveRuleBadRule(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("An error was not returned when one should have been.")
		} else if (err.(error)).Error() != "unable to resolve search terms for rule : \n\truntime error: index out of range [0] with length 0" {
			t.Errorf("%s", err.(error).Error())
		}
	}()
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("../examples/apt-term.log")

	currentRule := &Rule{}

	ResolveRule(logFile, currentRule)
}
