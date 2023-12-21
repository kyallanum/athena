package models

import (
	"fmt"
	"io"
	"net/http"
)

type ConfigWebSource struct {
	ConfigurationSource
}

func (config *ConfigWebSource) LoadConfig() ([]byte, error) {
	url := config.source

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ConfigWebSource -> LoadConfig: \n\t%w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("ConfigWebSource -> LoadConfig: \n\tReceived Status Code %d when attempting to get file: %s", response.StatusCode, url)
	}

	data, _ := io.ReadAll(response.Body)
	return data, nil
}

func (ConfigWebSource) New(source string) IConfigurationSource {
	return &ConfigWebSource{
		ConfigurationSource: ConfigurationSource{
			source_type: "web",
			source:      source,
		},
	}
}
