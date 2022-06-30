package cases_cmd

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/vcokltfre/volcan/src/commands"
	"github.com/vcokltfre/volcan/src/database"
)

func RegisterCommands() error {
	err := commands.Handler.Register(getCase)
	if err != nil {
		return err
	}

	return nil
}

func createCaseEmbed(dbCase *database.Case) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title: "Case #" + fmt.Sprint(dbCase.ID),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "User",
				Value: dbCase.UserName,
			},
			{
				Name:  "Reason",
				Value: dbCase.Reason,
			},
			{
				Name:  "Moderator",
				Value: dbCase.ModName,
			},
		},
	}
}
