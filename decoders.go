package goasp

type decoder_cb func(*parser, pobject) (pobject, error)

var decoders map[otype]decoder_cb = map[otype]decoder_cb{
	tSTR_NULLTERM: parse_str_nullterm,
}
