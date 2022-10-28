package bencode

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Decode(buf []byte) (input []byte, output any) {
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

func Encode(data any) (string, error) {
	switch v := data.(type) {
	case int:
		return fmt.Sprintf("i%de", v), nil
	case string:
		return fmt.Sprintf("%d:%s", len(v), v), nil
	case []any:
		var temp string
		for _, elem := range v {
			result, _ := Encode(elem)
			temp += result
		}

		return fmt.Sprintf("l%se", temp), nil
	case map[string]any:
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
			reflect.TypeOf(v), reflect.Int, reflect.String, reflect.TypeOf([]any{}), reflect.TypeOf(map[string]any{}))

		return "", errors.New(msg)
	}
}

func getString(buf []byte) (input []byte, output string) {
	length := ""

	for i, b := range buf {
		if b != ':' {
			length += string(b)
			continue
		}

		if num, err := strconv.Atoi(length); err == nil {
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

func getList(buf []byte) (input []byte, output []any) {
	var list []any
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

func getDict(buf []byte) (input []byte, output map[string]any) {
	dict := map[string]any{}
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
