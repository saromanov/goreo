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
	if err := builder.Run(nil); err != nil {
		return errors.Wrap(err, "unable to apply build")
	}
	
	if err := archive.Run("./"); err != nil {
		return errors.Wrap(err, "unable to archive files")
	}

	return nil
}
