package printer

import (
	"fmt"
)

// GrayPrinter implements Printer
var _ Printer = (*GrayPrinter)(nil)

type GrayPrinter struct {
	maxNameSize int
}

func NewGrayPrinter(possibleNames []string) *GrayPrinter {
	printer := &GrayPrinter{
		maxNameSize: len(possibleNames[0]),
	}

	for _, possibleName := range possibleNames {
		if len(possibleName) > printer.maxNameSize {
			printer.maxNameSize = len(possibleName)
		}
	}

	return printer
}

func (p *GrayPrinter) Println(name string, content string) {
	fmt.Printf(
		"%-[1]*[2]s | %[3]s",
		p.maxNameSize,
		name,
		content,
	)
}
