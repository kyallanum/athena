package models

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func TestGetLineAtIndex(t *testing.T) {
	currentFile, _ := os.Open("../examples/apt-term.log")
	scanner := bufio.NewScanner(currentFile)

	logFile := make([]string, 0)

	for scanner.Scan() {
		currentText := scanner.Text()
		logFile = append(logFile, currentText)
	}

	logFileObject := LogFile.New(LogFile{}, logFile)

	line, err := logFileObject.GetLineAtIndex(0)
	if err != nil {
		t.Errorf("Error improperly returned when getting a proper line")
	}

	if reflect.TypeOf(line).String() != "string" {
		t.Errorf("GetLineAtIndex returned the wrong data type")
	}

	_, err = logFileObject.GetLineAtIndex(-1)
	if err == nil {
		t.Errorf("Error was not returned when attempting to get an incorrect index from log")
	}

	if err.Error() != "line at index: -1 does not exist in the logfile" {
		t.Errorf("Error was improperly returned when attempting to get an incorrect index from log file")
	}
}
