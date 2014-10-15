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

d = {'color': 'red', 'size': 'XL', 'quant': 'few'}
assert ['color', 'quant', 'size'] == sorted(d.keys())
assert ['XL', 'few', 'red'] == sorted(d.values())
assert [('color', 'red'), ('quant', 'few'), ('size', 'XL')] == sorted(d.items())

d = {'color': 'red', 'size': True, 'quant': 123}
assert ['color', 'quant', 'size'] == sorted(d.keys())
assert len(d.values()) == 3
# TODO -- when we can compare heterogenous lists.
#assert 'red' in d.values()
#assert True in d.values()
#assert 123 in d.values()
# assert [('color', 'red'), ('quant', 123), ('size', True)] == sorted(d.items())

# del from dict.
del d['color']
# del d['bogus'] # TODO -- decide if to panic on Key Error
assert ['quant', 'size'] == sorted(d.keys())
del d['size']
assert ['quant'] == sorted(d.keys())

# del from list.
l = [ 2, 4, 6, 8, 10 ]
del l[2]
assert l == [ 2, 4, 8, 10 ]

# del from list by slice.
ll = [ 2, 4, 6, 8, 10 ]
del ll[1:3]
assert ll == [ 2, 8, 10 ]

assert len(ll) == 3
my_len = len
assert my_len(ll) == 3
ap = ll.append
ap(88)
ap(99)
assert len(ll) == 5
assert my_len(ll) == 5
assert ll == [ 2, 8, 10 ] + [88] + [99]

# test dict clear()
bob = {'hair': 10, 'eyes': 20}
assert len(bob) == 2
assert bob['hair'] == 10
bob.clear()
assert len(bob) == 0
bob['hair'] = 100
assert len(bob) == 1
assert bob['hair'] == 100

# test dict copy()
bob = {'hair': 10, 'eyes': 20}
alice = bob.copy()
assert len(alice) == 2
assert alice['hair'] == 10
bob.clear()
assert len(alice) == 2
assert alice['hair'] == 10

# test dict get()
bob = {'hair': 10, 'eyes': 20}
assert bob.get('hair') == 10
assert bob.get('hairy') is None
assert bob.get('hair', 99) == 10
assert bob.get('hairy', 99) == 99
assert bob.get('hair', default=99) == 10
assert bob.get('hairy', default=99) == 99
# assert bob.get(*['hair'], **{'default': 99}) == 10  # C Python cannot do this.
# assert bob.get(*['hairy'], **{'default': 99}) == 99  # C Python cannot do this.
