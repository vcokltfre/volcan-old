package core

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

var Session *discordgo.Session

func SetupSession() error {
	sess, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		return err
	}

	Session = sess

	// We make use of all intents at various stages of moderation
	// so it's not a sin to use them all here.
	Session.Identify.Intents = discordgo.IntentsAll

	logrus.Info("Connecting to the Discord gateway...")

	err = sess.Open()
	if err != nil {
		return err
	}

	logrus.Info("Successfully connected to the Discord gateway.")

	return nil
}
