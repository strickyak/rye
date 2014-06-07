package runt

import (
	"bytes"
	"fmt"
	"errors"
	"go/ast"
	"os"
	R "reflect"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"

	. "github.com/strickyak/yak"
)

var _ = debug.PrintStack

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
	Invoke(field string, aa ...P) P
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
	Contents() interface{}
	Bytes() []byte
}

// C_object is the root of inherited classes.
type C_object struct {
	PBase
}

func (o *C_object) PtrC_object() *C_object {
	return o
}

type EitherPOrError struct {
	Right P
	Left  error
}

type void struct{}

// C_generator is the channel begin a yielding producer and a consuming for loop.
type C_generator struct {
	C_object
	Ready  chan *void
	Result chan EitherPOrError
}

func NewGenerator() *C_generator {
	z := &C_generator{
		Ready:  make(chan *void, 1),
		Result: make(chan EitherPOrError, 1),
	}
	z.SetSelf(z)
	return z
}

func (o *C_generator) PtrC_generator() *C_generator {
	return o
}

func (o *C_generator) Iter() Nexter { return o }

// Next is called by the consumer.
// Next waits for next result from the generator.
// It returns either a result of type P and true,
// or if there are no more, it returns nil and false.
// If the generator goroutine died on an exception,
// that exception gets wrapped in a new error and rethrown here.
func (o *C_generator) Next() (P, bool) {
	o.Ready <- nil
	// That wakes up the generator goroutine.
	// Now we block, waiting on next result.
	either, ok := <-o.Result
	if !ok {
		return nil, false
	}
	if either.Left != nil {
		close(o.Ready)
		panic(errors.New(F("Generator threw exception: %q", either.Left)))
	}
	return either.Right, true
}

// Enough is called by the consumer, to tell the producer to stop because we've got enough.
func (o *C_generator) Enough() {
	close(o.Ready)
}

// Yield is called by the producer, to yield a value to the consumer.
func (o *C_generator) Yield(item P) {
	o.Result <- EitherPOrError{Right: item, Left: nil}
}

// Yield is called by the producer when it catches an exception, to yield it to the producer (as an Either Left).
func (o *C_generator) YieldError(err error) {
	o.Result <- EitherPOrError{Right: nil, Left: err}
}

// Finish is called by the producer when it is finished.
func (o *C_generator) Finish() {
	close(o.Result)
}

// Wait is called by the producer, once at the start, and once after each yield.
// Wait returns false if the consumer said Enough.
// TODO:  Don't wait, to achieve concurrency.  Let the user decide the Result channel buffer size.
func (o *C_generator) Wait() bool {
	_, ok := <-o.Ready
	return ok
}

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
func (o *PBase) Call(aa ...P) P              { panic(Bad("Receiver cannot Call", o.Self, aa)) }
func (o *PBase) Invoke(field string, aa ...P) P {
	panic(Bad("Receiver cannot invoke", o.Self, field, aa))
}
func (o *PBase) Len() int      { panic(Bad("Receiver cannot Len: ", o.Self)) }
func (o *PBase) GetItem(a P) P { panic(Bad("Receiver cannot GetItem", o.Self, o, a)) }
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

func (o *PBase) Bool() bool { return true } // Most things are true.
func (o *PBase) Neg() P     { panic(Bad("Receiver cannot Neg", o.Self)) }
func (o *PBase) Pos() P     { panic(Bad("Receiver cannot Pos", o.Self)) }
func (o *PBase) Abs() P     { panic(Bad("Receiver cannot Abs", o.Self)) }
func (o *PBase) Inv() P     { panic(Bad("Receiver cannot Inv", o.Self)) }

func (o *PBase) Int() int64            { panic(Bad("Receiver cannot Int", o.Self)) }
func (o *PBase) Float() float64        { panic(Bad("Receiver cannot Float", o.Self)) }
func (o *PBase) Complex() complex128   { panic(Bad("Receiver cannot Complex", o.Self)) }
func (o *PBase) Contents() interface{} { return o.Self }

