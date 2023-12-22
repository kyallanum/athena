package utils

import (
	"fmt"
	"slices"

	models "github.com/kyallanum/athena/v0.1.0/models"
	config "github.com/kyallanum/athena/v0.1.0/models/config"
	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func resolveSearchTerms(logFile *models.LogFile, rule *config.Rule, linesResolved *[]int) (*library.SearchTermData, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("utils/Rule -> ResolveSearchTerms: \n\t%w", err)
	}

	currentSearchTermData := library.SearchTermData.New(library.SearchTermData{})
	currentSearchTerm := rule.SearchTerms[0]
	searchTermTranslated := false
	for fileIndex, searchTermIndex := 0, 0; fileIndex < logFile.GetLen() && searchTermIndex < len(rule.SearchTerms); fileIndex++ {
		if slices.Contains(*linesResolved, fileIndex) {
			continue
		}

		if !searchTermTranslated {
			currentSearchTerm = translateSearchTermReference(currentSearchTerm, currentSearchTermData)
			searchTermTranslated = true
		}

		currentLine, err := logFile.GetLineAtIndex(fileIndex)
		if err != nil {
			return nil, wrap_error(err)
		}

		result := resolveLine(currentLine, currentSearchTerm)
		if result == nil {
			continue
		}

		if rule.PrintLog {
			fmt.Printf("%d: %s\n", fileIndex+1, currentLine)
		}

		*linesResolved = append(*linesResolved, fileIndex)
		searchTermIndex++
		if searchTermIndex == len(rule.SearchTerms) {
			break
		}
		currentSearchTerm = rule.SearchTerms[searchTermIndex]
		searchTermTranslated = false

		for key, value := range *result {
			err := currentSearchTermData.AddValue(key, value)
			if err != nil {
				return nil, wrap_error(err)
			}
		}
	}

	return currentSearchTermData, nil
}
