//

package rye

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"go/ast"
	"hash/crc64"
	"io"
	"math"
	"os"
	R "reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const SHOW_DEPTH = 6

var ONone = &PNone{}
var OTrue = &PBool{Z: true}
var OFalse = &PBool{Z: false}

var None M = MForge(ONone)
var True M = MForge(OTrue)
var False M = MForge(OFalse)
var EmptyStr M = M{X: Forge(&PStr{S: ""}).Self}

var G_rye_rye = True // Global var "rye_rye" is always True in Rye.
var Globals Scope = make(Scope)

var CounterMap = make(map[string]*int64)

func Shutdown() {
	if DebugCounters == 0 {
		return
	}










}

var RyeEnv string
var Debug int
var DebugCall int
var DebugExcept int
var DebugReflect int
var DebugGo int
var SkipAssert int
var DebugCounters int

func init() {
	RyeEnv := os.Getenv("RYE")
	for _, ch := range RyeEnv {
		switch ch {
		case 'a':
			SkipAssert++
		case 'c':
			DebugCall++
		case 'd':
			Debug++
		case 'e':
			DebugExcept++
		case 'r':
			DebugReflect++
		case 'i':
			DebugCounters++
		case 'g':
			DebugGo++
		}
	}
}

// P is the interface for every Pythonic value.
type P interface {
	Hash() int64
	Pickle(w *bytes.Buffer)
	Show() string
	String() string
	Repr() string
	PType() M
	RType() string // For recording.
	Callable() bool
	Is(a M) bool
	IsNot(a M) bool
	GetSelf() P
	SetSelf(a P)
	GetPBase() *PBase
	B() B
	M() M

	FetchField(field string) M
	StoreField(field string, p M)

	Call(aa ...M) M
	Invoke(field string, aa ...M) M
	Iter() Nexter
	List() []M
	Dict() Scope

	Len() int
	SetItem(i M, x M)
	DelItem(i M)
	DelItemSlice(i, j M)
	GetItem(a M) M
	GetItemSlice(a, b, c M) M
	Contains(a M) bool    // Reverse "in"
	NotContains(a M) bool // Reverse "not in"

	Add(a M) M
	Sub(a M) M
	Mul(a M) M
	Div(a M) M
	//IDiv(a M) M
	Mod(a M) M
	Pow(a M) M
	BitAnd(a M) M
	BitOr(a M) M
	BitXor(a M) M
	ShiftLeft(a M) M
	ShiftRight(a M) M
	UnsignedShiftRight(a M) M

	Bool() bool // a.k.a. nonzero()
	UnaryMinus() M
	UnaryPlus() M
	UnaryInvert() M

	EQ(a M) bool
	NE(a M) bool
	LT(a M) bool
	LE(a M) bool
	GT(a M) bool
	GE(a M) bool
	Compare(a M) int

	CanStr() bool
	Str() string
	// String() == ForceString()
	CanInt() bool
	Int() int64
	ForceInt() int64
	CanFloat() bool
	Float() float64
	ForceFloat() float64
	Contents() interface{}
	Bytes() []byte
	Object() *C_object
	Superclass() M
}

type B *PBase

