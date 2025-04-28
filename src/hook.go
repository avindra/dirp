package dirp

import (
	"embed"
	"os"
)

//go:embed hooks/*
var folder embed.FS

// PrintHook emits shell code for Bash, ZSH, sh, BusyBox, etc
func PrintHook() {
	hook, _ := folder.ReadFile("hooks/hook.sh")
	os.Stdout.Write(hook)
}

// PrintFishHook emits shell code for Fish
func PrintFishHook() {
	hook, _ := folder.ReadFile("hooks/hook.fish")
	os.Stdout.Write(hook)
}

// PrintRcHook emits code for rc, the plan 9 shell
func PrintRcHook() {
	hook, _ := folder.ReadFile("hooks/hook.rc")
	os.Stdout.Write(hook)
}

// PrintEsHook emits code for es, a shell based on rc
func PrintEsHook() {
	hook, _ := folder.ReadFile("hooks/hook.es")
	os.Stdout.Write(hook)
}

// PrintPwshHook emits code for PowerShell
func PrintPwshHook() {
	hook, _ := folder.ReadFile("hooks/hook.ps1")
	os.Stdout.Write(hook)
}
