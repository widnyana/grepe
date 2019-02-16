package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	// Mongo config schema
	Mongo struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Options  string `yaml:"options"`
	}

	//Config of Mongo
	Config struct {
		Mongo Mongo `yaml:"mongo"`
	}
)

var (
	// Conf ...
	Conf Config
)

// FromFile read file and pass it to parseConfig
func FromFile(filepath string) error {
	handle, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(fmt.Sprintf("Could not read file: %s", filepath))
	}

	err = parseConfig(&Conf, handle)
	return err
}

func parseConfig(c *Config, entry []byte) error {
	err := yaml.Unmarshal(entry, &c)
	if err != nil {
		return err
	}

	return nil
}
