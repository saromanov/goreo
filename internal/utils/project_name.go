package utils

import (
	"fmt"
	"os"
	"strings"
)

// GetProjectName retruns name of the project from directory
func GetProjectName() (string, error) {
	dirPath, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("unable to get directory path: %v", err)
	}
	splitDirs := strings.Split(dirPath, "/")
	if len(splitDirs) > 0 {
		dirPath = splitDirs[len(splitDirs)-1]
	}

	return dirPath, nil
}
