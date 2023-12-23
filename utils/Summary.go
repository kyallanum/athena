package utils

import (
	"regexp"
	"strings"

	library "github.com/kyallanum/athena/v0.1.0/models/library"
)

func resolveSummaryLine(summaryLine string, ruleData *library.RuleData) []string {
	keys := getSummaryKeys(summaryLine)
	ret_summary_line := make([]string, 0)
	expanded := false

	for _, key := range keys {
		operation := GetOperation(key[1], key[2])
		calculated := operation.CalculateOperation(key[2], *ruleData)
		for i := 0; i < len(calculated); i++ {
			// expand the first time
			if !expanded {
				ret_summary_line = append(ret_summary_line, summaryLine)
			}
			ret_summary_line[i] = strings.Replace(ret_summary_line[i], key[0], calculated[i], 1)
		}
		expanded = true
	}
	return ret_summary_line
}

func getSummaryKeys(summaryLine string) [][]string {
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
