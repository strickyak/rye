package rye

import (
	"bytes"
	"fmt"
	"hash/crc64"
	"strconv"
	"strings"
)

type M struct {
	N int64
	S string
	X P
}

var MissingM = M{666, "---", nil}

/*
func (m M) Me() M {
	return m
}

func (m M) XXX() M {
	if m.X != nil {
		// return m.X
		panic("cannot XXX")
	} else if len(m.S) == 0 {
		// return m.N
		panic("cannot XXX on int")
	}

	// return m.S
	panic("cannot XXX on str")
}
*/

func (m M) Superclass() M {
	//fmt.Printf("Superclass %#v\n", m)
	if m.X != nil {
		return m.X.Superclass()
	}
	//fmt.Printf("Superclass2  %#v\n", m)
	//panic("cannot Superclass on int or str")
  return None
}

func (m M) SetItem(i M, x M) {
	if m.X != nil {
		m.X.SetItem(i, x)
		return
	} else if len(m.S) == 0 {
		panic("cannot SetItem on int")
	}

	panic("cannot SetItem on str")
}

func (m M) UnaryMinus() M {
	if m.X != nil {
		return m.X.UnaryMinus()
	} else if len(m.S) == 0 {
		return M{N: -m.N}
	}

	panic("cannot UnaryMinus on str")
}

func (m M) UnaryInvert() M {
	if m.X != nil {
		return m.X.UnaryInvert()
	} else if len(m.S) == 0 {
		return M{N: ^m.N}
	}

	panic("cannot UnaryInvert on str")
}

func (m M) Dict() Scope {
	if m.X != nil {
		return m.X.Dict()
	} else if len(m.S) == 0 {
		panic("cannot Dict on int")
	}

	panic("cannot Dict on str")
}

func (m M) DelItem(i M) {
	if m.X != nil {
		m.X.DelItem(i)
		return
	} else if len(m.S) == 0 {
		panic("cannot DelItem on int")
	}

	panic("cannot DelItem on str")
}

func (m M) DelItemSlice(i, j M) {
	if m.X != nil {
		m.X.DelItemSlice(i, j)
		return
	} else if len(m.S) == 0 {
		panic("cannot DelItemSlice on int")
	}

	panic("cannot DelItemSlice on str")
}

func (m M) Call(aa ...M) M {
	if m.X != nil {
		return m.X.Call(aa...)
	} else if len(m.S) == 0 {
		panic("cannot Call on int")
	}

	panic("cannot Call on str")
}

func (m M) Invoke(field string, aa ...M) M {
	if m.X != nil {
		return m.X.Invoke(field, aa...)
	} else if len(m.S) == 0 {
		panic("cannot Invoke on int")
	}

	// TODO: dont MkBStr
	return MkBStr(m.S).Self.Invoke(field, aa...)
}

func (m M) IsNot(a M) bool {
	return !m.Is(a)
}

func (m M) Is(a M) bool {
	if m.X != nil {
		return m.X == a.X
	} else if len(m.S) == 0 {
		return (a.X == nil) && (len(a.S) == 0) && m.N == a.N
	}

	return (a.X == nil) && (len(a.S) > 0) && m.S == a.S
}

func (m M) ToP() P {
	if m.X != nil {
		return m.X
	} else if len(m.S) == 0 {
		return MkBInt(m.N).Self
	}

	return MkBStr(m.S).Self
}

func (m M) List() []M {
	if m.X != nil {
		return m.X.List()
	} else if len(m.S) == 0 {
		panic("cannot List() an int")
	}

	return MkBStr(m.S).Self.List() // Borrow PStr::List.
}

func (m M) FetchField(field string) M {
	if m.X != nil {
		return m.X.FetchField(field)
	}
	panic("cannot FetchField")
}

func (m M) StoreField(field string, p M) {
	if m.X != nil {
		m.X.StoreField(field, p)
		return
	}
	panic("cannot FetchField")
}

func (m M) Callable() bool {
	if m.X != nil {
		return m.X.Callable()
	}
	return false
}

func (m M) EQ(a M) bool {
	if m.X != nil {
		return m.X.EQ(a)
	} else if len(m.S) == 0 {
		// int ==
		if a.X == nil && len(a.S) == 0 {
			return m.N == a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int == int
			return m.N == t.N
		case *PFloat:
			// int == float
			return float64(m.N) == t.F
		case *PBool:
			// int == bool
			return m.N == t.Int()
		}
	}

	// str ==
	if a.X == nil && len(a.S) > 0 {
		return m.S == a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str == str
		return m.S == t.S
	case *PByt:
		// str == byt
		return m.S == string(t.YY)
	}

	return false
}

