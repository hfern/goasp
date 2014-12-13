package goasp

import (
	"testing"
)

func assertDecodedString(t *testing.T, encoded, expected string, token otype) {
	prs := newParserFromString(encoded)
	obj, err := decode_obj(&prs)

	if err != nil {
		t.Fatal(err)
	}

	assertTokenIs(t, obj, token)

	var val string
	var ok bool

	if val, ok = obj.value.(string); !ok {
		t.Fatalf("pobject(%v)'s value should be a castable to string.", otype_names[token])
	}

	if val != expected {
		t.Fatal("Expected \"" + expected + "\", got \"" + val + "\"")
	}
}

func TestDecode_str_nullterm(t *testing.T) {
	encoded := "\x0bA Test String\x00"
	decoded := "A Test String"
	assertDecodedString(t, encoded, decoded, tSTR_NULLTERM)
}

func TestDecode_length_string(t *testing.T) {
	assertDecodedString(t, "\x05\x0CHello World!", "Hello World!", tSTRING)
	assertDecodedString(t, "\x1E\x0CHello World!", "Hello World!", tSTRING2)
	assertDecodedString(t, "\x05\x00", "", tSTRING)
	assertDecodedString(t, "\x05\x01\x07", "\x07", tSTRING)
}
