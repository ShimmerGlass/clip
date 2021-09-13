package clip

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Print(v interface{}) error {
	return Fprint(os.Stdout, v)
}

func Fprint(w io.Writer, v interface{}) error {
	p := &printer{
		buf: &strings.Builder{},
	}
	err := p.write(v, fieldOptions{
		KeyStyle: style{
			Color: color.FgYellow,
		},
	})
	if err != nil {
		return err
	}
	fmt.Fprintln(w, p.buf.String())
	return nil
}
