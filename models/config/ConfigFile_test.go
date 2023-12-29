package models

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
)

func TestTranslateRegex(t *testing.T) {
	regexToTranslate := "(?<test_name>)"
	err := translateRegex(&regexToTranslate)

	if err != nil {
		t.Errorf("An error occurred while attempting to translate regex: \n\t%v", err)
	}

	if regexToTranslate != "(?P<test_name>)" {
		t.Errorf("Regex was not translated properly.")
	}
}

func TestTranslateRegexEmptyString(t *testing.T) {
	regexToTranslate := ""
	err := translateRegex(&regexToTranslate)

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
