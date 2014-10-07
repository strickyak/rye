package rye

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"go/ast"
	"math"
	"os"
	R "reflect"
	"runtime/debug"
	"strconv"
	"strings"

	. "github.com/strickyak/yak"
)

var _ = debug.PrintStack

const SHOW_DEPTH = 6

var None = &PNone{}
var True = &PBool{B: true}
var False = &PBool{B: false}

var Globals map[string]P = make(map[string]P)

func init() {
	None.Self = None
	True.Self = True
	False.Self = False
}

var RyeEnv string
var Debug int
var DebugAdapt int
var DebugGo int

func init() {
	RyeEnv := os.Getenv("RYE")
	for _, ch := range RyeEnv {
		switch ch {
		case 'd':
			Debug++
		case 'a':
			DebugAdapt++
		case 'g':
			DebugGo++
		}
	}
}

// P is the interface for every Pythonic value.
type P interface {
	Pickle(w *bytes.Buffer)
	Show() string
	String() string
	Repr() string
	Type() P
	Is(a P) bool
	IsNot(a P) bool
	GetSelf() P
	SetSelf(a P)
	GetPBase() *PBase

	Field(field string) P
	// FieldGets(field string, x P) P
	FieldForCall(field string) P
	Call(aa ...P) P
	Invoke(field string, aa ...P) P
	Iter() Nexter
	List() []P
	Dict() map[string]P

	Len() int
	SetItem(i P, x P)
	DelItem(i P)
	DelItemSlice(i, j P)
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
	BitAnd(a P) P
	BitOr(a P) P
	BitXor(a P) P
	ShiftLeft(a P) P
	ShiftRight(a P) P
	UnsignedShiftRight(a P) P

	IAdd(a P) // +=
	ISub(a P) // -=
	IMul(a P) // *=

	Bool() bool // a.k.a. nonzero()
	UnaryMinus() P
	UnaryPlus() P
	UnaryInvert() P

	EQ(a P) bool
	NE(a P) bool
	LT(a P) bool
	LE(a P) bool
	GT(a P) bool
	GE(a P) bool
	Compare(a P) int // Non-rich comparison.

	Int() int64
	Float() float64
	Complex() complex128
	Contents() interface{}
	Bytes() []byte
}

func Pickle(p P) []byte {
	var b bytes.Buffer
	p.GetSelf().Pickle(&b)
	z := b.Bytes()
	//@ Say("PICKLE", len(z))
	return z
}

// C_object is the root of inherited classes.
type C_object struct {
	PBase
}

func (o *C_object) PtrC_object() *C_object {
	return o
}

func MkPromise(fn func() P) *C_promise {
	z := &C_promise{Ch: make(chan EitherPOrError, 1)}
	z.SetSelf(z)
	if DebugGo > 0 {
		println("#go# Made Promise: ", z)
	}
	go func() {
		var x P
		defer func() {
			r := recover()
			if r != nil {
				if DebugGo > 0 {
					println("#go# BAD Promise: ", z, r)
					PrintStack(r)
				}
				z.Ch <- EitherPOrError{Left: r, Right: nil}
			} else {
				if DebugGo > 0 {
					println("#go# OK Promise: ", z, x)
				}
				z.Ch <- EitherPOrError{Left: nil, Right: x}
			}
		}()
		if DebugGo > 0 {
			println("#go# Running Promise: ", z)
		}
		x = fn()
		if DebugGo > 0 {
			println("#go# Ran Promise: ", z)
		}
	}()
	if DebugGo > 0 {
		println("#go# Started Promise: ", z)
	}
	return z
}

type C_promise struct {
	C_object
	Ch chan EitherPOrError
}

func (o *C_promise) PtrC_promise() *C_promise {
	return o
}

type meth_C_promise_Wait struct {
	PBase
	Promise *C_promise
}

func (o *C_promise) GET_Wait() P {
	z := &meth_C_promise_Wait{Promise: o}
	z.SetSelf(z)
	return z
}
func (o *meth_C_promise_Wait) Call0() P {
	ch := o.Promise.Ch
	if ch == nil {
		panic("Wait() called more than once on promise")
	}
	x := <-ch
	o.Promise.Ch = nil // Don't allow another Wait.
	if x.Left != nil {
		panic(x.Left)
	}
	return x.Right
}

type EitherPOrError struct {
	Right P
	Left  interface{}
}

type void struct{}

// C_generator is the channel begin a yielding producer and a consuming for loop.
type C_generator struct {
	C_object
	Ready    chan *void
	Result   chan EitherPOrError
	Finished bool
}

const GENERATOR_ASYNC = 5

func NewGenerator() *C_generator {
	z := &C_generator{
		Ready:  make(chan *void, GENERATOR_ASYNC),
		Result: make(chan EitherPOrError, GENERATOR_ASYNC),
	}
	z.SetSelf(z)
	// Signal the coroutine so it can run asynchronously.
	for i := 0; i < GENERATOR_ASYNC; i++ {
		z.Ready <- nil
	}
	return z
}

func (o *C_generator) PtrC_generator() *C_generator {
	return o
}

func (o *C_generator) Iter() Nexter { return o }
func (o *C_generator) List() []P {
	var z []P
	for {
		x, gotOne := o.Next()
		if !gotOne {
			break
		}
		z = append(z, x)
	}
	o.Enough()
	return z
}

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

	//@ Say("##### NEXT", either, ok)

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
	//@ Say("FINISH")
	if !o.Finished {
		o.Finished = true
		close(o.Result)
	}
}

// Wait is called by the producer, once at the start, and once after each yield.
// Wait returns false if the consumer said Enough.
// TODO:  Don't wait, to achieve concurrency.  Let the user decide the Result channel buffer size.
func (o *C_generator) Wait() bool {
	_, ok := <-o.Ready
	//@ Say("WAIT ->", ok)
	return ok
}

type PBase struct {
	Self P
}

func (o *PBase) GetPBase() *PBase     { return o }
func (o *PBase) GetSelf() P           { return o.Self }
func (o *PBase) SetSelf(a P)          { o.Self = a }
func (o *PBase) Field(field string) P { panic(Bad("Receiver cannot Field", o.Self, o, field)) }

//func (o *PBase) FieldGets(field string, x P) P {
//	panic(Bad("Receiver cannot FieldGets", o.Self, o, field, x))
//}
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
func (o *PBase) DelItemSlice(i, j P)  { panic(Bad("Receiver cannot DelItemSlice: ", o.Self)) }
func (o *PBase) Iter() Nexter         { panic(Bad("Receiver cannot Iter: ", o.Self)) }
func (o *PBase) List() []P            { panic(Bad("Receiver cannot List: ", o.Self)) }
func (o *PBase) Dict() map[string]P   { panic(Bad("Receiver cannot Dict: ", o.Self)) }

