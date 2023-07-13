package printer

import (
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// GrayPrinter implements Printer
var _ Printer = (*GrayPrinter)(nil)

type GrayPrinter struct {
	maxNameSize  int
	defaultColor DescribedColor
}

func NewGrayPrinter(possibleNames []string) *GrayPrinter {
	printer := &GrayPrinter{
		maxNameSize:  len(possibleNames[0]),
		defaultColor: DescribedColor{color.FgWhite, color.New(color.FgWhite)},
	}

	for _, possibleName := range possibleNames {
		if len(possibleName) > printer.maxNameSize {
			printer.maxNameSize = len(possibleName)
		}
	}

	return printer
}

func (p *GrayPrinter) Println(name string, content string) (int, error) {
	return p.Fprintln(os.Stdout, name, content)
}

func (p *GrayPrinter) Fprintln(w io.Writer, name string, content string) (int, error) {
	return fmt.Fprintf(
		w,
		"%-[1]*[2]s | %[3]s\n",
		p.maxNameSize,
		name,
		content,
	)
}

func (p *GrayPrinter) GetColorFor(name string) DescribedColor {
	return p.defaultColor
}
