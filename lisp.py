import strings

Table = {}

class Atom:
  def __init__(self, x):
    self.a = 1
    self.x = x
  def nullp(self):
    return self.x == 'nil'
  def atomp(self):
    return True
  def eq(self, a):
    return self.x == a.x

class Pair:
  def __init__(self, h, t):
    self.a = 0
    self.h = h
    self.t = t
  def nullp(self):
    return False
  def atomp(self):
    return False
  def eq(self, a):
    return self == a
  def hd(self):
    return self.h
  def tl(self):
    return self.t

def Intern(s):
  x = Table[s]
  if x is None:
    x = Atom(s)
    Table[s] = x
    return x
  return x

def List1(a):
  return Pair(a, Nil)
#def List2(a, b):
#  return Pair(a, Pair(b, Nil))
#def List3(a, b, c):
#  return Pair(a, Pair(b, Pair(c, Nil)))

Nil = Intern('nil')
Lambda = Intern('lambda')
#A = Intern('a')
#B = Intern('b')
#C = Intern('c')

#t1 = List3(A, B, C)
#print t1
#print t1.hd().x

print Lambda
#print Table
#print Pair(Lambda, Nil)
#t11 = List1(Lambda)
#print t11
