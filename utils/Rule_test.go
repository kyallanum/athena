package utils

import (
	"os"
	"reflect"
	"testing"

	config "github.com/kyallanum/athena/v1.0.0/models/config"
)

func TestResolveRule(t *testing.T) {
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := LoadLogFile("../examples/apt-term.log")

	currentConfig, _ := CreateConfiguration("../examples/apt-term-config.json")

	currentRule := currentConfig.Rules[0]

	ruleData, err := ResolveRule(logFile, &currentRule)
	if err != nil {
		t.Errorf("An error was returned when one should not have been: \n\t%s", err.Error())
	}

	if reflect.TypeOf(ruleData).String() != "*models.RuleData" {
		t.Errorf("The incorrect datatype was not returned: \n\t%s", reflect.TypeOf(ruleData).String())
	}
}

func TestResolveRuleBadRule(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("An error was not returned when one should have been.")
		} else if (err.(error)).Error() != "unable to resolve search terms for rule : \n\truntime error: index out of range [0] with length 0" {
			t.Errorf("%s", err.(error).Error())
		}
	}()
	os.Stdout, _ = os.Open(os.DevNull)
	defer os.Stdout.Close()

	logFile, _ := LoadLogFile("../examples/apt-term.log")

	currentRule := &config.Rule{}

	ResolveRule(logFile, currentRule)
}