func (m M) NE(a M) bool {
	if m.X != nil {
		return m.X.NE(a)
	} else if len(m.S) == 0 {
		// int !=
		if a.X == nil && len(a.S) == 0 {
			return m.N != a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int != int
			return m.N != t.N
		case *PFloat:
			// int != float
			return float64(m.N) != t.F
		case *PBool:
			// int != bool
			return m.N != t.Int()
		}
	}

	// str !=
	if a.X == nil && len(a.S) > 0 {
		return m.S != a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str != str
		return m.S != t.S
	case *PByt:
		// str != byt
		return m.S != string(t.YY)
	}

	return true
}

func (m M) LT(a M) bool {
	if m.X != nil {
		return m.X.LT(a)
	} else if len(m.S) == 0 {
		// int *
		if a.X == nil && len(a.S) == 0 {
			return m.N < a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int < int
			return m.N < t.N
		case *PFloat:
			// int * float
			return float64(m.N) < t.F
		case *PBool:
			// int < bool
			return m.N < t.Int()
		}
		return m.N < a.Int()
	}

	// str <
	if a.X == nil && len(a.S) > 0 {
		return m.S < a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str < str
		return m.S < t.S
	case *PByt:
		// str < byt
		return m.S < string(t.YY)
	}

	panic(F("Cannot LT: str < %s", a.PType()))
}

func (m M) LE(a M) bool {
	if m.X != nil {
		return m.X.LE(a)
	} else if len(m.S) == 0 {
		// int *
		if a.X == nil && len(a.S) == 0 {
			return m.N <= a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int <= int
			return m.N <= t.N
		case *PFloat:
			// int <= float
			return float64(m.N) <= t.F
		case *PBool:
			// int <= bool
			return m.N <= t.Int()
		}
		return m.N <= a.Int()
	}

	// str <=
	if a.X == nil && len(a.S) > 0 {
		return m.S <= a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str <= str
		return m.S <= t.S
	case *PByt:
		// str <= byt
		return m.S <= string(t.YY)
	}

	panic(F("Cannot LE: str <= %s", a.PType()))
}

func (m M) GT(a M) bool {
	if m.X != nil {
		return m.X.GT(a)
	} else if len(m.S) == 0 {
		// int >
		if a.X == nil && len(a.S) == 0 {
			return m.N > a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int > int
			return m.N > t.N
		case *PFloat:
			// int > float
			return float64(m.N) > t.F
		case *PBool:
			// int > bool
			return m.N > t.Int()
		}
		return m.N > a.Int()
	}

	// str >
	if a.X == nil && len(a.S) > 0 {
		return m.S > a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str > str
		return m.S > t.S
	case *PByt:
		// str > byt
		return m.S > string(t.YY)
	}

	panic(F("Cannot GT: str > %s", a.PType()))
}

func (m M) GE(a M) bool {
	if m.X != nil {
		return m.X.GE(a)
	} else if len(m.S) == 0 {
		// int >=
		if a.X == nil && len(a.S) == 0 {
			return m.N >= a.N
		}

		switch t := a.X.(type) {
		case *PInt:
			// int >= int
			return m.N >= t.N
		case *PFloat:
			// int * float
			return float64(m.N) >= t.F
		case *PBool:
			// int >= bool
			return m.N >= t.Int()
		}
		return m.N >= a.Int()
	}

	// str >=
	if a.X == nil && len(a.S) > 0 {
		return m.S >= a.S
	}
	switch t := a.X.(type) {
	case *PStr:
		// str >= str
		return m.S >= t.S
	case *PByt:
		// str >= byt
		return m.S >= string(t.YY)
	}

	panic(F("Cannot GE: str >= %s", a.PType()))
}

func (m M) CanInt() bool {
	if m.X != nil {
		return m.X.CanInt()
	} else if len(m.S) == 0 {
		return true
	}
	return false
}

func (m M) Int() int64 {
	if m.X != nil {
		return m.X.Int()
	} else if len(m.S) == 0 {
		return m.N
	}
	panic("cannot Int() a string")
}

