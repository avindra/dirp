package dirp

import (
	"testing"
)

func TestIsDir(t *testing.T) {
	// Test with a valid directory
	dir := "/tmp"
	if !IsDir(dir) {
		t.Errorf("Expected %s to be a directory", dir)
	}

	// Test with a non-directory path
	file := "/proc/self/exe"
	if IsDir(file) {
		t.Errorf("Expected %s to not be a directory", file)
	}

	// Test with a non-existent path
	nonExistent := "/tmp/nonexistentdir"
	if IsDir(nonExistent) {
		t.Errorf("Expected %s to not be a directory", nonExistent)
	}
}

func TestFindDirs(t *testing.T) {
	// Test with a valid directory
	dir := "/tmp"
	cfg := FindDirs(dir)
	if cfg == nil {
		t.Errorf("Expected to find directories in %s", dir)
	} else {
		for _, path := range cfg {
			if !IsDir(path) {
				t.Errorf("Expected %s to be a directory", path)
			}
		}
	}

	// Test with a non-existent directory
	nonExistent := "/tmp/nonexistentdir"
	cfg = FindDirs(nonExistent)
	if cfg != nil {
		t.Errorf("Expected to not find directories in %s", nonExistent)
	}
}
