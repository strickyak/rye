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

x = [4, 6, 8]
assert len(x) == 3

x = (4, 6, 8, 9)
assert len(x) == 4

x = {'foo': 10, 'bar': 20}
assert len(x) == 2

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
