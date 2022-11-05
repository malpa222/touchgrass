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
		return t, getErr(t, in)
	}
}

func ToList[T any](in any) ([]T, error) {
	if val, ok := in.([]T); ok {
		return val, nil
	} else {
		var t []T
		return nil, getErr(t, in)
	}
}

func ToDict[T1 comparable, T2 any](in any) (map[T1]T2, error) {
	if val, ok := in.(map[T1]T2); ok {
		return val, nil
	} else {
		var t map[T1]T2
		return nil, getErr(t, in)
	}
}

func getErr(wanted any, received any) error {
	return errors.New(fmt.Sprintf("cast failed, invalid types:\n"+
		"wanted: %v\n"+
		"received: %v",
		reflect.TypeOf(wanted), reflect.TypeOf(received)))
}
