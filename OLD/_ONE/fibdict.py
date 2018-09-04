def Add(a, b):
  return {'n': a['n'] + b['n']}
def Sub(a, b):
  return {'n': a['n'] - b['n']}
def LT(a, b):
  return a['n'] < b['n']
def Int(x):
  return {'n': x}

def fib(a):
  #print "fib(a) <<", a
  if LT(a, Int(2)):
    #print a, 'LESS'
    return a
  else:
    #print a, 'NOT LESS'
    x = fib(Sub(a, Int(1)))
    #print "x =", x, a
    y = fib(Sub(a, Int(2)))
    #print "y =", y, x, a
    return Add(x, y)

print fib(Int(5))
print fib(Int(6))
print fib(Int(7))
print fib(Int(8))
print fib(Int(9))
print fib(Int(30))
