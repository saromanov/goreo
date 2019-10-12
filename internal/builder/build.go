package builder

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/config"
)

// Run provides building of the project
// It output built paths to all binaries
func Run(c *config.Build) ([]string, error) {
	archs := []string{"linux", "windows"}
	platforms := []string{"amd64"}

	if c != nil && len(c.Archs) > 0 {
		archs = c.Archs
	}
	if c != nil && len(c.Platforms) > 0 {
		platforms = c.Platforms
	}

	names := []string{}
	for _, a := range archs {
		for _, p := range platforms {
			name, err := buildToArch(c.Name, a, p)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("unable to build %s to the platform %s", a, p))
			}
			names = append(names, name)
		}
	}
	return names, nil
}

// buildToArch provides building of the go package to the specific platform
func buildToArch(projectName, osName, platformName string) (string, error) {
	os.Setenv("GOOS", osName)
	os.Setenv("GOARCH", platformName)
	name, err := getProjectName(projectName)
	if err != nil {
		return "", errors.Wrap(err, "unable to get project name")
	}
	binaryName := fmt.Sprintf("%s_%s_%s", name, osName, platformName)
	err = exec.Command("go", "build", "-o", binaryName).Run()
	if err != nil {
		return "", errors.Wrap(err, "unable to execute go build")
	}
	return binaryName, nil
}

// if project name is not defined, then
// get name of the working directory
func getProjectName(projectName string) (string, error) {
	if projectName != "" {
		return projectName, nil
	}

	dirPath, err := os.Getwd()
	if err != nil {
		return "", errors.Wrap(err, "unable to get directory name")
	}
	splitDirs := strings.Split(dirPath, "/")
	if len(splitDirs) > 0 {
		dirPath = splitDirs[len(splitDirs)-1]
	}

	return dirPath, nil
}
