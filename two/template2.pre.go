// +build prego

package two
import . "github.com/strickyak/rye"
import "reflect"
import "log"
import "math"
import "unsafe"

var _ = unsafe.Sizeof(0)
var _ = log.Printf

// +build prego

// Expanding binop.tpl.go with Add


func FnAdd(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Add_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Add_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Add_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Add_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Add_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Add_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Add_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Add_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Add_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnAdd: %x %x %x %x", a1, a2, b1, b2))
}

var Add_work Add_Worker = Add_Worker{BinopWorkerBase{Name: "Add"}}

type Add_Worker struct {
	BinopWorkerBase
}

func (*Add_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Add(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Add(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Add(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Add(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Add with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Add_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Add(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Add(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Add with %s & int", a.TypeName()))
}

func (*Add_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Add(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Add(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Add with %s & int", b.TypeName()))
}

func (*Add_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Add(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Add_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Add(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Add(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Add with bool & %s", b.TypeName()))
}
func (o *JBool) Add_Int(b int64) (U, V) {
	z := int64(macro.Operator_Add(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Add_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Add(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Add_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Add(o.F, tb.F))
    //log.Printf("JFloat::Add_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Add(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Add with float & %s", b.TypeName()))
}
func (o *JFloat) Add_Int(b int64) (U, V) {
	z := float64(macro.Operator_Add(o.F, float64(b)))
    //log.Printf("JFloat::Add_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Add_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Add(float64(a), o.F))
    //log.Printf("JFloat::Add_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding binop.tpl.go with Sub


func FnSub(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Sub_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Sub_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Sub_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Sub_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Sub_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Sub_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Sub_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Sub_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Sub_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnSub: %x %x %x %x", a1, a2, b1, b2))
}

var Sub_work Sub_Worker = Sub_Worker{BinopWorkerBase{Name: "Sub"}}

type Sub_Worker struct {
	BinopWorkerBase
}

func (*Sub_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Sub(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Sub(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Sub(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Sub(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Sub with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Sub_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Sub(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Sub(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Sub with %s & int", a.TypeName()))
}

func (*Sub_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Sub(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Sub(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Sub with %s & int", b.TypeName()))
}

func (*Sub_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Sub(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Sub_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Sub(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Sub(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Sub with bool & %s", b.TypeName()))
}
func (o *JBool) Sub_Int(b int64) (U, V) {
	z := int64(macro.Operator_Sub(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Sub_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Sub(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Sub_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Sub(o.F, tb.F))
    //log.Printf("JFloat::Sub_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Sub(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Sub with float & %s", b.TypeName()))
}
func (o *JFloat) Sub_Int(b int64) (U, V) {
	z := float64(macro.Operator_Sub(o.F, float64(b)))
    //log.Printf("JFloat::Sub_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Sub_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Sub(float64(a), o.F))
    //log.Printf("JFloat::Sub_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding binop.tpl.go with Mul


func FnMul(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Mul_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Mul_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Mul_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Mul_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Mul_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Mul_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Mul_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Mul_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Mul_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnMul: %x %x %x %x", a1, a2, b1, b2))
}

var Mul_work Mul_Worker = Mul_Worker{BinopWorkerBase{Name: "Mul"}}

type Mul_Worker struct {
	BinopWorkerBase
}

func (*Mul_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Mul(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Mul(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Mul(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Mul(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Mul with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Mul_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mul(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mul(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Mul with %s & int", a.TypeName()))
}

func (*Mul_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mul(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mul(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Mul with %s & int", b.TypeName()))
}

func (*Mul_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Mul(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Mul_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mul(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mul(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Mul with bool & %s", b.TypeName()))
}
func (o *JBool) Mul_Int(b int64) (U, V) {
	z := int64(macro.Operator_Mul(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Mul_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Mul(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Mul_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mul(o.F, tb.F))
    //log.Printf("JFloat::Mul_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Mul(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Mul with float & %s", b.TypeName()))
}
func (o *JFloat) Mul_Int(b int64) (U, V) {
	z := float64(macro.Operator_Mul(o.F, float64(b)))
    //log.Printf("JFloat::Mul_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Mul_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Mul(float64(a), o.F))
    //log.Printf("JFloat::Mul_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding binop.tpl.go with Div


func FnDiv(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Div_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Div_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Div_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Div_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Div_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Div_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Div_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Div_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Div_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnDiv: %x %x %x %x", a1, a2, b1, b2))
}

var Div_work Div_Worker = Div_Worker{BinopWorkerBase{Name: "Div"}}

type Div_Worker struct {
	BinopWorkerBase
}

func (*Div_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Div(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Div(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Div(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Div(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Div with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Div_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Div(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Div(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Div with %s & int", a.TypeName()))
}

func (*Div_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Div(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Div(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Div with %s & int", b.TypeName()))
}

func (*Div_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Div(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Div_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Div(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Div(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Div with bool & %s", b.TypeName()))
}
func (o *JBool) Div_Int(b int64) (U, V) {
	z := int64(macro.Operator_Div(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Div_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Div(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Div_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Div(o.F, tb.F))
    //log.Printf("JFloat::Div_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Div(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Div with float & %s", b.TypeName()))
}
func (o *JFloat) Div_Int(b int64) (U, V) {
	z := float64(macro.Operator_Div(o.F, float64(b)))
    //log.Printf("JFloat::Div_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Div_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Div(float64(a), o.F))
    //log.Printf("JFloat::Div_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding binop.tpl.go with Mod


func FnMod(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Mod_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Mod_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Mod_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Mod_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Mod_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Mod_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Mod_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Mod_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Mod_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnMod: %x %x %x %x", a1, a2, b1, b2))
}

var Mod_work Mod_Worker = Mod_Worker{BinopWorkerBase{Name: "Mod"}}

type Mod_Worker struct {
	BinopWorkerBase
}

func (*Mod_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Mod(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Mod(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Mod(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Mod(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Mod with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Mod_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mod(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mod(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Mod with %s & int", a.TypeName()))
}

func (*Mod_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mod(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mod(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Mod with %s & int", b.TypeName()))
}

func (*Mod_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Mod(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Mod_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mod(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Mod(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Mod with bool & %s", b.TypeName()))
}
func (o *JBool) Mod_Int(b int64) (U, V) {
	z := int64(macro.Operator_Mod(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Mod_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Mod(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Mod_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Mod(o.F, tb.F))
    //log.Printf("JFloat::Mod_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Mod(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Mod with float & %s", b.TypeName()))
}
func (o *JFloat) Mod_Int(b int64) (U, V) {
	z := float64(macro.Operator_Mod(o.F, float64(b)))
    //log.Printf("JFloat::Mod_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Mod_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Mod(float64(a), o.F))
    //log.Printf("JFloat::Mod_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding binop.tpl.go with Pow


func FnPow(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.Pow_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.Pow_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.Pow_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Pow_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Pow_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Pow_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.Pow_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return Pow_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return Pow_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnPow: %x %x %x %x", a1, a2, b1, b2))
}

var Pow_work Pow_Worker = Pow_Worker{BinopWorkerBase{Name: "Pow"}}

type Pow_Worker struct {
	BinopWorkerBase
}

func (*Pow_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Pow(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_Pow(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_Pow(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_Pow(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot Pow with %s & %s", a.TypeName(), b.TypeName()))
}

func (*Pow_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_Pow(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Pow(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot Pow with %s & int", a.TypeName()))
}

func (*Pow_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Pow(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Pow(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot Pow with %s & int", b.TypeName()))
}

func (*Pow_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_Pow(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) Pow_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Pow(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_Pow(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot Pow with bool & %s", b.TypeName()))
}
func (o *JBool) Pow_Int(b int64) (U, V) {
	z := int64(macro.Operator_Pow(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) Pow_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_Pow(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) Pow_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_Pow(o.F, tb.F))
    //log.Printf("JFloat::Pow_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_Pow(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot Pow with float & %s", b.TypeName()))
}
func (o *JFloat) Pow_Int(b int64) (U, V) {
	z := float64(macro.Operator_Pow(o.F, float64(b)))
    //log.Printf("JFloat::Pow_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) Pow_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_Pow(float64(a), o.F))
    //log.Printf("JFloat::Pow_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with LT


func FnLT(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LT(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_LT(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_LT(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LT(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_LT(a, b)
		case Str:
			return inline.Compares_LT(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LT(-b.Compare_Str(a))
		case Int:
			return inline.Compares_LT(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_LT(a, b)
		}
	}
	panic(F("Bad switch FnLT: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with LE


func FnLE(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LE(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_LE(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_LE(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LE(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_LE(a, b)
		case Str:
			return inline.Compares_LE(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_LE(-b.Compare_Str(a))
		case Int:
			return inline.Compares_LE(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_LE(a, b)
		}
	}
	panic(F("Bad switch FnLE: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with GT


func FnGT(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GT(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_GT(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_GT(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GT(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_GT(a, b)
		case Str:
			return inline.Compares_GT(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GT(-b.Compare_Str(a))
		case Int:
			return inline.Compares_GT(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_GT(a, b)
		}
	}
	panic(F("Bad switch FnGT: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with GE


func FnGE(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GE(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_GE(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_GE(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GE(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_GE(a, b)
		case Str:
			return inline.Compares_GE(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_GE(-b.Compare_Str(a))
		case Int:
			return inline.Compares_GE(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_GE(a, b)
		}
	}
	panic(F("Bad switch FnGE: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with EQ


func FnEQ(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_EQ(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_EQ(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_EQ(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_EQ(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_EQ(a, b)
		case Str:
			return inline.Compares_EQ(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_EQ(-b.Compare_Str(a))
		case Int:
			return inline.Compares_EQ(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_EQ(a, b)
		}
	}
	panic(F("Bad switch FnEQ: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
// +build prego

// Expanding relop.tpl.go with NE


func FnNE(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_NE(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_NE(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_NE(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_NE(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_NE(a, b)
		case Str:
			return inline.Compares_NE(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_NE(-b.Compare_Str(a))
		case Int:
			return inline.Compares_NE(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_NE(a, b)
		}
	}
	panic(F("Bad switch FnNE: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
