# test generators

def CountBy(n):
  x = 0
  while True:
    yield x
    x += n

for x in CountBy(7):
  if x > 100:
    break
  print x

for x in CountBy(55):
  if x > 100:
    break
  print x

for x in CountBy(999):
  if x > 100:
    break
  print x

# Test nested function.

def AddTo(x):
  def augment(y):
    return x + y
  z = x*x
  return augment

assert AddTo(3)(4) == 7

def EvenK(n):
  if n & 1:
    def F():
      return 'odd'
  else:
    def F():
      return 'even'
  return F

assert EvenK(6)() == 'even'
assert EvenK(7)() == 'odd'

# The mistake programmers often make.
def GenBad(n):
  for i in range(n):
    def augment(x):
      return x + i
    yield augment
  i = n

assert sum([f(100) for f in list(GenBad(5))]) == 25 + 500

# Corrected version.
def GenGood(n):
  for i in range(n):
    def make_augment():
      copy_i = i
      def augment(x):
        return x + copy_i
      return augment
    yield make_augment()
  i = n

assert sum([f(100) for f in list(GenGood(5))]) == 10 + 500

# Corrected version, with a class.

class GoodClass:
  def __init__(self, n):
    self.n = n

  def Generate(self):
    for i in range(self.n):
      def make_Augment():
        copy_i = i
        def Augment(x):
          return x + copy_i
        return Augment
      yield make_Augment()
    i = self.n

assert sum([f(100) for f in list(GoodClass(5).Generate())]) == 10 + 500

# This is nondeterministic in Rye, but not in Python.
# Print it if you are curious about GenBad:
# print [f(0) for f in GenBad(100)]

fn = lambda x, y : x + y
assert fn(20, 22) == 42

fn = lambda x : x + x
assert fn(21) == 42

def AddToL(x):
  return lambda y: x + y
assert AddToL(100)(88) == 188

def ConstantL(x):
  return lambda : x
assert ConstantL(800)() == 800

class ConstantCls:
  def __init__(self, x):
    self.x = x
    self.lam = lambda : x
  def Get(self):
    return self.lam
  def Get2(self):
    return lambda: self.x
assert ConstantCls(800).Get()() == 800
assert ConstantCls(800).Get2()() == 800
assert ConstantCls(800).lam() == 800
