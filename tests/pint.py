assert type(1000) == int
assert type(0) == int
assert type(-1000) == int

assert -1000 < -2 < -1 < 0 < 1 < 2 < 1000
assert -1000 <= -2 <= -1 <= 0 <= 1 <= 2 <= 1000
assert -1000 <= -2 <= -1 <= -0 <= 0 <= 1 <= 2 <= 1000
assert -1000 <= -2 <= -1 <= -0 == 0 <= 1 <= 2 <= 1000
assert 1000 > 2 > 1 > -0 > -1 > -2 > -1000
assert 1000 >= 2 >= 1 >= -0 >= -1 >= -2 >= -1000
assert 1000 >= 2 >= 1 >= 0 >= -0 >= -1 >= -2 >= -1000
assert 1000 >= 2 >= 1 >= 0 == -0 >= -1 >= -2 >= -1000

assert 1 + 1 == 2
assert 1 + 1 == 2.0
assert 1 + 1 != 3
assert 1 + 1 != '2'
assert 1 + 1 != '2.0'
assert 1 + 1 == int('2')

#assert 1 + 1 == int(float('2.0'))

assert sum([1, 2, 3, 4]) == 10
assert sum((1, 2, 3, 4)) == 10
assert sum(range(6)) == 15
assert sum(xrange(7)) == 21

assert +100 == 100
assert 100 == +100
assert -100 == 0 - 100
assert 0 - 100 == -100

assert 20 - 15 == 5 
assert 15 - 20 == -5 
assert 1 - 1 + 2 - 2 + 3 - 3 == 0

assert 3 * 3 + 4 * 4 == 5 * 5
assert 5 * 5 - 4 * 4 == 3 * 3

assert 9 / 3 == 3
assert int(10 / 3) == 3
assert int(11 / 3) == 3
assert int(12 / 3) == 4
assert int(13 / 3) == 4

assert 9 % 3 == 0
assert 10 % 3 == 1
assert 11 % 3 == 2
assert 12 % 3 == 0
assert 13 % 3 == 1

assert not 0
assert not []
assert not ()
assert not set([])
assert not (3 * [])
assert 3 * [] == []
assert 3 * [] != ()
assert 3 * [] != None
#assert not ([] * 3)
#assert [] * 3 == []
#assert [] * 3 != ()
#assert [] * 3 != None
# assert 3 * () == ()

for i in range(30):
  assert str(i-12) == repr(i-12)
  assert len(range(i)) == i
  assert len(set(range(i))) == i
  assert len(list(range(i))) == i
  assert len(tuple(range(i))) == i
  assert len(dict([(x,x) for x in range(i)])) == i
  #assert i == int(float(i))
  #assert i == int(float(i)+0.4)
  #assert i == int(float(i) + 0.4)

assert '0123456789101112' == ''.join([str(i) for i in range(13)])
assert '0,1,2,3,4,5,6,7,8,9,10,11,12' == ','.join([str(i) for i in range(13)])

assert 255 & 255 == 255
assert 255 & 256 == 0
assert 255 & 257 == 1

assert 255 | 255 == 255
assert 255 | 256 == 511
assert 255 | 257 == 511

assert 255 ^ 255 == 0
assert 255 ^ 256 == 511
assert 255 ^ 257 == 510

assert 0 == 00
assert 1 == 01
assert 8 == 010
assert -8 == -010
assert 0 == 0x0
assert 1 == 0x1
assert 16 == 0x10
#assert -16 == -0x10

x = 1
for i in range(60):
  assert x == 1<<i
  x += x
y = x
for i in range(60):
  assert x == y
  assert y == 1<<(60-i)
  assert y == 1<<(60-i) == x
  assert y == 1<<(60-i) == x == y <= x <= y >= x >= y >= 0
  x = x/2
  y = y>>1



