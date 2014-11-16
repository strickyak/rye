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

#def AddTo(x):
#  def fn(y):
#    return x + y
#  return fn
#
#assert AddTo(3)(4) == 7
