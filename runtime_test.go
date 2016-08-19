package rye_test

import (
	"testing"

	. "github.com/strickyak/rye"
)

func MustEq(a, b string) {
	if a != b {
		println("MustEq fails: a=", a, "   b=", b)
	}
}

func Test_PStr_GetItem(t *testing.T) {
	MustEq("a", MkStr("abcdefg").GetItem(MkInt(0)).String())
	MustEq("g", MkStr("abcdefg").GetItem(MkInt(6)).String())
	MustEq("g", MkStr("abcdefg").GetItem(MkInt(-1)).String())
}
func Test_PStr_GetItemSlice(t *testing.T) {
	MustEq("bcd", MkStr("abcdefg").GetItemSlice(MkInt(1), MkInt(4), None).String())
	MustEq("bcdef", MkStr("abcdefg").GetItemSlice(MkInt(1), MkInt(-1), None).String())
	MustEq("abcd", MkStr("abcdefg").GetItemSlice(None, MkInt(-3), None).String())
	MustEq("cdefg", MkStr("abcdefg").GetItemSlice(MkInt(2), None, None).String())
	MustEq("abcdefg", MkStr("abcdefg").GetItemSlice(None, None, None).String())
}
