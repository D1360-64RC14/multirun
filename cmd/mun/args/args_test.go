package args_test

import (
	"reflect"
	"testing"

	"github.com/D1360-64RC14/multirun/cmd/mun/args"
	"github.com/D1360-64RC14/multirun/config"
)

func TestExtractCommandsField(t *testing.T) {
	testCases := []struct {
		description    string
		givenArgs      []string
		resultCommands config.Commands
	}{
		{
			"Basic single perfect-environment argument",
			[]string{
				"sass-preprocessor:npx sass -w --no-source-map src/styles:static/styles",
			},
			config.Commands{
				"sass-preprocessor": "npx sass -w --no-source-map src/styles:static/styles",
			},
		},
		{
			"Basic multiple perfect-environment arguments",
			[]string{
				"echo-starting:echo Starting...",
				"sass-preprocessor:npx sass -w --no-source-map src/styles:static/styles",
				"ts-compiller:npx tsc -w",
				"list-directory:ls",
			},
			config.Commands{
				"echo-starting":     "echo Starting...",
				"sass-preprocessor": "npx sass -w --no-source-map src/styles:static/styles",
				"ts-compiller":      "npx tsc -w",
				"list-directory":    "ls",
			},
		},
		{
			"Strange combinations",
			[]string{
				"text-1:hello:world",
				"text-2:\"hello world\"",
			},
			config.Commands{
				"text-1": "hello:world",
				"text-2": "\"hello world\"",
			},
		},
		{
			"Some that can't be catch",
			[]string{
				"this text:should not be catch",
				"this :also",
				"should-catch: only this",
			},
			config.Commands{
				"should-catch": " only this",
			},
		},
		{
			"Strange spaces",
			[]string{
				"space-in-command: testing space before command",
				" space-in-name:testing space before name",
				"      lot-space-in-name:testing [a lot of] space before name",
				"    man:            wtf?",
			},
			config.Commands{
				"space-in-command":  " testing space before command",
				"space-in-name":     "testing space before name",
				"lot-space-in-name": "testing [a lot of] space before name",
				"man":               "            wtf?",
			},
		},
	}

	for _, _case := range testCases {
		t.Run(_case.description, func(t *testing.T) {
			commands := args.ExtractCommandsField(_case.givenArgs)

			if !reflect.DeepEqual(commands, _case.resultCommands) {
				t.Errorf("Resulting map should be '%#v', but got '%#v'", _case.resultCommands, commands)
			}
		})
	}
}
