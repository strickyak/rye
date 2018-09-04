// +build prego

package two

import . "github.com/strickyak/rye"
import (
	"bytes"
	"fmt"
	"hash/crc64"
	"log"
	"math"
	//"reflect"
	R "reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

var ReifyString *string
var ReifyU U
var ReifyV V
var ReifyPJ PJ

var Log = log.Printf

type U uintptr
type V unsafe.Pointer
type W struct {
	U U
	V V
}

type InterfaceGuts struct {
	V V
	U U
}

var OJNone = &JNone{}
var OJTrue = &JBool{Z: true, I: 1, F: 1.0}
var OJFalse = &JBool{Z: false, I: 0, F: 0.0}

var NoneJ PJ = ForgeJ(OJNone)
var TrueJ PJ = ForgeJ(OJTrue)
var FalseJ PJ = ForgeJ(OJFalse)

var NoneJ_1 U = 16 | Py
var NoneJ_2 V = V(unsafe.Pointer(&OJNone.JBase))
var TrueJ_1 U = 16 | Py
var TrueJ_2 V = V(unsafe.Pointer(&OJTrue.JBase))
var FalseJ_1 U = 16 | Py
var FalseJ_2 V = V(unsafe.Pointer(&OJFalse.JBase))

var V0 V

var Filler struct {
	Addr [16]byte
}

func NOprintln(args ...interface{}) {}

//////////////////////////////////

func Assert(what string, b bool) {
	//#if asserts
	if !b {
		panic(F("Assert failed: %s", what))
	}
	//#endif
}

const Py = 0
const Int = 1
const Str = 2

func Slow() {
	panic("Slow Op Not Supported Yet")
}
func Wrong() {
	panic("Wrong types")
}
func Todo() {
	panic("TO DO")
}

func Peek(u U, v V, s string, args... interface{}) {
  msg := fmt.Sprintf(s, args...)
  switch inline.Tag(u, v) {
  case Py:
    a := inline.TakePJ(u, v)
    log.Printf("Peek{%s} [$%x,$%x] (%s) %s", msg, u, uintptr(unsafe.Pointer(v)), a.TypeName(), a.Repr())
  case Int:
    a := inline.TakeIntJ(u, v)
    log.Printf("Peek{%s} [$%x,$%x] Int: %d", msg, u, uintptr(unsafe.Pointer(v)), a)
  case Str:
    a := inline.TakeStrJ(u, v)
    log.Printf("Peek{%s} [$%x,$%x] [%d]Str: %q", msg, u, uintptr(unsafe.Pointer(v)), len(a), a)
    for i := 0; i < len(a); i++ {
      if a[i] == 0 {
        panic("STOP")
      }
    }
  }
}

func ForgeJ(p PJ) PJ {
	//#if c
	if DebugCounters > 0 {
		f_name := fmt.Sprintf("ForgeJ:%T", p)
		ptr := CounterMap[f_name]
		if ptr == nil {
			ptr = new(int64)
			CounterMap[f_name] = ptr
		}
		(*ptr)++
	}
	//#endif
	p.SetSelf(p)
	return p // .BJ()
}

func ForceIntJ(u U, v V) int64 {
	switch inline.Tag(u, v) {
	case Py:
		a := inline.TakePJ(u, v)
		return a.ForceInt()
	case Int:
		a := inline.TakeIntJ(u, v)
    return a
	case Str:
		a := inline.TakeStrJ(u, v)
		i, err := strconv.ParseInt(a, 10, 64)
		if err != nil {
			panic(err)
		}
    return i
  }
	panic(F("Bad: 0x%x 0x%x", u, v))
}
func FnLenJ(u U, v V) int {
	//println("FnLenJ", u, v)
	switch inline.Tag(u, v) {
	case Py:
		p := macro.TakePJ(u, v)
		return p.Len()
	case Int:
		panic("Cannot take len() of int")
	case Str:
		return int(u) >> 3
	}
	panic(F("Bad: 0x%x 0x%x", u, v))
}

func FnNotContainsJ(uc U, vc V, ue U, ve V) bool {
  return !FnContainsJ(uc, vc, ue, ve)
}
func FnContainsJ(uc U, vc V, ue U, ve V) bool {
	//println("FnContainsJ", uc, vc, ue, ve)
	switch inline.Tag(uc, vc) {
	case Py:
		p := macro.TakePJ(uc, vc)
		return p.Contains(ue, ve)
	case Int:
		panic("Cannot FnContainsJ from int")
	case Str:
    s := inline.TakeStrJ(uc, vc)
    e := string(StrJ(ue, ve))
    return strings.Contains(s, e)
	}
	panic(F("Bad: 0x%x 0x%x", uc, vc))
}

func FnSetItemJ(c PJ, ui U, vi V, ue U, ve V) {
	//println("FnSetItemJ", ui, vi, ue, ve)
	switch inline.Tag(ui, vi) {
	case Py:
    a := inline.TakePJ(ui, vi)
		panic(F("SetItem: Bad index type: %s", a.TypeName()))
	case Int:
    a := inline.TakeIntJ(ui, vi)
		c.SetItemInt(int(a), ue, ve)
    return
	case Str:
    a := inline.TakeStrJ(ui, vi)
		c.SetItemStr(a, ue, ve)
    return
	}
	panic(F("Bad: 0x%x 0x%x", ui, vi))
}

func FnIsJ(u1 U, v1 V, u2 U, v2 V) bool {
  return u1 == u2 && v1 == v2
}
func FnIsNotJ(u1 U, v1 V, u2 U, v2 V) bool {
  return u1 != u2 || v1 != v2
}

func FnGetItemJ(uc U, vc V, ui U, vi V) (U,V) {
	//println("FnGetItemJ", uc, vc, ui, vi)
	switch inline.Tag(uc, vc) {
	case Py:
		p := macro.TakePJ(uc, vc)
	  switch inline.Tag(ui, vi) {
      case Py:
        panic("TODO")
      case Int:
        key := int(IntJ(ui, vi))
		    return p.GetItemInt(key)
      case Str:
        key := inline.TakeStrJ(ui, vi)
		    return p.GetItemStr(key)
    }
		panic(F("Bad: 0x%x 0x%x", ui, vi))
	case Int:
		panic("Cannot GetItem from int")
	case Str:
    s := inline.TakeStrJ(uc, vc)
    i := int(IntJ(ui, vi))
    if i < 0 {
      i += len(s)
    }
		z := s[i:i+1]
    return inline.MkStrJ(z)
	}
	panic(F("Bad: 0x%x 0x%x", uc, vc))
}

func FnGetItemSliceJ(a1 U, a2 V, b1 U, b2 V, c1 U, c2 V) (U, V) {
	//println("FnGetItemSliceJ", a1, a2, b2, b2, c1, c2)
	n := int64(macro.LenJ(a1, a2))

	// start point b
	b, bok := macro.IntOkOrNone(b1, b2)
	if !bok {
		b = 0
	}
	if b < 0 {
		b = n + b
	}

	// limit point c
	c, cok := macro.IntOkOrNone(c1, c2)
	if !cok {
		c = n
	}
	if c < 0 {
		c = n + c
	}
	if c > n {
		c = n // Allow trailing number too big.
	}

	switch a1 & 7 {
	case Py:
		p := macro.TakePJ(a1, a2)
		return p.GetItemSlice(int(b), int(c))
	case Int:
		panic("Cannot take slice of int")
	case Str:
		var s string = macro.TakeStrJ(a1, a2)
		//println("FnGetItemSliceJ", s, len(s), b, c)
		z := s[b:c]
		return macro.MkStrJ(z)
	default:
		panic("Bad")
	}
}

func FnDelItemSliceJ(a1 U, a2 V, b1 U, b2 V, c1 U, c2 V) {
	//println("FnGetItemSliceJ", a1, a2, b2, b2, c1, c2)
	n := int64(macro.LenJ(a1, a2))

	// start point b
	b, bok := macro.IntOkOrNone(b1, b2)
	if !bok {
		b = 0
	}
	if b < 0 {
		b = n + b
	}

	// limit point c
	c, cok := macro.IntOkOrNone(c1, c2)
	if !cok {
		c = n
	}
	if c < 0 {
		c = n + c
	}
	if c > n {
		c = n // Allow trailing number too big.
	}

	switch a1 & 7 {
	case Py:
		p := macro.TakePJ(a1, a2)
		p.DelItemSlice(int(b), int(c))
	case Int:
		panic("Cannot del slice of int")
	case Str:
		panic("cannot del slice of str")
	default:
		panic("Bad")
	}
}

//////////////////////////////////

type PJ interface {
	Hash() int64
	Pickle(w *bytes.Buffer)
	Show() string
	String() string
	Repr() string
	PType() (U, V)
	TypeName() string
	Callable() bool
	Is(a_1 U, a_2 V) bool
	IsNot(a_1 U, a_2 V) bool
	GetSelf() PJ
	SetSelf(a PJ)
	GetJBase() *JBase

	FetchField(field string) (U, V)
	StoreField(field string, p_1 U, p_2 V, q_1 U, q_2 V)

	Call(aa ...U) (U, V)
	Invoke(field string, aa ...W) (U, V)
	Iter() JNexter
	List() []W
	Dict() JScope

	Len() int

	SetItemInt(i int, x_1 U, x_2 V)
	SetItemStr(s string, x_1 U, x_2 V)

	DelItemInt(i int)
	DelItemStr(s string)
	DelItemSlice(b, c int)

	GetItemInt(i int) (U, V)
	GetItemStr(s string) (U, V)
	GetItemSlice(b, c int) (U, V)

	Contains(a_1 U, a_2 V) bool    // Reverse "in"

	Add_Py(b PJ) (U, V)
	Sub_Py(b PJ) (U, V)
	Mul_Py(b PJ) (U, V)
	Div_Py(b PJ) (U, V)
	Mod_Py(b PJ) (U, V)
	Pow_Py(b PJ) (U, V)

	Add_Int(b int64) (U, V)
	Sub_Int(b int64) (U, V)
	Mul_Int(b int64) (U, V)
	Div_Int(b int64) (U, V)
	Mod_Int(b int64) (U, V)
	Pow_Int(b int64) (U, V)

	Add_Str(b string) (U, V)
	Sub_Str(b string) (U, V)
	Mul_Str(b string) (U, V)
	Div_Str(b string) (U, V)
	Mod_Str(b string) (U, V)
	Pow_Str(b string) (U, V)

	Add_IntRev(b int64) (U, V)
	Sub_IntRev(b int64) (U, V)
	Mul_IntRev(b int64) (U, V)
	Div_IntRev(b int64) (U, V)
	Mod_IntRev(b int64) (U, V)
	Pow_IntRev(b int64) (U, V)

	Add_StrRev(b string) (U, V)
	Sub_StrRev(b string) (U, V)
	Mul_StrRev(b string) (U, V)
	Div_StrRev(b string) (U, V)
	Mod_StrRev(b string) (U, V)
	Pow_StrRev(b string) (U, V)

	BitAnd(b int64) (U, V)
	BitOr(b int64) (U, V)
	BitXor(b int64) (U, V)
	ShiftLeft(b int64) (U, V)
	ShiftRight(b int64) (U, V)
	UnsignedShiftRight(b int64) (U, V)

	Bool() bool // a.k.a. nonzero()
	UnaryMinus() (U, V)
	UnaryPlus() (U, V)
	UnaryInvert() (U, V)

	Compare_Py(b PJ) int
	Compare_Int(b int64) int
	Compare_Str(b string) int

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
	Object() *J_object
	Superclass() (U, V)
}

type JBase struct {
	Self PJ
}

type JByt struct {
	JBase
	YY []byte
}

type JBool struct {
	JBase
	Z bool
	I int64
	F float64
}

type JFloat struct {
	JBase
	F float64
}

type JListIter struct {
	JBase
	PP []W
	I  int
}

type JGo struct {
	JBase
	V R.Value
}

type JList struct {
	JBase
	PP []W
}

type JTuple struct {
	JList
}

type JNone struct {
	JBase
}

type JScope map[string]W

type JDict struct {
	JBase
	ppp JScope
	mu  sync.Mutex
}
type JNexter interface {
	NextJ() (U, V, bool)
}

func (o *JListIter) Iter() JNexter {
	return o
}

func (o *JListIter) NextJ() (U, V, bool) {
	if o.I < len(o.PP) {
		z := o.PP[o.I]
		z1, z2 := z.U, z.V
		o.I++
		return z1, z2, true
	}
	return 0, V0, false
}

func (o *JBase) GetJBase() *JBase { return o }
func (o *JBase) GetSelf() PJ      { return o.Self }
func (o *JBase) SetSelf(a PJ) {
	o.Self = a
}

func (g *JBase) FetchField(field string) (U, V) {
	// Try using PGO reflection.
	//TODO return FetchFieldByNameForObject(R.ValueOf(g.Self), field)
	return 0, V0
}

func (o *JBase) StoreField(field string, p_1 U, p_2 V, q_1 U, q_2 V) {
	panic(F("Receiver %T cannot StoreField", o.Self))
}

func (o *JBase) Call(aa ...U) (U, V) { panic(F("Receiver %T cannot Call", o.Self)) }
func (o *JBase) Invoke(field string, aa ...W) (U, V) {
	panic(F("Receiver %T cannot invoke", o.Self))
}
func (o *JBase) Len() int { panic(F("Receiver %T cannot Len: ", o.Self)) }

func (o *JBase) SetItemInt(i int, x_1 U, x_2 V) {
	panic(F("Receiver %T cannot SetItemInt", o.Self))
}
func (o *JBase) SetItemStr(s string, x_1 U, x_2 V) {
	panic(F("Receiver %T cannot SetItemStr", o.Self))
}
func (o *JBase) DelItemInt(i int) {
	panic(F("Receiver %T cannot DelItemInt", o.Self))
}
func (o *JBase) DelItemStr(s string) {
	panic(F("Receiver %T cannot DelItemStr", o.Self))
}
func (o *JBase) DelItemSlice(b, c int) {
	panic(F("Receiver %T cannot DelItemSlice", o.Self))
}
func (o *JBase) GetItemInt(i int) (U, V) {
	panic(F("Receiver %T cannot GetItemInt", o.Self))
}
func (o *JBase) GetItemStr(s string) (U, V) {
	panic(F("Receiver %T cannot GetItemStr", o.Self))
}
func (o *JBase) GetItemSlice(b, c int) (U, V) {
	panic(F("Receiver %T cannot GetItemSlice", o.Self))
}

func (o *JBase) Is(a_1 U, a_2 V) bool {
	var o1 U
	var o2 V
	o1, o2 = macro.MkPJ(o)
	return o1 == a_1 && o2 == a_2
}
func (o *JBase) IsNot(a_1 U, a_2 V) bool {
	var o1 U
	var o2 V
	o1, o2 = macro.MkPJ(o)
	return o1 != a_1 || o2 != a_2
}

//func (o *PIntJ) Is(a_1 U, a_2 V) bool     { return o.N == a.N }
//func (o *PIntJ) IsNot(a_1 U, a_2 V) bool  { return o.N != a.N }
//func (o *PStrJ) Is(a_1 U, a_2 V) bool     { return o.S == a.S }
//func (o *PStrJ) IsNot(a_1 U, a_2 V) bool  { return o.S != a.S }

func (o *JBase) Contains(a_1 U, a_2 V) bool         { panic(F("Receiver %T cannot Contains: ", o.Self)) }
//func (o *JBase) SetItem(i_1 U, i_2 V, x_1 U, x_2 V) { panic(F("Receiver %T cannot SetItem: ", o.Self)) }
func (o *JBase) DelItem(i_1 U, i_2 V)               { panic(F("Receiver %T cannot DelItem: ", o.Self)) }
func (o *JBase) Iter() JNexter                      { panic(F("Receiver %T cannot Iter: ", o.Self)) }
func (o *JBase) List() []W                          { panic(F("Receiver %T cannot List: ", o.Self)) }
func (o *JBase) Dict() JScope                       { panic(F("Receiver %T cannot Dict: ", o.Self)) }

func (o *JBase) Add_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Add %T", o.Self, b.GetSelf())) }
func (o *JBase) Sub_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Sub %T", o.Self, b.GetSelf())) }
func (o *JBase) Mul_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Mul %T", o.Self, b.GetSelf())) }
func (o *JBase) Div_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Div %T", o.Self, b.GetSelf())) }
func (o *JBase) Mod_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Mod %T", o.Self, b.GetSelf())) }
func (o *JBase) Pow_Py(b PJ) (U, V) { panic(F("Receiver %T cannot Pow %T", o.Self, b.GetSelf())) }

func (o *JBase) Add_Int(b int64) (U, V) { panic(F("Receiver %T cannot Add int", o.Self)) }
func (o *JBase) Sub_Int(b int64) (U, V) { panic(F("Receiver %T cannot Sub int", o.Self)) }
func (o *JBase) Mul_Int(b int64) (U, V) { panic(F("Receiver %T cannot Mul int", o.Self)) }
func (o *JBase) Div_Int(b int64) (U, V) { panic(F("Receiver %T cannot Div int", o.Self)) }
func (o *JBase) Mod_Int(b int64) (U, V) { panic(F("Receiver %T cannot Mod int", o.Self)) }
func (o *JBase) Pow_Int(b int64) (U, V) { panic(F("Receiver %T cannot Pow int", o.Self)) }

func (o *JBase) Add_Str(b string) (U, V) { panic(F("Receiver %T cannot Add str", o.Self)) }
func (o *JBase) Sub_Str(b string) (U, V) { panic(F("Receiver %T cannot Sub str", o.Self)) }
func (o *JBase) Mul_Str(b string) (U, V) { panic(F("Receiver %T cannot Mul str", o.Self)) }
func (o *JBase) Div_Str(b string) (U, V) { panic(F("Receiver %T cannot Div str", o.Self)) }
func (o *JBase) Mod_Str(b string) (U, V) { panic(F("Receiver %T cannot Mod str", o.Self)) }
func (o *JBase) Pow_Str(b string) (U, V) { panic(F("Receiver %T cannot Pow str", o.Self)) }

func (o *JBase) Add_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot AddRev int", o.Self)) }
func (o *JBase) Sub_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot SubRev int", o.Self)) }
func (o *JBase) Mul_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot MulRev int", o.Self)) }
func (o *JBase) Div_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot DivRev int", o.Self)) }
func (o *JBase) Mod_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot ModRev int", o.Self)) }
func (o *JBase) Pow_IntRev(b int64) (U, V) { panic(F("Receiver %T cannot PowRev int", o.Self)) }

func (o *JBase) Add_StrRev(b string) (U, V) { panic(F("Receiver %T cannot AddRev str", o.Self)) }
func (o *JBase) Sub_StrRev(b string) (U, V) { panic(F("Receiver %T cannot SubRev str", o.Self)) }
func (o *JBase) Mul_StrRev(b string) (U, V) { panic(F("Receiver %T cannot MulRev str", o.Self)) }
func (o *JBase) Div_StrRev(b string) (U, V) { panic(F("Receiver %T cannot DivRev str", o.Self)) }
func (o *JBase) Mod_StrRev(b string) (U, V) { panic(F("Receiver %T cannot ModRev str", o.Self)) }
func (o *JBase) Pow_StrRev(b string) (U, V) { panic(F("Receiver %T cannot PowRev str", o.Self)) }

func (o *JBase) BitAnd(b int64) (U, V)     { panic(F("Receiver %T cannot BitAnd: ", o.Self)) }
func (o *JBase) BitOr(b int64) (U, V)      { panic(F("Receiver %T cannot BitOr: ", o.Self)) }
func (o *JBase) BitXor(b int64) (U, V)     { panic(F("Receiver %T cannot BitXor: ", o.Self)) }
func (o *JBase) ShiftLeft(b int64) (U, V)  { panic(F("Receiver %T cannot ShiftLeft: ", o.Self)) }
func (o *JBase) ShiftRight(b int64) (U, V) { panic(F("Receiver %T cannot ShiftRight: ", o.Self)) }
func (o *JBase) UnsignedShiftRight(b int64) (U, V) {
	panic(F("Receiver %T cannot UnsignedShiftRight: ", o.Self))
}

// The Order of things in Rye:  objects < numbers < strings
// TODO: make numbers sort together; make byt sort with str.
func (o *JBase) Compare_Int(b int64) int {
	return -1
}
func (o *JBase) Compare_Str(b string) int {
	return -1
}
func (o *JBase) Compare_Py(b PJ) int {
	//bp := b.GetJBase()
	//x := R.ValueOf(o).Pointer()
	//y := R.ValueOf(bp).Pointer()
  x := uintptr(unsafe.Pointer(o))
  y := uintptr(unsafe.Pointer(b.GetJBase()))
	return macro.OperatorCompare(x, y)
}
func (o *JBase) Bool() bool          { return true } // Most things are true.
func (o *JBase) UnaryMinus() (U, V)  { panic(F("Receiver %T cannot UnaryMinus", o.Self)) }
func (o *JBase) UnaryPlus() (U, V)   { panic(F("Receiver %T cannot UnaryPlus", o.Self)) }
func (o *JBase) UnaryInvert() (U, V) { panic(F("Receiver %T cannot UnaryInvert", o.Self)) }

func (o *JBase) CanStr() bool { return false }
func (o *JBase) Str() string  { panic(F("Receiver %T cannot Str", o.Self)) }

func (o *JBase) CanInt() bool    { return false }
func (o *JBase) Int() int64      { panic(F("Receiver %T cannot Int", o.Self)) }
func (o *JBase) ForceInt() int64 { panic(F("Receiver %T cannot ForceInt", o.Self)) }

func (o *JBase) CanFloat() bool        { return false }
func (o *JBase) Float() float64        { panic(F("Receiver %T cannot Float", o.Self)) }
func (o *JBase) ForceFloat() float64   { panic(F("Receiver %T cannot ForceFloat", o.Self)) }
func (o *JBase) Contents() interface{} { return o.Self }

func (o *JBase) Superclass() (U, V) { panic("todo") }
func (o *JBase) Object() *J_object  { return nil }
func (o *JBase) Callable() bool     { return false }
func (o *JBase) PType() (U, V) {
	s := F("%T", o.Self)
	return macro.MkStrJ(s)
}
func (o *JBase) TypeName() string { panic(F("Receiver %T cannot TypeName", o.Self)) }
func (o *JBase) Bytes() []byte    { panic(F("Receiver %T cannot Bytes", o.Self)) }
func (o *JBase) String() string {
	if o.Self == nil {
		panic("JBase: Self is nil")
	}
	return o.Self.Show()
}
func (o *JBase) Hash() int64 { return int64(R.ValueOf(o).Pointer()) }

const ShortHashModulusJ = 99999 // 3 * 3 * 41 * 271
func (o *JBase) ShortPointerHashString() string {
	return F("@%05d", o.ShortPointerHash())
}
func (o *JBase) ShortPointerHash() int {
	val := R.ValueOf(o.Self)
	return 1 + int(val.Pointer())%ShortHashModulusJ
}

func (o *JBase) Pickle(w *bytes.Buffer) { panic(F("Receiver %T cannot Pickle", o.Self)) }
func (o *JBase) Repr() string           { return o.Self.String() }
func (o *JBase) Show() string {
	panic("todo")
	/*
		if o.Self == nil {
			panic("OHNO: o.Self == nil")
		}
		return ShowP(o.Self, SHOW_DEPTH)
	*/
}

// Because go confuses empty lists with nil, rye does the same with None.
// This saves you writing `for x in vec if vec else []:`
func (o *JNone) TypeName() string              { return "NoneType" }
func (o *JNone) Hash() int64                   { return 23 }
func (o *JNone) Len() int                      { return 0 }
func (o *JNone) Contains(a_1 U, a_2 V) bool   { return false }
func (o *JNone) List() []W                     { return nil }
func (o *JNone) Dict() JScope                  { return make(JScope) }
func (o *JNone) Iter() JNexter {
	z := &JListIter{PP: nil}
	ForgeJ(z)
	return z
}

func (o *JNone) Bool() bool            { return false }
func (o *JNone) String() string        { return "None" }
func (o *JNone) Repr() string          { return "None" }
func (o *JNone) Contents() interface{} { return nil }

func (o *JNone) Pickle(w *bytes.Buffer) { w.WriteByte(RypNone) }

///////////////////////////////////////////////////////////////////

func (o *JBool) TypeName() string { return "bool" }
func (o *JBool) Hash() int64      { return o.Int() }
func (o *JBool) Pickle(w *bytes.Buffer) {
	if o.Z {
		w.WriteByte(RypTrue)
	} else {
		w.WriteByte(RypFalse)
	}
}
func (o *JBool) Contents() interface{} { return o.Z }
func (o *JBool) Bool() bool            { return o.Z }
func (o *JBool) CanInt() bool          { return true }
func (o *JBool) ForceInt() int64       { return o.Int() }
func (o *JBool) Int() int64 {
	if o.Z {
		return 1
	} else {
		return 0
	}
}
func (o *JBool) CanFloat() bool      { return true }
func (o *JBool) ForceFloat() float64 { return o.Float() }
func (o *JBool) Float() float64 {
	if o.Z {
		return 1.0
	} else {
		return 0.0
	}
}
func (o *JBool) String() string {
	if o.Z {
		return "True"
	} else {
		return "False"
	}
}
func (o *JBool) Repr() string {
	if o.Z {
		return "True"
	} else {
		return "False"
	}
}
func (o *JBool) PType() (U, V) { panic("todo") } // return G_bool
func (o *JBool) Compare_Py(b PJ) int {
	// Bool sorts like a number.
	switch t := b.(type) {
	case *JBool:
		return macro.OperatorCompare(o.I, t.I)
	case *JFloat:
		return macro.OperatorCompare(o.F, t.F)
	case *JByt: // byt sorts like str.
		return -1 // Numbers are less than strings.
	}
	return 1 // Numbers are greater than objects.
}

//////////////////////////////////////////

func (o *JFloat) TypeName() string { return "float" }
func (o *JFloat) Hash() int64      { return o.Int() }
func (o *JFloat) Pickle(w *bytes.Buffer) {
	x := int64(math.Float64bits(o.F))
	n := RypIntLenMinus1(x)
	w.WriteByte(byte(RypFloat + n))
	RypWriteInt(w, x)
}
func (o *JFloat) Contents() interface{} { return o.F }
func (o *JFloat) Bool() bool            { return o.F != 0.0 }
func (o *JFloat) CanInt() bool          { return false }
func (o *JFloat) ForceInt() int64       { return o.Int() }
func (o *JFloat) Int() int64            { return int64(o.F) }
func (o *JFloat) CanFloat() bool        { return true }
func (o *JFloat) ForceFloat() float64   { return o.F }
func (o *JFloat) Float() float64        { return o.F }
func (o *JFloat) String() string {
	{
		return strconv.FormatFloat(o.F, 'g', -1, 64)
	}
}
func (o *JFloat) Repr() string {
	{
		return strconv.FormatFloat(o.F, 'g', -1, 64)
	}
}
func (o *JFloat) Compare_Int(b int64) int {
  f := float64(b)
  return macro.OperatorCompare(o.F, f)
}
func (o *JFloat) Compare_Str(string) int {
  return -1 // Numbers < Str
}
func (o *JFloat) Compare_Py(b PJ) int {
	// Float sorts like a number.
	switch t := b.(type) {
	case *JBool:
		return macro.OperatorCompare(o.F, t.F)
	case *JFloat:
		z := macro.OperatorCompare(o.F, t.F)
		//log.Printf("JFloat::Compare_Py ---- %g %g -> %d", o.F, t.F, z)
		return z
	case *JByt: // byt sorts like str.
		return -2 // Numbers are less than strings.
	}
	return 9 // Numbers are greater than objects.
}

////////////////////////////////

type JCallSpec struct {
	Name     string
	Args     []string
	Defaults []W
	Star     string
	StarStar string
}

type JPNewCallable struct {
	JBase
	JCallSpec *JCallSpec
}

func (o *JPNewCallable) Callable() bool { return true }

func (o *JPNewCallable) String() string {
	return fmt.Sprintf("<func %s>", o.JCallSpec.Name)
}

func (o *JPNewCallable) Repr() string {
	return o.String()
}

/*
// Old XPCallableJ

type XPCallableJ struct {
	JBase
	Name     string
	Args     []string
	Defaults []W
	Star     string
	StarStar string
}

func (o *XPCallableJ) Callable() bool { return true }

func (o *XPCallableJ) String() string {
	return fmt.Sprintf("<func %s>", o.Name)
}

func (o *XPCallableJ) Repr() string {
	return o.String()
}
*/

////////////////////////////////

// Proto1J is ...
func Proto1J(a_1 U, a_2 V) (U, V) {
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("Proto1J: got a nil")
		}
		a := macro.TakePJ(a_1, a_2)
		_ = a
		panic("Proto1J: on Py & Py")
	case Int:
		a := macro.TakeIntJ(a_1, a_2)
		_ = a
		panic("Proto1J: on Int & Py")
	case Str:
		a := macro.TakeStrJ(a_1, a_2)
		_ = a
		panic("Proto1J: on Str & Py")
	}
	panic("Proto1J: WTF")
}

