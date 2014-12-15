package goasp

import (
	"io"
	"testing"
)

func assertDecodedString(t *testing.T, encoded, expected string, token otype, canError *error) parser {
	prs := newParserFromString(encoded)
	obj, err := decode_obj(&prs)

	if err != nil {
		if canError != nil && err == *canError {
			return prs
		} else {
			t.Fatal(err)
		}
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

	return prs
}

func TestDecode_str_nullterm(t *testing.T) {
	encoded := "\x0bA Test String\x00"
	decoded := "A Test String"
	assertDecodedString(t, encoded, decoded, tSTR_NULLTERM, nil)
	assertDecodedString(t, "\x0bThis is missing Nullterm", "", tSTR_NULLTERM, &io.EOF)
}

func TestDecode_length_string(t *testing.T) {
	assertDecodedString(t, "\x05\x0CHello World!", "Hello World!", tSTRING, nil)
	assertDecodedString(t, "\x1E\x0CHello World!", "Hello World!", tSTRING2, nil)
	assertDecodedString(t, "\x05\x00", "", tSTRING, nil)
	assertDecodedString(t, "\x05\x01\x07", "\x07", tSTRING, nil)
	assertDecodedString(t, "\x05\x09Too Shor", "", tSTRING, &io.EOF)
}

func TestDecode_string_array(t *testing.T) {
	encoded := "\x15\x05\x01A\x01B\x01C\x01D\x01E" // ABCDE
	expected := []string{"A", "B", "C", "D", "E"}

	prs := newParserFromString(encoded)
	obj, err := decode_obj(&prs)
	if err != nil {
		t.Fatal(err)
	}

	assertTokenIs(t, obj, tSTRING_ARRAY)

	var got []string
	var ok bool

	if got, ok = obj.value.([]string); !ok {
		t.Fatal("String Array node.value should be cast-able to []string")
	}

	if len(expected) != len(got) {
		t.Fatalf(
			"Length of expected != length of got! Expected %v[%v], Got: %v[%v]",
			len(expected),
			expected,
			len(got),
			got,
		)
	}

	mismatches := 0

	for i := 0; i < len(expected); i++ {
		if expected[i] != got[i] {
			t.Log("el[%v] mismatch: expected \"%v\", got \"%v\"", i, expected[i], got[i])
			mismatches++
		}
	}

	if mismatches > 0 {
		t.Fatalf("%v mismatches -- test failed.", mismatches)
	}
}

func TestDecode_string_array_fails_on_eof(t *testing.T) {
	encoded := "\x15\x05\x01A\x01B\x01C\x01D" // ABCD, expects 5th element, missing
	prs := newParserFromString(encoded)
	obj, err := decode_obj(&prs)
	assertTokenIs(t, obj, tSTRING_ARRAY)
	if err != io.EOF {
		t.Fatal("Decoding Str Array should have failed when missing items.")
	}
}
