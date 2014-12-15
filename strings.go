package goasp

func parse_length_string(p *parser, obj pobject) (pobject, error) {
	length, err := decode_uinteger_val(p)
	if err != nil {
		return obj, err
	}

	strbuf := make([]byte, length)

	for i := uint(0); i < length; i++ {
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
	if len(str) > 0 {
		str = str[0 : len(str)-1] // get rid of null terminator
	}
	obj.value = string(str)
	return obj, err
}

func parse_str_array(p *parser, obj pobject) (pobject, error) {
	length, err := decode_uinteger_val(p)
	if err != nil {
		return obj, err
	}

	str_array := make([]string, 0, length)

	for i := uint(0); i < length; i++ {
		node, err := parse_length_string(p, pobject{token: tSTRING})
		if err != nil {
			return obj, err
		}

		var v string
		var ok bool

		if v, ok = node.value.(string); !ok {
			panic("Type Error: stringNode.value should be cast-able to string")
		}

		str_array = append(str_array, v)
	}

	obj.value = str_array

	return obj, nil
}
