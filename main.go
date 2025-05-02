package main

import (
	"os"

	dirp "github.com/avindra/dirp/src"
)

func driller(path string) {
	cfg := dirp.FindDirs(path)
	dirp.Selector(cfg)
}

func main() {
	args := os.Args
	args = args[1:]

	if len(args) == 0 {
		handleNoArgs()
		return
	}

	arg0 := args[0]
	if dirp.IsDir(arg0) {
		driller(arg0)
	} else if arg0 == "hook" {
		if len(args) >= 2 {
			dirp.PrintShellHook(args[1])
		} else {
			dirp.PrintShellHook("sh")
		}
	} else if arg0 == "cfg" {
		os.Stdout.WriteString(dirp.GetConfigPath())
		os.Exit(2)
	} else {
		handleNoArgs()
	}
}

func handleNoArgs() {
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
	} else {
		// No config detected, default to driller
		driller(".")
	}
}
