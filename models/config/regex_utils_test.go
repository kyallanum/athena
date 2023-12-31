package models

import (
	"reflect"
	"testing"

	library "github.com/kyallanum/athena/v1.0.0/models/library"
)

func TestResolveLine(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("An error was returned improperly when calling resolveLine: \n\t%s", (err.(error)).Error())
		}
	}()
	line := "test line 1"
	regex := "stuff"

	result := resolveLine(line, regex)
	if result != nil {
		t.Errorf("An incorrect object was returned when calling resolveLine")
	}

	regex = `(?P<test_name>line \d+)`
	result = resolveLine(line, regex)

	if reflect.TypeOf(result).String() != "*map[string]string" {
		t.Errorf("resolveLine did not return the proper data type: \n\t%s", reflect.TypeOf(result).String())
	}
}

func TestResolveLineBadRegex(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("An error was not returned properly when it should have")
		} else if (err.(error)).Error() != "the provided regular expression cannot be compiled: \n\tregexp: Compile(`(stuff`): error parsing regexp: missing closing ): `(stuff`" {
			t.Errorf("The improper error was returned when calling with a bad regex: \n\t%s", (err.(error)).Error())
		}
	}()

	line := "test line 1"
	regex := "(stuff"

	_ = resolveLine(line, regex)
}

func TestTranslateSearchTermReference(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("An error was returned when it shouldn't have: \n\t%s", (err.(error).Error()))
		}
	}()

	regex := `Testing {{test}}`
	st_data := library.SearchTermData.New(library.SearchTermData{})
	st_data.AddValue("test", "Test1")

	newRegex, err := translateSearchTermReference(regex, st_data)
	if err != nil {
		t.Errorf("An error was returned when it shouldn't have: \n\t%s", err.Error())
	}

	if newRegex != "Testing Test1" {
		t.Errorf("TranslateSearchTermReference returned the wrong value: \n\t%s", newRegex)
	}
}

func TestTranslateSearchTermReferenceBadReference(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("An error was returned when it shouldn't have: \n\t%s", (err.(error).Error()))
		}
	}()

	regex := `Testing {{test}}`
	st_data := library.SearchTermData.New(library.SearchTermData{})
	st_data.AddValue("bad_test", "testing")

	_, err := translateSearchTermReference(regex, st_data)
	if err == nil {
		t.Errorf("An error was not returned when it should have.")
	}

	if err.Error() != "an error occurred when translating a search term reference. \n\tthe following key was not registered in a previous search term: test" {
		t.Errorf("An improper error was returned when attempting to translate search term reference: \n\t%s", err.Error())
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
