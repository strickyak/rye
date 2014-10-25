package rye

import "sort"

func N_dict(args, kw P) P {
	var d *PDict
	vec := args.(*PList)
	switch len(vec.PP) {
	case 0:
		d = MkDictFromPairs(nil)
	case 1:
		a := vec.PP[0]
		switch a.Flavor() {
		case ListLike:
			d = MkDictFromPairs(a.List())
		case DictLike:
			d = MkDictCopy(Scope(a.Dict()))
		case GoLike:
			d = MkDict(a.Dict())
		default:
			panic("Bad arg to dict()")
		}
	default:
		panic("Too many args to dict()")
	}
	for k, v := range kw.(*PDict).PPP {
		d.PPP[k] = v
	}
	return d
}

func N_byt(a P) P {
	switch x := a.(type) {
	case *PByt:
		// byt() can be used to copy a byte array.
		bb := make([]byte, len(x.YY))
		copy(bb, x.YY)
		return MkByt(bb)
	}
	// *PStr makes a copy already inside a.Bytes().
	// So does *PList.
	return MkByt(a.Bytes())
}

func N_mkbyt(a P) P {
	return MkByt(make([]byte, int(a.Int())))
}

func N_range(a P) P {
	n := a.Int()
	v := make([]P, n)
	for i := int64(0); i < n; i++ {
		v[i] = MkInt(i)
	}
	return MkList(v)
}

// For Sorting.

type sorter struct {
	pp      []P
	reverse bool
	cmp     func(a, b P) int
	key     func(a P) P
}

func newSorter(pp []P) *sorter {
	return &sorter{
		pp:  pp,
		cmp: func(a, b P) int { return a.Compare(b) },
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

func N_sorted(vec, cmp, key, reverse P) P {
	ps := vec.List()
	if len(ps) == 0 {
		return MkList(nil)
	}
	zs := CopyPs(ps)
	if len(zs) > 1 {
		o := newSorter(zs)
		if cmp != None {
			o.cmp = func(a, b P) int { return int(call_2(cmp, a, b).Int()) }
		}
		if key != None {
			o.key = func(a P) P { return call_1(key, a) }
		}
		o.reverse = reverse.Bool()

		sort.Sort(o)
	}
	return MkList(zs)
}
