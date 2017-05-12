// +build prego

package rye

import "fmt"
import R "reflect"
import "sort"
import "strings"

func N_rye_what(a M) M {
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
}

func N_set(a M) M {
	z := make(Scope)
	// If a is None or empty, leave it empty.
	if JBool(a) {
		for _, e := range JList(a) {
			z[JString(e)] = True
		}
	}
	return MkSet(z)
}

func N_dict(args, kw M) M {
	var d *PDict
	vec := JList(args)
	switch len(vec) {
	case 0:
		d = MkDict(make(Scope)).X.(*PDict)
	case 1:
		a := vec[0]
		switch t := a.X.(type) {
		case *PNone:
			d = MkDict(make(Scope)).X.(*PDict)
		case *PList:
			d = MkDictFromPairs(t.List()).X.(*PDict)
		case *PDict:
			d = MkDictCopy(Scope(t.Dict())).X.(*PDict)
		case *PModule:
			d = MkDictCopy(Scope(t.Dict())).X.(*PDict)
		case *PGo:
			d = MkDict(t.Dict()).X.(*PDict)
		default:
			d = MkDictFromPairs(t.List()).X.(*PDict)
			//?// panic(fmt.Sprintf("Bad arg to dict(), flavor=%d", a.Flavor()))
		}
	default:
		panic("Too many args to dict()")
	}
	kwd := kw.X.(*PDict)
	//#if m
	kwd.mu.Lock()
	//#endif
	for k, v := range kwd.ppp {
		d.ppp[k] = v
	}
	//#if m
	kwd.mu.Unlock()
	//#endif
	return MkX(&d.PBase)
}

func N_byt(a M) M {
	if a.X != nil {
		switch x := a.X.(type) {
		case *PByt:
			// byt() can be used to copy a byte array.
			bb := make([]byte, len(x.YY))
			copy(bb, x.YY)
			return MkByt(bb)
		}
	}
	// *PStr makes a copy already inside a.Bytes().
	// So does *PList.
	return MkByt(JBytes(a))
}

func N_mkbyt(a M) M {
	return MkByt(make([]byte, int(JInt(a))))
}

func N_range(a M) M {
	n := JInt(a)
	v := make([]M, n)
	for i := int64(0); i < n; i++ {
		v[i] = MkInt(i)
	}
	return MkList(v)
}

// For Sorting.

type sorter struct {
	pp      []M
	reverse bool
	cmp     func(a, b M) int
	key     func(a M) M
}

func newSorter(pp []M) *sorter {
	return &sorter{
		pp:  pp,
		cmp: func(a, b M) int { return TripCompare(a, b) },
	}
}

func (o *sorter) Len() int { return len(o.pp) }
func (o *sorter) Less(i, j int) bool {
	a := o.pp[i]
	b := o.pp[j]
	if o.key != nil {
		a = o.key(a)
		b = o.key(b)
	}
	c := o.cmp(a, b)
	if o.reverse {
		return c > 0
	} else {
		return c < 0
	}
}
func (o *sorter) Swap(i, j int) {
	o.pp[i], o.pp[j] = o.pp[j], o.pp[i]
}

func N_sorted(vec, cmp, key, reverse M) M {
	ps := JList(vec)
	if len(ps) == 0 {
		return MkList(nil)
	}
	zs := CopyPs(ps)
	if len(zs) > 1 {
		o := newSorter(zs)
		if cmp != None {
			o.cmp = func(a, b M) int { return int(JInt(CALL_2(cmp, a, b))) }
		}
		if key != None {
			o.key = func(a M) M { return CALL_1(key, a) }
		}
		o.reverse = JBool(reverse)

		sort.Sort(o)
	}
	return MkList(zs)
}
