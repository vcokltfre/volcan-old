package commands

import "github.com/bwmarrin/discordgo"

type CommandHandler struct {
	Commands []Command
}

func New() *CommandHandler {
	return &CommandHandler{
		Commands: []Command{},
	}
}

func (h *CommandHandler) ProcessMessage(message *discordgo.MessageCreate) error {
	return nil
}
