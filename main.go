package main

import (
	"os"

	dirp "github.com/avindra/dirp/src"
)

func main() {
	args := os.Args
	args = args[1:]
	// compat with go run . --
	if args[0] == "--" {
		args = args[1:]
	}

	if len(args) > 0 {
		arg0 := args[0]
		if dirp.IsDir(arg0) {
			cfg := dirp.FindDirs(arg0)
			dirp.Selector(cfg)
		} else if arg0 == "hook" {
			if len(args) >= 2 && args[1] == "bash" {
				dirp.PrintBashHook()
			} else {
				dirp.PrintHook() // fish
			}
		}
	} else {
		if dirp.InputHasData() {
			dirp.Feeder()
			return
		}
	}
}
