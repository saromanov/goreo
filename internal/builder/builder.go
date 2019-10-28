package builder

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/config"
	"github.com/saromanov/goreo/internal/git"
	"github.com/saromanov/goreo/internal/template"
)

// Run provides building of the project
// It output built paths to all binaries
func Run(c *config.Build) ([]string, error) {
	setEnvironmentVariables(c.Envs)

	names := []string{}
	for _, a := range c.Archs {
		if a == "arm" && len(c.Goarm) > 0 {
			os.Setenv("GOOS", c.Goarm[0])
		}
		for _, p := range c.Platforms {
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
	resultName, err := template.GetName(projectName, osName)
	if err != nil {
		return "", errors.Wrap(err, "unable to execute template")
	}
	binaryName, err := createProjectName(resultName, platformName, snapshot)
	if err != nil {
		return "", errors.Wrap(err, "unable to create project name")
	}
	err = exec.Command("go", "build", "-o", binaryName).Run()
	if err != nil {
		return "", errors.Wrap(err, "unable to execute go build")
	}
	return binaryName, nil
}

func createProjectName(projectName, platformName string, snapshot bool) (string, error) {
	binaryName := fmt.Sprintf("%s_%s", projectName, platformName)
	if snapshot {
		commit, err := git.GetLastCommitID()
		if err != nil {
			return "", errors.Wrap(err, "unable to get last commit id")
		}
		binaryName = fmt.Sprintf("%s_%s_%s", projectName, commit, platformName)
	}

	return binaryName, nil
}

// setting list of environment variables before the build
func setEnvironmentVariables(vars map[string]interface{}) {
	for k, v := range vars {
		os.Setenv(k, v.(string))
	}
}
