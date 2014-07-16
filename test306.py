def A(x):
  return x + x + x + x
def B(x):
  return A(x) + A(x) + A(x) + A(x)
def C(x):
  return B(x) + B(x) + B(x) + B(x)
def D(x):
  return C(x) + C(x) + C(x) + C(x)
def E(x):
  return D(x) + D(x) + D(x) + D(x)
def F(x):
  return E(x) + E(x) + E(x) + E(x)
def G(x):
  return F(x) + F(x) + F(x) + F(x)
def H(x):
  return G(x) + G(x) + G(x) + G(x)
def I(x):
  return H(x) + H(x) + H(x) + H(x)
def J(x):
  return I(x) + I(x) + I(x) + I(x)
def K(x):
  return J(x) + J(x) + J(x) + J(x)

print K(1)

def Abs(x):
  return x if x >= 0 else 0-x
def Sgn(x):
  return -1 if x < 0 else 0 if x == 0 else 1

assert Abs(10) == 10
assert Abs(0) == 0
assert Abs(-8) == 8
assert Sgn(10) == 1
assert Sgn(0) == 0
assert Sgn(-8) == -1

assert 'T' == ('T' if True else 'F')
assert 'F' == ('T' if False else 'F')

d = {}
d['a'], d['b'] = 50, 60
assert d['a'] == 50
assert d['b'] == 60
