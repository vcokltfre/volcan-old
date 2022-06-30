package bot

import cases_cmd "github.com/vcokltfre/volcan/src/bot/cases"

func RegisterCommands() error {
	err := cases_cmd.RegisterCommands()
	if err != nil {
		return err
	}

	return nil
}
