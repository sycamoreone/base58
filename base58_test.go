package base58

import (
	"bytes"
	"testing"
)

type test struct {
	decoded []byte
	encoded string
}

// Test cases have been take from the bitcoinj Java implementation of the Bitcoin protocol.
var testCases = []test{
	{ []byte("Hello World"), "JxF12TrwUP45BMd" },
	{ make([]byte, 1), "1" },
	{ make([]byte, 4), "1111" },
	{ make([]byte, 7), "1111111" },
	{ []byte{}, "" },
}

func TestEncode(t *testing.T) {
	for _, v := range testCases {
		if got := Encode(v.decoded); got != v.encoded {
			t.Logf("expected %s but got %s\n", v.encoded, got)
			t.Fail()
		}
	}
}

func TestDecode(t *testing.T) {
	for _, v := range testCases {
		if got, err := Decode(v.encoded); !bytes.Equal(got, v.decoded) || err != nil {
			t.Logf("expected %v but got %v\n", v.decoded, got)
			t.Fail()
		}
	}
}

func TestCorruptInputError(t *testing.T) {
	_, err := Decode("This isn't valid base58")
	if err == nil {
		t.Logf("decoded invalid data without returning non-nil error\n")
		t.Fail()
	}
}