package models

import (
	"fmt"

	library "github.com/kyallanum/athena/models/library"
	logs "github.com/kyallanum/athena/models/logs"
	"github.com/sirupsen/logrus"
)

func ResolveRule(contents *logs.LogFile, rule *Rule, logger *logrus.Logger) (*library.RuleData, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to resolve rule %s: \n\t%w", rule.Name, err)
	}

	allEntriesFound := false
	linesResolved := []int{}

	currentRuleData := library.NewRuleData()

	for !allEntriesFound {
		currentSearchTermData, err := resolveSearchTerms(contents, rule, &linesResolved, logger)
		if err != nil {
			return nil, wrapError(err)
		}

		if len(currentSearchTermData.Keys()) != 0 {
			currentRuleData.AppendSearchTermData(currentSearchTermData)
		} else {
			allEntriesFound = true
		}
	}

	for _, summaryLine := range rule.Summary {
		summaryData, err := resolveSummaryLine(summaryLine, &currentRuleData)
		if err != nil {
			return nil, wrapError(err)
		}
		for _, line := range summaryData {
			currentRuleData.AppendSummaryData(line)
		}
	}
	return &currentRuleData, nil
}
