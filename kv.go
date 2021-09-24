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
		p.buf.WriteString(":")

		if !field.Inline {
			p.buf.WriteString("\n")
		} else {
			p.buf.WriteString(" ")
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

		value := val.Field(i)
		if !value.CanInterface() {
			continue
		}

		if p.shouldInline(value) {
			opts.Inline = true
		}

		if value.Type().Kind() == reflect.Ptr {
			value = reflect.Indirect(value)
		}

		if field.Anonymous && value.Kind() == reflect.Struct {
			fields, err := p.structFields(value, base)
			if err != nil {
				return nil, err
			}

			res = append(res, fields...)
		} else {
			res = append(res, kvField{
				Value:        value.Interface(),
				fieldOptions: opts,
			})
		}
	}

	return res, nil
}

func (p *printer) mapFields(val reflect.Value) ([]kvField, error) {
	iter := val.MapRange()

	res := []kvField{}
	for iter.Next() {
		opts := fieldOptions{
			Name:   iter.Key().String(),
			Inline: p.shouldInline(iter.Value()),
		}

		res = append(res, kvField{
			Value:        iter.Value().Interface(),
			fieldOptions: opts,
		})
	}

	return res, nil
}

func (p *printer) shouldInline(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.Map:
		return val.Len() == 0
	case reflect.Float32, reflect.Float64, reflect.String, reflect.Int,
		reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Bool:
		return true
	}

	return false
}
