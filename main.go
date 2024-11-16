package main

import (
	"os"

	"github.com/themartes/venom/parsers"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Kind     string   `yaml:"kind"`
	Required []string `yaml:"required"` // list of req pckgs
	Settings Settings
}

type Settings struct {
	Git []string `yaml:"git,omitempty"`
}

func main() {
	config := loadConfig()
	requiredCheck(config.Required)

	// TODO: skip if empty
	parsers.ParseGit(config.Settings.Git)
}

func loadConfig() *Config {
	// TODO: change to default $HOME/.venom.yaml or cli arg
	file, err := os.ReadFile("./.dev/sample.yaml")
	if err != nil {
		panic(err)
	}

	var config *Config
	// TODO: this should be parsed either from url or local cfg provided by cli flags
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return config
}
