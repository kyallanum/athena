package cmd

import (
	"fmt"
	"os"

	models "github.com/kyallanum/athena/v1.0.0/models"
	config "github.com/kyallanum/athena/v1.0.0/models/config"
	library "github.com/kyallanum/athena/v1.0.0/models/library"
	"github.com/kyallanum/athena/v1.0.0/utils"
	"github.com/spf13/cobra"
)

var configFile, logFile string

var rootCmd = &cobra.Command{
	Use:   "athena [flags]",
	Short: "A text and log file parser to discern important information",
  SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(configFile) == 0 {
			return fmt.Errorf("")
		}
		if len(logFile) == 0 {
			return fmt.Errorf("")
		}
		fmt.Println("Athena v1.0.0 Starting")

		fmt.Println("Getting Configuration File: ", configFile, "...")
		configuration, err := utils.CreateConfiguration(configFile)
		if err != nil {
			return err
		}
		fmt.Println("Configuration Loaded")

		fmt.Println("Loading Log File: ", logFile, "... ")
		logFileContents, err := utils.LoadLogFile(logFile)
		if err != nil {
			return err
		}
		fmt.Println("Log File Loaded")

		library, err := resolveLogFile(logFileContents, configuration)
		if err != nil {
			return err
		}

		if err = printSummary(library); err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", os.Getenv("ATHENA_CONFIG_FILE"), "path to config file")
	rootCmd.Flags().StringVarP(&logFile, "log-file", "l", os.Getenv("ATHENA_LOG_FILE"), "path to log file")
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