func BoolToInt64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func BoolToFloat64(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func BoolToString(b bool) string {
	if b {
		return "True"
	} else {
		return "False"
	}
}

func Forge(p P) B {











	p.SetSelf(p)
	return p.B()
}

func MForge(p P) M {
	//println(fmt.Sprintf("MForge1: %T %#v", p, p))
	p.SetSelf(p)
	//println(fmt.Sprintf("MForge2: %T %#v", p, p))
	return p.M()
}

// Lock2 locks both mutexes, lowest address first, for deadlock avoidance.
func Lock2(x, y sync.Locker) {
	ax := R.ValueOf(x).Elem().UnsafeAddr()
	ay := R.ValueOf(y).Elem().UnsafeAddr()
	if ax == ay {
		x.Lock()
	} else if ax < ay {
		x.Lock()
		y.Lock()
	} else {
		y.Lock()
		x.Lock()
	}
}

// Unlock2 unlocks both mutexes, lowest address last, for deadlock avoidance.
func Unlock2(x, y sync.Locker) {
	ax := R.ValueOf(x).Elem().UnsafeAddr()
	ay := R.ValueOf(y).Elem().UnsafeAddr()
	if ax == ay {
		x.Unlock()
	} else if ax < ay {
		y.Unlock()
		x.Unlock()
	} else {
		x.Unlock()
		y.Unlock()
	}
}

func Pickle(p M) []byte {
	var b bytes.Buffer
	p.Pickle(&b)
	z := b.Bytes()
	return z
}

// C_object is the root of inherited classes.
type C_object struct {
	PBase
}

func (o *C_object) Object() *C_object {
	return o
}
func (o *C_object) PtrC_object() *C_object {
	return o
}

type I_GetSetAttr interface {
	M_1___getattr__(M) M
	M_2___setattr__(M, M) M
}

func (g *C_object) FetchField(field string) M {
	// Try using PGO reflection.
	return g.Self.(I_GetSetAttr).M_1___getattr__(MkStr(field))
}

func (g *C_object) StoreField(field string, p M) {
	// Try using PGO reflection.
	g.Self.(I_GetSetAttr).M_2___setattr__(MkStr(field), p)
}

func FetchFieldByNameForObject(v R.Value, field string) M {
	// First try for method:
	meth := v.MethodByName(field)
	if meth.IsValid() {
		return MkValue(meth)
	}

	_, okP := v.Interface().(P)
	if okP {
		// If it has a GET_<field> method, call it to get the field.
		meth = v.MethodByName("GET_" + field)
		if meth.IsValid() {
			fn := MkValue(meth).X.(*PGo)
			var zz []R.Value = fn.V.Call(nil)
			return zz[0].Interface().(M)
		}
	}

	// Then try for field:
	v2 := MaybeDerefTwice(v)
	if v2.Kind() != R.Struct {
		panic(F("FetchFieldByNameForObject: Cannot get field %q from non-Struct %#v", field, v2))
	}
	x := v2.FieldByName(field)
	if x.IsValid() {
		return AdaptForReturn(x)
	}
	if okP {
		x = v2.FieldByName("M_" + field)
		if x.IsValid() {
			return AdaptForReturn(x)
		}
	}
	panic(F("FetchFieldByNameForObject: No such field %q on %T %#v", field, v2.Interface(), v2))
}
func StoreFieldByNameForObject(v R.Value, field string, a M) {
	v = MaybeDeref(v) // Once for interface
	v = MaybeDeref(v) // Once for pointer
	if v.Kind() != R.Struct {
		panic(F("StoreFieldByNameForObject: Cannot set field %q on non-Struct %#v", field, v))
	}
	vf := v.FieldByName(field)
	if vf.IsValid() {
		va := AdaptForCall(a, vf.Type())
		vf.Set(va)
		return
	}
	vf = v.FieldByName("M_" + field)
	if vf.IsValid() {
		va := AdaptForCall(a, vf.Type())
		vf.Set(va)
		return
	}
	panic(F("StoreFieldByNameForObject: No such field %q on %T %#v", field, v.Interface(), v))
}

func MkPromise(fn func() M) M {
	z := &C_promise{Ch: make(chan Either, 1)}
	z.SetSelf(z)
	if DebugGo > 0 {
		println("#go# Made Promise: ", z)
	}
	go func() {
		var x M
		defer func() {
			r := recover()
			if r != nil {
				if DebugGo > 0 {
					println("#go# BAD Promise: ", z, r)
				}
				PrintStackFYI(r)
				z.Ch <- Either{Left: r, Valid: true, Right: MissingM}
			} else {
				if DebugGo > 0 {
					println("#go# OK Promise: ", z, x.String())
				}
				z.Ch <- Either{Left: nil, Valid: true, Right: x}
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
	return MkX(&z.PBase)
}

type C_promise struct {
	C_object
	Ch chan Either
}

func (o *C_promise) PtrC_promise() *C_promise {
	return o
}

func (o *C_promise) Wait() M {
	ch := o.Ch
	if ch == nil {
		panic("Wait() called more than once on promise")
	}
	x := <-ch
	o.Ch = nil // Don't allow another Wait.
	if x.Left != nil {
		panic(x.Left)
	}
	return x.Right
}

type C_rye_chan struct {
	C_object
	Chan    chan Either
	RevChan chan M
}

func (o *C_rye_chan) PtrC_chan() *C_rye_chan {
	return o
}

func make_rye_chan(size int64, revSize int64) M {
	z := new(C_rye_chan)
	z.Chan = make(chan Either, int(size))
	if revSize >= 0 {
		z.RevChan = make(chan M, int(revSize))
	}
	return MForge(z)
}

type Either struct {
	Left  interface{}
	Valid bool
	Right M
}

type void struct{}

// C_generator is the channel begin a yielding producer and a consuming for loop.
type C_generator struct {
	C_object
	Ready    chan *void
	Result   chan Either
	Finished bool
}

// GENERATOR_BUF_SIZE is the buffer chan size between Generator and consumer.
const GENERATOR_BUF_SIZE = 8

func NewGenerator() *C_generator {
	z := &C_generator{
		Ready:  make(chan *void, GENERATOR_BUF_SIZE),
		Result: make(chan Either, GENERATOR_BUF_SIZE),
	}
	z.SetSelf(z)
	// Signal the coroutine so it can run asynchronously.
	for i := 0; i < GENERATOR_BUF_SIZE; i++ {
		z.Ready <- nil
	}
	return z
}

func (o *C_generator) PtrC_generator() *C_generator {
	return o
}

func (o *C_generator) Iter() Nexter { return o }
func (o *C_generator) List() []M {
	var z []M
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
// It returns either a result of type M and true,
// or if there are no more, it returns nil and false.
// If the generator goroutine died on an exception,
// that exception gets rethrown here.
func (o *C_generator) Next() (M, bool) {
	o.Ready <- nil
	// That wakes up the generator goroutine.
	// Now we block, waiting on next result.
	either, ok := <-o.Result

	if !ok {
		return MissingM, false
	}
	if either.Left != nil {
		if o.Ready != nil {
			close(o.Ready)
			o.Ready = nil
		}
		panic(either.Left)
	}
	return either.Right, true
}

// Enough is called by the consumer, to tell the producer to stop because we've got enough.
func (o *C_generator) Enough() {
	if o.Ready != nil {
		close(o.Ready)
		o.Ready = nil
	}
}

// Yield is called by the producer, to yield a value to the consumer.
func (o *C_generator) Yield(item M) {
	o.Result <- Either{Left: nil, Valid: true, Right: item}
}

// Yield is called by the producer when it catches an exception, to yield it to the producer (as an Either Left).
func (o *C_generator) YieldException(ex interface{}) {
	o.Result <- Either{Left: ex, Valid: true, Right: MissingM}
}

// Finish is called by the producer when it is finished.
func (o *C_generator) Finish() {
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
	return ok
}

type GetSelfer interface {
	GetSelf() P
}

type PBase struct {
	Self P
}

func (o *PBase) B() B             { return o }
func (o *PBase) M() M             { return MkX(o) }
func (o *PBase) GetPBase() *PBase { return o }
func (o *PBase) GetSelf() P       { return o.Self }
func (o *PBase) SetSelf(a P) {
	o.Self = a
	//if Recording != nil {
	//	fmt.Fprintf(Recording, "Self\t%T\n", o.Self.Contents())
	//}
}

func (g *PBase) FetchField(field string) M {
	// Try using PGO reflection.
	return FetchFieldByNameForObject(R.ValueOf(g.Self), field)
}

func (o *PBase) StoreField(field string, p M) {
	panic(F("Receiver %T cannot StoreField", o.Self))
}

func (o *PBase) Call(aa ...M) M { panic(F("Receiver %T cannot Call", o.Self)) }
func (o *PBase) Invoke(field string, aa ...M) M {
	panic(F("Receiver %T cannot invoke", o.Self))
}
func (o *PBase) Len() int      { panic(F("Receiver %T cannot Len: ", o.Self)) }
func (o *PBase) GetItem(a M) M { panic(F("Receiver %T cannot GetItem", o.Self)) }
func (o *PBase) GetItemSlice(a, b, c M) M {
	panic(F("Receiver %T cannot GetItemSlice", o.Self))
}

func (o *PBase) Is(a M) bool    { return P(o) == a.X }
func (o *PBase) IsNot(a M) bool { return P(o) != a.X }
func (o *PInt) Is(a M) bool     { return o.N == a.N }
func (o *PInt) IsNot(a M) bool  { return o.N != a.N }
func (o *PStr) Is(a M) bool     { return o.S == a.S }
func (o *PStr) IsNot(a M) bool  { return o.S != a.S }

func (o *PBase) Contains(a M) bool    { panic(F("Receiver %T cannot Contains: ", o.Self)) }
func (o *PBase) NotContains(a M) bool { panic(F("Receiver %T cannot NotContains: ", o.Self)) }
func (o *PBase) SetItem(i M, x M)     { panic(F("Receiver %T cannot SetItem: ", o.Self)) }
func (o *PBase) DelItem(i M)          { panic(F("Receiver %T cannot DelItem: ", o.Self)) }
func (o *PBase) DelItemSlice(i, j M)  { panic(F("Receiver %T cannot DelItemSlice: ", o.Self)) }
func (o *PBase) Iter() Nexter         { panic(F("Receiver %T cannot Iter: ", o.Self)) }
func (o *PBase) List() []M            { panic(F("Receiver %T cannot List: ", o.Self)) }
func (o *PBase) Dict() Scope          { panic(F("Receiver %T cannot Dict: ", o.Self)) }

func (o *PBase) Add(a M) M { panic(F("Receiver %T cannot Add: ", o.Self)) }
func (o *PBase) Sub(a M) M { panic(F("Receiver %T cannot Sub: ", o.Self)) }
func (o *PBase) Mul(a M) M { panic(F("Receiver %T cannot Mul: ", o.Self)) }
func (o *PBase) Div(a M) M { panic(F("Receiver %T cannot Div: ", o.Self)) }

//func (o *PBase) IDiv(a M) M       { panic(F("Receiver %T cannot IDiv: ", o.Self)) }
func (o *PBase) Mod(a M) M        { panic(F("Receiver %T cannot Mod: ", o.Self)) }
func (o *PBase) Pow(a M) M        { panic(F("Receiver %T cannot Pow: ", o.Self)) }
func (o *PBase) BitAnd(a M) M     { panic(F("Receiver %T cannot BitAnd: ", o.Self)) }
func (o *PBase) BitOr(a M) M      { panic(F("Receiver %T cannot BitOr: ", o.Self)) }
func (o *PBase) BitXor(a M) M     { panic(F("Receiver %T cannot BitXor: ", o.Self)) }
func (o *PBase) ShiftLeft(a M) M  { panic(F("Receiver %T cannot ShiftLeft: ", o.Self)) }
func (o *PBase) ShiftRight(a M) M { panic(F("Receiver %T cannot ShiftRight: ", o.Self)) }
func (o *PBase) UnsignedShiftRight(a M) M {
	panic(F("Receiver %T cannot UnsignedShiftRight: ", o.Self))
}

func (o *PBase) EQ(a M) bool { return o.Self.Compare(a) == 0 }
func (o *PBase) NE(a M) bool { return o.Self.Compare(a) != 0 }
func (o *PBase) LT(a M) bool { return o.Self.Compare(a) < 0 }
func (o *PBase) LE(a M) bool { return o.Self.Compare(a) <= 0 }
func (o *PBase) GT(a M) bool { return o.Self.Compare(a) > 0 }
func (o *PBase) GE(a M) bool { return o.Self.Compare(a) >= 0 }
func (o *PBase) Compare(a M) int {
	// Default comparision uses address in memory.
	if a.X != nil {
		// cmp ... ...
		x := R.ValueOf(o).Pointer()
		y := R.ValueOf(a.X.B()).Pointer()
		switch {
		case x < y:
			return -1
		case x > y:
			return 1
		}
		return 0
	} else if len(a.S) == 0 {
		// cmp ... str
		return StrCmp(o.PType().String(), a.PType().String())
	} else {
		// cmp ... int
		return StrCmp(o.PType().String(), a.PType().String())
	}
}
func (o *PBase) Bool() bool     { return true } // Most things are true.
func (o *PBase) UnaryMinus() M  { panic(F("Receiver %T cannot UnaryMinus", o.Self)) }
func (o *PBase) UnaryPlus() M   { panic(F("Receiver %T cannot UnaryPlus", o.Self)) }
func (o *PBase) UnaryInvert() M { panic(F("Receiver %T cannot UnaryInvert", o.Self)) }

func (o *PBase) CanStr() bool { return false }
func (o *PBase) Str() string  { panic(F("Receiver %T cannot Str", o.Self)) }

func (o *PBase) CanInt() bool    { return false }
func (o *PBase) Int() int64      { panic(F("Receiver %T cannot Int", o.Self)) }
func (o *PBase) ForceInt() int64 { panic(F("Receiver %T cannot ForceInt", o.Self)) }

func (o *PBase) CanFloat() bool        { return false }
func (o *PBase) Float() float64        { panic(F("Receiver %T cannot Float", o.Self)) }
func (o *PBase) ForceFloat() float64   { panic(F("Receiver %T cannot ForceFloat", o.Self)) }
func (o *PBase) Contents() interface{} { return o.Self }

func (o *PBase) Superclass() M     { return None }
func (o *PBase) Object() *C_object { return nil }
func (o *PBase) Callable() bool    { return false }
func (o *PBase) PType() M          { return MkStr(F("%T", o.Self)) }
func (o *PBase) RType() string     { return "" }
func (o *PBase) Bytes() []byte     { panic(F("Receiver %T cannot Bytes", o.Self)) }
func (o *PBase) String() string {
	if o.Self == nil {
		panic("PBase: Self is nil")
	}
	return o.Self.Show()
}
func (o *PBase) Hash() int64 { return int64(R.ValueOf(o).Pointer()) }

const ShortHashModulus = 99999 // 3 * 3 * 41 * 271
func (o *PBase) ShortPointerHashString() string {
	return F("@%05d", o.ShortPointerHash())
}
func (o *PBase) ShortPointerHash() int {
	val := R.ValueOf(o.Self)
	return 1 + int(val.Pointer())%ShortHashModulus
}

func (o *PBase) Pickle(w *bytes.Buffer) { panic(F("Receiver %T cannot Pickle", o.Self)) }
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
	Z bool
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
	PP []M
}

type PListIter struct {
	PBase
	PP []M
	I  int
}

type PTuple struct {
	PBase
	PP []M
}

type PNone struct {
	PBase
}

type Scope map[string]M

type PDict struct {
	PBase
	ppp Scope
	mu  sync.Mutex
}

type PSet struct {
	PBase
	ppp Scope
	mu  sync.Mutex
}

func MkRecovered(a interface{}) M {
	switch x := a.(type) {
	case string:
		return MkStr(x)
	case []byte:
		return MkByt(x)
	case M:
		return x
	case P:
		return x.M()
	}
	return MkGo(a)
}

//type Ber interface {
//  B()
//}

func Mk(a interface{}) M {
	switch t := a.(type) {
	case M:
		return t
	case B:
		return MkX(t)
	case P:
		return MkX(t.B())
	case string:
		return MkStr(t)
	case int:
		return Mkint(t)
	case int64:
		return MkInt(t)
	case float64:
		return MkFloat(t)
	}
	panic(fmt.Sprintf("Cannot call Mk() on a %T : %#v", a, a))
}

var counterMkGo int64

func MkGo(a interface{}) M {
	counterMkGo++
	z := &PGo{V: R.ValueOf(a)}
	return MForge(z)
}

var counterMkValue int64

func MkValue(a R.Value) M {
	counterMkValue++
	z := &PGo{V: a}
	return MForge(z)
}

var counterMkint int64

func Mkint(n int) M {
	counterMkint++
	return M{N: int64(n)}
}

var counterMkInt int64

func MkInt(n int64) M {
	counterMkInt++
	return M{N: n}
}

func init() {













}

var counterMkBInt int64

func MkBInt(n int64) B {
	counterMkBInt++
	z := &PInt{N: n}
	return Forge(z)
}

var counterMkFloat int64

func MkFloat(f float64) M {
	counterMkFloat++
	z := &PFloat{F: f}
	return MForge(z)
}

var counterMkBStr int64

func MkBStr(s string) B {
	counterMkBStr++
	z := &PStr{S: s}
	return Forge(z)
}

var counterMkStr int64

func MkStr(s string) M {
	counterMkStr++
	if len(s) == 0 {
		return EmptyStr
	}
	return M{S: s}
	//z := &PStr{S: s}
	//return MForge(z)
}

func MkByt(yy []byte) M { z := &PByt{YY: yy}; return MForge(z) }
func MkStrs(ss []string) M {
	pp := make([]M, len(ss))
	for i, s := range ss {
		pp[i] = MkStr(s)
	}
	return MkList(pp)
}
func MkByts(ss [][]byte) M {
	pp := make([]M, len(ss))
	for i, s := range ss {
		pp[i] = MkByt(s)
	}
	return MkList(pp)
}

var counterMkList int64

func MkList(pp []M) M {
	counterMkList++
	z := &PList{PP: pp}
	return MForge(z)
}

var counterMkTuple int64

func MkTuple(pp []M) M {
	counterMkTuple++
	z := &PTuple{PP: pp}
	return MForge(z)
}

var counterMkDict int64

func MkDict(ppp Scope) M {
	counterMkDict++
	z := &PDict{ppp: ppp}
	return MForge(z)
}
func MkSet(ppp Scope) M { z := &PSet{ppp: ppp}; return MForge(z) }

func PMkList(pp []M) *PList    { z := &PList{PP: pp}; Forge(z); return z }
func PMkDict(ppp Scope) *PDict { z := &PDict{ppp: ppp}; Forge(z); return z }

func MkDictCopy(ppp Scope) M {
	z := &PDict{ppp: make(Scope)}
	for k, v := range ppp {
		z.ppp[k] = v
	}
	return MForge(z)
}

func MkDictFromPairs(pp []M) M {
	z := &PDict{ppp: make(Scope)}
	for _, x := range pp {
		sub := x.List()
		if len(sub) != 2 {
			panic(F("MkDictFromPairs: got sublist of size %d, wanted size 2", len(sub)))
		}
		k := sub[0].String()
		v := sub[1]
		z.ppp[k] = v
	}
	return MForge(z)
}

func MkListV(pp ...M) M  { z := &PList{PP: pp}; return MForge(z) }
func MkTupleV(pp ...M) M { z := &PTuple{PP: pp}; return MForge(z) }
func MkDictV(pp ...M) M {
	if (len(pp) % 2) == 1 {
		panic("MkDictV got odd len(pp)")
	}
	zzz := make(Scope)
	for i := 0; i < len(pp); i += 2 {
		zzz[pp[i].String()] = pp[i+1]
	}
	z := &PDict{ppp: zzz}
	return MForge(z)
}

func MkSetV(pp ...M) M {
	zzz := make(Scope)
	for i := 0; i < len(pp); i++ {
		zzz[pp[i].String()] = True
	}
	z := &PSet{ppp: zzz}
	return MForge(z)
}

func MkNone() M { return None }
func MkBool(b bool) M {
	if b {
		return True
	} else {
		return False
	}
}
func ListToStrings(a []M) []string {
	n := len(a)
	z := make([]string, n, n)
	for i, e := range a {
		z[i] = e.Str()
	}
	return z
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

// Because go confuses empty lists with nil, rye does the same with None.
// This saves you writing `for x in vec if vec else []:`
func (o *PNone) Hash() int64          { return 23 }
func (o *PNone) Len() int             { return 0 }
func (o *PNone) Contains(a M) bool    { return false }
func (o *PNone) NotContains(a M) bool { return true }
func (o *PNone) List() []M            { return nil }
func (o *PNone) Dict() Scope          { return make(Scope) }
func (o *PNone) Iter() Nexter {
	z := &PListIter{PP: nil}
	Forge(z)
	return z
}

func (o *PNone) Bool() bool            { return false }
func (o *PNone) String() string        { return "None" }
func (o *PNone) Repr() string          { return "None" }
func (o *PNone) Contents() interface{} { return nil }

func (o *PNone) Pickle(w *bytes.Buffer) { w.WriteByte(RypNone) }

func (o *PBool) Hash() int64 { return o.Int() }
func (o *PBool) Pickle(w *bytes.Buffer) {
	if o.Z {
		w.WriteByte(RypTrue)
	} else {
		w.WriteByte(RypFalse)
	}
}
func (o *PBool) Contents() interface{} { return o.Z }
func (o *PBool) Bool() bool            { return o.Z }
func (o *PBool) CanInt() bool          { return true }
func (o *PBool) ForceInt() int64       { return o.Int() }
func (o *PBool) Int() int64 {
	if o.Z {
		return 1
	} else {
		return 0
	}
}
func (o *PBool) CanFloat() bool      { return true }
func (o *PBool) ForceFloat() float64 { return o.Float() }
func (o *PBool) Float() float64 {
	if o.Z {
		return 1.0
	} else {
		return 0.0
	}
}
func (o *PBool) String() string {
	if o.Z {
		return "True"
	} else {
		return "False"
	}
}
func (o *PBool) Repr() string {
	if o.Z {
		return "True"
	} else {
		return "False"
	}
}
func (o *PBool) PType() M { return G_bool }
func (o *PBool) Compare(a M) int {
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
	return StrCmp(o.PType().String(), a.PType().String())
}

// Uncommon but useful arithmetic on bools doesn't need to be efficient.
func (o *PBool) Add(b M) M { return MkInt(o.Int()).Add(b) }
func (o *PBool) Sub(b M) M { return MkInt(o.Int()).Sub(b) }
func (o *PBool) Mul(b M) M { return MkInt(o.Int()).Mul(b) }
func (o *PBool) Div(b M) M { return MkInt(o.Int()).Div(b) }

//func (o *PBool) IDiv(b M) M { return MkInt(o.Int()).IDiv(b) }
func (o *PBool) Mod(b M) M { return MkInt(o.Int()).Mod(b) }

func (o *PInt) Hash() int64    { return o.Int() }
func (o *PInt) UnaryMinus() M  { return MkInt(0 - o.N) }
func (o *PInt) UnaryPlus() M   { return MkInt(o.N) }
func (o *PInt) UnaryInvert() M { return MkInt(int64(-1) ^ o.N) }
func (o *PInt) Add(a M) M {
	switch x := a.X.(type) {
	case *PFloat:
		return MkFloat(float64(o.N) + x.F)
	}
	return MkInt(o.N + a.Int())
}
func (o *PInt) Sub(a M) M {
	switch x := a.X.(type) {
	case *PFloat:
		return MkFloat(float64(o.N) - x.F)
	}
	return MkInt(o.N - a.Int())
}

func RepeatList(a []M, n int64) M {
	var z []M
	for i := int64(0); i < n; i++ {
		z = append(z, a...)
	}
	return MkList(z)
}
func RepeatTuple(a []M, n int64) M {
	var z []M
	for i := int64(0); i < n; i++ {
		z = append(z, a...)
	}
	return MkTuple(z)
}
func (o *PInt) Mul(a M) M {
	switch x := a.X.(type) {
	case *PInt:
		return MkInt(o.N * x.N)
	case *PFloat:
		return MkFloat(float64(o.N) * x.F)
	case *PStr:
		return MkStr(strings.Repeat(x.S, int(o.N)))
	case *PByt:
		sz := len(x.YY)
		n := int(o.N)
		z := make([]byte, n*sz)
		for i := 0; i < n; i++ {
			copy(z[i*sz:(i+1)*sz], x.YY)
		}
		return MkByt(z)
	case *PList:
		return RepeatList(x.PP, o.N)
	case *PTuple:
		return RepeatTuple(x.PP, o.N)
	case *PGo:
		bt := x.V.Type()
		switch bt.Kind() {
		case R.Int64, R.Int:
			return MkInt(o.N * x.V.Int())
		}
	}
	panic("Cannot multply int times whatever")
}
func (o *PInt) Div(a M) M {
	switch x := a.X.(type) {
	case *PFloat:
		return MkFloat(float64(o.N) / x.F)
	}
	return MkInt(o.N / a.Int())
}

//func (o *PInt) IDiv(a M) M {
//	switch x := a.Self.(type) {
//	case *PFloat:
//		return MkFloat(math.Floor(float64(o.N) / x.F))
//	}
//	return MkInt(o.N / a.Self.Int())
//}
func (o *PInt) Mod(a M) M {
	switch x := a.X.(type) {
	case *PFloat:
		_ = x
		panic("golang cannot mod with a float")
	}
	return MkInt(o.N % a.Int())
}
func (o *PInt) BitAnd(a M) M             { return MkInt(o.N & a.Int()) }
func (o *PInt) BitOr(a M) M              { return MkInt(o.N | a.Int()) }
func (o *PInt) BitXor(a M) M             { return MkInt(o.N ^ a.Int()) }
func (o *PInt) ShiftLeft(a M) M          { return MkInt(o.N << uint64(a.Int())) }
func (o *PInt) ShiftRight(a M) M         { return MkInt(o.N >> uint64(a.Int())) }
func (o *PInt) UnsignedShiftRight(a M) M { return MkInt(int64(uint64(o.N) >> uint64(a.Int()))) }
func (o *PInt) Compare(a M) int {
	switch b := a.X.(type) {
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
	return StrCmp(o.PType().String(), a.PType().String())
}
func (o *PInt) CanInt() bool          { return true }
func (o *PInt) Int() int64            { return o.N }
func (o *PInt) ForceInt() int64       { return o.N }
func (o *PInt) CanFloat() bool        { return true }
func (o *PInt) Float() float64        { return float64(o.N) }
func (o *PInt) ForceFloat() float64   { return float64(o.N) }
func (o *PInt) String() string        { return strconv.FormatInt(o.N, 10) }
func (o *PInt) Repr() string          { return o.String() }
func (o *PInt) Bool() bool            { return o.N != 0 }
func (o *PInt) PType() M              { return G_int }
func (o *PInt) RType() string         { return "int" }
func (o *PInt) Contents() interface{} { return o.N }
func (o *PInt) Pickle(w *bytes.Buffer) {
	n := RypIntLenMinus1(o.N)
	w.WriteByte(byte(RypInt + n))
	RypWriteInt(w, o.N)
}

func (o *PFloat) Hash() int64   { return int64(o.F) ^ int64(1000000000000000*o.F) } // TODO better.
func (o *PFloat) UnaryMinus() M { return MkFloat(0.0 - o.F) }
func (o *PFloat) UnaryPlus() M  { return MkX(&o.PBase) }
func (o *PFloat) Add(a M) M     { return MkFloat(o.F + a.Float()) }
func (o *PFloat) Sub(a M) M     { return MkFloat(o.F - a.Float()) }
func (o *PFloat) Mul(a M) M     { return MkFloat(o.F * a.Float()) }
func (o *PFloat) Div(a M) M     { return MkFloat(o.F / a.Float()) }

//func (o *PFloat) IDiv(a M) M    { return MkInt(int64(o.F / a.Float())) }
func (o *PFloat) Compare(a M) int {
	c := a.Float()
	switch {
	case o.F < c:
		return -1
	case o.F > c:
		return 1
	case o.F == c:
		return 0
	}
	return StrCmp(o.PType().String(), a.PType().String())
}
func (o *PFloat) ForceInt() int64       { return int64(o.F) }
func (o *PFloat) CanFloat() bool        { return true }
func (o *PFloat) Float() float64        { return o.F }
func (o *PFloat) ForceFloat() float64   { return o.F }
func (o *PFloat) String() string        { return strconv.FormatFloat(o.F, 'g', -1, 64) }
func (o *PFloat) Repr() string          { return o.String() }
func (o *PFloat) Bool() bool            { return o.F != 0 }
func (o *PFloat) PType() M              { return G_float }
func (o *PFloat) RType() string         { return "float" }
func (o *PFloat) Contents() interface{} { return o.F }
func (o *PFloat) Pickle(w *bytes.Buffer) {
	x := int64(math.Float64bits(o.F))
	n := RypIntLenMinus1(x)
	w.WriteByte(byte(RypFloat + n))
	RypWriteInt(w, x)
}

var CrcPolynomial *crc64.Table = crc64.MakeTable(crc64.ECMA)

func (o *PStr) Hash() int64 { return int64(crc64.Checksum([]byte(o.S), CrcPolynomial)) }
func (o *PStr) Iter() Nexter {
	var pp []M
	for _, r := range o.S {
		pp = append(pp, MkStr(string(r)))
	}
	z := &PListIter{PP: pp}
	Forge(z)
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
func (o *PStr) GetItem(x M) M {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.S))
	}
	return MkStr(o.S[i : i+1])
}

func (o *PStr) GetItemSlice(x, y, z M) M {
	var i, j int64
	n := int64(len(o.S))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += n
		}
		if i < 0 {
			panic(F("First slicing index on PStr too small: %d", i))
		}
	}
	if i > n {
		panic(F("First slicing index on PStr too large: %d > len: %d", i, n))
	}
	if y == None {
		j = n
	} else {
		j = y.Int()
		if j < 0 {
			j += n
		}
		if j < 0 {
			panic(F("Second slicing index on PStr too small: %d", y.Int()))
		}
	}
	if j > n {
		j = n // Python lets you specify too big second index.
		// panic(F("Second slicing index on PStr too large: %d > len: %d", j, n))
	}
	return MkStr(o.S[i:j])
}

func (o *PStr) Mod(a M) M {
	switch t := a.X.(type) {
	case *PTuple:
		z := make([]interface{}, len(t.PP))
		for i, e := range t.PP {
			z[i] = e.Contents()
		}
		return MkStr(F(o.S, z...))
	}
	return MkStr(F(o.S, a.Contents()))
}

func (o *PStr) Mul(a M) M {
	switch t := a.X.(type) {
	case nil:
		if len(a.S) == 0 {
			return MkStr(strings.Repeat(o.S, int(a.N)))
		}
	case *PInt:
		return MkStr(strings.Repeat(o.S, int(t.N)))
	}
	panic(F("Cannot multiply: str * %s", a.PType()))
}
func (o *PStr) NotContains(a M) bool { return !o.Contains(a) }
func (o *PStr) Contains(a M) bool {
	switch t := a.X.(type) {
	case nil:
		if len(a.S) > 0 {
			return strings.Contains(o.S, a.S)
		}
	case *PStr:
		return strings.Contains(o.S, t.S)
	}
	panic(F("str cannot Contain non-str: %s", a.PType()))
}
func (o *PStr) Add(a M) M { return MkStr(o.S + a.String()) }
func (o *PStr) Compare(a M) int {
	if a.X == nil {
		if len(a.S) > 0 {
			return StrCmp(o.S, a.S)
		}
	}
	switch b := a.X.(type) {
	case *PStr:
		switch {
		case o.S < b.S:
			return -1
		case o.S > b.S:
			return 1
		}
		return 0
	}
	return StrCmp(o.PType().String(), a.PType().String())
}
func (o *PStr) ForceInt() int64 {
	z, err := strconv.ParseInt(o.S, 10, 64)
	if err != nil {
		panic(F("PStr::ForceInt: ParseInt: %v", err))
	}
	return z
}
func (o *PStr) ForceFloat() float64 {
	z, err := strconv.ParseFloat(o.S, 64)
	if err != nil {
		panic(F("PStr::ForceFloat: ParseFloat: %v", err))
	}
	return z
}
func (o *PStr) String() string { return o.S }
func (o *PStr) Bytes() []byte  { return []byte(o.S) }
func (o *PStr) Len() int       { return len(o.S) }

//func (o *PStr) Repr() string   { return F("%q", o.S) }
func (o *PStr) Repr() string  { return ReprStringLikeInPython(o.S) }
func (o *PStr) PType() M      { return G_str }
func (o *PStr) RType() string { return "str" }

var hexTABLE []byte = []byte("0123456789abcdef")

func ReprStringLikeInPython(s string) string {
	var bb bytes.Buffer
	bb.WriteByte('\'')
	n := len(s)
	for i := 0; i < n; i++ {
		var c byte = s[i]
		if ' ' <= c && c <= '~' && c != '\'' && c != '\\' {
			bb.WriteByte(c)
		} else {
			bb.WriteByte('\\')
			bb.WriteByte('x')
			bb.WriteByte(hexTABLE[(c>>4)&0xF])
			bb.WriteByte(hexTABLE[(c>>0)&0xF])
		}
	}
	bb.WriteByte('\'')
	return bb.String()
}

func (o *PStr) CanStr() bool { return true }
func (o *PStr) Str() string  { return o.S }

func (o *PByt) Hash() int64  { return int64(crc64.Checksum(o.YY, CrcPolynomial)) }
func (o *PByt) CanStr() bool { return true }
func (o *PByt) Str() string  { return string(o.YY) }

func (o *PByt) Iter() Nexter {
	var pp []M
	for _, r := range o.YY {
		pp = append(pp, Mkint(int(r)))
	}
	z := &PListIter{PP: pp}
	Forge(z)
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
func (o *PByt) GetItem(a M) M {
	i := int(a.Int())
	if i < 0 {
		i += len(o.YY)
	}
	return Mkint(int(o.YY[i]))
}
func (o *PByt) SetItem(a M, x M) {
	i := int(a.Int())
	if i < 0 {
		i += len(o.YY)
	}
	o.YY[i] = byte(x.Int())
}

func (o *PByt) GetItemSlice(x, y, z M) M {
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
		j = n // Python lets you specify too big second index.
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkByt(o.YY[i:j])
	return r
}

func (o *PByt) Mul(a M) M {
	switch t := a.X.(type) {
	case *PInt:
		return MkByt(bytes.Repeat(o.YY, int(t.N)))
	}
	panic(F("Cannot multiply: byt * %s", a.PType()))
}

func (o *PByt) BitAnd(a M) M {
	n := len(o.YY)
	b := a.Bytes()
	if n != len(b) {
		panic(F("BitAnd on byt of different sizes: %d vs %d", n, len(b)))
	}
	z := make([]byte, n)
	for i, e := range o.YY {
		z[i] = e & b[i]
	}
	return MkByt(z)
}
func (o *PByt) BitOr(a M) M {
	n := len(o.YY)
	b := a.Bytes()
	if n != len(b) {
		panic(F("BitOr on byt of different sizes: %d vs %d", n, len(b)))
	}
	z := make([]byte, n)
	for i, e := range o.YY {
		z[i] = e | b[i]
	}
	return MkByt(z)
}
func (o *PByt) BitXor(a M) M {
	n := len(o.YY)
	b := a.Bytes()
	if n != len(b) {
		panic(F("BitXor on byt of different sizes: %d vs %d", n, len(b)))
	}
	z := make([]byte, n)
	for i, e := range o.YY {
		z[i] = e ^ b[i]
	}
	return MkByt(z)
}

func (o *PByt) NotContains(a M) bool { return !o.Contains(a) }
func (o *PByt) Contains(a M) bool {
	switch t := a.X.(type) {
	case *PByt:
		return bytes.Contains(o.YY, t.YY)
	case *PInt:
		n := t.N
		for _, e := range o.YY {
			if int64(e) == n {
				return true
			}
		}
		return false
	}
	panic(F("Byt cannot Contain %s", a.PType()))
}
func (o *PByt) Add(a M) M {
	aa := a.Bytes()
	var zz []byte
	zz = append(zz, o.YY...)
	zz = append(zz, aa...)
	return MkByt(zz)
}

func (o *PByt) String() string { return string(o.YY) }
func (o *PByt) Show() string   { return o.Repr() }
func (o *PByt) Bytes() []byte  { return o.YY }
func (o *PByt) Len() int       { return len(o.YY) }
func (o *PByt) Repr() string   { return F("byt(%s)", ReprStringLikeInPython(string(o.YY))) }
func (o *PByt) PType() M       { return G_byt }
func (o *PByt) RType() string  { return "byt" }
func (o *PByt) List() []M {
	zz := make([]M, len(o.YY))
	for i, x := range o.YY {
		zz[i] = Mkint(int(x))
	}
	return zz
}
func (o *PByt) Compare(a M) int {
	switch b := a.X.(type) {
	case *PByt:
		switch {
		case string(o.YY) < string(b.YY):
			return -1
		case string(o.YY) > string(b.YY):
			return 1
		}
		return 0
	}
	return StrCmp(o.PType().String(), a.PType().String())
}

func (o *PTuple) Hash() int64 {
	var z int64
	for _, e := range o.PP {
		z += e.Hash() // TODO better
	}
	return z
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
func (o *PTuple) NotContains(a M) bool  { return !o.Contains(a) }
func (o *PTuple) Contains(a M) bool {
	for _, x := range o.PP {
		if a.EQ(x) {
			return true
		}
	}
	return false
}
func (o *PTuple) Len() int { return len(o.PP) }
func (o *PTuple) GetItem(x M) M {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.PP))
	}
	return o.PP[i]
}
func (o *PTuple) GetItemSlice(x, y, z M) M {
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
		j = n
		// panic(F("Second slicing index on PTuple too large: %d > len: %d", j, n))
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkList(o.PP[i:j])
	return r
}
func (o *PTuple) String() string { return o.Repr() }
func (o *PTuple) PType() M       { return G_tuple }

//func (o *PTuple) RType() string         { return "tuple" }
func (o *PTuple) RType() string {
	buf := bytes.NewBufferString("tuple(")
	n := len(o.PP)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(o.PP[i].RType())
	}
	buf.WriteString(")")
	return buf.String()
}

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
	Forge(z)
	return z
}
func (o *PTuple) List() []M {
	return o.PP
}

