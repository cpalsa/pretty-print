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

	flag.BoolVar(&opts.clipboard, "c", false, "Clipboard mode. Output will also be copied to the clipboard. If stdin is connected to terminal then the current clipboard contents are formatted.")
	flag.BoolVar(&opts.notifications, "n", false, "Desktop notifications. Push desktop notifications on success/failure.")
	flag.IntVar(&opts.indent, "i", 2, "Indent. Number of spaces that each depth level should be formatted with.")
	flag.Parse()

	return opts
}

func getInput() (ipt input, err error) {
	ipt = input{options: getOptions()}

	// check if clipboard functionality supported
	err = clipboard.Init()
	if ipt.options.clipboard && err != nil {
		return ipt, err
	}

	// clipboard mode active and stdin connected to terminal
	if ipt.options.clipboard && !isInputFromPipe() {
		ipt.data = clipboard.Read(clipboard.FmtText)
	} else {
		// otherwise, just read what ever stdin points to
		ipt.data, err = io.ReadAll(os.Stdin)
	}

	// trim input data and assign format
	if err == nil {
		ipt.data = []byte(strings.TrimSpace(string(ipt.data)))
		ipt.format, err = getDataFormat(ipt.data)
	}

	return ipt, err
}
