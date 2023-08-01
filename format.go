package main

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/beevik/etree"
)

func format(ipt input) (formatted []byte, err error) {
	if ipt.format == "xml" {
		return formatXml(ipt)
	}

	if ipt.format == "json" {
		return formatJson(ipt)
	}

	return formatted, errors.New("unsupported input format")
}

func formatXml(ipt input) (formatted []byte, err error) {
	// I wish I could have just used encoding/xml -- but I don't know the data ahead of time and can't unmarshal into empty interface.
	// Maybe there is another way, but I'm unsure at this time.
	doc := etree.NewDocument()
	err = doc.ReadFromBytes(ipt.data)

	if err != nil {
		return nil, err
	}

	doc.Indent(ipt.options.indent)

	return doc.WriteToBytes()
}

func formatJson(ipt input) (formatted []byte, err error) {
	var buf interface{}
	err = json.Unmarshal(ipt.data, &buf)

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(buf, "", strings.Repeat(" ", ipt.options.indent))
}
