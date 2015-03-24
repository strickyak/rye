z = []
for x in range(10):
  z.append(x*x)
must sum(z) == 285
say z

z = 0
for x in range(12):
  if x == 10:
    break
  z += x*x
must z == 285
z = 0
for x in range(12):
  if x >= 10:
    continue
  z += x*x
must z == 285

z = []
for e in range(10):
  if e % 2 == 0:
    z = z + [e]
  else:
    z = z + [None]
must z == [0, None, 2, None, 4, None, 6, None, 8, None]
say z

i = 9
z = 0
while i > 0:
  z += i*i
  i -= 1
must z == 285

while False:
  must 1 == 2

z = 1
try:
  z = 2
except as ex:
  z = 3
must z == 2
try:
  z = 'foo' / 'bar'
except as ex:
  z = 3
  say ex
must z == 3
say ex 
say str(ex)
must str(ex).find('Receiver cannot Div:') >= 0 or str(ex).find('TypeError:')

def pi():
  return 3.14
must pi() == 3.14

def twice(x):
  return x + x
must twice(100) == 200
must twice('xyz') == 'xyzxyz'
must twice(range(3)) == [0, 1, 2, 0, 1, 2]

say "OKAY test201.py"
