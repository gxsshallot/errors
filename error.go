package errors

// enable or disable crash stack log
var EnableStack = true

// max stack depth for runtime caller
var MaxStackDepth = 100

type CODE interface {
	~int | ~string
}

/*
 * It is error struct to record business error information and crash stack.
 * Support recording two level of code-message.
 * Support recording a custom attach data.
 */
type errorBase[T CODE] struct {
	Code       T
	Message    string
	SubCode    T
	SubMessage string
	Attach     interface{}

	// crash stack log
	stack string
}

func newBase[T CODE](
	code T,
	message string,
	subCode T,
	subMessage string,
	attach interface{},
	enableStack bool,
) *errorBase[T] {
	return &errorBase[T]{
		Code:       code,
		Message:    message,
		SubCode:    subCode,
		SubMessage: subMessage,
		Attach:     attach,
		stack:      getStack(enableStack),
	}
}
