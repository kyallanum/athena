package models

import (
	"reflect"
	"testing"
)

func TestCreateNewSearchTermData(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})

	if reflect.TypeOf(st_data).String() != "*models.SearchTermData" {
		t.Errorf("Creating new SearchTermData returned the wrong type")
	}
}

func TestSearchTermDataGetKeys(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})
	st_data.AddValue("test1", "testing")
	st_data.AddValue("test2", "testing")

	keys_returned := st_data.Keys()

	if len(keys_returned) != 2 {
		t.Errorf("Get Keys returned the wrong number of keys")
	}

	if reflect.TypeOf(keys_returned).String() != "[]string" {
		t.Errorf("Get Keys returned the wrong datatype")
	}
}

func TestSearchTermDataGetKeysNoKeys(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})

	keys_returned := st_data.Keys()

	if keys_returned != nil {
		t.Errorf("Get Keys returned the wrong answer with no keys")
	}
}

func TestGetValue(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})
	st_data.AddValue("test1", "testing")

	currentValue, err := st_data.Value("test1")
	if err != nil {
		t.Errorf("Error improperly returned when getting value.")
	}

	if currentValue != "testing" {
		t.Errorf("Value improperly returned when getting value.")
	}

	if reflect.TypeOf(currentValue).String() != "string" {
		t.Errorf("GetValue returned improper datatype")
	}
}

func TestGetValueWrongValue(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})

	_, err := st_data.Value("test1")

	if err == nil {
		t.Errorf("Error not returned when it should have.")
	}

	if err.Error() != "could not find key: test1 in searchtermdata" {
		t.Errorf("Error improperly returned when getting new searchtermdata")
	}
}

func TestAddValue(t *testing.T) {
	st_data := SearchTermData.New(SearchTermData{})

	err := st_data.AddValue("test1", "testing")
	if err != nil {
		t.Errorf("Error returned when there should not have been")
	}

	err = st_data.AddValue("test1", "another test")
	if err.Error() != "unable to set value: another test for key: test1 - data is readonly" {
		t.Errorf("Error improperly returned when attempting to overwrite data")
	}
}
