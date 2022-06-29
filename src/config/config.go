package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var Config BotConfig

func LoadConfig() error {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &Config)
}

type Guild struct {
	Name   string `yaml:"name"`
	Id     string `yaml:"id"`
	Prefix string `yaml:"prefix"`
}

type BotConfig struct {
	Guilds []Guild `yaml:"guilds"`
}
