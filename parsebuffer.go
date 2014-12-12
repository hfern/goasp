package goasp

import (
	"bufio"
	"io"
)

type parser struct {
	b *bufio.Reader
}

func newParser(source io.Reader) parser {
	return parser{
		b: bufio.NewReader(source),
	}
}
