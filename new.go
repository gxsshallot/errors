package errors

import "fmt"

func New(code int) *errorBase {
	return newBase(code, GlobalCodes.Get(code), 0, "", EnableStack)
}

func Newf(
	code int,
	messageFormat string,
	messageArgs ...interface{},
) *errorBase {
	var message string
	if len(messageArgs) > 0 {
		message = fmt.Sprintf(messageFormat, messageArgs...)
	} else {
		message = messageFormat
	}
	return newBase(code, message, 0, "", EnableStack)
}

func NewFull(
	code int,
	message string,
	enableStack bool,
) *errorBase {
	return newBase(code, message, 0, "", enableStack)
}

func NewWithError(
	code int,
	err error,
) *errorBase {
	return newBase(code, err.Error(), 0, "", EnableStack)
}