func (o *PBase) Type() P { return MkStr(F("%t", o.Self)) }
func (o *PBase) Bytes() []byte { panic(Bad("Receiver cannot Bytes", o.Self)) }
func (o *PBase) String() string {
	if o.Self == nil {
		panic("PBase:  Why is o.Self NIL?")
	}
	return o.Self.Show()
}
func (o *PBase) Repr() string { return o.String() }
func (o *PBase) Show() string {
	if o.Self == nil {
		panic("OHNO: o.Self == nil")
	}
	return ShowP(o.Self, 3)
}

func ShowP(a P, depth int) string {
	r := R.ValueOf(a)
	if !r.IsValid() {
		panic("INVALID")
		return "$INVALID$ "
	}

	switch r.Kind() {
	case R.Interface:
		if r.IsNil() {
			return "$NIL_INTERFACE$ "
		}
		r = r.Elem()
	}

	// Deref pointers.
	switch r.Kind() {
	case R.Ptr:
		if r.IsNil() {
			return "$NIL_PTR$ "
		}
		r = r.Elem()
	}

	buf := bytes.NewBuffer(nil)
	t := r.Type()
	switch r.Kind() {
	case R.Struct:
		tn := t.Name()
		if tn == "" {
			tn = "?"
		}
		buf.WriteString(F("{%s ", tn))
		if depth > 0 {
			for i := 0; i < t.NumField(); i++ {
				k := t.Field(i).Name
				if k == "PBase" {
					continue
				}
				if k == "C_object" || k == "object" {
					continue
				}
				if !ast.IsExported(k) {
					buf.WriteString("$PRIVATE$ ")
					continue
				}
				v := r.Field(i)

				switch x := v.Interface().(type) {
				case R.Value:
					v = x
				}

				if !v.IsValid() {
					buf.WriteString(F("%s=Invalid ", k))
					continue
				}
				switch x := v.Interface().(type) {
				case *PInt:
					buf.WriteString(F("%s=%d ", k, x.N))
				case *PFloat:
					buf.WriteString(F("%s=%f ", k, x.F))
				case *PStr:
					buf.WriteString(F("%s=%q ", k, x.S))
				case *PNone:
					buf.WriteString(F("%s=None ", k))
				case *PGo:
					buf.WriteString(F("%s=%v ", k, x.V.Interface()))
				case P:
					buf.WriteString(F("%s=%s ", k, ShowP(v.Interface().(P), depth-1)))
				case int:
					buf.WriteString(F("%s=%d ", k, x))
				case int64:
					buf.WriteString(F("%s=%d ", k, x))
				case float64:
					buf.WriteString(F("%s=%f ", k, x))
				case string:
					buf.WriteString(F("%s=%q ", k, x))
				case fmt.Stringer:
					buf.WriteString(F("%s~%T~%v ", k, x, x))
				default:
					for v.Kind() == R.Interface {
						v = v.Elem()
					}
					for v.Kind() == R.Ptr {
						v = v.Elem()
					}
					if v.Kind() == R.Struct {
						ptr := v.Addr().Interface()
						if inner, ok := ptr.(P) ; ok {
							buf.WriteString(F("%s", ShowP(inner, depth)))
						} else {
							buf.WriteString(F("%v", v.Interface()))
						}
					} else {
						buf.WriteString(F("%s:%s ", k, v.Type().Name()))
					}
				}
			}
		}
		buf.WriteString("} ")
	default:
		buf.WriteString(F("Kind:%s", r.Kind()))
	}
	return buf.String()
}

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