// Proto2J is ...
func Proto2J(a_1 U, a_2 V, b_1 U, b_2 V) (U, V) {
	if b_1 == 0 {
		panic("Proto2J: on _ & nil")
	}
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("Proto2J: on nil")
		}
		a := macro.TakePJ(a_1, a_2)
		_ = a
		switch b_1 & 7 {
		case Py:
			panic("Proto2J: on Py & Py")
		case Int:
			panic("Proto2J: on Py & Int")
		case Str:
			Wrong()
			panic("Proto2J: on Py & Str")
		default:
			panic("Proto2J: on Py & WTF")
		}
	case Int:
		a := macro.TakeIntJ(a_1, a_2)
		_ = a
		switch b_1 & 7 {
		case Py:
			panic("Proto2J: on Int & Py")
		case Int:
			panic("Proto2J: on Int & Int")
		case Str:
			Wrong()
			panic("Proto2J: on Int & Str")
		default:
			panic("Proto2J: on Int & WTF")
		}
	case Str:
		a := macro.TakeStrJ(a_1, a_2)
		_ = a
		switch b_1 & 7 {
		case Py:
			panic("Proto2J: on Str & Py")
		case Int:
			panic("Proto2J: on Str & Int")
		case Str:
			Wrong()
			panic("Proto2J: on Str & Str")
		default:
			panic("Proto2J: on Str & WTF")
		}
	}
	panic("Proto2J: WTF")
}

