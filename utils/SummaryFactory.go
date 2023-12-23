package utils

import (
	"strings"

	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func GetOperation(operation string, key string) library.ISummaryOperation {
	switch strings.ToLower(strings.TrimSpace(operation)) {
	case "count":
		return library.Count.New(library.Count{}, key)
	case "print":
		return library.Print.New(library.Print{}, key)
	}

	return nil
}
