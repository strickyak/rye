#go import os
go import math/big
go import crypto/md5
go import encoding/base64

a = big.NewInt(100)
print a
print a.String()
b = big.NewInt(100)
c = big.NewInt(0)

c.Exp(a, b, None)
print c.String()

s = md5.Sum('abcdefg')
xx = 0
for x in s:
  xx += x
print xx
print type(s)
print repr(str(s))
print repr(str(byt(repr(str(s)))))

q = byt(' apple banana coconut durian ')
print 'q:', repr(q)
r = base64.StdEncoding.EncodeToString(q)
print 'r:', repr(r)
q2 = base64.StdEncoding.DecodeString(r)
print 'q2:', repr(q2)
assert q == q2
