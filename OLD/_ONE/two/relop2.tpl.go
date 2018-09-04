// +build prego

// Expanding relop.tpl.go with OPNAME
package rye

func FnOPNAME(a1 U, a2 V, b1 U, b2 V) bool {
	switch a1 & 7 {
	case Py:
		a := inline.TakePJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_OPNAME(a.Compare_Py(b))
		case Int:
			b := inline.TakeIntJ(b1, b2)
			return inline.Compares_OPNAME(a.Compare_Int(b))
		case Str:
			b := inline.TakeStrJ(b1, b2)
			return inline.Compares_OPNAME(a.Compare_Str(b))
		}
	case Int:
		a := inline.TakeIntJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_OPNAME(-b.Compare_Int(a))
		case Int:
			b := inline.TakeIntJ(a1, a2)
			return macro.Operator_OPNAME(a, b)
		case Str:
			return inline.Compares_OPNAME(-1)  // Int < Str
		}
	case Str:
		a := inline.TakeStrJ(a1, a2)
		switch b1 & 7 {
		case Py:
			b := inline.TakePJ(b1, b2)
			return inline.Compares_OPNAME(-b.Compare_Str(a))
		case Int:
			return inline.Compares_OPNAME(1)  // Str > Int
		case Str:
			b := inline.TakeStrJ(a1, a2)
			return macro.Operator_OPNAME(a, b)
		}
	}
	panic(F("Bad switch FnOPNAME: %x %x %x %x", a1, a2, b1, b2))
}

//////////////////////////////////////////
