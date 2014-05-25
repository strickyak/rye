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
        raise 'Cannot eval non-prim atom at head: ' + str(self.h.x) + ' ' + str(self.h.p)
    else:
      fn = self.h.Eval(env)
      args = []
      i = self.t
      print "Before While: i=", i
      while i != Nil:
        print "Weithin While: i=", i
        args.Append(i.h.Eval(env))
        i = i.t
        print "Next While: i=", i
      z = fn.Apply(args, env)
      return z

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

  def EvalArg1(self, env):
    return self.t.h.Eval(env)

  def EvalArg2(self, env):
    return self.t.h.Eval(env), self.t.t.h.Eval(env) 

  def EvalArg3(self, env):
    return self.t.h.Eval(env), self.t.t.h.Eval(env), self.t.t.t.h.Eval(env) 

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

A = Intern('a')
B = Intern('b')
C = Intern('c')
D = Intern('d')

def Truth(b):
  if b:
    return T
  else:
    return Nil

Lambda = Intern('lambda')
def SelfEvaluating(a, env):
  return a
Lambda.SetPrim(SelfEvaluating)

Plus = Intern('+')
def PrimPlus(a, env):
  b, c = a.EvalArg2(env)
  return Atom(int(b.x) + int(c.x))
Plus.SetPrim(PrimPlus)

Minus = Intern('-')
def PrimMinus(a, env):
  b, c = a.EvalArg2(env)
  return Atom(int(b.x) - int(c.x))
Minus.SetPrim(PrimMinus)

Times = Intern('*')
def PrimTimes(a, env):
  b, c = a.EvalArg2(env)
  return Atom(int(b.x) * int(c.x))
Times.SetPrim(PrimTimes)

Eq = Intern('==')
def PrimEq(a, env):
  b, c = a.EvalArg2(env)
  return Truth(int(b.x) == int(c.x))
Eq.SetPrim(PrimEq)

Lt = Intern('<')
def PrimLt(a, env):
  b, c = a.EvalArg2(env)
  return Truth(int(b.x) < int(c.x))
Lt.SetPrim(PrimLt)

Le = Intern('<=')
def PrimLe(a, env):
  b, c = a.EvalArg2(env)
  return Truth(int(b.x) <= int(c.x))
Le.SetPrim(PrimLe)

Hd = Intern('hd')
def PrimHd(a, env):
  b = a.EvalArg1(env)
  return b.h
Hd.SetPrim(PrimHd)

Tl = Intern('tl')
def PrimTl(a, env):
  b = a.EvalArg1(env)
  return b.t
Tl.SetPrim(PrimTl)

Nullp = Intern('nullp')
def PrimNullp(a, env):
  b = a.EvalArg1(env)
  return Truth(b is Nil)
Nullp.SetPrim(PrimNullp)

Atomp = Intern('atomp')
def PrimAtomp(a, env):
  b = a.EvalArg1(env)
  return Truth(b.a)
Atomp.SetPrim(PrimAtomp)

Quote = Intern('quote')
def PrimQuote(a, env):
  return a.t.h
Quote.SetPrim(PrimQuote)

Cons = Intern('cons')
def PrimCons(a, env):
  b, c = a.EvalArg2(env)
  return Pair(b, c)
Cons.SetPrim(PrimCons)

If = Intern('if')
def PrimIf(a, env):
  b = a.EvalArg1(env)
  if b is Nil:
    return a.t.t.t.h.Eval(env)
  else:
    return a.t.t.h.Eval(env)
If.SetPrim(PrimIf)

t1 = List1(A)
print "'(A) => ", t1.Show()

t2 = List2(A, B)
print "'(A B) => ", t2.Show()

t3 = List3(A, B, C)
print "'(A B C) => ", t3.Show()

t4 = List3(Lambda, A, Pair(B, C))
print "'(Lambda A B . C) => ", t4.Show()

print "(eval 'a '(a b c d)) => ", A.Eval(List4(A, B, C, D)).Show()
print "(eval 'c '(a b c d)) => ", C.Eval(List4(A, B, C, D)).Show()
print "(+ 4 19) => ", List3(Plus, Atom(4), Atom(19)).Eval(Nil).Show()
print "(- 4 19) => ", List3(Minus, Atom(4), Atom(19)).Eval(Nil).Show()
print "(* 4 19) => ", List3(Times, Atom(4), Atom(19)).Eval(Nil).Show()
print "(== 4 19) => ", List3(Eq, Atom(4), Atom(19)).Eval(Nil).Show()
print "(< 4 19) => ", List3(Lt, Atom(4), Atom(19)).Eval(Nil).Show()
print "(<= 4 19) => ", List3(Le, Atom(4), Atom(19)).Eval(Nil).Show()

x = [ 10, 20, 30 ]
x.Append(88)
print "x........ ", x

lambda10 = List3(Lambda, List2(A, B), List3(Plus, A, B))
call10 = List3(lambda10, Atom(39), Atom(3))
env10 = List4(A, Atom(111), B, Atom(222))
# print "((lambda (a b) (+ a b)) 39 3) => ", call10.Eval(env10).Show()

def SplitWhite(s):
  ww = []
  i = 0
  n = len(s)
  while i < n:
    while i < n and s[i] <= ' ':
      i += 1
    if i == n:
      break
    w = ''
    while i < n and s[i] > ' ':
      w += s[i]
      i += 1
    ww.Append(w)
  return ww

print SplitWhite("  old  man  river  ")

class LispParser:
  def __init__(self, s):
    self.s = s
    self.ww = SplitWhite(s)
    self.n = len(self.ww)
    self.i = 0

  def Next(self): 
    if self.i >= self.n:
      return None

    if self.ww[self.i] == "(":
      self.i += 1
      v = []
      while self.ww[self.i] != ")":
        x = self.Next()
        if x is None:
          raise 'Unexpected end of words, with open paren'
        v = v.Append(x)
      # Reverse the v 
      z = Nil
      for j in range(len(v)):
        z = Pair(v[0 - j - 1], z)
      return z

    z = Intern(self.ww[self.i])
    self.i += 1
    return z

print LispParser(" grep ").Next().Show()
print LispParser(" ( stdin stdout ( sigint sigkill ) ) ").Next().Show()


