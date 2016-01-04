import sys

print "HELLO"
print "WORLD"
print "HELLO",
print "WORLD"
print "HELLO", "",
print "WORLD"

assert 0666 == 438
assert 0xFF == 255

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

rye_rye = False
if rye_rye:
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
