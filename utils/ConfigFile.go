package utils

import (
	"encoding/json"
	"fmt"

	models "github.com/kyallanum/athena/v0.1.0/models/config"
)

func CreateConfiguration(source string) (config *models.Configuration, err error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("utils/ConfigFile -> CreateConfiguration: \n\t%w", err)
	}

	configSource, err := getSource(source)
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

	configObject.TranslateRegexGroups()

	return configObject, nil
}