func (o *PBase) Add(a P) P        { panic(Bad("Receiver cannot Add: ", o.Self, a)) }
func (o *PBase) Sub(a P) P        { panic(Bad("Receiver cannot Sub: ", o.Self, a)) }
func (o *PBase) Mul(a P) P        { panic(Bad("Receiver cannot Mul: ", o.Self, a)) }
func (o *PBase) Div(a P) P        { panic(Bad("Receiver cannot Div: ", o.Self, a)) }
func (o *PBase) IDiv(a P) P       { panic(Bad("Receiver cannot IDiv: ", o.Self, a)) }
func (o *PBase) Mod(a P) P        { panic(Bad("Receiver cannot Mod: ", o.Self, a)) }
func (o *PBase) Pow(a P) P        { panic(Bad("Receiver cannot Pow: ", o.Self, a)) }
func (o *PBase) BitAnd(a P) P     { panic(Bad("Receiver cannot BitAnd: ", o.Self, a)) }
func (o *PBase) BitOr(a P) P      { panic(Bad("Receiver cannot BitOr: ", o.Self, a)) }
func (o *PBase) BitXor(a P) P     { panic(Bad("Receiver cannot BitXor: ", o.Self, a)) }
func (o *PBase) ShiftLeft(a P) P  { panic(Bad("Receiver cannot ShiftLeft: ", o.Self, a)) }
func (o *PBase) ShiftRight(a P) P { panic(Bad("Receiver cannot ShiftRight: ", o.Self, a)) }
func (o *PBase) UnsignedShiftRight(a P) P {
	panic(Bad("Receiver cannot UnsignedShiftRight: ", o.Self, a))
}

func (o *PBase) IAdd(a P) { panic(Bad("Receiver cannot IAdd: ", o.Self, a)) }
func (o *PBase) ISub(a P) { panic(Bad("Receiver cannot ISub: ", o.Self, a)) }
func (o *PBase) IMul(a P) { panic(Bad("Receiver cannot IMul: ", o.Self, a)) }

func (o *PBase) EQ(a P) bool { return o.Self.Compare(a) == 0 }
func (o *PBase) NE(a P) bool { return o.Self.Compare(a) != 0 }
func (o *PBase) LT(a P) bool { return o.Self.Compare(a) < 0 }
func (o *PBase) LE(a P) bool { return o.Self.Compare(a) <= 0 }
func (o *PBase) GT(a P) bool { return o.Self.Compare(a) > 0 }
func (o *PBase) GE(a P) bool { return o.Self.Compare(a) >= 0 }
func (o *PBase) Compare(a P) int {
	// Default comparision uses address in memory.
	x := R.ValueOf(o).Pointer()
	y := R.ValueOf(a.GetPBase()).Pointer()
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	}
	return 0
}
func (o *PBase) Bool() bool     { return true } // Most things are true.
func (o *PBase) UnaryMinus() P  { panic(Bad("Receiver cannot UnaryMinus", o.Self)) }
func (o *PBase) UnaryPlus() P   { panic(Bad("Receiver cannot UnaryPlus", o.Self)) }
func (o *PBase) UnaryInvert() P { panic(Bad("Receiver cannot UnaryInvert", o.Self)) }

func (o *PBase) Int() int64            { panic(Bad("Receiver cannot Int", o.Self)) }
func (o *PBase) Float() float64        { panic(Bad("Receiver cannot Float", o.Self)) }
func (o *PBase) Complex() complex128   { panic(Bad("Receiver cannot Complex", o.Self)) }
func (o *PBase) Contents() interface{} { return o.Self }

func (o *PBase) Type() P       { return MkStr(F("%T", o.Self)) }
func (o *PBase) Bytes() []byte { panic(Bad("Receiver cannot Bytes", o.Self)) }
func (o *PBase) String() string {
	if o.Self == nil {
		panic("PBase:  Why is o.Self NIL?")
	}
	return o.Self.Show()
}
func (o *PBase) Pickle(w *bytes.Buffer) { panic(Bad("Receiver cannot Pickle", o.Self)) }
func (o *PBase) Repr() string           { return o.Self.String() }
func (o *PBase) Show() string {
	if o.Self == nil {
		panic("OHNO: o.Self == nil")
	}
	return ShowP(o.Self, SHOW_DEPTH)
}

func ShowP(a P, depth int) string {
	m := make(map[string]R.Value)
	return ShowPmap(a, depth, m, false)
}
func ShowPmap(a P, depth int, m map[string]R.Value, anon bool) string {
	if a == nil {
		return "P(nil) "
	}

	r := R.ValueOf(a) // TODO:  I don't like this code.
	if !r.IsValid() {
		panic("INVALID")
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
		ptr := r.Addr().Pointer()
		mapkey := F("%s@%x", t.Name(), ptr)

		if !anon {
			buf.WriteString(F(" %s@%04d{ ", t.Name(), (ptr/4)%9999))
		} // XYZZY

		_, ok := m[mapkey]
		m[mapkey] = r
		if !ok && depth > 0 {
			for i := 0; i < t.NumField(); i++ {
				f := t.Field(i)
				k := f.Name
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
					buf.WriteString(F("%s=%s ", k, ShowPmap(x, depth-1, m, false)))
				case []P:
					buf.WriteString(F("%s=[%d]{ ", k, len(x)))
					for _, xe := range x {
						buf.WriteString(F(" %s,", ShowPmap(xe, depth-1, m, false)))
					}
					buf.WriteString("} ")
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
						thing := v.Addr().Interface()
						if inner, ok := thing.(P); ok {
							buf.WriteString(F("%s", ShowPmap(inner, depth, m, f.Anonymous)))
						} else {
							buf.WriteString(F("%v", v.Interface()))
						}
					} else {
						buf.WriteString(F("%s:%s ", k, v.Type().Name()))
					}
				}
			}
		}
		if !anon {
			buf.WriteString("} ")
		} // XYZZY
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

func MkGo(a interface{}) *PGo { z := &PGo{V: R.ValueOf(a)}; z.Self = z; return z }
func MkValue(a R.Value) *PGo  { z := &PGo{V: a}; z.Self = z; return z }

