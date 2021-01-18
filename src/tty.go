package dir

import "os"

// https://stackoverflow.com/a/26567513/270302
func InputHasData() bool {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}
