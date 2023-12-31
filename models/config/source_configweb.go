package models

import (
	"fmt"
	"io"
	"net/http"
)

type ConfigWebSource struct {
	ConfigurationSource
}

func (config *ConfigWebSource) Config() ([]byte, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to create configuration for web source: \n\t%w", err)
	}
	url := config.source

	response, err := http.Get(url)
	if err != nil {
		return nil, wrapError(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, wrapError(fmt.Errorf("received status code %d when attempting to get file: %s", response.StatusCode, url))
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
