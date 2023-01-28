package errors

import (
	"reflect"
)

// convert error to *errorBase
func Revert(err error) (*errorBase, bool) {
	if err == nil {
		return nil, false
	}
	target := &errorBase{}
	val := reflect.ValueOf(&target)
	if reflect.TypeOf(err).AssignableTo(val.Type().Elem()) {
		val.Elem().Set(reflect.ValueOf(err))
		return target, true
	}
	return nil, false
}
