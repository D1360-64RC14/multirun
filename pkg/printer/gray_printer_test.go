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

func TestGrayPrinter_Fprintln_PerfectCases(t *testing.T) {
	testCases := []struct {
		inputName      string
		inputContent   string
		expectedOutput string
	}{
		{
			"Name 1",
			"Hello World",
			"Name 1 | Hello World\n",
		},
		{
			"Name 3",
			"This is fine",
			"Name 3 | This is fine\n",
		},
		{
			"Name 1",
			"Lorem ipsum",
			"Name 1 | Lorem ipsum\n",
		},
	}

	printer := prn.NewGrayPrinter([]string{
		"Name 1",
		"Name 2",
		"Name 3",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != color.FgWhite {
				t.Errorf("expected attribute '%d', got '%d'", color.FgWhite, colorAttr)
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

func TestGrayPrinter_Fprintln_LineSizes(t *testing.T) {
	testCases := []struct {
		inputName      string
		inputContent   string
		expectedOutput string
	}{
		{
			"typescript-compiler",
			"npx tsc -w",
			"typescript-compiler | npx tsc -w\n",
		},
		{
			"lorem-gen",
			"lipsum",
			"lorem-gen           | lipsum\n",
		},
		{
			"the-preprocessor",
			"npx sass -w --no-source-map src/styles:static/styles",
			"the-preprocessor    | npx sass -w --no-source-map src/styles:static/styles\n",
		},
	}

	printer := prn.NewGrayPrinter([]string{
		"lorem-gen",
		"typescript-compiler",
		"the-preprocessor",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != color.FgWhite {
				t.Errorf("expected attribute '%d', got '%d'", color.FgWhite, colorAttr)
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

func TestGrayPrinter_Fprintln_InvalidNames(t *testing.T) {
	testCases := []struct {
		inputName      string
		inputContent   string
		expectedOutput string
	}{
		{
			"Name 1",
			"Hello World",
			"Name 1 | Hello World\n",
		},
		{
			"Name 3",
			"This is fine",
			"Name 3 | This is fine\n",
		},
		{
			"An invalid name",
			"Lorem ipsum",
			"An invalid name | Lorem ipsum\n",
		},
		{
			"Name 2",
			"Dolor sit amet",
			"Name 2 | Dolor sit amet\n",
		},
		{
			"other",
			"Consectetur adipiscing elit",
			"other  | Consectetur adipiscing elit\n",
		},
	}

	printer := prn.NewGrayPrinter([]string{
		"Name 1",
		"Name 2",
		"Name 3",
	})

	for i, _case := range testCases {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			colorAttr := printer.GetColorFor(_case.inputName).Attribute

			if colorAttr != color.FgWhite {
				t.Errorf("expected attribute '%d', got '%d'", color.FgWhite, colorAttr)
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
