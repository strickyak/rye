package rye

import (
	"unsafe"
)

var _ = unsafe.Sizeof(0)

type W interface{}

func WFromint(i int) W {
	var w W
	wp := uintptr(unsafe.Pointer(&w))
	wp1 := (*uintptr)(unsafe.Pointer(wp))
	wp2 := (*uintptr)(unsafe.Pointer(wp + SZ))

	*wp1 = 1
	*wp2 = uintptr(i)
	return w
}

func WFromInt(i int64) W {
	var w W
	wp := uintptr(unsafe.Pointer(&w))
	wp1 := (*uintptr)(unsafe.Pointer(wp))
	wp2 := (*uintptr)(unsafe.Pointer(wp + SZ))

	*wp1 = 1
	*wp2 = uintptr(i)
	return w
}

func IntFromW(w W) int64 {
	p := uintptr(unsafe.Pointer(&w))
	wp2 := (*uintptr)(unsafe.Pointer(p + SZ))

	return int64(*wp2)
}

func AddW(a, b W) W {
	return WFromInt(IntFromW(a) + IntFromW(b))
}

func AddWQuick(a, b W) W {
	pa := uintptr(unsafe.Pointer(&a))
	pb := uintptr(unsafe.Pointer(&b))

	wpa2 := (*uintptr)(unsafe.Pointer(pa + SZ))
	wpb2 := (*uintptr)(unsafe.Pointer(pb + SZ))

	return WFromInt(int64(*wpa2) + int64(*wpb2))
}

func WFromIntWrapper(x int64) W {
	for i := 0; i < 10; i++ {
		if i == 5 {
			return WFromInt(x)
		}
	}
	var w W
	return w
}

func WFromStr(s string) W {
	sp := uintptr(unsafe.Pointer(&s))
	sp1 := (*uintptr)(unsafe.Pointer(sp))
	sp2 := (*uintptr)(unsafe.Pointer(sp + SZ))

	var w W
	wp := uintptr(unsafe.Pointer(&w))
	wp1 := (*uintptr)(unsafe.Pointer(wp))
	wp2 := (*uintptr)(unsafe.Pointer(wp + SZ))

	*wp1 = (*sp2 << 3) | 2 // String length.
	*wp2 = *sp1            // String pointer.
	return w
}

func WType(w W) uint {
	wp1 := (*uintptr)(unsafe.Pointer(&w))
	return uint(*wp1 & 7)
}

func WIsInt(w W) bool {
	wp1 := (*uintptr)(unsafe.Pointer(&w))
	return (*wp1 & 7) == 1
}

func WIsStr(w W) bool {
	wp1 := (*uintptr)(unsafe.Pointer(&w))
	return (*wp1 & 7) == 2
}

func WIsSlow(w W) bool {
	wp1 := (*uintptr)(unsafe.Pointer(&w))
	return (*wp1 & 7) == 0
}

func StrFromW(w W) string {
	wp := uintptr(unsafe.Pointer(&w))
	wp1 := (*uintptr)(unsafe.Pointer(wp))
	wp2 := (*uintptr)(unsafe.Pointer(wp + SZ))

	var s string
	sp := uintptr(unsafe.Pointer(&s))
	sp1 := (*uintptr)(unsafe.Pointer(sp))
	sp2 := (*uintptr)(unsafe.Pointer(sp + SZ))

	*sp1 = *wp2
	*sp2 = *wp1 >> 3
	return s
}

func init() {
	println("WFromInt 0:", WFromInt(0))
	println("WFromInt 100:", WFromInt(100))
	println("WFromInt -100:", WFromIntWrapper(-100))
	println("WFromInt 111:", IntFromW(AddW(WFromInt(100), WFromInt(11))))
	s := "hello world"
	s1 := s[:5]
	s2 := s[6:]
	PeekStr(s)
	PeekStr(s1)
	PeekStr(s2)
}
