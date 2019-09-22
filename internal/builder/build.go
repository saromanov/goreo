package builder

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

// Build provides building of the project
func Build() error {
	if err := buildToArch("linux", "amd64"); err != nil {
		return err
	}
	if err := buildToArch("windows", "amd64"); err != nil {
		return err
	}
	return nil
}

// buildToArch provides building of the go package to the specific platform
func buildToArch(osName, platformName string) error {
	os.Setenv("GOOS", osName)
	os.Setenv("GOARCH", platformName)
	err := exec.Command("go", "build", "-o", filepath.Dir(".")).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute go build")
	}
	return nil
}
