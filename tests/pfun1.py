def Lsd(l, s, d):
  #return 240*l+12*s+d
  return 240*l + 12*s + d

def Lsd0(l=0, s=0, d=0):
  return 240*l + 12*s + d

C = 0
def LsdC(l=C, s=C, d=C):
  return 240*l + 12*s + d

assert 267 == Lsd(1,2,3)
assert 267 == Lsd(1,2,3.0)
assert 267.0 == Lsd(1,2,3)
assert 267.0 == Lsd(1,2,3.0)

assert 0 == Lsd0()
assert 240 == Lsd0(1)
assert 264 == Lsd0(1,2)
assert 267 == Lsd0(1,2,3)

assert 0 == LsdC()
assert 240 == LsdC(1)
assert 264 == LsdC(1,2)
assert 267 == LsdC(1,2,3)

# Bad
#C = 10
#print LsdC()
#print LsdC(1)
#print LsdC(1,2)
#print LsdC(1,2,3)
#assert 0 == LsdC()
#assert 240 == LsdC(1)
#assert 264 == LsdC(1,2)
#assert 267 == LsdC(1,2,3)
