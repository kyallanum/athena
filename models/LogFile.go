package models

import "fmt"

// This struct purely exists to make the logfile contents readonly after initialization.
type LogFile struct {
	contents []string
	length   int
}

func (logfile *LogFile) GetLen() int {
	return logfile.length
}

func (logfile *LogFile) GetLineAtIndex(index int) (string, error) {
	if index < 0 || index >= logfile.length {
		return "", fmt.Errorf("line at index: %d does not exist in the logfile", index)
	}

	return logfile.contents[index], nil
}

func (LogFile) New(contents []string) *LogFile {
	return &LogFile{
		contents: contents,
		length:   len(contents),
	}
}
