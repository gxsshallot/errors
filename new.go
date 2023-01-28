package errors

import "fmt"

func New(
	code int,
	message string,
) *errorBase {
	return newBase(code, message, 0, "", nil, EnableStack)
}

func Newf(
	code int,
	messageFormat string,
	messageArgs ...interface{},
) *errorBase {
	message := fmt.Sprintf(messageFormat, messageArgs...)
	return newBase(code, message, 0, "", nil, EnableStack)
}

func NewFull(
	code int,
	message string,
	attach interface{},
	enableStack bool,
) *errorBase {
	return newBase(code, message, 0, "", attach, enableStack)
}

func NewWithError(
	code int,
	err error,
) *errorBase {
	return newBase(code, err.Error(), 0, "", nil, EnableStack)
}
