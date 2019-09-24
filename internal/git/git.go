package git

type Git struct {
	Path string
	Release string
}

func Publish(tag string) error {
	err := exec.Command("git", "push", "origin", tag).Run()
	if err != nil {
		return errors.Wrap(err, "unable to execute git push")
	}
	return nil
}