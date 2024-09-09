package io

import (
	"fmt"
	"os"
)

func Delete(path string) error {
	// Get file information to determine if it's a file or directory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("error accessing path %s: %w", path, err)
	}

	// Check if it's a directory
	if fileInfo.IsDir() {
		err = os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("failed to delete directory %s: %w", path, err)
		}
		fmt.Printf("Deleted directory: %s\n", path)
	} else {
		// Remove file
		err = os.Remove(path)
		if err != nil {
			return fmt.Errorf("failed to delete file %s: %w", path, err)
		}
		fmt.Printf("Deleted file: %s\n", path)
	}

	return nil
}
