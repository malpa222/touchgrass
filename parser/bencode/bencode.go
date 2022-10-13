package bencode

import (
	"strconv"
)

type Box interface{}
type List []Box
type Dictionary map[string]Box

func GetString(buf []byte) (input []byte, output string) {
	length := ""

	for i, b := range buf {
		if b == ':' {
			temp := i + 1
			strlen, err := strconv.Atoi(length)

			if err == nil {
				strlen += temp
				return buf[strlen:], string(buf[temp:strlen])
			}
		}

		length += string(b)
	}

	return buf, ""
}

func GetInt(buf []byte) (input []byte, output int) {
	if buf[0] == 'i' {
		str := ""

		for i := 1; i < len(buf); i++ {
			if buf[i] == 'e' {
				num, err := strconv.Atoi(str)

				if err == nil {
					return buf[i+1:], num
				}
			}

			str += string(buf[i])
		}
	}

	return nil, 0
}

func GetList(data []byte) (input []byte, output *List) {
	// l4:spam4:eggse => ['spam', 'eggs']

	return nil, nil
}

func GetDict(data []byte) (input []byte, output *Dictionary) {
	// d3:cow3:moo4:spam4:eggse =>  {'cow': 'moo', 'spam': 'eggs'}
	// d4:spaml1:a1:bee => {'spam': ['a', 'b']}

	return nil, nil
}
