rye_true = False
if rye_true:
  from . import pfun1 as F
else:
  import pfun1 as F

def Lsd4(l, s, d):
  return 1000 * F.Lsd(l, s, d)

def Lsd2(l=0, s=0, d=0):
  return 1000000 * F.Lsd0(l, s, d)

assert 267 == F.Lsd(1,2,3)
assert 267 == F.Lsd(1,2,3.0)
assert 267.0 == F.Lsd(1,2,3)
assert 267.0 == F.Lsd(1,2,3.0)

assert 0 == F.Lsd0()
assert 240 == F.Lsd0(1)
assert 264 == F.Lsd0(1,2)
assert 267 == F.Lsd0(1,2,3)

assert 267000 == Lsd4(1,2,3)
assert 267000 == Lsd4(1,2,3.0)
assert 267000.0 == Lsd4(1,2,3)
assert 267000.0 == Lsd4(1,2,3.0)

assert 0 == Lsd2()
assert 240000000 == Lsd2(1)
assert 264000000 == Lsd2(1,2)
assert 267000000 == Lsd2(1,2,3)