func (m M) CanFloat() bool {
	if m.X != nil {
		return m.X.CanFloat()
	} else if len(m.S) == 0 {
		return true
	}
	return false
}

func (m M) Float() float64 {
	if m.X != nil {
		return m.X.Float()
	} else if len(m.S) == 0 {
		return float64(m.N)
	}
	panic("cannot Int() a string")
}

func (m M) Hash() int64 {
	if m.X != nil {
		return m.X.Hash()
	} else if len(m.S) == 0 {
		return m.N
	}

	return int64(crc64.Checksum([]byte(m.S), CrcPolynomial))
}
func (m M) Iter() Nexter {
	if m.X != nil {
		return m.X.Iter()
	} else if len(m.S) == 0 {
		panic("cannot Iter() on int")
	}

	var pp []M
	for _, r := range m.S {
		pp = append(pp, MkStr(string(r)))
	}
	z := &PListIter{PP: pp}
	Forge(z)
	return z
}
func (m M) Pickle(w *bytes.Buffer) {
	if m.X != nil {
		m.X.Pickle(w)
		return
	} else if len(m.S) == 0 {
		n := RypIntLenMinus1(m.N)
		w.WriteByte(byte(RypInt + n))
		RypWriteInt(w, m.N)
		return
	}

	l := int64(len(m.S))
	n := RypIntLenMinus1(l)
	w.WriteByte(byte(RypStr + n))
	RypWriteInt(w, l)
	w.WriteString(m.S)
}

func (m M) Contents() interface{} {
	if m.X != nil {
		return m.X.Contents()
	} else if len(m.S) == 0 {
		return m.N
	}
	return m.S
}
func (m M) Bool() bool {
	if m.X != nil {
		return m.X.Bool()
	} else if len(m.S) == 0 {
		return m.N != 0
	}
	return true
}
func (m M) GetItem(x M) M {
	if m.X != nil {
		return m.X.GetItem(x)
	} else if len(m.S) == 0 {
		panic("cannot GetItem(0 on int")
	}
	i := x.Int()
	if i < 0 {
		i += int64(len(m.S))
	}
	return MkStr(m.S[i : i+1])
}

func (m M) GetItemSlice(x, y, z M) M {
	if m.X != nil {
		return m.X.GetItemSlice(x, y, z)
	} else if len(m.S) == 0 {
		panic("cannot GetItemSlice(0 on int")
	}
	var i, j int64
	n := int64(len(m.S))
	if x == None {
		i = 0
	} else {
		i = x.Int()
		if i < 0 {
			i += n
		}
		if i < 0 {
			panic(F("First slicing index on PStr too small: %d", i))
		}
	}
	if i > n {
		panic(F("First slicing index on PStr too large: %d > len: %d", i, n))
	}
	if y == None {
		j = n
	} else {
		j = y.Int()
		if j < 0 {
			j += n
		}
		if j < 0 {
			panic(F("Second slicing index on PStr too small: %d", y.Int()))
		}
	}
	if j > n {
		j = n // Python lets you specify too big second index.
		// panic(F("Second slicing index on PStr too large: %d > len: %d", j, n))
	}
	return MkStr(m.S[i:j])
}

func (m M) Mod(a M) M {
	if m.X != nil {
		return m.X.Mod(a)
	} else if len(m.S) == 0 {
		if a.CanInt() {
			return MkInt(m.N % a.Int())
		} else {
			panic("cannot Mod() int with non-int")
		}
	}

	if a.X != nil {
		switch t := a.X.(type) {
		case *PTuple:
			z := make([]interface{}, len(t.PP))
			for i, e := range t.PP {
				z[i] = e.Contents()
			}
			return MkStr(F(m.S, z...))
		}
	}
	return MkStr(F(m.S, a.Contents()))
}

