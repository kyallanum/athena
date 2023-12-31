package models

import (
	"bufio"
	"fmt"
	"os"
)

// This struct purely exists to make the logfile contents readonly after initialization.
type LogFile struct {
	contents []string
	length   int
}

func (logfile *LogFile) Len() int {
	return logfile.length
}

func (logfile *LogFile) LineAtIndex(index int) (string, error) {
	if index < 0 || index >= logfile.length {
		return "", fmt.Errorf("line at index: %d does not exist in the logfile", index)
	}

	return logfile.contents[index], nil
}

func New(contents []string) *LogFile {
	return &LogFile{
		contents: contents,
		length:   len(contents),
	}
}

func LoadLogFile(fileName string) (*LogFile, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to load log from file: \n\t%w", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, wrapError(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, wrapError(err)
	}

	logFileContents := New(lines)
	return logFileContents, nil
}
