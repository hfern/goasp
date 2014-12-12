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

	if obj.token != token {
		gotTokenName, _ := otype_names[obj.token]
		t.Fatalf(
			"Expected %v (0x%02x), got %v (0x%02x)",
			otype_names[token],
			token,
			gotTokenName,
			obj.token,
		)
	}

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

}
