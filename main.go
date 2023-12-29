package main

import (
	"flag"
	"fmt"
	"os"

	models "github.com/kyallanum/athena/v1.0.0/models"
	config "github.com/kyallanum/athena/v1.0.0/models/config"
	library "github.com/kyallanum/athena/v1.0.0/models/library"
	"github.com/kyallanum/athena/v1.0.0/utils"
)

func err_check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseFlags(configFile *string, logFile *string) error {
	var config string
	var logfile string

	flag.StringVar(&config, "c", "", "")
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&logfile, "l", "", "")
	flag.StringVar(&logfile, "log-file", "", "")
	flag.Parse()

	if *configFile == "" {
		*configFile = config
	}
	if *logFile == "" {
		*logFile = logfile
	}

	if *configFile == "" {
		return fmt.Errorf("configuration file was not specified")
	}

	if *logFile == "" {
		return fmt.Errorf("log file was not specified")
	}

	return nil
}

func resolveLogFile(contents *models.LogFile, config *config.Configuration) (*library.Library, error) {
	wrap_error := func(err error) error {
		return fmt.Errorf("unable to resolve log file: \n\t%w", err)
	}

	if contents == nil || contents.GetLen() == 0 {
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
			return nil, wrap_error(err)
		}

		ret_library.AddRuleData(config.Rules[i].Name, currentRuleData)
	}

	fmt.Println("Log File Resolved")

	return ret_library, nil
}

func printSummary(library *library.Library) error {
	wrap_error := func(err error) error {
		return fmt.Errorf("unable to print summary: \n\t%w", err)
	}

	libraryName, err := library.GetName()
	if err != nil {
		return wrap_error(err)
	}

	fmt.Printf("\n--------------- %s Log File Summary ---------------\n", libraryName)
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

	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\nAn issue occurred:", err)
		}
	}()

	CONFIG_FILE := os.Getenv("ATHENA_CONFIG_FILE")
	LOG_FILE := os.Getenv("ATHENA_LOG_FILE")

	parseFlags(&CONFIG_FILE, &LOG_FILE)

	fmt.Println("Athena v1.0.0 Starting")

	fmt.Println("Getting Configuration File: ", CONFIG_FILE, "...")
	configuration, err := utils.CreateConfiguration(CONFIG_FILE)
	err_check(err)
	fmt.Println("Configuration Loaded")

	fmt.Println("Loading Log File: ", LOG_FILE, "... ")
	logFileContents, err := utils.LoadLogFile(LOG_FILE)
	err_check(err)
	fmt.Println("Log File Loaded")

	library, err := resolveLogFile(logFileContents, configuration)
	err_check(err)

	err = printSummary(library)
	err_check(err)
}
