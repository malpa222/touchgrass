package bencode

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestStringSimple(t *testing.T) {
	str1 := []byte("6:qwerty")
	in, out := Decode(str1)

	want1 := "qwerty"

	if len(in) != 0 || out != want1 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestStringTrailing(t *testing.T) {
	str1 := []byte("6:qwertyi14e")
	in, out := Decode(str1)

	want1 := []byte("i14e")
	want2 := "qwerty"

	if !bytes.Equal(in, want1) || out != want2 {
		t.Errorf("Expected: %v, %v\nGot: %v, %v", want1, want2, in, out)
	}
}

func TestInt(t *testing.T) {
	int1 := []byte("i13e")
	in, out := Decode(int1)

	want1 := 13

	if len(in) != 0 || out != want1 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", 0, want1, len(in), out)
	}
}

func TestIntTrailing(t *testing.T) {
	int1 := []byte("i13e4:test")
	in, out := Decode(int1)

	want1 := []byte("4:test")
	want2 := 13

	if !bytes.Equal(in, want1) || out != want2 {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, out)
	}
}

func TestListSimple(t *testing.T) {
	testVal1 := []byte("li32e7:andrzeje")
	in, list := Decode(testVal1)

	var want1 []byte
	want2 := List{32, "andrzej"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, list)
	}

}

func TestListNested(t *testing.T) {
	testVal1 := []byte("li32e4:testli14ee6:qwertye")
	in, list := Decode(testVal1)

	var want1 []byte
	want2 := List{32, "test", List{14}, "qwerty"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, list, in)
	}
}

func TestListTrailing(t *testing.T) {
	testVal1 := []byte("li32e4:testli14ee6:qwertyei123e")
	in, list := Decode(testVal1)

	want1 := []byte("i123e")
	want2 := List{32, "test", List{14}, "qwerty"}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, list) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, list, in)
	}
}

func TestDict(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key26:qwertye")
	in, dict := Decode(testVal1)

	var want1 []byte
	want2 := Dictionary{
		"key1": 15,
		"key2": "qwerty",
	}

	if !bytes.Equal(in, want1) || !reflect.DeepEqual(want2, dict) {
		t.Errorf("\nExpected: %v, %v\nGot: %v, %v", want1, want2, in, dict)
	}
}

func TestDictNested(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key2d4:key1li13e4:testeee")
	in, dict := Decode(testVal1)

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

func TestDictTrailing(t *testing.T) {
	testVal1 := []byte("d4:key1i15e4:key2d4:key1li13e4:testeeeli420e4:teste")
	in, dict := Decode(testVal1)

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

func TestEncodeInt(t *testing.T) {
	want := "i123456e"
	res, _ := Encode(123456)

	if res != want {
		t.Errorf("\nThe strings don't match.\nExpected: %s\nGot: %s", want, res)
	}
}

func TestEncodeString(t *testing.T) {
	test := "tentekstma18znakow"
	want := fmt.Sprintf("%d:%s", len(test), test)

	res, _ := Encode(test)

	if res != want {
		t.Errorf("\nThe strings don't match.\nExpected: %s\nGot: %s", want, res)
	}
}

func TestEncodeList(t *testing.T) {
	str1 := "andrzejmatus"
	str2 := "szioq"
	test := List{12345, str1, str2}
	want := fmt.Sprintf("li%de%d:%s%d:%se", test[0], len(str1), str1, len(str2), str2)

	res, err := Encode(test)
	if res != want || err != nil {
		t.Errorf("\nThe strings don't match.\nExpected: %s\nGot: %s", want, res)
		t.Errorf("%v", err)
	}
}

func TestEncodeDict(t *testing.T) {
	str1 := "andrzejmatus"
	str2 := "szioq"
	test := Dictionary{
		"key1": str1,
		"key2": str2,
		"key3": 123456,
	}
	want := fmt.Sprintf("d4:key1%d:%s4:key2%d:%s4:key3i%dee", len(str1), str1, len(str2), str2, test["key3"])

	res, err := Encode(test)
	if res != want || err != nil {
		t.Errorf("\nThe strings don't match.\nExpected: %s\nGot: %s", want, res)
		t.Errorf("%v", err)
	}
}

func TestEncodeDictEmbedList(t *testing.T) {
	str1 := "andrzejmatus"
	list := List{123456, str1}
	test := Dictionary{
		"key1": str1,
		"key2": list,
		"key3": 123456,
	}
	want := fmt.Sprintf("d4:key1%d:%s4:key2li%de%d:%se4:key3i%dee",
		len(str1), str1, list[0], len(str1), str1, test["key3"])

	res, err := Encode(test)
	if res != want || err != nil {
		t.Errorf("\nThe strings don't match.\nExpected: %s\nGot: %s", want, res)
		t.Errorf("%v", err)
	}
}

func TestEncodeStruct(t *testing.T) {
	_, err := Encode(bytes.Buffer{})
	if err == nil {
		t.Errorf("\nExpected an error, got nil")
	}
}

func TestToBytes(t *testing.T) {
	_, dict := Decode([]byte("d4:key1i15e4:key2d4:key1li13e4:testeeeli420e4:teste"))
	if rawBytes, err := ToBytes(dict); err != nil {
		t.Errorf("\nAaaawaria:\n%#v", err)
	} else {
		fmt.Printf("%#v", rawBytes)
	}
}
