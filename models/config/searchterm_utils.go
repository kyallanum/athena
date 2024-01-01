package models

import (
	"fmt"
	"slices"

	library "github.com/kyallanum/athena/models/library"
	logs "github.com/kyallanum/athena/models/logs"
	"github.com/sirupsen/logrus"
)

func resolveSearchTerms(logFile *logs.LogFile, rule *Rule, linesResolved *[]int, logger *logrus.Logger) (*library.SearchTermData, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to resolve search terms for rule %s: \n\t%w", rule.Name, err)
	}

	defer func() {
		if err := recover(); err != nil {
			panic(wrapError(fmt.Errorf("%s", err)))
		}
	}()

	currentSearchTermData := library.NewSearchTermData()
	currentSearchTerm := rule.SearchTerms[0]
	searchTermTranslated := false
	for fileIndex, searchTermIndex := 0, 0; fileIndex < logFile.Len() && searchTermIndex < len(rule.SearchTerms); fileIndex++ {
		if slices.Contains(*linesResolved, fileIndex) {
			continue
		}

		if !searchTermTranslated {
			newSearchTerm, err := translateSearchTermReference(currentSearchTerm, currentSearchTermData)
			if err != nil {
				return nil, wrapError(err)
			}
			currentSearchTerm = newSearchTerm
			searchTermTranslated = true
		}

		currentLine, err := logFile.LineAtIndex(fileIndex)
		if err != nil {
			return nil, wrapError(err)
		}

		result := resolveLine(currentLine, currentSearchTerm)
		if result == nil {
			continue
		}

		if rule.PrintLog {
			logger.Infof("%d: %s", fileIndex+1, currentLine)
		}

		*linesResolved = append(*linesResolved, fileIndex)

		for key, value := range *result {
			err := currentSearchTermData.AddValue(key, value)
			if err != nil {
				return nil, wrapError(err)
			}
		}

		searchTermIndex++
		if searchTermIndex == len(rule.SearchTerms) {
			break
		}
		currentSearchTerm = rule.SearchTerms[searchTermIndex]
		searchTermTranslated = false
	}

	return currentSearchTermData, nil
}
