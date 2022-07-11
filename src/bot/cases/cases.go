package cases_cmd

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/vcokltfre/volcan/src/commands"
	"github.com/vcokltfre/volcan/src/database"
	"github.com/vcokltfre/volcan/src/utils"
)

var colours = map[string]int{
	"warn": utils.ColourWarn,
}

func RegisterCommands() error {
	err := commands.Handler.Register(getCase)
	if err != nil {
		return err
	}

	return nil
}

func createCaseEmbed(dbCase *database.Case) *discordgo.MessageEmbed {
	caseDetail := []string{
		fmt.Sprintf("**User:** %s (%s)", dbCase.UserName, dbCase.UserID),
		fmt.Sprintf("**Moderator:** %s (%s)", dbCase.ModName, dbCase.ModID),
		fmt.Sprintf("**Created:** <t:%d:F>", dbCase.CreatedAt),
		fmt.Sprintf("**Reason:** %s", dbCase.Reason),
	}

	colour, ok := colours[dbCase.Type]
	if !ok {
		colour = utils.ColourInfo
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Case #%d: %s", dbCase.ID, dbCase.Type),
		Description: strings.Join(caseDetail, "\n"),
		Color:       colour,
	}
}
