package clip

import (
	"reflect"
	"strings"
)

type printer struct {
	buf *strings.Builder

	indent         int
	indentPrewrite bool
}

func (p *printer) write(v interface{}, opt fieldOptions) error {
	val := reflect.ValueOf(v)
	if val.Type().Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	switch val.Type().Kind() {
	case reflect.Struct:
		fields, err := p.structFields(val, opt)
		if err != nil {
			return err
		}
		err = p.writeFields(fields, opt)
		if err != nil {
			return err
		}

	case reflect.Map:
		fields, err := p.mapFields(val)
		if err != nil {
			return err
		}
		err = p.writeFields(fields, opt)
		if err != nil {
			return err
		}

	case reflect.Array, reflect.Slice:
		err := p.writeList(val, opt)
		if err != nil {
			return err
		}

	case reflect.Float32, reflect.Float64, reflect.String, reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Bool:
		err := p.writeScalar(val.Interface(), opt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *printer) writeIndent() {
	if p.indentPrewrite {
		p.indentPrewrite = false
		return
	}
	p.buf.WriteString(strings.Repeat("  ", p.indent))
}
