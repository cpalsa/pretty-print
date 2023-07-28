package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
)

func notify(msg string) {
	cmd := exec.Command("notify-send", "Pretty Print", msg)
	cmd.Run()
}

func exit(msg string, code int) {
	notify(msg)
	os.Exit(code)
}

func main() {
	input, err := clipboard.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Could not read from clipboard: %v", err), 1)
	}

	formatted, err := format(input)
	if err != nil {
		exit(fmt.Sprintf("Could not format clipboard content: %v", err), 1)
	}

	err = clipboard.WriteAll(string(formatted))
	if err != nil {
		exit(fmt.Sprintf("Could not write to clipboard: %v", err), 1)
	}

	format := "XML"
	if isJson(input) {
		format = "JSON"
	}

	notify(fmt.Sprintf("Clipboard formatted as %v!", format))
}
