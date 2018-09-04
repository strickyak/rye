// +build prego

package rye

func (inline) JCanInt(m M) bool {
	var __z__ bool
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.CanInt()
	} else if len(__a__.S) == 0 {
		__z__ = true
	} else {
		__z__ = false
	}
	return __z__
}

func (inline) JInt(m M) int64 {
	var __z__ int64
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.Int()
	} else if len(__a__.S) == 0 {
		__z__ = __a__.N
	} else {
		panic("cannot Int() a string")
	}
	return __z__
}

func (inline) JLen(m M) int {
	var __z__ int
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.Len()
	} else if len(__a__.S) == 0 {
		panic("cannot Len() on an integer")
	} else {
		__z__ = len(__a__.S)
	}
	return __z__
}

func (inline) JRepr(m M) string {
	var __z__ string
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.Repr()
	} else if len(__a__.S) == 0 {
		__z__ = JString(__a__)
	} else {
		__z__ = ReprStringLikeInPython(__a__.S)
	}
	return __z__
}
func (inline) JCanStr(m M) bool {
	var __z__ bool
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.CanStr()
	} else if len(__a__.S) == 0 {
		__z__ = false
	} else {
		__z__ = true
	}
	return __z__
}
func (inline) JStr(m M) string {
	var __z__ string
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.Str()
	} else if len(__a__.S) == 0 {
		panic("cannot Str() an int")
	} else {
		__z__ = __a__.S
	}
	return __z__
}
func (inline) JString(m M) string {
	var __z__ string
	__a__ := m
	if __a__.X != nil {
		__z__ = __a__.X.String()
	} else if len(__a__.S) == 0 {
		__z__ = fmt.Sprintf("%d", __a__.N)
	} else {
		__z__ = __a__.S
	}
	return __z__
}

//#if DONT
func (inline) MkX(x B) M {
	__a__ := x
	switch t := __a__.Self.(type) {
	case *PStr:
		if len(t.S) == 0 {
			return EmptyStr
		}
		return M{S: t.S}
	}
	return M{X: __a__.Self}
}

func (inline) Mkint(n int) M {
	//#if cc
	counterMkint++
	//#endif
	return M{N: int64(n)}
}

func (inline) MkInt(n int64) M {
	//#if cc
	counterMkInt++
	//#endif
	return M{N: n}
}

func (inline) MkStr(s string) M {
	__a__ := s
	//#if cc
	counterMkStr++
	//#endif
	if len(__a__) == 0 {
		return EmptyStr
	}
	return M{S: __a__}
	//z := &PStr{S: s}
	//return MForge(z)
}

func (inline) MkList(pp []M) M {
	//#if cc
	counterMkList++
	//#endif
	z := &PList{PP: pp}
	return MForge(z)
}

func (inline) MkTuple(pp []M) M {
	//#if cc
	counterMkTuple++
	//#endif
	z := &PTuple{PP: pp}
	return MForge(z)
}

func (inline) MkDict(ppp Scope) M {
	//#if cc
	counterMkDict++
	//#endif
	z := &PDict{ppp: ppp}
	return MForge(z)
}

func (inline) MkDict(ppp Scope) M {
	//#if cc
	counterMkDict++
	//#endif
	z := &PDict{ppp: ppp}
	return MForge(z)
}
//#endif DONT
