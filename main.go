package main

import (
	"fmt"
	"os"

	"github.com/0xAX/notificator"
	"golang.design/x/clipboard"
)

func notify(msg string) {
	notify := notificator.New(notificator.Options{
		AppName: "Pretty Print",
	})

	//"icon/default.png"
	notify.Push("Formatting Result", msg, "", notificator.UR_NORMAL)
}

func main() {
	input, err := getInput()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if input.options.notifications {
			notify(fmt.Sprint(err))
		}
		os.Exit(1)
	}

	formatted, err := format(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if input.options.notifications {
			notify(fmt.Sprint(err))
		}
		os.Exit(1)
	}

	if input.options.clipboard {
		clipboard.Write(clipboard.FmtText, formatted)
	}

	fmt.Println(string(formatted))

	if input.options.notifications {
		notify(fmt.Sprintf("Data formatted as %v", input.format))
	}

	os.Exit(0)
}
