package models

import (
	"fmt"
	"strings"
)

type ISummaryOperation interface {
	CalculateOperation(ruleData RuleData) ([]string, error)
	Operation() string
	SetOperation(operation string)
	Key() string
	SetKey(key string)
}

type SummaryOperation struct {
	operation string
	key       string
}

func (summaryKey *SummaryOperation) Operation() string {
	return summaryKey.operation
}

func (summaryKey *SummaryOperation) SetOperation(operation string) {
	summaryKey.operation = operation
}

func (summaryKey *SummaryOperation) Key() string {
	return summaryKey.key
}

func (summaryKey *SummaryOperation) SetKey(key string) {
	summaryKey.key = key
}

func Operation(operation string, key string) (ISummaryOperation, error) {
	switch strings.ToLower(strings.TrimSpace(operation)) {
	case "count":
		return Count.New(Count{}, key), nil
	case "print":
		return Print.New(Print{}, key), nil
	}

	return nil, fmt.Errorf("the given operation is not implemented: %s\n\t", operation)
}
