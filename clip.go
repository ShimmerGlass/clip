package clip

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func Print(v interface{}) error {
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
	fmt.Println(p.buf.String())
	return nil
}
