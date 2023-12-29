package utils

import (
	"reflect"
	"testing"
)

func TestLoadLogFile(t *testing.T) {
	filename := "../examples/apt-term.log"

	logFile, err := LoadLogFile(filename)

	if err != nil {
		t.Errorf("An error was returned from LoadLogFile when one shouldn't have: \n\t%s", err.Error())
	}

	if reflect.TypeOf(logFile).String() != "*models.LogFile" {
		t.Errorf("LoadLogFile returned the wrong data type: \n\t%s", reflect.TypeOf(logFile).String())
	}
}

func TestLoadLogFileBadFile(t *testing.T) {
	filename := "../examples/apt-term-bad.log"

	_, err := LoadLogFile(filename)

	if err == nil {
		t.Errorf("An error was not returned when attempting to Load a nonexistant log file.")
	}

	if err.Error() != "unable to load log from file: \n\topen ../examples/apt-term-bad.log: no such file or directory" {
		t.Errorf("An improper error was returned when attempting to load a nonexistant log file: \n\t%s", err.Error())
	}
}