func Mkint(n int) *PInt         { z := &PInt{N: int64(n)}; z.Self = z; return z }
func MkInt(n int64) *PInt       { z := &PInt{N: n}; z.Self = z; return z }
func MkFloat(f float64) *PFloat { z := &PFloat{F: f}; z.Self = z; return z }
func MkStr(s string) *PStr      { z := &PStr{S: s}; z.Self = z; return z }
func MkByt(yy []byte) *PByt     { z := &PByt{YY: yy}; z.Self = z; return z }
func MkStrs(ss []string) *PList {
	pp := make([]P, len(ss))
	for i, s := range ss {
		pp[i] = MkStr(s)
	}
	return MkList(pp)
}

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
	if (len(pp) % 2) == 1 {
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

func WriteC(w *bytes.Buffer, c rune) {
	n, err := w.Write([]byte{byte(c)})
	if n != 1 {
		panic("WriteB: could not Write")
	}
	if err != nil {
		panic(F("WriteB: error during Write: %s", err))
	}
}

func (o *PNone) Bool() bool             { return false }
func (o *PNone) String() string         { return "None" }
func (o *PNone) Repr() string           { return "None" }
func (o *PNone) Contents() interface{}  { return nil }
func (o *PNone) Pickle(w *bytes.Buffer) { w.WriteByte(RypNone) }

func (o *PBool) Pickle(w *bytes.Buffer) {
	if o.B {
		w.WriteByte(RypTrue)
	} else {
		w.WriteByte(RypFalse)
	}
}
func (o *PBool) Contents() interface{} { return o.B }
func (o *PBool) Bool() bool            { return o.B }
func (o *PBool) Int() int64 {
	if o.B {
		return 1
	} else {
		return 0
	}
}
func (o *PBool) Float() float64 {
	if o.B {
		return 1.0
	} else {
		return 0.0
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
func (o *PBool) Type() P { return G_bool }
func (o *PBool) Compare(a P) int {
	x := o.Float()
	y := a.Float()
	switch {
	case x < y:
		return -1
	case x > y:
		return 1
	case x == y:
		return 0
	}
	panic(F("Cannot compare *PFloat to %T", a))
}

func (o *PInt) UnaryMinus() P  { return MkInt(0 - o.N) }
func (o *PInt) UnaryPlus() P   { return o }
func (o *PInt) UnaryInvert() P { return MkInt(int64(-1) ^ o.N) }
func (o *PInt) Add(a P) P      { return MkInt(o.N + a.Int()) }
func (o *PInt) Sub(a P) P      { return MkInt(o.N - a.Int()) }
func (o *PInt) Mul(a P) P {
	switch x := a.(type) {
	case *PInt:
		return MkInt(o.N * x.N)
	case *PFloat:
		return MkFloat(float64(o.N) * x.F)
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
	case *PGo:
		bt := x.V.Type()
		switch bt.Kind() {
		case R.Int64, R.Int:
			return MkInt(o.N * x.V.Int())
		}
	}
	panic("Cannot multply int times whatever")
}
func (o *PInt) Div(a P) P                { return MkInt(o.N / a.Int()) }
func (o *PInt) IDiv(a P) P               { return MkInt(o.N / a.Int()) }
func (o *PInt) Mod(a P) P                { return MkInt(o.N % a.Int()) }
func (o *PInt) BitAnd(a P) P             { return MkInt(o.N & a.Int()) }
func (o *PInt) BitOr(a P) P              { return MkInt(o.N | a.Int()) }
func (o *PInt) BitXor(a P) P             { return MkInt(o.N ^ a.Int()) }
func (o *PInt) ShiftLeft(a P) P          { return MkInt(o.N << uint64(a.Int())) }
func (o *PInt) ShiftRight(a P) P         { return MkInt(o.N >> uint64(a.Int())) }
func (o *PInt) UnsignedShiftRight(a P) P { return MkInt(int64(uint64(o.N) >> uint64(a.Int()))) }
func (o *PInt) EQ(a P) bool              { return (o.N == a.Int()) }
func (o *PInt) NE(a P) bool              { return (o.N != a.Int()) }
func (o *PInt) LT(a P) bool              { return (o.N < a.Int()) }
func (o *PInt) LE(a P) bool              { return (o.N <= a.Int()) }
func (o *PInt) GT(a P) bool              { return (o.N > a.Int()) }
func (o *PInt) GE(a P) bool              { return (o.N >= a.Int()) }
func (o *PInt) Compare(a P) int {
	switch b := a.(type) {
	case *PInt:
		switch {
		case o.N < b.N:
			return -1
		case o.N > b.N:
			return 1
		}
		return 0
	case *PFloat:
		switch {
		case float64(o.N) < b.F:
			return -1
		case float64(o.N) > b.F:
			return 1
		}
		return 0
	case *PBool:
		c := b.Int()
		switch {
		case o.N < c:
			return -1
		case o.N > c:
			return 1
		}
		return 0
	}
	panic(F("Cannot compare *PInt to %T", a))
}
func (o *PInt) Int() int64            { return o.N }
func (o *PInt) Float() float64        { return float64(o.N) }
func (o *PInt) String() string        { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string          { return o.String() }
func (o *PInt) Bool() bool            { return o.N != 0 }
func (o *PInt) Type() P               { return G_int }
func (o *PInt) Contents() interface{} { return o.N }
func (o *PInt) Pickle(w *bytes.Buffer) {
	n := RypIntLenMinus1(o.N)
	w.WriteByte(byte(RypInt + n))
	RypWriteInt(w, o.N)
}

func (o *PFloat) Add(a P) P   { return MkFloat(o.F + a.Float()) }
func (o *PFloat) Sub(a P) P   { return MkFloat(o.F - a.Float()) }
func (o *PFloat) Mul(a P) P   { return MkFloat(o.F * a.Float()) }
func (o *PFloat) Div(a P) P   { return MkFloat(o.F / a.Float()) }
func (o *PFloat) IDiv(a P) P  { return MkInt(int64(o.F / a.Float())) }
func (o *PFloat) EQ(a P) bool { return (o.F == a.Float()) }
func (o *PFloat) NE(a P) bool { return (o.F != a.Float()) }
func (o *PFloat) LT(a P) bool { return (o.F < a.Float()) }
func (o *PFloat) LE(a P) bool { return (o.F <= a.Float()) }
func (o *PFloat) GT(a P) bool { return (o.F > a.Float()) }
func (o *PFloat) GE(a P) bool { return (o.F >= a.Float()) }
func (o *PFloat) Compare(a P) int {
	c := a.Float()
	switch {
	case o.F < c:
		return -1
	case o.F > c:
		return 1
	case o.F == c:
		return 0
	}
	panic(F("Cannot compare *PFloat to %T", a))
}
func (o *PFloat) Int() int64            { return int64(o.F) }
func (o *PFloat) Float() float64        { return o.F }
func (o *PFloat) String() string        { return strconv.FormatFloat(o.F, 'g', -1, 64) }
func (o *PFloat) Repr() string          { return o.String() }
func (o *PFloat) Bool() bool            { return o.F != 0 }
func (o *PFloat) Type() P               { return G_float }
func (o *PFloat) Contents() interface{} { return o.F }
func (o *PFloat) Pickle(w *bytes.Buffer) {
	x := int64(math.Float64bits(o.F))
	n := RypIntLenMinus1(x)
	w.WriteByte(byte(RypFloat + n))
	RypWriteInt(w, x)
}
func (o *PStr) Iter() Nexter {
	var pp []P
	for _, r := range o.S {
		pp = append(pp, MkStr(string(r)))
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}
func (o *PStr) Pickle(w *bytes.Buffer) {
	l := int64(len(o.S))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypStr + n))
	RypWriteInt(w, l)
	w.WriteString(o.S)
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
	n := int64(len(o.S))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.S))
		}
		if i < 0 {
			panic(F("First slicing index on PStr too small: %d", i))
		}
	}
	if i > n {
		panic(F("First slicing index on PStr too large: %d > len: %d", i, n))
	}
	if y == None {
		j = int64(len(o.S))
	} else {
		j = y.Int()
		if j < 0 {
			j += int64(len(o.S))
		}
		if j < 0 {
			panic(F("Second slicing index on PStr too small: %d", j))
		}
	}
	if j > n {
		panic(F("Second slicing index on PStr too large: %d > len: %d", j, n))
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
			// THIS NEEDS WORK.
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
		// THIS NEEDS WORK.
		panic(Show("Bad value on rhs in Mod:", a))
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
func (o *PStr) Add(a P) P   { return MkStr(o.S + a.String()) }
func (o *PStr) EQ(a P) bool { return (o.S == a.String()) }
func (o *PStr) NE(a P) bool { return (o.S != a.String()) }
func (o *PStr) LT(a P) bool { return (o.S < a.String()) }
func (o *PStr) LE(a P) bool { return (o.S <= a.String()) }
func (o *PStr) GT(a P) bool { return (o.S > a.String()) }
func (o *PStr) GE(a P) bool { return (o.S >= a.String()) }
func (o *PStr) Compare(a P) int {
	switch b := a.(type) {
	case *PStr:
		switch {
		case o.S < b.S:
			return -1
		case o.S > b.S:
			return 1
		}
		return 0
	}
	panic(F("Cannot compare *PStr to %T", a))
}
func (o *PStr) Int() int64     { return CI(strconv.ParseInt(o.S, 10, 64)) }
func (o *PStr) String() string { return o.S }
func (o *PStr) Bytes() []byte  { return []byte(o.S) }
func (o *PStr) Len() int       { return len(o.S) }
func (o *PStr) Repr() string   { return F("%q", o.S) }
func (o *PStr) Type() P        { return G_str }

func (o *PByt) Iter() Nexter {
	var pp []P
	for _, r := range o.YY {
		pp = append(pp, Mkint(int(r)))
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}
func (o *PByt) Pickle(w *bytes.Buffer) {
	l := int64(len(o.YY))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypByt + n))
	RypWriteInt(w, l)
	w.Write(o.YY)
}
func (o *PByt) Contents() interface{} { return o.YY }
func (o *PByt) Bool() bool            { return len(o.YY) != 0 }
func (o *PByt) GetItem(a P) P {
	i := int(a.Int())
	if i < 0 {
		i += len(o.YY)
	}
	return Mkint(int(o.YY[i]))
}
func (o *PByt) SetItem(a P, x P) {
	i := int(a.Int())
	if i < 0 {
		i += len(o.YY)
	}
	o.YY[i] = byte(x.Int())
}

