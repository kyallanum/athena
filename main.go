package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	config "github.com/kyallanum/athena/v1.0.0/models/config"
	library "github.com/kyallanum/athena/v1.0.0/models/library"
	logs "github.com/kyallanum/athena/v1.0.0/models/logs"
	"github.com/kyallanum/athena/v1.0.0/utils"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func parseFlags(configFile *string, logFile *string) error {
	var configFlag string
	var logFlag string

	flag.StringVar(&configFlag, "c", "", "")
	flag.StringVar(&configFlag, "config", "", "")
	flag.StringVar(&logFlag, "l", "", "")
	flag.StringVar(&logFlag, "log-file", "", "")
	flag.Parse()

	if *configFile == "" {
		*configFile = configFlag
	}
	if *logFile == "" {
		*logFile = logFlag
	}

	if *configFile == "" {
		return fmt.Errorf("configuration file was not specified")
	}

	if *logFile == "" {
		return fmt.Errorf("log file was not specified")
	}

	return nil
}

func resolveLogFile(contents *logs.LogFile, config *config.Configuration) (*library.Library, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to resolve log file: \n\t%w", err)
	}

	if contents == nil || contents.Len() == 0 {
		return nil, fmt.Errorf("log file contains no contents")
	} else if config == nil || (config.Name == "" && config.Rules == nil) {
		return nil, fmt.Errorf("configuration file has no contents")
	} else if config.Name == "" {
		return nil, fmt.Errorf("configuration file contains no log name")
	} else if config.Rules == nil || len(config.Rules) == 0 {
		return nil, fmt.Errorf("configuration does not have any rules")
	}

	ret_library := library.Library.New(library.Library{}, config.Name)

	fmt.Println("Resolving Log File")
	for i := 0; i < len(config.Rules); i++ {
		currentRuleData, err := utils.ResolveRule(contents, &config.Rules[i])
		if err != nil {
			return nil, wrapError(err)
		}

		ret_library.AddRuleData(config.Rules[i].Name, currentRuleData)
	}

	fmt.Println("Log File Resolved")

	return ret_library, nil
}

func printSummary(library *library.Library) error {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to print summary: \n\t%w", err)
	}

	libraryName, err := library.Name()
	if err != nil {
		return wrapError(err)
	}

	fmt.Printf("\n--------------- %s Log File Summary ---------------\n", libraryName)
	libraryKeys := library.LibraryKeys()
	for _, rule := range libraryKeys {
		fmt.Printf("Rule: %s\n", rule)
		ruleData, _ := library.RuleData(rule)
		summaryDataLen := ruleData.SummaryDataLen()
		if summaryDataLen == 0 {
			fmt.Println("No summary lines provided.")
		} else {
			for i := 0; i < summaryDataLen; i++ {
				fmt.Println("\t", ruleData.SummaryData(i))
			}
		}
		fmt.Println()
	}

	return nil
}

func main() {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	defer func() {
		if err := recover(); err != nil {
			logger.Fatal("\nAn issue occurred:", err)
		}
	}()

	CONFIG_FILE := os.Getenv("ATHENA_CONFIG_FILE")
	LOG_FILE := os.Getenv("ATHENA_LOG_FILE")

	parseFlags(&CONFIG_FILE, &LOG_FILE)

	logger.Print("Athena v1.0.0 Starting")

	fmt.Println("Getting Configuration File: ", CONFIG_FILE, "...")
	configuration, err := config.CreateConfiguration(CONFIG_FILE)
	errCheck(err)
	fmt.Println("Configuration Loaded")

	fmt.Println("Loading Log File: ", LOG_FILE, "... ")
	logFileContents, err := logs.LoadLogFile(LOG_FILE)
	errCheck(err)
	fmt.Println("Log File Loaded")

	library, err := resolveLogFile(logFileContents, configuration)
	errCheck(err)

	err = printSummary(library)
	errCheck(err)
}