////////////////////////////////

// BoolJ is ...
func BoolJ(a_1 U, a_2 V) bool {
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("BoolJ: got a nil")
		}
		a := macro.TakePJ(a_1, a_2)
		return a.Bool()
	case Int:
		a := macro.TakeIntJ(a_1, a_2)
		return a != 0
	case Str:
		a := macro.TakeStrJ(a_1, a_2)
		return len(a) != 0
	}
	panic("BoolJ: WTF")
}

// IntJ is ...
func IntJ(a_1 U, a_2 V) int64 {
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("IntJ: got a nil")
		}
		a := macro.TakePJ(a_1, a_2)
		return a.Int()
		panic("IntJ: on Py & Py")
	case Int:
		a := macro.TakeIntJ(a_1, a_2)
		return a
	case Str:
		a := macro.TakeStrJ(a_1, a_2)
		_ = a
		panic("IntJ: on Str & Py")
	}
	panic("IntJ: WTF")
}

// StrJ is the weak string converter.
func StrJ(a_1 U, a_2 V) string {
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("StrJ: got a nil")
		}
		s := macro.TakePJ(a_1, a_2)
		return s.Str()
	case Int:
		panic("StrJ: got an Int")
	case Str:
		return macro.TakeStrJ(a_1, a_2)
	}
	panic("StrJ: WTF")
}