func (o *PByt) GetItemSlice(x, y, z P) P {
	var i, j int64
	n := int64(len(o.YY))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.YY))
		}
		if i < 0 {
			panic(F("First slicing index on PByt too small: %d", i))
		}
	}
	if i > n {
		panic(F("First slicing index on PByt too large: %d > len: %d", i, n))
	}
	if y == None {
		j = int64(len(o.YY))
	} else {
		j = y.Int()
		if j < 0 {
			j += int64(len(o.YY))
		}
		if j < 0 {
			panic(F("Second slicing index on PByt too small: %d", j))
		}
	}
	if j > n {
		panic(F("Second slicing index on PByt too large: %d > len: %d", j, n))
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
func (o *PByt) Add(a P) P {
	aa := a.Bytes()
	var zz []byte
	zz = append(zz, o.YY...)
	zz = append(zz, aa...)
	return MkByt(zz)
}

func (o *PByt) EQ(a P) bool { return (string(o.YY) == a.String()) }
func (o *PByt) NE(a P) bool { return (string(o.YY) != a.String()) }
func (o *PByt) LT(a P) bool { return (string(o.YY) < a.String()) }
func (o *PByt) LE(a P) bool { return (string(o.YY) <= a.String()) }
func (o *PByt) GT(a P) bool { return (string(o.YY) > a.String()) }
func (o *PByt) GE(a P) bool { return (string(o.YY) >= a.String()) }

func (o *PByt) String() string { return string(o.YY) }
func (o *PByt) Show() string   { return o.Repr() }
func (o *PByt) Bytes() []byte  { return o.YY }
func (o *PByt) Len() int       { return len(o.YY) }
func (o *PByt) Repr() string   { return F("byt(%q)", string(o.YY)) }
func (o *PByt) Type() P        { return G_byt }
func (o *PByt) List() []P {
	zz := make([]P, len(o.YY))
	for i, x := range o.YY {
		zz[i] = Mkint(int(x))
	}
	return zz
}
func (o *PByt) Compare(a P) int {
	switch b := a.(type) {
	case *PByt:
		switch {
		case string(o.YY) < string(b.YY):
			return -1
		case string(o.YY) > string(b.YY):
			return 1
		}
		return 0
	}
	panic(F("Cannot compare *PByt to %T", a))
}

func (o *PTuple) Pickle(w *bytes.Buffer) {
	l := int64(len(o.PP))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypTuple + n))
	RypWriteInt(w, l)
	for _, x := range o.PP {
		x.Pickle(w)
	}
}
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
func (o *PTuple) GetItemSlice(x, y, z P) P {
	var i, j int64
	n := int64(len(o.PP))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.PP))
			if i < 0 {
				panic(F("First slicing index on PTuple too small: %d", i))
			}
		}
	}
	if i > n {
		panic(F("First slicing index on PTuple too large: %d > len: %d", i, n))
	}
	if y == None {
		j = int64(len(o.PP))
	} else {
		j = y.Int()
		if j < 0 {
			j += int64(len(o.PP))
			if j < 0 {
				panic(F("Second slicing index on PTuple too small: %d", j))
			}
		}
	}
	if j > n {
		panic(F("Second slicing index on PTuple too large: %d > len: %d", j, n))
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkList(o.PP[i:j])
	return r
}
func (o *PTuple) String() string { return o.Repr() }
func (o *PTuple) Type() P        { return G_tuple }
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

func (o *PTuple) Compare(a P) int {
	switch b := a.(type) {
	case *PTuple:
		on := len(o.PP)
		bn := len(b.PP)
		for i := 0; true; i++ {
			if on <= i {
				if bn <= i {
					// Both ended before now, so they were equal.
					return 0
				} else {
					// bn ls longer, so we are less than it.
					return -1
				}
			} else {
				if bn < i {
					// we are longer, so we are greater than it.
					return 1
				} else {
					// Neither ended yet.
					cmp := o.PP[i].Compare(b.PP[i])
					if cmp != 0 {
						return cmp
					}
					// But if the are equal, continue with next slot.
				}
			}
		}
	}
	panic(F("Cannot compare *PTuple to %T", a))
}

func (o *PList) Compare(a P) int {
	switch b := a.(type) {
	case *PList:
		on := len(o.PP)
		bn := len(b.PP)
		for i := 0; true; i++ {
			if on <= i {
				if bn <= i {
					// Both ended before now, so they were equal.
					return 0
				} else {
					// bn ls longer, so we are less than it.
					return -1
				}
			} else {
				if bn <= i {
					// we are longer, so we are greater than it.
					return 1
				} else {
					// Neither ended yet.
					cmp := o.PP[i].Compare(b.PP[i])
					if cmp != 0 {
						return cmp
					}
					// But if the are equal, continue with next slot.
				}
			}
		}
	}
	panic(F("Cannot compare *PList to %T", a))
}

func (o *PList) Pickle(w *bytes.Buffer) {
	l := int64(len(o.PP))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypList + n))
	RypWriteInt(w, l)
	for _, x := range o.PP {
		x.Pickle(w)
	}
}
func (o *PList) Add(a P) P {
	b := a.List()
	z := make([]P, 0, len(o.PP)+len(b))
	z = append(z, o.PP...)
	z = append(z, b...)
	return MkList(z)
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
func (o *PList) Bytes() []byte {
	zz := make([]byte, len(o.PP))
	for i, x := range o.PP {
		zz[i] = byte(x.Int())
	}
	return zz
}
func (o *PList) Len() int { return len(o.PP) }
func (o *PList) GetItem(x P) P {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.PP))
	}
	return o.PP[i]
}
func (o *PList) GetItemSlice(x, y, z P) P {
	var i, j int64
	n := int64(len(o.PP))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += int64(len(o.PP))
			if i < 0 {
				panic(F("First slicing index on PList too small: %d", i))
			}
		}
	}
	if i > n {
		panic(F("First slicing index on PList too large: %d > len: %d", i, n))
	}
	if y == None {
		j = int64(len(o.PP))
	} else {
		j = y.Int()
		if j < 0 {
			j += int64(len(o.PP))
		}
		if j < 0 {
			panic(F("Second slicing index on PList too small: %d", j))
		}
	}
	if j > n {
		panic(F("Second slicing index on PList too large: %d > len: %d", j, n))
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkList(o.PP[i:j])
	return r
}

