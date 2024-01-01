package cmd

import (
	"fmt"
	"os"

	config "github.com/kyallanum/athena/models/config"
	library "github.com/kyallanum/athena/models/library"
	logger_pkg "github.com/kyallanum/athena/models/logger"
	logs "github.com/kyallanum/athena/models/logs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configFile, logFile, logOutput string

var rootCmd = &cobra.Command{
	Use:           "athena [flags]",
	Short:         "A text and log file parser to discern important information",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(configFile) == 0 {
			return fmt.Errorf("exiting: required flag(s) \"config\" not set")
		}
		if len(logFile) == 0 {
			return fmt.Errorf("exiting: required flag(s) \"log-file\" not set")
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

	logger := logger_pkg.New()
	defer func() {
		if err := recover(); err != nil {
			logger.Fatalf("An error occured: \n\t%s", err)
		}
	}()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	} else if rootCmd.Flags().Changed("help") {
		os.Exit(1)
	}

	if logOutput != "" {
		logger_pkg.AddFileLogger(logger, logOutput)
	}
	errCheck(err)

	logger.Info("Athena v1.0.0 Starting")

	logger.Info("Getting Configuration File: ", configFile, "...")
	configuration, err := config.CreateConfiguration(configFile)
	errCheck(err)
	logger.Info("Configuration Loaded")

	logger.Info("Loading Log File: ", logFile, "... ")
	logFileContents, err := logs.LoadLogFile(logFile)
	errCheck(err)
	logger.Info("Log File Loaded")

	library, err := resolveLogFile(logFileContents, configuration, logger)
	errCheck(err)

	err = printSummary(library, logger)
	errCheck(err)

	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&configFile, "config", "c", os.Getenv("ATHENA_CONFIG_FILE"), "")
	rootCmd.Flags().StringVarP(&logFile, "log-file", "l", os.Getenv("ATHENA_LOG_FILE"), "")
	rootCmd.Flags().StringVarP(&logOutput, "log-output", "o", os.Getenv("ATHENA_LOG_OUTPUT"), "")
}

func resolveLogFile(contents *logs.LogFile, configuration *config.Configuration, logger *logrus.Logger) (*library.Library, error) {

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

	logger.Info("Resolving Log File")
	for i := 0; i < len(configuration.Rules); i++ {
		currentRuleData, err := config.ResolveRule(contents, &configuration.Rules[i], logger)
		if err != nil {
			return nil, wrapError(err)
		}

		ret_library.AddRuleData(configuration.Rules[i].Name, currentRuleData)
	}

	logger.Info("Log File Resolved")

	return ret_library, nil
}

func printSummary(library *library.Library, logger *logrus.Logger) error {
	wrapError := func(err error) error {
		return fmt.Errorf("unable to print summary: \n\t%w", err)
	}

	libraryName, err := library.Name()
	if err != nil {
		return wrapError(err)
	}

	logger.Infof("\n--------------- %s Log File Summary ---------------\n", libraryName)
	libraryKeys := library.LibraryKeys()
	for _, rule := range libraryKeys {
		logger.Infof("Rule: %s\n", rule)
		ruleData, _ := library.RuleData(rule)
		summaryDataLen := ruleData.SummaryDataLen()
		if summaryDataLen == 0 {
			logger.Info("No summary lines provided.")
		} else {
			for i := 0; i < summaryDataLen; i++ {
				logger.Info("\t", ruleData.SummaryData(i))
			}
		}
		logger.Info()
	}

	return nil
}
