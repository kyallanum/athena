package models

import (
	"fmt"
	"os"

	urlverifier "github.com/davidmytton/url-verifier"
)

type IConfigurationSource interface {
	SetSourceType(sourceType string)
	SetSource(source string)
	SourceType() string
	Source() string
	Config() ([]byte, error)
}

type ConfigurationSource struct {
	source_type string
	source      string
}

func (configSource *ConfigurationSource) SourceType() string {
	return configSource.source_type
}

func (configSource *ConfigurationSource) Source() string {
	return configSource.source
}

func (configSource *ConfigurationSource) SetSourceType(sourceType string) {
	configSource.source_type = sourceType
}

func (configSource *ConfigurationSource) SetSource(source string) {
	configSource.source = source
}

func verifyFilePath(source string) (bool, error) {
	fileInfo, err := os.Stat(source)
	if err != nil {
		return false, fmt.Errorf("unable to get file information for file: %s. error: %w", source, err)
	}

	if fileInfo.Mode().Perm()&0444 != 0444 {
		return false, fmt.Errorf("%s: file does not have read permissions", source)
	}

	return true, nil
}

func verifyUrl(source string) (bool, error) {
	verifier := urlverifier.NewVerifier()

	verifiedUrl, err := verifier.Verify(source)
	if err != nil {
		return false, fmt.Errorf("url provided could not be verified: \n\t%w", err)
	}
	if !verifiedUrl.IsURL {
		return false, nil
	}

	_, err = verifier.CheckHTTP(source)
	if err != nil {
		return false, fmt.Errorf("url provided is not reachable. please check the URL and try again")
	}

	return true, nil
}

func Source(source string) (IConfigurationSource, error) {
	isUrl, err := verifyUrl(source)
	if err != nil {
		return nil, err
	}

	if isUrl {
		return ConfigWebSource.New(ConfigWebSource{}, source), nil
	}

	isFile, err := verifyFilePath(source)
	if err != nil {
		return nil, err
	}

	if isFile {
		return ConfigFileSource.New(ConfigFileSource{}, source), nil
	}

	//Should never reach this
	return nil, nil
}
