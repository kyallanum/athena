package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kyallanum/athena/cmd"
	"github.com/kyallanum/athena/models/config"
	"github.com/kyallanum/athena/models/library"
	"github.com/kyallanum/athena/models/logs"
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

func resolveLogFile(contents *logs.LogFile, configuration *config.Configuration) (*library.Library, error) {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to resolve log file: \n\t%w", err)
	}

	if contents == nil || contents.Len() == 0 {
		return nil, fmt.Errorf("log file contains no contents")
	} else if configuration == nil || (configuration.Name == "" && configuration.Rules == nil) {
		return nil, fmt.Errorf("configuration file has no contents")
	} else if configuration.Name == "" {
		return nil, fmt.Errorf("configuration file contains no log name")
	} else if configuration.Rules == nil || len(configuration.Rules) == 0 {
		return nil, fmt.Errorf("configuration does not have any rules")
	}

	ret_library := library.New(configuration.Name)

	fmt.Println("Resolving Log File")
	for i := 0; i < len(configuration.Rules); i++ {
		currentRuleData, err := config.ResolveRule(contents, &configuration.Rules[i])
		if err != nil {
			return nil, wrapError(err)
		}

		ret_library.AddRuleData(configuration.Rules[i].Name, currentRuleData)
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
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, "\nAn error occured: ", err)
		}
	}()

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
