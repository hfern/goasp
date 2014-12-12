package goasp

import (
	"testing"
)

func TestDecode_str_nullterm(t *testing.T) {
	encoded := "\x0bA Test String\x00"

	prs := newParserFromString(encoded)
	obj, err := decode_obj(&prs)

	if err != nil {
		t.Fatal(err)
	}

	if obj.token != tSTR_NULLTERM {
		t.Fatalf("Expected tSTR_NULLTERM (0x%02x), got 0x%02x", tSTR_NULLTERM, obj.token)
	}

	var val string
	var ok bool

	if val, ok = obj.value.(string); !ok {
		t.Fatal("pobject tSTR_NULLTERM's value should be a string.")
	}

	expected := "A Test String"

	if val != expected {
		t.Fatal("Expected \"" + expected + "\", got \"" + val + "\"")
	}
}
