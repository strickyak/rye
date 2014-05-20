package runt

import (
	"bytes"
	"errors"
	"os"
	R "reflect"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	. "github.com/strickyak/yak"
)

var None = &PNone{}
var True = &PBool{B: true}
var False = &PBool{B: false}

func init() {
	None.Self = None
	True.Self = True
	False.Self = False
}

var RyeEnv string
var Debug int

func init() {
	RyeEnv := os.Getenv("RYE")
	for _, ch := range RyeEnv {
		if ch == 'd' {
			Debug++
		}
	}
}

func VP(a interface{}) P {
	if Debug < 1 {
		return a.(P)
	}
	if a == nil {
		Say("VP", "<nil>")
		return nil
	}
	Say("VP", a)
	if Debug >= 3 {
		debug.PrintStack()
	}
	return a.(P)
}

func VSP(s string, a interface{}) P {
	if Debug < 1 {
		return a.(P)
	}
	if a == nil {
		Say("VSP", s, "<nil>")
		return nil
	}
	Say("VSP", s, a)
	if Debug >= 3 {
		debug.PrintStack()
	}
	return a.(P)
}

// P is the interface for every Pythonic value.
type P interface {
	Show() string
	String() string
	Repr() string
	Type() P
	Is(a P) bool
	IsNot(a P) bool
	GetSelf() P
	SetSelf(a P)

	Field(field string) P
	FieldGets(field string, x P) P
	FieldForCall(field string) P
	Call(aa ...P) P
	Iter() Nexter
	List() []P

	Len() int
	SetItem(i P, x P)
	DelItem(i P)
	GetItem(a P) P
	GetItemSlice(a, b, c P) P
	Contains(a P) bool    // Reverse "in"
	NotContains(a P) bool // Reverse "not in"

	Add(a P) P
	Sub(a P) P
	Mul(a P) P
	Div(a P) P
	IDiv(a P) P
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
}

// C_object is the root of inherited classes.
type C_object struct {
	PBase
	////Rye_Self I_object
}

////// I_object is the interface for C_object*.
////type I_object interface {
////	P
////	PtrC_object() *C_object
////	MySelf() I_object
////}

func (o *C_object) PtrC_object() *C_object {
	return o
}

////func (o *C_object) MySelf() I_object {
////	return I_object(o)
////}

type PBase struct {
	Self P
}

func (o *PBase) GetSelf() P           { return o.Self }
func (o *PBase) SetSelf(a P)          { o.Self = a }
func (o *PBase) Field(field string) P { panic(Bad("Receiver cannot Field", o.Self, o, field)) }
func (o *PBase) FieldGets(field string, x P) P {
	panic(Bad("Receiver cannot FieldGets", o.Self, o, field, x))
}
func (o *PBase) FieldForCall(field string) P { panic(Bad("Receiver cannot FieldForCall", o.Self)) }
func (o *PBase) Call(aa ...P) P              { panic(Bad("Receiver cannot Call", o.Self, o, aa)) }
func (o *PBase) Len() int                    { panic(Bad("Receiver cannot Len: ", o.Self)) }
func (o *PBase) GetItem(a P) P               { panic(Bad("Receiver cannot GetItem", o.Self, o, a)) }
func (o *PBase) GetItemSlice(a, b, c P) P {
	panic(Bad("Receiver cannot GetItemSlice", o.Self, o, a, b, c))
}
func (o *PBase) Is(a P) bool          { return o.GetSelf() == a.GetSelf() }
func (o *PBase) IsNot(a P) bool       { return o.GetSelf() != a.GetSelf() }
func (o *PBase) Contains(a P) bool    { panic(Bad("Receiver cannot Contains: ", o.Self)) }
func (o *PBase) NotContains(a P) bool { panic(Bad("Receiver cannot NotContains: ", o.Self)) }
func (o *PBase) SetItem(i P, x P)     { panic(Bad("Receiver cannot SetItem: ", o.Self)) }
func (o *PBase) DelItem(i P)          { panic(Bad("Receiver cannot DelItem: ", o.Self)) }
func (o *PBase) Iter() Nexter         { panic(Bad("Receiver cannot Iter: ", o.Self)) }
func (o *PBase) List() []P            { panic(Bad("Receiver cannot List: ", o.Self)) }

