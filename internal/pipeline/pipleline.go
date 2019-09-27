package pipeline

import "github.com/saromanov/goreo/internal/config"

type Pipleline struct {
	conf *config.Config
}

// New initialize new pipleline
func New(c *config.Config) *Pipleline {
	return &Pipleline{
		conf: c,
	}
}