// StringJ is ...
func StringJ(a_1 U, a_2 V) string {
	switch a_1 & 7 {
	case Py:
		if a_1 == 0 {
			panic("StringJ: got a nil")
		}
		a := macro.TakePJ(a_1, a_2)
		return a.String()
	case Int:
		a := macro.TakeIntJ(a_1, a_2)
		return strconv.FormatInt(a, 10)
	case Str:
		return macro.TakeStrJ(a_1, a_2)
	}
	panic("StringJ: WTF")
}

// ReprJ is python's repr().
func ReprJ(u U, v V) string {
	switch u & 7 {
	case Py:
		if u == 0 {
			panic("ReprJ: got a nil")
		}
		a := macro.TakePJ(u, v)
		return a.Repr()
	case Int:
		a := macro.TakeIntJ(u, v)
		return strconv.FormatInt(a, 10)
	case Str:
		a := macro.TakeStrJ(u, v)
		return fmt.Sprintf("%q", a)
	}
	panic("ReprJ: WTF")
}

type EitherJ struct {
	Left   interface{}
	Valid  bool
	Right1 U
	Right2 V
}

type void int

// J_generator is the channel begin a yielding producer and a consuming for loop.
type J_generator struct {
	J_object
	Ready    chan *void
	Result   chan EitherJ
	Finished bool
}

