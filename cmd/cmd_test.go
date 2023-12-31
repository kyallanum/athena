package cmd

import (
	"io"
	"os"
	"testing"

	config "github.com/kyallanum/athena/models/config"
	logs "github.com/kyallanum/athena/models/logs"
	"github.com/sirupsen/logrus"
)

func TestResolveFile(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()
	logFile, _ := logs.LoadLogFile("../examples/apt-term.log")
	configuration, _ := config.CreateConfiguration("../examples/apt-term-config.json")

	_, err := resolveLogFile(logFile, configuration, logger)

	if err != nil {
		t.Errorf("An error occurred while resolving log file: %s", err)
	}
}

func TestResolveLogFileBadLog(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("../examples/apt-term-bad.log")
	configuration, _ := config.CreateConfiguration("../examples/apt-term-config.json")

	_, err := resolveLogFile(logFile, configuration, nil)

	if err.Error() != "log file contains no contents" {
		t.Errorf("Error was not properly returned when checking log file contents")
	}
}

func TestResolveLogFileBadConfig(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := logs.LoadLogFile("../examples/apt-term.log")
	configuration, _ := config.CreateConfiguration("../examples/apt-term-config-bad.json")

	_, err := resolveLogFile(logFile, configuration, logger)

	if err.Error() != "configuration file has no contents" {
		t.Errorf("Error was not properly returned when checking configuration contents")
	}
}

func TestResolveLogFileNoConfigName(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	logFile, _ := logs.LoadLogFile("../examples/apt-term.log")
	configuration := &config.Configuration{
		Rules: make([]config.Rule, 1),
	}

	_, err := resolveLogFile(logFile, configuration, logger)

	if err.Error() != "configuration file contains no log name" {
		t.Errorf("Error was not properly returned when checking configuration name")
	}
}

func TestResolveLogFileNoConfigRules(t *testing.T) {
	logger := logrus.New()
	logger.SetOutput(io.Discard)

	logFile, _ := logs.LoadLogFile("../examples/apt-term.log")
	configuration := &config.Configuration{
		Name: "test",
	}

	_, err := resolveLogFile(logFile, configuration, logger)
	if err.Error() != "configuration does not have any rules" {
		t.Errorf("Error was not properly returned when checking configuration rules")
	}
}