func (m M) Mul(a M) M {
	if m.X != nil {
		// X *
		return m.X.Mul(a)
	} else if len(m.S) == 0 {
		// int *
		if a.X == nil {
			if len(a.S) == 0 {
				// int * int
				return M{N: m.N * a.N}
			} else {
				// int * str
				return MkStr(strings.Repeat(a.S, int(m.N)))
			}
		}

		switch t := a.X.(type) {
		case *PInt:
			// int * int
			return M{N: m.N * t.N}
		case *PList:
			// int * list
			return RepeatList(t.PP, m.N)
		case *PFloat:
			// int * float
			return MkFloat(float64(m.N) * t.F)
		case *PStr:
			// int * str
			return MkStr(strings.Repeat(t.S, int(m.N)))
		case *PByt:
			// int * byt
			return MkByt(bytes.Repeat(t.YY, int(m.N)))
		}
		return MkInt(m.N * a.Int())
	}

	// str *
	switch t := a.X.(type) {
	case nil:
		if a.S == "" {
			// str * int
			return M{N: m.N * a.N}
		} else {
			// str * str
			panic("cannot Mul: str * str")
		}
	case *PInt:
		return MkStr(strings.Repeat(m.S, int(t.N)))
	}

	panic(F("Cannot multiply: str * %s", a.PType()))
}
func (m M) NotContains(a M) bool {
	return !m.Contains(a)
}
func (m M) Contains(a M) bool {
	if m.X != nil {
		return m.X.Contains(a)
	} else if len(m.S) == 0 {
		panic("connot Contains() on int")
	}

	if a.CanStr() {
		return strings.Contains(m.S, a.Str())
	}
	panic(F("str cannot Contains() non-str: %s", a.PType()))
}
func (m M) Add(a M) M {
	if m.X != nil {
		return m.X.Add(a)
	} else if len(m.S) == 0 {
		// int +
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N + a.N}
		}

		switch t := a.X.(type) {
		case *PInt:
			// int + int
			return M{N: m.N + t.N}
		case *PBool:
			// int + bool
			return M{N: m.N + t.Int()}
		case *PFloat:
			// int + float
			return MkFloat(float64(m.N) + t.F)
		}
		//println("Add...int...", m.N)
		//println("Add...int...", a.Int())
		return MkInt(m.N + a.Int())
	}

	// str +
	switch t := a.X.(type) {
	case nil:
		if a.S == "" {
			// str + int
			panic("cannot Add: str + str")
		} else {
			// str + str
			return MkStr(m.S + a.S)
		}
	case *PStr:
		return MkStr(m.S + t.S)
	case *PByt:
		return MkStr(m.S + string(t.YY))
	}

	panic(F("Cannot add: str + %s", a.PType()))
}

func (m M) Sub(a M) M {
	if m.X != nil {
		return m.X.Sub(a)
	} else if len(m.S) == 0 {
		// int -
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N - a.N}
		}

		switch t := a.X.(type) {
		case nil:
			if len(a.S) == 0 {
				// int - int
				return M{N: m.N - a.N}
			} else {
				// int - str
				panic("cannot Sub: int - str")
			}
		case *PInt:
			// int - int
			return M{N: m.N - t.N}
		case *PFloat:
			// int - float
			return MkFloat(float64(m.N) - t.F)
		}
		return MkInt(m.N - a.Int())
	}

	panic(F("Cannot Sub: str - %s", a.PType()))
}

func (m M) Div(a M) M {
	if m.X != nil {
		return m.X.Sub(a)
	} else if len(m.S) == 0 {
		if a.CanInt() {
			return MkInt(m.N / a.Int())
		} else if a.CanFloat() {
			return MkFloat(float64(m.N) / a.Float())
		}
	}
	panic("cannot Div on str")
}

func (m M) BitAnd(a M) M {
	if m.X != nil {
		return m.X.BitAnd(a)
	} else if len(m.S) == 0 {
		if a.CanInt() {
			return MkInt(m.N & a.Int())
		} else {
			panic("cannot BitAnd on int with non-int")
		}
	}
	panic("cannot BitAnd on str")
}

func (m M) BitOr(a M) M {
	if m.X != nil {
		return m.X.BitOr(a)
	} else if len(m.S) == 0 {
		if a.CanInt() {
			return MkInt(m.N | a.Int())
		} else {
			panic("cannot BitOr on int with non-int")
		}
	}
	panic("cannot BitOr on str")
}

func (m M) BitXor(a M) M {
	if m.X != nil {
		return m.X.BitXor(a)
	} else if len(m.S) == 0 {
		if a.CanInt() {
			return MkInt(m.N ^ a.Int())
		} else {
			panic("cannot BitXor on int with non-int")
		}
	}
	panic("cannot BitXor on str")
}

func (m M) ShiftLeft(a M) M {
	if m.X != nil {
		return m.X.ShiftLeft(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N << uint64(a.N)}
		}
		if a.CanInt() {
			return MkInt(m.N << uint64(a.Int()))
		} else {
			panic("cannot ShiftLeft on int with non-int")
		}
	}
	panic("cannot ShiftLeft on str")
}

