package models

type Configuration struct {
	Name  string `json:"name"`
	Rules []Rule `json:"rules"`
}

type Rule struct {
	Name        string   `json:"name"`
	PrintLog    bool     `json:"printLog"`
	SearchTerms []string `json:"searchTerms"`
	Summary     []string `json:"summary"`
}
