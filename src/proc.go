package dirp

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// execWith runs a command, passes through data and returns stdout to caller
// src: https://github.com/junegunn/fzf/issues/1270#issuecomment-504000372
func execWith(data io.Reader, command []string) (string, error) {
	var result strings.Builder
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = &result
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(stdin, data)
	//_, err = data.WriteTo(stdin)
	if err != nil {
		return "", err
	}
	err = stdin.Close()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil
}

// Fuzz handles fzf: it does not want to expose itself as library
//
// -0 and -1, short for --select-{1,0} are scripting options that
// automatically selects for {0,1} matches.
//
// ref: https://github.com/junegunn/fzf/issues/2097#issuecomment-650682010
// src: https://github.com/junegunn/fzf/issues/1270#issuecomment-504000372
func Fuzz(data io.Reader) (string, error) {
	return execWith(data, []string{"fzf", "-0", "-1"})
}
