package models

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCreateNewSearchTermDatas(t *testing.T) {
	testTable := []struct {
		name           string
		expectedOutput *SearchTermData
	}{
		{
			name: "Test_New_Search_Term_Data",
			expectedOutput: &SearchTermData{
				data: map[string]string{},
			},
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			stData := NewSearchTermData()
			if reflect.TypeOf(stData).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Creating new SearchTermData returned the wrong data type: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(stData).String())
			}

			if !reflect.DeepEqual(stData.data, test.expectedOutput.data) {
				t.Errorf("Creating new SearchTermData returned the wrong data: \n\tExpected: %v\n\tReceived: %v", test.expectedOutput.data, stData.data)
			}
		})
	}
}

func TestSearchTermDataGetKeys(t *testing.T) {
	testTable := []struct {
		name           string
		searchTermData *SearchTermData
		expectedOutput []string
	}{
		{
			name: "Test_Get_Keys_Good_Data",
			searchTermData: &SearchTermData{
				data: map[string]string{
					"test1": "Testing 1",
					"test2": "Testing 2",
				},
			},
			expectedOutput: []string{"test1", "test2"},
		},
		{
			name: "Test_Get_Keys_No_Keys",
			searchTermData: &SearchTermData{
				data: map[string]string{},
			},
			expectedOutput: nil,
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			keysReturned := test.searchTermData.Keys()

			if reflect.TypeOf(keysReturned).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("Get Keys returned the wrong data type: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(keysReturned).String())
			}

			sort.Strings(keysReturned)
			sort.Strings(test.expectedOutput)

			if !reflect.DeepEqual(keysReturned, test.expectedOutput) {
				t.Errorf("Get Keys returned the wrong data: \n\tExpected: %v\n\tReceived: %v", test.expectedOutput, keysReturned)
			}
		})
	}
}

func TestGetValues(t *testing.T) {
	testTable := []struct {
		name           string
		key            string
		searchTermData *SearchTermData
		expectedOutput string
		expectedError  error
	}{
		{
			name: "Test_Get_Valuee_Good_Data",
			key:  "test1",
			searchTermData: &SearchTermData{
				data: map[string]string{
					"test1": "testing",
				},
			},
			expectedOutput: "testing",
			expectedError:  nil,
		},
		{
			name: "Test_Get_Value_Wrong_Key",
			key:  "test1",
			searchTermData: &SearchTermData{
				data: map[string]string{},
			},
			expectedOutput: "",
			expectedError:  fmt.Errorf("could not find key: test1 in searchtermdata"),
		},
	}

	for _, test := range testTable {
		currentValue, err := test.searchTermData.Value(test.key)
		if !checkExpectedError(err, test.expectedError) {
			t.Errorf("Get Value returned an unexpected error: \n\tExpected: %v\n\tReceived: %v", test.expectedError, err)
		}

		if reflect.TypeOf(currentValue).String() != reflect.TypeOf(test.expectedOutput).String() {
			t.Errorf("Get Value returned unexpected data type: \n\tExpected: %v\n\tReceived: %v", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(currentValue).String())
		}

		if currentValue != test.expectedOutput {
			t.Errorf("Get Value returned incorrect data: \n\tExpected: %v\n\tReceived: %v", test.expectedOutput, currentValue)
		}
	}
}

func TestAddValues(t *testing.T) {
	testTable := []struct {
		name          string
		keysToAdd     []string
		valuesToAdd   []string
		expectedError error
	}{
		{
			name:          "Test_Good_Values",
			keysToAdd:     []string{"test1"},
			valuesToAdd:   []string{"testing"},
			expectedError: nil,
		},
		{
			name:          "Test_Overwrite_Error",
			keysToAdd:     []string{"test1", "test1"},
			valuesToAdd:   []string{"testing", "another test"},
			expectedError: fmt.Errorf("unable to set value: another test for key: test1 - data is readonly"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			var err error
			st_data := NewSearchTermData()

			for index, keyToAdd := range test.keysToAdd {
				err = st_data.AddValue(keyToAdd, test.valuesToAdd[index])

				if err != nil {
					break
				}
			}

			if !checkExpectedError(err, test.expectedError) {
				t.Errorf("SearchTermData.AddValue did not return the expected error: \n\tExpected: %v\n\tReceived: %v", test.expectedError, err)
			}
		})
	}
}
