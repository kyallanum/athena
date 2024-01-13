package models

import (
	"fmt"
	"reflect"
	"testing"

	library "github.com/kyallanum/athena/models/library"
)

func createSearchTermTestData(reference, replace string) *library.SearchTermData {
	st_data := library.NewSearchTermData()
	st_data.AddValue(reference, replace)
	return st_data
}

func TestResolveLine(t *testing.T) {
	testTable := []struct {
		name           string
		line           string
		regex          string
		expectedOutput *map[string]string
		expectedError  any
	}{
		{
			name:           "Test_Good_Regex_No_Match",
			line:           "test line 1",
			regex:          "stuff",
			expectedOutput: nil,
			expectedError:  nil,
		},
		{
			name:           "Test_Good_Regex_With_Match",
			line:           "test line 1",
			regex:          `(?P<test_name>line \d+)`,
			expectedOutput: &map[string]string{},
			expectedError:  nil,
		},
		{
			name:           "Test_Bad_Regex",
			line:           "test line 1",
			regex:          `(stuff`,
			expectedOutput: nil,
			expectedError:  fmt.Errorf("the provided regular expression cannot be compiled: \n\tregexp: Compile(`(stuff`): error parsing regexp: missing closing ): `(stuff`"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); !checkExpectedError(err, test.expectedError) {
					t.Errorf("An error was returned improperly when calling resolveLine: \n\tExpected: %s\n\tReceived: %s", test.expectedError.(error).Error(), err.(error).Error())
				}
			}()

			result := resolveLine(test.line, test.regex)

			if reflect.TypeOf(result).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("resolveLine returned unexpected output type: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(test.expectedError).String())
			}
		})
	}
}

func TestTranslateSearchTermReference(t *testing.T) {
	testTable := []struct {
		name           string
		regex          string
		st_data        *library.SearchTermData
		expectedOutput string
		expectedError  error
	}{
		{
			name:           "Test_Good_Search_Term",
			regex:          `Testing {{test}}`,
			st_data:        createSearchTermTestData("test", "test1"),
			expectedOutput: "Testing test1",
			expectedError:  nil,
		},
		{
			name:           "Test_Bad_Search_Term",
			regex:          `Testing {{test}}`,
			st_data:        createSearchTermTestData("bad_test", "test"),
			expectedOutput: "",
			expectedError:  fmt.Errorf("an error occurred when translating a search term reference. \n\tthe following key was not registered in a previous search term: test"),
		},
	}

	for _, test := range testTable {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if err := recover(); !checkExpectedError(err, test.expectedError) {
					t.Errorf("An incorrect error was returned: \n\tExpected: %s\n\tReceived: %s", test.expectedError.Error(), err.(error).Error())
				}
			}()

			newRegex, err := translateSearchTermReference(test.regex, test.st_data)
			if err != nil {
				panic(err)
			}

			if reflect.TypeOf(newRegex).String() != reflect.TypeOf(test.expectedOutput).String() {
				t.Errorf("TranslateSearchTermReference did not output the correct data type: \n\tExpected: %s\n\tReceived: %s", reflect.TypeOf(test.expectedOutput).String(), reflect.TypeOf(newRegex).String())
			}

			if newRegex != test.expectedOutput {
				t.Errorf("TranslateSearchTermReference did not output the correct result: \n\tExpected: %s\n\tReceived: %s", test.expectedOutput, newRegex)
			}
		})
	}
}

func TestValidateString(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("An error occurred when it shouldn't have: \n\t%s", (err.(error).Error()))
		}
	}()

	stringToValidate := `*+?\\.^[]$&|`
	validatedString := escapeSpecialCharacters(stringToValidate)

	if validatedString != "\\*\\+\\?\\\\\\\\\\.\\^\\[\\]\\$\\&\\|" {
		t.Errorf("String was not validated properly: %s", validatedString)
	}
}
