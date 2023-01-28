package errors

import "fmt"

// enable or disable crash stack log
var EnableStack = true

// max stack depth for runtime caller
var MaxStackDepth = 100

// It is error struct to record business error information and crash stack.
type errorBase struct {
	Code       int
	Message    string
	SubCode    int
	SubMessage string
	Attach     map[string]string

	// crash stack log
	Stack string
}

func (e errorBase) Error() (result string) {
	result += fmt.Sprintf("code=%v message=%s", e.Code, e.Message)
	if e.SubCode != 0 {
		result += fmt.Sprintf("\nsubcode=%v submessage=%s", e.SubCode, e.SubMessage)
	}
	if len(e.Stack) > 0 {
		result += fmt.Sprintf("\n\nstacks:\n%s", e.Stack)
	}
	return
}

func newBase(
	code int,
	message string,
	subCode int,
	subMessage string,
	enableStack bool,
) *errorBase {
	return &errorBase{
		Code:       code,
		Message:    message,
		SubCode:    subCode,
		SubMessage: subMessage,
		Attach:     map[string]string{},
		Stack:      getStack(enableStack),
	}
}
