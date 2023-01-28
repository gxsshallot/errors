package errors

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(1000, "test")
	fmt.Println(err.stack)
}
