package utils

import (
	"bufio"
	"io"
	"os"
)

func LoadFileBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	bytes := make([]byte, stat.Size())
	_, err = bufio.NewReader(file).Read(bytes)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return bytes, nil
}

func LoadFileStringSlice(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
