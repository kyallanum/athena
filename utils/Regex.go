package utils

import (
	"fmt"
	"regexp"
	"strings"

	library "github.com/kyallanum/athena/v1.0.0/models/library"
)

// func resolveLine determines whether the current log line matches a
// matches a given log line
func resolveLine(line string, regex string) *map[string]string {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("the provided regular expression cannot be compiled: \n\t%s", err.(string)))
		}
	}()

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

// func translateSearchTermReference finds all references of {{key}} and replaces it with a key
// in the current search term data.
func translateSearchTermReference(regex string, currentSearchTermData *library.SearchTermData) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("the search term could not be translated. this is most likely an internal error: \n\t%s", err.(string)))
		}
	}()

	// Matches the pattern: {{name_to_replace}}
	nameExtractRegex := `(\{\{(?P<name_to_replace>[\w]+?)\}\})`
	re := regexp.MustCompile(nameExtractRegex)
	matches := re.FindAllStringSubmatch(regex, -1)
	numMatches := len(matches)

	keysToLookup := []string{}

	// The first three groups are useless for this case. So get the last one.
	for _, match := range matches {
		keysToLookup = append(keysToLookup, match[2])
	}

	// Validate all strings to replace and then replace them one by one.
	for i := 0; i < numMatches; i++ {
		stringToReplace, err := currentSearchTermData.GetValue(strings.TrimSpace(strings.ToLower(keysToLookup[i])))
		if err != nil {
			return "", fmt.Errorf("an error occurred when translating a search term reference. \n\tthe following key was not registered in a previous search term: %s", keysToLookup[i])
		}
		stringToReplace = validateStringToReplace(stringToReplace)
		foundString := re.FindString(regex)
		if foundString != "" {
			regex = strings.Replace(regex, foundString, stringToReplace, 1)
		}
	}

	// fmt.Println(regex)
	return regex, nil
}

// func validateStringToReplace takes any string that will get replaced
// by strings.Replace and and escapes any special characters with a "\"
func validateStringToReplace(regex string) string {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("a string could not be escaped. this is most likely an internal error: \n\t%s", err.(string)))
		}
	}()
	charactersToEscape := `([\*\+\?\\\.\^\[\]\$\&\|]{1})`
	re := regexp.MustCompile(charactersToEscape)

	regex = re.ReplaceAllString(regex, `\${1}`)
	return regex
}
