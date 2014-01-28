package runt

import (
	"bytes"
	. "fmt"
	"os"
	R "reflect"
	"strconv"
)

type Any interface{}

type P interface {
	Show() string
	String() string
	Repr() string

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

func (o PBase) Add(a P) P    { panic("cannot Add: %#v") }
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
	return Sprintf("<%s:%u>", R.ValueOf(o).Type(), R.ValueOf(o).Addr().Pointer())
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

type Scope map[string]P

type PDict struct {
	PBase
	PPP Scope
}

type PFunc struct {
	PBase
	F func(args []P) P
}

func Mkint(n int) *PInt       { return &PInt{N: int64(n)} }
func MkInt(n int64) *PInt     { return &PInt{N: n} }
func MkStr(s string) *PStr    { return &PStr{S: s} }
func MkList(pp []P) *PList    { return &PList{PP: pp} }
func MkDict(ppp Scope) *PDict { return &PDict{PPP: ppp} }

func (o *PInt) Add(a P) P      { return MkInt(o.N + a.Int()) }
func (o *PInt) Int() int64     { return o.N }
func (o *PInt) String() string { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string   { return o.String() }

func (o *PStr) Add(a P) P      { return MkStr(o.S + a.String()) }
func (o *PStr) Int() int64     { return CI(strconv.ParseInt(o.S, 10, 64)) }
func (o *PStr) String() string { return o.S }
func (o *PStr) Len() int       { return len(o.S) }
func (o *PStr) Repr() string   { return Sprintf("%q", o.S) }

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

func Show(aa ...Any) string {
	buf := bytes.NewBuffer(nil)
	for _, a := range aa {
		switch x := a.(type) {
		case string:
			buf.WriteString(Sprintf("%q ", x))
		case []byte:
			buf.WriteString(Sprintf("%q ", string(x)))
		case int:
			buf.WriteString(Sprintf("%d ", x))
		case int64:
			buf.WriteString(Sprintf("%d ", x))
		case float32:
			buf.WriteString(Sprintf("%f ", x))
		case Stringer:
			buf.WriteString(Sprintf("%s ", x))
		default:
			v := R.ValueOf(a)
			switch v.Kind() {
			case R.Slice:
				n := v.Len()
				buf.WriteString(Sprintf("%d[ ", n))
				for i := 0; i < n; i++ {
					buf.WriteString(Show(v.Index(i).Interface()))
				}
				buf.WriteString("] ")
			default:
				buf.WriteString(Sprintf("{%s:%s:%v} ", R.ValueOf(x).Kind(), R.ValueOf(x).Type(), x))
			}
		}
	}
	return buf.String()
}
func Say(aa ...Any) {
	Fprintf(os.Stderr, "## %s\n", Show(aa...))
}

func Ci(x int, err error) int {
	if err != nil {
		panic(err)
	}
	return x
}

func CI(x int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return x
}

func Cs(x string, err error) string {
	if err != nil {
		panic(err)
	}
	return x
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