func (o *PBase) Add(a P) P    { panic(Bad("Receiver cannot Add: ", o.Self, a)) }
func (o *PBase) Sub(a P) P    { panic(Bad("Receiver cannot Sub: ", o.Self, a)) }
func (o *PBase) Mul(a P) P    { panic(Bad("Receiver cannot Mul: ", o.Self, a)) }
func (o *PBase) Div(a P) P    { panic(Bad("Receiver cannot Div: ", o.Self, a)) }
func (o *PBase) IDiv(a P) P   { panic(Bad("Receiver cannot IDiv: ", o.Self, a)) }
func (o *PBase) Mod(a P) P    { panic(Bad("Receiver cannot Mod: ", o.Self, a)) }
func (o *PBase) Pow(a P) P    { panic(Bad("Receiver cannot Pow: ", o.Self, a)) }
func (o *PBase) And(a P) P    { panic(Bad("Receiver cannot And: ", o.Self, a)) }
func (o *PBase) Or(a P) P     { panic(Bad("Receiver cannot Or: ", o.Self, a)) }
func (o *PBase) Xor(a P) P    { panic(Bad("Receiver cannot Xor: ", o.Self, a)) }
func (o *PBase) LShift(a P) P { panic(Bad("Receiver cannot LShift: ", o.Self, a)) }
func (o *PBase) RShift(a P) P { panic(Bad("Receiver cannot RShift: ", o.Self, a)) }

func (o *PBase) IAdd(a P) { panic(Bad("Receiver cannot IAdd: ", o.Self, a)) }
func (o *PBase) ISub(a P) { panic(Bad("Receiver cannot ISub: ", o.Self, a)) }
func (o *PBase) IMul(a P) { panic(Bad("Receiver cannot IMul: ", o.Self, a)) }

func (o *PBase) EQ(a P) bool { return P(o) == a }
func (o *PBase) NE(a P) bool { return P(o) != a }
func (o *PBase) LT(a P) bool { panic(Bad("Receiver cannot LT: ", o.Self, a)) }
func (o *PBase) LE(a P) bool { panic(Bad("Receiver cannot LE: ", o.Self, a)) }
func (o *PBase) GT(a P) bool { panic(Bad("Receiver cannot GT: ", o.Self, a)) }
func (o *PBase) GE(a P) bool { panic(Bad("Receiver cannot GE: ", o.Self, a)) }

func (o *PBase) Bool() bool { panic(Bad("Receiver cannot Bool", o.Self)) }
func (o *PBase) Neg() P     { panic(Bad("Receiver cannot Neg", o.Self)) }
func (o *PBase) Pos() P     { panic(Bad("Receiver cannot Pos", o.Self)) }
func (o *PBase) Abs() P     { panic(Bad("Receiver cannot Abs", o.Self)) }
func (o *PBase) Inv() P     { panic(Bad("Receiver cannot Inv", o.Self)) }

func (o *PBase) Int() int64          { panic(Bad("Receiver cannot Int", o.Self)) }
func (o *PBase) Float() float64      { panic(Bad("Receiver cannot Float", o.Self)) }
func (o *PBase) Complex() complex128 { panic(Bad("Receiver cannot Complex", o.Self)) }

func (o *PBase) String() string {
	if o.Self == nil {
		panic("PBase:  Why is o.Self NIL?")
	}
	return F("<%#v>", o.Self)
}
func (o *PBase) Repr() string { return o.String() }
func (o *PBase) Show() string { return o.Self.String() }
func (o *PBase) Type() P      { return MkStr(F("%t", o.Self)) }

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

type PListIter struct {
	PBase
	PP []P
	I  int
}

