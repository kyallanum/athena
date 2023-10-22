package models

type Print struct {
	SummaryOperation
}

func (print *Print) CalculateOperation(key string, ruleData RuleData) []string {
	ret_value := make([]string, 0)

	st_data_len := ruleData.GetSearchTermDataLen()
	for i := 0; i < st_data_len; i++ {
		currentRuleData := ruleData.GetSearchTermData(i)
		searchTermData, _ := currentRuleData.GetValue(key)
		ret_value = append(ret_value, searchTermData)
	}
	return ret_value
}

func (Print) New(key string) ISummaryOperation {
	return &Print{
		SummaryOperation: SummaryOperation{
			operation: "print",
			key:       key,
		},
	}
}
