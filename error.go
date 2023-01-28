package errors

import "fmt"

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
	Stack string
}

func (e errorBase[T]) Error() (result string) {
	var zero T
	result += fmt.Sprintf("code=%v message=%s", e.Code, e.Message)
	if e.SubCode != zero {
		result += fmt.Sprintf("\nsubcode=%v submessage=%s", e.SubCode, e.SubMessage)
	}
	if len(e.Stack) > 0 {
		result += fmt.Sprintf("\n\nstacks:\n%s", e.Stack)
	}
	return
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
		Stack:      getStack(enableStack),
	}
}
