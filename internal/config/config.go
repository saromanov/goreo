package config

// Config defines configuration for builders
type Config struct {
	build *Build
}

// Build defines configuration for the build
type Build struct {
	Archs     []string `yaml:"archs"`
	Platforms []string `yaml:"platforms"`
}
