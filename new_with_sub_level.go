package errors

func NewWithSubLevel[T CODE](
	code T,
	message string,
	subCode T,
	subMessage string,
) *errorBase[T] {
	return newBase(code, message, subCode, subMessage, nil, EnableStack)
}

func NewWithSubLevelFull[T CODE](
	code T,
	message string,
	subCode T,
	subMessage string,
	attach interface{},
	enableStack bool,
) *errorBase[T] {
	return newBase(code, message, subCode, subMessage, attach, enableStack)
}
