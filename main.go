package main

import (
	"fmt"

	models "github.com/kyallanum/athena/v0.1.0/models"
	config "github.com/kyallanum/athena/v0.1.0/models/config"
	library "github.com/kyallanum/athena/v0.1.0/models/library"
	"github.com/kyallanum/athena/v0.1.0/utils"
)

var CONFIG_FILE = "/home/klanum/athena/examples/apt-term-config.json"
var LOG_FILE = "/home/klanum/athena/examples/apt-term.log"

func err_check(err error) {
	if err != nil {
		panic(err)
	}
}

func resolveLogFile(contents *models.LogFile, config *config.Configuration) (*library.Library, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("unable to resolve log file: \n\t%w", err)
	}

	ret_library := library.Library.New(library.Library{}, config.Name)

	fmt.Println("Resolving Log File")
	for i := 0; i < len(config.Rules); i++ {
		currentRuleData, err := utils.ResolveRule(contents, &config.Rules[i])
		if err != nil {
			return nil, wrap_error(err)
		}

		ret_library.AddRuleData(config.Rules[i].Name, currentRuleData)
	}

	fmt.Println("Log File Resolved")

	return ret_library, nil
}

func printSummary(library *library.Library) {
	fmt.Printf("\n--------------- %s Log File Summary ---------------\n", library.GetName())
	libraryKeys := library.GetLibraryKeys()
	for _, rule := range libraryKeys {
		fmt.Printf("Rule: %s\n", rule)
		ruleData, _ := library.GetRuleData(rule)
		summaryDataLen := ruleData.GetSummaryDataLen()
		if summaryDataLen == 0 {
			fmt.Println("No summary lines provided.")
		} else {
			for i := 0; i < summaryDataLen; i++ {
				fmt.Println("\t", ruleData.GetSummaryData(i))
			}
		}
		fmt.Println()
	}
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
	fmt.Println("Configuration Loaded")

	fmt.Println("Loading Log File: ", LOG_FILE, "... ")
	logFileContents, err := utils.LoadLogFile(LOG_FILE)
	err_check(err)
	fmt.Println("Log File Loaded")

	library, err := resolveLogFile(logFileContents, configuration)
	if err != nil {
		err_check(err)
	}
	printSummary(library)
}
