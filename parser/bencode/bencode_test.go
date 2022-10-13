package bencode

import (
	"bytes"
	"testing"
)

func TestGetString(t *testing.T) {
	str1 := []byte("6:qwerty")
	str2 := []byte("6:qwertyi14e")

	want1 := "qwerty"
	want2 := []byte("i14e")

	in, out := GetString(str1)

	if len(in) != 0 || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}

	in, out = GetString(str2)

	if !bytes.Equal(in, want2) || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", want2, want1, in, out)
	}
}

func TestGetInt(t *testing.T) {
	str1 := []byte("i13e")
	str2 := []byte("i13e4:test")

	want1 := 13
	want2 := []byte("4:test")

	in, out := GetInt(str1)

	if len(in) != 0 || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}

	in, out = GetInt(str2)

	if !bytes.Equal(in, want2) || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", want2, want1, in, out)
	}
}
