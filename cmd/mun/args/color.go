package args

import "github.com/D1360-64RC14/multirun/internal"

var colorArgument = internal.Argument{
	Name:     "--color",
	Required: false,
	Default:  "both",
	Validation: func(value string) bool {
		switch value {
		case "none", "mun", "command", "both":
			return true
		default:
			return false
		}
	},
}
