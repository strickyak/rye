Table = {}

class Atom:
  def __init__(self, x):
    self.a = 1
    self.x = x
  def Nullp(self):
    return self.x == 'nil'
  def Atomp(self):
    return True
  def Eq(self, a):
    return self.x == a.x
  def Show(self):
    return self.x + " "
  def Eval(self, env):
    if type(self.x) == str:
      return env.Find(self)
    else:
      return self

class Pair:
  def __init__(self, h, t):
    self.a = 0
    self.h = h
    self.t = t
  def Nullp(self):
    return False
  def Atomp(self):
    return False
  def Eq(self, a):
    return self == a
  def Hd(self):
    return self.h
  def Tl(self):
    return self.t
  def Show(self):
    if self.t.a:
      if self.t.x == 'nil':
        return "( " + self.h.Show() + ") "
      else:
        return "( " + self.h.Show() + ". " + self.t.Show() + ") "
    else:
      return "( " + self.h.Show() + self.t.ShowRest() + ") "
  def ShowRest(self):
    if self.t.a:
      if self.t.x == 'nil':
        return self.h.Show()
      else:
        return self.h.Show() + ". " + self.t.Show()
    else:
      return self.h.Show() + self.t.ShowRest()
  def Find(self, a):
    if self.h == a:
      return self.t.h
    else:
      return self.t.t.Find(a)

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
def List4(a, b, c, d):
  return Pair(a, Pair(b, Pair(c, Pair(d, Nil))))

Nil = Intern('nil')
Lambda = Intern('lambda')
A = Intern('a')
B = Intern('b')
C = Intern('c')
D = Intern('d')

t1 = List1(A)
print "(A) => ", t1.Show()

t2 = List2(A, B)
print "(A B) => ", t2.Show()

t3 = List3(A, B, C)
print "(A B C) => ", t3.Show()

t4 = List3(Lambda, A, Pair(B, C))
print "(Lambda A B . C) => ", t4.Show()

print "(eval 'a '(a b c d)) => ",A.Eval(List4(A, B, C, D)).Show()
print "(eval 'c '(a b c d)) => ",C.Eval(List4(A, B, C, D)).Show()
