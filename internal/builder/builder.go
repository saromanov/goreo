package builder

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/config"
	"github.com/saromanov/goreo/internal/git"
)

// Run provides building of the project
// It output built paths to all binaries
func Run(c *config.Build) ([]string, error) {
	archs := []string{"linux", "windows"}
	platforms := []string{"amd64"}
	goarmVersions := []string{"6"}

	setEnvironmentVariables(c.Envs)
	if c != nil && len(c.Archs) > 0 {
		archs = c.Archs
	}
	if c != nil && len(c.Platforms) > 0 {
		platforms = c.Platforms
	}
	if c != nil && len(c.Goarm) > 0 {
		goarmVersions = c.Goarm
	}
	fmt.Println(goarmVersions)

	names := []string{}
	for _, a := range archs {
		for _, p := range platforms {
			name, err := buildToArch(c.Name, a, p, c.Snapshot)
			if err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("unable to build %s to the platform %s", a, p))
			}
			names = append(names, name)
		}
	}
	return names, nil
}

// buildToArch provides building of the go package to the specific platform
func buildToArch(projectName, osName, platformName string, snapshot bool) (string, error) {
	os.Setenv("GOOS", osName)
	os.Setenv("GOARCH", platformName)
	binaryName, err := createProjectName(projectName, osName, platformName, snapshot)
	if err != nil {
		return "", errors.Wrap(err, "unable to create project name")
	}
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

func createProjectName(projectName, osName, platformName string, snapshot bool) (string, error) {
	name, err := getProjectName(projectName)
	if err != nil {
		return "", errors.Wrap(err, "unable to get project name")
	}
	binaryName := fmt.Sprintf("%s_%s_%s", name, osName, platformName)
	if snapshot {
		commit, err := git.GetLastCommitID()
		if err != nil {
			return "", err
		}
		binaryName = fmt.Sprintf("%s_%s_%s_%s", name, osName, commit, platformName)
	}

	return binaryName, nil
}

// setting list of environment variables before the build
func setEnvironmentVariables(vars map[string]interface{}) {
	for k, v := range vars {
		os.Setenv(k, v.(string))
	}
}
