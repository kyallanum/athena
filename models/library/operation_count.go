package models

import (
	"fmt"
	"strconv"
)

type Count struct {
	SummaryOperation
}

func (count *Count) CalculateOperation(ruleData RuleData) ([]string, error) {
	searchTermLen := ruleData.SearchTermDataLen()

	if searchTermLen == 0 {
		return nil, fmt.Errorf("rule does not have any search term data stored")
	}

	searchTermCount := 0

	for index := 0; index < searchTermLen; index++ {
		currentSearchTermData := ruleData.SearchTermData(index)
		_, err := currentSearchTermData.Value(count.key)
		if err != nil {
			continue
		}
		searchTermCount++
	}

	if searchTermCount == 0 {
		return nil, fmt.Errorf("this value was never extracted during search phase")
	}

	ret_value := make([]string, 0)
	ret_value = append(ret_value, strconv.Itoa(searchTermCount))
	return ret_value, nil
}

func NewCountOperation(key string) ISummaryOperation {
	return &Count{
		SummaryOperation: SummaryOperation{
			operation: "count",
			key:       key,
		},
	}
}
