package config

// Config defines configuration for builders
type Config struct {
	build   *Build
	publish *Publish
}

func (c *Config) GetBuild() *Build {
	if c == nil {
		return &Build{}
	}
	if c.build == nil {
		return &Build{}
	}
	return c.build
}

func (c *Config) GetPublish() *Publish {
	if c.publish == nil {
		return &Publish{}
	}
	return c.publish
}

// Build defines configuration for the build
type Build struct {
	Name      string   `yaml:"name"`
	Archs     []string `yaml:"archs"`
	Platforms []string `yaml:"platforms"`
	Version   string   `yaml:"version"`
}

type Publish struct {
}
