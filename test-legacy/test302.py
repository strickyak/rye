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
ex = ll.extend
ex([42, 43, 44])
assert ll == [2, 8, 10, 88, 99, 42, 43, 44]

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
# assert bob.get('hair', default=99) == 10
# assert bob.get('hairy', default=99) == 99
# assert bob.get(*['hair'], **{'default': 99}) == 10  # C Python cannot do this.
# assert bob.get(*['hairy'], **{'default': 99}) == 99  # C Python cannot do this.

assert 'one,two,three'.split(',') == ['one', 'two', 'three']
assert 'one,two,three'.split(',', 1) == ['one', 'two,three']
assert 'one,two,three'.split(',', -1) == ['one', 'two', 'three']
assert 'one,two,three'.split(';') == ['one,two,three']
assert ''.split(';') == ['']

assert 'Once UPON a time'.lower() == 'once upon a time'
# broken in Go? # assert 'Once UPON a time'.title() == 'Once Upon A Time'
assert 'Once UPON a time'.upper() == 'ONCE UPON A TIME'

assert ' ; Once upon a time; \n'.strip(' \t,;\n') == 'Once upon a time'
assert ' ; Once upon a time; \n'.lstrip(' \t,;\n') == 'Once upon a time; \n'
assert ' ; Once upon a time; \n'.rstrip(' \t,;\n') == ' ; Once upon a time'

assert ' x x x x x '.replace('x', 'y', 3) == ' y y y x x '
assert ' x x x x x '.replace('x', 'y', -1) == ' y y y y y '
assert ' x x x x x '.replace('x', 'y') == ' y y y y y '

stuff = [[4, 5, 6], 11, 22, 33, 44, 'five', 'six', 'seven', 22, 44, 55]
assert stuff.count(22) == 2
assert stuff.count('five') == 1
assert stuff.count('nine') == 0
assert stuff.index(22) == 2
assert stuff.index('five') == 5

stuff.remove(44)
assert stuff == [[4, 5, 6], 11, 22, 33, 'five', 'six', 'seven', 22, 44, 55]

stuff.remove(44)
assert stuff == [[4, 5, 6], 11, 22, 33, 'five', 'six', 'seven', 22, 55]

stuff.insert(3, 25)
assert stuff == [[4, 5, 6], 11, 22, 25, 33, 'five', 'six', 'seven', 22, 55]

stuff.insert(0, None)
assert stuff == [None, [4, 5, 6], 11, 22, 25, 33, 'five', 'six', 'seven', 22, 55]

z = stuff.pop()
assert stuff == [None, [4, 5, 6], 11, 22, 25, 33, 'five', 'six', 'seven', 22]
assert z == 55

z = stuff.pop(2)
assert stuff == [None, [4, 5, 6], 22, 25, 33, 'five', 'six', 'seven', 22]
assert z == 11

stuff.reverse()
assert stuff == [22, 'seven', 'six', 'five', 33, 25, 22, [4, 5, 6], None]

z = []
z.reverse()
assert z == []

z = [5]
z.reverse()
assert z == [5]

z = [4, 7]
z.reverse()
assert z == [7, 4]

z = [3, 6, 9]
z.reverse()
assert z == [9, 6, 3]

d = {'one': 1, 'two': 2, 'three': 3}
assert sorted(d.keys()) == sorted(['one', 'two', 'three'])
assert sorted(d.values()) == sorted([1, 2, 3])
assert sorted(d.items()) == sorted([('one', 1), ('two', 2), ('three', 3)])
assert sorted(d.iterkeys()) == sorted(['one', 'two', 'three'])
assert sorted(d.itervalues()) == sorted([1, 2, 3])
assert sorted(d.iteritems()) == sorted([('one', 1), ('two', 2), ('three', 3)])
assert d.has_key('one')
assert not d.has_key('zero')
assert 'one' in d
assert 'zero' not in d
assert d.setdefault('one') == 1
assert d.setdefault('one', 'foo') == 1
assert d.setdefault('zero') is None
assert d.setdefault('nine', 'foo') == 'foo'
assert 'zero' in d
assert 'nine' in d
assert d['one'] == 1
assert d['zero'] is None
assert d['nine'] == 'foo'

d.update({'five': 5, 'six': 6})
assert sorted(d.iteritems()) == sorted([('one', 1), ('two', 2), ('three', 3), ('five', 5), ('six', 6), ('zero', None), ('nine', 'foo')])

assert len(dict()) == 0
assert len(dict(color='violet')) == 1
assert len(dict(color='violet', format='text')) == 2
assert dict(color='violet', format='text')['color'] == 'violet'
assert dict({'color': 'blue'}, color='violet', format='text')['color'] == 'violet'
assert dict([('color', 'blue'), ('format', 'text')])['color'] == 'blue'

assert max([3, 5, 2, 9, 6]) == 9
assert max(3, 5, 2, 9, 6) == 9
assert min([3, 5, 2, 9, 6]) == 2
assert min(3, 5, 2, 9, 6) == 2
assert max('nando') == 'o'
assert min('nando') == 'a'
assert max(['nando']) == 'nando'
assert min(['nando']) == 'nando'
assert max('nando', 'nando') == 'nando'
assert min('nando', 'nando') == 'nando'

assert range(5) == [0, 1, 2, 3, 4]
assert list(xrange(5)) == [0, 1, 2, 3, 4]

def cmpLenStr(a, b):
  if len(str(a)) < len(str(b)):
    return -1
  else:
    return 1

assert sorted([10, 1, 2, 100]) == [1, 2, 10, 100]
assert sorted([10, 1, 2, 100], reverse=True) == [100, 10, 2, 1]
assert sorted([10, 1, 2, 100], key=str) == [1, 10, 100, 2]
assert sorted([-1, 5, 2222, 300], cmp=cmpLenStr) == [5, -1, 300, 2222]

vec = [10, 1, 2, 100]
vec.sort()
assert vec == [1, 2, 10, 100]
vec.sort(reverse=True)
assert vec == [100, 10, 2, 1]

s = """Multi""
'Line
"String"""

assert s.split('\n') == ['Multi""', "'Line", '"String']

h = (3
  + 4) * (
    (3
    )*
  2
  )
assert h == 42

assert not ''.isupper()
assert 'MONDAY'.isupper()
assert not 'MONDAY'.islower()
assert 'tuesday'.islower()
assert not 'tuesday'.isupper()
assert not 'tuesday'.isdigit()
assert not 'tuesday'.isspace()
assert 'Wednesday'.isalpha()
assert 'Wednesday'.isalnum()
assert 'Wednesday404'.isalnum()
assert not 'Wednesday404'.isalpha()
assert not 'Wednesday404'.isdigit()
assert '404'.isalnum()
assert '404'.isdigit()
assert ' \t\r\n'.isspace()
assert not ''.isspace()

assert ':a:bx::c\t\n:'.split(':') == ['', 'a', 'bx', '', 'c\t\n', '']
assert ' a bx  c\t\n '.split(' ') == ['', 'a', 'bx', '', 'c\t\n', '']
assert ' a bx  c\t\n '.split() == ['a', 'bx', 'c' ]
assert '()'.split(':') == ['()']
assert '()'.split() == ['()']
assert ':'.split(':') == ['', '']
assert ' '.split() == []
assert '\r'.split(':') == ['\r']
assert '\r'.split() == []
assert ''.split(':') == ['']
assert ''.split() == []

print "302 OKAY."
