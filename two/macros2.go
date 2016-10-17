//

package two


import . "github.com/strickyak/rye"
import "reflect"
import "unsafe"
import "log"
var _ = log.Printf


func Tag ( u U, v V )  int  {
  return int((u)&7)
}
func IsPy ( u U, v V )  bool  {
  return (((u)&7) == Py)
}
func IsInt ( u U, v V )  bool  {
  return (((u)&7) == Int)
}
func IsStr ( u U, v V )  bool  {
  return (((u)&7) == Str)
}
func IsEggs ( u U, v V )  bool  {
  return ((uintptr(u)==0) && (uintptr(unsafe.Pointer(v))==0))
}










































func Compares_LT ( cmp int )  bool  {
  return ((cmp) < 0)
}
func Compares_LE ( cmp int )  bool  {
  return ((cmp) <= 0)
}
func Compares_GT ( cmp int )  bool  {
  return ((cmp) > 0)
}
func Compares_GE ( cmp int )  bool  {
  return ((cmp) >= 0)
}
func Compares_EQ ( cmp int )  bool  {
  return ((cmp) == 0)
}
func Compares_NE ( cmp int )  bool  {
  return ((cmp) != 0)
}

func MkBoolJ ( b bool )  (U,V)  {
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
func TakeBoolJ ( j1 U, j2 V )  bool  {
  return (j2 == TrueJ_2)
}

  //Assert ("inline TakeIntJ", (j1&7) == Int)
func BigTakeIntJ ( j1 U, j2 V )  int64  {
  return (int64((j1) &^ 7) + int64(7 & uintptr(unsafe.Pointer(j2))))
}

func BigMkIntJ ( x int64 )  (U, V)  {
  var __u U = U(((x) &^ 7) | Int)
  var __v V = V(&Filler.Addr[(x) & 7])
  return __u, __v
}

func BigMkintJ ( x int )  (U, V)  {
  var __u U = U(((x) &^ 7) | Int)
  var __v V = V(&Filler.Addr[(x) & 7])
  return __u, __v
}

  //Assert ("inline TakeIntJ", (j1&7) == Int)
func TakeIntJ ( j1 U, j2 V )  int64  {
  return (int64((j1) &^ 7) >> 3)
}

func MkIntJ ( x int64 )  (U, V)  {
  return U(((x<<3) ) | Int), V(nil)
}

func MkintJ ( x int )  (U, V)  {
  return U(((x<<3) ) | Int), V(nil)
}

  //Assert ("inline TakeStrJ", (j1&7) == Str)
func TakeStrJ ( j1 U, j2 V )  string  {
  var __s string
  __sh := (*reflect.StringHeader)(unsafe.Pointer(&__s))
  __sh.Data = uintptr(unsafe.Pointer(j2))
  __sh.Len = int(j1 >> 3)
  return __s
}

func MkStrJ ( s string )  (U, V)  {
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

func MkStrJX ( s string )  (U, V)  {
  var __tmp string = s
  return /*noinline:*/MkStrJ(__tmp)
}

  //Assert ("inline TakePJ", j1 == Py + 16)
func TakePJ ( j1 U, j2 V )  PJ  {
  var __b *JBase = (*JBase)(unsafe.Pointer(j2))
  _ = __b
  return __b.Self
}
func CheckTakePJ ( j1 U, j2 V )  PJ  {
  if ((j1) != Py + 16) {
    panic("int or str now allowed here")
  }
  var __b *JBase = (*JBase)(unsafe.Pointer(j2))
  _ = __b
  return __b.Self
}

// MkPJ is never not-inlined because it is used to force Pair().








func MkFloatJ ( f float64 )  (U, V)  {
  __z := &JFloat{F: f};
  __z.JBase.Self = __z
 /*macro:MkPJ{*/   var _91__b *JBase = (__z).GetJBase();   var _91__u U = Py + 16;   var _91__v V = V(unsafe.Pointer(_91__b));   _,_ = _91__u, _91__v;  /*macro}*/   return _91__u, _91__v
}

func HashJ ( u U, v V )  int64  {
  var __z int64
  if ((u)&7) == Int {
    __z = /*noinline:*/TakeIntJ(u,  v)
  } else if ((u)&7) == Str {
    __s := /*noinline:*/TakeStrJ(u,  v)
    __z = StrHash(__s)
  } else {
    __pj := /*noinline:*/TakePJ(u,  v)
    __z = __pj.Hash()
  }
  return __z
}











func CompareJ ( a1 U, a2 V, b1 U, b2 V )  int  {
  var __z int
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
 /*macro:OperatorCompare{*/   var _120__z int;   if (__ta) < ( __tb) {     _120__z = -1;   } else if (__ta) > ( __tb) {     _120__z = 1;   };  /*macro}*/     __z = _120__z
  } else if ((a1)&7) == Str && ((b1)&7) == Str {
    __ta := /*noinline:*/TakeStrJ(a1,  a2)
    __tb := /*noinline:*/TakeStrJ(b1,  b2)
 /*macro:OperatorCompare{*/   var _133__z int;   if (__ta) < ( __tb) {     _133__z = -1;   } else if (__ta) > ( __tb) {     _133__z = 1;   };  /*macro}*/     __z = _133__z
  } else {
    __z = FnCompare(a1, a2, b1, b2)
  }
  return __z
}

func AddJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
    __tz := (int64(__ta) + int64(__tb))
    __z1, __z2 = /*noinline:*/MkIntJ(__tz)
  } else if ((a1)&7) == Str && ((b1)&7) == Str {
    __ta := /*noinline:*/TakeStrJ(a1,  a2)
    __tb := /*noinline:*/TakeStrJ(b1,  b2)
    __tz := __ta + __tb
    ReifyString = &__tz
    __z1, __z2 = /*noinline:*/MkStrJ(__tz)
  } else {
    __z1, __z2 = FnAdd(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func SubJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
    __tz := (int64(__ta) - int64(__tb))
    __z1, __z2 = /*noinline:*/MkIntJ(__tz)
  } else {
    __z1, __z2 = FnSub(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func MulJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
    __tz := (int64(__ta) * int64(__tb))
    __z1, __z2 = /*noinline:*/MkIntJ(__tz)
  } else {
    __z1, __z2 = FnMul(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func DivJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
    __tz := (int64(__ta) / int64(__tb))
    __z1, __z2 = /*noinline:*/MkIntJ(__tz)
  } else {
    __z1, __z2 = FnDiv(a1, a2, b1, b2)
  }
  return __z1, __z2
}
func ModJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
  var __z1 U
  var __z2 V
  if ((a1)&7) == Int && ((b1)&7) == Int {
    __ta := /*noinline:*/TakeIntJ(a1,  a2)
    __tb := /*noinline:*/TakeIntJ(b1,  b2)
    __tz := (int64(__ta) % int64(__tb))
    __z1, __z2 = /*noinline:*/MkIntJ(__tz)
  } else {
    __z1, __z2 = FnMod(a1, a2, b1, b2)
  }
  return __z1, __z2
}

func PowJ ( a1 U, a2 V, b1 U, b2 V )  (U,V)  {
 return FnPow(a1, a2, b1, b2)
}




































func LenJ ( j1 U, j2 V )  int  {
  var __len int
  if ((j1) & 7) == Str {
    __len = int(j1) >> 3
  } else {
    __len = FnLenJ((j1), (j2))
  }
  return __len
}

func NullJ ( u U, v V )  bool  {
  return ((u) == NoneJ_1 && (v) == NoneJ_2)
}

func IntOkOrNone ( a1 U, a2 V )  (int64, bool)  {
  var __intor int64
  var __ok bool

  if ((a1)&7) == Int {
    __intor, __ok = /*noinline:*/TakeIntJ(a1, a2), true
  } else if a1 == NoneJ_1 && a2 == NoneJ_2 {
    __intor, __ok = 0, false
  } else {
    __p := /*noinline:*/TakePJ(a1,  a2)
    __intor, __ok = __p.Int(), true
  }
  return __intor, __ok
}
func ForgeUV ( p PJ )  (U, V)  {
  p.SetSelf(p)
 /*macro:MkPJ{*/   var _255__b *JBase = (p).GetJBase();   var _255__u U = Py + 16;   var _255__v V = V(unsafe.Pointer(_255__b));   _,_ = _255__u, _255__v;  /*macro}*/ 	return _255__u, _255__v
}