func (o *PList) String() string { return o.Repr() }
func (o *PList) Type() P        { return G_list }
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

func (o *PList) DelItem(x P) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	i := int(x.Int())
	copy(a[i:], a[i+1:])
	a[len(a)-1] = nil // or the zero value of T
	o.PP = a[:len(a)-1]
}

func (o *PList) DelItemSlice(x, y P) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	i, j := int(x.Int()), int(y.Int())
	copy(a[i:], a[j:])
	for k, n := len(a)-j+i, len(a); k < n; k++ {
		a[k] = nil // or the zero value of T
	} // for k
	o.PP = a[:len(a)-j+i]
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

func (o *PDict) Pickle(w *bytes.Buffer) {
	l := int64(len(o.PPP))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypDict + n))
	RypWriteInt(w, l)
	for k, v := range o.PPP {
		MkStr(k).Pickle(w)
		v.Pickle(w)
	}
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
	key := a.String()
	o.PPP[key] = x
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
func (o *PDict) Type() P        { return G_dict }
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
		val, ok := o.PPP[keys[i]]
		if !ok {
			panic("PDict Repr Not Ok")
		}
		buf.WriteString(F("%q: %s", keys[i], val.Repr()))
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
func (o *PDict) Dict() map[string]P {
	return o.PPP
}
func (o *PDict) DelItem(i P) {
	delete(o.PPP, i.String())
}

type meth_PDict_keys struct {
	PBase
	dict *PDict
}

func (o *PDict) GET_keys() P {
	z := &meth_PDict_keys{dict: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PDict_keys) Call0() P {
	pp := make([]P, 0, len(o.dict.PPP))
	for k, _ := range o.dict.PPP {
		pp = append(pp, MkStr(k))
	}
	return MkList(pp)
}

type meth_PDict_values struct {
	PBase
	dict *PDict
}

func (o *PDict) GET_values() P {
	z := &meth_PDict_values{dict: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PDict_values) Call0() P {
	pp := make([]P, 0, len(o.dict.PPP))
	for _, v := range o.dict.PPP {
		pp = append(pp, v)
	}
	return MkList(pp)
}

type meth_PDict_items struct {
	PBase
	dict *PDict
}

func (o *PDict) GET_items() P {
	z := &meth_PDict_items{dict: o}
	z.SetSelf(z)
	return z
}

func (o *meth_PDict_items) Call0() P {
	pp := make([]P, 0, len(o.dict.PPP))
	for k, v := range o.dict.PPP {
		pp = append(pp, MkTuple([]P{MkStr(k), v}))
	}
	return MkList(pp)
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

func (o *C_object) Repr() string {
	return ShowP(o.Self, SHOW_DEPTH)
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

var PBaseType = R.TypeOf(PBase{})
var ByteSliceType = R.TypeOf([]byte{})

func (o *C_object) PickleFields(w *bytes.Buffer, v R.Value) {
	t := v.Type()
	if t.Kind() != R.Struct {
		panic(F("PickleFields expected Struct: %q", t.String()))
	}
	nf := t.NumField()
	for i := 0; i < nf; i++ {
		f := t.Field(i)
		if f.Anonymous {
			if f.Type != PBaseType {
				o.PickleFields(w, v.Field(i))
			}
		} else {
			RypWriteLabel(w, f.Name)
			v.Field(i).Interface().(P).GetSelf().Pickle(w)
		}
	}
}

func (o *C_object) Pickle(w *bytes.Buffer) {
	w.WriteByte(RypClass)
	RypWriteLabel(w, R.ValueOf(o.Self).Type().Elem().String())
	o.PickleFields(w, R.ValueOf(o.Self).Elem())
	w.WriteByte(0) // Like a 0-length label, to terminate fields.
}

func NewList() *PList {
	z := &PList{PP: make([]P, 0)}
	z.Self = z
	return z
}

func CopyPs(pp []P) []P {
	zz := make([]P, len(pp))
	copy(zz, pp)
	return zz
}
func CopyList(aa *PList) *PList {
	z := &PList{PP: CopyPs(aa.PP)}
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
	if t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	return t
}

func MaybeDerefAll(t R.Value) R.Value {
	for t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	return t
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

func GetItemMap(r R.Value, x P) P {
	k := AdaptForCall(x, r.Type().Key())
	v := r.MapIndex(k)
	if !v.IsValid() {
		panic(F("Map key not found"))
	}
	return AdaptForReturn(v)
}

func DemoteInt64(x64 int64) int {
	x := int(x64)
	if int64(x) != x64 {
		panic(F("Cannot demote int64: %d vs %d", x64, x))
	}
	return x
}
func SliceGetItem(r R.Value, x P) P {
	n := r.Len()
	k := DemoteInt64(x.Int())
	if k < -n || n <= k {
		panic(F("Slice key out of range: %d, len = %d", k, n))
	}
	if k < 0 {
		k += n
	}
	v := r.Index(k)
	if !v.IsValid() {
		panic(F("Map key not found"))
	}
	return AdaptForReturn(v)
}

func (o *PGo) Contents() interface{} { return o.V.Interface() }
func (o *PGo) Bool() bool {
	r := o.V
	if SafeIsNil(r) {
		// Nil is false, as in Python.
		return false
	}
	r = MaybeDeref(r)
	switch r.Kind() {
	case R.Map, R.Slice, R.Array:
		// Zero length is false, as in Python.
		return r.Len() > 0
	}
	return true
}

func (o *PGo) Len() int {
	r := MaybeDerefAll(o.V)
	switch r.Kind() {
	case R.Map, R.Slice, R.Array:
		return r.Len()
	}

	panic(F("Cannot get length of PGo type %t", o.V.Type()))
}
func (o *PGo) GetItem(x P) P {
	r := MaybeDerefAll(o.V)
	switch r.Kind() {
	case R.Map:
		return GetItemMap(r, x)
	case R.Slice, R.Array:
		return SliceGetItem(r, x)
	}

	panic(F("Cannot GetItem on PGo type %t", o.V.Type()))
}
func (o *PGo) GetItemSlice(a, b, c P) P {
	r := MaybeDerefAll(o.V)
	switch r.Kind() {
	case R.Slice, R.Array:
		n := r.Len()

		var a2 int
		if a != None {
			a2 = int(a.Int())
		}
		if a2 < 0 {
			a2 += n
			if a2 < 0 {
				panic(F("First slicing index on PGo too small: %d", a2))
			}
		}
		if a2 > n {
			panic(F("First slicing index on PGo too large: %d > len: %d", a2, n))
		}

		var b2 int
		if b == None {
			b2 = n
		} else {
			b2 = int(b.Int())
		}
		if b2 < 0 {
			b2 += n
			if b2 < 0 {
				panic(F("Second slicing index on PGo too small: %d", b2))
			}
		}
		if b2 > n {
			panic(F("Second slicing index on PGo too large: %d > len: %d", b2, n))
		}

		// TODO: c2 = int(c.Int())
		z := r.Slice(a2, b2)
		return MkValue(z)
	}

	panic(F("Cannot GetItemSlice on PGo type %t", o.V.Type()))
}
func (g *PGo) Repr() string {
	return F("PGo.Repr{%#v}", g.V.Interface())
}
func (g *PGo) String() string {
	// g0 := MaybeDeref(g.V)
	g0 := g.V
	i0 := g0.Interface()

	switch x := i0.(type) {
	case fmt.Stringer:
		return x.String()
	case []byte:
		return string(x)
	case int:
		return F("%d", x)
	case int64:
		return F("%d", x)
	case uint:
		return F("%d", x)
	case uint64:
		return F("%d", x)
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
	return F("PGo{%#v}", i0)
}

var Int64Type = R.TypeOf(int64(0))
var IntType = R.TypeOf(int(0))
var PType = R.TypeOf(new(P)).Elem()

func (o *PGo) Int() int64 {
	x := o.V
	t := x.Type()
	switch {
	case t.ConvertibleTo(Int64Type):
		return x.Convert(Int64Type).Int()
	case t.ConvertibleTo(IntType):
		return x.Convert(IntType).Int()
	}
	x = MaybeDeref(x)
	t = x.Type()
	switch {
	case t.ConvertibleTo(Int64Type):
		return x.Convert(Int64Type).Int()
	case t.ConvertibleTo(IntType):
		return x.Convert(IntType).Int()
	}
	panic(F("PGo cannot convert to int64: %s", o.Show()))
}
func InvokeMap(r R.Value, field string, aa []P) P {
	switch {
	case field == "keys" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]P, len(keys))
		for i, e := range keys {
			pp[i] = AdaptForReturn(e)
		}
		return MkList(pp)
	case field == "values" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]P, len(keys))
		for i, e := range keys {
			pp[i] = AdaptForReturn(r.MapIndex(e))
		}
		return MkList(pp)
	case field == "items" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]P, len(keys))
		for i, e := range keys {
			pp[i] = MkTuple([]P{AdaptForReturn(e), AdaptForReturn(r.MapIndex(e))})
		}
		return MkList(pp)
	case field == "get" && (len(aa) == 1 || len(aa) == 2):
		key := AdaptForCall(aa[0], r.Type().Key())
		v := r.MapIndex(key)
		if v.IsValid() {
			return AdaptForReturn(v)
		} else {
			if len(aa) == 1 {
				return None
			} else {
				return aa[1]
			}
		}
	}
	panic(F("Method on Map Type %q does not exist: %s", r.Type(), field))
}