type PTuple struct {
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

func MkP(a interface{}) P {
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

func MkGo(a interface{}) *PGo { z := &PGo{V: R.ValueOf(a)}; z.Self = z; return z }

func Mkint(n int) *PInt         { z := &PInt{N: int64(n)}; z.Self = z; return z }
func MkInt(n int64) *PInt       { z := &PInt{N: n}; z.Self = z; return z }
func MkFloat(f float64) *PFloat { z := &PFloat{F: f}; z.Self = z; return z }
func MkStr(s string) *PStr      { z := &PStr{S: s}; z.Self = z; return z }

func MkList(pp []P) *PList    { z := &PList{PP: pp}; z.Self = z; return z }
func MkTuple(pp []P) *PTuple  { z := &PTuple{PP: pp}; z.Self = z; return z }
func MkDict(ppp Scope) *PDict { z := &PDict{PPP: ppp}; z.Self = z; return z }
func MkDictFromPairs(pp []P) *PDict {
	z := &PDict{PPP: make(Scope)}
	z.Self = z
	for _, x := range pp {
		sub := x.List()
		if len(sub) != 2 {
			Bad("Expected sublist of size 2, but got size %d", len(sub))
		}
		k := sub[0].String()
		v := sub[1]
		z.PPP[k] = v
	}
	return z
}

func MkListV(pp ...P) *PList   { z := &PList{PP: pp}; z.Self = z; return z }
func MkTupleV(pp ...P) *PTuple { z := &PTuple{PP: pp}; z.Self = z; return z }
func MkDictV(pp ...P) *PDict {
	if (len(pp) & 1) == 1 {
		panic("MkDictV got odd len(pp)")
	}
	zzz := make(Scope)
	for i := 0; i < len(pp); i += 2 {
		zzz[pp[i].String()] = pp[i+1]
	}
	z := &PDict{PPP: zzz}
	z.Self = z
	return z
}

func MkNone() *PNone { return None }
func MkBool(b bool) *PBool {
	if b {
		return True
	} else {
		return False
	}
}

func (o *PNone) Bool() bool     { return false }
func (o *PNone) String() string { return "None" }
func (o *PNone) Repr() string   { return "None" }

func (o *PBool) Bool() bool { return o.B }
func (o *PBool) Int() int64 {
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
func (o *PBool) Type() P { return B_bool }

func (o *PInt) Add(a P) P      { return MkInt(o.N + a.Int()) }
func (o *PInt) Sub(a P) P      { return MkInt(o.N - a.Int()) }
func (o *PInt) Mul(a P) P      { return MkInt(o.N * a.Int()) }
func (o *PInt) Div(a P) P      { return MkInt(o.N / a.Int()) }
func (o *PInt) Mod(a P) P      { return MkInt(o.N % a.Int()) }
func (o *PInt) And(a P) P      { return MkInt(o.N & a.Int()) }
func (o *PInt) Or(a P) P       { return MkInt(o.N | a.Int()) }
func (o *PInt) Xor(a P) P      { return MkInt(o.N ^ a.Int()) }
func (o *PInt) LShift(a P) P   { return MkInt(o.N << uint64(a.Int())) }
func (o *PInt) RShift(a P) P   { return MkInt(o.N >> uint64(a.Int())) }
func (o *PInt) EQ(a P) bool    { return (o.N == a.Int()) }
func (o *PInt) NE(a P) bool    { return (o.N != a.Int()) }
func (o *PInt) LT(a P) bool    { return (o.N < a.Int()) }
func (o *PInt) LE(a P) bool    { return (o.N <= a.Int()) }
func (o *PInt) GT(a P) bool    { return (o.N > a.Int()) }
func (o *PInt) GE(a P) bool    { return (o.N >= a.Int()) }
func (o *PInt) Int() int64     { return o.N }
func (o *PInt) Float() float64 { return float64(o.N) }
func (o *PInt) String() string { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string   { return o.String() }
func (o *PInt) Bool() bool     { return o.N != 0 }
func (o *PInt) Type() P        { return B_int }

func (o *PFloat) Add(a P) P      { return MkFloat(o.F + a.Float()) }
func (o *PFloat) Sub(a P) P      { return MkFloat(o.F - a.Float()) }
func (o *PFloat) Mul(a P) P      { return MkFloat(o.F * a.Float()) }
func (o *PFloat) Div(a P) P      { return MkFloat(o.F / a.Float()) }
func (o *PFloat) EQ(a P) bool    { return (o.F == a.Float()) }
func (o *PFloat) NE(a P) bool    { return (o.F != a.Float()) }
func (o *PFloat) LT(a P) bool    { return (o.F < a.Float()) }
func (o *PFloat) LE(a P) bool    { return (o.F <= a.Float()) }
func (o *PFloat) GT(a P) bool    { return (o.F > a.Float()) }
func (o *PFloat) GE(a P) bool    { return (o.F >= a.Float()) }
func (o *PFloat) Int() int64     { return int64(o.F) }
func (o *PFloat) Float() float64 { return o.F }
func (o *PFloat) String() string { return strconv.FormatFloat(o.F, 'g', -1, 64) }
func (o *PFloat) Repr() string   { return o.String() }
func (o *PFloat) Bool() bool     { return o.F != 0 }
func (o *PFloat) Type() P        { return B_float }

func (o PStr) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.S))
	}
	return MkStr(o.S[i : i+1])
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
		j = y.Int()
		if j < 0 {
			j += int64(len(o.S))
		}
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkStr(o.S[i:j])
	return r
}

