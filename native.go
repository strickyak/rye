package rye

import "sort"

func N_byt(a P) P {
	switch x := a.(type) {
	case *PStr:
		bb := make([]byte, len(x.S))
		copy(bb, x.S)
		return MkByt(bb)
	case *PByt:
		return a
	case *PInt:
		return MkByt(make([]byte, int(x.N)))
	}
	return MkByt(a.Bytes())
}

func N_range(a P) P {
	n := a.Int()
	v := make([]P, n)
	for i := int64(0); i < n; i++ {
		v[i] = MkInt(i)
	}
	return MkList(v)
}

// Types for sorting.
type AnyPs []P

func (o AnyPs) Len() int { return len(o) }
func (o AnyPs) Less(i, j int) bool {
	return (o[i].Compare(o[j]) < 0)
}
func (o AnyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type StringyPs []P

func (o StringyPs) Len() int { return len(o) }
func (o StringyPs) Less(i, j int) bool {
	return (o[i].(*PStr).S < o[j].(*PStr).S)
}
func (o StringyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type IntyPs []P

func (o IntyPs) Len() int { return len(o) }
func (o IntyPs) Less(i, j int) bool {
	return (o[i].(*PInt).N < o[j].(*PInt).N)
}
func (o IntyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

type FloatyPs []P

func (o FloatyPs) Len() int { return len(o) }
func (o FloatyPs) Less(i, j int) bool {
	return (o[i].(*PFloat).F < o[j].(*PFloat).F)
}
func (o FloatyPs) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func N_sorted(a P) P {
	ps := CopyPs(a.List())
	if len(ps) == 0 {
		return MkList([]P{})
	}
	switch ps[0].(type) {
	// TODO -- Make heterogenous lists work.
	case *PStr:
		sort.Sort(StringyPs(ps))
	case *PInt:
		sort.Sort(IntyPs(ps))
	case *PFloat:
		sort.Sort(FloatyPs(ps))
	default:
		sort.Sort(AnyPs(ps))
	}
	return MkList(ps)
}
