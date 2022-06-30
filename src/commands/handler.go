package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/vcokltfre/volcan/src/config"
	"github.com/vcokltfre/volcan/src/core"
	"github.com/vcokltfre/volcan/src/utils"
)

var Handler CommandHandler

type CommandHandler struct {
	Commands map[string]*Command
	guilds   []string
}

func Setup() error {
	guilds := []string{}

	for id := range config.Config.Guilds {
		guilds = append(guilds, id)
	}

	Handler = CommandHandler{
		Commands: map[string]*Command{},
		guilds:   guilds,
	}

	return setupEventListeners()
}

func (h *CommandHandler) Register(command *Command) error {
	if h.Commands[command.Name] != nil {
		return fmt.Errorf("Command %s is already defined.", command.Name)
	}

	h.Commands[command.Name] = command
	return nil
}

func (h *CommandHandler) ProcessMessage(message *discordgo.MessageCreate) error {
	if !utils.Contains(h.guilds, message.GuildID) {
		return nil
	}

	guild := config.Config.Guilds[message.GuildID]
	if !strings.HasPrefix(message.Content, guild.Prefix) {
		return nil
	}

	withoutPrefix := strings.TrimPrefix(message.Content, guild.Prefix)

	commandParts := strings.Split(withoutPrefix, " ")
	if len(commandParts) < 1 {
		return nil
	}

	command := h.Commands[commandParts[0]]
	if command == nil {
		return nil
	}

	parts, err := parseCommandParts(withoutPrefix)
	if err != nil {
		core.Session.ChannelMessageSend(message.ChannelID, err.Error())
		return err
	}

	ctx, err := ConstructContext(parts[1:], command, message)
	if err != nil {
		core.Session.ChannelMessageSend(message.ChannelID, err.Error())
		return err
	}

	for _, check := range command.Checks {
		err := check(ctx)
		if err != nil {
			core.Session.ChannelMessageSend(message.ChannelID, err.Error())
			return err
		}
	}

	err = command.Callback(ctx)
	if err != nil {
		core.Session.ChannelMessageSend(message.ChannelID, err.Error())
		return err
	}

	return nil
}
