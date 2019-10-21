package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
		Name: getProjectName(),
	}
}

// if project name is not defined, then
// get name of the working directory
func getProjectName() string {
	dirPath, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("unable to get directory path: %v", err))
	}
	splitDirs := strings.Split(dirPath, "/")
	if len(splitDirs) > 0 {
		dirPath = splitDirs[len(splitDirs)-1]
	}

	return dirPath
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
