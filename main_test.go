package main

import (
	"os"
	"testing"

	config "github.com/kyallanum/athena/v1.0.0/models/config"
	logs "github.com/kyallanum/athena/v1.0.0/models/logs"
)

func TestResolveFile(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()
	logFile, _ := logs.LoadLogFile("./examples/apt-term.log")
	configuration, _ := config.CreateConfiguration("./examples/apt-term-config.json")

	_, err := resolveLogFile(logFile, configuration)

	if err != nil {
		t.Errorf("An error occurred while resolving log file: %s", err)
	}
}

func TestResolveLogFileBadLog(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("./examples/apt-term-bad.log")
	configuration, _ := config.CreateConfiguration("./examples/apt-term-config.json")

	_, err := resolveLogFile(logFile, configuration)

	if err.Error() != "log file contains no contents" {
		t.Errorf("Error was not properly returned when checking log file contents")
	}
}

func TestResolveLogFileBadConfig(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("./examples/apt-term.log")
	configuration, _ := config.CreateConfiguration("./examples/apt-term-config-bad.json")

	_, err := resolveLogFile(logFile, configuration)

	if err.Error() != "configuration file has no contents" {
		t.Errorf("Error was not properly returned when checking configuration contents")
	}
}

func TestResolveLogFileNoConfigName(t *testing.T) {
	logFile, _ := logs.LoadLogFile("./examples/apt-term.log")
	configuration := &config.Configuration{
		Rules: make([]config.Rule, 1),
	}

	_, err := resolveLogFile(logFile, configuration)

	if err.Error() != "configuration file contains no log name" {
		t.Errorf("Error was not properly returned when checking configuration name")
	}
}

func TestResolveLogFileNoConfigRules(t *testing.T) {
	logFile, _ := logs.LoadLogFile("./examples/apt-term.log")
	configuration := &config.Configuration{
		Name: "test",
	}

	_, err := resolveLogFile(logFile, configuration)
	if err.Error() != "configuration does not have any rules" {
		t.Errorf("Error was not properly returned when checking configuration rules")
	}
}
