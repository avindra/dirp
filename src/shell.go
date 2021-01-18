package dir

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// IsDir tests to see if a path is a dir or not
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// ExecWired runs a command, passes through data and returns stdout to caller
// src: https://github.com/junegunn/fzf/issues/1270#issuecomment-504000372
func ExecWired(data io.Reader, phrase string) (string, error) {
	var result strings.Builder
	cmd := exec.Command(phrase)
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
// ref: https://github.com/junegunn/fzf/issues/2097#issuecomment-650682010
// src: https://github.com/junegunn/fzf/issues/1270#issuecomment-504000372
func Fuzz(data io.Reader) (string, error) {
	return ExecWired(data, "fzf")
}
