package errors

func NewSubLevel[T CODE](
	code T,
	message string,
	subCode T,
	subMessage string,
) *errorBase[T] {
	return newBase(code, message, subCode, subMessage, nil, EnableStack)
}

func NewSubLevelFull[T CODE](
	code T,
	message string,
	subCode T,
	subMessage string,
	attach interface{},
	enableStack bool,
) *errorBase[T] {
	return newBase(code, message, subCode, subMessage, attach, enableStack)
}
