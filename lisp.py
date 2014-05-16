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
  def show(self):
    return self.x + " "

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
  def show(self):
    if self.t.a:
      if self.t.x == 'nil':
        return "( " + self.h.show() + ") "
      else:
        return "( " + self.h.show() + ". " + self.t.show() + ") "
    else:
      return "( " + self.h.show() + self.t.showRest() + ") "
  def showRest(self):
    if self.t.a:
      if self.t.x == 'nil':
        return self.h.show()
      else:
        return self.h.show() + ". " + self.t.show()
    else:
      return self.h.show() + self.t.showRest()

def Intern(s):
  x = Table[s]
  if x is None:
    x = Atom(s)
    Table[s] = x
    return x
  return x

def List1(a):
  return Pair(a, Nil)
def List2(a, b):
  return Pair(a, Pair(b, Nil))
def List3(a, b, c):
  return Pair(a, Pair(b, Pair(c, Nil)))

Nil = Intern('nil')
Lambda = Intern('lambda')
A = Intern('a')
B = Intern('b')
C = Intern('c')

t1 = List1(A)
print "(A) => ", t1.show()

t2 = List2(A, B)
print "(A B) => ", t2.show()

t3 = List3(A, B, C)
print "(A B C) => ", t3.show()

t4 = List3(Lambda, A, Pair(B, C))
print "(Lambda A B . C) => ", t4.show()
