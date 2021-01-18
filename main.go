package main

import (
	"os"

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
			cfg := dir.FindDirs(arg0)
			dir.Selector(cfg)
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
