package internal

import (
	"errors"
	"strings"
)

var (
	ErrArgumentNotFound = errors.New("argument not found")
	ErrArgumentInvalid  = errors.New("invalid argument")
)

type Argument struct {
	Name       string
	Required   bool
	Default    string
	Validation func(value string) bool
}

func (a *Argument) ExtractFrom(args []string) (string, error) {
	value, err := a.extractValue(args)
	if err != nil && a.Required {
		return "", err
	}

	valid := a.Validation(value)

	if !valid && a.Required {
		return "", ErrArgumentInvalid
	}

	if err != nil || !valid {
		return a.Default, nil
	}

	return value, nil
}

func (a *Argument) extractValue(args []string) (string, error) {
	combinedName := a.Name + "="

	for i := 0; i < len(args); i++ {
		arg := args[i]

		if !strings.HasPrefix(arg, a.Name) {
			continue
		}

		if strings.HasPrefix(arg, combinedName) {
			sepIndex := strings.IndexRune(arg, '=')
			return arg[sepIndex+1:], nil
		}

		if len(args) >= i+1 {
			return args[i+1], nil
		}
	}

	return "", ErrArgumentNotFound
}
