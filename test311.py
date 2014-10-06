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

d = MakeDict7(color='red')
assert d['foo'] == 7
assert d['bar'] == 77
assert d['color'] == 'red'

d = MakeDict7(8, color='red', bar=88)
assert d['foo'] == 8
assert d['bar'] == 88
assert d['color'] == 'red'

d = MakeDict7(color=None, bar=88)
assert d['foo'] == 7
assert d['bar'] == 88
assert d['color'] == None

d = MakeDict7(8, flavor='lime', **{'bar': 888, 'size': 'XL', 'level': 'duck', 'color': 'cyan'})
assert d['foo'] == 8
assert d['bar'] == 888
assert d['color'] == 'cyan'
assert d['size'] == 'XL'
assert d['flavor'] == 'lime'

d = MakeDict7(flavor='lime', *[3], **{'bar': 888, 'size': 'XL', 'level': 'duck', 'color': 'cyan'})
assert d['foo'] == 3
assert d['bar'] == 888
assert d['color'] == 'cyan'
assert d['size'] == 'XL'
assert d['flavor'] == 'lime'

d = MakeDict7(flavor='lime', *(3, 33), **{'size': 'XL', 'level': 'duck', 'color': 'cyan'})
assert d['foo'] == 3
assert d['bar'] == 33
assert d['color'] == 'cyan'
assert d['size'] == 'XL'
assert d['flavor'] == 'lime'

pass
