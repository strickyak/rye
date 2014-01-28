package main

import (
	. "fmt"
	. "github.com/strickyak/rye/runt"
)

func main() {
	x := MkInt(39)
	y := MkInt(3)
	Printf("%v\n", x.Add(y))

	a := MkStr("hello")
	b := MkStr("world")
	Printf("%q\n", a.Add(b))

	l := Enlist(x, y, x.Add(y), b)
	i := MkInt(1)
	Printf("%v\n", l.GetItem(i))
	Printf("%v\n", l.GetItem(y))
}