type PByt struct {
	PBase
	YY []byte
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
func MkValue(a R.Value) *PGo  { z := &PGo{V: a}; z.Self = z; return z }

func Mkint(n int) *PInt         { z := &PInt{N: int64(n)}; z.Self = z; return z }
func MkInt(n int64) *PInt       { z := &PInt{N: n}; z.Self = z; return z }
func MkFloat(f float64) *PFloat { z := &PFloat{F: f}; z.Self = z; return z }
func MkStr(s string) *PStr      { z := &PStr{S: s}; z.Self = z; return z }
func MkByt(yy []byte) *PByt     { z := &PByt{YY: yy}; z.Self = z; return z }

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

func (o *PNone) Bool() bool            { return false }
func (o *PNone) String() string        { return "None" }
func (o *PNone) Repr() string          { return "None" }
func (o *PNone) Contents() interface{} { return nil }

func (o *PBool) Contents() interface{} { return o.B }
func (o *PBool) Bool() bool            { return o.B }
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

func (o *PInt) Add(a P) P { return MkInt(o.N + a.Int()) }
func (o *PInt) Sub(a P) P { return MkInt(o.N - a.Int()) }
func (o *PInt) Mul(a P) P {
	switch x := a.(type) {
	case *PInt:
		return MkInt(o.N * x.N)
	case *PStr:
		return MkStr(strings.Repeat(x.S, int(o.N)))
	case *PList:
		var z []P
		for i := 0; i < int(o.N); i++ {
			z = append(z, x.PP...)
		}
		return MkList(z)
	case *PTuple:
		var z []P
		for i := 0; i < int(o.N); i++ {
			z = append(z, x.PP...)
		}
		return MkTuple(z)
	}
	panic("Cannot multply int times whatever")
}
func (o *PInt) Div(a P) P             { return MkInt(o.N / a.Int()) }
func (o *PInt) Mod(a P) P             { return MkInt(o.N % a.Int()) }
func (o *PInt) And(a P) P             { return MkInt(o.N & a.Int()) }
func (o *PInt) Or(a P) P              { return MkInt(o.N | a.Int()) }
func (o *PInt) Xor(a P) P             { return MkInt(o.N ^ a.Int()) }
func (o *PInt) LShift(a P) P          { return MkInt(o.N << uint64(a.Int())) }
func (o *PInt) RShift(a P) P          { return MkInt(o.N >> uint64(a.Int())) }
func (o *PInt) EQ(a P) bool           { return (o.N == a.Int()) }
func (o *PInt) NE(a P) bool           { return (o.N != a.Int()) }
func (o *PInt) LT(a P) bool           { return (o.N < a.Int()) }
func (o *PInt) LE(a P) bool           { return (o.N <= a.Int()) }
func (o *PInt) GT(a P) bool           { return (o.N > a.Int()) }
func (o *PInt) GE(a P) bool           { return (o.N >= a.Int()) }
func (o *PInt) Int() int64            { return o.N }
func (o *PInt) Float() float64        { return float64(o.N) }
func (o *PInt) String() string        { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string          { return o.String() }
func (o *PInt) Bool() bool            { return o.N != 0 }
func (o *PInt) Type() P               { return B_int }
func (o *PInt) Contents() interface{} { return o.N }

func (o *PFloat) Add(a P) P             { return MkFloat(o.F + a.Float()) }
func (o *PFloat) Sub(a P) P             { return MkFloat(o.F - a.Float()) }
func (o *PFloat) Mul(a P) P             { return MkFloat(o.F * a.Float()) }
func (o *PFloat) Div(a P) P             { return MkFloat(o.F / a.Float()) }
func (o *PFloat) EQ(a P) bool           { return (o.F == a.Float()) }
func (o *PFloat) NE(a P) bool           { return (o.F != a.Float()) }
func (o *PFloat) LT(a P) bool           { return (o.F < a.Float()) }
func (o *PFloat) LE(a P) bool           { return (o.F <= a.Float()) }
func (o *PFloat) GT(a P) bool           { return (o.F > a.Float()) }
func (o *PFloat) GE(a P) bool           { return (o.F >= a.Float()) }
func (o *PFloat) Int() int64            { return int64(o.F) }
func (o *PFloat) Float() float64        { return o.F }
func (o *PFloat) String() string        { return strconv.FormatFloat(o.F, 'g', -1, 64) }
func (o *PFloat) Repr() string          { return o.String() }
func (o *PFloat) Bool() bool            { return o.F != 0 }
func (o *PFloat) Type() P               { return B_float }
func (o *PFloat) Contents() interface{} { return o.F }

func (o *PStr) Iter() Nexter {
	var pp []P
	for _, r := range o.S {
		pp = append(pp, MkStr(string(r)))
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}
func (o *PStr) Contents() interface{} { return o.S }
func (o *PStr) Bool() bool            { return len(o.S) != 0 }
func (o *PStr) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.S))
	}
	return MkStr(o.S[i : i+1])
}

