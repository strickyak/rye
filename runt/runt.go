package runt

import (
	"bytes"
	R "reflect"
	"strconv"

	. "github.com/strickyak/yak"
)

var None = &PNone{}
var True = &PBool{B: true}
var False = &PBool{B: false}

func VP(a interface {}) P {
	if a == nil {
		Say("VP", "<nil>")
		return nil
	}
	Say("VP", a)
	return a.(P)
}

func VSP(s string, a interface {}) P {
	if a == nil {
		Say("VSP", s, "<nil>")
		return nil
	}
	Say("VSP", s, a)
	return a.(P)
}

// P is the interface for every Pythonic value.
type P interface {
	Show() string
	String() string
	Repr() string

	Field(field string) P
	FieldGets(field string, x P) P
	FieldForCall(field string) P
	Call(aa ...P) P

	Len() int
	SetItem(i P, x P)
	DelItem(i P)
	GetItem(a P) P
	GetItemSlice(a, b, c P) P

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

	EQ(a P) P
	NE(a P) P
	LT(a P) P
	LE(a P) P
	GT(a P) P
	GE(a P) P

	Int() int64
	Float() float64
	Complex() complex128
}

// C_object is the root of inherited classes.
type C_object struct {
	PBase
	Rye_Self I_object
}

// I_object is the interface for C_object*.
type I_object interface {
	P
	C_object() *C_object
	Self() I_object
}

func (o *C_object) C_object() *C_object {
	return o
}
func (o *C_object) Self() I_object {
	return I_object(o)
}

type PBase struct {
}

func (o PBase) Field(field string) P          { panic(Bad("PBase cannot Field", o, field)) }
func (o PBase) FieldGets(field string, x P) P { panic(Bad("PBase cannot FieldGets", o, field, x)) }
func (o PBase) FieldForCall(field string) P   { panic(Bad("PBase cannot FieldForCall")) }
func (o PBase) Call(aa ...P) P                { panic(Bad("PBase cannot Call", o, aa)) }
func (o PBase) Len() int         { panic(Bad("PBase cannot Len: %#v")) }
func (o PBase) GetItem(a P) P                 { panic(Bad("PBase cannot GetItem", o, a)) }
func (o PBase) GetItemSlice(a, b, c P) P      { panic(Bad("PBase cannot GetItemSlice", o, a, b, c)) }
func (o PBase) SetItem(i P, x P) { panic(Bad("PBase cannot SetItem: %#v")) }
func (o PBase) DelItem(i P)      { panic(Bad("PBase cannot DelItem: %#v")) }


func (o PBase) Add(a P) P    { panic(Bad("PBase cannot Add: %#v", a)) }
func (o PBase) Sub(a P) P    { panic(Bad("PBase cannot Sub: %#v", a)) }
func (o PBase) Mul(a P) P    { panic(Bad("PBase cannot Mul: %#v", a)) }
func (o PBase) Div(a P) P    { panic(Bad("PBase cannot Div: %#v", a)) }
func (o PBase) Mod(a P) P    { panic(Bad("PBase cannot Mod: %#v", a)) }
func (o PBase) Pow(a P) P    { panic(Bad("PBase cannot Pow: %#v", a)) }
func (o PBase) And(a P) P    { panic(Bad("PBase cannot And: %#v", a)) }
func (o PBase) Or(a P) P     { panic(Bad("PBase cannot Or: %#v", a)) }
func (o PBase) Xor(a P) P    { panic(Bad("PBase cannot Xor: %#v", a)) }
func (o PBase) LShift(a P) P { panic(Bad("PBase cannot LShift: %#v", a)) }
func (o PBase) RShift(a P) P { panic(Bad("PBase cannot RShift: %#v", a)) }

func (o PBase) IAdd(a P) { panic(Bad("PBase cannot IAdd: %#v", a)) }
func (o PBase) ISub(a P) { panic(Bad("PBase cannot ISub: %#v", a)) }
func (o PBase) IMul(a P) { panic(Bad("PBase cannot IMul: %#v", a)) }

func (o PBase) EQ(a P) P { panic(Bad("PBase cannot EQ: %#v", a)) }
func (o PBase) NE(a P) P { panic(Bad("PBase cannot NE: %#v", a)) }
func (o PBase) LT(a P) P { panic(Bad("PBase cannot LT: %#v", a)) }
func (o PBase) LE(a P) P { panic(Bad("PBase cannot LE: %#v", a)) }
func (o PBase) GT(a P) P { panic(Bad("PBase cannot GT: %#v", a)) }
func (o PBase) GE(a P) P { panic(Bad("PBase cannot GE: %#v", a)) }

