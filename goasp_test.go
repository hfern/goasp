package goasp

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"testing"
)

func TestDecode_SrjcComplex(t *testing.T) {
	encoded, err := ioutil.ReadFile("./testdata/srjc_schedule.b64.txt")
	if err != nil {
		t.Fatal(err)
	}

	binary, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		t.Fatal(err)
	}

	_, err = Decode(bytes.NewBuffer(binary))

	if err != nil {
		t.Fatal(err)
	}

}