func (o *PTuple) Compare(a M) int {
	switch b := a.X.(type) {
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
	return StrCmp(o.PType().String(), a.PType().String())
}

func (o *PList) Hash() int64 {
	var z int64
	for _, e := range o.PP {
		z += e.Hash() // TODO better
	}
	return z
}
func (o *PList) Compare(a M) int {
	switch b := a.X.(type) {
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
	return StrCmp(o.PType().String(), a.PType().String())
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
func (o *PList) Add(a M) M {
	b := a.List()
	z := make([]M, 0, len(o.PP)+len(b))
	z = append(z, o.PP...)
	z = append(z, b...)
	return MkList(z)
}
func (o *PList) Contents() interface{} { return o.PP }
func (o *PList) Bool() bool            { return len(o.PP) != 0 }
func (o *PList) NotContains(a M) bool  { return !o.Contains(a) }
func (o *PList) Contains(a M) bool {
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
func (o *PList) SetItem(a M, x M) {
	if !a.CanInt() {
		panic("index to PList::SetItem should be an integer")
	}
	i := int(a.Int())
	if i < 0 {
		i += len(o.PP)
	}
	o.PP[i] = x
}

func (o *PList) GetItem(x M) M {
	i := x.Int()
	if i < 0 {
		i += int64(len(o.PP))
	}
	return o.PP[i]
}
func (o *PList) GetItemSlice(x, y, z M) M {
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
		j = n
		// panic(F("Second slicing index on PList too large: %d > len: %d", j, n))
	}
	// TODO: Step by z.
	if z != None {
		panic("GetItemSlice: step not imp")
	}
	r := MkList(o.PP[i:j])
	return r
}

func (o *PList) String() string { return o.Repr() }
func (o *PList) PType() M       { return G_list }
func (o *PList) RType() string  { return "list" }
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
	Forge(z)
	return z
}
func (o *PList) List() []M {
	return o.PP
}

func (o *PList) DelItem(x M) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	i := int(x.Int())
	n := len(a)
	if n == 0 {
		panic("cannot del item in empty list")
	}
	if i < 0 {
		i += n
	}
	if i < n-1 {
		copy(a[i:], a[i+1:])
	}
	o.PP = a[:n-1]
}
func (o *PList) DelItemSlice(x, y M) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	i, j := int(x.Int()), int(y.Int())
	copy(a[i:], a[j:])
	//? for k, n := len(a)-j+i, len(a); k < n; k++ {
	//? 	a[k] = nil // or the zero value of T
	//? } // for k
	o.PP = a[:len(a)-j+i]
}

