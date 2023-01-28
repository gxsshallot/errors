package errors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	msg, subMsg, data := "failed", "sub info", "attach data"
	errListInt := []*errorBase[int]{
		New(1001, msg),
		Newf(1002, "%s %s", msg, "args"),
		NewFull(1003, msg, data, true),
		NewFull(1003, msg, data, false),
		NewSubLevel(2001, msg, 2001001, subMsg),
		NewSubLevelFull(2002, msg, 2001002, subMsg, data, true),
		NewSubLevelFull(2002, msg, 2001002, subMsg, data, false),
	}
	for _, err := range errListInt {
		if e := testErr[int](err); e != nil {
			t.Error(e)
		}
	}
	errListStr := []*errorBase[string]{
		New("1001", msg),
		Newf("1002", "%s %s", msg, "args"),
		NewFull("1003", msg, data, true),
		NewFull("1003", msg, data, false),
		NewSubLevel("2001", msg, "2001001", subMsg),
		NewSubLevelFull("2002", msg, "2001002", subMsg, data, true),
		NewSubLevelFull("2002", msg, "2001002", subMsg, data, false),
	}
	for _, err := range errListStr {
		if e := testErr[string](err); e != nil {
			t.Error(e)
		}
	}
}

func testErr[T CODE](err error) error {
	output := err.Error()
	if len(output) == 0 {
		return errors.New("invalid Error() output")
	}
	return nil
}
