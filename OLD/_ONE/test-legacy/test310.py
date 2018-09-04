def combine(a, b, c, d=10, e=100, f=1000):
  return a+b+c+d+e+f

assert 1116 == combine(1, 2, 3)
assert 1126 == combine(1, 2, 3, 20)
assert 1326 == combine(1, 2, 3, 20, 300)
assert 9326 == combine(1, 2, 3, 20, 300, 9000)

assert 1116 == combine(1, 2, 3, d=10, e=100, f=1000)
assert 1126 == combine(1, 2, 3, 20, e=100, f=1000)
assert 1326 == combine(1, 2, 3, 20, 300, f=1000)
assert 9326 == combine(1, 2, 3, 20, 300, 9000)
