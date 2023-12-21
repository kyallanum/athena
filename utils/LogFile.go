package utils

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kyallanum/athena/v0.1.0/models"
)

func LoadLogFile(fileName string) (*models.LogFile, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("LoadInputFile -> Open File: \n\t%w", err)
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, fmt.Errorf("LoadInputFile -> Read File: \n\t%w", scanner.Err())
	}

	logFileContents := models.LogFile.New(models.LogFile{}, lines)
	return logFileContents, nil
}
