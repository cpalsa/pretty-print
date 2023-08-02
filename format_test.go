package main

import (
	"os"
	"testing"
)

func TestFormatXml(t *testing.T) {
	testFile := "./testdata/unformatted.xml"
	testData, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}

	controlFile := "./testdata/formatted.xml"
	controlData, err := os.ReadFile(controlFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", controlFile)
	}

	input := input{
		data: testData,
		options: options{
			indent:        2,
			clipboard:     false,
			notifications: false,
		},
	}

	formatted, err := formatXml(input)
	if err != nil {
		t.Errorf("formatting xml failed: %v", err)
	}

	if len(formatted) != len(controlData) {
		t.Errorf("formatted xml is a different length than control data: %v vs %v", len(formatted), len(controlData))
	}

	if string(formatted) != string(controlData) {
		t.Errorf("formatted xml does not match control data")
	}
}

func TestFormatJson(t *testing.T) {
	testFile := "./testdata/unformatted.json"
	testData, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}

	controlFile := "./testdata/formatted.json"
	controlData, err := os.ReadFile(controlFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", controlFile)
	}

	input := input{
		data: testData,
		options: options{
			indent:        2,
			clipboard:     false,
			notifications: false,
		},
	}

	formatted, err := formatJson(input)
	if err != nil {
		t.Errorf("formatting json failed: %v", err)
	}

	if len(formatted) != len(controlData) {
		t.Errorf("formatted json is a different length than control data: %v vs %v", len(formatted), len(controlData))
	}

	if string(formatted) != string(controlData) {
		t.Errorf("formatted json does not match control data")
	}
}
