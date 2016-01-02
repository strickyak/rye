def Zero():
  return 499999500000

def One():
  s = 0
  for i in range(1000000):
    s += i
  return s

def OneA():
  s = 0
  for i in range(0 + 1000000):
    s += i
  return s

def OneB():
  s = 0
  r = range(0 + 1000000)
  for i in r:
    s += i
  return s

def XOne():
  s = 0
  for i in xrange(1000000):
    s += i
  return s

def XOneA():
  s = 0
  for i in xrange(0 + 1000000):
    s += i
  return s

def XOneB():
  s = 0
  r = xrange(0 + 1000000)
  for i in r:
    s += i
  return s

def XXOne():
  def xxrange(n):
    for i in xrange(n):
      yield i
  s = 0
  for i in xxrange(1000000):
    s += i
  return s

def Two():
  s = 0
  i = 0
  while i < 1000000:
    s += i
    i += 1
  return s

def Three():
  native: `
    s := 0
    i := 0
    for i < 1000000 {
      s += i
      i += 1
    }
    return Mkint(s)
  `

must Zero() == 499999500000
must One() == 499999500000
must OneA() == 499999500000
must OneB() == 499999500000
must XOne() == 499999500000
must XOneA() == 499999500000
must XOneB() == 499999500000
must XXOne() == 499999500000
must Two() == 499999500000
must Three() == 499999500000

from lib import microbench
say microbench.Run(Zero, 3)
say microbench.Run(One, 10)
say microbench.Run(OneA, 10)
say microbench.Run(OneB, 10)
say microbench.Run(XOne, 10)
say microbench.Run(XOneA, 10)
say microbench.Run(XOneB, 10)
say microbench.Run(XXOne, 10)
say microbench.Run(Two, 10)
say microbench.Run(Three, 3)
