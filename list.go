package clip

import "reflect"

func (p *printer) writeList(val reflect.Value, opt fieldOptions) error {
	for i := 0; i < val.Len(); i++ {
		if !opt.Inline {
			p.writeIndent()
			p.buf.WriteString("- ")
			p.indentPrewrite = true
		}

		p.indent++
		err := p.write(val.Index(i).Interface(), opt)
		if err != nil {
			return err
		}
		p.indent--

		if !opt.Inline {
			p.buf.WriteByte('\n')
		} else {
			p.buf.WriteByte('\t')
		}
	}

	return nil
}
