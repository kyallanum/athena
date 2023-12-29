package models

import (
	"reflect"
	"testing"
)

func TestCreateNewCount(t *testing.T) {
	testCountOperation := Count.New(Count{}, "test")

	if testCountOperation.GetOperation() != "count" {
		t.Errorf("Operation type is not correct for count.")
	}

	if reflect.TypeOf(testCountOperation).String() != "*models.Count" {
		t.Errorf("Creating new operation did not return the appropriate type.")
	}
}

func TestCountCalculateOperation(t *testing.T) {
	st_data1 := SearchTermData.New(SearchTermData{})
	st_data1.AddValue("test", "testing 1")

	st_data2 := SearchTermData.New(SearchTermData{})
	st_data2.AddValue("test", "testing 2")

	testRuleData := RuleData.New(RuleData{})
	testRuleData.AppendSearchTermData(st_data1)
	testRuleData.AppendSearchTermData(st_data2)

	testCountOperation := Count.New(Count{}, "test")

	numTests, err := testCountOperation.CalculateOperation(testRuleData)

	if err != nil {
		t.Errorf("An error occurred when calculating for operation: count %v", err)
	}

	if numTests[0] != "2" {
		t.Errorf("An issue occurred when calculating for operation: count")
	}

	if reflect.TypeOf(numTests).String() != "[]string" {
		t.Errorf("Calculate Operation returned the wrong value type")
	}
}

func TestCountCalculateOperationEmptyRuleData(t *testing.T) {
	testRuleData := RuleData.New(RuleData{})
	testCountOperation := Count.New(Count{}, "test")
	_, err := testCountOperation.CalculateOperation(testRuleData)

	if err.Error() != "rule does not have any search term data stored" {
		t.Errorf("Error not returned properly for empty RuleData")
	}
}

func TestCountCalculateOperationSearchTermNotExist(t *testing.T) {
	st_data1 := SearchTermData.New(SearchTermData{})
	st_data1.AddValue("bad_test", "testing")

	testRuleData := RuleData.New(RuleData{})
	testRuleData.AppendSearchTermData(st_data1)

	testCountOperation := Count.New(Count{}, "test")
	_, err := testCountOperation.CalculateOperation(testRuleData)
	if err.Error() != "this value was never extracted during search phase" {
		t.Errorf("Error not returned properly for search term data that was never created")
	}
}
