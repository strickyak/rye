package rye

type FInt struct {
	// PBase  // Self points either to Fast or to Slow.
	Fast PInt
	Slow B
}

func (o *FInt) B() B {
	if o.Slow == nil {
		return &o.Fast.PBase
	}
	return o.Slow
}

type FFloat struct {
	// PBase  // Self points either to Fast or to Slow.
	Fast PFloat
	Slow B
}

func (o *FFloat) B() B {
	if o.Slow == nil {
		return &o.Fast.PBase
	}
	return o.Slow
}

type FStr struct {
	// PBase  // Self points either to Fast or to Slow.
	Fast PStr
	Slow B
}

func (o *FStr) B() B {
	if o.Slow == nil {
		return &o.Fast.PBase
	}
	return o.Slow
}

/*
var Counter_SLOW_FIB int64
func SLOW_FIB(a_n B) B {
        Counter_SLOW_FIB++
        if a_n.Self.LT(litI_c81e728d9d4c2f636f067f89cc14862c) {
                return a_n
        } else {
                var doAdd_101_left B = SLOW_FIB(( a_n.Self.Sub(litI_c4ca4238a0b923820dcc509a6f75849b)))
                var doAdd_101_right B = SLOW_FIB(( a_n.Self.Sub(litI_c81e728d9d4c2f636f067f89cc14862c)))
                var doAdd_101_sum B = ( doAdd_101_left.Self.Add(doAdd_101_right))
                return doAdd_101_sum
        }
        return None
}

func NSLOW_FIB(a_n B) B {
        if a_n.Self.LT(litI_c81e728d9d4c2f636f067f89cc14862c) {
                return a_n
        } else {
                var doAdd_101_left B = NSLOW_FIB(( a_n.Self.Sub(litI_c4ca4238a0b923820dcc509a6f75849b)))
                var doAdd_101_right B = NSLOW_FIB(( a_n.Self.Sub(litI_c81e728d9d4c2f636f067f89cc14862c)))
                var doAdd_101_sum B = ( doAdd_101_left.Self.Add(doAdd_101_right))
                return doAdd_101_sum
        }
        return None
}
*/

func FIntLT(a *FInt, b *FInt) bool {
	if a.Slow == nil && b.Slow == nil {
		return a.Fast.N < b.Fast.N
	}
	return a.Slow.Self.LT(b.Slow)
}
func FIntAdd(a *FInt, b *FInt, ret *FInt) {
	if a.Slow == nil && b.Slow == nil {
		ret.Fast.N = a.Fast.N + b.Fast.N
		ret.Slow = nil
	} else {
		ret.Slow = a.Slow.Self.Add(b.Slow)
	}
}
func FIntSub(a *FInt, b *FInt, ret *FInt) {
	if a.Slow == nil && b.Slow == nil {
		ret.Fast.N = a.Fast.N - b.Fast.N
		ret.Slow = nil
	} else {
		ret.Slow = a.Slow.Self.Sub(b.Slow)
	}
}
func FIntAssign(a *FInt, ret *FInt) {
	if a.Slow == nil {
		ret.Fast.N = a.Fast.N
		ret.Slow = nil
	} else {
		ret.Slow = a.Slow
	}
}

var Zero_FAST_FIB_VARS FAST_FIB_VARS

type FAST_FIB_VARS struct {
	v_lit_1 FInt
	v_lit_2 FInt
	v_t1    FInt
	v_t2    FInt
	v_left  FInt
	v_right FInt
	v_sum   FInt
}

func FAST_FIB(a_n *FInt, ret *FInt) {
	//println("<<<", a_n.Fast.N)
	var v FAST_FIB_VARS

	v.v_lit_1.Fast.N = 1
	v.v_lit_2.Fast.N = 2

	v.v_lit_1.Fast.Self = &v.v_lit_1.Fast
	v.v_lit_2.Fast.Self = &v.v_lit_2.Fast
	v.v_t1.Fast.Self = &v.v_t1.Fast
	v.v_t2.Fast.Self = &v.v_t2.Fast
	v.v_left.Fast.Self = &v.v_left.Fast
	v.v_right.Fast.Self = &v.v_right.Fast
	v.v_sum.Fast.Self = &v.v_sum.Fast

	//v.v_lit_1.Slow = & v.v_lit_1.Fast.PBase
	//v.v_lit_2.Slow = & v.v_lit_2.Fast.PBase
	//v.v_t1.Slow = & v.v_t1.Fast.PBase
	//v.v_t2.Slow = & v.v_t2.Fast.PBase
	//v.v_left.Slow = & v.v_left.Fast.PBase
	//v.v_right.Slow = & v.v_right.Fast.PBase
	//v.v_sum.Slow = & v.v_sum.Fast.PBase

	if FIntLT(a_n, &v.v_lit_2) {
		FIntAssign(a_n, ret)
		//println(">>>", ret.Fast.N)
		//v = Zero_FAST_FIB_VARS
		return
	} else {
		FIntSub(a_n, &v.v_lit_1, &v.v_t1)
		FAST_FIB(&v.v_t1, &v.v_left)
		FIntSub(a_n, &v.v_lit_2, &v.v_t2)
		FAST_FIB(&v.v_t2, &v.v_right)
		FIntAdd(&v.v_left, &v.v_right, &v.v_sum)

		FIntAssign(&v.v_sum, ret)
		//println(">>>>>>", ret.Fast.N)
		//v = Zero_FAST_FIB_VARS
		return

	}
	ret.Slow = None
	panic("bottom")
	//v = Zero_FAST_FIB_VARS
	return
}

