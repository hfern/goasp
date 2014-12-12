package goasp

func parse_length_string(p *parser, obj pobject) (pobject, error) {
	length, err := p.b.ReadByte()
	if err != nil {
		return obj, err
	}

	strbuf := make([]byte, length)

	for i := byte(0); i < length; i++ {
		strbuf[i], err = p.b.ReadByte()
		if err != nil {
			return pobject{}, err
		}
	}

	obj.value = string(strbuf)
	return obj, nil
}

func parse_str_nullterm(p *parser, obj pobject) (pobject, error) {
	str, err := p.b.ReadBytes(0x0)
	str = str[0 : len(str)-1] // get rid of null terminator
	obj.value = string(str)
	return obj, err
}
