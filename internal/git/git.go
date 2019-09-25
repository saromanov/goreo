package git

import (
	"os/exec"

	"github.com/pkg/errors"
)

func Publish(tag string) error {
	err := exec.Command("git", "push", "origin", tag).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute git push")
	}
	return nil
}
