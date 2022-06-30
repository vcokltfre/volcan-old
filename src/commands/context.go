package commands

import (
	"fmt"
	"strings"
)

type Context struct {
	Args      map[string]string
	Flags     map[string]string
	BoolFlags map[string]bool
}

func ConstructContext(parts []string, command *Command) (*Context, error) {
	ctx := &Context{
		Args:      map[string]string{},
		Flags:     map[string]string{},
		BoolFlags: map[string]bool{},
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
				idx++
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
