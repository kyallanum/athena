package models

type IConfigurationSource interface {
	SetSourceType(sourceType string)
	SetSource(source string)
	GetSourceType() string
	GetSource() string
	LoadConfig() ([]byte, error)
}

type ConfigurationSource struct {
	source_type string
	source      string
}

func (configSource *ConfigurationSource) GetSourceType() string {
	return configSource.source_type
}

func (configSource *ConfigurationSource) GetSource() string {
	return configSource.source
}

func (configSource *ConfigurationSource) SetSourceType(sourceType string) {
	configSource.source_type = sourceType
}

func (configSource *ConfigurationSource) SetSource(source string) {
	configSource.source = source
}
