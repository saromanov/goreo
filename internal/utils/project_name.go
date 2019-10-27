package utils

import (
	"fmt"
	"os"
	"strings"
)

// GetProjectName retruns name of the project
// as directory
func GetProjectName() string {
	dirPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("unable to get directory path: %v", err))
	}
	splitDirs := strings.Split(dirPath, "/")
	if len(splitDirs) > 0 {
		dirPath = splitDirs[len(splitDirs)-1]
	}

	return dirPath
}
