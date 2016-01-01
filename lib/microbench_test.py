
def Zero():
  pass

def One():
  s = 0
  for i in range(1000000):
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

from lib import microbench
say microbench.Run(Zero, 1)
say microbench.Run(One, 5)
say microbench.Run(Two, 5)
say microbench.Run(Three, 1)
