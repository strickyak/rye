package rye

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func AddIntFast(a int, b int) int {
	return add(a, b)
}

func AddIntSlow(a int, b int) int {
	a1 := a
	a2 := a1
	a3 := a2

	b1 := b
	b2 := b1
	b3 := b2

	z := add(a3, b3)

	z1 := z
	z2 := z1
	z3 := z2

	return z3
}

/*
func AddDFast(a D, b D) D {
	return AddD(a, b)
}
*/

func AddDSlow(a D, b D) D {
	a1 := a
	a2 := a1
	a3 := a2

	b1 := b
	b2 := b1
	b3 := b2

	z := AddD(a3, b3)

	z1 := z
	z2 := z1
	z3 := z2

	return z3
}

func AddMFast(a M, b M) M {
	return a.Add(b)
}

func AddMSlow(a M, b M) M {
	a1 := a
	a2 := a1
	a3 := a2

	b1 := b
	b2 := b1
	b3 := b2

	z := a3.Add(b3)

	z1 := z
	z2 := z1
	z3 := z2

	return z3
}

func AddMSlowVar(a M, b M) M {
  var a1 M
	a1 = a
  var a2 M
	a2 = a1
  var a3 M
	a3 = a2

  var b1 M
	b1 = b
  var b2 M
	b2 = b1
  var b3 M
	b3 = b2

  var z M
	z = a3.Add(b3)

  var z1 M
	z1 = z
  var z2 M
	z2 = z1
  var z3 M
	z3 = z2

	return z3
}

func AddMSlowSink(a M, b M) M {
	a1 := a
	a2 := a1
	a3 := a2
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;

	b1 := b
	b2 := b1
	b3 := b2
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;

	z := a3.Add(b3)

	z1 := z
	z2 := z1
	z3 := z2
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;

	return z3
}

func AddDSlowSink(a D, b D) D {
	a1 := a
	a2 := a1
	a3 := a2
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;
  _ = a1 ; _ = a2 ; _ = a3 ;

	b1 := b
	b2 := b1
	b3 := b2
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;
  _ = b1 ; _ = b2 ; _ = b3 ;

	z := AddD(a3, b3)

	z1 := z
	z2 := z1
	z3 := z2
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;
  _ = z1 ; _ = z2 ; _ = z3 ;

	return z3
}

func BenchmarkAssignDSlow(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += IntFromD(AddDSlow(DFromint(i), DFromint(i*i)))
	}
	println(sum)
}

func BenchmarkAssignDSlowSink(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += IntFromD(AddDSlowSink(DFromint(i), DFromint(i*i)))
	}
	println(sum)
}

func BenchmarkAssignDFast(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += IntFromD(AddD(DFromint(i), DFromint(i*i)))
	}
	println(sum)
}

func BenchmarkAssignDFastQuick(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += IntFromD(AddDQuick(DFromint(i), DFromint(i*i)))
	}
	println(sum)
}

func BenchmarkAssignMSlow(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += AddMSlow(Mkint(i), Mkint(i*i)).Int()
	}
	println(sum)
}

func BenchmarkAssignMSlowVar(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += AddMSlowVar(Mkint(i), Mkint(i*i)).Int()
	}
	println(sum)
}

func BenchmarkAssignMSlowSink(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += AddMSlowSink(Mkint(i), Mkint(i*i)).Int()
	}
	println(sum)
}

func BenchmarkAssignMFast(b *testing.B) {
	var sum int64
	for i := 0; i < b.N; i++ {
		sum += AddMFast(Mkint(i), Mkint(i*i)).Int()
	}
	println(sum)
}

func BenchmarkAssignIntSlow(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		sum += AddIntSlow(i, i*i)
	}
	println(sum)
}

func BenchmarkAssignIntFast(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		sum += AddIntFast(i, i*i)
	}
	println(sum)
}
