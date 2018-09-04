def fib(n):
  native: `
    var a, ret FInt

    a.Fast.N = a_n.Self.Int()
    a.Fast.Self = & a.Fast
    a.Slow = nil

    ret.Fast.Self = & ret.Fast
    ret.Slow = nil

     FAST_FIB_Inlined(&a, &ret)

    return &ret.Fast.PBase
`

print fib(38)
