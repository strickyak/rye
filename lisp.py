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
      if self.x in ['nil', 'true']:
        return self
      else:
        return env.Find(self)
    else:
      return self
  def Apply(self, args, env):
    return self.p(args, env)
  def Find(self, a):
    if self == Nil:
      raise "cannot find: " + a.Show()
    else:
      raise "Sending 'find' to Atom: " + self.Show()

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
    prim = None
    fn = None
    if self.h.a:
      if self.h.x in SPECIAL_FORMS:
        z = SPECIAL_FORMS[self.h.x](self, env)
        return z
      if self.h.p:
        prim = self.h.p
    if not prim:
      fn = self.h.Eval(env)
    args = []
    i = self.t
    while i is not Nil:
      args.append(i.h.Eval(env))
      i = i.t
    if prim:
      z = prim(args, env)
    else:
      z = fn.Apply(args, env)
    return z

  def Apply(self, args, env):
    if self.h is not Lambda:
      raise 'Cannot apply without Lambda: ' + str(self.h.h)
    formals = self.t.h
    expr = self.t.t.h
    end = self.t.t.t
    if end is not Nil:
      raise 'Too much at end of Lambda:' + str(self)

    for a in args:
      env = Pair(formals.h, Pair(a, env))
      formals = formals.t
    z = expr.Eval(env)
    return z

def Intern(s):
  p = Table.get(s)
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
True_ = Intern('true')
Defun = Intern('defun')

A = Intern('a')
B = Intern('b')
C = Intern('c')
D = Intern('d')

def Truth(b):
  if b:
    return True_
  else:
    return Nil

Plus = Intern('+')
def PrimPlus(args, env):
  b, c = args
  return Atom(int(b.x) + int(c.x))
Plus.SetPrim(PrimPlus)

Minus = Intern('-')
def PrimMinus(args, env):
  b, c = args
  return Atom(int(b.x) - int(c.x))
Minus.SetPrim(PrimMinus)

Times = Intern('*')
def PrimTimes(args, env):
  b, c = args
  return Atom(int(b.x) * int(c.x))
Times.SetPrim(PrimTimes)

Eq = Intern('==')
def PrimEq(args, env):
  b, c = args
  return Truth(int(b.x) == int(c.x))
Eq.SetPrim(PrimEq)

Lt = Intern('<')
def PrimLt(args, env):
  b, c = args
  return Truth(int(b.x) < int(c.x))
Lt.SetPrim(PrimLt)

Le = Intern('<=')
def PrimLe(args, env):
  b, c = args
  return Truth(int(b.x) <= int(c.x))
Le.SetPrim(PrimLe)

Hd = Intern('hd')
def PrimHd(args, env):
  b, = args
  return b.h
Hd.SetPrim(PrimHd)

Tl = Intern('tl')
def PrimTl(args, env):
  b, = args
  return b.t
Tl.SetPrim(PrimTl)

Nullp = Intern('nullp')
def PrimNullp(args, env):
  b, = args
  return Truth(b is Nil)
Nullp.SetPrim(PrimNullp)

Atomp = Intern('atomp')
def PrimAtomp(args, env):
  b, = args
  return Truth(b.a)
Atomp.SetPrim(PrimAtomp)

Cons = Intern('cons')
def PrimCons(args, env):
  b, c = args
  return Pair(b, c)
Cons.SetPrim(PrimCons)

def SpecialIf(a, env):
  b = a.t.h.Eval(env)
  if b is Nil:
    return a.t.t.t.h.Eval(env)
  else:
    return a.t.t.h.Eval(env)

def SpecialQuote(a, env):
  return a.t.h

def SpecialLambda(a, env):
  return a

If = Intern('if')
Quote = Intern('quote')
Lambda = Intern('lambda')
SPECIAL_FORMS = {
    "if": SpecialIf,
    "quote": SpecialQuote,
    "lambda": SpecialLambda,
}

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
x.append(88)
print "x........ ", x

lambda10 = List3(Lambda, List2(A, B), List3(Plus, A, B))
call10 = List3(lambda10, Atom(39), Atom(3))
env10 = List4(A, Atom(111), B, Atom(222))

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
    ww.append(w)
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
        v.append(x)
      # Reverse the v 
      self.i += 1
      z = Nil
      for j in range(len(v)):
        z = Pair(v[0 - j - 1], z)
      return z

    else:
      w = self.ww[self.i]
      w0 = w[0]
      if '0' <= w0 and w0 <= '9':
        z = Atom(int(w))
      else:
        z = Intern(self.ww[self.i])
      self.i += 1
      return z

def Run(s, env):
  z = Nil
  p = LispParser(s)
  x = p.Next()
  while x is not None:
    if not x.a and x.h is Defun:
      name = x.t.h
      args = x.t.t.h
      body = x.t.t.t.h
      env = Pair(name, Pair(List3(Lambda, args, body), env))
      x = p.Next()
      continue
    z = x.Eval(env)
    print "RUN: ", x.Show(), "  -->  ", z.Show(), "  ;; ENV: ", env.Show()
    x = p.Next()
  return z

print SplitWhite("( defun double ( x ) ( + x x ) )  ( double 333 )")
print LispParser("( defun double ( x ) ( + x x ) )  ( double 333 )").Next().Show()
Run("( defun double ( x ) ( + x x ) )  ( double 333 )", Nil)
Run("( defun triangle ( x ) ( if ( < x 1 ) 0 ( + x ( triangle ( - x 1 ) ) ) ) )  ( triangle 6 )", Nil)
Run("( defun fib ( x y ) ( if ( < x 2 ) x ( + ( fib ( - x 1 ) ) ( fib ( - x 2 ) ) ) ) )  ( fib 10 )", Nil)
