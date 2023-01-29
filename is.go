package errors

import (
	"reflect"
)

func Is(err error) bool {
	if err == nil {
		return false
	}
	target := &errorBase{}
	return reflect.TypeOf(err) == reflect.TypeOf(target)
}
