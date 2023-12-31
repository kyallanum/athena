package models

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
