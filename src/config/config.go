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
	Name    string `yaml:"name"`
	Prefix  string `yaml:"prefix"`
	Primary bool   `yaml:"primary"`
}

type BotConfig struct {
	Guilds map[string]Guild `yaml:"guilds"`
}
