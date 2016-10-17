// +build prego

package two

//#if noinline
import . "github.com/strickyak/rye"
import "reflect"
import "unsafe"
import "log"
var _ = log.Printf
//#endif

func (inline) Tag(u U, v V) int {
  return int((u)&7)
}
func (inline) IsPy(u U, v V) bool {
  return (((u)&7) == Py)
}
func (inline) IsInt(u U, v V) bool {
  return (((u)&7) == Int)
}
func (inline) IsStr(u U, v V) bool {
  return (((u)&7) == Str)
}
func (inline) IsEggs(u U, v V) bool {
  return ((uintptr(u)==0) && (uintptr(unsafe.Pointer(v))==0))
}

func (macro) Operator_Add(a, b) {
  return ((a) + (b))
}
func (macro) Operator_Sub(a, b) {
  return ((a) - (b))
}
func (macro) Operator_Mul(a, b) {
  return ((a) * (b))
}
func (macro) Operator_Div(a, b) {
  return ((a) / (b))
}
func (macro) Operator_Mod(a, b) {
  __kind := reflect.ValueOf(a).Kind()
  if __kind == reflect.Float64 {
    panic(F("Cannot use % on floats"))
  }
  return (int64(a) % int64(b))
}
func (macro) Operator_Pow(a, b) {
  return math.Pow(float64(a), float64(b))
}
func (macro) Operator_LT(a, b) {
  return ((a) < (b))
}
func (macro) Operator_LE(a, b) {
  return ((a) <= (b))
}
func (macro) Operator_GT(a, b) {
  return ((a) > (b))
}
func (macro) Operator_GE(a, b) {
  return ((a) >= (b))
}
func (macro) Operator_EQ(a, b) {
  return ((a) == (b))
}
func (macro) Operator_NE(a, b) {
  return ((a) != (b))
}

func (inline) Compares_LT(cmp int) bool {
  return ((cmp) < 0)
}
func (inline) Compares_LE(cmp int) bool {
  return ((cmp) <= 0)
}
func (inline) Compares_GT(cmp int) bool {
  return ((cmp) > 0)
}
func (inline) Compares_GE(cmp int) bool {
  return ((cmp) >= 0)
}
func (inline) Compares_EQ(cmp int) bool {
  return ((cmp) == 0)
}
func (inline) Compares_NE(cmp int) bool {
  return ((cmp) != 0)
}

func (inline) MkBoolJ(b bool) (U,V) {
  var z1 U
  var z2 V
  if b {
    z1, z2 = TrueJ_1, TrueJ_2
  } else {
    z1, z2 = FalseJ_1, FalseJ_2
  }
  return z1, z2
}

  //Assert ("inline TakeBoolJ", j1 == TrueJ_1 && j2 == TrueJ_2 || j1 == FalseJ_1 && j2 == FalseJ_2)
func (inline) TakeBoolJ(j1 U, j2 V) bool {
  return (j2 == TrueJ_2)
}

  //Assert ("inline TakeIntJ", (j1&7) == Int)
func (inline) BigTakeIntJ(j1 U, j2 V) int64 {
  return (int64((j1) &^ 7) + int64(7 & uintptr(unsafe.Pointer(j2))))
}

func (inline) BigMkIntJ(x int64) (U, V) {
  var __u U = U(((x) &^ 7) | Int)
  var __v V = V(&Filler.Addr[(x) & 7])
  return __u, __v
}

func (inline) BigMkintJ(x int) (U, V) {
  var __u U = U(((x) &^ 7) | Int)
  var __v V = V(&Filler.Addr[(x) & 7])
  return __u, __v
}

  //Assert ("inline TakeIntJ", (j1&7) == Int)
func (inline) TakeIntJ(j1 U, j2 V) int64 {
  return (int64((j1) &^ 7) >> 3)
}

func (inline) MkIntJ(x int64) (U, V) {
  return U(((x<<3) ) | Int), V(nil)
}

func (inline) MkintJ(x int) (U, V) {
  return U(((x<<3) ) | Int), V(nil)
}

  //Assert ("inline TakeStrJ", (j1&7) == Str)
func (inline) TakeStrJ(j1 U, j2 V) string {
  var __s string
  __sh := (*reflect.StringHeader)(unsafe.Pointer(&__s))
  __sh.Data = uintptr(unsafe.Pointer(j2))
  __sh.Len = int(j1 >> 3)
  return __s
}

func (inline) MkStrJ(s string) (U, V) {
  var __tmp string = s
  __sh := (*reflect.StringHeader)(unsafe.Pointer(&__tmp))
  var __u U
  var __v V
  __u, __v = U(Str|(__sh.Len) << 3), V(unsafe.Pointer(__sh.Data))
  if false {
    Peek( __u, __v, "MkStrJ")
  }
  return __u, __v
}

func (inline) MkStrJX(s string) (U, V) {
  var __tmp string = s
  return inline.MkStrJ(__tmp)
}

  //Assert ("inline TakePJ", j1 == Py + 16)
func (inline) TakePJ(j1 U, j2 V) PJ {
  var __b *JBase = (*JBase)(unsafe.Pointer(j2))
  _ = __b
  return __b.Self
}
func (inline) CheckTakePJ(j1 U, j2 V) PJ {
  if ((j1) != Py + 16) {
    panic("int or str now allowed here")
  }
  var __b *JBase = (*JBase)(unsafe.Pointer(j2))
  _ = __b
  return __b.Self
}

