package clip

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	inline = "inline"

	formatPercent = "percent"

	colorBlack   = "black"
	colorRed     = "red"
	colorGreen   = "green"
	colorYellow  = "yellow"
	colorBlue    = "blue"
	colorMagenta = "magenta"
	colorCyan    = "cyan"
	colorWhite   = "white"

	keyColorBlack   = "key_" + colorBlack
	keyColorRed     = "key_" + colorRed
	keyColorGreen   = "key_" + colorGreen
	keyColorYellow  = "key_" + colorYellow
	keyColorBlue    = "key_" + colorBlue
	keyColorMagenta = "key_" + colorMagenta
	keyColorCyan    = "key_" + colorCyan
	keyColorWhite   = "key_" + colorWhite
)

type style struct {
	Color      color.Attribute
	Bold       bool
	Italic     bool
	Underline  bool
	CrossedOut bool
}

func (s style) color() *color.Color {
	attrs := []color.Attribute{s.Color}
	if s.Bold {
		attrs = append(attrs, color.Bold)
	}
	if s.Italic {
		attrs = append(attrs, color.Italic)
	}
	if s.Underline {
		attrs = append(attrs, color.Underline)
	}
	if s.CrossedOut {
		attrs = append(attrs, color.CrossedOut)
	}

	return color.New(attrs...)
}

type fieldOptions struct {
	Name       string
	Inline     bool
	KeyStyle   style
	ValueStyle style
	Format     string
}

func parseTag(base fieldOptions, s string) (fieldOptions, error) {
	parts := strings.Split(s, ",")

	for i, part := range parts {
		if i == 0 {
			base.Name = part
			continue
		}

		switch part {
		case inline:
			base.Inline = true

		case keyColorBlack, keyColorRed, keyColorGreen, keyColorYellow,
			keyColorBlue, keyColorMagenta, keyColorCyan, keyColorWhite:
			base.KeyStyle = styleColor(base.KeyStyle, strings.TrimPrefix(part, "key_"))

		case colorBlack, colorRed, colorGreen, colorYellow,
			colorBlue, colorMagenta, colorCyan, colorWhite:
			base.ValueStyle = styleColor(base.ValueStyle, part)

		case formatPercent:
			base.Format = part

		default:
			return base, fmt.Errorf("unknow option %q", part)

		}
	}

	return base, nil
}

func styleColor(s style, v string) style {
	switch v {
	case colorBlack:
		s.Color = color.FgBlack
	case colorRed:
		s.Color = color.FgRed
	case colorGreen:
		s.Color = color.FgGreen
	case colorYellow:
		s.Color = color.FgYellow
	case colorBlue:
		s.Color = color.FgBlue
	case colorMagenta:
		s.Color = color.FgMagenta
	case colorCyan:
		s.Color = color.FgCyan
	case colorWhite:
		s.Color = color.FgWhite
	}

	return s
}
