package main

import (
	"fmt"
	"os"

	"github.com/0xAX/notificator"
	"golang.design/x/clipboard"
)

func main() {
	notify := notificator.New(notificator.Options{
		AppName: "Pretty Print",
	})

	input, err := getInput()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if input.options.notifications {
			notify.Push("Pretty Print: Error", fmt.Sprint(err), "", notificator.UR_CRITICAL)
		}
		os.Exit(1)
	}

	formatted, err := format(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if input.options.notifications {
			notify.Push("Pretty Print: Error", fmt.Sprint(err), "", notificator.UR_CRITICAL)
		}
		os.Exit(1)
	}

	if input.options.clipboard {
		clipboard.Write(clipboard.FmtText, formatted)
	}

	fmt.Print(string(formatted))

	if input.options.notifications {
		notify.Push("Pretty Print: Success", fmt.Sprintf("Data formatted as %v", input.format), "", notificator.UR_NORMAL)
	}

	os.Exit(0)
}
