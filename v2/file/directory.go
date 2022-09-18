package file

import (
	"fmt"
	"os"
)

// CreateDirIfNotExists creates a directory, and all its parents, if it
// does not exist.
func CreateDirIfNotExists(path string) error {
	const mode = 0755

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, mode)
		if err != nil {
			return fmt.Errorf("CreateDirIfNotExists(%q): %w: %v", path, ErrCreateDir, err)
		}
	}

	return nil
}
