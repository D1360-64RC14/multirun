package printer

import "io"

type Printer interface {
	Println(name, content string) (int, error)
	Fprintln(w io.Writer, name, content string) (int, error)
	GetColorFor(name string) DescribedColor
}
