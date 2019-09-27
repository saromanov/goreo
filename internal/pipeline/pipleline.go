package pipeline

import (
	"github.com/pkg/errors"
	"github.com/saromanov/goreo/internal/archive"
	"github.com/saromanov/goreo/internal/builder"
	"github.com/saromanov/goreo/internal/config"
)

type Pipleline struct {
	conf *config.Config
}

// New initialize new pipleline
func New(c *config.Config) *Pipleline {
	return &Pipleline{
		conf: c,
	}
}

// Run provides executing of the builder
func Run() error {
	if err := archive.Run("."); err != nil {
		return errors.Wrap(err, "unable to archive files")
	}

	if err := builder.Run(nil); err != nil {
		return errors.Wrap(err, "unable to apply build")
	}
}