func NewJGenerator() *J_generator {
	z := &J_generator{
		Ready:  make(chan *void, GENERATOR_BUF_SIZE),
		Result: make(chan EitherJ, GENERATOR_BUF_SIZE),
	}
	z.SetSelf(z)
	// Signal the coroutine so it can run asynchronously.
	for i := 0; i < GENERATOR_BUF_SIZE; i++ {
		z.Ready <- nil
	}
	return z
}

func (o *J_generator) PtrJ_generator() *J_generator {
	return o
}

func (o *J_generator) Iter() JNexter { return o }
func (o *J_generator) List() []W {
	var z []W
	for {
		x1, x2, gotOne := o.NextJ()
		if !gotOne {
			break
		}
		z = append(z, W{x1, x2})
	}
	o.Enough()
	return z
}

// NextJ is called by the consumer.
// NextJ waits for next result from the generator.
// It returns either a result of type (U,V) and true,
// or if there are no more, it returns (0,V0) and false.
// If the generator goroutine died on an exception,
// that exception gets rethrown here.
func (o *J_generator) NextJ() (U, V, bool) {
	o.Ready <- nil
	// That wakes up the generator goroutine.
	// Now we block, waiting on next result.
	either, ok := <-o.Result

	if !ok {
		return 0, V0, false
	}
	if either.Left != nil {
		if o.Ready != nil {
			close(o.Ready)
			o.Ready = nil
		}
		panic(either.Left)
	}
	return either.Right1, either.Right2, true
}

// Enough is called by the consumer, to tell the producer to stop because we've got enough.
func (o *J_generator) Enough() {
	if o.Ready != nil {
		close(o.Ready)
		o.Ready = nil
	}
}

// Yield is called by the producer, to yield a value to the consumer.
func (o *J_generator) Yield(item_1 U, item_2 V) {
	o.Result <- EitherJ{Left: nil, Valid: true, Right1: item_1, Right2: item_2}
}

// Yield is called by the producer when it catches an exception, to yield it to the producer (as an EitherJ Left).
func (o *J_generator) YieldException(ex interface{}) {
	o.Result <- EitherJ{Left: ex, Valid: true, Right1: 0, Right2: V0}
}

// Finish is called by the producer when it is finished.
func (o *J_generator) Finish() {
	if !o.Finished {
		o.Finished = true
		close(o.Result)
	}
}

// Wait is called by the producer, once at the start, and once after each yield.
// Wait returns false if the consumer said Enough.
// TODO:  Don't wait, to achieve concurrency.  Let the user decide the Result channel buffer size.
func (o *J_generator) Wait() bool {
	_, ok := <-o.Ready
	return ok
}

// J_object is the root of inherited classes.
type J_object struct {
	JBase
}

func (o *J_object) Object() *J_object {
	return o
}
func (o *J_object) PtrJ_object() *J_object {
	return o
}

type PtrJ_object_er interface {
	PtrJ_object() *J_object
}

func (o *J_object) PType() (U, V) {
	panic("return xacro.MkPJ(nil)")
}
func (o *J_object) TypeName() string { return "object" }

func (o *J_object) Repr() string {
	x1, x2 := o.Self.(ij__repr__).M_0___repr__()
	return StrJ(x1, x2)
}
func (o *J_object) String() string {
	x1, x2 := o.Self.(ij__str__).M_0___str__()
	return StrJ(x1, x2)
}
func (o *J_object) M_0___str__() (U, V) {
	val := R.ValueOf(o.Self)
	cname := val.Type().Elem().Name()
	if strings.HasPrefix(cname, "J_") {
		cname = cname[2:] // Demangle.
	}
	z := F("<%s%s>", cname, o.ShortPointerHashString())
	return macro.MkStrJ(z)
}
func (o *J_object) M_0___repr__() (U, V) {
	// Todo: use ShowP(o.Self, SHOW_DEPTH)
	z := fmt.Sprintf("%#v", o.Self)
	return macro.MkStrJ(z)
}

type ij__str__ interface {
	M_0___str__() (U, V)
}
type ij__repr__ interface {
	M_0___repr__() (U, V)
}

func (o *J_object) PickleFields(w *bytes.Buffer, v R.Value) {
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
			v.Field(i).Interface().(PJ).Pickle(w)
		}
	}
}

func (o *J_object) Pickle(w *bytes.Buffer) {
	w.WriteByte(RypClass)
	RypWriteLabel(w, R.ValueOf(o.Self).Type().Elem().String())
	o.PickleFields(w, R.ValueOf(o.Self).Elem())
	w.WriteByte(0) // Like a 0-length label, to terminate fields.
}

