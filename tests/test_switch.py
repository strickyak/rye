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
  switch:
    case x<0:
      return -8
    case x>0:
      return 8
    case x==0:
      return 0
    default:
      return None

pi = float(math.Pi)
e = float(math.E)
must -8 == g(e-pi)
must +8 == g(pi-e)
must 0 == g(e - e)
