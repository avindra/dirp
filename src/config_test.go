package dirp

import (
	"strings"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cfg := ReadConfig(strings.NewReader(`
rootfs|/
temp fs|/tmp
Downloads|/home/dolores/Downloads
`))

	if len(cfg) != 3 {
		t.Error("Not enough entries")
	}

	if cfg["temp fs"] != "/tmp" {
		t.Error("Corrupted parse")
	}

	if cfg["Downloads"] != "/home/dolores/Downloads" {
		t.Error("Corrupted parse")
	}
}

func TestConfigComments(t *testing.T) {
	cfg := ReadConfig(strings.NewReader(`
rootfs|/
sys library|/Library
#temp fs|/tmp
`))

	if len(cfg) > 2 {
		t.Error("Too many entries")
	}
}

// commented lines MUST begin with a #
func TestConfigCommentLimitation(t *testing.T) {
	cfg := ReadConfig(strings.NewReader(`
rootfs|/
sys library|/Library
 #temp fs|/tmp
`))

	if len(cfg) != 3 {
		t.Error("Too many entries")
	}

	if cfg[" #temp fs"] != "/tmp" {
		t.Error("Boo")
	}
}
