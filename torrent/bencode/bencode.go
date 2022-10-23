package bencode

import (
	"strconv"
)

type Box any
type List []Box
type Dictionary map[string]Box

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
	var list List
	var temp = buf[1:]

	for i := 1; i < len(temp); i++ {
		if temp[0] == 'e' {
			break
		}

		var val any

		temp, val = inferType(temp)
		list = append(list, val)
	}

	return temp[1:], list
}

func GetDict(buf []byte) (input []byte, output Dictionary) {
	dict := Dictionary{}
	var temp = buf[1:]

	for i := 1; i < len(temp); i++ {
		if temp[0] == 'e' {
			break
		}

		var val any
		var key string

		temp, key = GetString(temp)
		temp, val = inferType(temp)

		dict[key] = val
	}

	return temp[1:], dict
}

func inferType(buf []byte) (input []byte, output Box) {
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
