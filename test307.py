def TriangleWithRecursion(n):
  if n < 1:
    return 0
  return n + TriangleWithRecursion(n-1)

print TriangleWithRecursion(6)

def TriangleWithWhile(n):
  z = 0
  while n > 0:
    z = z + n
    n = n - 1
  return z

print TriangleWithWhile(6)

def TriangleWithRecursionStrings(s):
  if len(s) < 1:
    return ""
  z = s + TriangleWithRecursionStrings(s[:-1])
  print 'RETURNING'
  print z
  return z

print TriangleWithRecursionStrings("abcdef")

# Test tuples, lists, dicts.

x = [8, 4, 6]
assert len(x) == 3
assert 6 in x
assert 7 not in x
s = 0
for x0 in x:
  s = s + x0
assert s == 18, s
sx = sorted(x)
assert len(x) == 3
assert x[0] == 8
assert x[1] == 4
assert x[2] == 6
assert len(sx) == 3
assert sx[0] == 4
assert sx[1] == 6
assert sx[2] == 8

x = (9, 6, 8, 4)
assert len(x) == 4
assert 6 in x
assert 7 not in x
s = 0
for x0 in x:
  s = s + x0
assert s == 27, s
sx = sorted(x)
assert len(x) == 4
assert x[0] == 9
assert x[1] == 6
assert x[2] == 8
assert len(sx) == 4
assert sx[0] == 4
assert sx[1] == 6
assert sx[2] == 8


x = {'foo': 10, 'bar': 20}
assert len(x) == 2
assert 'foo' in x
assert 'mumble' not in x
s = 0
for x0 in x:
  s = s + len(x0)
assert s == 6, s
sx = sorted(list(x))
assert len(sx) == 2
assert sx[0] == 'bar'
assert sx[1] == 'foo'

assert type('foo') == str
assert type(4) == int
assert type(float(3)) == float
assert type([2,3,4]) == list
assert type({'xyz':17}) == dict

def Loop():
  z = 0
  v = [4, 6, 8]
  for x in v:
    print x
    z = z + x
  assert z == 18

Loop()

if True:
  print "YES"
if False:
  print "NO"
if None:
  print "NONE"

s = 0
for i in range(10):
  for j in range(10):
    n = i * 10 + j + 1
    s = s + n
assert s == 5050
print s, s, s, 'Woot!'

for (x, y) in [(1, 10), (2, 20), (3, 30)]:
  s *= x
  s += y
print s

for ((x0, x1), (y0, y1)) in [((1, 10), (2, 20)), ((3, 30), (4, 40))]:
  s *= (x0 + x1)
  s += (y0 + 3*y1)
print s

s = 1001
for x, y in [(1, 10), (2, 20), (3, 30)]:
  s *= x
  s += y
print s

for (x0, x1), (y0, y1) in [((1, 10), (2, 20)), ((3, 30), (4, 40))]:
  s *= (x0 + x1)
  s += (y0 + 3*y1)
print s

s = 2001
for (x0, _), _, (_, y1) in [((1, 10), 7, (2, 20)), ((3, 30), 7, (4, 40))]:
  s *= x0
  s += y1
print s
