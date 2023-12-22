package models

import (
	"bufio"
	"fmt"
	"os"
)

type ConfigFileSource struct {
	ConfigurationSource
}

func (config *ConfigFileSource) LoadConfig() ([]byte, error) {
	file, err := os.Open(config.source)
	if err != nil {
		return nil, fmt.Errorf("ConfigFileSource -> LoadConfig: \n\t%w", err)
	}

	defer file.Close()

	bytes := make([]byte, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		bytes = append(bytes, scanner.Bytes()...)
	}

	return bytes, nil
}

func (ConfigFileSource) New(source string) IConfigurationSource {
	return &ConfigFileSource{
		ConfigurationSource: ConfigurationSource{
			source_type: "file",
			source:      source,
		},
	}
}
