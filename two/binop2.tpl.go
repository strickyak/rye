// +build prego

// Expanding binop.tpl.go with OPNAME
package rye

func FnOPNAME(a1 U, a2 V, b1 U, b2 V) (U, V) {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return a.OPNAME_Py(b)
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return a.OPNAME_Int(b)
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return a.OPNAME_Str(b)
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.OPNAME_IntRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return OPNAME_work.IntInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return OPNAME_work.IntStr(a, b)
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return b.OPNAME_StrRev(a)
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return OPNAME_work.StrInt(a, b)
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return OPNAME_work.StrStr(a, b)
		}
	}
	panic(F("Bad switch FnOPNAME: %x %x %x %x", a1, a2, b1, b2))
}

var OPNAME_work OPNAME_Worker = OPNAME_Worker{BinopWorkerBase{Name: "OPNAME"}}

type OPNAME_Worker struct {
	BinopWorkerBase
}

func (*OPNAME_Worker) PyPy(a PJ, b PJ) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_OPNAME(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := float64(macro.Operator_OPNAME(ta.F, tb.F))
			return macro.MkFloatJ(z)
		}
	case *JBool:
		switch tb := b.(type) {
		case *JFloat:
			z := float64(macro.Operator_OPNAME(ta.F, tb.F))
			return macro.MkFloatJ(z)
		case *JBool:
			z := int64(macro.Operator_OPNAME(ta.I, tb.I))
			return macro.MkIntJ(z)
		}
	}
	panic(F("cannot OPNAME with %s & %s", a.TypeName(), b.TypeName()))
}

func (*OPNAME_Worker) PyInt(a PJ, b int64) (U, V) {
	switch ta := a.(type) {
	case *JFloat:
		z := float64(macro.Operator_OPNAME(ta.F, float64(b)))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_OPNAME(ta.I, b))
		return MkIntJ(z)
	}
	panic(F("cannot OPNAME with %s & int", a.TypeName()))
}

func (*OPNAME_Worker) IntPy(a int64, b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_OPNAME(float64(a), tb.F))
		return MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_OPNAME(a, tb.I))
		return MkIntJ(z)
	}
	panic(F("cannot OPNAME with %s & int", b.TypeName()))
}

func (*OPNAME_Worker) IntInt(a int64, b int64) (U, V) {
	z := int64(macro.Operator_OPNAME(a, b))
	return macro.MkIntJ(z)
}

func (o *JBool) OPNAME_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_OPNAME(o.F, tb.Float()))
		return inline.MkFloatJ(z)
	case *JBool:
		z := int64(macro.Operator_OPNAME(o.I, tb.I))
		return inline.MkIntJ(z)
	}
	panic(F("Cannot OPNAME with bool & %s", b.TypeName()))
}
func (o *JBool) OPNAME_Int(b int64) (U, V) {
	z := int64(macro.Operator_OPNAME(o.I, b))
	return inline.MkIntJ(z)
}
func (o *JBool) OPNAME_IntRev(a int64) (U, V) {
	z := int64(macro.Operator_OPNAME(a, o.I))
	return inline.MkIntJ(z)
}

func (o *JFloat) OPNAME_Py(b PJ) (U, V) {
	switch tb := b.(type) {
	case *JFloat:
		z := float64(macro.Operator_OPNAME(o.F, tb.F))
    //log.Printf("JFloat::OPNAME_Py %g %g -> %g", o.F, tb.F, z)
		return inline.MkFloatJ(z)
	case *JBool:
		z := float64(macro.Operator_OPNAME(o.F, tb.F))
		return inline.MkFloatJ(z)
	}
	panic(F("Cannot OPNAME with float & %s", b.TypeName()))
}
func (o *JFloat) OPNAME_Int(b int64) (U, V) {
	z := float64(macro.Operator_OPNAME(o.F, float64(b)))
    //log.Printf("JFloat::OPNAME_Int %g %d -> %g", o.F, b, z)
	return inline.MkFloatJ(z)
}
func (o *JFloat) OPNAME_IntRev(a int64) (U, V) {
	z := float64(macro.Operator_OPNAME(float64(a), o.F))
    //log.Printf("JFloat::OPNAME_IntRev %d %g -> %g", a, o.F, z)
	return inline.MkFloatJ(z)
}

//////////////////////////////////////////
