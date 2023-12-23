package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kyallanum/athena/v0.1.0/models"
)

func LoadLogFile(fileName string) (*models.LogFile, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("LogFile -> LoadLogFile: \n\t%w", err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, wrap_error(err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, wrap_error(err)
	}

	logFileContents := models.LogFile.New(models.LogFile{}, lines)
	return logFileContents, nil
}
