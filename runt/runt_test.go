package runt_test

import (
	"testing"

	. "github.com/strickyak/rye/runt"
	. "github.com/strickyak/yak"
)

func Test_PStr_GetItemSlice(t *testing.T) {
	MustEq("xbcd", MkStr("abcdefg").GetItemSlice(MkInt(1), MkInt(4), None).String())
}
