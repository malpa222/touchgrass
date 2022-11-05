package cast

import (
	"errors"
	"fmt"
	"reflect"
)

func To[T any](in any) (T, error) {
	switch in.(type) {
	case T:
		return in.(T), nil
	default:
		var t T // stupid hack
		return nil, getErrString(reflect.TypeOf(t), in)
	}
}

func ToList[T any](in any) ([]T, error) {
	switch in.(type) {
	case T:
		return in.([]T), nil
	default:
		return nil, getErrString(reflect.TypeOf([]T{}), in)
	}
}

func ToDict[T1 comparable, T2 any](in any) (map[T1]T2, error) {
	switch in.(type) {
	case map[T1]T2:
		return in.(map[T1]T2), nil
	default:
		return nil, getErrString(reflect.TypeOf(map[T1]T2{}), in)
	}
}

func getErrString(wanted any, received any) error {
	return errors.New(fmt.Sprintf("cast failed, invalid types:\n"+
		"wanted: %v\n"+
		"received: %v",
		wanted, reflect.TypeOf(received)))
}
