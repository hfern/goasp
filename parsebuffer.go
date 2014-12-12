package goasp

import (
	"bufio"
	"bytes"
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

func newParserFromBytes(source []byte) parser {
	return newParser(bytes.NewBuffer(source))
}

func newParserFromString(source string) parser {
	return newParserFromBytes([]byte(source))
}
