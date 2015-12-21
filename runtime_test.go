package rye_test

import (
	"testing"

	. "github.com/strickyak/rye"
)

func Test_PStr_GetItem(t *testing.T) {
	MustEq("a", MkStr("abcdefg").Self.GetItem(MkInt(0)).Self.String())
	MustEq("g", MkStr("abcdefg").Self.GetItem(MkInt(6)).Self.String())
	MustEq("g", MkStr("abcdefg").Self.GetItem(MkInt(-1)).Self.String())
}
func Test_PStr_GetItemSlice(t *testing.T) {
	MustEq("bcd", MkStr("abcdefg").Self.GetItemSlice(MkInt(1), MkInt(4), None).Self.String())
	MustEq("bcdef", MkStr("abcdefg").Self.GetItemSlice(MkInt(1), MkInt(-1), None).Self.String())
	MustEq("abcd", MkStr("abcdefg").Self.GetItemSlice(None, MkInt(-3), None).Self.String())
	MustEq("cdefg", MkStr("abcdefg").Self.GetItemSlice(MkInt(2), None, None).Self.String())
	MustEq("abcdefg", MkStr("abcdefg").Self.GetItemSlice(None, None, None).Self.String())
}
