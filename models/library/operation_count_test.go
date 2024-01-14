package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateNewCounts(t *testing.T) {
	testTable := []struct {
		name           string
		key            string
		expectedOutput ISummaryOperation
	}{
		{
			name: "Test_New_Count",
			key:  "test",
			expectedOutput: &Count{
				SummaryOperation{
					operation: "count",
				},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			testCountOperation := NewCountOperation(test.key)

			if testCountOperation.Operation() != test.expectedOutput.Operation() {
				t.Errorf("Operation did not return properly: Expected: %s\n\tReceived: %s\n\t", test.expectedOutput.Operation(), testCountOperation.Operation())
			}

			if reflect.TypeOf(testCountOperation).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Operation did not return the correct data type: Expected: %s\n\tReceived: %s\n\t", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(testCountOperation).String())
			}
		})
	}
}

func TestCountCalculateOperations(t *testing.T) {
	testTable := []struct {
		name           string
		ruleData       RuleData
		expectedOutput []string
		expectedError  error
	}{
		{
			name: "Test_Count_Good_Data",
			ruleData: RuleData{
				st_data_collection: []SearchTermData{
					{
						data: map[string]string{
							"test": "testing 1",
						},
					},
					{
						data: map[string]string{
							"test": "testing 2",
						},
					},
				},
			},
			expectedOutput: []string{"2"},
			expectedError:  nil,
		},
		{
			name: "Test_Count_No_RuleData",
			ruleData: RuleData{
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
			testCountOperation := NewCountOperation("test")
			numReturned, err := testCountOperation.CalculateOperation(test.ruleData)

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("An error occurred when calculating operation: Count\n\tExpected: %v, Received: %v", err, test.expectedError)
			}

			if reflect.TypeOf(numReturned).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("The calculation did not return the proper datatype: \n\tExpected: %s, Received: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(numReturned).String())
			}

			if !reflect.DeepEqual(numReturned, test.expectedOutput) {
				t.Errorf("Count did not return the expected value: \n\tExpected: %s\n\tReceived: %s", numReturned[0], test.expectedOutput[0])
			}
		})
	}
}
