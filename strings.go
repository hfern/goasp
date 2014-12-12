package goasp

func parse_str_nullterm(p *parser, obj pobject) (pobject, error) {
	str, err := p.b.ReadBytes(0x0)
	str = str[0:len(str)] // get rid of null terminator
	obj.value = string(str)
	return obj, err
}
