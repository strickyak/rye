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
except:
  assert len(sys.exc_info()) == 3
  msg, value, trace = sys.exc_info()

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
