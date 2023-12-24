package utils

import (
	"fmt"
	"strings"

	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func GetOperation(operation string, key string) (library.ISummaryOperation, error) {
	switch strings.ToLower(strings.TrimSpace(operation)) {
	case "count":
		return library.Count.New(library.Count{}, key), nil
	case "print":
		return library.Print.New(library.Print{}, key), nil
	}

	return nil, fmt.Errorf("the given operation is not implemented: %s\n\t", operation)
}