func (o *PStr) GetItemSlice(x, y, z P) P {
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
		z := make([]interface{}, len(t.PP))
		for i, e := range t.PP {
			z[i] = e.Contents()
		}
		return MkStr(F(o.S, z...))
	case *PInt:
		return MkStr(F(o.S, t.N))
	case *PFloat:
		return MkStr(F(o.S, t.F))
	case *PStr:
		return MkStr(F(o.S, t.S))
	default:
		return MkStr(F(o.S, a.Contents()))
	}
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
func (o *PStr) Bytes() []byte { return []byte(o.S) }
func (o *PStr) Len() int       { return len(o.S) }
func (o *PStr) Repr() string   { return F("%q", o.S) }
func (o *PStr) Type() P        { return B_str }

func (o *PByt) Iter() Nexter {
	var pp []P
	for _, r := range o.YY {
		pp = append(pp, Mkint(int(r)))
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}
func (o *PByt) Contents() interface{} { return o.YY }
func (o *PByt) Bool() bool            { return len(o.YY) != 0 }
func (o *PByt) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.YY))
	}
	return Mkint(int(o.YY[i]))
}

func (o *PByt) GetItemSlice(x, y, z P) P {
	var i, j int64
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.YY))
		}
	}
	if y == None {
		j = int64(len(o.YY))
	} else {
		j = y.Int()
		if j < 0 {
			j += int64(len(o.YY))
		}
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkByt(o.YY[i:j])
	return r
}

func (o *PByt) Mul(a P) P {
	switch t := a.(type) {
	case *PInt:
		return MkByt(bytes.Repeat(o.YY, int(t.Int())))
	}
	panic(Badf("Cannot multiply: byt * %t", a))
}
func (o *PByt) NotContains(a P) bool { return !o.Contains(a) }
func (o *PByt) Contains(a P) bool {
	switch t := a.(type) {
	case *PByt:
		return bytes.Contains(o.YY, t.YY)
	}
	panic(Bad("Byt cannot Contain non-byt:", a))
}
func (o *PByt) Add(a P) P      {
	aa := a.Bytes()
	var zz []byte
	zz = append(zz, o.YY...)
	zz = append(zz, aa...)
	return MkByt(zz)
}

func (o *PByt) EQ(a P) bool    { return (string(o.YY) == a.String()) }
func (o *PByt) NE(a P) bool    { return (string(o.YY) != a.String()) }
func (o *PByt) LT(a P) bool    { return (string(o.YY) < a.String()) }
func (o *PByt) LE(a P) bool    { return (string(o.YY) <= a.String()) }
func (o *PByt) GT(a P) bool    { return (string(o.YY) > a.String()) }
func (o *PByt) GE(a P) bool    { return (string(o.YY) >= a.String()) }

func (o *PByt) String() string { return string(o.YY) }
func (o *PByt) Bytes() []byte { return o.YY }
func (o *PByt) Len() int       { return len(o.YY) }
func (o *PByt) Repr() string   { return F("byt(%q)", string(o.YY)) }
func (o *PByt) Type() P        { return B_byt }

