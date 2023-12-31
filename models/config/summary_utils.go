package models

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	library "github.com/kyallanum/athena/models/library"
)

func summaryKeys(summaryLine string) [][]string {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("could not get summary keys. this is most likely an internal error: \n\t%s", err.(string)))
		}
	}()

	ret_keys := make([][]string, 0)
	keyRegex := `\{\{(?P<value_1>[\w]+)(\()(?P<value_2>[\w]+)(\))\}\}`
	re := regexp.MustCompile(keyRegex)
	matches := re.FindAllStringSubmatch(summaryLine, -1)

	for _, match := range matches {
		original := match[0]
		operation := match[1]
		key := match[3]
		keyToAdd := []string{original, operation, key}
		ret_keys = append(ret_keys, keyToAdd)
	}

	return ret_keys
}

func resolveSummaryLine(summaryLine string, ruleData *library.RuleData) ([]string, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to resolve summary line: \n\t%w", err)
	}

	keys := summaryKeys(summaryLine)
	if !isOneUniqueOperation(keys) {
		return nil, fmt.Errorf("could not resolve summary. mixing operations is currently not implemented")
	}

	ret_summary_line := make([]string, 0)
	expanded := false

	for _, key := range keys {
		operation, err := library.Operation(key[1], key[2])
		if err != nil {
			return nil, wrapError(err)
		}

		calculated, err := operation.CalculateOperation(*ruleData)
		if err != nil {
			return nil, wrapError(err)
		}

		for i := 0; i < len(calculated); i++ {
			// expand the first time
			if !expanded {
				ret_summary_line = append(ret_summary_line, summaryLine)
			}
			ret_summary_line[i] = strings.Replace(ret_summary_line[i], key[0], calculated[i], 1)
		}
		expanded = true
	}
	return ret_summary_line, nil
}

func isOneUniqueOperation(keys [][]string) bool {
	currentKeys := make([]string, 0)
	for _, key := range keys {
		currentKeys = append(currentKeys, key[1])
	}

	ret_unique := slices.CompactFunc(currentKeys, func(i, j string) bool {
		return strings.EqualFold(i, j)
	})

	return len(ret_unique) == 1
}
