package models

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCreateNewPrint(t *testing.T) {
	testTable := []struct {
		name           string
		key            string
		expectedOutput ISummaryOperation
	}{
		{
			name: "Test_Create_New_Print_Operation",
			key:  "testing",
			expectedOutput: &Print{
				SummaryOperation: SummaryOperation{
					operation: "print",
					key:       "testing",
				},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			testPrint := NewPrintOperation(test.key)

			if testPrint.Operation() != test.expectedOutput.Operation() {
				t.Errorf("Creating new Operation object returned improper operation: \n\tExpected: %s\n\tReceived: %s", test.expectedOutput.Operation(), testPrint.Operation())
			}

			if testPrint.Key() != test.expectedOutput.Key() {
				t.Errorf("Creating new Operation object does not assign the proper key: \n\tExpected: %s\n\tReceived: %s", test.expectedOutput.Key(), testPrint.Key())
			}

			if reflect.TypeOf(testPrint).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Operation that was created is not of the right type: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(testPrint).String())
			}
		})
	}
}

func TestPrintCalculateOperation(t *testing.T) {
	testTable := []struct {
		name           string
		keyToAdd       string
		testRuleData   RuleData
		expectedOutput []string
		expectedError  error
	}{
		{
			name:     "Test_Print_Calculate_Operation_Good_Data",
			keyToAdd: "testing",
			testRuleData: RuleData{
				st_data_collection: []SearchTermData{
					{
						data: map[string]string{
							"testing": "test 1",
						},
					},
					{
						data: map[string]string{
							"testing": "test 2",
						},
					},
				},
			},
			expectedOutput: []string{"test 1", "test 2"},
			expectedError:  nil,
		},
		{
			name:     "Test_Print_Calculate_Operation_No_Rule_Data",
			keyToAdd: "testing",
			testRuleData: RuleData{
				st_data_collection: []SearchTermData{
					{
						data: map[string]string{},
					},
				},
			},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("this value was never extracted during search phase"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			testOperation := NewPrintOperation("testing")
			printOutput, err := testOperation.CalculateOperation(test.testRuleData)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("An error occurred when attempting to resolve lines to print: \n\tExpected: %v\n\tReceived: %v", test.expectedError, err)
			}

			if reflect.TypeOf(printOutput).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Print Operation did not return the proper output datatype: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(printOutput).String())
			}

			sort.Strings(printOutput)
			sort.Strings(test.expectedOutput)

			if !reflect.DeepEqual(printOutput, test.expectedOutput) {
				t.Errorf("Print Operation did not return the proper output: \n\tExpected: %v\n\tReceived: %v", test.expectedOutput, printOutput)
			}
		})
	}
}
