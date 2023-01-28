package errors

import (
	"errors"
)

// convert error to *errorBase
func Revert(err error) (*errorBase, bool) {
	target := &errorBase{}
	ok := errors.As(err, &target)
	return target, ok
}
