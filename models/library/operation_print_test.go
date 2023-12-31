package models

import (
	"reflect"
	"testing"
)

func TestCreateNewPrint(t *testing.T) {
	testPrint := NewPrintOperation("testing")

	if testPrint.Operation() != "print" {
		t.Errorf("Returned object does not have the right operation")
	}

	if reflect.TypeOf(testPrint).String() != "*models.Print" {
		t.Errorf("Returned object not of the correct type")
	}
}

func TestPrintCalculateOperation(t *testing.T) {
	st_data1 := NewSearchTermData()
	st_data1.AddValue("testing", "test 1")

	st_data2 := NewSearchTermData()
	st_data2.AddValue("testing", "test 2")

	testRuleData := NewRuleData()
	testRuleData.AppendSearchTermData(st_data1)
	testRuleData.AppendSearchTermData(st_data2)

	testOperation := NewPrintOperation("testing")
	printOutput, err := testOperation.CalculateOperation(testRuleData)
	if err != nil {
		t.Errorf("An error occurred when attempting to resolve lines to print: \n\t%s", err.Error())
	}

	if len(printOutput) != 2 {
		t.Errorf("Returned incorrect input length for print statement")
	}

	if reflect.TypeOf(printOutput).String() != "[]string" {
		t.Errorf("Returned object of incorrect type")
	}
}

func TestPrintCalculateOperationEmptyRuleData(t *testing.T) {
	testRuleData := NewRuleData()
	testPrintOperation := NewPrintOperation("testing")

	_, err := testPrintOperation.CalculateOperation(testRuleData)
	if err.Error() != "rule does not have any search term data stored" {
		t.Errorf("error not properly returned when ruledata is empty")
	}
}

func TestPrintCalculateOperationSearchTermNotExist(t *testing.T) {
	st_data1 := NewSearchTermData()
	st_data1.AddValue("bad_test", "test 1")

	testRuleData := NewRuleData()
	testRuleData.AppendSearchTermData(st_data1)

	testPrintOperation := NewPrintOperation("testing")
	_, err := testPrintOperation.CalculateOperation(testRuleData)

	if err.Error() != "this value was never extracted during search phase" {
		t.Errorf("Error was not properly returned for search term data that was never created")
	}
}
