package goasp

import (
	"errors"
)

var MissingFirstByte = errors.New("Missing first byte following object type ID.")

func decode_uinteger_val(p *parser) (uint, error) {
	val := uint(0)

	shift := uint(0)
	cont := true

	var bits byte
	var err error

	for err == nil && cont {
		bits, err = p.b.ReadByte()
		cont = bits & ^byte(0x7F) > 0
		val += uint(bits&0x7F) << shift
		shift += 7
	}

	if err != nil {
		return 0, err
	}

	return val, nil
}

func parse_uinteger(p *parser, obj pobject) (pobject, error) {
	val, err := decode_uinteger_val(p)
	if err != nil {
		return obj, err
	}

	obj.value = val
	return obj, nil
}
