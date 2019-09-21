package builder

import (
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

// Build provides building of the project
func Build() error {
	err := exec.Command("go", "build", "-o", filepath.Dir(".")).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute go build")
	}
	return nil
}
