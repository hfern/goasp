package goasp

import (
	"testing"
)

func assertTokenIs(t *testing.T, obj pobject, token otype) {
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
}
