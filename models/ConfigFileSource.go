package models

type ConfigFileSource struct {
	ConfigurationSource
}

func (ConfigFileSource) New(source string) IConfigurationSource {
	return &ConfigFileSource{
		ConfigurationSource: ConfigurationSource{
			source_type: "file",
			source:      source,
		},
	}
}
