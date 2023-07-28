package main

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/beevik/etree"
)

func format(input string) (formatted []byte, err error) {
	if isXml(input) {
		formatted, err = formatXml(input)
	}

	if isJson(input) {
		formatted, err = formatJson(input)
	}

	if len(formatted) == 0 {
		err = errors.New("invalid input")
	}

	return formatted, err
}

func formatXml(input string) (formatted []byte, err error) {
	// etree will just echo the input if it is not XML instead of erroring, so lets fix that here
	if trimmed := strings.TrimSpace(input); !strings.HasPrefix(trimmed, "<") {
		err = errors.New("input is not xml")
		return
	}

	// I wish I could have just used encoding/xml -- but I don't know the data ahead of time and can't unmarshal into empty interface.
	// Maybe there is another way, but I'm unsure at this time.
	doc := etree.NewDocument()
	err = doc.ReadFromString(input)

	if err == nil {
		doc.Indent(2)
		tmp, err := doc.WriteToString()

		// remove extra new line added by potential stream
		if err == nil {
			tmp = strings.TrimSpace(tmp)
			formatted = []byte(tmp)
		}
	}

	return formatted, err
}

func formatJson(input string) (formatted []byte, err error) {
	var data interface{}
	err = json.Unmarshal([]byte(input), &data)

	if err != nil {
		return
	}

	formatted, err = json.MarshalIndent(data, "", "  ")

	return formatted, err
}

func isXml(input string) bool {
	return strings.HasPrefix(strings.TrimSpace(input), "<")
}
func isJson(input string) bool {
	trimmed := strings.TrimSpace(input)
	return strings.HasPrefix(trimmed, "{") || strings.HasPrefix(trimmed, "[")
}
