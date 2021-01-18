package main

import (
	"fmt"
	"os"
	"strings"

	dir "github.com/avindra/dir/src"
)

func main() {
	args := os.Args
	args = args[1:]
	// compat with go run . --
	if args[0] == "--" {
		args = args[1:]
	}

	switch len(args) {
	case 1:
		arg0 := args[0]
		if dir.IsDir(arg0) {
			routine := "find " + arg0 + " -maxdepth 1 -type d"
			target, err := dir.ExecWired(strings.NewReader(""), routine)
			if err != nil {
				fmt.Println("tgt", target, routine)
			} else {
				fmt.Print("tgterr", err)
			}
		} else if arg0 == "hook" {
			dir.PrintHook()
		}
	default:
		if dir.InputHasData() {
			dir.Feeder()
			return
		}
	}
}
