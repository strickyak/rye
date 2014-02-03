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
