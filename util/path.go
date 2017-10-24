package util

import (
	"fmt"
	"os"
)

func IsGitDir(path string) error {
	// Check the path is existed
	status, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("The status of the path %s is not expected: %v", path, err)
	}
	if status.IsDir() != true {
		return fmt.Errorf("The path %s is not a directory: %v", path)
	}
	// TODO: Check if the directory is a git repo.
	return nil
}
