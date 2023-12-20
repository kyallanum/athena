package utils

import (
	"github.com/kyallanum/athena/v0.1.0/models"
)

func CreateConfiguration() (config *models.Configuration, err error) {

}

// func LoadConfigurationFile(fileName string) ([]byte, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, fmt.Errorf("LoadConfigurationFile -> File Open: \n\t%w", err)
// 	}

// 	defer file.Close()

// 	stat, err := file.Stat()
// 	if err != nil {
// 		return nil, fmt.Errorf("LoadConfigurationFile -> GetFileSize: \n\t%w", err)
// 	}

// 	bytes := make([]byte, stat.Size())
// 	_, err = bufio.NewReader(file).Read(bytes)
// 	if err != nil && err != io.EOF {
// 		return nil, fmt.Errorf("LoadConfigurationFile -> Read File: \n\t%w", err)
// 	}
// 	return bytes, nil
// }

// func LoadInputFile(fileName string) ([]string, error) {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return nil, fmt.Errorf("LoadInputFile -> Open File: \n\t%w", err)
// 	}
// 	defer file.Close()

// 	lines := make([]string, 0)
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if scanner.Err() != nil {
// 		return nil, fmt.Errorf("LoadInputFile -> Read File: \n\t%w", scanner.Err())
// 	}
// 	return lines, nil
// }
