package utils

import (
	"encoding/json"
	"fmt"

	"github.com/kyallanum/athena/v0.1.0/models"
)

func CreateConfiguration(source string) (config *models.Configuration, err error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("utils/ConfigFile -> CreateConfiguration: \n\t%w", err)
	}

	configSource, err := GetSource(source)
	if err != nil {
		return nil, wrap_error(err)
	}

	configFile, err := configSource.LoadConfig()
	if err != nil {
		return nil, wrap_error(err)
	}

	configObject := &models.Configuration{}
	err = json.Unmarshal(configFile, configObject)
	if err != nil {
		return nil, wrap_error(err)
	}

	return configObject, nil
}

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
