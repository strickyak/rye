package rye_test

import (
	"regexp"
	"testing"

	. "github.com/strickyak/rye/rye__"
)

func MustEq(a, b string) {
	if a != b {
		println("MustEq fails: a=", a, "   b=", b)
	}
}

func Test_PStr_GetItem(t *testing.T) {
	MustEq("a", JString(MkStr("abcdefg").GetItem(MkInt(0))))
	MustEq("g", JString(MkStr("abcdefg").GetItem(MkInt(6))))
	MustEq("g", JString(MkStr("abcdefg").GetItem(MkInt(-1))))
}
func Test_PStr_GetItemSlice(t *testing.T) {
	MustEq("bcd", JString(MkStr("abcdefg").GetItemSlice(MkInt(1), MkInt(4), None)))
	MustEq("bcdef", JString(MkStr("abcdefg").GetItemSlice(MkInt(1), MkInt(-1), None)))
	MustEq("abcd", JString(MkStr("abcdefg").GetItemSlice(None, MkInt(-3), None)))
	MustEq("cdefg", JString(MkStr("abcdefg").GetItemSlice(MkInt(2), None, None)))
	MustEq("abcdefg", JString(MkStr("abcdefg").GetItemSlice(None, None, None)))
}

const P1 = "/home/strick/gocode/src/github.com/strickyak/rye/rye__CMmty/z/rye_module.go"

var M2 = regexp.MustCompile(`^(.*/src/)(.+)/(rye__[A-Za-z0-9]*)/([^/]+)/rye_module[.]go$`)

func Test_MatchGoFilenameToRyeFilenameOrEmpty(t *testing.T) {
	m := M2.FindStringSubmatch(P1)
	println("m", len(m), m)
	a, b := MatchGoFilenameToRyeFilenameOrEmpty(P1)
	println("a", a)
	println("b", b)
	if a != `/home/strick/gocode/src/github.com/strickyak/rye/z.py` {
		t.Errorf("BAD %q %q", a, b)
	}
}
