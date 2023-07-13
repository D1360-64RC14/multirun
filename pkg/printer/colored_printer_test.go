package printer_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	prn "github.com/D1360-64RC14/multirun/pkg/printer"
	"github.com/fatih/color"
)

func init() {
	color.NoColor = false
}

func TestColoredPrinter_Fprintln_PerfectCases(t *testing.T) {
	testCases := []struct {
		inputName         string
		inputContent      string
		expectedOutput    string
		expectedColorAttr color.Attribute
	}{
		{
			"Name 1",
			"Hello World",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("Name 1 |"), "Hello World"),
			color.FgWhite,
		},
		{
			"Name 3",
			"This is fine",
			fmt.Sprintf("%s %s\n", color.New(color.FgCyan).Sprint("Name 3 |"), "This is fine"),
			color.FgCyan,
		},
		{
			"Name 1",
			"Lorem ipsum",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("Name 1 |"), "Lorem ipsum"),
			color.FgWhite,
		},
		{
			"Name 2",
			"Dolor sit amet",
			fmt.Sprintf("%s %s\n", color.New(color.FgGreen).Sprint("Name 2 |"), "Dolor sit amet"),
			color.FgGreen,
		},
		{
			"Name 3",
			"Consectetur adipiscing elit",
			fmt.Sprintf("%s %s\n", color.New(color.FgCyan).Sprint("Name 3 |"), "Consectetur adipiscing elit"),
			color.FgCyan,
		},
	}

	printer := prn.NewColoredPrinter([]string{
		"Name 1",
		"Name 2",
		"Name 3",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != _case.expectedColorAttr {
				t.Errorf("expected attribute '%d', got '%d'", _case.expectedColorAttr, colorAttr)
			}

			buffer := bytes.NewBufferString("")

			_, err := printer.Fprintln(buffer, _case.inputName, _case.inputContent)
			if err != nil {
				t.Errorf("unexpected error while printing to buffer: '%s'", err)
			}

			actualContent, err := io.ReadAll(buffer)
			if err != nil {
				t.Errorf("unexpected error while reading from buffer: '%s'", err)
			}

			if string(actualContent) != _case.expectedOutput {
				t.Errorf("expected content '%s', got '%s'", _case.expectedOutput, actualContent)
			}
		})
	}
}

func TestColoredPrinter_Fprintln_LineSizes(t *testing.T) {
	testCases := []struct {
		inputName         string
		inputContent      string
		expectedOutput    string
		expectedColorAttr color.Attribute
	}{
		{
			"typescript-compiler",
			"npx tsc -w",
			fmt.Sprintf("%s %s\n", color.New(color.FgGreen).Sprint("typescript-compiler |"), "npx tsc -w"),
			color.FgGreen,
		},
		{
			"lorem-gen",
			"lipsum",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("lorem-gen           |"), "lipsum"),
			color.FgWhite,
		},
		{
			"the-preprocessor",
			"npx sass -w --no-source-map src/styles:static/styles",
			fmt.Sprintf("%s %s\n", color.New(color.FgCyan).Sprint("the-preprocessor    |"), "npx sass -w --no-source-map src/styles:static/styles"),
			color.FgCyan,
		},
	}

	printer := prn.NewColoredPrinter([]string{
		"lorem-gen",
		"typescript-compiler",
		"the-preprocessor",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != _case.expectedColorAttr {
				t.Errorf("expected attribute '%d', got '%d'", _case.expectedColorAttr, colorAttr)
			}

			buffer := bytes.NewBufferString("")

			_, err := printer.Fprintln(buffer, _case.inputName, _case.inputContent)
			if err != nil {
				t.Errorf("unexpected error while printing to buffer: '%s'", err)
			}

			actualContent, err := io.ReadAll(buffer)
			if err != nil {
				t.Errorf("unexpected error while reading from buffer: '%s'", err)
			}

			if string(actualContent) != _case.expectedOutput {
				t.Errorf("expected content '%s', got '%s'", _case.expectedOutput, actualContent)
			}
		})
	}
}

func TestColoredPrinter_Fprintln_InvalidNames(t *testing.T) {
	testCases := []struct {
		inputName         string
		inputContent      string
		expectedOutput    string
		expectedColorAttr color.Attribute
	}{
		{
			"Name 1",
			"Hello World",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("Name 1 |"), "Hello World"),
			color.FgWhite,
		},
		{
			"Name 3",
			"This is fine",
			fmt.Sprintf("%s %s\n", color.New(color.FgCyan).Sprint("Name 3 |"), "This is fine"),
			color.FgCyan,
		},
		{
			"An invalid name",
			"Lorem ipsum",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("An invalid name |"), "Lorem ipsum"),
			color.FgWhite,
		},
		{
			"Name 2",
			"Dolor sit amet",
			fmt.Sprintf("%s %s\n", color.New(color.FgGreen).Sprint("Name 2 |"), "Dolor sit amet"),
			color.FgGreen,
		},
		{
			"other",
			"Consectetur adipiscing elit",
			fmt.Sprintf("%s %s\n", color.New(color.FgWhite).Sprint("other  |"), "Consectetur adipiscing elit"),
			color.FgWhite,
		},
	}

	printer := prn.NewColoredPrinter([]string{
		"Name 1",
		"Name 2",
		"Name 3",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != _case.expectedColorAttr {
				t.Errorf("expected attribute '%d', got '%d'", _case.expectedColorAttr, colorAttr)
			}

			buffer := bytes.NewBufferString("")

			_, err := printer.Fprintln(buffer, _case.inputName, _case.inputContent)
			if err != nil {
				t.Errorf("unexpected error while printing to buffer: '%s'", err)
			}

			actualContent, err := io.ReadAll(buffer)
			if err != nil {
				t.Errorf("unexpected error while reading from buffer: '%s'", err)
			}

			if string(actualContent) != _case.expectedOutput {
				t.Errorf("expected content '%s', got '%s'", _case.expectedOutput, actualContent)
			}
		})
	}
}
