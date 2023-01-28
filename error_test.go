package errors

import (
	"testing"
)

func TestNew(t *testing.T) {
	msg, subMsg, data := "failed", "sub info", "attach data"
	errList := []*errorBase{
		New(1001, msg),
		Newf(1002, "%s %s", msg, "args"),
		NewFull(1003, msg, data, true),
		NewFull(1003, msg, data, false),
		NewSubLevel(2001, msg, 2001001, subMsg),
		NewSubLevelFull(2002, msg, 2001002, subMsg, data, true),
		NewSubLevelFull(2002, msg, 2001002, subMsg, data, false),
	}
	for _, err := range errList {
		if len(err.Error()) == 0 {
			t.Error("invalid Error() output")
		}
	}
}
