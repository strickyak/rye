package rye_test

// Benchmark interface dispatch vs type switch.
// Results:
//  1. They're almost exactly the same time.
//  2. Multiplying is faster than adding!?

import (
	"testing"
)

type Operator interface {
  Operate(a, b int) int
}

type Sum struct {}
func (*Sum) Operate(a, b int) int { return a + b }
func SumOperate(s *Sum, a, b int) int { return a + b }

type Product struct {}
func (*Product) Operate(a, b int) int { return a * b }
func ProductOperate(p *Product, a, b int) int { return a * b }

func Calculate(n int, op Operator) int {
  if n < 2 {
    return n
  } else {
    return op.Operate(n, Calculate(n-1, op))
  }
}

func CalculateSwitch(n int, op Operator) int {
  if n < 2 {
    return n
  } else {
    switch t := op.(type) {
    case *Sum:
      return SumOperate(t, n, Calculate(n-1, op))
    case *Product:
      return ProductOperate(t, n, Calculate(n-1, op))
    }
    panic(666)
  }
}

var Reify int

func BenchmarkCalculateSum(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		z += Calculate(100, &Sum{})
	}
  Reify = z
}

func BenchmarkCalculateProduct(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		z += Calculate(100, &Product{})
	}
  Reify = z
}

func BenchmarkCalculateSumSwitch(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		z += CalculateSwitch(100, &Sum{})
	}
  Reify = z
}

func BenchmarkCalculateProductSwitch(b *testing.B) {
	var z int
	for i := 0; i < b.N; i++ {
		z += CalculateSwitch(100, &Product{})
	}
  Reify = z
}

func Test(*testing.T) {}
