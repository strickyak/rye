print "HELLO"
print "WORLD"
print "HELLO" + "WORLD"

assert [1, 2, 3] == [1, 2, 3]
assert [1, 2, 3] != [1, 2, 0]
assert [1, 2, 3] != [1, 2, 3, 4]
assert [1, 2, 3, 4] != [1, 2, 3]

assert [1, 2, 3] <= [1, 2, 4]
assert [1, 2, 3] <= [1, 2, 3, 0]

assert [1, 2, 3] >= [1, 2, 2]
assert [1, 2, 3, 0] >= [1, 2, 3]

assert False == False
assert False == 0
assert False == 0.0

assert True == True
assert True == 1
assert True == 1.0

assert False < True
assert False < 1
assert False < 1.0

###########################
def TwentyThree():
  z = 0
  for i in range(4):
    if i == 2:
      return 23
    z += i
  return z

assert 23 == TwentyThree()
###########################

def Starred(a, b, *v):
  return len(v)

assert 0 == Starred(5, 5)
assert 1 == Starred(5, 5, 5)
assert 3 == Starred(5, 5, 5, 5, 5)

###########################

def Stooges(larry, moe, curly):
  return larry*100 + moe*10 + curly

assert 135 == Stooges(1, 3, 5)
assert 135 == Stooges(larry=1, moe=3, curly=5)
assert 135 == Stooges(moe=3, larry=1, curly=5)
assert 135 == Stooges(1, moe=3, curly=5)
assert 135 == Stooges(1, 3, curly=5)
assert 135 == Stooges(*[1, 3, 5])
assert 135 == Stooges(1, *[3, 5])
assert 135 == Stooges(1, 3, *[5])
assert 135 == Stooges(1, 3, 5, *[])
assert 135 == Stooges(1, 3, **{'curly': 5})

def VStooges(*v, **kw):
  larry = v[0] if len(v) > 0 else kw['larry']
  moe = v[1] if len(v) > 1 else kw['moe']
  curly = v[2] if len(v) > 2 else kw['curly']
  return larry*100 + moe*10 + curly

assert 135 == VStooges(1, 3, 5)
assert 135 == VStooges(larry=1, moe=3, curly=5)
assert 135 == VStooges(moe=3, larry=1, curly=5)
assert 135 == VStooges(1, moe=3, curly=5)
assert 135 == VStooges(1, 3, curly=5)
assert 135 == VStooges(*[1, 3, 5])
assert 135 == VStooges(1, *[3, 5])
assert 135 == VStooges(1, 3, *[5])
assert 135 == VStooges(1, 3, 5, *[])
assert 135 == VStooges(1, 3, **{'curly': 5})
pass