func (o *PStr) Mod(a P) P {
	switch t := a.(type) {
	case *PTuple:
		return MkStr(strings.Repeat(o.S, int(t.Int())))
	case *PInt:
		return MkStr(F(o.S, t.N))
	case *PStr:
		return MkStr(F(o.S, t.S))
	}
	panic(Badf("Not Imp: str %% %t", a))
}

func (o *PStr) Mul(a P) P {
	switch t := a.(type) {
	case *PInt:
		return MkStr(strings.Repeat(o.S, int(t.Int())))
	}
	panic(Badf("Cannot multiply: str * %t", a))
}
func (o *PStr) NotContains(a P) bool { return !o.Contains(a) }
func (o *PStr) Contains(a P) bool {
	switch t := a.(type) {
	case *PStr:
		return strings.Contains(o.S, t.S)
	}
	panic(Bad("string cannot Contain non-string:", a))
}
func (o *PStr) Add(a P) P      { return MkStr(o.S + a.String()) }
func (o *PStr) EQ(a P) bool    { return (o.S == a.String()) }
func (o *PStr) NE(a P) bool    { return (o.S != a.String()) }
func (o *PStr) LT(a P) bool    { return (o.S < a.String()) }
func (o *PStr) LE(a P) bool    { return (o.S <= a.String()) }
func (o *PStr) GT(a P) bool    { return (o.S > a.String()) }
func (o *PStr) GE(a P) bool    { return (o.S >= a.String()) }
func (o *PStr) Int() int64     { return CI(strconv.ParseInt(o.S, 10, 64)) }
func (o *PStr) String() string { return o.S }
func (o *PStr) Len() int       { return len(o.S) }
func (o *PStr) Repr() string   { return F("%q", o.S) }
func (o *PStr) Type() P        { return B_str }

func (o *PTuple) Len() int       { return len(o.PP) }
func (o *PTuple) GetItem(a P) P  { return o.PP[a.Int()] }
func (o *PTuple) String() string { return o.Repr() }
func (o *PTuple) Type() P        { return B_tuple }
func (o *PTuple) Repr() string {
	buf := bytes.NewBufferString("(")
	n := len(o.PP)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(o.PP[i].Repr())
	}
	if n == 0 {
		buf.WriteString(", ") // Special just for singleton tuples.
	}
	buf.WriteString(")")
	return buf.String()
}
func (o *PTuple) Iter() Nexter {
	z := &PListIter{PP: o.PP}
	z.Self = z
	return z
}
func (o *PTuple) List() []P {
	return o.PP
}

func (o *PList) NotContains(a P) bool { return !o.Contains(a) }
func (o *PList) Contains(a P) bool {
	for _, x := range o.PP {
		if a.EQ(x) {
			return true
		}
	}
	return false
}
func (o *PList) Len() int       { return len(o.PP) }
func (o *PList) GetItem(a P) P  { return o.PP[a.Int()] }
func (o *PList) String() string { return o.Repr() }
func (o *PList) Type() P        { return B_list }
func (o *PList) Repr() string {
	buf := bytes.NewBufferString("[")
	n := len(o.PP)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(o.PP[i].Repr())
	}
	buf.WriteString("]")
	return buf.String()
}
func (o *PList) Iter() Nexter {
	z := &PListIter{PP: o.PP}
	z.Self = z
	return z
}
func (o *PList) List() []P {
	return o.PP
}

type meth_PList_Append struct {
	PBase
	pl *PList
}

func (o *PList) GET_Append() P {
	z := &meth_PList_Append{pl: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PList_Append) Call1(item P) P {
	o.pl.PP = append(o.pl.PP, item)
	return o.pl
}

func (o *PListIter) Iter() Nexter {
	return o
}

type Nexter interface {
	Next() P
}

func (o *PListIter) Next() P {
	if o.I < len(o.PP) {
		z := o.PP[o.I]
		o.I++
		return z
	}
	panic(G_StopIterationSingleton)
}

func (o *PDict) NotContains(a P) bool { return !o.Contains(a) }
func (o *PDict) Contains(a P) bool {
	for x, _ := range o.PPP {
		if a.EQ(MkStr(x)) {
			return true
		}
	}
	return false
}
func (o *PDict) Len() int { return len(o.PPP) }
func (o *PDict) SetItem(a P, x P) {
	o.PPP[a.String()] = x
}
func (o *PDict) GetItem(a P) P {
	z, ok := o.PPP[a.String()]
	if !ok {
		return None
	}
	return z
}
func (o *PDict) String() string { return o.Repr() }
func (o *PDict) Type() P        { return B_dict }
func (o *PDict) Repr() string {
	keys := make([]string, 0, len(o.PPP))
	for k, _ := range o.PPP {
		keys = append(keys, k)
	}
	buf := bytes.NewBufferString("{")
	n := len(keys)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(F("%q: %s", keys[i], o.PPP[keys[i]].Repr()))
	}
	buf.WriteString("}")
	return buf.String()
}
func (o *PDict) Iter() Nexter {
	var keys []P
	for k, _ := range o.PPP {
		keys = append(keys, MkStr(k))
	}
	z := &PListIter{PP: keys}
	z.Self = z
	return z
}
func (o *PDict) List() []P {
	var keys []P
	for k, _ := range o.PPP {
		keys = append(keys, MkStr(k))
	}
	return keys
}

