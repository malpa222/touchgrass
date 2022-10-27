package bencode

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type Box any
type List []Box
type Dictionary map[string]Box

func Decode(buf []byte) (input []byte, output Box) {
	switch buf[0] {
	case 'i':
		return getInt(buf)
	case 'l':
		return getList(buf)
	case 'd':
		return getDict(buf)
	default:
		if _, err := strconv.Atoi(string(buf[0])); err == nil {
			return getString(buf)
		}
	}

	return buf, buf
}

func Encode(data Box) (string, error) {
	switch v := data.(type) {
	case int:
		return fmt.Sprintf("i%de", v), nil
	case string:
		return fmt.Sprintf("%d:%s", len(v), v), nil
	case []Box:
	case List:
		var temp string
		for _, elem := range v {
			result, _ := Encode(elem)
			temp += result
		}

		return fmt.Sprintf("l%se", temp), nil
	case map[string]Box:
	case Dictionary:
		var temp string
		for key, elem := range v {
			result, _ := Encode(key)
			temp += result

			result, _ = Encode(elem)
			temp += result
		}

		return fmt.Sprintf("d%se", temp), nil
	default:
		msg := fmt.Sprintf("Cannot encode the type %v.\n Supported types: %v, %v, %v, %v",
			reflect.TypeOf(v), reflect.Int, reflect.String, reflect.TypeOf(List{}), reflect.TypeOf(Dictionary{}))

		return "", errors.New(msg)
	}

	return "", nil
}

func ToBytes(box Box) (out *[]byte, err error) {
	var buf bytes.Buffer

	gob.Register(List{})
	gob.Register(Dictionary{})

	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(box); err != nil {
		return nil, err
	}

	temp := buf.Bytes()
	return &temp, nil
}

func getString(buf []byte) (input []byte, output string) {
	length := ""

	for i, b := range buf {
		if b != ':' {
			length += string(b)
			continue
		}

		num, err := strconv.Atoi(length)
		if err != nil {
			temp := i + 1
			num += temp

			return buf[num:], string(buf[temp:num])
		}
	}

	return buf, ""
}

func getInt(buf []byte) (input []byte, output int) {
	str := ""

	for i := 1; i < len(buf); i++ {
		if buf[i] != 'e' {
			str += string(buf[i])
			continue
		}

		if num, err := strconv.Atoi(str); err == nil {
			return buf[i+1:], num
		}
	}

	return buf, 0
}

func getList(buf []byte) (input []byte, output List) {
	var list List
	var temp = buf[1:]

	for i := 1; i < len(temp); i++ {
		if temp[0] == 'e' {
			break
		}

		var val any

		temp, val = Decode(temp)
		list = append(list, val)
	}

	return temp[1:], list
}

func getDict(buf []byte) (input []byte, output Dictionary) {
	dict := Dictionary{}
	var temp = buf[1:]

	for i := 1; i < len(temp); i++ {
		if temp[0] == 'e' {
			break
		}

		var val any
		var key string

		temp, key = getString(temp)
		temp, val = Decode(temp)

		dict[key] = val
	}

	return temp[1:], dict
}
