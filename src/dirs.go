package dirp

import (
	"os"
	"path/filepath"
)

// IsDir tests to see if a path is a dir or not
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// FindDirs at path. Used by driller
func FindDirs(path string) ConfigSelection {
	fullPath, _ := filepath.Abs(path)
	dir, err := os.Open(fullPath)
	if err != nil {
		return nil
	}

	names, _ := dir.Readdirnames(-1)
	defer dir.Close()

	cfg := make(ConfigSelection)
	for _, name := range names {
		resolvedPath := filepath.Join(fullPath, name)
		if IsDir(resolvedPath) {
			cfg[resolvedPath] = resolvedPath
		}
	}
	return cfg
}
