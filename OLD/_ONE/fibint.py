def fib(n):
  #print 'n <<<', n
  z = fib1(n)
  #print 'n >>>', n, z
  return z

def fib1(n):
  if n < 2:
    #print n, 'LESS'
    return n
  else:
    #print n, 'NOT LESS'
    x = fib(n-1)
    #print "x =", x, n
    y = fib(n-2)
    #print "y =", y, x, n
    return x + y

#print fib(5)
#print fib(6)
#print fib(7)
#print fib(8)
#print fib(9)
#print fib(10)
#print fib(20)
#print fib(30)
#print fib(35)
#print fib(38)
print fib(38)