func FAST_FIB_Inlined(a_n *FInt, ret *FInt) {
	//println("<<<", a_n.Fast.N)
	var v FAST_FIB_VARS

	v.v_lit_1.Fast.N = 1
	v.v_lit_2.Fast.N = 2

	//v.v_lit_1.Fast.Self = & v.v_lit_1.Fast
	//v.v_lit_2.Fast.Self = & v.v_lit_2.Fast
	//v.v_t1.Fast.Self = & v.v_t1.Fast
	//v.v_t2.Fast.Self = & v.v_t2.Fast
	//v.v_left.Fast.Self = & v.v_left.Fast
	//v.v_right.Fast.Self = & v.v_right.Fast
	//v.v_sum.Fast.Self = & v.v_sum.Fast

	var cond bool
	if a_n.Slow == nil && v.v_lit_2.Slow == nil {
		cond = (a_n.Fast.N < v.v_lit_2.Fast.N)
	} else {
		cond = a_n.Slow.Self.LT(v.v_lit_2.Slow)
	}

	// if FIntLT(a_n, &v.v_lit_2) //
	if cond {
		// FIntAssign(a_n, ret)
		ret.Fast.N = a_n.Fast.N
		ret.Slow = nil
		//println(">>>", ret.Fast.N)
		return
	} else {

		// FIntSub(a_n, &v.v_lit_1, &v.v_t1)
		if a_n.Slow == nil && v.v_lit_1.Slow == nil {
			v.v_t1.Fast.N = a_n.Fast.N - v.v_lit_1.Fast.N
		} else {
			v.v_t1.Slow = a_n.Slow.Self.Sub(v.v_lit_1.Slow)
		}

		FAST_FIB_Inlined(&v.v_t1, &v.v_left)

		// FIntSub(a_n, &v.v_lit_2, &v.v_t2)
		if a_n.Slow == nil && v.v_lit_2.Slow == nil {
			v.v_t2.Fast.N = a_n.Fast.N - v.v_lit_2.Fast.N
		} else {
			v.v_t2.Slow = a_n.Slow.Self.Sub(v.v_lit_2.Slow)
		}

		FAST_FIB_Inlined(&v.v_t2, &v.v_right)

		// FIntAdd(&v.v_left, &v.v_right, &v.v_sum)
		if v.v_left.Slow == nil && v.v_right.Slow == nil {
			v.v_sum.Fast.N = v.v_left.Fast.N + v.v_right.Fast.N
		} else {
			v.v_sum.Slow = v.v_left.Slow.Self.Add(v.v_right.Slow)
		}

		// FIntAssign(&v.v_sum, ret)
		ret.Fast.N = v.v_sum.Fast.N
		ret.Slow = nil
		//println(">>>>>>", ret.Fast.N)
		return

	}
	ret.Slow = None
	panic("bottom")
	return
}

// fib(38):
// $ time python fibint.py // 39088169 // real  0m13.442s
// $ GOMAXPROCS=1 time ./fibint.bin // 39088169 // 16.73user 0.07system 0:16.78elapsed
// $ GOMAXPROCS=1 time ./fibfast.bin // 39088169 // 19.32user 0.15system 0:19.23elapsed

/*
$  time ./fibfast_inlined.bin
39088169
real  0m2.627s user  0m2.627s sys 0m0.004s

$ GOMAXPROCS=1  time ./fibfast_inlined.bin
39088169
2.61user 0.00system 0:02.61elapsed 100%CPU (0avgtext+0avgdata 2500maxresident)k
0inputs+0outputs (0major+659minor)pagefaults 0swaps

$ GOMAXPROCS=1 RYE_PPROF=_pprof time ./fibfast_inlined.bin
39088169
2.13user 0.01system 0:02.14elapsed 100%CPU (0avgtext+0avgdata 4644maxresident)k

*/
