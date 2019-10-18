package pipeline

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/archive"
	"github.com/saromanov/goreo/internal/builder"
	"github.com/saromanov/goreo/internal/config"
)

type Pipeline struct {
	conf *config.Config
}

// New initialize new pipleline
func New(c *config.Config) *Pipeline {
	return &Pipeline{
		conf: c,
	}
}

// Run provides executing of the builder
func (p *Pipeline) Run() error {
	names, err := builder.Run(p.conf.GetBuild())
	if err != nil {
		return errors.Wrap(err, "unable to apply build")
	}

	for _, name := range names {
		if err := makeArchive(name, p.conf); err != nil {
			return errors.Wrap(err, "unable to archive files")
		}
	}

	return nil
}

func makeArchive(name string, c *config.Config) error {
	archiveConf := c.GetArchive()
	if err := os.Mkdir(name, 777); err != nil {
		return err
	}
	fileName := filepath.Base(name)
	if len(archiveConf.Files) > 0 {
		for _, fileName := range archiveConf.Files {
			copyFile(fileName, "../")
		}
	}
	if err := archive.Run("./", name, fileName); err != nil {
		return errors.Wrap(err, "unable to archive files")
	}

	return nil
}

func copyFile(fileName, dest string) error {
	srcFile, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