func JNewSpecCall(cs *JCallSpec, a1 []W, a2 []W, kv []JKV, kv2 map[string]W) ([]W, *JList, *JDict) {
	n := len(cs.Defaults)
  //log.Printf("JNewSpecCall n=%d cs.Defaults=%v a1=%v", n, cs.Defaults, a1)
  //log.Printf("cs=%#v", cs)
	argv := make([]W, n)
	var star []W
	var starstar map[string]W

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
		u := e.U
		v := e.V
		stored := false
		for ni, ne := range cs.Args { // O(n^2), probably not a problem.
			if k == ne {
				argv[ni] = W{u, v}
				stored = true
				break
			}
		}
		if !stored {
			if starstar == nil {
				starstar = make(map[string]W)
			}
			starstar[k] = W{u, v}
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
				starstar = make(map[string]W)
			}
			starstar[k] = v
		}
	}

	for i := 0; i < len(argv); i++ {
		if argv[i] == (W{0, V0}) {
			panic(F("The %dth fixed formal argument '%s' of function %q has no assigned value (fixed formal args are: %v)", i+1, cs.Args[i], cs.Name, cs.Args))
		}
	}

	if cs.Star == "" && len(star) > 0 {
		panic(F("Function %q wants %d args, but got %d (no * arg)", cs.Name, len(cs.Args), len(cs.Args)+len(star)))
	}

	if cs.StarStar == "" && len(starstar) > 0 {
		panic(F("Function %q cannot take %d extra named args", cs.Name, len(starstar)))
	}

	return argv, MkListJ(star), MkDictJ(JScope(starstar))
}

////////////////////////////
//
//  Reusable Callers.

// PCall0

type JCall0 struct {
	JPNewCallable
	Fn   func() (U, V)
}

func (o *JCall0) Contents() interface{} {
	return o.Fn
}
func (o JCall0) JCall0() (U, V) {
	return o.Fn()
}

func (o JCall0) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U, V) {
	argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn()
}

// JCall1

type JCall1 struct {
	JPNewCallable
	Fn   func(a0_1 U, a0_2 V) (U, V)
}

func (o *JCall1) Contents() interface{} {
	return o.Fn
}
func (o JCall1) JCall1(a0_1 U, a0_2 V) (U, V) {
  //log.Printf("JCall1 FAST: %q", o.JPNewCallable.JCallSpec.Name)
	return o.Fn(a0_1, a0_2)
}

func (o JCall1) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U, V) {
	argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
  //log.Printf("JCall1 SLOW: %q -- %d %d %d %d", o.JPNewCallable.JCallSpec.Name, len(a1), len(a2), len(kv1), len(kv2))
	return o.Fn(argv[0].U, argv[0].V)
}

// JCall2

type JCall2 struct {
	JPNewCallable
	Fn   func(a0_1 U, a0_2 V, a1_1 U, a1_2 V) (U, V)
}

func (o *JCall2) Contents() interface{} {
	return o.Fn
}
func (o JCall2) JCall2(a0_1 U, a0_2 V, a1_1 U, a1_2 V) (U, V) {
	return o.Fn(a0_1, a0_2, a1_1, a1_2)
}

func (o JCall2) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U, V) {
	argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn(argv[0].U, argv[0].V, argv[1].U, argv[1].V)
}

// JCall3

type JCall3 struct {
	JPNewCallable
	Fn   func(a0_1 U, a0_2 V, a1_1 U, a1_2 V, a2_1 U, a2_2 V) (U, V)
}

func (o *JCall3) Contents() interface{} {
	return o.Fn
}
func (o JCall3) JCall3(a0_1 U, a0_2 V, a1_1 U, a1_2 V, a2_1 U, a2_2 V) (U, V) {
	return o.Fn(a0_1, a0_2, a1_1, a1_2, a2_1, a2_2)
}

func (o JCall3) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U, V) {
	argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.Fn(argv[0].U, argv[0].V, argv[1].U, argv[1].V, argv[2].U, argv[2].V)
}

///////////////////

type IJCallV interface {
	JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U, V)
}

func MkTupleJ(pp []W) *JTuple {
	z := &JTuple{JList{PP: pp}}
	ForgeJ(z)
	return z
}
func MkTupleVUV(pp ...W) (U, V) {
	z := &JTuple{JList{PP: pp}}
	ForgeJ(z)
	return inline.MkPJ(z)
}

func MkListJ(pp []W) *JList {
	z := &JList{PP: pp}
	ForgeJ(z)
	return z
}
func MkListVUV(pp ...W) (U, V) {
	z := &JList{PP: pp}
	ForgeJ(z)
	return inline.MkPJ(z)
}

func MkDictJ(ppp JScope) *JDict {
	z := &JDict{ppp: ppp}
	ForgeJ(z)
	return z
}
func MkDictVJ(pp ...W) (U,V) {
	if (len(pp) % 2) == 1 {
		panic("MkDictV got odd len(pp)")
	}
	zzz := make(JScope)
	for i := 0; i < len(pp); i += 2 {
    k := pp[i]
    key := StringJ(k.U, k.V)
		zzz[key] = pp[i+1]
	}
	z := &JDict{ppp: zzz}
	ForgeJ(z)
  return inline.MkPJ(z)
}

func MkStrsJ(ss []string) *JList {
	pp := make([]W, len(ss))
	for i, s := range ss {
		u1, u2 := macro.MkStrJ(s)
		pp[i] = W{u1, u2}
	}
	return MkListJ(pp)
}

func MkBytUV(bb []byte) (U,V) {
  z := &JByt{YY: bb}
  ForgeJ(z)
  return inline.MkPJ(z)
}

func BytesJ(u U, v V) []byte {
  panic("TODO")
}

////////////////////////////////////////

type BinopWorkerBase struct {
	Name string
}

func (op *BinopWorkerBase) PyPy(a PJ, b PJ) (U, V) {
	panic(F("operation %s not defined on %s & %s", op.Name, a.TypeName(), b.TypeName()))
}
func (op *BinopWorkerBase) PyInt(a PJ, b int64) (U, V) {
	panic(F("operation %s not defined on %s & int", op.Name, a.TypeName()))
}
func (op *BinopWorkerBase) PyStr(a PJ, b string) (U, V) {
	panic(F("operation %s not defined on %s & str", op.Name, a.TypeName()))
}
func (op *BinopWorkerBase) IntPy(a int64, b PJ) (U, V) {
	panic(F("operation %s not defined on int & %s", op.Name, b.TypeName()))
}
func (op *BinopWorkerBase) IntInt(a int64, b int64) (U, V) {
	panic(F("operation %s not defined on int & int", op.Name))
}
func (op *BinopWorkerBase) IntStr(a int64, b string) (U, V) {
	panic(F("operation %s not defined on int & str", op.Name))
}
func (op *BinopWorkerBase) StrPy(a string, b PJ) (U, V) {
	panic(F("operation %s not defined on str & %s", op.Name, b.TypeName()))
}
func (op *BinopWorkerBase) StrInt(a string, b int64) (U, V) {
	panic(F("operation %s not defined on str & int", op.Name))
}
func (op *BinopWorkerBase) StrStr(a string, b string) (U, V) {
	panic(F("operation %s not defined on str & str", op.Name))
}

type RelopWorkerBase struct {
	Name string
}

