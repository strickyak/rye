// +build TEMPLATE

package runtime

import (
	"bytes"
	"hash/crc64"
	"sort"
	//if 'm' {
	"sync"
	//}
)

type T struct {
	PBase
	ppp Scope
	//if 'm' {
	mu sync.Mutex
	//}
}

func (o *T) Hash() int64 {
	var z int64
	if 'm' {
		o.mu.Lock()
	}
	for k, v := range o.ppp {
		z += int64(crc64.Checksum([]byte(k), CrcPolynomial))
		z += JHash(v) // TODO better
	}
	if 'm' {
		o.mu.Unlock()
	}
	return z
}
func (o *T) Pickle(w *bytes.Buffer) {
	if 'm' {
		o.mu.Lock()
		defer o.mu.Unlock()
	}
	l := int64(len(o.ppp))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypDict + n))
	RypWriteInt(w, l)
	for k, v := range o.ppp {
		MkStr(k).Pickle(w)
		v.Pickle(w)
	}
}
func (o *T) Contents() interface{} { return o.ppp }
func (o *T) Bool() bool            { return len(o.ppp) != 0 }
func (o *T) NotContains(a M) bool  { return !o.Contains(a) }
func (o *T) Contains(a M) bool {
	key := JString(a)
	if 'm' {
		o.mu.Lock()
	}
	_, ok := o.ppp[key]
	if 'm' {
		o.mu.Unlock()
	}
	return ok
}
func (o *T) Len() int { return len(o.ppp) }
func (o *T) SetItem(a M, x M) {
	key := JString(a)
	if 'm' {
		o.mu.Lock()
	}
	o.ppp[key] = x
	if 'm' {
		o.mu.Unlock()
	}
}
func (o *T) GetItem(a M) M {
	key := JString(a)
	if 'm' {
		o.mu.Lock()
	}
	z, ok := o.ppp[key]
	if 'm' {
		o.mu.Unlock()
	}
	if !ok {
		panic(F("T: KeyError: %q", key))
	}
	return z
}
func (o *T) String() string { return o.Repr() }
func (o *T) PType() M       { return G_dict }
func (o *T) RType() string  { return "dict" }
func (o *T) Repr() string {
	if 'm' {
		o.mu.Lock()
	}
	vec := make(KVSlice, 0, len(o.ppp))
	for k, v := range o.ppp {
		vec = append(vec, KV{k, v})
	}
	if 'm' {
		o.mu.Unlock()
	}

	sort.Sort(vec)
	buf := bytes.NewBufferString("{")
	n := len(vec)
	for i := 0; i < n; i++ {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(F("%q: %s", vec[i].Key, JRepr(vec[i].Value)))
	}
	buf.WriteString("}")
	return buf.String()
}
func (o *T) Start(int) {}
func (o *T) Enough()   {}
func (o *T) Iter() Receiver {
	var keys []M
	if 'm' {
		o.mu.Lock()
	}
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	if 'm' {
		o.mu.Unlock()
	}
	z := &PListIter{PP: keys}
	Forge(z)
	return z
}
func (o *T) List() []M {
	var keys []M
	if 'm' {
		o.mu.Lock()
	}
	for k, _ := range o.ppp {
		keys = append(keys, MkStr(k))
	}
	if 'm' {
		o.mu.Unlock()
	}
	return keys
}
func (o *T) Dict() Scope {
	return o.ppp
}
func (o *T) DelItem(i M) {
	key := JString(i)
	if 'm' {
		o.mu.Lock()
	}
	delete(o.ppp, key)
	if 'm' {
		o.mu.Unlock()
	}
}
func (o *T) Compare(a M) int {
	switch b := a.X.(type) {
	case *T:
		okeys := o.List()
		akeys := b.List()
		ostrs := make([]string, len(okeys))
		astrs := make([]string, len(akeys))
		for i, x := range okeys {
			ostrs[i] = JString(x)
		}
		for i, x := range akeys {
			astrs[i] = JString(x)
		}
		sort.Strings(ostrs)
		sort.Strings(astrs)
		olist := make([]M, len(okeys)*2)
		alist := make([]M, len(akeys)*2)
		if 'm' {
			o.mu.Lock()
		}
		for i, x := range ostrs {
			olist[i*2] = MkStr(x)
			olist[i*2+1] = o.ppp[x]
		}
		if 'm' {
			o.mu.Unlock()
			b.mu.Lock()
		}
		for i, x := range astrs {
			alist[i*2] = MkStr(x)
			alist[i*2+1] = b.ppp[x]
		}
		if 'm' {
			b.mu.Unlock()
		}
		return JCompare(MkList(olist), MkList(alist))
	}
	return StrCmp(JString(o.PType()), JString(JPType(a)))
}

func (self *T) clear() {
	if 'm' {
		self.mu.Lock()
	}
	self.ppp = make(map[string]M)
	if 'm' {
		self.mu.Unlock()
	}
}
func (self *T) copy() M {
	z := make(map[string]M)
	if 'm' {
		self.mu.Lock()
	}
	for k, v := range self.ppp {
		z[k] = v
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkDict(z)
}
func (self *T) items() M {
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for k, v := range self.ppp {
		z = append(z, MkTuple([]M{MkStr(k), v}))
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
}
func (self *T) keys() M {
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for k, _ := range self.ppp {
		z = append(z, MkStr(k))
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
}
func (self *T) values() M {
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for _, v := range self.ppp {
		z = append(z, v)
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
}
func (self *T) get(key M, dflt M) M {
	k := JString(key)
	if 'm' {
		self.mu.Lock()
	}
	z, ok := self.ppp[k]
	if 'm' {
		self.mu.Unlock()
	}
	if ok {
		return z
	}
	return dflt
}
