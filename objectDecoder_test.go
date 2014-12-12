package goasp

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func Test_ComplexSRJC(t *testing.T) {
	encoded, err := ioutil.ReadFile("./testdata/srjc_schedule.b64.txt")
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	t.Log(encoded)

	binary, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}

}
