package errors

func NewSubLevel(
	code int,
	message string,
	subCode int,
	subMessage string,
) *errorBase {
	return newBase(code, message, subCode, subMessage, nil, EnableStack)
}

func NewSubLevelFull(
	code int,
	message string,
	subCode int,
	subMessage string,
	attach interface{},
	enableStack bool,
) *errorBase {
	return newBase(code, message, subCode, subMessage, attach, enableStack)
}
