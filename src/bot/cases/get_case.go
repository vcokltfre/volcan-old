package cases_cmd

import (
	"encoding/json"

	"github.com/vcokltfre/volcan/src/commands"
	"github.com/vcokltfre/volcan/src/impl"
)

var getCase = &commands.Command{
	Name:        "case",
	Description: "Get a case by ID",
	Args: []commands.Arg{
		{
			Name:      "id",
			Validator: commands.Validator(commands.ValidateInt(0, 9999999)),
			Required:  true,
		},
	},
	BoolFlags: []commands.BoolFlag{
		{
			Name:    "json",
			Aliases: []string{"j"},
		},
	},
	Callback: func(ctx *commands.Context) error {
		val, _ := ctx.Int("id")
		dbCase, err := impl.Interface.Cases.GetCase(val)
		if err != nil {
			return err
		}

		if ctx.BoolFlags["json"] {
			data, err := json.MarshalIndent(dbCase, "", "  ")
			if err != nil {
				return err
			}

			_, err = ctx.Send("```json\n" + string(data) + "\n```")
			return err
		}

		_, err = ctx.SendEmbed(createCaseEmbed(dbCase))

		return err
	},
}
