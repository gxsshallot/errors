package errors

func NewSub(code int, subCode int) *errorBase {
	return newBase(code, GlobalCodes.Get(code), subCode, GlobalSubCodes.Get(subCode), EnableStack)
}

func NewSubLevel(
	code int,
	message string,
	subCode int,
	subMessage string,
) *errorBase {
	return newBase(code, message, subCode, subMessage, EnableStack)
}

func NewSubLevelFull(
	code int,
	message string,
	subCode int,
	subMessage string,
	enableStack bool,
) *errorBase {
	return newBase(code, message, subCode, subMessage, enableStack)
}
