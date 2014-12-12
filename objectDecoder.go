package goasp

import (
	"errors"
)

var (
	LacksPreamble = errors.New("Source lacks viewstate preamble")
)

// decode_full attempts to decode the binary-serialized
// viewstate code. Checks for preamble.
func decode_full(st *parser) (node pobject, err error) {
	preamble, err := st.b.Peek(2)

	if err != nil {
		return pobject{}, err
	}

	if string(preamble) != string(view_state_preamble) {
		return pobject{}, LacksPreamble
	}

	// consume preamble
	st.b.ReadByte()
	st.b.ReadByte()

	return decode_obj(st)
}

// decode_obj tries to decode an object using
// its ID byte. Consumes the buffer until the cursor is
// before a byte NOT in the object
func decode_obj(st *parser) (node pobject, err error) {
	var symbol byte
	var handler decoder_cb

	if symbol, err = st.b.ReadByte(); err != nil {
		return
	}

	tok := otype(symbol)
	var ok bool

	if handler, ok = decoders[tok]; !ok {
		err = UnknownObjectIdError{token: tok}
		return
	}

	node, err = handler(st, pobject{token: tok})
	return
}
