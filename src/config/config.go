package config

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"
)

var Config BotConfig

func LoadConfig() error {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	return Config.Validate()
}

type Guild struct {
	Name    string `yaml:"name"`
	Prefix  string `yaml:"prefix"`
	Primary bool   `yaml:"primary"`
}

type BotConfig struct {
	Guilds       map[string]Guild `yaml:"guilds"`
	Levels       map[string]int   `yaml:"levels"`
	PrimaryGuild string
}

// Get the permission level of a member.
// This takes the highest available level they qualify for,
// which can be set per user, level, and guild.
func (c *BotConfig) GetLevel(member discordgo.Member) int {
	level := 0

	// Check if the user has a level set.
	userLevel, ok := c.Levels[member.User.ID]
	if ok {
		level = userLevel
	}

	// Check each role to see if there's a higher level.
	for _, role := range member.Roles {
		roleLevel, ok := c.Levels[role]
		if ok && roleLevel > level {
			level = roleLevel
		}
	}

	// Check if the guild has a level set.
	guildLevel, ok := c.Levels[member.GuildID]
	if ok && guildLevel > level {
		level = guildLevel
	}

	return level
}

func (c *BotConfig) Validate() error {
	primaryGuilds := 0

	for id, guild := range c.Guilds {
		if guild.Primary {
			primaryGuilds++
			c.PrimaryGuild = id
		}
	}

	if primaryGuilds != 1 {
		return fmt.Errorf("There must be exactly one primary guild.")
	}

	return nil
}
