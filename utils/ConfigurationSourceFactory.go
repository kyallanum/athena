package utils

import (
	"fmt"
	"os"

	urlverifier "github.com/davidmytton/url-verifier"
	"github.com/kyallanum/athena/v0.1.0/models"
)

func GetSource(source string) (models.IConfigurationSource, error) {
	isUrl, err := verifyUrl(source)
	if err != nil {
		return nil, err
	}

	if isUrl {
		return models.ConfigWebSource.New(models.ConfigWebSource{}, source), nil
	}

	isFile, err := verifyFilePath(source)
	if err != nil {
		return nil, err
	}

	if isFile {
		return models.ConfigFileSource.New(models.ConfigFileSource{}, source), nil
	}
}

func verifyUrl(source string) (bool, error) {
	verifier := urlverifier.NewVerifier()

	_, err := verifier.Verify(source)
	if err != nil {
		return false, nil
	}

	_, err = verifier.CheckHTTP(source)
	if err != nil {
		return false, fmt.Errorf("URL provided is not reachable. Please check the URL and try again.")
	}

	return true, nil
}

func verifyFilePath(source string) (bool, error) {
	fileInfo, err := os.Stat(source)
	if err != nil {
		return false, fmt.Errorf("Unable to get file information for file: %s. Error: %w", source, err)
	}

	if fileInfo.Mode().Perm()&0444 != 0444 {
		return false, fmt.Errorf("%s: File does not have read permissions.", source)
	}

	return true, nil
}
