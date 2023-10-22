package models

import "strconv"

type Count struct {
	SummaryOperation
}

func (count *Count) CalculateOperation(key string, ruleData RuleData) []string {
	searchTermLen := ruleData.GetSearchTermDataLen()

	searchTermCount := 0

	for index := 0; index < searchTermLen; index++ {
		currentSearchTermData := ruleData.GetSearchTermData(index)
		_, err := currentSearchTermData.GetValue(key)
		if err != nil {
			continue
		}
		searchTermCount++
	}

	ret_value := make([]string, 0)
	ret_value = append(ret_value, strconv.Itoa(searchTermCount))
	return ret_value
}

func (Count) New(key string) ISummaryOperation {
	return &Count{
		SummaryOperation: SummaryOperation{
			operation: "count",
			key:       key,
		},
	}
}
