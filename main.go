package main

import (
	"fmt"

	"github.com/kyallanum/athena/v0.1.0/utils"
)

var CONFIG_FILE = "https://raw.githubusercontent.com/kyallanum/athena/main/examples/apt-term-config.json"
var LOG_FILE = "/home/klanum/athena/examples/apt-term.log"

func err_check(err error) {
	if err != nil {
		panic(err)
	}
}

// func resolveRules(contents []string, config *models.Configuration) map[string](map[string][]string) {
// 	fmt.Println("Resolving Log File")

// 	// This is a map containing all strings that were resolved from each line.
// 	// Grouped by key in regex, which is further grouped by Rule Name
// 	ruleLibrary := make(map[string](map[string][]string))

// 	for _, rule := range *config.Rules {
// 		fmt.Println("Resolving Rule:", *rule.Name)
// 		linesResolved := make([]int, 0)
// 		for {
// 			currentRuleDict := resolveRuleForFileContents(contents, &linesResolved, &rule)
// 			if currentRuleDict == nil {
// 				break
// 			}
// 			// fmt.Println(currentRuleDict)

// 			utils.AddDictToLibrary(currentRuleDict, &ruleLibrary)
// 		}
// 	}

// 	return ruleLibrary
// }

// func resolveRuleForFileContents(contents []string, linesResolved *([]int), currentRule *models.Rule) *map[string](map[string]string) {
// 	currentSearchTermIndex := 0
// 	currentKeys := make(map[string](map[string]string))

// 	for index, lineContent := range contents {
// 		if currentSearchTermIndex == len(*currentRule.SearchTerms) {
// 			break
// 		}
// 		if slices.Contains(*linesResolved, index) {
// 			continue
// 		}

// 		keyTranslatedRegex := utils.TranslateNames((*currentRule.SearchTerms)[currentSearchTermIndex], currentKeys)

// 		result := utils.ResolveRegexpNames(lineContent, keyTranslatedRegex)
// 		if result == nil {
// 			continue
// 		}

// 		if *currentRule.PrintLog {
// 			fmt.Printf("%d: %s\n", index+1, lineContent)
// 		}

// 		*linesResolved = append(*linesResolved, index)
// 		currentSearchTermIndex++

// 		for key, value := range *result {
// 			if currentKeys[*currentRule.Name] == nil {
// 				currentKeys[*currentRule.Name] = make(map[string]string)
// 			}
// 			currentKeys[*currentRule.Name][key] = value
// 		}
// 	}

// 	if len(currentKeys) > 0 {
// 		return &currentKeys
// 	} else {
// 		return nil
// 	}
// }

// func printSummary(config *models.Configuration, library map[string](map[string][]string)) {
// 	fmt.Printf("\n--------------- %s Log File Summary ---------------\n", *config.Name)
// 	for _, rule := range *config.Rules {
// 		fmt.Printf("---------- %s Rule ----------\n", *rule.Name)
// 		for _, summaryString := range *rule.Summary {
// 			fmt.Println(summaryString)
// 			// utils.TranslateSummaryLine(summaryString, library)
// 		}
// 	}
// }

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\nAn issue occurred:", err)
		}
	}()

	fmt.Println("Athena v0.1.0 Starting")
	fmt.Println("Getting Configuration File: ", CONFIG_FILE, "... ")

	configuration, err := utils.CreateConfiguration(CONFIG_FILE)
	err_check(err)

	fmt.Println(configuration)

	// // Parse JSON and validate
	// configuration := &models.Configuration{}
	// json.Unmarshal(configBytes, configuration)
	// err = configuration.ValidateConfig()
	// err_check(err)

	// fmt.Println("Translating Regular Expressions to Go Standards.")

	// // Transform Regex if necessary
	// for ruleIndex := range *configuration.Rules {
	// 	err_check(err)
	// 	for searchTermIndex, currentSearchTerm := range *(*configuration.Rules)[ruleIndex].SearchTerms {
	// 		utils.TranslateRegex(&currentSearchTerm)
	// 		(*(*configuration.Rules)[ruleIndex].SearchTerms)[searchTermIndex] = currentSearchTerm
	// 	}
	// }
	// fmt.Println("Configuration Loaded")
	// fmt.Print("Loading Log File ", LOG_FILE, "... ")
	// logFileContents, err := utils.LoadInputFile(LOG_FILE)
	// err_check(err)
	// fmt.Println("Loaded")
	// fmt.Println()
	// library := resolveRules(logFileContents, configuration)
	// printSummary(configuration, library)
	// // fmt.Println(library)
}
