def Foo(a, b):
  i1 = a ##i
  i2 = b ##i
  return i1 + i2

x = 120 ##i
y = 3   ##i
print Foo(x, y)

def MakeDict(**kw):
  return kw

d = MakeDict(abc=123, xyz=789)
assert d['abc'] == 123
assert d['xyz'] == 789

def MakeDict7(foo=7, bar=77, **kw):
  return MakeDict(foo=foo, bar=bar, **kw)

d2 = MakeDict7(color='red')
print 'd2 = ', repr(d2)
assert d2['foo'] == 7
assert d2['bar'] == 77
assert d2['color'] == 'red'

d3 = MakeDict7(8, color='red', bar=88)
print 'd3 = ', repr(d3)
assert d3['foo'] == 8
assert d3['bar'] == 88
assert d3['color'] == 'red'
