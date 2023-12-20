package models

type ConfigWebSource struct {
	ConfigurationSource
}

func (ConfigWebSource) New(source string) IConfigurationSource {
	return &ConfigWebSource{
		ConfigurationSource: ConfigurationSource{
			source_type: "web",
			source:      source,
		},
	}
}
