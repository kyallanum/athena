package utils

import (
	"fmt"

	"github.com/kyallanum/athena/v0.1.0/models"
)

func ResolveRule(contents *models.LogFile, rule *models.Rule) (*models.RuleData, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("utils/Rule -> ResolveRule: \n\t%w", err)
	}

	allEntriesFound := false
	linesResolved := []int{}

	currentRuleData := models.RuleData.New(models.RuleData{})

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
