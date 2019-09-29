package pipeline

import (
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
	names, err := builder.Run(nil)
	if err != nil {
		return errors.Wrap(err, "unable to apply build")
	}

	for _, name := range names {
		if err := archive.Run("./", name, "release.zip"); err != nil {
			return errors.Wrap(err, "unable to archive files")
		}

	}

	return nil
}
