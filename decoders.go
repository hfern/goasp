package goasp

type decoder_cb func(*parser, pobject) (pobject, error)

var decoders map[otype]decoder_cb = map[otype]decoder_cb{
	tSTR_NULLTERM: parse_str_nullterm,
}

func parse_str_nullterm(p *parser, obj pobject) (pobject, error) {
	str, err := p.b.ReadBytes(0x0)
	str = str[0:len(str)] // get rid of null terminator
	obj.value = string(str)
	return obj, err
}
