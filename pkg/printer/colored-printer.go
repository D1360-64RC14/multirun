package printer

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// ColoredPrinter implements Printer
var _ Printer = (*ColoredPrinter)(nil)

type DescribedColor struct {
	attribute color.Attribute
	color     *color.Color
}

type ColoredPrinter struct {
	maxNameSize    int
	nameColorPairs map[string]DescribedColor
	defaultColor   DescribedColor
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

func NewColoredPrinter(possibleNames []string) *ColoredPrinter {
	printer := &ColoredPrinter{
		maxNameSize:    len(possibleNames[0]),
		nameColorPairs: make(map[string]DescribedColor, len(possibleNames)),
		defaultColor:   DescribedColor{color.FgWhite, color.New(color.FgWhite)},
	}

	for i, possibleName := range possibleNames {
		// check maximum name size
		if len(possibleName) > printer.maxNameSize {
			printer.maxNameSize = len(possibleName)
		}

		// assign name with color
		colorIndex := i % len(Colors)
		printer.nameColorPairs[possibleName] = DescribedColor{
			attribute: Colors[colorIndex],
			color:     color.New(Colors[colorIndex]),
		}
	}

	return printer
}

func (p *ColoredPrinter) Println(name string, content string) (int, error) {
	return p.Fprintln(os.Stdout, name, content)
}

func (p *ColoredPrinter) Fprintln(w io.Writer, name, content string) (int, error) {
	colorToUse := p.GetColorFor(name).color

	bytesWritten, err := fmt.Fprintf(
		w,
		"%-[1]*[2]s | %[3]s\n",
		p.maxNameSize,
		colorToUse.Sprint(name),
		content,
	)
	if err != nil {
		return bytesWritten, err
	}

	return bytesWritten, nil
}

func (p *ColoredPrinter) GetColorFor(name string) DescribedColor {
	pair, ok := p.nameColorPairs[name]

	if !ok {
		return p.defaultColor
	}

	return pair
}
