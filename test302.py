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

aa, (ab, cd), (ef, gh) = 5, (10, 20), (30, (40, 44))
assert aa == 5
assert ab == 10
assert cd == 20
assert ef == 30
assert gh == (40, 44)

(xx, yy) = (7, 8)
assert xx == 7
assert yy == 8

def AFactor(x):
  if x % 2 == 0:
    return 2
  elif x % 3 == 0:
    return 3
  elif x % 5 == 0:
    return 5
  else:
    return 0

assert [AFactor(x+2) for x in range(8)] == [2, 3, 2, 5, 2, 0, 2, 3]
