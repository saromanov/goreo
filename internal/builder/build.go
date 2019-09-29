package builder

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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

	for _, a := range archs {
		for _, p := range platforms {
			if err := buildToArch(a, p); err != nil {
				return nil, errors.Wrap(err, fmt.Sprintf("unable to build %s to the platform %s", a, p))
			}
		}
	}
	return []string{}, nil
}

// buildToArch provides building of the go package to the specific platform
func buildToArch(osName, platformName string) error {
	os.Setenv("GOOS", osName)
	os.Setenv("GOARCH", platformName)
	binaryName := fmt.Sprintf("%s_%s_%s.%s", filepath.Dir("."), osName, platformName)
	err := exec.Command("go", "build", "-o", binaryName).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute go build")
	}
	return nil
}
