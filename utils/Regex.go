package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsRegexInLine(line string, regex string) bool {
	re := regexp.MustCompile(regex)
	return re.FindAllString(line, -1) != nil
}

func ResolveRegexpNames(line string, regex string) *map[string]string {
	currentRegexp := regexp.MustCompile(regex)
	match := currentRegexp.FindStringSubmatch(line)
	result := make(map[string]string)

	if len(match) > 0 {
		for index, name := range currentRegexp.SubexpNames() {
			if index != 0 && name != "" {
				result[name] = match[index]
			}
		}
		return &result
	}

	return nil
}

func TranslateRegex(regex *string) {
	regexAddGolangGroupName := `(\(\?)(\<[\w\W]+?\>)`
	compiledRegex := regexp.MustCompile(regexAddGolangGroupName)

	*regex = compiledRegex.ReplaceAllString(*regex, "${1}P${2}")
}

func TranslateNames(regex string, names map[string](map[string]string)) string {
	// Matches the pattern: {{RuleName[ReplacedName]}}
	nameExtractRegex := `({\{(?P<RuleName>[\w]+?)\[(?P<ReplacedName>[\w]+?)\]\}\})`
	re := regexp.MustCompile(nameExtractRegex)
	matches := re.FindAllStringSubmatch(regex, -1)
	numMatches := len(matches)

	myMap := make(map[string][]string)

	// The first three groups are useless for this case. So get everything after that.
	for index, match := range matches {
		myMap["value"+strconv.Itoa(index+1)] = match[2:]
	}

	// Validate all strings to replace and then replace them one by one.
	for i := 0; i < numMatches; i++ {
		stringToReplace := names[myMap["value"+strconv.Itoa(i+1)][0]][myMap["value"+strconv.Itoa(i+1)][1]]
		stringToReplace = validateStringToReplace(stringToReplace)
		foundString := re.FindString(regex)
		if foundString != "" {
			regex = strings.Replace(regex, foundString, stringToReplace, 1)
		}
	}

	// fmt.Println(regex)
	return regex
}

func TranslateSummaryLine(summaryString string, library map[string](map[string][]string)) {
	extractionRegex := `(\{\{[\w\[\]\(\)]+?\}\})`
	re := regexp.MustCompile(extractionRegex)
	matches := re.FindAllString(summaryString, -1)

	for _, reference := range matches {
		translationRegex := `\{\{(Count|)(\(|)([\w]+)\[([\w]+)\](\)|)\}\}`
		re = regexp.MustCompile(translationRegex)
		translationMatches := re.FindAllStringSubmatch(reference, -1)
		fmt.Println(translationMatches)
	}
}

func validateStringToReplace(regex string) string {
	charactersToEscape := `[\*\+\?\\\.\^\[\]\$\&\|]{1}`
	re := regexp.MustCompile(charactersToEscape)

	regex = re.ReplaceAllString(regex, `\\$1`)
	return regex
}