func (m M) ShiftRight(a M) M {
	if m.X != nil {
		return m.X.ShiftRight(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N >> uint64(a.N)}
		}
		if a.CanInt() {
			return MkInt(m.N >> uint64(a.Int()))
		} else {
			panic("cannot ShiftRight on int with non-int")
		}
	}
	panic("cannot ShiftRight on str")
}

func (m M) UnsignedShiftRight(a M) M {
	if m.X != nil {
		return m.X.UnsignedShiftRight(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: int64(uint64(m.N) >> uint64(a.N))}
		}
		if a.CanInt() {
			return MkInt(int64(uint64(m.N) >> uint64(a.Int())))
		} else {
			panic("cannot UnsignedShiftRight on int with non-int")
		}
	}
	panic("cannot UnsignedShiftRight on str")
}

func (m M) Compare(a M) int {
	if m.X != nil {
		return m.X.Compare(a)
	} else if len(m.S) == 0 {
		// TODO
		x := a.Int()
		switch {
		case m.N < x:
			return -1
		case m.N > x:
			return 1
		}
		return 0
	}
	// string:
	if a.CanStr() {
		x := a.Str()
		switch {
		case m.S < x:
			return -1
		case m.S > x:
			return 1
		}
		return 0
	}
	return StrCmp("str", a.PType().String())
}
func (m M) ForceInt() int64 {
	if m.X != nil {
		return m.X.ForceInt()
	} else if len(m.S) == 0 {
		return m.N
	}
	z, err := strconv.ParseInt(m.S, 10, 64)
	if err != nil {
		panic(F("PStr::ForceInt: ParseInt: %v", err))
	}
	return z
}
func (m M) ForceFloat() float64 {
	if m.X != nil {
		return m.X.ForceFloat()
	} else if len(m.S) == 0 {
		return float64(m.N)
	}
	z, err := strconv.ParseFloat(m.S, 64)
	if err != nil {
		panic(F("PStr::ForceFloat: ParseFloat: %v", err))
	}
	return z
}
func (m M) Bytes() []byte {
	if m.X != nil {
		return m.X.Bytes()
	} else if len(m.S) == 0 {
		panic("cannot Bytes() on an integer")
	}
	return []byte(m.S)
}
func (m M) Len() int {
	if m.X != nil {
		return m.X.Len()
	} else if len(m.S) == 0 {
		panic("cannot Len() on an integer")
	}
	return len(m.S)
}

func (m M) Repr() string {
	if m.X != nil {
		return m.X.Repr()
	} else if len(m.S) == 0 {
		return m.String()
	}
	return ReprStringLikeInPython(m.S)
}
func (m M) PType() M {
	if m.X != nil {
		return m.X.PType()
	} else if len(m.S) == 0 {
		return G_int
	}
	return G_str
}
func (m M) RType() string {
	if m.X != nil {
		return m.X.RType()
	} else if len(m.S) == 0 {
		return "int"
	}
	return "str"
}
func (m M) CanStr() bool {
	if m.X != nil {
		return m.X.CanStr()
	} else if len(m.S) == 0 {
		return false
	}
	return true
}
func (m M) Str() string {
	if m.X != nil {
		return m.X.Str()
	} else if len(m.S) == 0 {
		panic("cannot Str() an int")
	}
	return m.S
}
func (m M) String() string {
	if m.X != nil {
		return m.X.String()
	} else if len(m.S) == 0 {
		return fmt.Sprintf("%d", m.N)
	}
	return m.S
}

//////////////////////

func MMkStr(s string) M {
	if len(s) == 0 {
	}
	return M{S: s}
}
func MMkint(n int) M {
	return M{N: int64(n)}
}
func MMkInt(n int64) M {
	return M{N: n}
}
func MkX(x B) M {
	switch t := x.Self.(type) {
	case *PInt:
		return MMkInt(t.N)
	case *PStr:
		if len(t.S) == 0 {
			return EmptyStr
		}
		return MMkStr(t.S)
	}
	return M{X: x.Self}
}

func MPMkList(pp []M) *PList    { z := &PList{PP: pp}; Forge(z); return z }
func MPMkDict(ppp Scope) *PDict { z := &PDict{ppp: ppp}; Forge(z); return z }

//////////
