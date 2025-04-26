package dirp

import (
	"embed"
)

//go:embed hooks/*
var folder embed.FS

// PrintHook emits shell code for Bash, ZSH, sh, BusyBox, etc
func PrintHook() {
	hook, _ := folder.ReadFile("hooks/hook.sh")
	print(string(hook))
}

// PrintFishHook emits shell code for Fish
func PrintFishHook() {
	hook, _ := folder.ReadFile("hooks/hook.fish")
	print(string(hook))
}

// PrintRcHook emits code for rc, the plan 9 shell
func PrintRcHook() {
	hook, _ := folder.ReadFile("hooks/hook.rc")
	print(string(hook))
}

// PrintEsHook emits code for es, a shell based on rc
func PrintEsHook() {
	hook, _ := folder.ReadFile("hooks/hook.es")
	print(string(hook))
}