func (o PBase) Bool() bool { panic(Bad("PBase cannot Bool")) }
func (o PBase) Neg() P     { panic(Bad("PBase cannot Neg")) }
func (o PBase) Pos() P     { panic(Bad("PBase cannot Pos")) }
func (o PBase) Abs() P     { panic(Bad("PBase cannot Abs")) }
func (o PBase) Inv() P     { panic(Bad("PBase cannot Inv")) }

func (o PBase) Int() int64          { panic(Bad("PBase cannot Int")) }
func (o PBase) Float() float64      { panic(Bad("PBase cannot Float")) }
func (o PBase) Complex() complex128 { panic(Bad("PBase cannot Complex")) }

func (o PBase) String() string {
	return F("<%s:%u>", R.ValueOf(o).Type(), R.ValueOf(o).Addr().Pointer())
}
func (o PBase) Repr() string { return o.String() }
func (o PBase) Show() string { return o.String() }

type PInt struct {
	PBase
	N int64
}

type PBool struct {
	PBase
	B bool
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
func MkNone() *PNone          { return None }
func MkBool(b bool) *PBool {
	if b {
		return True
	} else {
		return False
	}
}

func (o *PBool) Bool() bool     { return o.B }
func (o *PBool) Int() int64     {
	if o.B {
		return 1
	} else {
		return 0
	}
}
func (o *PBool) String() string {
	if o.B {
		return "True"
	} else {
		return "False"
	}
}
func (o *PBool) Repr() string {
	if o.B {
		return "True"
	} else {
		return "False"
	}
}

func (o *PInt) Add(a P) P      { return MkInt(o.N + a.Int()) }
func (o *PInt) Sub(a P) P      { return MkInt(o.N - a.Int()) }
func (o *PInt) Mul(a P) P      { return MkInt(o.N * a.Int()) }
func (o *PInt) Div(a P) P      { return MkInt(o.N / a.Int()) }
func (o *PInt) Mod(a P) P      { return MkInt(o.N % a.Int()) }
func (o *PInt) And(a P) P      { return MkInt(o.N & a.Int()) }
func (o *PInt) Or(a P) P      { return MkInt(o.N | a.Int()) }
func (o *PInt) Xor(a P) P      { return MkInt(o.N ^ a.Int()) }
func (o *PInt) LShift(a P) P      { return MkInt(o.N << uint64(a.Int())) }
func (o *PInt) RShift(a P) P      { return MkInt(o.N >> uint64(a.Int())) }
func (o *PInt) EQ(a P) P       { return MkBool(o.N == a.Int()) }
func (o *PInt) NE(a P) P       { return MkBool(o.N != a.Int()) }
func (o *PInt) LT(a P) P       { return MkBool(o.N < a.Int()) }
func (o *PInt) LE(a P) P       { return MkBool(o.N <= a.Int()) }
func (o *PInt) GT(a P) P       { return MkBool(o.N > a.Int()) }
func (o *PInt) GE(a P) P       { return MkBool(o.N >= a.Int()) }
func (o *PInt) Int() int64     { return o.N }
func (o *PInt) String() string { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string   { return o.String() }
func (o *PInt) Bool() bool     { return o.N != 0 }

func (o PStr) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.S))
	}
	return MkStr(o.S[i:i+1])
}

func (o PStr) GetItemSlice(x, y, z P) P {
	var i, j int64
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.S))
		}
	}
	if y == None {
		j = int64(len(o.S))
	} else {
		j := y.Int()
		if j < 0 {
			j += int64(len(o.S))
		}
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	return MkStr(o.S[i:j])
}

func (o *PStr) Add(a P) P      { return MkStr(o.S + a.String()) }
func (o *PStr) EQ(a P) P       { return MkBool(o.S == a.String()) }
func (o *PStr) NE(a P) P       { return MkBool(o.S != a.String()) }
func (o *PStr) LT(a P) P       { return MkBool(o.S < a.String()) }
func (o *PStr) LE(a P) P       { return MkBool(o.S <= a.String()) }
func (o *PStr) GT(a P) P       { return MkBool(o.S > a.String()) }
func (o *PStr) GE(a P) P       { return MkBool(o.S >= a.String()) }
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

func F_len(a P) P { return MkInt(int64(a.Len())) }
var G_len = &PFunc1{ Fn: F_len }

type PFunc1 struct {
	PBase
	Fn func(a P) P
}
func (p *PFunc1) Call1(a1 P) P {
	return p.Fn(a1)
}

/*
func (p *PFunc) Call(aa ...P) P {
	return p.Fn(aa)
}
func (p *PFunc) Call0() P {
	return p.Call()
}
func (p *PFunc) Call1(a1 P) P {
	return p.Call(a1)
}
func (p *PFunc) Call2(a1, a2 P) P {
	return p.Call(a1, a2)
}
func (p *PFunc) Call3(a1, a2, a3 P) P {
	return p.Call(a1, a2, a3)
}
*/

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
	return MkGo(z)
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
