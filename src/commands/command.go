package commands

import (
	"fmt"
	"strings"

	"github.com/vcokltfre/volcan/src/core"
	"github.com/vcokltfre/volcan/src/utils"
)

type CallbackFunction func(*Context) error

type CheckFunction func(*Context) error

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Args        []Arg
	Flags       []Flag
	BoolFlags   []BoolFlag
	Callback    CallbackFunction
	Checks      []CheckFunction
}

type Arg struct {
	Name        string
	Type        string
	Description string
	Validator   Validator
	Required    bool
}

type Flag struct {
	Name        string
	Type        string
	Description string
	Aliases     []string
	Validator   Validator
}

type BoolFlag struct {
	Name        string
	Description string
	Aliases     []string
}

func (c *Command) Validate() error {
	argNames := []string{}
	flagNames := []string{}

	for _, arg := range c.Args {
		if utils.Contains(argNames, arg.Name) {
			return fmt.Errorf("Argument %s is already defined for command %s.", arg.Name, c.Name)
		}

		argNames = append(argNames, arg.Name)
	}

	// TODO: ensure no optional args follow required args

	for _, flag := range c.Flags {
		if utils.Contains(flagNames, flag.Name) {
			return fmt.Errorf("Flag %s is already defined for command %s.", flag.Name, c.Name)
		}

		flagNames = append(flagNames, flag.Name)

		for _, alias := range flag.Aliases {
			if utils.Contains(flagNames, flag.Name) {
				return fmt.Errorf("Flag alias %s is already defined for command %s", alias, c.Name)
			}

			flagNames = append(flagNames, alias)
		}
	}

	for _, flag := range c.BoolFlags {
		if utils.Contains(flagNames, flag.Name) {
			return fmt.Errorf("Flag %s is already defined for command %s.", flag.Name, c.Name)
		}

		flagNames = append(flagNames, flag.Name)

		for _, alias := range flag.Aliases {
			if utils.Contains(flagNames, flag.Name) {
				return fmt.Errorf("Flag alias %s is already defined for command %s", alias, c.Name)
			}

			flagNames = append(flagNames, alias)
		}
	}

	return nil
}

// Given a flag name this will return it's unaliased name
// and whether or not it is a boolean flag.
func (c *Command) getCanonicalFlagName(name string) (string, bool) {
	for _, flag := range c.Flags {
		if flag.Name == name || utils.Contains(flag.Aliases, name) {
			return flag.Name, false
		}
	}

	for _, flag := range c.BoolFlags {
		if flag.Name == name || utils.Contains(flag.Aliases, name) {
			return flag.Name, true
		}
	}

	return "", false
}

func (c *Command) sendHelp(channelID string) error {
	usage := c.Name
	argHelp := []string{}

	for _, arg := range c.Args {
		if arg.Required {
			usage += fmt.Sprintf(" <%s:%s>", arg.Name, arg.Type)
		} else {
			usage += fmt.Sprintf(" [%s:%s]", arg.Name, arg.Type)
		}

		argHelp = append(argHelp, fmt.Sprintf("(`%s`) `%s`: %s", arg.Type, arg.Name, arg.Description))
	}

	argHelp = append(argHelp, "")

	for _, flag := range c.Flags {
		usage += fmt.Sprintf(" [--%s=%s]", flag.Name, flag.Type)
		argHelp = append(argHelp, fmt.Sprintf("(`%s`) `--%s`: %s", flag.Type, flag.Name, flag.Description))
	}

	for _, flag := range c.BoolFlags {
		usage += fmt.Sprintf(" [--%s]", flag.Name)
		argHelp = append(argHelp, fmt.Sprintf("(`bool`) `--%s`: %s", flag.Name, flag.Description))
	}

	commandHelp := fmt.Sprintf("**%s** - *%s*\n\nUsage: `%s`\n\n", c.Name, c.Description, usage)

	commandHelp += strings.Join(argHelp, "\n")

	_, err := core.Session.ChannelMessageSend(channelID, commandHelp)
	return err
}
