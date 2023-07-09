package internal_test

import (
	"fmt"
	"testing"

	"github.com/D1360-64RC14/multirun/internal"
)

var argument_OptionalColor = internal.Argument{
	Name:     "--color",
	Required: false,
	Default:  "green",
	Validation: func(value string) bool {
		switch value {
		case "red", "green", "blue":
			return true
		default:
			return false
		}
	},
}
var argument_RequiredColor = internal.Argument{
	Name:     "--color",
	Required: true,
	Default:  "",
	Validation: func(value string) bool {
		switch value {
		case "red", "green", "blue":
			return true
		default:
			return false
		}
	},
}

func TestArgument_ExtractFrom_Required(t *testing.T) {
	testCases := []struct {
		givenArgs      []string
		expectedString string
		expectedError  error
	}{
		// --- Valid Color ---
		{[]string{"some-text", "--color", "red"}, "red", nil},
		{[]string{"some-text", "--color", "green"}, "green", nil},
		{[]string{"some-text", "--color", "blue"}, "blue", nil},
		{[]string{"some-text", "--color=red"}, "red", nil},
		{[]string{"some-text", "--color=green"}, "green", nil},
		{[]string{"some-text", "--color=blue"}, "blue", nil},
		{[]string{"--color", "blue", "some-text"}, "blue", nil},
		{[]string{"--color=blue", "some-text"}, "blue", nil},
		{[]string{"some-text", "--color=blue", "red"}, "blue", nil},
		// --- Invalid Color ---
		{[]string{"some-text", "--color="}, "", internal.ErrArgumentInvalid},
		{[]string{"some-text", "--color=", "red"}, "", internal.ErrArgumentInvalid},
		{[]string{"some-text", "--color=yellow", "blue"}, "", internal.ErrArgumentInvalid},
		{[]string{"some-text", "--color", "cyan", "blue"}, "", internal.ErrArgumentInvalid},
		// --- Missing Argument ---
		{[]string{"some-text", "color"}, "", internal.ErrArgumentNotFound},
		{[]string{"read", "some-text"}, "", internal.ErrArgumentNotFound},
		{[]string{}, "", internal.ErrArgumentNotFound},
		{[]string{"--read-only=true", "some-text"}, "", internal.ErrArgumentNotFound},
	}

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			color, err := argument_RequiredColor.ExtractFrom(_case.givenArgs)

			if err != _case.expectedError {
				t.Errorf("Error should be '%!', got '%!'", _case.expectedError, err)
			}
			if color != _case.expectedString {
				t.Errorf("Color should be '%s', got '%s'", _case.expectedString, color)
			}
		})
	}
}

func TestArgument_ExtractFrom_Optional(t *testing.T) {
	testCases := []struct {
		givenArgs      []string
		expectedString string
		expectedError  error
	}{
		// --- Valid Color ---
		{[]string{"some-text", "--color", "red"}, "red", nil},
		{[]string{"some-text", "--color", "green"}, "green", nil},
		{[]string{"some-text", "--color", "blue"}, "blue", nil},
		{[]string{"some-text", "--color=red"}, "red", nil},
		{[]string{"some-text", "--color=green"}, "green", nil},
		{[]string{"some-text", "--color=blue"}, "blue", nil},
		{[]string{"--color", "blue", "some-text"}, "blue", nil},
		{[]string{"--color=blue", "some-text"}, "blue", nil},
		{[]string{"some-text", "--color=blue", "red"}, "blue", nil},
		// --- Invalid Color ---
		{[]string{"some-text", "--color="}, "green", nil},
		{[]string{"some-text", "--color=", "red"}, "green", nil},
		{[]string{"some-text", "--color=yellow", "blue"}, "green", nil},
		{[]string{"some-text", "--color", "cyan", "blue"}, "green", nil},
		// --- Missing Argument ---
		{[]string{"some-text", "color"}, "green", nil},
		{[]string{"read", "some-text"}, "green", nil},
		{[]string{}, "green", nil},
		{[]string{"--read-only=true", "some-text"}, "green", nil},
	}

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			color, err := argument_OptionalColor.ExtractFrom(_case.givenArgs)

			if err != _case.expectedError {
				t.Errorf("Error should be '%!', got '%!'", _case.expectedError, err)
			}
			if color != _case.expectedString {
				t.Errorf("Color should be '%s', got '%s'", _case.expectedString, color)
			}
		})
	}
}
