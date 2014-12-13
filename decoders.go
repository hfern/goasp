package goasp

type decoder_cb func(*parser, pobject) (pobject, error)

var decoders map[otype]decoder_cb = map[otype]decoder_cb{
	tSTRING:       parse_length_string,
	tSTRING2:      parse_length_string,
	tSTR_NULLTERM: parse_str_nullterm,
	tUINT:         parse_uinteger,
}
