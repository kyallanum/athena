package models

type IConfigurationSource interface {
	setSourceType(sourceType string)
	setSource(source string)
	getSourceType() string
	getSource() string
}

type ConfigurationSource struct {
	source_type string
	source      string
}

func (configSource *ConfigurationSource) getSourceType() string {
	return configSource.source_type
}

func (configSource *ConfigurationSource) getSource() string {
	return configSource.source
}

func (configSource *ConfigurationSource) setSourceType(sourceType string) {
	configSource.source_type = sourceType
}

func (configSource *ConfigurationSource) setSource(source string) {
	configSource.source = source
}
