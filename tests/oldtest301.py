import sys  # rye_pragma from "github.com/strickyak/rye/emulation"

print "HELLO"
print "WORLD"
print "HELLO",
print "WORLD"
print "HELLO", "",
print "WORLD"

print 'bilbo'
assert 0666 == 438
print 'frodo'
assert 0xFF == 255
print 'nando'

assert 3 < 4 < 5
assert 8 <= 8 <= 8
assert not (3 < 4 < 1)
assert not (7 < 4 < 5)
assert not (8 > 8 > 8)
assert 3 < 4 < 5 < 6 < 7 < 8

try:
  raise Exception('FooBar')
except Exception as e:
  got_ex = e
assert str(got_ex) == 'FooBar'

rye_true = False
if rye_true:
  print >>sys.stderr, 'I am rye.'
else:
  print >>sys.stderr, 'I am python.'

print __name__

class Foo:
  def __init__(self):
    pass
  def __str__(self):
    return 'nando'

# TODO: why do I need tuple for Rye?
print '%s' % ( Foo(), )

if True: print 'YES'
else: assert False

print "line46"
print "" == "foo"
print len("xyzzy")
print "line49"
for i in range(4): print i
print

class OneLiner:
  def __init__(self, x):  print 'OneLiner', x
OneLiner(42)

# Test map().
squares4 = map((lambda x: x*x), range(4))
assert squares4 == [0, 1, 4, 9]
cubes4 = map((lambda x, y: x*y), range(4), squares4)
assert cubes4 == [0, 1, 8, 27]

# Test that map() uses zip_padding_with_None() correctly, as in python.
assert map((lambda x, y, z: (x, y, z)), range(3), range(5), range(4)) == [
    (0, 0, 0), (1, 1, 1), (2, 2, 2), (None, 3, 3), (None, 4, None)]

# Test reduce().
assert 15 == reduce((lambda a, b: a+b), range(6), 0)
assert 15 == reduce((lambda a, b: a+b), range(6))
assert "OneTwoThree" == reduce((lambda a, b: a+b), ["One", "Two", "Three"], "")
assert "OneTwoThree" == reduce((lambda a, b: a+b), ["One", "Two", "Three"])

assert 15 == 1 + 2 + 3 + 4 + 5
assert 15.0 == 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert 95 != 1 + 2 + 3 + 4 + 5
assert 95.0 != 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert 95 >= 1 + 2 + 3 + 4 + 5
assert 95.0 >= 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert 95 > 1 + 2 + 3 + 4 + 5
assert 95.0 > 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert -5 != 1 + 2 + 3 + 4 + 5
assert -5.0 != 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert -5 <= 1 + 2 + 3 + 4 + 5
assert -5.0 <= 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert -5 < 1 + 2 + 3 + 4 + 5
assert -5.0 < 1.0 + 2.0 + 3.0 + 4.0 + 5.0

assert 3 == True + True + 1
assert 3 == 1 + True + True
assert 3.0 == 1.0 + True + True

assert 404 == 403 + True

assert 15 == 20 - 5
assert 15 == 20.0 - 5
assert 15 == 20 - 5.0
assert 15 == 20.0 - 5.0

assert 100 == 20 * 5
assert 100 == 20.0 * 5
assert 100 == 20 * 5.0
assert 100 == 20.0 * 5.0

assert 4 == 20 / 5
assert 4 == 20.0 / 5
assert 4 == 20 / 5.0
assert 4 == 20.0 / 5.0

assert 4 == 24 % 5
# note: float cannot mod in golang, nor in rye.

# Test short circuit evaluations.
a, b = 0, 0
def A(x):
  global a
  a += 1
  return x
def B(x):
  global b
  b += 1
  return x
assert 88 == (A(14) and B(88))
assert (a, b) == (1, 1)
assert '' == (A('') and B(88))
assert (a, b) == (2, 1)
assert 88 == (A('') or B(88))
assert (a, b) == (3, 2)
assert 77 == (A(77) or B(88))
assert (a, b) == (4, 2)
