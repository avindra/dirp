package main

import (
	"fmt"
	"os"

	dirp "github.com/avindra/dirp/src"
)

func main() {
	args := os.Args
	args = args[1:]

	if len(args) > 0 {
		// compat with go run . --
		if args[0] == "--" {
			args = args[1:]
			if len(args) == 0 {
				goto NOARGS
			}
		}

		arg0 := args[0]
		if dirp.IsDir(arg0) {
			cfg := dirp.FindDirs(arg0)
			dirp.Selector(cfg)
		} else if arg0 == "hook" {
			if len(args) >= 2 {
				switch args[1] {
				case "fish":
					dirp.PrintFishHook()
				case "rc":
					dirp.PrintRcHook()
				default:
					panic("I don't know about " + args[1])
				}
			} else {
				dirp.PrintHook()
			}
		} else if arg0 == "cfg" {
			fmt.Print(dirp.GetConfigPath())
			os.Exit(2)
		}
		return
	}

NOARGS:
	var cfg dirp.ConfigSelection
	if dirp.InputHasData() {
		cfg = dirp.ReadConfig(os.Stdin)
	} else {
		f, err := os.Open(dirp.GetConfigPath())
		if err == nil {
			cfg = dirp.ReadConfig(f)
		}
	}

	if len(cfg) > 0 {
		dirp.Selector(cfg)
	}
}
