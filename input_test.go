package main

import (
	"os"
	"testing"
)

// TODO: Figure out how to test this
func TestIsInputFromPipe(t *testing.T) {}

func TestGetDataFormat(t *testing.T) {
	testFileXml := "./testdata/unformatted.xml"
	testDataXml, err := os.ReadFile(testFileXml)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFileXml)
	}
	format, err := getDataFormat(testDataXml)
	if err != nil {
		t.Errorf("failed to determine data format as xml: %v", err)
	}
	if format != "xml" {
		t.Errorf("expected data format to be 'xml' but got %v", format)
	}

	testFileJson := "./testdata/unformatted.json"
	testDataJson, err := os.ReadFile(testFileJson)
	if err != nil {
		t.Errorf("failed loading test data from %v", testFileJson)
	}
	format, err = getDataFormat(testDataJson)
	if err != nil {
		t.Errorf("failed to determine data format as json: %v", err)
	}
	if format != "json" {
		t.Errorf("expected data format to be 'json' but got %v", format)
	}

	testDataInvalid := "Hello, World!"
	format, err = getDataFormat([]byte(testDataInvalid))
	if err == nil {
		t.Errorf("expected an error, but got data format %v", format)
	}
}

// TODO: Figure out how to test this
func TestGetOptions(t *testing.T) {}

// TODO: Figure out how to test this
func TestGetInput(t *testing.T) {}
