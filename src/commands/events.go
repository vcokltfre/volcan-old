package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/volcan/src/core"
)

func setupEventListeners() error {
	core.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		err := Handler.ProcessMessage(m)
		if err != nil {
			logrus.Errorf("Error processing message: %v", err)
		}
	})

	return nil
}
