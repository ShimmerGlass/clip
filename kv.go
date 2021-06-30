package clip

import (
	"fmt"
	"reflect"
)

type kvField struct {
	fieldOptions
	Value interface{}
}

func (p *printer) writeFields(fields []kvField, opt fieldOptions) error {
	for i, field := range fields {
		p.writeIndent()
		opt.KeyStyle.color().Fprint(p.buf, field.Name)
		p.buf.WriteString(": ")

		if !field.fieldOptions.Inline {
			p.buf.WriteByte('\n')
		}

		p.indent++
		err := p.write(field.Value, field.fieldOptions)
		if err != nil {
			return err
		}
		p.indent--

		if i != len(fields)-1 {
			if !opt.Inline {
				p.buf.WriteByte('\n')
			} else {
				p.buf.WriteByte('\t')
			}
		}
	}

	return nil
}

func (p *printer) structFields(val reflect.Value, base fieldOptions) ([]kvField, error) {
	typ := val.Type()

	res := []kvField{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		opts, err := parseTag(base, field.Tag.Get("clip"))
		if err != nil {
			return nil, fmt.Errorf("%T.%s: %w", typ.Name(), field.Name, err)
		}
		if opts.Name == "" {
			opts.Name = field.Name
		}

		res = append(res, kvField{
			Value:        val.Field(i).Interface(),
			fieldOptions: opts,
		})
	}

	return res, nil
}

func (p *printer) mapFields(val reflect.Value) ([]kvField, error) {
	iter := val.MapRange()

	res := []kvField{}
	for iter.Next() {
		res = append(res, kvField{
			Value: iter.Value().Interface(),
			fieldOptions: fieldOptions{
				Name: iter.Key().String(),
			},
		})
	}

	return res, nil
}