type PtrC_object_er interface {
	PtrC_object() *C_object
}

func (o *C_object) Bool() bool { // In python, all objects are true.
	return true
}

func (o *C_object) EQ(a P) bool {
	switch a2 := a.(type) {
	case PtrC_object_er:
		a3 := a2.PtrC_object()
		if o == a3 {
			return true
		}
	}
	return false
}

func NewList() *PList {
	z := &PList{PP: make([]P, 0)}
	z.Self = z
	return z
}

func CopySlice(pp []P) []P {
	zz := make([]P, len(pp))
	copy(zz, pp)
	return zz
}
func CopyList(aa *PList) *PList {
	z := &PList{PP: CopySlice(aa.PP)}
	z.Self = z
	return z
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

type PStopIteration struct{ PBase }

func F_StopIteration() P { return new(PStopIteration) }

// TODO: convert these.
var G_StopIteration = &PFunc0{Fn: F_StopIteration}
var G_StopIterationSingleton = F_StopIteration()

func F_len(a P) P   { return MkInt(int64(a.Len())) }
func F_repr(a P) P  { return MkStr(a.Repr()) }
func F_str(a P) P   { return MkStr(a.String()) }
func F_int(a P) P   { return MkInt(a.Int()) }
func F_float(a P) P { return MkFloat(a.Float()) }
func F_list(a P) P  { return MkList(a.List()) }
func F_tuple(a P) P { return MkTuple(a.List()) }
func F_dict(a P) P  { return MkDictFromPairs(a.List()) }
func F_bool(a P) P  { return MkBool(a.Bool()) }
func F_type(a P) P  { return a.Type() }

func F_range(a P) P {
	n := a.Int()
	v := make([]P, n)
	for i := int64(0); i < n; i++ {
		v[i] = MkInt(i)
	}
	return MkList(v)
}

// Types for sorting.
type StringyPs []P

func (o StringyPs) Len() int { return len(o) }
func (o StringyPs) Less(i, j int) bool {
	return (o[i].(*PStr).S < o[j].(*PStr).S)
}
func (o StringyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type IntyPs []P

func (o IntyPs) Len() int { return len(o) }
func (o IntyPs) Less(i, j int) bool {
	return (o[i].(*PInt).N < o[j].(*PInt).N)
}
func (o IntyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type FloatyPs []P

func (o FloatyPs) Len() int { return len(o) }
func (o FloatyPs) Less(i, j int) bool {
	return (o[i].(*PFloat).F < o[j].(*PFloat).F)
}
func (o FloatyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func F_sorted(a P) P {
	ps := CopySlice(a.List())
	if len(ps) == 0 {
		return MkList([]P{})
	}
	switch ps[0].(type) {
	case *PStr:
		sort.Sort(StringyPs(ps))
	case *PInt:
		sort.Sort(IntyPs(ps))
	case *PFloat:
		sort.Sort(FloatyPs(ps))
	default:
		panic(Badf("sorted: cannot sort list beginning with type %t", ps[0]))
	}
	return MkList(ps)
}

// Builting functions.
var B_len *PFunc1
var B_repr *PFunc1
var B_str *PFunc1
var B_int *PFunc1
var B_float *PFunc1
var B_range *PFunc1
var B_sorted *PFunc1
var B_list *PFunc1
var B_dict *PFunc1
var B_tuple *PFunc1
var B_bool *PFunc1
var B_type *PFunc1

func init() {
	B_len = &PFunc1{Fn: F_len}
	B_repr = &PFunc1{Fn: F_repr}
	B_str = &PFunc1{Fn: F_str}
	B_int = &PFunc1{Fn: F_int}
	B_float = &PFunc1{Fn: F_float}
	B_range = &PFunc1{Fn: F_range}
	B_sorted = &PFunc1{Fn: F_sorted}
	B_list = &PFunc1{Fn: F_list}
	B_dict = &PFunc1{Fn: F_dict}
	B_tuple = &PFunc1{Fn: F_tuple}
	B_bool = &PFunc1{Fn: F_bool}
	B_type = &PFunc1{Fn: F_type}

	B_len.Self = B_len
	B_repr.Self = B_repr
	B_str.Self = B_str
	B_int.Self = B_int
	B_float.Self = B_float
	B_range.Self = B_range
	B_sorted.Self = B_sorted
	B_list.Self = B_list
	B_dict.Self = B_dict
	B_tuple.Self = B_tuple
	B_bool.Self = B_bool
	B_type.Self = B_type
}

type PModule struct {
	PBase
}

func (o *PModule) Init_PModule() {
}

func init() {
	// TODO: convert these to members.
	var G_StopIteration = &PFunc0{Fn: F_StopIteration}
	G_StopIteration.Self = G_StopIteration
}

type PFunc0 struct {
	PBase
	Fn func() P
}

func (p *PFunc0) Call0() P {
	return p.Fn()
}

type PFunc1 struct {
	PBase
	Fn func(a P) P
}

func (o *PFunc1) EQ(a P) bool { return (o == a.(*PFunc1)) }
func (p *PFunc1) Call1(a1 P) P {
	return p.Fn(a1)
}

func (g *PGo) Call(aa ...P) P {
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
		args[i] = AdaptForCall(a, t.In(i))
	}
	outs := f.Call(args)
	// TODO: strip off error.
	switch numOut {
	case 0:
		return None
	case 1:
		return AdaptForReturn(outs[0])
	}
	panic(Bad("Multi-arg returns no imp yet"))
}

func AdaptForCall(v P, t R.Type) R.Value {
	switch t.Kind() {
	case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
		pgo, ok := v.(*PGo)
		if ok && pgo.V.IsNil() {
			return R.Zero(t)
		}
	}
	switch t.Kind() {
	case R.Int:
		return R.ValueOf(int(v.Int()))
	case R.Int64:
		return R.ValueOf(v.Int())
	case R.String:
		return R.ValueOf(v.String())
	}
	panic(Bad("Cannot AdaptForCall: %s TO %s", v, t))
}

func AdaptForReturn(v R.Value) P {
	switch v.Kind() {
	case R.String:
		return MkStr(v.String())
	case R.Int:
		return MkInt(v.Int())
	case R.Int64:
		return MkInt(v.Int())
	case R.Bool:
		if v.Bool() {
			return True
		}
		return False
	}
	panic(Bad("Cannot AdaptForReturn: %s: %#v", v.Kind(), v.Interface()))
}

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

// See ../../yak-labs/chirp-lang/reflect.go
var Roots map[string]Per = make(map[string]Per)

var errorInterfaceType R.Type = R.TypeOf(errors.New).Out(0)

type Per interface {
	P() P
}
type FuncRoot struct{ Func R.Value }
type VarRoot struct{ Var R.Value }
type TypeRoot struct{ Type R.Type }
type ConstRoot struct{ Const interface{} }

func (r FuncRoot) P() P  { return MkGo(r.Func.Interface()) }
func (r VarRoot) P() P   { return MkGo(r.Var.Interface()) }
func (r TypeRoot) P() P  { return MkGo(r.Type) }
func (r ConstRoot) P() P { return MkGo(R.ValueOf(r.Const).Interface()) }

type PGoModule struct {
	PBase
	SimpleName string // TODO make this more general
	RootPrefix string // Append what you're looking for.
}

func (o PGoModule) FieldForCall(field string) P {
	per, ok := Roots[o.RootPrefix+field]
	if !ok {
		panic(Bad("No field %q on PGoModule %q", field, o.SimpleName))
	}
	return per.P()
}

func GoImport(im string) *PGoModule {
	z := &PGoModule{
		SimpleName: im,
		RootPrefix: "/" + im + "/",
	}
	z.Self = z
	return z
}

func init() {
	var tmp P = new(PBase)
	// Demonstrate these things implement P.
	tmp = new(PInt)
	tmp = new(PFloat)
	tmp = new(PList)
	tmp = new(PDict)
	tmp = new(PModule)
	tmp = new(PGoModule)
	_ = tmp
}
