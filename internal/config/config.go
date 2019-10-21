package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config defines configuration for builders
type Config struct {
	Build   *Build   `yaml:"build"`
	Publish *Publish `yaml:"publish"`
	Archive *Archive `yaml:"archive"`
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
	fmt.Println("CONF: ", conf)
	return conf, nil
}

func (c *Config) GetBuild() *Build {
	if c == nil {
		return &Build{}
	}
	if c.Build == nil {
		return &Build{}
	}
	return c.Build
}

func (c *Config) GetPublish() *Publish {
	if c.Publish == nil {
		return &Publish{}
	}
	return c.Publish
}

func (c *Config) GetArchive() *Archive {
	if c.Archive == nil {
		return &Archive{}
	}
	return c.Archive
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
}

type Publish struct {
	Description string `yaml:"description"`
}

type Archive struct {
	Files []string `yaml:"files"`
	Name  string   `yaml:"name"`
}
