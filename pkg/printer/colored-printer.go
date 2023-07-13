package printer

import (
	"fmt"

	"github.com/fatih/color"
)

// ColoredPrinter implements Printer
var _ Printer = (*ColoredPrinter)(nil)

type ColoredPrinter struct {
	maxNameSize    int
	nameColorPairs map[string]*color.Color
}

var Colors = [...]color.Attribute{
	color.FgWhite,
	color.FgGreen,
	color.FgCyan,
	color.FgYellow,
	color.FgMagenta,
	color.FgRed,
	color.FgBlue,
}
var DefaultColor = color.FgWhite

func NewColoredPrinter(possibleNames []string) *ColoredPrinter {
	printer := &ColoredPrinter{
		maxNameSize:    len(possibleNames[0]),
		nameColorPairs: make(map[string]*color.Color, len(possibleNames)),
	}

	for i, possibleName := range possibleNames {
		// check maximum name size
		if len(possibleName) > printer.maxNameSize {
			printer.maxNameSize = len(possibleName)
		}

		// assign name with color
		colorIndex := i % len(Colors)
		printer.nameColorPairs[possibleName] = color.New(Colors[colorIndex])
	}

	return printer
}

func (p *ColoredPrinter) Println(name string, content string) {
	colorToUse := p.GetColorFor(name)

	fmt.Printf(
		"%-[1]*[2]s | %[3]s",
		p.maxNameSize,
		colorToUse.Sprint(name),
		content,
	)
}

func (p *ColoredPrinter) GetColorFor(name string) *color.Color {
	colorToUse, ok := p.nameColorPairs[name]

	if !ok {
		colorToUse = color.New(DefaultColor)
	}

	return colorToUse
}
