Table = {}

class Atom:
  def __init__(self, x):
    self.a = 1      # Is an atom.
    self.x = x      # Value of the atom.  If type is str, it is a symbol.
    self.p = None   # Not a primative (unless you override it).
  def X(self):
    return self.x
  def P(self):
    return self.p
  def SetPrim(self, p):
    self.p = p
  def Nullp(self):
    return self.x == 'nil'
  def Atomp(self):
    return True
  def Eq(self, a):
    return self.x == a.x
  def Show(self):
    return str(self.x) + " "
  def Eval(self, env):
    if type(self.x) == str:
      return env.Find(self)
    else:
      return self

class Pair:
  def __init__(self, h, t):
    self.a = 0    # Not an atom.
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

  def Eval(self, env):
    if self.h.a:
      if self.h.p:
        return self.h.p(self, env) # Primative.
    else:
      fn = self.h.Eval(env)
      args = []
      i = self.t
      while i != Nil:
        args.append(i.h.Eval(env))
        i = i.t
      fn.Apply(args, env)

  def Apply(self, args, env):
    if self.h.a:
      raise 'Cannot apply an atom: ' + str(self.h.a)
    if self.h.h isnot Lambda:
      raise 'Cannot apply without Lambda: ' + str(self.h.h)
    formals = self.h.t.h
    expr = self.h.t.t.h
    end = self.h.t.t.t
    if end isnot Nil:
      raise 'Too much at end of Lambda:' + str(self.h.h)

    for a in args:
      env = Pair(formals.h, Pair(a, env))
      formals = formals.t
    return expr.Eval(env)

def Intern(s):
  p = Table[s]
  if p is None:
    p = Atom(s)
    Table[s] = p
    return p
  return p

def List1(a):
  return Pair(a, Nil)
def List2(a, b):
  return Pair(a, Pair(b, Nil))
def List3(a, b, c):
  return Pair(a, Pair(b, Pair(c, Nil)))
def List4(a, b, c, d):
  return Pair(a, Pair(b, Pair(c, Pair(d, Nil))))

Nil = Intern('nil')
T = Intern('t')
Lambda = Intern('lambda')
A = Intern('a')
B = Intern('b')
C = Intern('c')
D = Intern('d')

Plus = Intern('plus')
def PrimPlus(a, env):
  return Atom(a.t.h.x + a.t.t.h.x)
Plus.SetPrim(PrimPlus)

Car = Intern('car')
def PrimCar(a, env):
  return a.t.h.h
Plus.SetPrim(PrimCar)

Cdr = Intern('cdr')
def PrimCdr(a, env):
  return a.t.h.t
Plus.SetPrim(PrimCdr)

Nullp = Intern('nullp')
def PrimNullp(a, env):
  if a.t.h is Nil:
    return T
  else:
    return Nil
Plus.SetPrim(PrimNullp)

Atomp = Intern('atomp')
def PrimAtomp(a, env):
  if a.t.h.a:
    return T
  else:
    return Nil
Plus.SetPrim(PrimAtomp)

Cons = Intern('cons')
def PrimCons(a, env):
  return Pair(a.t.h, a.t.t.h)
Plus.SetPrim(PrimCons)

If = Intern('if')
def PrimIf(a, env):
  if a.t.h is Nil:
    return a.t.t.t.h.Eval(env)
  else:
    return a.t.t.h.Eval(env)
Plus.SetPrim(PrimIf)

t1 = List1(A)
print "(A) => ", t1.Show()

t2 = List2(A, B)
print "(A B) => ", t2.Show()

t3 = List3(A, B, C)
print "(A B C) => ", t3.Show()

t4 = List3(Lambda, A, Pair(B, C))
print "(Lambda A B . C) => ", t4.Show()

print "(eval 'a '(a b c d)) => ", A.Eval(List4(A, B, C, D)).Show()
print "(eval 'c '(a b c d)) => ", C.Eval(List4(A, B, C, D)).Show()
print "(plus 19 4) => ", List3(Plus, Atom(19), Atom(4)).Eval(Nil).Show()
