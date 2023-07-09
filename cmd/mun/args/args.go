package args

import (
	"strings"

	"github.com/D1360-64RC14/multirun/config"
)

func ExtractConfig(args []string) (*config.Config, error) {
	settings, err := ExtractSettingsField(args)
	if err != nil {
		return nil, err
	}

	commands := ExtractCommandsField(args)

	return &config.Config{
		Commands: commands,
		Settings: settings,
	}, nil
}

func ExtractSettingsField(args []string) (config.Settings, error) {
	color, err := colorArgument.ExtractFrom(args)
	if err != nil {
		return config.Settings{}, err
	}

	return config.Settings{
		Color: color,
	}, nil
}

func ExtractCommandsField(args []string) config.Commands {
	foundCommands := make(config.Commands, len(args))

	for i := 0; i < len(args); i++ {
		arg := strings.TrimLeft(args[i], " ")

		firstSpaceIndex := strings.IndexRune(arg, ' ')
		firstColonIndex := strings.IndexRune(arg, ':')

		haveSpaces := firstSpaceIndex != -1
		haveColons := firstColonIndex != -1

		anySpaceAfterColon := firstSpaceIndex > firstColonIndex

		if haveColons && (!haveSpaces || anySpaceAfterColon) {
			name := arg[:firstColonIndex]
			command := arg[firstColonIndex+1:]

			foundCommands[name] = command
		}
	}

	return foundCommands
}
