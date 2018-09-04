// +build prego

package two

import . "github.com/strickyak/rye"
import "fmt"

//import R "reflect"
import "reflect"
import "sort"

//import "strings"
import "unsafe"

var _ = Shutdown        // USE rye
var _ = reflect.ValueOf // USE reflect

func N2_rye_what(u U, v V) (U, V) {
	s := fmt.Sprintf("{%s|%#v}", TypeNameJ(u, v), ContentsJ(u, v))
	return inline.MkStrJ(s)
	//#if 0
	if a.X != nil {
		switch x := a.X.(type) {
		case *PGo:
			z := fmt.Sprintf("PGo|%s|%s", x.V.Type().Kind(), x.V.Type())
			switch y := x.V.Interface().(type) {
			case R.Value:
				z += fmt.Sprintf("|%s|%s", y.Type().Kind(), y.Type())
			}
			return MkStr(z)
		}
	}
	return MkStr(strings.TrimLeft(fmt.Sprintf("%T", a), "*"))
	//#endif
}

func N2_set(u U, v V) (U, V) {
	panic("TODO")
	//#if 0
	z := make(JScope)
	// If a is None or empty, leave it empty.
	if a.Bool() {
		for _, e := range a.List() {
			z[e.String()] = True
		}
	}
	return MkSet(z)
	//#endif
}

func N2_dict(args_1 U, args_2 V, kw_1 U, kw_2 V) (U, V) {
	var d *JDict
	vec := ListJ(args_1, args_2)
	switch len(vec) {
	case 0:
		d = MkDictJ(make(JScope))
	case 1:
		u, v := vec[0].U, vec[0].V
		switch inline.Tag(u, v) {
		case Py:
			a := inline.TakePJ(u, v)
			switch t := a.(type) {
			case *JNone:
				d = MkDictJ(make(JScope))
			case *JList:
				d = MkDictFromPairsJ(t.List())
			case *JDict:
				d = MkDictCopyJ(JScope(t.Dict()))
			//case *JModule:
			//d = MkDictCopyJ(JScope(t.Dict()))
			case *JGo:
				d = MkDictJ(t.Dict())
			default:
				d = MkDictFromPairsJ(t.List())
			}
		case Int:
			panic("Cannot make dict from int")
		case Str:
			panic("Cannot make dict from str")
		}
	default:
		panic("Too many args to dict()")
	}
	kwd := inline.TakePJ(kw_1, kw_2).(*JDict)
	//#if mudict
	kwd.mu.Lock()
	//#endif
	for k, v := range kwd.ppp {
		d.ppp[k] = v
	}
	//#if mudict
	kwd.mu.Unlock()
	//#endif
	return macro.MkPJ(&d.JBase)
}

func N2_byt(u U, v V) (U, V) {
	switch inline.Tag(u, v) {
	case Py:
		a := inline.TakePJ(u, v)
		switch t := a.(type) {
		case *JByt:
			// byt() can be used to copy a byte array.
			bb := make([]byte, len(t.YY))
			copy(bb, t.YY)
			return MkBytUV(bb)
		}
	}

	// *PStr makes a copy already inside a.Bytes().
	// So does *PList.
	bb := BytesJ(u, v)
	return MkBytUV(bb)
}

func N2_mkbyt(u U, v V) (U, V) {
	return MkBytUV(make([]byte, int(IntJ(u, v))))
}

func N2_range(u U, v V) (U, V) {
	n := IntJ(u, v)
	vec := make([]W, n)
	for i := int64(0); i < n; i++ {
		iu, iv := inline.MkIntJ(i)
		vec[i] = W{iu, iv}
	}
	z := MkListJ(vec)
	return macro.MkPJ(z)
}

// For Sorting.

type jsorter struct {
	pp      []W
	reverse bool
	cmp     func(a1 U, a2 V, b1 U, b2 V) int
	key     func(u U, v V) (U, V)
}

func newJSorter(pp []W) *jsorter {
	return &jsorter{
		pp: pp,
		cmp: func(a1 U, a2 V, b1 U, b2 V) int {
			return inline.CompareJ(a1, a2, b1, b2)
		},
	}
}

func (o *jsorter) Len() int { return len(o.pp) }
func (o *jsorter) Less(i, j int) bool {
	aw := o.pp[i]
	au, av := aw.U, aw.V
	bw := o.pp[j]
	bu, bv := bw.U, bw.V
	if o.key != nil {
		au, av = o.key(au, av)
		bu, bv = o.key(bu, bv)
	}
	c := o.cmp(au, av, bu, bv)
	if o.reverse {
		return c > 0
	} else {
		return c < 0
	}
}
func (o *jsorter) Swap(i, j int) {
	o.pp[i], o.pp[j] = o.pp[j], o.pp[i]
}

func N2_sorted(vec_1 U, vec_2 V, cmp_1 U, cmp_2 V, key_1 U, key_2 V, reverse_1 U, reverse_2 V) (U, V) {
	ps := ListJ(vec_1, vec_2)
	if len(ps) == 0 {
		z := MkListJ(nil)
		return inline.MkPJ(z)
	}
	zs := CopyPJs(ps)
	if len(zs) > 1 {
		o := newJSorter(zs)
		if !inline.NullJ(cmp_1, cmp_2) {
			o.cmp = func(a1 U, a2 V, b1 U, b2 V) int {
				u, v := JCALL_2(cmp_1, cmp_2, a1, a2, b1, b2)
				z := IntJ(u, v)
				return int(z)
			}
		}
		if !inline.NullJ(key_1, key_2) {
			o.key = func(u U, v V) (U, V) {
				return JCALL_1(key_1, key_2, u, v)
			}
		}
		o.reverse = BoolJ(reverse_1, reverse_2)

		sort.Sort(o)
	}
	z := MkListJ(zs)
	return inline.MkPJ(z)
}
