package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/vcokltfre/volcan/src/core"
)

type Context struct {
	Args      map[string]string
	Flags     map[string]string
	BoolFlags map[string]bool
	ChannelID string
	GuildId   string
	Author    *discordgo.Member
}

func ConstructContext(parts []string, command *Command, message *discordgo.MessageCreate) (*Context, error) {
	ctx := &Context{
		Args:      map[string]string{},
		Flags:     map[string]string{},
		BoolFlags: map[string]bool{},
		ChannelID: message.ChannelID,
		GuildId:   message.GuildID,
		Author:    message.Member,
	}

	requiredArgs := []string{}
	optionalArgs := []string{}

	argValidators := map[string]Validator{}
	flagValidators := map[string]Validator{}

	for _, arg := range command.Args {
		if arg.Required {
			requiredArgs = append(requiredArgs, arg.Name)
		} else {
			optionalArgs = append(optionalArgs, arg.Name)
		}

		if arg.Validator == nil {
			continue
		}

		argValidators[arg.Name] = arg.Validator
	}

	for _, flag := range command.Flags {
		if flag.Validator == nil {
			continue
		}

		flagValidators[flag.Name] = flag.Validator
	}

	idx := 0

	for idx < len(parts) {
		part := parts[idx]

		if strings.HasPrefix(part, "--") {
			flagName := strings.TrimPrefix(part, "--")
			flag, isBool := command.getCanonicalFlagName(flagName)

			if flag == "" {
				return nil, fmt.Errorf("Flag %s is not defined for command %s.", flagName, command.Name)
			}

			if isBool {
				ctx.BoolFlags[flag] = true
			} else {
				if idx >= len(parts) {
					return nil, fmt.Errorf("Flag %s requires an argument.", flag)
				}

				validator, ok := flagValidators[flag]
				if ok {
					err := validator(parts[idx])
					if err != nil {
						return nil, err
					}
				}

				ctx.Flags[flag] = parts[idx]
			}

			idx++

			continue
		}

		if strings.HasPrefix(part, "-") {
			flagNames := strings.TrimPrefix(part, "-")

			for _, flagName := range flagNames {
				flag, isBool := command.getCanonicalFlagName(string(flagName))

				if flag == "" {
					return nil, fmt.Errorf("Flag %d is not defined for command %s.", flagName, command.Name)
				}

				if isBool {
					ctx.BoolFlags[flag] = true
				} else {
					if idx >= len(parts) {
						return nil, fmt.Errorf("Flag %s requires an argument.", flag)
					}

					validator, ok := flagValidators[flag]
					if ok {
						err := validator(parts[idx])
						if err != nil {
							return nil, err
						}
					}

					ctx.Flags[flag] = parts[idx]
				}
			}

			idx++

			continue
		}

		if len(requiredArgs) > 0 {
			argName := requiredArgs[0]
			requiredArgs = requiredArgs[1:]

			validator, ok := argValidators[argName]
			if ok {
				err := validator(part)
				if err != nil {
					return nil, err
				}
			}

			ctx.Args[argName] = part
			idx += 1
			continue
		}

		if len(optionalArgs) > 0 {
			argName := optionalArgs[0]
			optionalArgs = optionalArgs[1:]

			validator, ok := argValidators[argName]
			if ok {
				err := validator(part)
				if err != nil {
					return nil, err
				}
			}

			ctx.Args[argName] = part
			idx += 1
			continue
		}

		return nil, fmt.Errorf("Too many arguments for command %s.", command.Name)
	}

	if len(requiredArgs) > 0 {
		return nil, fmt.Errorf("Missing required arguments: %s", strings.Join(requiredArgs, ", "))
	}

	return ctx, nil
}

func (c *Context) Int(name string) (int, error) {
	str, ok := c.Args[name]
	if !ok {
		return 0, fmt.Errorf("Argument %s is not defined.", name)
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return i, nil
}

func (c *Context) Send(data string) (*discordgo.Message, error) {
	return core.Session.ChannelMessageSend(c.ChannelID, data)
}

func (c *Context) SendEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return core.Session.ChannelMessageSendEmbed(c.ChannelID, embed)
}
