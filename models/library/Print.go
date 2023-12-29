package models

import "fmt"

type Print struct {
	SummaryOperation
}

func (print *Print) CalculateOperation(ruleData RuleData) ([]string, error) {
	ret_value := make([]string, 0)

	st_data_len := ruleData.GetSearchTermDataLen()

	if st_data_len == 0 {
		return nil, fmt.Errorf("rule does not have any search term data stored")
	}

	for i := 0; i < st_data_len; i++ {
		currentRuleData := ruleData.GetSearchTermData(i)
		searchTermData, err := currentRuleData.GetValue(print.key)
		if err != nil {
			continue
		}
		ret_value = append(ret_value, searchTermData)
	}

	if len(ret_value) == 0 {
		return nil, fmt.Errorf("this value was never extracted during search phase")
	}
	return ret_value, nil
}

func (Print) New(key string) ISummaryOperation {
	return &Print{
		SummaryOperation: SummaryOperation{
			operation: "print",
			key:       key,
		},
	}
}
