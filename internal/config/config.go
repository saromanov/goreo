package config

// Config defines configuration for builders
type Config struct {
	build *Build
}

func (c *Config) GetBuild() *Build {
	if c.build == nil {
		return &Build{}
	}
	return c.build
}

// Build defines configuration for the build
type Build struct {
	Archs     []string `yaml:"archs"`
	Platforms []string `yaml:"platforms"`
	Version   string   `yaml:"version"`
}
