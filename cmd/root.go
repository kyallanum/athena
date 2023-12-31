package cmd

import (
	"fmt"
	"os"

	config "github.com/kyallanum/athena/v1.0.0/models/config"
	library "github.com/kyallanum/athena/v1.0.0/models/library"
	logs "github.com/kyallanum/athena/v1.0.0/models/logs"
	"github.com/spf13/cobra"
)

var configFile, logFile string

var rootCmd = &cobra.Command{
	Use:           "athena [flags]",
	Short:         "A text and log file parser to discern important information",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(configFile) == 0 {
			return fmt.Errorf(`required flag(s) "config" not set`)
		}
		if len(logFile) == 0 {
			return fmt.Errorf(`required flag(s) "log-file" not set`)
		}
		return nil
	},
}

func Execute() error {
	errCheck := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		return err
	}
	fmt.Println("Athena v1.0.0 Starting")

	fmt.Println("Getting Configuration File: ", configFile, "...")
	configuration, err := config.CreateConfiguration(configFile)
	errCheck(err)
	fmt.Println("Configuration Loaded")

	fmt.Println("Loading Log File: ", logFile, "... ")
	logFileContents, err := logs.LoadLogFile(logFile)
	errCheck(err)
	fmt.Println("Log File Loaded")

	library, err := resolveLogFile(logFileContents, configuration)
	errCheck(err)

	err = printSummary(library)
	errCheck(err)

	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", os.Getenv(""), "")
	rootCmd.Flags().StringVarP(&logFile, "log-file", "l", os.Getenv(""), "")
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

	ret_library := library.Library.New(library.Library{}, configuration.Name)

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
