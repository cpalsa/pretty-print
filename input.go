package main

import (
	"errors"
	"flag"
	"io"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

type options struct {
	indent        int
	clipboard     bool
	notifications bool
}

type input struct {
	data    []byte
	format  string
	options options
}

func isInputFromPipe() bool {
	stat, err := os.Stdin.Stat()

	if err != nil {
		return false
	}

	return (stat.Mode() & os.ModeCharDevice) == 0
}

func getDataFormat(data []byte) (format string, err error) {
	str := string(data)
	if strings.HasPrefix(str, "<") {
		format = "xml"
	}
	if strings.HasPrefix(str, "{") || strings.HasPrefix(str, "[") {
		format = "json"
	}

	if len(format) == 0 {
		err = errors.New("unsupported data format")
	}

	return format, err
}

func getOptions() (opts options) {
	opts = options{}

	flag.BoolVar(&opts.clipboard, "c", !isInputFromPipe(), "Clipboard mode. Output will also be copied to the clipboard. If stdin is connected to terminal then the current clipboard contents are formatted.")
	flag.BoolVar(&opts.notifications, "n", opts.clipboard, "Desktop notifications. Defaults to true if using clipboard mode.")
	flag.IntVar(&opts.indent, "i", 2, "Indent. Number of spaces that each depth level should be formatted with.")
	flag.Parse()

	return opts
}

func getInput() (ipt input, err error) {
	ipt = input{options: getOptions()}

	// Read data from the pipe/redirection
	if isInputFromPipe() {
		ipt.data, err = io.ReadAll(os.Stdin)
	} else if ipt.options.clipboard {
		// grab data from clipboard
		ipt.data = clipboard.Read(clipboard.FmtText)
	} else {
		err = errors.New("no data rcved from stdin and clipboard mode is disabled")
	}

	if err == nil {
		// trim input data and assign format
		ipt.data = []byte(strings.TrimSpace(string(ipt.data)))
		ipt.format, err = getDataFormat(ipt.data)
	}

	return ipt, err
}
