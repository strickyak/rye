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