func (o *PListIter) Iter() Nexter {
	return o
}

type Nexter interface {
	Next() (M, bool)
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

func (o *PListIter) Next() (M, bool) {
	if o.I < len(o.PP) {
		z := o.PP[o.I]
		o.I++
		return z, true
	}
	return MissingM, false
}

// ========== dict

func (o *PDict) Hash() int64 {
	var z int64
	o.mu.Lock()
	for k, v := range o.ppp {
		z += int64(crc64.Checksum([]byte(k), CrcPolynomial))
		z += v.Hash() // TODO better
	}
	o.mu.Unlock()
	return z
}
func (o *PDict) Pickle(w *bytes.Buffer) {
	o.mu.Lock()
	defer o.mu.Unlock()
	l := int64(len(o.ppp))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypDict + n))
	RypWriteInt(w, l)
	for k, v := range o.ppp {
		MkStr(k).Pickle(w)
		v.Pickle(w)
	}
}
func (o *PDict) Contents() interface{} { return o.ppp }
func (o *PDict) Bool() bool            { return len(o.ppp) != 0 }
func (o *PDict) NotContains(a M) bool  { return !o.Contains(a) }
func (o *PDict) Contains(a M) bool {
	key := a.String()
	o.mu.Lock()
	_, ok := o.ppp[key]
	o.mu.Unlock()
	return ok
}
func (o *PDict) Len() int { return len(o.ppp) }
func (o *PDict) SetItem(a M, x M) {
	key := a.String()
	o.mu.Lock()
	o.ppp[key] = x
	o.mu.Unlock()
}
func (o *PDict) GetItem(a M) M {
	key := a.String()
	o.mu.Lock()
	z, ok := o.ppp[key]
	o.mu.Unlock()
	if !ok {
		panic(F("PDict: KeyError: %q", key))
	}
	return z
}
func (o *PDict) String() string { return o.Repr() }
func (o *PDict) PType() M       { return G_dict }
func (o *PDict) RType() string  { return "dict" }
func (o *PDict) Repr() string {
	o.mu.Lock()
	vec := make(KVSlice, 0, len(o.ppp))
	for k, v := range o.ppp {
		vec = append(vec, KV{k, v})
	}
	o.mu.Unlock()

	sort.Sort(vec)
	buf := bytes.NewBufferString("{")
	n := len(vec)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(F("%q: %s", vec[i].Key, vec[i].Value.Repr()))
	}
	buf.WriteString("}")
	return buf.String()
}
func (o *PDict) Enough() {}
func (o *PDict) Iter() Nexter {
	var keys []M
	o.mu.Lock()
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	o.mu.Unlock()
	z := &PListIter{PP: keys}
	Forge(z)
	return z
}
func (o *PDict) List() []M {
	var keys []M
	o.mu.Lock()
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	o.mu.Unlock()
	return keys
}
func (o *PDict) Dict() Scope {
	return o.ppp
}
func (o *PDict) DelItem(i M) {
	key := i.String()
	o.mu.Lock()
	delete(o.ppp, key)
	o.mu.Unlock()
}
func (o *PDict) Compare(a M) int {
	switch b := a.X.(type) {
	case *PDict:
		okeys := o.List()
		akeys := b.List()
		ostrs := make([]string, len(okeys))
		astrs := make([]string, len(akeys))
		for i, x := range okeys {
			ostrs[i] = x.String()
		}
		for i, x := range akeys {
			astrs[i] = x.String()
		}
		sort.Strings(ostrs)
		sort.Strings(astrs)
		olist := make([]M, len(okeys)*2)
		alist := make([]M, len(akeys)*2)
		o.mu.Lock()
		for i, x := range ostrs {
			olist[i*2] = MkStr(x)
			olist[i*2+1] = o.ppp[x]
		}
		o.mu.Unlock()
		b.mu.Lock()
		for i, x := range astrs {
			alist[i*2] = MkStr(x)
			alist[i*2+1] = b.ppp[x]
		}
		b.mu.Unlock()
		return MkList(olist).Compare(MkList(alist))
	}
	return StrCmp(o.PType().String(), a.PType().String())
}

