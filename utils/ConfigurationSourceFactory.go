package utils

import (
	"fmt"
	"os"

	urlverifier "github.com/davidmytton/url-verifier"
	"github.com/kyallanum/athena/v0.1.0/models"
)

func getSource(source string) (models.IConfigurationSource, error) {
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

	//Should never reach this
	return nil, nil
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
