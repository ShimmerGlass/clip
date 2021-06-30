package clip

func (p *printer) writeScalar(v interface{}, opt fieldOptions) error {
	opt.ValueStyle.color().Fprint(p.buf, v)
	return nil
}
