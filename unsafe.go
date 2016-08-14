package rye

import (
	"unsafe"
)

var _ = unsafe.Sizeof(0)

var PointerToInt *int
const SZ = unsafe.Sizeof(PointerToInt)

type D interface{}

func DFromint(i int) D {
  var d D
  dp := uintptr(unsafe.Pointer(&d))
  dp1 := (*uintptr)(unsafe.Pointer(dp))
  dp2 := (*uintptr)(unsafe.Pointer(dp + SZ))

  *dp1 = 1
  *dp2 = uintptr(i)
  return d
}

func DFromInt(i int64) D {
  var d D
  dp := uintptr(unsafe.Pointer(&d))
  dp1 := (*uintptr)(unsafe.Pointer(dp))
  dp2 := (*uintptr)(unsafe.Pointer(dp + SZ))

  *dp1 = 1
  *dp2 = uintptr(i)
  return d
}

func IntFromD(d D) int64 {
  p := uintptr(unsafe.Pointer(&d))
  dp2 := (*uintptr)(unsafe.Pointer(p + SZ))

  return int64(*dp2)
}

func AddD(a, b D) D {
  return DFromInt(IntFromD(a) + IntFromD(b))
}

func AddDQuick(a, b D) D {
  pa := uintptr(unsafe.Pointer(&a))
  pb := uintptr(unsafe.Pointer(&b))

  dpa2 := (*uintptr)(unsafe.Pointer(pa + SZ))
  dpb2 := (*uintptr)(unsafe.Pointer(pb + SZ))

  return DFromInt(int64(*dpa2) + int64(*dpb2))
}

func DFromIntWrapper(x int64) D {
  for i := 0; i < 10; i++ {
    if i == 5 {
      return DFromInt(x)
    }
  }
  var d D
  return d
}

func PeekStr(s string) {
  p := uintptr(unsafe.Pointer(&s))
  dp1 := (*uintptr)(unsafe.Pointer(p))
  dp2 := (*uintptr)(unsafe.Pointer(p + SZ))

  println("PeekStr:", s, ":::", *dp1, *dp2)
}

func DFromStr(s string) D {
  sp := uintptr(unsafe.Pointer(&s))
  sp1 := (*uintptr)(unsafe.Pointer(sp))
  sp2 := (*uintptr)(unsafe.Pointer(sp + SZ))

  var d D
  dp := uintptr(unsafe.Pointer(&d))
  dp1 := (*uintptr)(unsafe.Pointer(dp))
  dp2 := (*uintptr)(unsafe.Pointer(dp + SZ))

  *dp1 = (*sp2 << 3) | 2 // String length.
  *dp2 = *sp1            // String pointer.
  return d
}

func DType(d D) uint {
  dp1 := (*uintptr)(unsafe.Pointer(&d))
  return uint(*dp1 & 7)
}

func DIsInt(d D) bool {
  dp1 := (*uintptr)(unsafe.Pointer(&d))
  return (*dp1 & 7) == 1
}

func DIsStr(d D) bool {
  dp1 := (*uintptr)(unsafe.Pointer(&d))
  return (*dp1 & 7) == 2
}

func DIsSlow(d D) bool {
  dp1 := (*uintptr)(unsafe.Pointer(&d))
  return (*dp1 & 7) == 0
}

func StrFromD(d D) string {
  dp := uintptr(unsafe.Pointer(&d))
  dp1 := (*uintptr)(unsafe.Pointer(dp))
  dp2 := (*uintptr)(unsafe.Pointer(dp + SZ))

  var s string
  sp := uintptr(unsafe.Pointer(&s))
  sp1 := (*uintptr)(unsafe.Pointer(sp))
  sp2 := (*uintptr)(unsafe.Pointer(sp + SZ))

  *sp1 = *dp2
  *sp2 = *dp1 >> 3
  return s
}

func init() {
  println("DFromInt 0:", DFromInt(0))
  println("DFromInt 100:", DFromInt(100))
  println("DFromInt -100:", DFromIntWrapper(-100))
  println("DFromInt 111:", IntFromD(AddD(DFromInt(100), DFromInt(11))))
  s := "hello world"
  s1 := s[:5]
  s2 := s[6:]
  PeekStr(s)
  PeekStr(s1)
  PeekStr(s2)
}
