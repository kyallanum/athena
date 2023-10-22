package models

type ISummaryOperation interface {
	CalculateOperation(key string, ruleData RuleData) []string
	GetOperation() string
	SetOperation(operation string)
	GetKey() string
	SetKey(key string)
}

type SummaryOperation struct {
	operation string
	key       string
}

func (summaryKey *SummaryOperation) GetOperation() string {
	return summaryKey.operation
}

func (summaryKey *SummaryOperation) SetOperation(operation string) {
	summaryKey.operation = operation
}

func (summaryKey *SummaryOperation) GetKey() string {
	return summaryKey.key
}

func (summaryKey *SummaryOperation) SetKey(key string) {
	summaryKey.key = key
}
