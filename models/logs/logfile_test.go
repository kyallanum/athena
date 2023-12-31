package models

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestGetLineAtIndex(t *testing.T) {
	currentFile, _ := os.Open("../../examples/apt-term.log")
	scanner := bufio.NewScanner(currentFile)

	logFile := make([]string, 0)

	for scanner.Scan() {
		currentText := scanner.Text()
		logFile = append(logFile, currentText)
	}

	logFileObject := New(logFile)

	line, err := logFileObject.LineAtIndex(0)
	if err != nil {
		t.Errorf("Error improperly returned when getting a proper line")
	}

	if reflect.TypeOf(line).String() != "string" {
		t.Errorf("GetLineAtIndex returned the wrong data type")
	}

	_, err = logFileObject.LineAtIndex(-1)
	if err == nil {
		t.Errorf("Error was not returned when attempting to get an incorrect index from log")
	}

	if err.Error() != "line at index: -1 does not exist in the logfile" {
		t.Errorf("Error was improperly returned when attempting to get an incorrect index from log file")
	}
}

func TestLoadLogFile(t *testing.T) {
	filename := "../../examples/apt-term.log"

	logFile, err := LoadLogFile(filename)

	if err != nil {
		t.Errorf("An error was returned from LoadLogFile when one shouldn't have: \n\t%s", err.Error())
	}

	if reflect.TypeOf(logFile).String() != "*models.LogFile" {
		t.Errorf("LoadLogFile returned the wrong data type: \n\t%s", reflect.TypeOf(logFile).String())
	}
}

func TestLoadLogFileBadFile(t *testing.T) {
	filename := "../../examples/apt-term-bad.log"

	_, err := LoadLogFile(filename)

	if err == nil {
		t.Errorf("An error was not returned when attempting to Load a nonexistant log file.")
	}

	if err.Error() != "unable to load log from file: \n\topen ../../examples/apt-term-bad.log: no such file or directory" {
		t.Errorf("An improper error was returned when attempting to load a nonexistant log file: \n\t%s", err.Error())
	}
}
