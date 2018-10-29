from go import math, os

def f(x):
  switch x:
    case 1, 3, 5:
      return 10*x
    case 2, 4, 6:
      return 100*x
    default:
      return 0

p = os.Getpid()
i = 4 * p / p
must 400 == f(i)
must 50 == f(i+1)

def g(x):
  say x
  switch:
    case x is None, x==0:
      return 'zero'
    case x<-1, x>1:
      return 'big'
    case x<0, x>0:
      return 'small'
    default:
      return None

pi = float(math.Pi)
e = float(math.E)
must 'small' == g(e-pi)
must 'small' == g(pi-e)
must 'big' == g(pi+e)
must 'big' == g(-pi-e)
must 'zero' == g(e - e)
must 'zero' == g(None)
