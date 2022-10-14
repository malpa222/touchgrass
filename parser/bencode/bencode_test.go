package bencode

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGetStringSimple(t *testing.T) {
	str1 := []byte("6:qwerty")

	want1 := "qwerty"

	in, out := GetString(str1)

	if len(in) != 0 || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestGetStringTrailing(t *testing.T) {
	str1 := []byte("6:qwertyi14e")

	want1 := "qwerty"
	want2 := []byte("i14e")

	in, out := GetString(str1)

	if !bytes.Equal(in, want2) || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", want2, want1, in, out)
	}
}

func TestGetInt(t *testing.T) {
	int1 := []byte("i13e")

	want1 := 13

	in, out := GetInt(int1)

	if len(in) != 0 || out != want1 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestGetIntTrailing(t *testing.T) {
	int1 := []byte("i13e4:test")

	want1 := 13
	want2 := []byte("4:test")

	in, out := GetInt(int1)

	if !bytes.Equal(in, want2) || out != want1 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want2, want1, in, out)
	}
}

func TestGetListSimple(t *testing.T) {
	testVal1 := []byte("li32e7:andrzeje")

	out, list := GetList(testVal1)

	var want1 []byte
	want2 := List{32, "andrzej"}

	if !bytes.Equal(out, want1) || !reflect.DeepEqual(&want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", &want2, want1, list, out)
	}

}

func TestGetListNested(t *testing.T) {
	testVal1 := []byte("li32e4:testli14ee6:qwertye")

	out, list := GetList(testVal1)

	var want1 []byte
	want2 := List{32, "test", List{14}, "qwerty"}

	if !bytes.Equal(out, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want2, want1, list, out)
	}
}