func (o *PTuple) Contents() interface{} { return o.PP }
func (o *PTuple) Bool() bool            { return len(o.PP) != 0 }
func (o *PTuple) NotContains(a P) bool  { return !o.Contains(a) }
func (o *PTuple) Contains(a P) bool {
	for _, x := range o.PP {
		if a.EQ(x) {
			return true
		}
	}
	return false
}
func (o *PTuple) Len() int { return len(o.PP) }
func (o *PTuple) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.PP))
	}
	return o.PP[i]
}
func (o *PTuple) String() string { return o.Repr() }
func (o *PTuple) Type() P        { return B_tuple }
func (o *PTuple) Repr() string {
	n := len(o.PP)
	if n == 0 {
		return "()"
	}
	buf := bytes.NewBufferString("(")
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(o.PP[i].Repr())
	}
	if n == 1 {
		buf.WriteString(",") // Special just for singleton tuples.
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

func (o *PList) Contents() interface{} { return o.PP }
func (o *PList) Bool() bool            { return len(o.PP) != 0 }
func (o *PList) NotContains(a P) bool  { return !o.Contains(a) }
func (o *PList) Contains(a P) bool {
	for _, x := range o.PP {
		if a.EQ(x) {
			return true
		}
	}
	return false
}
func (o *PList) Len() int { return len(o.PP) }
func (o *PList) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.PP))
	}
	return o.PP[i]
}
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

type meth_PList_append struct {
	PBase
	list *PList
}

