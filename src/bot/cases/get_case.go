package cases_cmd

import (
	"github.com/vcokltfre/volcan/src/commands"
	"github.com/vcokltfre/volcan/src/impl"
	"github.com/vcokltfre/volcan/src/utils"
)

var getCase = &commands.Command{
	Name:        "case",
	Description: "Get a case by ID",
	Args: []commands.Arg{
		{
			Name:        "id",
			Type:        "number",
			Description: "The ID of the case to get.",
			Validator:   commands.ValidateInt(0, 9999999),
			Required:    true,
		},
	},
	BoolFlags: []commands.BoolFlag{
		{
			Name:        "json",
			Description: "Output in JSON format.",
			Aliases:     []string{"j"},
		},
	},
	Callback: func(ctx *commands.Context) error {
		val, _ := ctx.Int("id")
		dbCase, err := impl.Interface.Cases.GetCase(val)
		if err != nil {
			return err
		}

		if ctx.BoolFlags["json"] {
			data, err := utils.PrettifyJSON(dbCase)
			if err != nil {
				return err
			}

			_, err = ctx.Send(data)
			return err
		}

		_, err = ctx.SendEmbed(createCaseEmbed(dbCase))

		return err
	},
}
