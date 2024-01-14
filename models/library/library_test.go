package models

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func checkExpectedError(actualError, expectedError any) bool {
	if actualError == nil && expectedError == nil {
		return true
	}

	if actualError != nil {
		if expectedError != nil {
			if strings.Contains(actualError.(error).Error(), expectedError.(error).Error()) {
				return true
			}
		}
	}
	return false
}

func TestLibraryGetName(t *testing.T) {
	testTable := []struct {
		name           string
		library_name   string
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Test_Library_With_Name",
			library_name:   "test_library",
			expectedOutput: "test_library",
			expectedError:  nil,
		},
		{
			name:           "Test_Library_No_Name",
			library_name:   "",
			expectedOutput: "",
			expectedError:  fmt.Errorf("Library does not have a name assigned"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			current_library := New(test.library_name)
			current_library_name, err := current_library.Name()

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("An unexpected error occurred when getting the name for the Library: \n\tExpected: %v\n\tReceived: %v", test.expectedError.Error(), err.Error())
			}

			if current_library_name != test.expectedOutput {
				t.Errorf("Library name was not returned properly: \n\tExpected: %v\n\tReceived: %v", test.expectedOutput, current_library_name)
			}
		})
	}
}

func TestGetLibraryKeys(t *testing.T) {
	testTable := []struct {
		name           string
		library        *Library
		expectedOutput []string
		expectedError  error
	}{
		{
			name:           "Test_No_Keys",
			library:        New("test_library_no_keys"),
			expectedOutput: nil,
		},
		{
			name: "Test_With_Keys",
			library: &Library{
				name: "test_library_with_keys",
				rule_data_collection: map[string]RuleData{
					"ruleData1": NewRuleData(),
					"ruleData2": NewRuleData(),
				},
			},
			expectedOutput: []string{"ruleData1", "ruleData2"},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			libraryKeys := test.library.LibraryKeys()

			sort.Strings(libraryKeys)
			sort.Strings(test.expectedOutput)

			if !reflect.DeepEqual(libraryKeys, test.expectedOutput) {

				t.Errorf("Unexpected return from getting library keys: \n\tExpected: %v\n\tReceived: %v", libraryKeys, test.expectedOutput)
			}
		})
	}
}

func TestGetRuleData(t *testing.T) {
	testTable := []struct {
		name           string
		ruleToExtract  string
		library        *Library
		expectedOutput RuleData
		expectedError  error
	}{
		{
			name:          "Test_Get_Good_Rule_Data",
			ruleToExtract: "ruleData1",
			library: &Library{
				name: "testing",
				rule_data_collection: map[string]RuleData{
					"ruleData1": NewRuleData(),
				},
			},
			expectedOutput: RuleData{},
			expectedError:  nil,
		},
		{
			name:          "Test_Get_Rule_Data_No_Keys",
			ruleToExtract: "badRuleData",
			library: &Library{
				name:                 "testing",
				rule_data_collection: map[string]RuleData{},
			},
			expectedOutput: RuleData{},
			expectedError:  fmt.Errorf("could not find key: badRuleData in library"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			newRuleData, err := test.library.RuleData(test.ruleToExtract)
			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("Library did not return the right error data: \n\tExpected: %v\n\tReceived: %v", test.expectedError, err)
			}

			if reflect.TypeOf(newRuleData).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Library did not return the proper data type: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(newRuleData).String())
			}
		})
	}
}

func TestAddRuleDatas(t *testing.T) {
	testTable := []struct {
		name          string
		library       *Library
		keysToAdd     []string
		rulesToAdd    []RuleData
		expectedError error
	}{
		{
			name: "Add_Good_Rule_Data",
			library: &Library{
				name:                 "test_library",
				rule_data_collection: map[string]RuleData{},
			},
			keysToAdd: []string{"ruleData1"},
			rulesToAdd: []RuleData{
				NewRuleData(),
			},
			expectedError: nil,
		},
		{
			name: "Add_Rule_Data_Overwrite",
			library: &Library{
				name:                 "test_library",
				rule_data_collection: map[string]RuleData{},
			},
			keysToAdd: []string{"ruleData1", "ruleData1"},
			rulesToAdd: []RuleData{
				NewRuleData(),
				NewRuleData(),
			},
			expectedError: fmt.Errorf("could not set ruledata for key: ruleData1 - data is readonly"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			var err error
			for index, keyToAdd := range test.keysToAdd {
				err = test.library.AddRuleData(keyToAdd, &test.rulesToAdd[index])
				if err != nil {
					break
				}
			}

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("AddRuleData returned an unexpected error: \n\tExpected: %v\n\tReceived: %v", test.expectedError.Error(), err.Error())
			}
		})
	}
}