func (op *RelopWorkerBase) PyPy(a PJ, b PJ) bool {
	panic(F("operation %s not defined on %s & %s", op.Name, a.TypeName(), b.TypeName()))
}
func (op *RelopWorkerBase) PyInt(a PJ, b int64) bool {
	panic(F("operation %s not defined on %s & int", op.Name, a.TypeName()))
}
func (op *RelopWorkerBase) PyStr(a PJ, b string) bool {
	panic(F("operation %s not defined on %s & str", op.Name, a.TypeName()))
}
func (op *RelopWorkerBase) IntPy(a int64, b PJ) bool {
	panic(F("operation %s not defined on int & %s", op.Name, b.TypeName()))
}
func (op *RelopWorkerBase) IntInt(a int64, b int64) bool {
	panic(F("operation %s not defined on int & int", op.Name))
}
func (op *RelopWorkerBase) IntStr(a int64, b string) bool {
	panic(F("operation %s not defined on int & str", op.Name))
}
func (op *RelopWorkerBase) StrPy(a string, b PJ) bool {
	panic(F("operation %s not defined on str & %s", op.Name, b.TypeName()))
}
func (op *RelopWorkerBase) StrInt(a string, b int64) bool {
	panic(F("operation %s not defined on str & int", op.Name))
}
func (op *RelopWorkerBase) StrStr(a string, b string) bool {
	panic(F("operation %s not defined on str & str", op.Name))
}

///////////////////////
func FnCompare(a1 U, a2 V, b1 U, b2 V) int {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Compare_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Compare_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Compare_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return -b.Compare_Int(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return inline.OperatorCompare(a, b)
		case Str:
			// Int < Str
			return -1
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return -b.Compare_Str(a)
		case Int:
			// Str > Int
			return 1
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return inline.OperatorCompare(a, b)
		}
	}
	panic(F("Bad switch FnCompare: %x %x %x %x", a1, a2, b1, b2))
}

///////////////////////

// #### JList

func ContentsJ(u U, v V) interface{} {
	switch inline.Tag(u, v) {
	case Py:
		p := inline.TakePJ(u, v)
		return p.Contents()
	case Int:
		return TakeIntJ(u, v)
	case Str:
		return TakeStrJ(u, v)
	}
	panic("bad TypeNameJ")
}

func TypeNameJ(u U, v V) string {
	switch inline.Tag(u, v) {
	case Py:
		p := inline.TakePJ(u, v)
		return p.TypeName()
	case Int:
		return "int"
	case Str:
		return "str"
	}
	panic("bad TypeNameJ")
}

func ListJ(u U, v V) []W {
	switch inline.Tag(u, v) {
	case Py:
		p := inline.TakePJ(u, v)
		return p.List()
	case Str:
		var ww []W
		s := inline.TakeStrJ(u, v)
		for _, ch := range s {
			u9, v9 := inline.MkintJ(int(ch))
			ww = append(ww, W{u9, v9})
		}
		return ww
	}
	panic(F("Cannot ListJ on %s", TypeNameJ(u, v)))
}

func StrHash(s string) int64 { return int64(crc64.Checksum([]byte(s), CrcPolynomial)) }

func (o *JList) Hash() int64 {
	var z int64
	for _, e := range o.PP {
		// TODO: https://godoc.org/github.com/leemcloughlin/gofarmhash ?
		z += inline.HashJ(e.U, e.V)
	}
	return z
}

func (o *JList) Compare_Py(a PJ) int {
	switch b := a.(type) {
	case *JList:
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
					oppi := o.PP[i]
					bppi := b.PP[i]
					cmp := inline.CompareJ(oppi.U, oppi.V, bppi.U, bppi.V)
					if cmp != 0 {
						return cmp
					}
					// But if the are equal, continue with next slot.
				}
			}
		}
	}
	return StrCmp(o.TypeName(), a.TypeName())
}

func PickleJ(u U, v V, w *bytes.Buffer) {
	panic("TODO")
}

func (o *JTuple) Pickle(w *bytes.Buffer) {
	l := int64(len(o.PP))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypTuple + n))
	RypWriteInt(w, l)
	for _, x := range o.PP {
		panic("TODO") // x.Pickle(w)
		PickleJ(x.U, x.V, w)
	}
}

func (o *JTuple) Mod_StrRev(s string) (U, V) {
  n := len(o.PP)
  aa := make([]interface{}, n)
  for i:=0; i<n; i++ {
    w := o.PP[i]
    switch inline.Tag(w.U, w.V) {
    case Py:
      p := inline.TakePJ(w.U, w.V)
      //log.Println("Mod_StrRev: i=", i, "u=", w.U, "v=", w.V, "p=", p, "p.TypeName=", p.TypeName())
      if (p != nil) {
        aa[i] = p.Contents()
      }
    case Int:
      aa[i] = inline.TakeIntJ(w.U, w.V)
    case Str:
      aa[i] = inline.TakeStrJ(w.U, w.V)
    }
  }
  z := fmt.Sprintf(s, aa...)
  //log.Printf("Mod_StrRev  fmt %q", s)
  //for i, e := range aa {
    //log.Printf("Mod_StrRev  arg[%d] = %#v", i, e)
  //}
  //log.Printf("Mod_StrRev  >>>>> %q", z)
  return inline.MkStrJ(z)
}

//#if 0
func (o *JTuple) ModRev(u U, v V) (U, V) {
  n := len(o.PP)
  switch inline.Tag(u,v) {
  case Str:
    s := inline.TakeStrJ(u,v)
    aa := make([]interface{}, n)
    for i:=0; i<n; i++ {
      w := o.PP[i]
      p := inline.TakePJ(w.U, w.V)
      aa[i] = p.Contents()
    }
    z := fmt.Sprintf(s, aa...)
    return inline.MkStrJ(z)
  }
  panic("Cannot non-string MOD tuple")
}
//#endif

func (o *JTuple) Add_Py(a PJ) (U, V) {
	b := a.List()
	v := make([]W, 0, len(o.PP)+len(b))
	v = append(v, o.PP...)
	v = append(v, b...)
	z := MkTupleJ(v)
	return inline.MkPJ(z)
}
func (o *JTuple) Repr() string {
	buf := bytes.NewBufferString("(")
	n := len(o.PP)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		e := o.PP[i]
		buf.WriteString(ReprJ(e.U, e.V))
	}
  if n == 1 {
	  buf.WriteString(",")
  }
	buf.WriteString(")")
	return buf.String()
}

func (o *JTuple) PType() (U, V)    { return H_tuple_1, H_tuple_2 }
func (o *JTuple) TypeName() string { return "tuple" }

func (o *JList) Pickle(w *bytes.Buffer) {
	l := int64(len(o.PP))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypList + n))
	RypWriteInt(w, l)
	for _, x := range o.PP {
		panic("TODO") // x.Pickle(w)
		PickleJ(x.U, x.V, w)
	}
}
func (o *JList) Add_Py(a PJ) (U, V) {
	b := a.List()
	v := make([]W, 0, len(o.PP)+len(b))
	v = append(v, o.PP...)
	v = append(v, b...)
	z := MkListJ(v)
	return inline.MkPJ(z)
}
func (o *JList) Contents() interface{}     { return o.PP }
func (o *JList) Bool() bool                { return len(o.PP) != 0 }
func (o *JList) Contains(u U, v V) bool {
	for _, x := range o.PP {
		if inline.EQJ(x.U, x.V, u, v) {
			return true
		}
	}
	return false
}
func (o *JList) Bytes() []byte {
	zz := make([]byte, len(o.PP))
	for i, x := range o.PP {
		zz[i] = byte(IntJ(x.U, x.V))
	}
	return zz
}
func (o *JList) Len() int { return len(o.PP) }
func (o *JList) SetItemInt(i int, x1 U, x2 V) {
	if i < 0 {
		i += len(o.PP)
	}
	o.PP[i] = W{x1, x2}
}

