package goasp

import (
	"testing"
)

func assertUint(t *testing.T, encoded string, expected uint, token otype) {
	prs := newParserFromString(encoded)
	obj, err := parse_uinteger(&prs, pobject{})

	if err != nil {
		t.Fatal(err)
	}

	var val uint
	var ok bool

	if val, ok = obj.value.(uint); !ok {
		t.Fatalf("pobject(%v)'s value should be a castable to uint.", otype_names[token])
	}

	if val != expected {
		t.Fatalf("Expected %v, got %v", expected, val)
	}
}

func TestDecode_uinteger_singleZero(t *testing.T) {
	assertUint(t, "\x00", 0, tUINT)

}

func TestDecode_uinteger_extendedZeroZero(t *testing.T) {
	assertUint(t, "\x80\x00", 0, tUINT)

}

func TestDecode_uinteger_normalOne(t *testing.T) {
	assertUint(t, "\x01", 1, tUINT)
}

func TestDecode_uinteger_normalExtension(t *testing.T) {
	assertUint(t, "\x81\x01", 129, tUINT)
}