// ========== set

func (o *PSet) BitOr(a M) M { // Union.
	switch t := a.X.(type) {
	case *PSet:
		z := make(Scope)
		Lock2(&o.mu, &t.mu)
		for k, _ := range o.ppp {
			z[k] = True
		}
		for k, _ := range t.ppp {
			z[k] = True
		}
		Unlock2(&o.mu, &t.mu)
		return MkSet(z)
	}
	panic("Operator l|r expects r is set when l is set")
}

func (o *PSet) BitAnd(a M) M { // Intersection.
	switch t := a.X.(type) {
	case *PSet:
		z := make(Scope)
		Lock2(&o.mu, &t.mu)
		for k, _ := range o.ppp {
			if _, ok := t.ppp[k]; ok {
				z[k] = True
			}
		}
		Unlock2(&o.mu, &t.mu)
		return MkSet(z)
	}
	panic("Operator l&r expects r is set when l is set")
}

func (o *PSet) Sub(a M) M { // Subtract Set.
	switch t := a.X.(type) {
	case *PSet:
		z := make(Scope)
		Lock2(&o.mu, &t.mu)
		for k, _ := range o.ppp {
			if _, ok := t.ppp[k]; !ok {
				z[k] = True
			}
		}
		Unlock2(&o.mu, &t.mu)
		return MkSet(z)
	}
	panic("Operator l-r expects r is set when l is set")
}

func (o *PSet) BitXor(a M) M { // Symmetric Difference.
	switch t := a.X.(type) {
	case *PSet:
		z := make(Scope)
		Lock2(&o.mu, &t.mu)
		for k, _ := range o.ppp {
			if _, ok := t.ppp[k]; !ok {
				z[k] = True
			}
		}
		for k, _ := range t.ppp {
			if _, ok := o.ppp[k]; !ok {
				z[k] = True
			}
		}
		Unlock2(&o.mu, &t.mu)
		return MkSet(z)
	}
	panic("Operator l^r expects r is set when l is set")
}

// Partial Ordering (sub- & super-set):

func (o *PSet) LE(a M) bool { // Subset?
	switch t := a.X.(type) {
	case *PSet:
		Lock2(&o.mu, &t.mu)
		for k, _ := range o.ppp {
			if _, ok := t.ppp[k]; !ok {
				Unlock2(&o.mu, &t.mu)
				return false
			}
		}
		Unlock2(&o.mu, &t.mu)
		return true
	}
	panic("Relational Operators expect rhs is set when lhs is set")
}
func (o *PSet) GE(a M) bool { // Superset?
	return a.LE(MkX(&o.PBase))
}
func (o *PSet) LT(a M) bool { // Proper Subset?
	return o.LE(a) && !o.EQ(a)
}
func (o *PSet) GT(a M) bool { // Proper Superset?
	return o.GE(a) && !o.EQ(a)
}

func (o *PSet) Hash() int64 {
	var z int64
	o.mu.Lock()
	for k, _ := range o.ppp {
		z += int64(crc64.Checksum([]byte(k), CrcPolynomial))
	}
	o.mu.Unlock()
	return z
}
func (o *PSet) Pickle(w *bytes.Buffer) {
	o.mu.Lock()
	defer o.mu.Unlock()
	l := int64(len(o.ppp))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypSet + n))
	RypWriteInt(w, l)
	for k, _ := range o.ppp {
		MkStr(k).Pickle(w)
	}
}
func (o *PSet) Contents() interface{} { return o.ppp }
func (o *PSet) Bool() bool            { return len(o.ppp) != 0 }
func (o *PSet) NotContains(a M) bool  { return !o.Contains(a) }
func (o *PSet) Contains(a M) bool {
	key := a.String()
	o.mu.Lock()
	_, ok := o.ppp[key]
	o.mu.Unlock()
	return ok
}
func (o *PSet) Len() int       { return len(o.ppp) }
func (o *PSet) String() string { return o.Repr() }
func (o *PSet) PType() M       { return G_set }
func (o *PSet) RType() string  { return "set" }
func (o *PSet) Repr() string {
	o.mu.Lock()
	vec := make([]string, 0, len(o.ppp))
	for k, _ := range o.ppp {
		vec = append(vec, k)
	}
	o.mu.Unlock()

	sort.Strings(vec)
	buf := bytes.NewBufferString("set([")
	n := len(vec)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(F("%q", vec[i]))
	}
	buf.WriteString("])")
	return buf.String()
}
func (o *PSet) Enough() {}
func (o *PSet) Iter() Nexter {
	var keys []M
	o.mu.Lock()
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	o.mu.Unlock()
	z := &PListIter{PP: keys}
	Forge(z)
	return z
}
func (o *PSet) List() []M {
	var keys []M
	o.mu.Lock()
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	o.mu.Unlock()
	return keys
}
func (o *PSet) Compare(a M) int {
	switch b := a.X.(type) {
	case *PSet:
		o2 := N_sorted(MkList(o.List()), None, None, False)
		a2 := N_sorted(MkList(b.List()), None, None, False)
		return o2.Compare(a2)
	}
	return StrCmp(o.PType().String(), a.PType().String())
}

// ========== object

// TODO: change PtrC_object_er to PtrC_object ?
type PtrC_object_er interface {
	PtrC_object() *C_object
}

func (o *C_object) PType() M      { return G_object }
func (o *C_object) RType() string { return "object" }

func (o *C_object) Repr() string {
	return o.Self.(i__repr__).M_0___repr__().Str()
}
func (o *C_object) String() string {
	return o.Self.(i__str__).M_0___str__().Str()
}
func (o *C_object) M_0___str__() M {
	val := R.ValueOf(o.Self)
	cname := val.Type().Elem().Name()
	if strings.HasPrefix(cname, "C_") {
		cname = cname[2:] // Demangle.
	}
	return MkStr(F("<%s%s>", cname, o.ShortPointerHashString()))
}
func (o *C_object) M_0___repr__() M {
	return MkStr(ShowP(o.Self, SHOW_DEPTH))
}

type i__str__ interface {
	M_0___str__() M
}
type i__repr__ interface {
	M_0___repr__() M
}

func (o *C_object) NE(a M) bool {
	return !(o.Self.EQ(a))
}
func (o *C_object) EQ(a M) bool {
	switch a2 := a.X.(type) {
	case PtrC_object_er:
		a3 := a2.PtrC_object()
		if o == a3 {
			return true
		}
	}
	return false
}

var StringType = R.TypeOf("")
var PBaseType = R.TypeOf(PBase{})
var BType = R.TypeOf(B(&PBase{}))
var MType = R.TypeOf(M{})
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
			v.Field(i).Interface().(M).Pickle(w)
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
	z := &PList{PP: make([]M, 0)}
	Forge(z)
	return z
}

func CopyPs(pp []M) []M {
	zz := make([]M, len(pp))
	copy(zz, pp)
	return zz
}
func CopyList(aa *PList) *PList {
	z := &PList{PP: CopyPs(aa.PP)}
	Forge(z)
	return z
}

/*
func RepeatList(aa *PList, n int64) *PList {
	zz := NewList()
	for i := int64(0); i < n; i++ {
		zz.PP = append(zz.PP, aa.PP...)
	}
	return zz
}
*/

func Enlist(args ...M) M {
	zz := make([]M, 0)
	for _, a := range args {
		zz = append(zz, a)
	}
	return MkList(zz)
}

func Entuple(args ...M) M {
	zz := make([]M, 0)
	for _, a := range args {
		zz = append(zz, a)
	}
	return MkTuple(zz)
}

