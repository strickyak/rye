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