func (o *PList) GET_append() P {
	z := &meth_PList_append{list: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PList_append) Call1(item P) P {
	o.list.PP = append(o.list.PP, item)
	return o.list
}

func (o *PListIter) Iter() Nexter {
	return o
}

type Nexter interface {
	Next() (P, bool)
}

// A Nexter may or may not implement Enough.
// It is not built in to Nexter, because if it
// does not exist, we can avoid a defer,
// which is still expensive, as of go1.2.
// When a goroutine is generating, we do want to
// defer Enough() so the goroutine won't be leaked.
type Enougher interface {
	Enough()
}

func (o *PListIter) Next() (P, bool) {
	if o.I < len(o.PP) {
		z := o.PP[o.I]
		o.I++
		return z, true
	}
	return nil, false
}

func (o *PDict) Contents() interface{} { return o.PPP }
func (o *PDict) Bool() bool            { return len(o.PPP) != 0 }
func (o *PDict) NotContains(a P) bool  { return !o.Contains(a) }
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
	key := a.String()
	z, ok := o.PPP[key]
	if !ok {
		panic(F("PDict: KeyError: %q", key))
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
func (o *PDict) Enough() {}
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

type meth_PDict_get struct {
	PBase
	dict *PDict
}

func (o *PDict) GET_get() P {
	z := &meth_PDict_get{dict: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PDict_get) Call1(item P) P {
	key := item.String()
	if z, ok := o.dict.PPP[key]; ok {
		return z
	}
	return None
}

func (o *meth_PDict_get) Call2(item P, dflt P) P {
	key := item.String()
	if z, ok := o.dict.PPP[key]; ok {
		return z
	}
	return dflt
}

type PtrC_object_er interface {
	PtrC_object() *C_object
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

func Entuple(args ...P) *PTuple {
	zz := make([]P, 0)
	for _, a := range args {
		zz = append(zz, a)
	}
	return MkTuple(zz)
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

func B_1_len(a P) P   { return MkInt(int64(a.Len())) }
func B_1_repr(a P) P  { return MkStr(a.Repr()) }
func B_1_str(a P) P   { return MkStr(a.String()) }
func B_1_int(a P) P   { return MkInt(a.Int()) }
func B_1_float(a P) P { return MkFloat(a.Float()) }
func B_1_list(a P) P  { return MkList(a.List()) }
func B_1_tuple(a P) P { return MkTuple(a.List()) }
func B_1_dict(a P) P  { return MkDictFromPairs(a.List()) }
func B_1_bool(a P) P  { return MkBool(a.Bool()) }
func B_1_type(a P) P  { return a.Type() }
func B_1_byt(a P) P  {
	switch x := a.(type) {
	case *PStr:
		bb := make([]byte, len(x.S))
		copy(bb, x.S)
		return MkGo(bb)
	}
	panic(F("Cannot make bytes from a %T", a))
}

func B_1_range(a P) P {
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

func B_1_sorted(a P) P {
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
var B_byt *PFunc1

func init() {
	B_len = &PFunc1{Fn: B_1_len}
	B_repr = &PFunc1{Fn: B_1_repr}
	B_str = &PFunc1{Fn: B_1_str}
	B_int = &PFunc1{Fn: B_1_int}
	B_float = &PFunc1{Fn: B_1_float}
	B_range = &PFunc1{Fn: B_1_range}
	B_sorted = &PFunc1{Fn: B_1_sorted}
	B_list = &PFunc1{Fn: B_1_list}
	B_dict = &PFunc1{Fn: B_1_dict}
	B_tuple = &PFunc1{Fn: B_1_tuple}
	B_bool = &PFunc1{Fn: B_1_bool}
	B_type = &PFunc1{Fn: B_1_type}
	B_byt = &PFunc1{Fn: B_1_byt}

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
	B_byt.Self = B_byt
}

type PModule struct {
	PBase
}

func (o *PModule) Init_PModule() {
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

func (g *PGo) String() string {
	g0 := MaybeDeref(g.V)
	i0 := g0.Interface()
	switch x := i0.(type) {
	case fmt.Stringer:
		return x.String()
	case []byte:
		return string(x)
	}

	switch g0.Kind() {
	case R.Array:
		switch g0.Type().Elem().Kind() {
		case R.Uint8:
			bb := make([]byte, g0.Len())
			R.Copy(R.ValueOf(bb), g0)
			return string(bb)
		}
	}

	if g0.CanAddr() {
		g1 := g0.Addr()
		i1 := g1.Interface()
		switch x := i1.(type) {
		case fmt.Stringer:
			return x.String()
		}
	}
	// Fallback on ShowP
	return "fallback:" + ShowP(g, 3)
}
func (g *PGo) Invoke(field string, aa ...P) P {
	g0 := MaybeDeref(g.V)

	meth, ok := g0.Type().MethodByName(field)
	if !ok {
		if g0.CanAddr() {
			g0 = g0.Addr()
			meth, ok = g0.Type().MethodByName(field)
		}
		if !ok {
			panic(F("Method on type %q does not exist: %s", g0.Type(), field))
		}
	}

	f := meth.Func
	return FinishInvokeOrCall(f, g0, aa)
}

func (g *PGo) Call(aa ...P) P {
	f := MaybeDeref(g.V)
	if f.Kind() != R.Func {
		z, ok := FunCallN(f, aa)
		if !ok {
			Bad("cannot Call when Value not a func and FunCallN fails", f)
		}
		return z
	}
	var zeroValue R.Value
	return FinishInvokeOrCall(f, zeroValue, aa)
}
func (g *PGo) Iter() Nexter {
	a := MaybeDeref(g.V)
	var pp []P

	switch a.Kind() {
	case R.Array, R.Slice:
		n := a.Len()
		for i := 0; i < n; i++ {
			pp = append(pp, AdaptForReturn(a.Index(i)))
		}
	default:
		Bad("*PGo cannot Iter() on kind %s", a.Kind())
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}

var errorType = R.TypeOf(new(error)).Elem()

func FinishInvokeOrCall(f R.Value, rcvr R.Value, aa []P) P {
	hasRcvr := rcvr.IsValid()
	lenRcvr := 0
	if hasRcvr {
		lenRcvr = 1
	}
	lenArgs := len(aa)
	lenIns := lenRcvr + lenArgs
	t := f.Type()
	numIn := t.NumIn()

	args := make([]R.Value, lenIns)
	if t.IsVariadic() {
		if lenIns < numIn-1 {
			Bad("call got %d args, want %d or more args", lenIns, numIn-1)
		}
		args[0] = rcvr
		for i, a := range aa {
			var desiredType R.Type
			if i >= numIn-1 {
				desiredType = t.In(numIn - 1).Elem()
			} else {
				desiredType = t.In(i)
			}
			args[i+lenRcvr] = AdaptForCall(a, desiredType)
		}
	} else {
		if lenIns != numIn {
			Bad("call got %d args, want %d args", lenIns, numIn)
		}
		args[0] = rcvr
		for i, a := range aa {
			args[i+lenRcvr] = AdaptForCall(a, t.In(i))
		}
	}

	outs := f.Call(args)

	numOut := t.NumOut()
	if numOut > 0 && t.Out(numOut-1) == errorType {
		// Check for error.
		if !outs[numOut-1].IsNil() {
			// Panic the error.
			panic(outs[numOut-1].Interface())
		}
		// Forget the nil error.
		numOut--
	}

	switch numOut {
	case 0:
		return None
	case 1:
		return AdaptForReturn(outs[0])
	default:
		slice := make([]P, numOut)
		for i := 0; i < numOut; i++ {
			slice[i] = AdaptForReturn(outs[i])
		}
		return MkTuple(slice)
	}
}

var typeInterfaceEmpty = R.TypeOf(new(interface{})).Elem()

func AdaptForCall(v P, want R.Type) R.Value {
	// None & nil.
	switch want.Kind() {
	case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
		// Convert Python None to nil go thing.
		if v == None {
			return R.Zero(want)
		}
		// Convert Go Nil (in a *PGo) to nil go thing.
		pgo, ok := v.(*PGo)
		if ok && pgo.V.IsNil() {
			return R.Zero(want)
		}
	}

	switch want.Kind() {
	case R.Uint8:
		return R.ValueOf(uint8(v.Int()))
	case R.Uint16:
		return R.ValueOf(uint16(v.Int()))
	case R.Uint32:
		return R.ValueOf(uint32(v.Int()))
	case R.Uint64:
		return R.ValueOf(uint64(v.Int()))
	case R.Int:
		return R.ValueOf(int(v.Int()))
	case R.Int8:
		return R.ValueOf(int8(v.Int()))
	case R.Int16:
		return R.ValueOf(int16(v.Int()))
	case R.Int32:
		return R.ValueOf(int32(v.Int()))
	case R.Int64:
		return R.ValueOf(v.Int())
	case R.String:
		return R.ValueOf(v.String())
	case R.Func:
		return MakeFunction(v, want) // This is hard.
	case R.Slice:
		switch want.Elem().Kind() {
		case R.Uint8:
			var bb []byte 
			copy(bb, v.String())
			return R.ValueOf(bb)
		}
	}

	switch vx := v.(type) {
	case *PGo:
		return vx.V.Convert(want)
	}

	if want == typeInterfaceEmpty {
		switch x := v.(type) {
		case *PInt:
			return R.ValueOf(x.N)
		case *PStr:
			return R.ValueOf(x.S)
		case *PBool:
			return R.ValueOf(x.B)
		}
	}
	panic(F("Cannot AdaptForCall: %s [%s] TO %s [%s]", v, R.TypeOf(v), want, want.Kind()))
}

func MakeFunction(v P, t R.Type) R.Value {
	nin := t.NumIn()
	if nin > 3 {
		panic(F("Not implemented: MakeFunction for %d args", nin))
	}

	return R.MakeFunc(t, func(aa []R.Value) (zz []R.Value) {
		var r P
		switch nin {
		case 0:
			r = v.(i_0).Call0()
		case 1:
			r = v.(i_1).Call1(AdaptForReturn(aa[0]))
		case 2:
			r = v.(i_2).Call2(AdaptForReturn(aa[0]), AdaptForReturn(aa[1]))
		case 3:
			r = v.(i_3).Call3(AdaptForReturn(aa[0]), AdaptForReturn(aa[1]), AdaptForReturn(aa[2]))
		default:
			panic(F("Not implemented: MakeFunction for %d args", nin))
		}

		// TODO: final error case.
		nout := t.NumOut()
		switch nout {
		case 0: // pass
		case 1:
			zz = append(zz, AdaptForCall(r, t.Out(0)))
		default:
			zz = make([]R.Value, nout)
			for i := 0; i < nout; i++ {
				zz[i] = AdaptForCall(r.GetItem(Mkint(i)), t.Out(i))
			}
		}
		return
	})
}

func AdaptForReturn(v R.Value) P {
	switch v.Kind() {
	case R.String:
		return MkStr(v.String())
	case R.Int:
		return MkInt(v.Int())
	case R.Int8:
		return MkInt(v.Int())
	case R.Int16:
		return MkInt(v.Int())
	case R.Int32:
		return MkInt(v.Int())
	case R.Int64:
		return MkInt(v.Int())
	case R.Uint8:
		return MkInt(int64(v.Uint()))
	case R.Uint16:
		return MkInt(int64(v.Uint()))
	case R.Uint32:
		return MkInt(int64(v.Uint()))
	case R.Uint64:
		return MkInt(int64(v.Uint()))
	case R.Bool:
		if v.Bool() {
			return True
		}
		return False
	case R.Slice:
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			return MkByt(v.Interface().([]byte))
		}
	case R.Array:
		n := v.Len()
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			bb := make([]byte, n)
			R.Copy(R.ValueOf(bb), v)
			return MkByt(bb)
			// return MkByt(v.Slice(0, n).Interface().([]byte))
		}
	}
	return MkValue(v)
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

var Imports = make(map[string]*PImport)
var ImportsEvalled = make(map[string]bool)

func EvalRyeModuleOnce(path string) bool {
	if ImportsEvalled[path] {
		return false
	}
	ImportsEvalled[path] = true
	return true
}

type PImport struct {
	PBase
	Go         bool
	Path       string // TODO make this more general
	RootPrefix string // Append what you're looking for.
	Reflect    R.Value
}

func GoImport(path string) *PImport {
	z, ok := Imports[path]
	if ok {
		if !z.Go {
			panic(F("GoImport: already imported %q but it is not a Go import", path))
		}
		return z
	}
	z = &PImport{
		Go:         true,
		Path:       path,
		RootPrefix: "/" + path + "/",
	}
	z.Self = z
	Imports[path] = z
	return z
}

func RyeImport(path string, mod interface{}) *PImport {
	z, ok := Imports[path]
	if ok {
		if z.Go {
			panic(F("RyeImport: already imported %q but it is a Go import", path))
		}
		return z
	}
	z = &PImport{
		Go:      false,
		Path:    path,
		Reflect: R.ValueOf(mod.(P).GetSelf()).Elem(),
	}
	z.Self = z
	Imports[path] = z
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
	tmp = new(PImport)
	tmp = new(C_object)
	_ = tmp
}

type i_0 interface {
	Call0() P
}
type i_1 interface {
	Call1(P) P
}
type i_2 interface {
	Call2(P, P) P
}
type i_3 interface {
	Call3(P, P, P) P
}
type i_4 interface {
	Call4(P, P, P, P) P
}

func FunCallN(f R.Value, aa []P) (P, bool) {
	switch len(aa) {
	case 0:
		if c, ok := f.Interface().(i_0); ok {
			return c.Call0(), true
		}
	case 1:
		if c, ok := f.Interface().(i_1); ok {
			return c.Call1(aa[0]), true
		}
	case 2:
		if c, ok := f.Interface().(i_2); ok {
			return c.Call2(aa[0], aa[1]), true
		}
	case 3:
		if c, ok := f.Interface().(i_3); ok {
			return c.Call3(aa[0], aa[1], aa[2]), true
		}
	case 4:
		if c, ok := f.Interface().(i_4); ok {
			return c.Call4(aa[0], aa[1], aa[2], aa[3]), true
		}
		// TODO: Fall back to reflection.
	}

	return None, false
}
