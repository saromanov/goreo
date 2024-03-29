package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config defines configuration for builders
type Config struct {
	// Before/After provides executing of commands before/after
	// starting of the pipeline
	Before   []string  `yaml:"before"`
	After    []string  `yaml:"after"`
	Build    *Build    `yaml:"build"`
	Publish  *Publish  `yaml:"publish"`
	Archive  *Archive  `yaml:"archive"`
	Checksum *Checksum `yaml:"checksum"`
}

// Unmarshal provides unmarshaling of the config
func Unmarshal(path string) (*Config, error) {
	var conf *Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

// GetBuild returns config for build or default
func (c *Config) GetBuild() *Build {
	if c == nil {
		return makeDefaultBuild()
	}
	if c.Build == nil {
		return makeDefaultBuild()
	}
	return c.Build
}

func makeDefaultBuild() *Build {
	return &Build{
		Name:      "project",
		Archs:     []string{"linux", "windows"},
		Platforms: []string{"amd64"},
		Goarm:     []string{"6"},
	}
}

func (c *Config) GetPublish() *Publish {
	if c.Publish == nil {
		return &Publish{}
	}
	return c.Publish
}

func (c *Config) GetArchive() *Archive {
	if c.Archive == nil {
		return &Archive{
			Path: "./",
		}
	}
	return c.Archive
}

func (c *Config) GetChecksum() *Checksum {
	if c.Checksum == nil {
		return &Checksum{}
	}

	return c.Checksum
}

// Build defines configuration for the build
type Build struct {
	Name      string                 `yaml:"name"`
	Envs      map[string]interface{} `yaml:"envs"`
	Snapshot  bool                   `yaml:"snapshot"`
	Archs     []string               `yaml:"archs"`
	Platforms []string               `yaml:"platforms"`
	Version   string                 `yaml:"version"`
	Goarm     []string               `yaml:"goarm"`
	Flags     []string               `yaml:"flags"`
}

type Publish struct {
	Description string `yaml:"description"`
}

type Archive struct {
	Files []string `yaml:"files"`
	Name  string   `yaml:"name"`
	Path  string   `yaml:"path"`
}

type Checksum struct {
	Name      string `yaml:"name"`
	Algorithm string `yaml:"algorithm"`
}