func (oo *PList) Append(aa M) {
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

func MaybeDerefTwice(t R.Value) R.Value {
	if t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	if t.Kind() == R.Ptr || t.Kind() == R.Interface {
		t = t.Elem()
	}
	return t
}

type PFunc0 struct {
	PBase
	Fn func() M
}

func (p *PFunc0) Call0() M {
	return p.Fn()
}

type PFunc1 struct {
	PBase
	Fn func(a M) M
}

/////func (o *PFunc1) EQ(a M) bool { return (o == a.(*PFunc1)) }
func (p *PFunc1) Call1(a1 M) M {
	return p.Fn(a1)
}

func GetItemMap(r R.Value, x M) M {
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
func SliceGetItem(r R.Value, x M) M {
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

func (o *PGo) Hash() int64           { return int64(o.V.UnsafeAddr()) }
func (o *PGo) Contents() interface{} { return o.V.Interface() }
func (o *PGo) PType() M {
	return MkGo(o.V.Type())
}
func (o *PGo) RType() string { return "go" }

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
	r := MaybeDerefTwice(o.V)
	switch r.Kind() {
	case R.Map, R.Slice, R.Array:
		return r.Len()
	}

	panic(F("Cannot get length of PGo type %t", o.V.Type()))
}
func (o *PGo) GetItem(x M) M {
	r := MaybeDerefTwice(o.V)
	switch r.Kind() {
	case R.Map:
		return GetItemMap(r, x)
	case R.Slice, R.Array:
		return SliceGetItem(r, x)
	}

	panic(F("Cannot GetItem on PGo type %T", o.V.Interface()))
}
func (o *PGo) GetItemSlice(a, b, c M) M {
	r := MaybeDerefTwice(o.V)
	switch r.Kind() {
	case R.Slice, R.Array:
		n := r.Len()

		var i, j int
		if a != None {
			i = int(a.Int())
		}
		if i < 0 {
			i += n
			if i < 0 {
				panic(F("First slicing index on PGo too small: %d", i))
			}
		}
		if i > n {
			panic(F("First slicing index on PGo too large: %d > len: %d", i, n))
		}

		if b == None {
			j = n
		} else {
			j = int(b.Int())
		}
		if j < 0 {
			j += n
			if j < 0 {
				panic(F("Second slicing index on PGo too small: %d", j))
			}
		}
		if j > n {
			j = n
			// panic(F("Second slicing index on PGo too large: %d > len: %d", j, n))
		}

		// TODO: c2 = int(c.Int())
		z := r.Slice(i, j)
		return MkValue(z)
	}

	panic(F("Cannot GetItemSlice on PGo type %t", o.V.Type()))
}
func (g *PGo) Repr() string {
	return F("PGo.Repr{%#v}", g.V.Interface())
}
func (g *PGo) Str() string {
	g0 := g.V
	//i0 := g0.Interface()

	if g0.Type().Kind() == R.Interface {
		g0 = g0.Elem()
	}

	switch g0.Type().Kind() {
	case R.String:
		return g0.Convert(StringType).Interface().(string)
	case R.Slice:
		switch g0.Elem().Type().Kind() {
		case R.Uint8:
			return string(g0.Convert(ByteSliceType).Interface().([]byte))
		}
	}
	panic(fmt.Sprintf("PGo object cannot Str, type: %T", g0.Interface()))
}

func (g *PGo) String() string {
	g0 := g.V
	i0 := g0.Interface()

	// TODO: This should be based on kind.
	switch x := i0.(type) {
	case fmt.Stringer:
		return x.String()
	case string:
		return x
	case error:
		return x.Error()
	case []byte:
		return string(x)
	}

	switch g0.Type().Kind() {
	case R.String:
		return g0.Convert(StringType).Interface().(string)
	case R.Int:
		return F("%d", i0)
	case R.Int64:
		return F("%d", i0)
	case R.Uint:
		return F("%d", i0)
	case R.Uint64:
		return F("%d", i0)
	}

	switch g0.Kind() {
	case R.Array, R.Slice:
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

var Float64Type = R.TypeOf(float64(0))
var Int64Type = R.TypeOf(int64(0))
var IntType = R.TypeOf(int(0))
var PType = R.TypeOf(new(P)).Elem()

func (o *PGo) ForceInt() int64 {
	return o.Self.Int()
}
func (o *PGo) CanInt() bool {
	x := o.V
	t := x.Type()
	switch {
	case t.ConvertibleTo(Int64Type):
		return true
	case t.ConvertibleTo(IntType):
		return true
	}
	x = MaybeDeref(x)
	t = x.Type()
	switch {
	case t.ConvertibleTo(Int64Type):
		return true
	case t.ConvertibleTo(IntType):
		return true
	}
	return false
}
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
	panic(F("PGo cannot convert to int64: %s", o.V.Type().String()))
}

func (o *PGo) ForceFloat() float64 {
	return o.Self.Float()
}
func (o *PGo) CanFloat() bool {
	x := o.V
	t := x.Type()
	switch {
	case t.ConvertibleTo(Float64Type):
		return true
	}
	x = MaybeDeref(x)
	t = x.Type()
	switch {
	case t.ConvertibleTo(Float64Type):
		return true
	}
	return false
}
func (o *PGo) Float() float64 {
	x := o.V
	t := x.Type()
	switch {
	case t.ConvertibleTo(Float64Type):
		return x.Convert(Float64Type).Float()
	}
	x = MaybeDeref(x)
	t = x.Type()
	switch {
	case t.ConvertibleTo(Float64Type):
		return x.Convert(Float64Type).Float()
	}
	panic(F("PGo cannot convert to float64: %s", o.V.Type().String()))
}

func InvokeMap(r R.Value, field string, aa []M) M {
	switch {
	case field == "Keys" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]M, len(keys))
		for i, e := range keys {
			pp[i] = AdaptForReturn(e)
		}
		return MkList(pp)
	case field == "Values" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]M, len(keys))
		for i, e := range keys {
			pp[i] = AdaptForReturn(r.MapIndex(e))
		}
		return MkList(pp)
	case field == "Items" && len(aa) == 0:
		keys := r.MapKeys()
		pp := make([]M, len(keys))
		for i, e := range keys {
			pp[i] = MkTuple([]M{AdaptForReturn(e), AdaptForReturn(r.MapIndex(e))})
		}
		return MkList(pp)
	case field == "Get" && (len(aa) == 1 || len(aa) == 2):
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

func (g *PGo) Invoke(field string, aa ...M) M {
	// We cannot invoke private field, so change the first letter to upper case.
	// This smooths over 'flush' vs 'Flush' and 'write' vs 'Write' for builtin
	// stdout & stderr objects.
	if 'a' <= field[0] && field[0] <= 'z' {
		field = string([]byte{byte(field[0] - 32)}) + field[1:]
	}

	r := g.V
	if meth, ok := r.Type().MethodByName(field); ok && meth.Func.IsValid() {
		return FinishInvokeOrCall(field, meth.Func, r, aa)
	}
	if r.Kind() == R.Map {
		return InvokeMap(r, field, aa)
	}

	r = MaybeDeref(r)
	if meth, ok := r.Type().MethodByName(field); ok && meth.Func.IsValid() {
		return FinishInvokeOrCall(field, meth.Func, r, aa)
	}
	if r.Kind() == R.Map {
		return InvokeMap(r, field, aa)
	}

	panic(F("Method on kind %q type %q origkind %q origtype %q does not exist: %s", r.Kind(), r.Type(), g.V.Kind(), g.V.Type(), field))
}

func (g *PGo) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	if len(kv1) > 0 || len(kv2) > 0 {
		panic("Cannot call GO with named (keyword) parameters.")
	}
	if a2 == nil {
		return g.Call(a1...)
	}
	if a1 == nil {
		return g.Call(a2...)
	}
	var aa []M
	aa = append(aa, a1...)
	aa = append(aa, a2...)
	return g.Call(aa...)
}

func (g *PGo) Call(aa ...M) M {
	f := MaybeDeref(g.V)
	if f.Kind() != R.Func {
		z, ok := FunCallN(f, aa)
		if !ok {
			panic(F("cannot Call when Value not a func and FunCallN fails: %T", f.Interface()))
		}
		return z
	}
	var zeroValue R.Value
	return FinishInvokeOrCall("?", f, zeroValue, aa)
}

// TODO -- make List() the primary, instead of Iter()
func (g *PGo) List() []M {
	return g.Iter().(*PListIter).PP
}
func (g *PGo) Iter() Nexter {
	a := MaybeDeref(g.V)
	var pp []M

	switch a.Kind() {
	case R.Array, R.Slice:
		n := a.Len()
		pp = make([]M, n)
		for i := 0; i < n; i++ {
			pp[i] = AdaptForReturn(a.Index(i))
		}
	case R.Map:
		keys := a.MapKeys()
		pp = make([]M, len(keys))
		for i, k := range keys {
			pp[i] = AdaptForReturn(k)
		}
	default:
		panic(F("*PGo cannot Iter() on %T", a.Interface()))
	}
	z := &PListIter{PP: pp}
	Forge(z)
	return z
}
func (g *PGo) Dict() Scope {
	z := make(Scope)
	a := MaybeDeref(g.V)

	switch a.Kind() {
	case R.Map:
		keys := a.MapKeys()
		for _, k := range keys {
			v := a.MapIndex(k)
			if !v.IsValid() {
				continue // It disappeared while iterating.
			}
			z[AdaptForReturn(k).String()] = AdaptForReturn(v)
		}
	default:
		panic(F("*PGo cannot Dict() on %T", a.Interface()))
	}
	return z
}

func (g *PGo) SetItem(i M, x M) {
	a := MaybeDeref(g.V)

	switch a.Kind() {
	case R.Array:
		i2 := int(i.Int())
		x2 := AdaptForCall(x, a.Type().Elem())
		a.Slice(0, a.Len()).Index(i2).Set(x2)
	case R.Slice:
		i2 := int(i.Int())
		x2 := AdaptForCall(x, a.Type().Elem())
		a.Index(i2).Set(x2)
	case R.Map:
		i2 := R.ValueOf(i.String())
		x2 := AdaptForCall(x, a.Type().Elem())
		a.SetMapIndex(i2, x2)
	default:
		panic(F("*PGo cannot Iter() on %T", a.Interface()))
	}
}

var errorType = R.TypeOf(new(error)).Elem()

func FinishInvokeOrCall(field string, f R.Value, rcvr R.Value, aa []M) M {
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
			panic(F("call got %d args, want %d or more args", lenIns, numIn-1))
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
			panic(F("call got %d args, want %d args", lenIns, numIn))
		}
		if lenIns > 0 {
			args[0] = rcvr
		}
		for i, a := range aa {
			args[i+lenRcvr] = AdaptForCall(a, ft.In(i+lenRcvr))
		}
	}

	if DebugReflect > 0 {
		for ci, ce := range args {
			fmt.Fprintf(os.Stderr, "r::  <<< CALL %v <<< [arg%d] <<< %#v\n", f.Interface(), ci, ce.Interface())
		}
		fmt.Fprintf(os.Stderr, "r::  --- CALL %v ...\n", f.Interface())
	}
	outs := f.Call(args)
	if DebugReflect > 0 {
		for oi, oe := range outs {
			fmt.Fprintf(os.Stderr, "r::  >>> CALL %v >>> [ret%d] >>> %#v\n", f.Interface(), oi, oe.Interface())
		}
		fmt.Fprintf(os.Stderr, "r::  ======\n")
	}

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
		slice := make([]M, numOut)
		for i := 0; i < numOut; i++ {
			slice[i] = AdaptForReturn(outs[i])
		}
		return MkTuple(slice)
	}
}

var typeInterfaceEmpty = R.TypeOf(new(interface{})).Elem()
var typeP = R.TypeOf(new(P)).Elem()

func GoDeref(p M) M {
	switch x := p.X.(type) {
	case *PGo:
		switch x.V.Kind() {
		case R.Ptr:
			return MkValue(x.V.Elem())
		}
		panic(F("Cannot goderef non-pointer: %s", x.V.Kind()))
	}
	panic(F("Cannot goderef non-*Go value: %T", p))
}

func GoCast(want M, p M) M {
	typ := want.Contents().(R.Type)
	return MkValue(AdaptForCall(p, typ))
}

func GoAppend(slice M, a M) M {
	return MkValue(R.Append(R.ValueOf(slice.Contents()), R.ValueOf(a.Contents())))
}

