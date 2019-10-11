package git

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Publish provodes publishing of the repo with tags
func Publish(tag string) error {
	err := exec.Command("git", "push", "--follow-tags", tag).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute git push")
	}
	return nil
}
