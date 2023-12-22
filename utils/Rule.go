package utils

import (
	"fmt"

	models "github.com/kyallanum/athena/v0.1.0/models"
	config "github.com/kyallanum/athena/v0.1.0/models/config"
	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func ResolveRule(contents *models.LogFile, rule *config.Rule) (*library.RuleData, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("utils/Rule -> ResolveRule: \n\t%w", err)
	}

	allEntriesFound := false
	linesResolved := []int{}

	currentRuleData := library.RuleData.New(library.RuleData{})

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
	return &currentRuleData, nil
}
