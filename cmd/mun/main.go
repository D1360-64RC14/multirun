package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/D1360-64RC14/multirun/cmd/mun/args"
	"github.com/D1360-64RC14/multirun/config"
	prn "github.com/D1360-64RC14/multirun/pkg/printer"
	"github.com/D1360-64RC14/multirun/pkg/runner"
	"github.com/fatih/color"
)

func init() {
	color.NoColor = false
}

func main() {
	config, err := loadConfiguration()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	commandNames := make([]string, 0, len(config.Commands))

	for name := range config.Commands {
		commandNames = append(commandNames, name)
	}

	var printer prn.Printer

	switch config.Settings.Color {
	case "both":
		printer = prn.NewColoredPrinter(commandNames)
	default:
		printer = prn.NewGrayPrinter(commandNames)
	}

	runner := runner.NewSimpleRunner(config.Commands, printer)

	err = runner.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	kill := make(chan os.Signal, 1)
	signal.Notify(kill, syscall.SIGINT)
	<-kill

	err = runner.Cancel()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	time.Sleep(50 * time.Millisecond)
}

func loadConfiguration() (*config.Config, error) {
	configByArgs, err := loadConfigFromArgs()
	if err != nil {
		return nil, err
	}
	configByFile, err := loadConfigFromFile()
	if err != nil {
		return nil, err
	}

	return &config.Config{
		Commands: appendedMap(configByArgs.Commands, configByFile.Commands),
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

func loadConfigFromArgs() (*config.Config, error) {
	return args.ExtractConfig(os.Args)
}
func loadConfigFromFile() (*config.Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	fullpath := filepath.Join(cwd, "multirun.yaml")

	_, err = os.Stat(fullpath)
	if err != nil {
		return &config.Config{}, nil
	}

	return config.LoadConfig(fullpath)
}

func appendedMap(base config.Commands, plus config.Commands) config.Commands {
	appended := make(config.Commands, len(base)+len(plus))

	for name, command := range base {
		appended[name] = command
	}

	for name, command := range plus {
		if _, ok := appended[name]; !ok {
			appended[name] = command
		}
	}

	return appended
}