// MkPJ is never not-inlined because it is used to force Pair().
func (macro) MkPJ(a PJ) (U, V) {
  var __b *JBase = (a).GetJBase()
  var __u U = Py + 16
  var __v V = V(unsafe.Pointer(__b))
  _,_ = __u, __v
  return __u, __v
}

func (inline) MkFloatJ(f float64) (U, V) {
  __z := &JFloat{F: f};
  __z.JBase.Self = __z
  return macro.MkPJ(__z)
}

func (inline) HashJ(u U, v V) int64 {
  var __z int64
  if ((u)&7) == Int {
    __z = inline.TakeIntJ(u, v)
  } else if ((u)&7) == Str {
    __s := inline.TakeStrJ(u, v)
    __z = StrHash(__s)
  } else {
    __pj := inline.TakePJ(u, v)
    __z = __pj.Hash()
  }
  return __z
}

func (macro) OperatorCompare(x, y) {
  var __z int
  if (x) < (y) {
    __z = -1
  } else if (x) > (y) {
    __z = 1
  }
  return __z
}

func (inline) CompareJ(a1 U, a2 V, b1 U, b2 V) int {
  var __z int
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __z = inline.OperatorCompare(__ta, __tb)
  } else if ((a1)&7) == Str && ((b1)&7) == Str {
    __ta := inline.TakeStrJ(a1, a2)
    __tb := inline.TakeStrJ(b1, b2)
    __z = inline.OperatorCompare(__ta, __tb)
  } else {
    __z = FnCompare(a1, a2, b1, b2)
  }
  return __z
}

func (inline) AddJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __tz := (int64(__ta) + int64(__tb))
    __z1, __z2 = inline.MkIntJ(__tz)
  } else if ((a1)&7) == Str && ((b1)&7) == Str {
    __ta := inline.TakeStrJ(a1, a2)
    __tb := inline.TakeStrJ(b1, b2)
    __tz := __ta + __tb
    ReifyString = &__tz
    __z1, __z2 = inline.MkStrJ(__tz)
  } else {
    __z1, __z2 = FnAdd(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func (inline) SubJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __tz := (int64(__ta) - int64(__tb))
    __z1, __z2 = inline.MkIntJ(__tz)
  } else {
    __z1, __z2 = FnSub(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func (inline) MulJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __tz := (int64(__ta) * int64(__tb))
    __z1, __z2 = inline.MkIntJ(__tz)
  } else {
    __z1, __z2 = FnMul(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func (inline) DivJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __tz := (int64(__ta) / int64(__tb))
    __z1, __z2 = inline.MkIntJ(__tz)
  } else {
    __z1, __z2 = FnDiv(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func (inline) ModJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := inline.TakeIntJ(a1, a2)
    __tb := inline.TakeIntJ(b1, b2)
    __tz := (int64(__ta) % int64(__tb))
    __z1, __z2 = inline.MkIntJ(__tz)
  } else {
    __z1, __z2 = FnMod(a1, a2, b1, b2)
  }
  return __z1, __z2
}

func (inline) PowJ(a1 U, a2 V, b1 U, b2 V) (U,V) {
 return FnPow(a1, a2, b1, b2)
}

func (macro) JRelOp(a1, a2, b1, b2, relop, fn) {
  var __rel bool
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ia := inline.TakeIntJ(a1, a2)
    __ib := inline.TakeIntJ(b1, b2)
    __rel = inline.relop(__ia, __ib)
  } else if ((a1)&7) == Str && ((b1)&7) == Str {
    __sa := inline.TakeStrJ(a1, a2)
    __sb := inline.TakeStrJ(b1, b2)
    __rel = inline.relop(__sa, __sb)
  } else {
    __rel = fn(a1, a2, b1, b2)
  }
  return __rel
}

func (macro) LTJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_LT, FnLT)
}
func (macro) LEJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_LE, FnLE)
}
func (macro) GTJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_GT, FnGT)
}
func (macro) GEJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_GE, FnGE)
}
func (macro) EQJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_EQ, FnEQ)
}
func (macro) NEJ(a1, a2, b1, b2) {
  return macro.JRelOp(a1, a2, b1, b2, Operator_NE, FnNE)
}

func (inline) LenJ(j1 U, j2 V) int {
  var __len int
  if ((j1) & 7) == Str {
    __len = int(j1) >> 3
  } else {
    __len = FnLenJ((j1), (j2))
  }
  return __len
}

func (inline) NullJ(u U, v V) bool {
  return ((u) == NoneJ_1 && (v) == NoneJ_2)
}

func (inline) IntOkOrNone(a1 U, a2 V) (int64, bool) {
  var __intor int64
  var __ok bool

  if ((a1)&7) == Int {
    __intor, __ok = macro.TakeIntJ(a1,a2), true
  } else if a1 == NoneJ_1 && a2 == NoneJ_2 {
    __intor, __ok = 0, false
  } else {
    __p := macro.TakePJ(a1, a2)
    __intor, __ok = __p.Int(), true
  }
  return __intor, __ok
}
func (inline) ForgeUV(p PJ) (U, V) {
  p.SetSelf(p)
	return macro.MkPJ(p)
}
