// +build prego

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

func JSuperclass(m M) M {
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

func JUnaryMinus(m M) M {
	if m.X != nil {
		return m.X.UnaryMinus()
	} else if len(m.S) == 0 {
		return M{N: -m.N}
	}

	panic("cannot UnaryMinus on str")
}

func JUnaryInvert(m M) M {
	if m.X != nil {
		return m.X.UnaryInvert()
	} else if len(m.S) == 0 {
		return M{N: ^m.N}
	}

	panic("cannot UnaryInvert on str")
}

func JDict(m M) Scope {
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

func TripIsNot(m M, a M) bool {
	return !TripIs(m, a)
}

func TripIs(m M, a M) bool {
	if m.X != nil {
		return m.X == a.X
	} else if len(m.S) == 0 {
		return (a.X == nil) && (len(a.S) == 0) && m.N == a.N
	}

	return (a.X == nil) && (len(a.S) > 0) && m.S == a.S
}

func JList(m M) []M {
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

func JCallable(m M) bool {
	if m.X != nil {
		return m.X.Callable()
	}
	return false
}

func TripEQ(m M, a M) bool {
	if m.X != nil {
		return m.X.EQ(a)
	} else if len(m.S) == 0 {
		// int ==
		if a.X == nil && len(a.S) == 0 {
			return m.N == a.N
		}

		switch t := a.X.(type) {
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

func TripNE(m M, a M) bool {
	if m.X != nil {
		return m.X.NE(a)
	} else if len(m.S) == 0 {
		// int !=
		if a.X == nil && len(a.S) == 0 {
			return m.N != a.N
		}

		switch t := a.X.(type) {
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

func TripLT(m M, a M) bool {
	if m.X != nil {
		return m.X.LT(a)
	} else if len(m.S) == 0 {
		// int *
		if a.X == nil && len(a.S) == 0 {
			return m.N < a.N
		}

		switch t := a.X.(type) {
		case *PFloat:
			// int * float
			return float64(m.N) < t.F
		case *PBool:
			// int < bool
			return m.N < t.Int()
		}
		return m.N < JInt(a)
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

	panic(F("Cannot LT: str < %s", JPType(a)))
}

func TripLE(m M, a M) bool {
	if m.X != nil {
		return m.X.LE(a)
	} else if len(m.S) == 0 {
		// int *
		if a.X == nil && len(a.S) == 0 {
			return m.N <= a.N
		}

		switch t := a.X.(type) {
		case *PFloat:
			// int <= float
			return float64(m.N) <= t.F
		case *PBool:
			// int <= bool
			return m.N <= t.Int()
		}
		return m.N <= JInt(a)
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

	panic(F("Cannot LE: str <= %s", JPType(a)))
}

func TripGT(m M, a M) bool {
	if m.X != nil {
		return m.X.GT(a)
	} else if len(m.S) == 0 {
		// int >
		if a.X == nil && len(a.S) == 0 {
			return m.N > a.N
		}

		switch t := a.X.(type) {
		case *PFloat:
			// int > float
			return float64(m.N) > t.F
		case *PBool:
			// int > bool
			return m.N > t.Int()
		}
		return m.N > JInt(a)
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

	panic(F("Cannot GT: str > %s", JPType(a)))
}

func TripGE(m M, a M) bool {
	if m.X != nil {
		return m.X.GE(a)
	} else if len(m.S) == 0 {
		// int >=
		if a.X == nil && len(a.S) == 0 {
			return m.N >= a.N
		}

		switch t := a.X.(type) {
		case *PFloat:
			// int * float
			return float64(m.N) >= t.F
		case *PBool:
			// int >= bool
			return m.N >= t.Int()
		}
		return m.N >= JInt(a)
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

	panic(F("Cannot GE: str >= %s", JPType(a)))
}

func JCanInt(m M) bool {
	if m.X != nil {
		return m.X.CanInt()
	} else if len(m.S) == 0 {
		return true
	}
	return false
}

func JInt(m M) int64 {
	if m.X != nil {
		return m.X.Int()
	} else if len(m.S) == 0 {
		return m.N
	}
	panic("cannot Int() a string")
}

func JCanFloat(m M) bool {
	if m.X != nil {
		return m.X.CanFloat()
	} else if len(m.S) == 0 {
		return true
	}
	return false
}

func JFloat(m M) float64 {
	if m.X != nil {
		return m.X.Float()
	} else if len(m.S) == 0 {
		return float64(m.N)
	}
	panic("cannot Int() a string")
}

func JHash(m M) int64 {
	if m.X != nil {
		return m.X.Hash()
	} else if len(m.S) == 0 {
		return m.N
	}

	return int64(crc64.Checksum([]byte(m.S), CrcPolynomial))
}
func JIter(m M) Nexter {
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

func JContents(m M) interface{} {
	if m.X != nil {
		return m.X.Contents()
	} else if len(m.S) == 0 {
		return m.N
	}
	return m.S
}
func JBool(m M) bool {
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
	i := JInt(x)
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
		i = JInt(x)
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
		j = JInt(y)
		if j < 0 {
			j += n
		}
		if j < 0 {
			panic(F("Second slicing index on PStr too small: %d", JInt(y)))
		}
	}
	if j > n {
		j = n // Python lets you specify too big second index.
		// panic(F("Second slicing index on PStr too large: %d > len: %d", j, n))
	}
	return MkStr(m.S[i:j])
}

func TripMod(m M, a M) M {
	if m.X != nil {
		return m.X.Mod(a)
	} else if len(m.S) == 0 {
		if JCanInt(a) {
			return MkInt(m.N % JInt(a))
		} else {
			panic("cannot Mod() int with non-int")
		}
	}

	if a.X != nil {
		switch t := a.X.(type) {
		case *PTuple:
			z := make([]interface{}, len(t.PP))
			for i, e := range t.PP {
				z[i] = JContents(e)
			}
			return MkStr(F(m.S, z...))
		}
	}
	return MkStr(F(m.S, JContents(a)))
}

func TripMul(m M, a M) M {
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
		return MkInt(m.N * JInt(a))
	}

	// str *
	if a.X == nil {
		if a.S == "" {
			// str * int
			return M{N: m.N * a.N}
		} else {
			// str * str
			panic("cannot Mul: str * str")
		}
	}

	panic(F("Cannot multiply: str * %s", JPType(a)))
}
func TripNotContains(m M, a M) bool {
	return !TripContains(m, a)
}
func TripContains (m M, a M) bool {
	if m.X != nil {
		return m.X.Contains(a)
	} else if len(m.S) == 0 {
		panic("connot Contains() on int")
	}

	if JCanStr(a) {
		return strings.Contains(m.S, JStr(a))
	}
	panic(F("str cannot Contains() non-str: %s", JPType(a)))
}
func TripAdd(m M, a M) M {
	if m.X != nil {
		return m.X.Add(a)
	} else if len(m.S) == 0 {
		// int +
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N + a.N}
		}

		switch t := a.X.(type) {
		case *PBool:
			// int + bool
			return M{N: m.N + t.Int()}
		case *PFloat:
			// int + float
			return MkFloat(float64(m.N) + t.F)
		}
		//println("Add...int...", m.N)
		//println("Add...int...", JInt(a))
		return MkInt(m.N + JInt(a))
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

	panic(F("Cannot add: str + %s", JPType(a)))
}

func TripSub(m M, a M) M {
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
		case *PFloat:
			// int - float
			return MkFloat(float64(m.N) - t.F)
		}
		return MkInt(m.N - JInt(a))
	}

	panic(F("Cannot Sub: str - %s", JPType(a)))
}

func TripDiv(m M, a M) M {
	if m.X != nil {
		return m.X.Sub(a)
	} else if len(m.S) == 0 {
		if JCanInt(a) {
			return MkInt(m.N / JInt(a))
		} else if JCanFloat(a) {
			return MkFloat(float64(m.N) / JFloat(a))
		}
	}
	panic("cannot Div on str")
}

func TripBitAnd(m M, a M) M {
	if m.X != nil {
		return m.X.BitAnd(a)
	} else if len(m.S) == 0 {
		if JCanInt(a) {
			return MkInt(m.N & JInt(a))
		} else {
			panic("cannot BitAnd on int with non-int")
		}
	}
	panic("cannot BitAnd on str")
}

func TripBitOr(m M, a M) M {
	if m.X != nil {
		return m.X.BitOr(a)
	} else if len(m.S) == 0 {
		if JCanInt(a) {
			return MkInt(m.N | JInt(a))
		} else {
			panic("cannot BitOr on int with non-int")
		}
	}
	panic("cannot BitOr on str")
}

func TripBitXor(m M, a M) M {
	if m.X != nil {
		return m.X.BitXor(a)
	} else if len(m.S) == 0 {
		if JCanInt(a) {
			return MkInt(m.N ^ JInt(a))
		} else {
			panic("cannot BitXor on int with non-int")
		}
	}
	panic("cannot BitXor on str")
}

func TripShiftLeft(m M, a M) M {
	if m.X != nil {
		return m.X.ShiftLeft(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N << uint64(a.N)}
		}
		if JCanInt(a) {
			return MkInt(m.N << uint64(JInt(a)))
		} else {
			panic("cannot ShiftLeft on int with non-int")
		}
	}
	panic("cannot ShiftLeft on str")
}

func TripShiftRight(m M, a M) M {
	if m.X != nil {
		return m.X.ShiftRight(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: m.N >> uint64(a.N)}
		}
		if JCanInt(a) {
			return MkInt(m.N >> uint64(JInt(a)))
		} else {
			panic("cannot ShiftRight on int with non-int")
		}
	}
	panic("cannot ShiftRight on str")
}

func TripUnsignedShiftRight(m M, a M) M {
	if m.X != nil {
		return m.X.UnsignedShiftRight(a)
	} else if len(m.S) == 0 {
		if a.X == nil && len(a.S) == 0 {
			return M{N: int64(uint64(m.N) >> uint64(a.N))}
		}
		if JCanInt(a) {
			return MkInt(int64(uint64(m.N) >> uint64(JInt(a))))
		} else {
			panic("cannot UnsignedShiftRight on int with non-int")
		}
	}
	panic("cannot UnsignedShiftRight on str")
}

func TripCompare(m M, a M) int {
	if m.X != nil {
		return m.X.Compare(a)
	} else if len(m.S) == 0 {
		// TODO
		x := JInt(a)
		switch {
		case m.N < x:
			return -1
		case m.N > x:
			return 1
		}
		return 0
	}
	// string:
	if JCanStr(a) {
		x := JStr(a)
		switch {
		case m.S < x:
			return -1
		case m.S > x:
			return 1
		}
		return 0
	}
	return StrCmp("str", JString(JPType(a)))
}
func JForceInt(m M) int64 {
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
func JForceFloat(m M) float64 {
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
func JBytes(m M) []byte {
	if m.X != nil {
		return m.X.Bytes()
	} else if len(m.S) == 0 {
		panic("cannot Bytes() on an integer")
	}
	return []byte(m.S)
}
func JLen(m M) int {
	if m.X != nil {
		return m.X.Len()
	} else if len(m.S) == 0 {
		panic("cannot Len() on an integer")
	}
	return len(m.S)
}

func JRepr(m M) string {
	if m.X != nil {
		return m.X.Repr()
	} else if len(m.S) == 0 {
		return JString(m)
	}
	return ReprStringLikeInPython(m.S)
}
func JPType(m M) M {
	if m.X != nil {
		return m.X.PType()
	} else if len(m.S) == 0 {
		return G_int
	}
	return G_str
}
func JRType(m M) string {
	if m.X != nil {
		return m.X.RType()
	} else if len(m.S) == 0 {
		return "int"
	}
	return "str"
}
func JCanStr(m M) bool {
	if m.X != nil {
		return m.X.CanStr()
	} else if len(m.S) == 0 {
		return false
	}
	return true
}
func JStr(m M) string {
	if m.X != nil {
		return m.X.Str()
	} else if len(m.S) == 0 {
		panic("cannot Str() an int")
	}
	return m.S
}
func JString(m M) string {
	if m.X != nil {
		return m.X.String()
	} else if len(m.S) == 0 {
		return fmt.Sprintf("%d", m.N)
	}
	return m.S
}

//////////////////////

func MkX(x B) M {
	switch t := x.Self.(type) {
	case *PStr:
		if len(t.S) == 0 {
			return EmptyStr
		}
		return M{S: t.S}
	}
	return M{X: x.Self}
}

func MPMkList(pp []M) *PList    { z := &PList{PP: pp}; Forge(z); return z }
func MPMkDict(ppp Scope) *PDict { z := &PDict{ppp: ppp}; Forge(z); return z }

//////////