func (g *PGo) Invoke(field string, aa ...P) P {
	// println(F("## Invoking Method %q On PGo type %T kind %v", field, g.V.Interface(), g.V.Kind()))
	r := g.V
	// println(F("TYPE1 %q", r.Type()))
	if meth, ok := r.Type().MethodByName(field); ok && meth.Func.IsValid() {
		return FinishInvokeOrCall(meth.Func, r, aa)
	}
	if r.Kind() == R.Map {
		return InvokeMap(r, field, aa)
	}

	r = MaybeDeref(r)
	// println(F("TYPE2 %q", r.Type()))
	if meth, ok := r.Type().MethodByName(field); ok && meth.Func.IsValid() {
		return FinishInvokeOrCall(meth.Func, r, aa)
	}
	if r.Kind() == R.Map {
		return InvokeMap(r, field, aa)
	}

	panic(F("Method on type %q does not exist: %s", g.V.Type(), field))
}

func (g *PGo) CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P {
	if len(kv1) > 0 || len(kv2) > 0 {
		panic("Cannot call GO with named (keyword) parameters.")
	}
	if a2 == nil {
		return g.Call(a1...)
	}
	if a1 == nil {
		return g.Call(a2...)
	}
	var aa []P
	aa = append(aa, a1...)
	aa = append(aa, a2...)
	return g.Call(aa...)
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

// TODO -- make List() the primary, instead of Iter()
func (g *PGo) List() []P {
	return g.Iter().(*PListIter).PP
}
func (g *PGo) Iter() Nexter {
	a := MaybeDeref(g.V)
	var pp []P

	switch a.Kind() {
	case R.Array, R.Slice:
		n := a.Len()
		pp = make([]P, n)
		for i := 0; i < n; i++ {
			pp[i] = AdaptForReturn(a.Index(i))
		}
	case R.Map:
		keys := a.MapKeys()
		pp = make([]P, len(keys))
		for i, e := range keys {
			pp[i] = AdaptForReturn(e)
		}
	default:
		Bad(F("*PGo cannot Iter() on kind %s, type=%T", a.Kind(), a.Interface()))
	}
	z := &PListIter{PP: pp}
	z.Self = z
	return z
}
func (g *PGo) SetItem(i P, x P) {
	a := MaybeDeref(g.V)

	switch a.Kind() {
	case R.Array, R.Slice:
		i2 := int(i.Int())
		x2 := AdaptForCall(x, a.Type().Elem())
		a.Index(i2).Set(x2)
	case R.Map:
		i2 := R.ValueOf(i.String())
		x2 := AdaptForCall(x, a.Type().Elem())
		a.SetMapIndex(i2, x2)
	default:
		Bad("*PGo cannot Iter() on kind %s", a.Kind())
	}
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
	ft := f.Type()
	numIn := ft.NumIn()

	args := make([]R.Value, lenIns)
	if ft.IsVariadic() {
		if lenIns < numIn-1 {
			Bad("call got %d args, want %d or more args", lenIns, numIn-1)
		}
		args[0] = rcvr
		for i, a := range aa {
			var desiredType R.Type
			if i >= numIn-1 {
				desiredType = ft.In(numIn - 1).Elem()
			} else {
				desiredType = ft.In(i)
			}
			args[i+lenRcvr] = AdaptForCall(a, desiredType)
		}
	} else {
		if lenIns != numIn {
			Bad("call got %d args, want %d args", lenIns, numIn)
		}
		if lenIns > 0 {
			args[0] = rcvr
		}
		for i, a := range aa {
			args[i+lenRcvr] = AdaptForCall(a, ft.In(i+lenRcvr))
		}
	}

	//@ println(F("##"))
	//@ for k, v := range aa {
	//@ println(F("##Arg %d was %q", k, v.Show()))
	//@ }

	//@ println(F("##"))
	//@ for k, v := range args {
	//@ println(F("##Arg %d is %#v", k, v.Interface()))
	//@ }

	//@ println(F("## CALLING %#v", f.Interface()))
	outs := f.Call(args)

	//@ for k, v := range outs {
	//@ println(F("##Result %d is %#v", k, v.Interface()))
	//@ }
	//@ println(F("##"))

	numOut := ft.NumOut()
	if numOut > 0 && ft.Out(numOut-1) == errorType {
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

func GoDeref(p P) P {
	switch x := p.(type) {
	case *PGo:
		switch x.V.Kind() {
		case R.Ptr:
			return MkValue(x.V.Elem())
		}
		panic(F("Cannot goderef non-pointer: %s", x.V.Kind()))
	}
	panic(F("Cannot goderef non-*Go value: %T", p))
}

func GoCast(want P, p P) P {
	typ := want.Contents().(R.Type)
	return MkValue(AdaptForCall(p, typ))
}

func AdaptForCall(v P, want R.Type) R.Value {
	if DebugAdapt > 0 {
		Say("AdaptForCall <<<<<<", v, want, F("%#v", v))
	}
	z := adaptForCall2(v, want)
	if DebugAdapt > 0 {
		Say("AdaptForCall >>>>>>", z)
	}
	return z
}
func adaptForCall2(v P, want R.Type) R.Value {
	// None & nil.
	switch want.Kind() {
	case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
		if v.Contents() == nil {
			if DebugAdapt > 0 {
				Say("AdaptForCall :::::: R.Zero")
			}
			return R.Zero(want)
		}
	}

	// Try builtin conversion:
	contents := v.Contents()
	vcontents := R.ValueOf(contents)
	tcontents := vcontents.Type()
	if tcontents.ConvertibleTo(want) {
		if DebugAdapt > 0 {
			Say("AdaptForCall :::::: vcontents.Convert")
		}
		return vcontents.Convert(want)
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
			switch vx := v.(type) {
			case *PStr:
				bb := make([]byte, v.Len())
				copy(bb, v.String())
				return R.ValueOf(bb)
			case *PByt:
				return R.ValueOf(vx.YY)
			}
			// panic(F("AdaptForCall: Cannot convert %T to []byte", v))
		}

		// For the "in" case.  TODO: "in out"?
		if tcontents.Kind() == R.Slice {
			n := vcontents.Len()
			sl := R.MakeSlice(want, n, n)
			for i := 0; i < n; i++ {
				v1 := vcontents.Index(i)
				var v2 R.Value
				if vp, ok := v1.Interface().(P); ok {
					v2 = AdaptForCall(vp, want.Elem())
				} else {
					v2 = AdaptForCall(MkValue(v1), want.Elem())
				}
				sl.Index(i).Set(v2)
			}
			return sl
		}
	case R.Map:
		// For the "in" case.  TODO: "in out"?
		if tcontents.Kind() == R.Map {
			m := R.MakeMap(want)
			for _, k := range vcontents.MapKeys() {
				k2 := AdaptForCall(MkValue(k), want.Key())
				v1 := vcontents.MapIndex(k)
				// Broken?
				//if want.Elem() == PType {
				//  v1 = R.ValueOf(v1.Interface().(P).Contents())
				//}
				var v2 R.Value
				if vp, ok := v1.Interface().(P); ok {
					// v1 = R.ValueOf(vp.Contents())
					// Say("Converted inner P", vp, v1)
					v2 = AdaptForCall(vp, want.Elem())
				} else {
					v2 = AdaptForCall(MkValue(v1), want.Elem())
				}
				m.SetMapIndex(k2, v2)
			}
			return m
		}
	}
	if DebugAdapt > 0 {
		Say("AdaptForCall :::::: Not Case")
	}

	if want == typeInterfaceEmpty {
		if DebugAdapt > 0 {
			Say("AdaptForCall :::::: Interface Empty")
		}
		return R.ValueOf(v.Contents())
	}
	if DebugAdapt > 0 {
		Say("AdaptForCall :::::: Panic.")
	}
	panic(F("Cannot AdaptForCall: %s [%s] %q [%s] TO %s [%s]", v, R.TypeOf(v), v.Repr(), R.TypeOf(v.Contents()), want, want.Kind()))
}

func MakeFunction(v P, ft R.Type) R.Value {
	nin := ft.NumIn()
	if nin > 3 {
		panic(F("Not implemented: MakeFunction for %d args", nin))
	}

	return R.MakeFunc(ft, func(aa []R.Value) (zz []R.Value) {
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
		nout := ft.NumOut()
		switch nout {
		case 0: // pass
		case 1:
			zz = append(zz, AdaptForCall(r, ft.Out(0)))
		default:
			zz = make([]R.Value, nout)
			for i := 0; i < nout; i++ {
				zz[i] = AdaptForCall(r.GetItem(Mkint(i)), ft.Out(i))
			}
		}
		return
	})
}

func AdaptForReturn(v R.Value) P {
	if !v.IsValid() {
		panic("Invalid Value in AdaptForReturn")
	}

	switch v.Kind() {
	case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
		if v.IsNil() {
			return None
		}
	}

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
	case R.Float32:
		return MkFloat(v.Float())
	case R.Float64:
		return MkFloat(v.Float())
	case R.Bool:
		if v.Bool() {
			return True
		}
		return False
	case R.Slice:
		switch v.Type().Elem().Kind() {
		case R.Uint8:
			return MkByt(v.Convert(ByteSliceType).Interface().([]byte))
			// return MkByt(v.Interface().([]byte))
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
	if p, ok := v.Interface().(P); ok {
		return p
	}
	return MkValue(v)
}

func (g *PGo) Field(field string) P {
	t := g.V
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

var Classes map[string]R.Type

func init() {
	Classes = make(map[string]R.Type)
}

func init() {
	var tmp P = new(PBase)
	// Demonstrate these things implement P.
	tmp = new(PInt)
	tmp = new(PFloat)
	tmp = new(PList)
	tmp = new(PDict)
	tmp = new(C_object)
	_ = tmp
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

func GoElemType(pointedTo interface{}) P {
	return MkGo(R.TypeOf(pointedTo).Elem())
}

const (
	RypNone = (17 + iota) << 3
	RypTrue
	RypFalse
	RypInt
	RypFloat
	RypStr
	RypByt
	RypTuple
	RypList
	RypDict
	RypSet
	RypClass
	RypGob
)
const RypMask = 31 << 3

func RypLegend() {
	/*
		Say("RypNone", RypNone)
		Say("RypTrue", RypTrue)
		Say("RypFalse", RypFalse)
		Say("RypInt", RypInt)
		Say("RypFloat", RypFloat)
		Say("RypStr", RypStr)
		Say("RypByt", RypByt)
		Say("RypTuple", RypTuple)
		Say("RypList", RypList)
		Say("RypDict", RypDict)
		Say("RypSet", RypSet)
		Say("RypClass", RypClass)
		Say("RypGob", RypGob)
	*/
}

func RypIntLenMinus1(x int64) int {
	u := uint64(x)
	z := 0
	if u == 0 {
		return 0
	}
	for u != 0 {
		u >>= 8
		z++
	}
	return z - 1
}

func RypWriteInt(b *bytes.Buffer, x int64) {
	u := uint64(x)
	if u == 0 {
		b.WriteByte(0)
		return
	}
	for u != 0 {
		b.WriteByte(byte(u & 255))
		u >>= 8
	}
}

func RypReadInt(b *bytes.Buffer, n int) int64 {
	n++ // I was "minus 1"; now it will be accurate.
	var u uint64
	var shift uint
	for i := 0; i < n; i++ {
		a, err := b.ReadByte()
		if err != nil {
			panic(err)
		}
		u |= (uint64(a) << shift)
		shift += 8
	}
	//Say("RypReadInt", int64(u))
	return int64(u)
}

func RypWriteLabel(b *bytes.Buffer, s string) {
	n := len(s)
	if n > 255 {
		panic("RypWriteLabel: Cannot write label longer than 255 bytes")
	}
	b.WriteByte(byte(n & 255))
	b.WriteString(s)
}

func RypReadLabel(b *bytes.Buffer) string {
	n, err := b.ReadByte()
	if err != nil {
		panic(err)
	}

	bb := b.Next(int(n))
	if len(bb) != int(n) {
		panic("bytes.Buffer.Next: Short read")
	}

	//Say("RypReadLabel", string(bb))
	return string(bb)
}

func UnPickle(s string) P {
	//@ Say("UNPICKLE", len(s))
	RypLegend()
	return RypUnPickle(bytes.NewBufferString(s))
}

func RypUnPickle(b *bytes.Buffer) P {
	tag, err := b.ReadByte()
	if err != nil {
		panic(err)
	}

	kind := int(tag & RypMask)
	arg := int(tag & 7)

	//Say("RypUnPickle tag, kind, arg, n", tag, kind, arg, b.Len())

	switch kind {
	case RypNone:
		return None
	case RypTrue:
		return True
	case RypFalse:
		return False
	case RypInt:
		return MkInt(RypReadInt(b, arg))
	case RypFloat:
		return MkFloat(math.Float64frombits(uint64(RypReadInt(b, arg))))
	case RypStr:
		n := int(RypReadInt(b, arg))
		bb := b.Next(n)
		if len(bb) != n {
			panic(F("bytes.Buffer.Next: Short read, got %d, wanted %d", len(bb), n))
		}
		return MkStr(string(bb[:]))
	case RypByt:
		n := int(RypReadInt(b, arg))
		bb := b.Next(n)
		if len(bb) != n {
			panic("bytes.Buffer.Next: Short read")
		}
		return MkByt(bb)
	case RypTuple:
		n := int(RypReadInt(b, arg))
		pp := make([]P, n)
		for i := 0; i < n; i++ {
			pp[i] = RypUnPickle(b)
		}
		return MkTuple(pp)
	case RypList:
		n := int(RypReadInt(b, arg))
		pp := make([]P, n)
		for i := 0; i < n; i++ {
			pp[i] = RypUnPickle(b)
		}
		return MkList(pp)
	case RypDict:
		n := int(RypReadInt(b, arg))
		ppp := make(Scope)
		for i := 0; i < n; i++ {
			k := RypUnPickle(b)
			v := RypUnPickle(b)
			ppp[k.String()] = v
		}
		return MkDict(ppp)
	case RypClass:
		cname := RypReadLabel(b)
		cls, ok := Classes[cname]
		if !ok {
			panic(F("Class not found in Classes: %q", cname))
		}
		obj := R.New(cls)
		cobj := obj.Convert(R.TypeOf(VarOfStarP).Elem()).Interface().(P)
		cobj.SetSelf(obj.Interface().(P))
		for {
			fname := RypReadLabel(b)
			if len(fname) == 0 {
				break
			}
			x := RypUnPickle(b)
			RypSetField(obj.Elem(), fname, x)
		}
		return cobj
	}
	panic(F("Bad tag: %d", tag))
}

var VarOfStarP *P

func RypSetField(obj R.Value, fname string, x P) {
	t := obj.Type()
	nf := t.NumField()
	for i := 0; i < nf; i++ {
		f := t.Field(i)
		if f.Anonymous {
			if f.Type != PBaseType {
				RypSetField(obj.Field(i), fname, x)
			}
		} else {
			if f.Name == fname {
				obj.Field(i).Set(R.ValueOf(x))
			}
		}
	}
}

func HexDecode(a []byte) []byte {
	var z []byte
	_, err := hex.Decode(a, z)
	if err != nil {
		panic(err)
	}
	return z
}
func SafeIsNil(v R.Value) bool {
	switch v.Kind() {
	case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
		return v.IsNil()
	}
	return false
}

func PrintStackUnlessEOF(e interface{}) {
	s := fmt.Sprintf("%s", e)
	if s == "EOF" {
		return
	}
	PrintStack(e)
}
func PrintStack(e interface{}) {
	fmt.Fprintf(os.Stderr, "\n")
	Say("PrintStack:", e)
	debug.PrintStack()
	fmt.Fprintf(os.Stderr, "\n")
}

func FetchFieldByName(v R.Value, field string) P {
	// First try for method:
	meth := v.MethodByName(field)
	if meth.IsValid() {
		return MkValue(meth)
	}

	// Then try for field:
	v1 := MaybeDerefAll(v)
	v2 := MaybeDerefAll(v1)
	if v2.Kind() != R.Struct {
		panic(F("FetchFieldByName: Cannot get field %q from non-Struct %#v", field, v2))
	}
	x := v2.FieldByName(field)
	if !x.IsValid() {
		panic(F("FetchFieldByName: No such field %q on %T %#v", field, v2.Interface(), v2))
	}
	return AdaptForReturn(x)
}
func StoreFieldByName(v R.Value, field string, a P) {
	v = MaybeDeref(v) // Once for interface
	v = MaybeDeref(v) // Once for pointer
	if v.Kind() == R.Struct {
		vf := v.FieldByName(field)
		if vf.IsValid() {
			va := AdaptForCall(a, vf.Type())
			vf.Set(va)
			return
		}
		panic(F("StoreFieldByName: No such field %q on %T %#v", field, v.Interface(), v))
	}
	panic(F("StoreFieldByName: Cannot set field %q on non-Struct %#v", field, v))
}

func Sez(args ...interface{}) {
	Say(args...)
}

func GoReify(a P) P {
	switch t := a.(type) {
	case *PGo:
		v := t.V
		switch v.Kind() {
		case R.String:
			return MkStr(v.String())
		case R.Int:
			return MkInt(v.Int())
		case R.Int64:
			return MkInt(v.Int())
		case R.Float64:
			return MkFloat(v.Float())
		case R.Float32:
			return MkFloat(v.Float())
		}
		p, ok := v.Interface().(P)
		if ok {
			return p
		}
		panic(F("Cannot GoReify value of inner type %T", v.Interface()))
	}
	panic(F("Cannot GoReify P of type %T", a))
}

type PCallSpec struct {
	PBase
	Name     string
	Args     []string
	Defaults []P
	Star     string
	StarStar string
}

func (o *PCallSpec) String() string {
	return fmt.Sprintf("<func %s>", o.Name)
}
func (o *PCallSpec) Repr() string {
	return fmt.Sprintf(o.Name)
}

type KV struct {
	Key   string
	Value P
}

func SpecCall(cs *PCallSpec, a1 []P, a2 []P, kv []KV, kv2 map[string]P) ([]P, *PList, *PDict) {
	n := len(cs.Defaults)
	argv := make([]P, n)
	star := make([]P, 0)
	starstar := make(map[string]P)

	Say("### CCCCCC <<< name, args, defaults, *, ** :", cs.Name, cs.Args, cs.Defaults, cs.Star, cs.StarStar)
	Say("### CCCCCC <<<", a1, a2, kv, kv2)

	copy(argv, cs.Defaults) // Copy defaults first, any of which may be nil.

	j := 0
	for a1 != nil {
		for _, a := range a1 {
			if j < n {
				argv[j] = a
				Say("### argv [] = ", j, a)
				j++
			} else {
				star = append(star, a)
				Say("### star append ", a)
			}
		}
		a1 = a2
		a2 = nil
	}

	for _, e := range kv {
		k := e.Key
		v := e.Value
		Say("### kv1: k, v", k, v)
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				Say("### kv1: Stored argv[]", ni)
				break
			}
		}
		if !stored {
			starstar[k] = v
			Say("### kv1: starstar <- ", k, v)
		}
	}

	for k, v := range kv2 {
		Say("### kv2: k, v", k, v)
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				Say("### kv2: Stored argv[]", ni)
				break
			}
		}
		if !stored {
			starstar[k] = v
			Say("### kv2: starstar <- ", k, v)
		}
	}

	Say("### CCCCCC === argv, star, starstar:", argv, star, starstar)
	for i, e := range argv {
		if e == nil {
			panic(F("The %dth fixed argument has no assigned value", i))
		}
	}

	if cs.Star == "" && len(star) > 0 {
		panic(F("Function %q wants %d args, but got %d (no * arg)", cs.Name, len(cs.Args), len(cs.Args)+len(star)))
	}

	if cs.StarStar == "" && len(starstar) > 0 {
		panic(F("Function %q cannot take %d extra named args", cs.Name, len(starstar)))
	}

	Say("### CCCCCC >>> argv, star, starstar:", argv, star, starstar)
	return argv, MkList(star), MkDict(starstar)
}

type ICallV interface {
	CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P
}
