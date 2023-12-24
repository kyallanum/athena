package utils

import (
	"fmt"

	models "github.com/kyallanum/athena/v0.1.0/models"
	config "github.com/kyallanum/athena/v0.1.0/models/config"
	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func ResolveRule(contents *models.LogFile, rule *config.Rule) (*library.RuleData, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("unable to resolve rule %s: \n\t%w", rule.Name, err)
	}

	allEntriesFound := false
	linesResolved := []int{}

	currentRuleData := library.RuleData.New(library.RuleData{}, rule.Name)

	for !allEntriesFound {
		currentSearchTermData, err := resolveSearchTerms(contents, rule, &linesResolved)
		if err != nil {
			return nil, wrap_error(err)
		}

		if len(currentSearchTermData.GetKeys()) != 0 {
			currentRuleData.AppendSearchTermData(currentSearchTermData)
		} else {
			allEntriesFound = true
		}
	}

	for _, summaryLine := range rule.Summary {
		summaryData, err := resolveSummaryLine(summaryLine, &currentRuleData)
		if err != nil {
			return nil, wrap_error(err)
		}
		for _, line := range summaryData {
			currentRuleData.AppendSummaryData(line)
		}
	}
	return &currentRuleData, nil
}
