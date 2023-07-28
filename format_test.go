package main

import (
	"os"
	"strings"
	"testing"
)

func TestIsXml(t *testing.T) {
	// isXml == true
	testFile := "./testdata/unformatted-xml.xml"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval := isXml(strings.TrimSpace(string(data)))
	if tval != true {
		t.Errorf("Expected %v to contain valid XML but isXml returned %v", testFile, tval)
	}

	// isXml == false
	testFile = "./testdata/unformatted-json.json"
	data, err = os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval = isXml(strings.TrimSpace(string(data)))
	if tval != false {
		t.Errorf("Expected %v to contain invalid XML but isXml returned %v", testFile, tval)
	}

	// isXml == false
	testFile = "./testdata/unformatted-json-array.json"
	data, err = os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval = isXml(strings.TrimSpace(string(data)))
	if tval != false {
		t.Errorf("Expected %v to contain invalid XML but isXml returned %v", testFile, tval)
	}
}

func TestIsJson(t *testing.T) {
	// isJson == false
	testFile := "./testdata/unformatted-xml.xml"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval := isJson((string(data)))
	if tval != false {
		t.Errorf("Expected %v to contain invalid JSON but isJson returned %v", testFile, tval)
	}

	// isJson == true
	testFile = "./testdata/unformatted-json.json"
	data, err = os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval = isJson((string(data)))
	if tval != true {
		t.Errorf("Expected %v to contain invalid XML but isJson returned %v", testFile, tval)
	}

	// isJson == true
	testFile = "./testdata/unformatted-json-array.json"
	data, err = os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}
	tval = isJson((string(data)))
	if tval != true {
		t.Errorf("Expected %v to contain invalid XML but isJson returned %v", testFile, tval)
	}
}

func TestFormatXml(t *testing.T) {
	testFile := "./testdata/unformatted-xml.xml"
	testData, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}

	controlFile := "./testdata/formatted-xml.xml"
	controlData, err := os.ReadFile(controlFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", controlFile)
	}

	formattedTestData, err := formatXml(string(testData))
	if err != nil {
		t.Errorf("formatting xml failed: %v", err)
	}

	if len(formattedTestData) != len(controlData) {
		t.Errorf("formatted xml is a different length than control data: %v vs %v", len(formattedTestData), len(controlData))
	}

	if string(formattedTestData) != string(controlData) {
		t.Errorf("formatted xml does not match control data: %v vs %v", len(formattedTestData), len(controlData))
	}
}

func TestFormatJson(t *testing.T) {
	// top level object case
	testFile := "./testdata/unformatted-json.json"
	testData, err := os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}

	controlFile := "./testdata/formatted-json.json"
	controlData, err := os.ReadFile(controlFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", controlFile)
	}

	formattedTestData, err := formatJson(string(testData))
	if err != nil {
		t.Errorf("formatting json failed: %v", err)
	}

	if len(formattedTestData) != len(controlData) {
		t.Errorf("formatted json is a different length than control data: %v vs %v", len(formattedTestData), len(controlData))
	}

	if string(formattedTestData) != string(controlData) {
		t.Errorf("formatted json does not match control data: %v vs %v", len(formattedTestData), len(controlData))
	}

	// array test case
	testFile = "./testdata/unformatted-json-array.json"
	testData, err = os.ReadFile(testFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFile)
	}

	controlFile = "./testdata/formatted-json-array.json"
	controlData, err = os.ReadFile(controlFile)
	if err != nil {
		t.Errorf("failed loading test data from %v", controlFile)
	}

	formattedTestData, err = formatJson(string(testData))
	if err != nil {
		t.Errorf("formatting json array failed: %v", err)
	}

	if len(formattedTestData) != len(controlData) {
		t.Errorf("formatted json array is a different length than control data: %v vs %v", len(formattedTestData), len(controlData))
	}

	if string(formattedTestData) != string(controlData) {
		t.Errorf("formatted json array does not match control data: %v vs %v", len(formattedTestData), len(controlData))
	}
}
