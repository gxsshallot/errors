package errors

func NewSub(code int, subCode int) *errorBase {
	return newBase(code, GlobalCodes.Get(code), subCode, GlobalSubCodes.Get(subCode), nil, EnableStack)
}

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
