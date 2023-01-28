// +build !go1.7

package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// fix: runtime.CallersFrames not support when go < 1.7
func getStack(enableStack bool) (stack string) {
	if !enableStack {
		return
	}
	ptr := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(3, ptr) // skip "runtime.Callers", "getStack", "newBase"
	for i := 0; i < length; i++ {
		ptrItem := ptr[i]
		itemFunc := runtime.FuncForPC(ptrItem)
		fileName, fileLine := itemFunc.FileLine(ptrItem)
		stack += fmt.Sprintf("%s:\n\t%s:%d (0x%x)\n", itemFunc.Name(), fileName, fileLine, itemFunc.Entry())
	}
	stack = strings.TrimSuffix(stack, "\n")
	return
}
