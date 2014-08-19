from go import strings as foo
from go import strconv
from go import fmt as F

import github.com/strickyak/rye/twice as Doppel

print foo.Contains("Hello", "ell"), Doppel.Twice(21)

print strconv.ParseBool("true")
print strconv.Atoi("1234")

try:
  print strconv.Atoi("23skidoo")
except:
  print 'Expected.'

print strconv.UnquoteChar("\\'foo", 39)
print F.Sprintf("%d %d %d", 111, 222, 333)

plus300 = Doppel.Adder(300)
print plus300.Plus(69)

class Augmentor(Doppel.Adder):
  def __init__(self, some, more):
    super(some + more)

aug123 = Augmentor(100, 23)
print aug123.Plus(4000)
print aug123

class Augmentarian(Augmentor):
  def __init__(self, some, even, more):
    super(some + more, even)

aug246 = Augmentarian(200, 40, 6)
print aug246.Plus(8000)
print aug246

class DoubleAugmentarian(Augmentor):
  def __init__(self, some, even, more):
    super(some + more, even)
  def Plus(self, x):
    return 2 * super.Plus(x)

da246 = DoubleAugmentarian(200, 40, 6)
print da246.Plus(8000)
print da246

assert (-1 >> 3) < 0
assert (-1 >> 3) == -1
assert (-1 >>> 3) > 0
assert (-1 >>> 3) == 2305843009213693951

def Lion(x):
  global cat
  defer cat += x
  assert cat == 100
  cat += x
  assert cat == 110

def Tiger():
  global cat
  cat += 100
  Lion(10)
  assert cat == 120

cat = 0
Tiger()

def TwentyThree():
  z = 0
  for i in range(4):
    if i == 2:
      return 23
    z += i
  return z

assert 23 == TwentyThree()

def Stooges(larry, moe, curly):
  return larry*100 + moe*10 + curly

assert 135 == Stooges(1, 3, 5)
#assert 135 == Stooges(larry=1, moe=3, curly=5)

def Starred(a, b, *v):
  return len(v)

assert 0 == Starred(5, 5)
assert 1 == Starred(5, 5, 5)
assert 3 == Starred(5, 5, 5, 5, 5)


