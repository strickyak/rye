package runt

import (
	"bytes"
	R "reflect"
	"strconv"

	. "github.com/strickyak/yak"
)

var None = MkNone()

type P interface {
	Show() string
	String() string
	Repr() string

	Field(field string) P
	FieldGets(field string, x P) P
	Call(aa ...P) P

	Add(a P) P
	Sub(a P) P
	Mul(a P) P
	Div(a P) P
	Mod(a P) P
	Pow(a P) P
	And(a P) P
	Or(a P) P
	Xor(a P) P
	LShift(a P) P
	RShift(a P) P

	IAdd(a P) // +=
	ISub(a P) // -=
	IMul(a P) // *=

	Bool() bool // a.k.a. nonzero()
	Neg() P
	Pos() P
	Abs() P
	Inv() P

	EQ(a P) bool
	NE(a P) bool
	LT(a P) bool
	LE(a P) bool
	GT(a P) bool
	GE(a P) bool

	Int() int64
	Float() float64
	Complex() complex128

	Len() int
	GetItem(i P) P
	SetItem(i P, x P)
	DelItem(i P)
}

type PBase struct {
}

func (o PBase) Field(field string) P          { panic(Bad("cannot Field", o, field)) }
func (o PBase) FieldGets(field string, x P) P { panic(Bad("cannot FieldGets", o, field, x)) }
func (o PBase) Call(aa ...P) P                { panic(Bad("cannot Call", o, aa)) }

func (o PBase) Add(a P) P    { panic(F("cannot Add: %#v", a)) }
func (o PBase) Sub(a P) P    { panic("cannot Sub: %#v") }
func (o PBase) Mul(a P) P    { panic("cannot Mul: %#v") }
func (o PBase) Div(a P) P    { panic("cannot Div: %#v") }
func (o PBase) Mod(a P) P    { panic("cannot Mod: %#v") }
func (o PBase) Pow(a P) P    { panic("cannot Pow: %#v") }
func (o PBase) And(a P) P    { panic("cannot And: %#v") }
func (o PBase) Or(a P) P     { panic("cannot Or: %#v") }
func (o PBase) Xor(a P) P    { panic("cannot Xor: %#v") }
func (o PBase) LShift(a P) P { panic("cannot LShift: %#v") }
func (o PBase) RShift(a P) P { panic("cannot RShift: %#v") }

func (o PBase) IAdd(a P) { panic("cannot IAdd: %#v") }
func (o PBase) ISub(a P) { panic("cannot ISub: %#v") }
func (o PBase) IMul(a P) { panic("cannot IMul: %#v") }

func (o PBase) EQ(a P) bool { panic("cannot EQ: %#v") }
func (o PBase) NE(a P) bool { panic("cannot NE: %#v") }
func (o PBase) LT(a P) bool { panic("cannot LT: %#v") }
func (o PBase) LE(a P) bool { panic("cannot LE: %#v") }
func (o PBase) GT(a P) bool { panic("cannot GT: %#v") }
func (o PBase) GE(a P) bool { panic("cannot GE: %#v") }

func (o PBase) Bool() bool { panic("cannot Bool: %#v") }
func (o PBase) Neg() P     { panic("cannot Neg: %#v") }
func (o PBase) Pos() P     { panic("cannot Pos: %#v") }
func (o PBase) Abs() P     { panic("cannot Abs: %#v") }
func (o PBase) Inv() P     { panic("cannot Inv: %#v") }

func (o PBase) Int() int64          { panic("cannot Int: %#v") }
func (o PBase) Float() float64      { panic("cannot Float: %#v") }
func (o PBase) Complex() complex128 { panic("cannot Complex: %#v") }

func (o PBase) String() string {
	return F("<%s:%u>", R.ValueOf(o).Type(), R.ValueOf(o).Addr().Pointer())
}
func (o PBase) Repr() string { return o.String() }
func (o PBase) Show() string { return o.String() }

func (o PBase) Len() int         { panic("cannot Len: %#v") }
func (o PBase) GetItem(i P) P    { panic("cannot GetItem: %#v") }
func (o PBase) SetItem(i P, x P) { panic("cannot SetItem: %#v") }
func (o PBase) DelItem(i P)      { panic("cannot DelItem: %#v") }

type PInt struct {
	PBase
	N int64
}

type PFloat struct {
	PBase
	F float64
}

type PStr struct {
	PBase
	S string
}

type PGo struct {
	PBase
	V R.Value
}

type PList struct {
	PBase
	PP []P
}

type PNone struct {
	PBase
}

type Scope map[string]P

type PDict struct {
	PBase
	PPP Scope
}

