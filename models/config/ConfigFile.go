package models

import "regexp"

// A Configuration struct represents the top level of a JSON configuration file.
// It has two elements, Name (The type of file to be resolved), and Rules.
type Configuration struct {
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}

// A Rule Struct represents different pieces of data that need to be found in
// a log file.
// "Name" should describe what the utility should be looking for.
// "PrintLog" should describe whether a line should be printed out
// after it resolved to a search term in the rule.
//
// "SearchTerms" a list of regular expressions that the rule should look for.
// SearchTerms can save information resolved from a regex using a named group (?<group_name>).
// SearchTerms can also reference data saved from named groups with {{rule_name[group_name]}}
//
// The summary is a list of strings that should be printed out at the end of log file
// analysis. These can implement the same naming conventions to reference data as SearchTerms.
// They can also implement their own data manipulation functions such as {{Count(rule_name[group_name])}}
// to count the number of instances of <group_name> in that rule.
type Rule struct {
	Name        string   `json:"name"`
	PrintLog    bool     `json:"printLog"`
	SearchTerms []string `json:"searchTerms"`
	Summary     []string `json:"summary"`
}

// Translates any regex groups with names from (?<name>) syntax to (?P<name>) for
// Processing in Go
func (config *Configuration) TranslateRegexGroups() {
	for ruleIndex, currentRule := range config.Rules {
		for searchTermIndex, currentSearchTerm := range currentRule.SearchTerms {
			translateRegex(&currentSearchTerm)
			config.Rules[ruleIndex].SearchTerms[searchTermIndex] = currentSearchTerm
		}
	}
}

func translateRegex(regex *string) {
	regexAddGolangGroupName := `(\(\?)(\<[\w\W]+?\>)`
	compiledRegex := regexp.MustCompile(regexAddGolangGroupName)

	*regex = compiledRegex.ReplaceAllString(*regex, "${1}P${2}")
}
