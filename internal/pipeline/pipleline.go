package pipeline

import (
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
		fileName := filepath.Base(name)
		if err := archive.Run("./", name, fileName); err != nil {
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
	if err := archive.Run("./", name, fileName); err != nil {
		return errors.Wrap(err, "unable to archive files")
	}

	return nil
}
