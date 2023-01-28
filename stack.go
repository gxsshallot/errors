package errors

import (
	"fmt"
	"runtime"
	"strings"
)

func getStack(enableStack bool) (stack string) {
	if !enableStack {
		return
	}
	ptr := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(3, ptr) // skip "runtime.Callers", "getStack", "newBase"
	frames := runtime.CallersFrames(ptr[:length])
	for {
		item, more := frames.Next()
		if !more {
			break
		}
		stack += fmt.Sprintf("%s:\n\t%s:%d (0x%x)\n", item.Function, item.File, item.Line, item.Entry)
	}
	stack = strings.TrimSuffix(stack, "\n")
	return
}
