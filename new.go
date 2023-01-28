package errors

import "fmt"

func New[T CODE](
	code T,
	message string,
) *errorBase[T] {
	var zero T
	return newBase(code, message, zero, "", nil, EnableStack)
}

func Newf[T CODE](
	code T,
	messageFormat string,
	messageArgs ...interface{},
) *errorBase[T] {
	var zero T
	message := fmt.Sprintf(messageFormat, messageArgs...)
	return newBase(code, message, zero, "", nil, EnableStack)
}

func NewFull[T CODE](
	code T,
	message string,
	attach interface{},
	enableStack bool,
) *errorBase[T] {
	var zero T
	return newBase(code, message, zero, "", attach, enableStack)
}
