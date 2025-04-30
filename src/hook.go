package dirp

import (
	"embed"
	"os"
)

//go:embed hooks/*
var folder embed.FS

func PrintShellHook(shell string) {
	ext := shell
	if ext == "pwsh" {
		ext = "ps1"
	}

	switch shell {
	case "fish", "pwsh", "es", "rc", "sh":
		profile := "hook." + ext
		hook, _ := folder.ReadFile("hooks/" + profile)
		os.Stdout.Write(hook)
	default:
		panic("I don't know about " + shell)
	}
}
