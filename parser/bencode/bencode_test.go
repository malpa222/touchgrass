package bencode

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGetStringSimple(t *testing.T) {
	str1 := []byte("6:qwerty")
	in, out := GetString(str1)

	want1 := "qwerty"

	if len(in) != 0 || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestGetStringTrailing(t *testing.T) {
	str1 := []byte("6:qwertyi14e")
	in, out := GetString(str1)

	want1 := []byte("i14e")
	want2 := "qwerty"

	if !bytes.Equal(in, want1) || out != want2 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", want1, want2, in, out)
	}
}

func TestGetInt(t *testing.T) {
	int1 := []byte("i13e")
	in, out := GetInt(int1)

	want1 := 13

	if len(in) != 0 || out != want1 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestGetIntTrailing(t *testing.T) {
	int1 := []byte("i13e4:test")
	in, out := GetInt(int1)

	want1 := []byte("4:test")
	want2 := 13

	if !bytes.Equal(in, want1) || out != want2 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, out)
	}
}

func TestGetListSimple(t *testing.T) {
	testVal1 := []byte("li32e7:andrzeje")
	in, list := GetList(testVal1)

	var want1 []byte
	want2 := List{32, "andrzej"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, list)
	}

}

func TestGetListNested(t *testing.T) {
	testVal1 := []byte("li32e4:testli14ee6:qwertye")
	in, list := GetList(testVal1)

	var want1 []byte
	want2 := List{32, "test", List{14}, "qwerty"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, list, in)
	}
}

func TestGetListTrailing(t *testing.T) {
	testVal1 := []byte("li32e4:testli14ee6:qwertyei123e")
	in, list := GetList(testVal1)

	want1 := []byte("i123e")
	want2 := List{32, "test", List{14}, "qwerty"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, list, in)
	}
}

func TestGetDict(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key26:qwertye")
	in, dict := GetDict(testVal1)

	var want1 []byte
	want2 := Dictionary{
		"key1": 15,
		"key2": "qwerty",
	}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, dict) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, dict)
	}
}

func TestGetDictNested(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key2d4:key1li13e4:testeee")
	in, dict := GetDict(testVal1)

	var want1 []byte
	want2 := Dictionary{
		"key1": 15,
		"key2": Dictionary{
			"key1": List{13, "test"},
		},
	}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, dict) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, dict)
	}
}

func TestGetDictTrailing(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key2d4:key1li13e4:testeeeli420e4:teste")
	in, dict := GetDict(testVal1)

	want1 := []byte("li420e4:teste")
	want2 := Dictionary{
		"key1": 15,
		"key2": Dictionary{
			"key1": List{13, "test"},
		},
	}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, dict) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, dict)
	}
}
