package main

import (
	"fmt"

	models "github.com/kyallanum/athena/v0.1.0/models"
	config "github.com/kyallanum/athena/v0.1.0/models/config"
	library "github.com/kyallanum/athena/v0.1.0/models/library"
	"github.com/kyallanum/athena/v0.1.0/utils"
)

var CONFIG_FILE = "/home/klanum/athena/examples/apt-term.log"
var LOG_FILE = "/home/klanum/athena/examples/apt-term.log"

func err_check(err error) {
	if err != nil {
		panic(err)
	}
}

// func printSummary(config *models.Configuration, library *models.Library) {
// 	fmt.Printf("\n--------------- %s Log File Summary ---------------\n", config.Name)
// 	for _, rule := range config.Rules {
// 		fmt.Printf("---------- %s Rule ----------\n", rule.Name)
// 		for _, summaryString := range rule.Summary {
// 			fmt.Println(summaryString)
// 			// utils.TranslateSummaryLine(summaryString, library)
// 		}
// 	}
// }

func resolveLogFile(contents *models.LogFile, config *config.Configuration) (*library.Library, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("main -> resolveLogFile: \n\t%w", err)
	}

	ret_library := library.Library.New(library.Library{})

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
	fmt.Println()

	fmt.Println("Loading Log File: ", LOG_FILE, "... ")
	logFileContents, err := utils.LoadLogFile(LOG_FILE)
	err_check(err)
	fmt.Println("Loaded")
	fmt.Println()

	library, err := resolveLogFile(logFileContents, configuration)
	if err != nil {
		err_check(err)
	}

	resolveSummary(configuration, library)

	printSummary(configuration, library)
	// // fmt.Println(library)
}
