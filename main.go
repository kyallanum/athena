package main

import (
	"fmt"

	"github.com/kyallanum/athena/v0.1.0/models"
	"github.com/kyallanum/athena/v0.1.0/utils"
)

var CONFIG_FILE = "/home/klanum/athena/examples/apt-term-config.json"
var LOG_FILE = "/home/klanum/athena/examples/apt-term.log"

func err_check(err error) {
	if err != nil {
		panic(err)
	}
}

//
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

func resolveLogFile(contents *models.LogFile, config *models.Configuration) (*models.Library, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("main -> resolveLogFile: \n\t%w", err)
	}

	ret_library := models.Library.New(models.Library{})

	fmt.Println("Resolving Log File")
	for i := 0; i < len(config.Rules); i++ {
		currentRuleData, err := utils.ResolveRule(contents, &config.Rules[i])
		if err != nil {
			return nil, wrap_error(err)
		}

		ret_library.AddRuleData(config.Rules[i].Name, currentRuleData)
	}

	return ret_library, nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\nAn issue occurred:", err)
		}
	}()

	fmt.Println("Athena v0.1.0 Starting")
	fmt.Println("Getting Configuration File: ", CONFIG_FILE, "...")

	configuration, err := utils.CreateConfiguration(CONFIG_FILE)
	err_check(err)

	configuration.TranslateRegexGroups()
	fmt.Println("Configuration Loaded")

	fmt.Println("Loading Log File: ", LOG_FILE, "... ")
	logFileContents, err := utils.LoadLogFile(LOG_FILE)
	err_check(err)
	fmt.Println("Loaded")

	resolveLogFile(logFileContents, configuration)
	// library := resolveRules(logFileContents, configuration)
	// printSummary(configuration, library)
	// // fmt.Println(library)
}
