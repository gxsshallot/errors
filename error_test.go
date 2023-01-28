package errors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	msg, subMsg, data := "failed", "sub info", "attach data"
	errListManual := []*errorBase{
		Newf(1001, msg),
		Newf(1002, "%s %s", msg, "args"),
		NewFull(1003, msg, data, true),
		NewFull(1004, msg, data, false),
		NewWithError(1005, errors.New(msg)),
		NewSubLevel(2001, msg, 2001001, subMsg),
		NewSubLevelFull(2002, msg, 2002002, subMsg, data, true),
		NewSubLevelFull(2003, msg, 2003003, subMsg, data, false),
	}
	for _, err := range errListManual {
		if len(err.Error()) == 0 {
			t.Error("invalid Error() output")
			continue
		}
	}
	// map test
	code, subCode := 3000, 3000000
	GlobalCodes.Add(code, msg)
	GlobalSubCodes.Add(subCode, subMsg)
	errListAuto := []struct {
		Code            int
		SubCode         int
		ValidMessage    bool
		ValidSubMessage bool
	}{
		{code, 0, true, false},
		{code, subCode, true, true},
		{0, 0, false, false}, // split to delete code and subcode
		{code, 0, false, false},
		{code, subCode, false, false},
	}
	for _, item := range errListAuto {
		if item.Code == 0 && item.SubCode == 0 {
			GlobalCodes.Del(code)
			GlobalSubCodes.Del(subCode)
			continue
		}
		var err *errorBase
		if item.SubCode == 0 {
			err = New(item.Code)
		} else {
			err = NewSub(item.Code, item.SubCode)
		}
		if item.ValidMessage == (len(err.Message) == 0) {
			t.Error("message invalid")
			continue
		}
		if item.ValidSubMessage == (len(err.SubMessage) == 0) {
			t.Error("sub_message invalid")
			continue
		}
	}
	// revert test
	errListRevert := []struct {
		Err     error
		IsValid bool
	}{
		{nil, false},
		{New(code), true},
		{NewSub(code, subCode), true},
		{errors.New(msg), false},
	}
	for _, item := range errListRevert {
		target, ok := Revert(item.Err)
		if ok != item.IsValid {
			t.Error("revert failed")
			continue
		}
		if ok && target.Code != code {
			t.Error("revert code invalid")
			continue
		}
	}
}
