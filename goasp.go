package goasp

import (
	"io"
)

// Encoder encodes values to ASP.net
type Encoder struct {
}

type Decoder struct {
}

func Decode(source io.Reader) (interface{}, error) {
	prs := newParser(source)
	return decode_full(&prs)
}
