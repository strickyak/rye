package rye

import "fmt"
import R "reflect"
import "sort"
import "strings"

func N_rye_what(a B) B {
	switch x := a.Self.(type) {
	case *PGo:
		z := fmt.Sprintf("PGo|%s|%s", x.V.Type().Kind(), x.V.Type())
		switch y := x.V.Interface().(type) {
		case R.Value:
			z += fmt.Sprintf("|%s|%s", y.Type().Kind(), y.Type())
		}
		return MkStr(z)
	}
	return MkStr(strings.TrimLeft(fmt.Sprintf("%T", a), "*"))
}

func N_set(a B) B {
  z := make(Scope)
  a_self := a.Self
  // If a is None or empty, leave it empty.
  if (a_self.Bool()) {
    for _, e := range a_self.List() {
      z[e.Self.String()] = True
    }
  }
  return MkSet(z)
}

func N_dict(args, kw B) B {
	var d *PDict
	vec := args.Self.(*PList)
	switch len(vec.PP) {
	case 0:
		d = MkDict(make(Scope)).Self.(*PDict)
	case 1:
		a := vec.PP[0]
		switch t := a.Self.(type) {
		case *PNone:
			d = MkDict(make(Scope)).Self.(*PDict)
		case *PList:
			d = MkDictFromPairs(t.List()).Self.(*PDict)
		case *PDict:
			d = MkDictCopy(Scope(t.Dict())).Self.(*PDict)
		case *PGo:
			d = MkDict(t.Dict()).Self.(*PDict)
		default:
			d = MkDictFromPairs(t.List()).Self.(*PDict)
			//?// panic(fmt.Sprintf("Bad arg to dict(), flavor=%d", a.Flavor()))
		}
	default:
		panic("Too many args to dict()")
	}
	kwd := kw.Self.(*PDict)
	kwd.mu.Lock()
	for k, v := range kwd.ppp {
		d.ppp[k] = v
	}
	kwd.mu.Unlock()
	return &d.PBase
}

func N_byt(a B) B {
	switch x := a.Self.(type) {
	case *PByt:
		// byt() can be used to copy a byte array.
		bb := make([]byte, len(x.YY))
		copy(bb, x.YY)
		return MkByt(bb)
	}
	// *PStr makes a copy already inside a.Bytes().
	// So does *PList.
	return MkByt(a.Self.Bytes())
}

func N_mkbyt(a B) B {
	return MkByt(make([]byte, int(a.Self.Int())))
}

func N_range(a B) B {
	n := a.Self.Int()
	v := make([]B, n)
	for i := int64(0); i < n; i++ {
		v[i] = MkInt(i)
	}
	return MkList(v)
}

// For Sorting.

type sorter struct {
	pp      []B
	reverse bool
	cmp     func(a, b B) int
	key     func(a B) B
}

func newSorter(pp []B) *sorter {
	return &sorter{
		pp:  pp,
		cmp: func(a, b B) int { return a.Self.Compare(b) },
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

func N_sorted(vec, cmp, key, reverse B) B {
	ps := vec.Self.List()
	if len(ps) == 0 {
		return MkList(nil)
	}
	zs := CopyPs(ps)
	if len(zs) > 1 {
		o := newSorter(zs)
		if cmp != None {
			o.cmp = func(a, b B) int { return int(call_2(cmp, a, b).Self.Int()) }
		}
		if key != None {
			o.key = func(a B) B { return call_1(key, a) }
		}
		o.reverse = reverse.Self.Bool()

		sort.Sort(o)
	}
	return MkList(zs)
}
