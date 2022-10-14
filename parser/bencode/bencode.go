package bencode

import (
	"strconv"
)

type List []any
type Dictionary map[string]any

func GetString(buf []byte) (input []byte, output string) {
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

func GetInt(buf []byte) (input []byte, output int) {
	if buf[0] != 'i' {
		return nil, 0
	}

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

func GetList(buf []byte) (input []byte, output List) {
	if buf[0] != 'l' {
		return nil, nil
	}

	var list List
	var temp = buf[1:]

	for i := 1; i < len(temp); i++ {
		if temp[0] == 'e' {
			break
		}

		var val any

		temp, val = InferType(temp)
		list = append(list, val)
	}

	return temp[1:], list
}

func GetDict(buf []byte) (input []byte, output Dictionary) {
	if buf[0] != 'd' {
		return nil, nil
	}
	return nil, nil
}

func InferType(buf []byte) (input []byte, output any) {
	switch buf[0] {
	case 'i':
		return GetInt(buf)
	case 'l':
		return GetList(buf)
	case 'd':
		return GetDict(buf)
	default:
		if _, err := strconv.Atoi(string(buf[0])); err == nil {
			return GetString(buf)
		}
	}

	return buf, buf
}