func (o *JList) GetItemInt(i int) (U, V) {
  if i < 0 {
    i = i + len(o.PP)
  }
  z := o.PP[i]
  return z.U, z.V
}
func (o *JList) GetItemStr(s string) (U, V) {
  panic("Cannot GetItem on list with str")
}
func (o *JList) GetItemSlice(i, j int) (U, V) {
	n := len(o.PP)
	if j > n {
		j = n // Allow second index too big?
	}
	r := MkListJ(o.PP[i:j])
	return inline.MkPJ(r)
}

func (o *JList) String() string   { return o.Repr() }
func (o *JList) PType() (U, V)    { return H_list_1, H_list_2 }
func (o *JList) TypeName() string { return "list" }
func (o *JList) Repr() string {
	buf := bytes.NewBufferString("[")
	n := len(o.PP)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		e := o.PP[i]
		buf.WriteString(ReprJ(e.U, e.V))
	}
	buf.WriteString("]")
	return buf.String()
}
func (o *JList) Iter() JNexter {
	z := &JListIter{PP: o.PP}
	ForgeJ(z)
	return z
}
func (o *JList) List() []W {
	return o.PP
}

func (o *JList) DelItem(x1 U, x2 V) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	i := int(IntJ(x1, x2))
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
func (o *JList) DelItemSlice(i, j int) {
	// Check out: https://code.google.com/p/go-wiki/wiki/SliceTricks
	a := o.PP
	copy(a[i:], a[j:])
	o.PP = a[:len(a)-j+i]
}

// ========== dict

func (o *JDict) Hash() int64 {
  panic("TODO")
	//#if 0
	var z int64
	//#if mudict
	o.mu.Lock()
	//#endif
	for k, v := range o.ppp {
		z += int64(crc64.Checksum([]byte(k), CrcPolynomial))
		z += v.Hash() // TODO better
	}
	//#if mudict
	o.mu.Unlock()
	//#endif
	return z
	//#endif
}
func (o *JDict) Pickle(w *bytes.Buffer) {
  panic("TODO")
	//#if 0
	//#if mudict
	o.mu.Lock()
	defer o.mu.Unlock()
	//#endif
	l := int64(len(o.ppp))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypDict + n))
	RypWriteInt(w, l)
	for k, v := range o.ppp {
		MkStr(k).Pickle(w)
		v.Pickle(w)
	}
	//#endif
}
func (o *JDict) Contents() interface{} { return o.ppp }
func (o *JDict) Bool() bool            { return len(o.ppp) != 0 }
func (o *JDict) Contains(u U, v V) bool {
	key := StringJ(u, v)  // Coerse to string.
	//#if mudict
	o.mu.Lock()
	//#endif
	_, ok := o.ppp[key]
	//#if mudict
	o.mu.Unlock()
	//#endif
	return ok
}
func (o *JDict) Len() int { return len(o.ppp) }
func (o *JDict) SetItemStr(key string, ux U, vx V) {
	//#if mudict
	o.mu.Lock()
	//#endif
	o.ppp[key] = W{ux, vx}
	//#if mudict
	o.mu.Unlock()
	//#endif
}
func (o *JDict) GetItemInt(i int) (U, V) {
	s := F("%d", i)
  return o.GetItemStr(s)
}
func (o *JDict) GetItemStr(s string) (U, V) {
	//#if mudict
	o.mu.Lock()
	//#endif
	z, ok := o.ppp[s]
	//#if mudict
	o.mu.Unlock()
	//#endif
	if !ok {
		panic("JDict: KeyError")
	}
	return z.U, z.V
}
func (o *JDict) String() string { return o.Repr() }
func (o *JDict) PType() (U,V)       { return H_dict_1, H_dict_2 }
func (o *JDict) RType() string  { return "dict" }
func (o *JDict) takeJKVSnapshot() JKVSlice {
	//#if mudict
	o.mu.Lock()
	//#endif
	vec := make(JKVSlice, 0, len(o.ppp))
	for k, v := range o.ppp {
		vec = append(vec, JKV{k, v.U, v.V})
	}
	//#if mudict
	o.mu.Unlock()
	//#endif
  return vec
}
func (o *JDict) Repr() string {
  vec := o.takeJKVSnapshot()

	sort.Sort(vec)
	buf := bytes.NewBufferString("{")
	n := len(vec)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(F("%q: %s", vec[i].Key, ReprJ(vec[i].U, vec[i].V)))
	}
	buf.WriteString("}")
	return buf.String()
}
func (o *JDict) Enough() {}
func (o *JDict) takeKeysWSnapshot() []W {
  keys := make([]W, 0, len(o.ppp))
	//#if mudict
	o.mu.Lock()
	//#endif
	for k, _ := range o.ppp {
    ku, kv := inline.MkStrJ(k)
		keys = append(keys, W{ku, kv})
	}
	//#if mudict
	o.mu.Unlock()
	//#endif
  return keys
}
func (o *JDict) Iter() JNexter {
	z := &JListIter{PP: o.takeKeysWSnapshot()}
	ForgeJ(z)
	return z
}
func (o *JDict) List() []W {
  return o.takeKeysWSnapshot()
}
func (o *JDict) Dict() JScope {
	return o.ppp
}
func (o *JDict) DelItem(i M) {
	key := i.String()
	//#if mudict
	o.mu.Lock()
	//#endif
	delete(o.ppp, key)
	//#if mudict
	o.mu.Unlock()
	//#endif
}
func (o *JDict) Compare(u U, v V) int {
 panic("TODO -- JDict Compare")
 //#if 0
	switch b := a.X.(type) {
	case *JDict:
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
		//#if mudict
		o.mu.Lock()
		//#endif
		for i, x := range ostrs {
			olist[i*2] = MkStr(x)
			olist[i*2+1] = o.ppp[x]
		}
		//#if mudict
		o.mu.Unlock()
		b.mu.Lock()
		//#endif
		for i, x := range astrs {
			alist[i*2] = MkStr(x)
			alist[i*2+1] = b.ppp[x]
		}
		//#if mudict
		b.mu.Unlock()
		//#endif
		return MkList(olist).Compare(MkList(alist))
	}
	return StrCmp(o.PType().String(), a.PType().String())
  //#endif
}

func MkDictCopyJ(ppp JScope) *JDict {
	z := &JDict{ppp: make(JScope)}
	for k, v := range ppp {
		z.ppp[k] = v
	}
	ForgeJ(z)
  return z
}

func MkDictFromPairsJ(pp []W) *JDict {
	z := &JDict{ppp: make(JScope)}
	for _, x := range pp {
		sub := ListJ(x.U, x.V)
		if len(sub) != 2 {
			panic(F("MkDictFromPairs: got sublist of size %d, wanted size 2", len(sub)))
		}
		k := StringJ(sub[0].U, sub[0].V)
		v := sub[1]
		z.ppp[k] = v
	}
	ForgeJ(z)
  return z
}

func NewListJ() *JList {
	z := &JList{PP: make([]W, 0)}
	ForgeJ(z)
	return z
}

func CopyPJs(pp []W) []W {
	zz := make([]W, len(pp))
	copy(zz, pp)
	return zz
}
func CopyListJ(aa *JList) *JList {
	z := &JList{PP: CopyPJs(aa.PP)}
	ForgeJ(z)
	return z
}

type JKVSlice []JKV            // Can be sorted by Key.
func (vec JKVSlice) Len() int { return len(vec) }
func (vec JKVSlice) Less(i, j int) bool {
	return vec[i].Key < vec[j].Key
}
func (vec JKVSlice) Swap(i, j int) {
	vec[i], vec[j] = vec[j], vec[i]
}

type JKV struct {
	Key   string
	U U
  V V
}
