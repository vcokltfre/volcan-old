package core

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/volcan/src/commands"
)

func setupEventListeners() error {
	handler = commands.New()

	Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		err := handler.ProcessMessage(m)
		if err != nil {
			logrus.Errorf("Error processing message: %v", err)
		}
	})

	return nil
}