type PFunc struct {
	PBase
	Fn func(args []P) P
}

type PObj struct {
	PBase
	Obj interface{}
}

func MkP(a Any) P {
	if a == nil {
		return None
	}
	switch x := a.(type) {
	case int:
		return Mkint(x)
	case int64:
		return MkInt(x)
	case string:
		return MkStr(x)
	}
	return MkGo(a)
}

func MkGo(a Any) *PGo { return &PGo{V: R.ValueOf(a)} }

func Mkint(n int) *PInt       { return &PInt{N: int64(n)} }
func MkInt(n int64) *PInt     { return &PInt{N: n} }
func MkStr(s string) *PStr    { return &PStr{S: s} }
func MkList(pp []P) *PList    { return &PList{PP: pp} }
func MkDict(ppp Scope) *PDict { return &PDict{PPP: ppp} }
func MkNone() *PNone          { return &PNone{} }

func (o *PInt) Add(a P) P      { return MkInt(o.N + a.Int()) }
func (o *PInt) Int() int64     { return o.N }
func (o *PInt) String() string { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string   { return o.String() }

func (o *PStr) Add(a P) P      { return MkStr(o.S + a.String()) }
func (o *PStr) Int() int64     { return CI(strconv.ParseInt(o.S, 10, 64)) }
func (o *PStr) String() string { return o.S }
func (o *PStr) Len() int       { return len(o.S) }
func (o *PStr) Repr() string   { return F("%q", o.S) }

func (o *PList) Len() int      { return len(o.PP) }
func (o *PList) GetItem(a P) P { return o.PP[a.Int()] }
func (o *PList) Repr() string {
	buf := bytes.NewBufferString("[ ")
	for i := 0; i < len(o.PP); i++ {
		buf.WriteString(o.PP[i].Repr())
		buf.WriteString(", ")
	}
	buf.WriteString("]")
	return buf.String()
}

func (o *PDict) Len() int      { return len(o.PPP) }
func (o *PDict) GetItem(a P) P { return o.PPP[a.String()] }

func NewList() *PList {
	return &PList{PP: make([]P, 0)}
}

func CopyList(aa *PList) *PList {
	zz := make([]P, 0)
	copy(zz, aa.PP)
	return &PList{PP: zz}
}

func RepeatList(aa *PList, n int64) *PList {
	zz := NewList()
	for i := int64(0); i < n; i++ {
		zz.PP = append(zz.PP, aa.PP...)
	}
	return zz
}

func Enlist(args ...P) *PList {
	zz := make([]P, 0)
	for _, a := range args {
		zz = append(zz, a)
	}
	return MkList(zz)
}

func (oo *PList) Append(aa P) {
	oo.PP = append(oo.PP, aa)
}

func (oo *PList) AppendElements(aa *PList) {
	oo.PP = append(oo.PP, aa.PP...)
}

func MaybeDeref(t R.Value) R.Value {
	for t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	return t
}

func (p *PFunc) Call(aa ...P) P {
	return p.Fn(aa)
}

/*
func (g *PGo) Call(aa...P) P {
	f := MaybeDeref(g.V)
	if f.Kind() != R.Func {
		Bad("cannot Call when Value not a func", f)
	}
	t := f.Type()
	if t.IsVariadic() {
		Bad("cannot call Variadic functions (yet)")
	}
	numIn := t.NumIn()
	numOut := t.NumOut()
	if len(aa) != numIn {
		Bad("call got %d args, want %d args", numIn, len(aa))
	}
	args := make([]R.Value, len(aa))
	for i, a := range aa {
		args[i] = AdaptToType(a, t.In(i))
	}
	outs := f.Call(args)
	// TODO: strip off error.
	switch numOut {
	case 0:
		return None
	case 1:
		return MkP(outs[0].Interface())
	}
	panic(Bad("Multi-arg returns no imp yet"))
}
*/

/*
func AdaptToType(v P, t R.Type) R.Value {
	z = Zero(t)
	if v == nil {
		return z
	}
	switch v.Kind() {
	case R.Int:
		return
	}
}
*/

func (g *PGo) Field(field string) P {
	t := MaybeDeref(g.V)
	for t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	if t.Kind() != R.Struct {
		Bad("cannot Field when Value not a struct", t)
	}
	z := t.FieldByName(field)
	if !z.IsValid() {
		Bad("field not found", t, field)
	}
	return MkP(z.Interface())
}

func init() {
	var tmp P = new(PBase)
	// Demonstrate these things implement P.
	tmp = new(PInt)
	tmp = new(PFloat)
	tmp = new(PList)
	tmp = new(PDict)
	_ = tmp
}
