package rye

import (
	"unsafe"
)

// So that 'go get' will fetch prego as well:
import _ "github.com/strickyak/prego"

var _ = unsafe.Sizeof(0)

var PointerToInt *int

const SZ = unsafe.Sizeof(PointerToInt)

type D uintptr

func DFromint(i int) (D, D) {
	return D(1), D(i)
}

func DFromInt(i int64) (D, D) {
	return D(1), D(i)
}

func IntFromD(dx, dy D) int64 {
	return int64(dy)
}

func AddD(ax, ay, bx, by D) (D, D) {
	return DFromInt(IntFromD(ax, ay) + IntFromD(bx, by))
}

func AddDQuick(ax, ay, bx, by D) (D, D) {
	return D(1), ay + by
}

func DFromIntWrapper(x int64) (D, D) {
	for i := 0; i < 10; i++ {
		if i == 5 {
			return DFromInt(x)
		}
	}
	var dx, dy D
	return dx, dy
}

func PeekStr(s string) {
	p := uintptr(unsafe.Pointer(&s))
	dp1 := (*uintptr)(unsafe.Pointer(p))
	dp2 := (*uintptr)(unsafe.Pointer(p + SZ))

	println("PeekStr:", s, ":::", *dp1, *dp2)
}

/*
func DFromStr(s string) (D,D) {
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
*/
