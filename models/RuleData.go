package models

// The RuleData struct has a name which corresponds with the configuration file
// and a collection of search term data that corresponds to each instance where a
// search term is resolved, with named groups that should be accessed later.
type RuleData struct {
	st_data_collection []SearchTermData
	summary_data       []string
}

func (RuleData) New() RuleData {
	return RuleData{
		st_data_collection: []SearchTermData{},
	}
}

func (rule_data *RuleData) AppendSearchTermData(stdata SearchTermData) {
	rule_data.st_data_collection = append(rule_data.st_data_collection, stdata)
}

func (rule_data *RuleData) AppendSummaryData(summary_line string) {
	rule_data.summary_data = append(rule_data.summary_data, summary_line)
}
