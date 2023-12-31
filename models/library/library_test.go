package models

import (
	"reflect"
	"testing"
)

func TestLibraryGetName(t *testing.T) {
	current_library := New("test_stuff")
	library_name, err := current_library.Name()
	if err != nil {
		t.Errorf("an error occurred when getting name for library")
	}

	if library_name != "test_stuff" {
		t.Errorf("library name returned improperly")
	}
}

func TestLibraryGetNameNoName(t *testing.T) {
	current_library := New("")
	_, err := current_library.Name()

	if err == nil {
		t.Errorf("Error not returned when Library name is nil")
	}

	if err.Error() != "Library does not have a name assigned" {
		t.Errorf("Error not returned properly when Library name is nil")
	}
}

func TestGetLibraryKeys(t *testing.T) {
	ruleData1 := NewRuleData()
	ruleData2 := NewRuleData()

	library := New("test_library")
	library.AddRuleData("ruleData1", &ruleData1)
	library.AddRuleData("ruleData2", &ruleData2)

	libraryKeys := library.LibraryKeys()

	if len(libraryKeys) != 2 {
		t.Errorf("Library did not return the correct number of keys")
	}

	if reflect.TypeOf(libraryKeys).String() != "[]string" {
		t.Errorf("Library did not return the correct datatype")
	}
}

func TestGetLibraryKeysNoKeys(t *testing.T) {
	library := New("testing")

	libraryKeys := library.LibraryKeys()

	if libraryKeys != nil {
		t.Errorf("Returned the incorrect value when there are no library keys")
	}
}

func TestGetRuleData(t *testing.T) {
	ruleData1 := NewRuleData()

	library := New("testing")
	library.AddRuleData("ruleData1", &ruleData1)

	newRuleData, err := library.RuleData("ruleData1")
	if err != nil {
		t.Errorf("Library did not return any rule data when it should have")
	}

	if reflect.TypeOf(newRuleData).String() != "models.RuleData" {
		t.Errorf("Library did not return the proper object")
	}
}

func TestGetRuleDataNoKeys(t *testing.T) {
	library := New("testing")

	_, err := library.RuleData("testingData")

	if err.Error() != "could not find key: testingData in library" {
		t.Errorf("Error improperly returned when testing getRuleData without a proper key")
	}
}

func TestAddRuleData(t *testing.T) {
	ruleData1 := NewRuleData()

	library := New("testing")
	err := library.AddRuleData("ruleData1", &ruleData1)

	if err != nil {
		t.Errorf("Adding rule data returned an error when it shouldn't have")
	}
}

func TestAddRuleDataOverwrite(t *testing.T) {
	ruleData1 := NewRuleData()
	ruleData2 := NewRuleData()

	library := New("testing")
	library.AddRuleData("ruleData1", &ruleData1)
	err := library.AddRuleData("ruleData1", &ruleData2)

	if err.Error() != "could not set ruledata for key: ruleData1 - data is readonly" {
		t.Errorf("Error improperly returned when overwriting data")
	}
}