func AdaptForCall(v M, want R.Type) R.Value {
	if DebugReflect > 0 {
		Say("AdaptForCall <<<<<<", v, want, F("%#v", v))
	}
	z := adaptForCall2(v, want)
	if DebugReflect > 0 {
		Say("AdaptForCall >>>>>>", z)
		if z.Kind() == R.Interface {
			println("R.Value--Interface", z.InterfaceData()[0], z.InterfaceData()[1])
			if (z.InterfaceData()[1] & 7) != 0 {
				panic("z.InterfaceData()[1] not aligned")
			}
		}
	}
	return z
}
func adaptForCall2(v M, want R.Type) R.Value {
	// None & nil.
	contents := v.Contents()
	if contents == nil {
		switch want.Kind() {

		case R.Chan, R.Func, R.Interface, R.Map, R.Ptr, R.Slice:
			if DebugReflect > 0 {
				Say("AdaptForCall :::::: R.Zero")
			}
			return R.Zero(want)

		default:
			if DebugReflect > 0 {
				Say("AdaptForCall :::::: contents is nil, want", want)
			}
			panic(F("Cannot convert None to go type %v", want))

		}
	}

	if want == BType {
		return R.ValueOf(v)
	}

	if want == MType {
		return R.ValueOf(v)
	}

	// Try builtin conversion:
	vcontents := R.ValueOf(contents)
	tcontents := vcontents.Type()
	if tcontents.ConvertibleTo(want) {
		if DebugReflect > 0 {
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
		return R.ValueOf(v.Str())
	case R.Func:
		return MakeFunction(v, want) // This is hard.
	case R.Array:
		switch want.Elem().Kind() {
		case R.Uint8:
			switch v.X.(type) {

			case nil:
				if v.S != "" {
					// Same as below str & byt cases.
					bb := v.Bytes()
					wl := want.Len()
					if len(bb) != wl {
						panic(F("Cannot convert []byte len %d to array len %d", len(bb), wl))
					}
					arr := R.New(want).Elem()
					for i := 0; i < wl; i++ {
						arr.Index(i).Set(R.ValueOf(bb[i]))
					}
					return arr
				}

			case *PStr, *PByt:
				// Same as above nil case.
				bb := v.Bytes()
				wl := want.Len()
				if len(bb) != wl {
					panic(F("Cannot convert []byte len %d to array len %d", len(bb), wl))
				}
				arr := R.New(want).Elem()
				for i := 0; i < wl; i++ {
					arr.Index(i).Set(R.ValueOf(bb[i]))
				}
				return arr

			}
		}
	case R.Slice:
		switch want.Elem().Kind() {
		case R.Uint8:
			switch vx := v.X.(type) {
			case nil:
				if v.S != "" {
					bb := make([]byte, v.Len())
					copy(bb, v.String())
					return R.ValueOf(bb)
				}
			case *PStr:
				bb := make([]byte, v.Len())
				copy(bb, v.String())
				return R.ValueOf(bb)
			case *PByt:
				return R.ValueOf(vx.YY)
			}
		}

		// For the "in" case.  TODO: "in out"?
		if tcontents.Kind() == R.Slice {
			n := vcontents.Len()
			sl := R.MakeSlice(want, n, n)
			for i := 0; i < n; i++ {
				v1 := vcontents.Index(i)
				var v2 R.Value
				if vp, ok := v1.Interface().(M); ok {
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
					v2 = AdaptForCall(vp.M(), want.Elem())
				} else {
					v2 = AdaptForCall(MkValue(v1), want.Elem())
				}
				m.SetMapIndex(k2, v2)
			}
			return m
		}
	}
	if DebugReflect > 0 {
		Say("AdaptForCall :::::: Not Case")
	}

	if want == typeInterfaceEmpty {
		if DebugReflect > 0 {
			Say("AdaptForCall :::::: Interface Empty")
		}
		return R.ValueOf(v.Contents())
	}

	if want == typeP {
		if DebugReflect > 0 {
			Say("AdaptForCall :::::: Interface P")
		}
		return R.ValueOf(v)
	}

	if DebugReflect > 0 {
		Say("AdaptForCall :::::: Panic.")
	}
	panic(F("Cannot AdaptForCall: %s [%s] %q [%s] TO %s [%s]", v, R.TypeOf(v), v.Repr(), R.TypeOf(v.Contents()), want, want.Kind()))
}

func MakeFunction(v M, ft R.Type) R.Value {
	nin := ft.NumIn()
	if nin > 3 {
		panic(F("Not implemented: MakeFunction for %d args", nin))
	}

	//println("MF: Inside MakeFunction", v, ft, nin)
	return R.MakeFunc(ft, func(aa []R.Value) (zz []R.Value) {
		var r M
		var err error = error(nil)
		//println("MF: Running Made Function", nin, aa)

		func() {
			defer func() {
				rec := recover()
				//println("MF: MakeFunction recovered=", Show(rec))
				if rec != nil {
					err = errors.New(F("%v", rec))
					//println("MF: MakeFunction err=", Show(err))
				}
			}()
			//println("MF: Running Inner Function", nin, Show(aa))
			switch nin {
			case 0:
				r = v.X.(I_0).Call0()
			case 1:
				r = v.X.(I_1).Call1(AdaptForReturn(aa[0]))
			case 2:
				r = v.X.(I_2).Call2(AdaptForReturn(aa[0]), AdaptForReturn(aa[1]))
			case 3:
				r = v.X.(I_3).Call3(AdaptForReturn(aa[0]), AdaptForReturn(aa[1]), AdaptForReturn(aa[2]))
			default:
				panic(F("Not implemented: MakeFunction for %d args", nin))
			}
			//println("MF: Set R to", Show(r))
		}()

		orig_nout := ft.NumOut() // orig_nout counts a final error return.
		nout := orig_nout        // nout ignores a final error return.
		if orig_nout > 0 && ft.Out(orig_nout-1) == errorType {
			nout -= 1 // Ignore final error result temporarily.
		}
		if err != nil {
			zz = make([]R.Value, nout)
			for i := 0; i < nout; i++ {
				zz[i] = R.Zero(ft.Out(i))
			}
		} else {
			switch nout {
			case 0:
				// pass
			case 1:
				zz = []R.Value{AdaptForCall(r, ft.Out(0))}
			default:
				zz = make([]R.Value, nout)
				for i := 0; i < nout; i++ {
					zz[i] = AdaptForCall(r.GetItem(Mkint(i)), ft.Out(i))
				}
			}
		}
		if orig_nout > nout {
			// Now append the final error slot.
			if err == nil {
				zz = append(zz, R.Zero(errorType))
			} else {
				zz = append(zz, R.ValueOf(err).Convert(errorType))
			}
		}
		//println("MF: Returning", Show(zz))
		return
	})
}

func AdaptForReturn(v R.Value) M {
	if DebugReflect > 0 {
		if v.Kind() == R.Interface {
			println("R.Value--Interface", v.InterfaceData()[0], v.InterfaceData()[1])
			if (v.InterfaceData()[1] & 7) != 0 {
				panic("v.InterfaceData()[1] not aligned")
			}
		}
	}
	// Say("AdaptForReturn <<<", v, v.Type())
	z := AdaptForReturn9(v)
	// Say("AdaptForReturn >>>", z, z.Type())
	return z
}

func AdaptForReturn9(v R.Value) M {
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
	if b, ok := v.Interface().(M); ok {
		return b
	}
	if p, ok := v.Interface().(P); ok {
		return p.M()
	}
	return MkValue(v)
}

func (g *PGo) FetchField(field string) M {
	return FetchFieldByName(g.V, field)
}

func (g *PGo) StoreField(field string, p M) {
	StoreFieldByName(g.V, field, p)
}

var Classes map[string]R.Type

func init() {
	Classes = make(map[string]R.Type)
}

func FunCallN(f R.Value, aa []M) (M, bool) {
	switch len(aa) {
	case 0:
		if c, ok := f.Interface().(I_0); ok {
			return c.Call0(), true
		}
	case 1:
		if c, ok := f.Interface().(I_1); ok {
			return c.Call1(aa[0]), true
		}
	case 2:
		if c, ok := f.Interface().(I_2); ok {
			return c.Call2(aa[0], aa[1]), true
		}
	case 3:
		if c, ok := f.Interface().(I_3); ok {
			return c.Call3(aa[0], aa[1], aa[2]), true
		}
	case 4:
		if c, ok := f.Interface().(I_4); ok {
			return c.Call4(aa[0], aa[1], aa[2], aa[3]), true
		}
		// TODO: Fall back to reflection.
	}

	return None, false
}

func GoElemType(pointedTo interface{}) M {
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

	return string(bb)
}

func UnPickle(b []byte) M {
	return RypUnPickle(bytes.NewBuffer(b))
}

func RypUnPickle(b *bytes.Buffer) M {
	tag, err := b.ReadByte()
	if err != nil {
		panic(err)
	}

	kind := int(tag & RypMask)
	arg := int(tag & 7)

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
		pp := make([]M, n)
		for i := 0; i < n; i++ {
			pp[i] = RypUnPickle(b)
		}
		return MkTuple(pp)
	case RypList:
		n := int(RypReadInt(b, arg))
		pp := make([]M, n)
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
	case RypSet:
		n := int(RypReadInt(b, arg))
		ppp := make(Scope)
		for i := 0; i < n; i++ {
			k := RypUnPickle(b)
			ppp[k.String()] = True
		}
		return MkSet(ppp)
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
		return cobj.M()
	}
	panic(F("RypUnPickle: bad tag: %d", tag))
}

var VarOfStarP *P

func RypSetField(obj R.Value, fname string, x M) {
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

func PrintStackFYIUnlessEOFBecauseExcept(e interface{}) {
	if DebugExcept > 0 {
		switch t := e.(type) {
		case *PGo:
			e = t.Contents()
		}
		switch t := e.(type) {
		case error:
			if t.Error() == "EOF" {
				return
			}
		case fmt.Stringer:
			e = t.String()
		}
		s := fmt.Sprintf("%s", e)
		if s == "EOF" {
			return
		}
		PrintStackFYI(e)
	}
}

func PrintStackFYI(e interface{}) {
	Flushem()
	fmt.Fprintf(os.Stderr, "\nFYI{{{[[[((( %v\n", e)
	if DebugExcept > 1 {
		debug.PrintStack()
		fmt.Fprintf(os.Stderr, "\n######\n")
	}

	rs := RyeStack()
	fmt.Fprintf(os.Stderr, "\n%s\n", rs)

	fmt.Fprintf(os.Stderr, "FYI)))]]]}}}\n")
}

var RYEMODULE_GO_FILENAME = regexp.MustCompile(`^(/.*/src/)(.+)/rye__/([^/].+)/ryemodule[.]go$`)

func MatchGoFilenameToRyeFilenameOrEmpty(gofile string) (pyfile string, pkg string) {
	m := RYEMODULE_GO_FILENAME.FindStringSubmatch(gofile)
	if m != nil {
		pyfile = m[1] + m[2] + "/" + m[3] + ".py"
		pkg = m[2] + "/rye__/" + m[3]
	}
	return
}

func RyeStack() string {
	var bb bytes.Buffer
	var prevFunc string
	for i := 0; i < 100; i++ {
		_, filename, lineno, ok := runtime.Caller(i)
		if !ok {
			break
		}
		pyFile, pkg := MatchGoFilenameToRyeFilenameOrEmpty(filename)
		if pyFile == "" {
			continue
		}
		if info, ok := LineInfoRegistry[pkg]; ok {
			lm := info.LookupLineNumber
			if 1 <= lineno && lineno < len(lm) {
				pylineno := int(lm[lineno])
				if pylineno > 0 {
					// Linear search for current func.
					for _, pair := range info.SourceWhats {
						if pair.N == pylineno {
							if pair.S != prevFunc {
								fmt.Fprintf(&bb, "%s\n", pair.S)
							}
							prevFunc = pair.S
							break
						}
					}

					// Linear search for source line.
					for _, pair := range info.SourceLines {
						if pair.N == pylineno {
							fmt.Fprintf(&bb, "   >>>> %s\n", pair.S)
							break
						}
					}

					fmt.Fprintf(&bb, "\t\t\t\t[%5d] %s:%d\n", i, pyFile, pylineno)
				}
			}
		}
		if DebugExcept > 1 {
			fmt.Fprintf(&bb, "%s:%d\n", filename, lineno)
		}
	}
	return bb.String()
}

// source returns a space-trimmed slice of the n'th line.
func nthLine(lines [][]byte, n int) []byte {
	n-- // in stack trace, lines are 1-indexed but our array is 0-indexed
	if n < 0 || n >= len(lines) {
		return []byte{}
	}
	return bytes.Trim(lines[n], " \t\r")
}

func FetchFieldByName(v R.Value, field string) M {
	// First try for method:
	meth := v.MethodByName(field)
	if meth.IsValid() {
		return MkValue(meth)
	}

	// Then try for field:
	v2 := MaybeDerefTwice(v)
	if v2.Kind() != R.Struct {
		panic(F("FetchFieldByName: Cannot get field %q from non-Struct %#v", field, v2))
	}
	x := v2.FieldByName(field)
	if !x.IsValid() {
		panic(F("FetchFieldByName: No such field %q on %T %#v", field, v2.Interface(), v2))
	}
	return AdaptForReturn(x)
}
func StoreFieldByName(v R.Value, field string, a M) {
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

// Old PCallable

type CallSpec struct {
	Name     string
	Args     []string
	Defaults []M
	Star     string
	StarStar string
}

type PNewCallable struct {
	PBase
	CallSpec *CallSpec
}

func (o *PNewCallable) Callable() bool { return true }

func (o *PNewCallable) String() string {
	return fmt.Sprintf("<func %s>", o.CallSpec.Name)
}

func (o *PNewCallable) Repr() string {
	return o.String()
}

// Old PCallable

type PCallable struct {
	PBase
	Name     string
	Args     []string
	Defaults []M
	Star     string
	StarStar string
}

func (o *PCallable) Callable() bool { return true }

func (o *PCallable) String() string {
	return fmt.Sprintf("<func %s>", o.Name)
}

func (o *PCallable) Repr() string {
	return o.String()
}

type KVSlice []KV            // Can be sorted by Key.
func (vec KVSlice) Len() int { return len(vec) }
func (vec KVSlice) Less(i, j int) bool {
	return vec[i].Key < vec[j].Key
}
func (vec KVSlice) Swap(i, j int) {
	vec[i], vec[j] = vec[j], vec[i]
}

type KV struct {
	Key   string
	Value M
}

func NewSpecCall(cs *CallSpec, a1 []M, a2 []M, kv []KV, kv2 map[string]M) ([]M, *PList, *PDict) {
	n := len(cs.Defaults)
	argv := make([]M, n)
	var star []M
	var starstar map[string]M

	copy(argv, cs.Defaults) // Copy defaults first, any of which may be nil.

	j := 0
	for a1 != nil {
		for _, a := range a1 {
			if j < n {
				argv[j] = a
				j++
			} else {
				star = append(star, a)
			}
		}
		a1 = a2
		a2 = nil
	}

	for _, e := range kv {
		k := e.Key
		v := e.Value
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				break
			}
		}
		if !stored {
			if starstar == nil {
				starstar = make(map[string]M)
			}
			starstar[k] = v
		}
	}

	for k, v := range kv2 {
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				break
			}
		}
		if !stored {
			if starstar == nil {
				starstar = make(map[string]M)
			}
			starstar[k] = v
		}
	}

	for i, e := range argv {
		if e == MissingM {
			panic(F("The %dth fixed formal argument '%s' of function %q has no assigned value (fixed formal args are: %v)", i+1, cs.Args[i], cs.Name, cs.Args))
		}
	}

	if cs.Star == "" && len(star) > 0 {
		panic(F("Function %q wants %d args, but got %d (no * arg)", cs.Name, len(cs.Args), len(cs.Args)+len(star)))
	}

	if cs.StarStar == "" && len(starstar) > 0 {
		panic(F("Function %q cannot take %d extra named args", cs.Name, len(starstar)))
	}

	return argv, PMkList(star), PMkDict(starstar)
}

