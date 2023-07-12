package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/D1360-64RC14/multirun/cmd/mun/args"
	"github.com/D1360-64RC14/multirun/config"
	prn "github.com/D1360-64RC14/multirun/pkg/printer"
)

func main() {
	config, err := loadConfiguration()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	commandNames := make([]string, 0, len(config.Commands))

	for name, _ := range config.Commands {
		commandNames = append(commandNames, name)
	}

	var printer prn.Printer

	switch config.Settings.Color {
	case "both":
		printer = prn.NewColoredPrinter(commandNames)
	default:
		printer = prn.NewGrayPrinter(commandNames)
	}

	fmt.Println("TODO: command runner")
	os.Exit(1)
}

func loadConfiguration() (*config.Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	configByArgs, err := args.ExtractConfig(os.Args)
	if err != nil {
		return nil, err
	}
	configByFile, err := config.LoadConfig(filepath.Join(cwd, "multirun.yaml"))
	if err != nil {
		return nil, err
	}

	return &config.Config{
		Commands: appendedMap(&configByArgs.Commands, &configByFile.Commands),
		Settings: config.Settings{
			Color: func() string {
				if !args.ColorArgument.Validation(configByArgs.Settings.Color) {
					return configByFile.Settings.Color
				}
				return configByArgs.Settings.Color
			}(),
		},
	}, nil
}

func appendedMap(base *config.Commands, plus *config.Commands) config.Commands {
	appended := make(config.Commands, len(*base)+len(*plus))

	for name, command := range *plus {
		if _, ok := appended[name]; !ok {
			appended[name] = command
		}
	}

	return appended
}
