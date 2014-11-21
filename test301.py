import sys

print "HELLO"
print "WORLD"
print "HELLO",
print "WORLD"
print "HELLO", "",
print "WORLD"

assert 0666 == 438
assert 0xFF == 255

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