func SpecCall(cs *PCallable, a1 []M, a2 []M, kv []KV, kv2 map[string]M) ([]M, *PList, *PDict) {
	n := len(cs.Defaults)
	argv := make([]M, n)
	var star []M
	var starstar map[string]M

	copy(argv, cs.Defaults) // Copy defaults first, any of which may be nil.

	j := 0
	for a1 != nil {
		for _, a := range a1 {
			if j < n {
				argv[j] = a
				j++
			} else {
				star = append(star, a)
			}
		}
		a1 = a2
		a2 = nil
	}

	for _, e := range kv {
		k := e.Key
		v := e.Value
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				break
			}
		}
		if !stored {
			if starstar == nil {
				starstar = make(map[string]M)
			}
			starstar[k] = v
		}
	}

	for k, v := range kv2 {
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = v
				stored = true
				break
			}
		}
		if !stored {
			if starstar == nil {
				starstar = make(map[string]M)
			}
			starstar[k] = v
		}
	}

	for i, e := range argv {
		if e == MissingM {
			panic(F("The %dth fixed formal argument '%s' of function %q has no assigned value (fixed formal args are: %v)", i+1, cs.Args[i], cs.Name, cs.Args))
		}
	}

	if cs.Star == "" && len(star) > 0 {
		panic(F("Function %q wants %d args, but got %d (no * arg)", cs.Name, len(cs.Args), len(cs.Args)+len(star)))
	}

	if cs.StarStar == "" && len(starstar) > 0 {
		panic(F("Function %q cannot take %d extra named args", cs.Name, len(starstar)))
	}

	return argv, PMkList(star), PMkDict(starstar)
}

type ICallV interface {
	CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M
}

func StrCmp(a, b string) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	case a == b:
		return 0
	}
	panic("StrCmp Failure")
}

// If pye/sys is linked in, it will change these pointers to its std{in,out,err}.
// If not, then nobody can change sys.std{in,out,err}, and these remain nil.
var PtrSysStdin *M
var PtrSysStdout *M
var PtrSysStderr *M

type PythonWriter interface {
	M_1_write(M) M
}

type AdaptPythonWriter struct {
	PythonW PythonWriter
}

func (self AdaptPythonWriter) Write(p []byte) (n int, err error) {
	self.PythonW.M_1_write(MkByt(p))
	return len(p), nil
}

func CurrentStdout() io.Writer {
	if PtrSysStdout == nil || M(*PtrSysStdout).Bool() == false {
		return os.Stdout
	}
	if w, ok := M(*PtrSysStdout).Contents().(io.Writer); ok {
		return w
	}
	if pw, ok := M(*PtrSysStdout).Contents().(PythonWriter); ok {
		return AdaptPythonWriter{PythonW: pw}
	}
	panic(F("CurrentStdout: not an io.Writer: %#v", M(*PtrSysStdout).Contents()))
}

func CurrentStderr() io.Writer {
	if PtrSysStderr == nil || M(*PtrSysStderr).Bool() == false {
		return os.Stderr
	}
	if w, ok := M(*PtrSysStderr).Contents().(io.Writer); ok {
		return w
	}
	if pw, ok := M(*PtrSysStderr).Contents().(PythonWriter); ok {
		return AdaptPythonWriter{PythonW: pw}
	}
	panic(F("CurrentStderr: not an io.Writer: %#v", M(*PtrSysStderr).Contents()))
}

type Flusher interface {
	Flush() error
}

func Flushem() {
	if PtrSysStdout != nil {
		if fl, ok := M(*PtrSysStdout).X.(Flusher); ok {
			err := fl.Flush()
			if err != nil {
				panic(F("Flushem: PtrSysStdout Flush: %s", err))
			}
		}
	}
	if PtrSysStderr != nil {
		if fl, ok := M(*PtrSysStderr).X.(Flusher); ok {
			err := fl.Flush()
			if err != nil {
				panic(F("Flushem: PtrSysStderr Flush: %s", err))
			}
		}
	}
}

type PModule struct {
	PBase
	ModName string
	Map     map[string]*M
}

func MakeModuleObject(m map[string]*M, modname string) M {
	z := &PModule{
		ModName: modname,
		Map:     m,
	}
	return MForge(z)
	//return &z.PBase
}

func (o *PModule) String() string { return F("<module %s>", o.ModName) }
func (o *PModule) Repr() string   { return F("<module %s>", o.ModName) }
func (o *PModule) FetchField(field string) M {
	if ptr, ok := o.Map[field]; ok {
		return *ptr
	}
	return MissingM
}
func (o *PModule) Iter() Nexter {
	return MkList(o.List()).Iter()
}
func (o *PModule) List() []M {
	var z []M
	for k, _ := range o.Map {
		z = append(z, MkStr(k))
	}
	return z
}
func (o *PModule) Dict() Scope {
	z := make(Scope)
	for k, ptr := range o.Map {
		z[k] = *ptr
	}
	return z
}

// NewErrorOrEOF take anything (like something recovered) and converts it to error, using io.EOF if needed.
func NewErrorOrEOF(r interface{}) error {
	if r != nil {
		var s string
		switch t := r.(type) {
		case error:
			return t
		case string:
			s = t
		case *PGo:
			rr := t.V.Interface()
			switch tt := rr.(type) {
			case error:
				return tt
			}
			s = fmt.Sprintf("%v", rr)
		case P:
			s = t.String()
		default:
			s = fmt.Sprintf("%v", r)
		}
		if s == "EOF" {
			return io.EOF
		}
		return errors.New(s)
	}
	return nil
}

//##################################//

type IntStringPair struct {
	N int
	S string
}

type LineInfo struct {
	LookupLineNumber []int32
	SourceLines      []IntStringPair
	SourceWhats      []IntStringPair
}

var LineInfoRegistry = make(map[string]*LineInfo)

func RegisterLineInfo(longmod string, info *LineInfo) {
	LineInfoRegistry[longmod] = info
}

//##################################//

// CheckTyp wants obj to be one of the types in typs.
func CheckTyp(name string, obj M, typs ...M) {
	ot := obj.PType()
	// Check quickly without subclasses.
	for _, t := range typs {
		// HACK around a special case.  TODO: fix type(None)!
		if t == None && obj == None {
			return
		} else if ot == t {
			return
		}
	}
	// Check with subclass.
	for _, t := range typs {
		if IsSubclass(ot, t) {
			return
		}
	}
	if obj == None {
		// Be soft about None, for now.
		// TODO: if IsSubclass(ot, G_object) // did not work.
		// Allow nil objects.
		return
		//}
	}
	panic(F("For %s, got object of type %s, wanted any of %v", name, obj.PType().String(), typs))
}

func IsSubclass(subcls, cls M) bool {
	for {
		if subcls == cls {
			return true
		}
		subcls = subcls.Superclass()
		if subcls == None {
			break
		}
	}
	return false
}

//##################################//

// It would be nice to trim this down.

var F = fmt.Sprintf

// Show objects as a string.
func Show(aa ...interface{}) string {
	buf := bytes.NewBuffer(nil)
	for _, a := range aa {
		switch x := a.(type) {
		case string:
			buf.WriteString(F("string %q ", x))
		case []byte:
			buf.WriteString(F("[]byte [%d] %q ", len(x), string(x)))
		case int:
			buf.WriteString(F("int %d ", x))
		case int64:
			buf.WriteString(F("int64 %d ", x))
		case float32:
			buf.WriteString(F("float32 %f ", x))
		case fmt.Stringer:
			buf.WriteString(F("Stringer %T %q ", a, x))
		case error:
			buf.WriteString(F("{error:%s} ", x))
		default:
			v := R.ValueOf(a)
			switch v.Kind() {
			case R.Slice:
				n := v.Len()
				buf.WriteString(F("%d[ ", n))
				for i := 0; i < n; i++ {
					buf.WriteString(Show(v.Index(i).Interface()))
					buf.WriteString(" , ")
				}
				buf.WriteString("] ")
			case R.Map:
				n := v.Len()
				buf.WriteString(F("%d{ ", n))
				kk := v.MapKeys()
				for _, k := range kk {
					buf.WriteString(Show(k.Interface()))
					buf.WriteString(": ")
					buf.WriteString(Show(v.MapIndex(k).Interface()))
					buf.WriteString(", ")
				}
				buf.WriteString("} ")
			default:
				buf.WriteString(F("WUT{%#v} ", x))
			}
		}
	}
	return buf.String()
}

// Say arguments on stderr.
func Say(aa ...interface{}) {
	buf := bytes.NewBuffer([]byte("## "))
	for _, a := range aa {
		buf.WriteString(Show(a))
		buf.WriteString(" ; ")
	}

	fmt.Fprintf(os.Stderr, "## %s\n", buf)
}

////////////////////////////
//
//  Reusable Callers.

// PCall0

type PCall0 struct {
	PNewCallable
	Fn   func() M
}

func (o *PCall0) Contents() interface{} {
	return o.Fn
}
func (o PCall0) Call0() M {
	return o.Fn()
}

func (o PCall0) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn()
}

// PCall1

type PCall1 struct {
	PNewCallable
	Fn   func(a0 M) M
}

func (o *PCall1) Contents() interface{} {
	return o.Fn
}
func (o PCall1) Call1(a0 M) M {
	return o.Fn(a0)
}

func (o PCall1) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn(argv[0])
}

// PCall2

type PCall2 struct {
	PNewCallable
	Fn   func(a0 M, a1 M) M
}

func (o *PCall2) Contents() interface{} {
	return o.Fn
}
func (o PCall2) Call2(a0 M, a1 M) M {
	return o.Fn(a0, a1)
}

func (o PCall2) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn(argv[0], argv[1])
}

// PCall3

type PCall3 struct {
	PNewCallable
	Fn   func(a0 M, a1 M, a2 M) M
}

func (o *PCall3) Contents() interface{} {
	return o.Fn
}
func (o PCall3) Call3(a0 M, a1 M, a2 M) M {
	return o.Fn(a0, a1, a2)
}

func (o PCall3) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn(argv[0], argv[1], argv[2])
}

// END.
