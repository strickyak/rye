z = []
for x in range(10):
  z.append(x*x)
assert sum(z) == 285
print z

z = 0
for x in range(12):
  if x == 10:
    break
  z += x*x
assert z == 285
z = 0
for x in range(12):
  if x >= 10:
    continue
  z += x*x
assert z == 285

z = []
for e in range(10):
  if e % 2 == 0:
    z = z + [e]
  else:
    z = z + [None]
assert z == [0, None, 2, None, 4, None, 6, None, 8, None]
print z

i = 9
z = 0
while i > 0:
  z += i*i
  i -= 1
assert z == 285

while False:
  assert 1 == 2

z = 1
try:
  z = 2
except as ex:
  z = 3
assert z == 2
try:
  z = 'foo' / 'bar'
except as ex:
  z = 3
  print ex
assert z == 3
print ex 
print str(ex)
assert str(ex).find('Receiver cannot Div:') >= 0 or str(ex).find('TypeError:')

def pi():
  return 3.14
assert pi() == 3.14

def twice(x):
  return x + x
assert twice(100) == 200
assert twice('xyz') == 'xyzxyz'
assert twice(range(3)) == [0, 1, 2, 0, 1, 2]

def forward_triangle(n):
  return triangle(n)

def triangle(n):
  if n < 1:
    return 0
  else:
    return n + triangle(n - 1)
assert forward_triangle(5) == 15
assert triangle(6) == 21

vec = [6]
assert triangle(*vec) == 21

print "OKAY test201.py"
