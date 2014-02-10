import re
import sys

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')

RE_KEYWORDS = re.compile(
    '\\b(class|def|if|else|while|True|False|None|print|and|or|try|except|raise|return|break|continue|pass)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[*]=|//|<<|>>|==|<=|>=|[*][*]')
RE_OPS = re.compile('[-.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_NUM = re.compile('[+-]?[0-9]+[-+.0-9_e]*')
RE_STR = re.compile('(["](([^"\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\n]|[\\\\].)*)[\'])')

TAB_WIDTH = 8

DETECTERS = [
  [RE_KEYWORDS, 'K'],
  [RE_ALFA, 'A'],
  [RE_NUM, 'N'],
  [RE_LONG_OPS, 'L'],
  [RE_OPS, 'O'],
  [RE_GROUP, 'G'],
  [RE_STR, 'S'],
]

ADD_OPS = {
  '+': 'Add',
  '-': 'Sub'
}
REL_OPS = {
  '==': 'EQ',
  '!=': 'NE',
  '<': 'LT',
  '<=': 'LE',
  '>': 'GT',
  '>=': 'GE',
}

IvDone = {}
MaxNumCallArgs = -1

def Bad(format, *args):
  raise Exception(format % args)

class Lex(object):
  def __init__(self, program):
    self.buf = program
    self.i = 0
    self.line_num = 1
    self.indents = [1]
    self.tokens = []
    n = len(self.buf)
    while self.i < n:
      self.DoWhite()
      if self.i < n:
        self.DoBlack()

  def Add(self, x):
    self.tokens.append(x)

  def DoBlack(self):
    rest = self.buf[self.i:]
    for reg, kind in DETECTERS:
      m = reg.match(rest)
      if m:
        got = m.group(0)
        self.Add((kind, got, self.i))
        self.i += len(got)
        return
    raise Bad("Cannot parse (at %d): %s", self.i, repr(rest))

  def DoWhite(self):
    m = RE_WHITE.match(self.buf[self.i:])
    # blank_lines includes all the newlines;
    #   if blank_lines is empty, we're not on a new line.
    # white is the remnant at the front of partial line;
    #   white is the new indentation level.
    blank_lines, _, white = m.groups()
    # both is the entire match.
    both = m.group(0)
    i = self.i
    self.i += len(both)

    if not blank_lines:
      return # White space is inconsequential if not after \n

    self.Add((';;', ';;', i))
    self.line_num += sum([c == '\n' for c in blank_lines])

    col = 1 + TabWidth(white)  # Conventionally, columns start at 1.
    if col < self.indents[-1]:
      # outdent (i.e. un-indent).
      j = len(self.indents) - 1  # For iterating backwards thru indents.
      outage = both  # For recording the white space.
      while col < self.indents[j]:
        # Not back far enough yet.
        self.Add(('OUT', outage, i))
        outage = ''  # We put all the white space in the first OUT.
        j -= 1
        if j < 0 or self.indents[j] < col:
          raise Bad('Cannot un-indent: New column is %d; previous columns are %s', col, repr(self.indents))
        if self.indents[j] == col:
            self.indents[j+1:] = []  # Trim tail to index j.

    elif col > self.indents[-1]:
        # indent
        self.Add(('IN', both, i))
        self.indents.append(col)

    print 'DoWhite', self.i, self.indents, self.tokens, repr(blank_lines), repr(white), repr(self.buf[self.i:])

def TabWidth(s):
  z = 0
  for c in s:
    if c == '\t':
      z = ((z+TAB_WIDTH-1) // TAB_WIDTH) * TAB_WIDTH
    else:
      z += 1
  return z

SerialNum = 10
def Serial(s):
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (s, SerialNum)

class Generator(object):
  def __init__(self, up):
    self.up = up
    self.glbls = {}         # name -> goName
    self.scopes = []
    self.tail = []
    self.cls = ''

  def GenModule(self, name, path, suite):
    print '@@ package main'
    print '@@ import "fmt"'
    print '@@ import . "github.com/strickyak/rye/runt"'
    print '@@ var _ = fmt.Sprintf'
    print '@@ var _ = MkInt'
    print '@@'

    print '@@ func Rye_Module(__name__ P) P {'
    for th in suite.things:
      th.visit(self)
    print '@@   return None'
    print '@@ }'
    print '@@'
    print '\n\n'.join(self.tail)
    print '@@'
    for g in self.glbls:
      print '@@ var %s P' % self.glbls[g]
    print '@@'
    for i in range(MaxNumCallArgs + 1):
      print '@@  type I_%d interface { Call%d(%s) P }' % (i, i, ", ".join(i * ['P']))
    print '@@'
    print '@@ func main() { Rye_Module(MkStr("__main__")) }'

  def Vexpr(self, p):
    print '@@ _ = %s' % p.a.visit(self)

  def Vassign(self, p):
    # a, b
    # Resolve rhs first.
    rhs = p.b.visit(self)

    a = p.a
    if a.__class__ is Tfield:
      # p, field
      x = a.p.visit(self)
      if x == 'self':  # Special optimization for self.
        self.instvars[a.field] = True
	lhs = 'self.S_%s' % a.field
      else:
        lhs = "%s.S_%s" % (x, a.field)

    elif a.__class__ is Tvar:
      # zzz -- wrong.
      if len(self.scopes):
        # Inside a function.
        scope = self.scopes[0]
	if scope.get(a.name):
	  lhs = scope[a.name]
	else:
          lhs = scope[a.name] = 'v_%s' % a.name
      else:
        lhs = a.visit(self)
        # At the modulde level.
        self.glbls[a.name] = lhs

    print '@@   %s = %s' % (lhs, rhs)

  def Vprint(self, p):
    print 'Print p.aa', p.aa
    vv = [a.visit(self) for a in p.aa]
    print 'Print vv', vv
    print '@@   println(%s.String())' % ', '.join(vv)

  def Vassert(self, p):
    print 'Assert', p.x, p.y, p.code
    print '@@   if ! P(%s).Bool() {' % p.x.visit(self)
    print '@@     panic("Assertion Failed:  %s ;  message=" + P(%s).String() )' % (
       p.code.encode('unicode_escape'), "None" if p.y is None else p.y.visit(self) )
    print '@@   }'

  def Vif(self, p):
    print 'IF: p.t', p.t
    print 'IF: p.yes', p.yes
    print 'IF: p.no', p.no
    print '@@   if VP(%s).Bool() {' % p.t.visit(self)
    p.yes.visit(self)
    print '@@   }'
    if p.no:
      print '@@   else {'
      p.no.visit(self)
      print '@@   }'

  def Vwhile(self, p):
    print 'WHILE: p.t', p.t
    print 'WHILE: p.yes', p.yes
    print '@@   for VP(%s).Bool() {' % p.t.visit(self)
    p.yes.visit(self)
    print '@@   }'

  def Vreturn(self, p):
    vv = [a.visit(self) for a in p.aa]
    print '@@   return %s ' % ', '.join(vv)

  def Vlit(self, p):
    if p.k == 'N':
      return 'MkInt(%s)' % p.v
    elif p.k == 'S':
      return 'MkStr("%s")' % eval(p.v).encode('unicode_escape')
    else:
      Bad('Unknown Vlit', p.k, p.v)

  def Vop(self, p):
    if p.b:
      return ' VSP("%s", VP(%s).%s(VP(%s))) ' % (p.op, p.a.visit(self), p.op, p.b.visit(self))
    else:
      raise Bad('Monadic %d not imp' % p.op)

  def Vgetitem(self, p):
    return ' VSP("GetItem", VP(%s).GetItem(VP(%s))) ' % (p.a.visit(self), p.x.visit(self))

  def Vgetitemslice(self, p):
    return ' VSP("GetItemSlice", VP(%s).GetItemSlice(VP(%s), VP(%s), VP(%s))) ' % (
        p.a.visit(self),
	None if p.x is None else p.x.visit(self),
	None if p.y is None else p.y.visit(self),
	None if p.z is None else p.z.visit(self))

  def Vtuple(self, p):
    return 'MkTupleV( %s )' % ', '.join([x.visit(self) for x in p.xx])

  def Vlist(self, p):
    return 'MkListV( %s )' % ', '.join([x.visit(self) for x in p.xx])

  def Vdict(self, p):
    return 'MkDictV( %s )' % ', '.join([x.visit(self) for x in p.xx])

  def Vvar(self, p):
    if p.name == 'self':
      return 'self'
    for s in self.scopes:
      if p.name in s:
        return s[p.name]
    return 'G_%s' % p.name

  def Vcall(self, p):
    # fn, args
    global MaxNumCallArgs
    n = len(p.args)
    MaxNumCallArgs = max(MaxNumCallArgs, n)
    arglist = ', '.join(["VP(%s)" % a.visit(self) for a in p.args])
    return 'VSP("CALL", VP(%s).(I_%d).Call%d(%s))' % (p.fn.visit(self), n, n, arglist)

  def Vfield(self, p):
    # p, field
    x = p.p.visit(self)
    if x == 'self' and self.instvars.get(p.field):  # Special optimization for self instvars.
      return '%s.S_%s' % (x, p.field)
    else:
      return 'P(%s).(I_GET_%s).GET_%s()' % (x, p.field, p.field)

  def Vdef(self, p):
    # name, args, body.
    buf = PushPrint()

    # Tweak args.  Record meth, if meth.
    args = p.args
    if self.cls and len(p.args):
      if p.args[0] != 'self':
        Bad('first arg to method %s is %s; should be self', p.name, p.args[0])
      args = p.args[1:]  # Skip self.
      self.meths[p.name] = [ args, ]  # Could add more, but args will do.

    # prepend new scope dictionary, containing just the args, so far.
    self.scopes = [ dict([(a, 'a_%s' % a) for a in args]) ] + self.scopes

    #################
    # Render the body, but hold it in buf2, because we will prepend the vars.
    buf2 = PushPrint()
    p.body.visit(self)
    PopPrint()
    code2 = buf2.String()
    #################

    if self.cls:
      func = 'func (self *C_%s) M_%d_%s' % (self.cls, len(p.args)-1, p.name)
    else:
      func = 'func F_%d_%s' % (len(p.args), p.name)

    print '@@'
    print '@@ %s(%s) P {' % (func, ', '.join(['a_%s P' % a for a in args]))

    scope = self.scopes[0]
    for v, v2 in scope.items():
      if v2[0] != 'a':  # Skip args
        print "@@   var %s P = None" % v2
    print code2

    # Pop that scope.
    self.scopes = self.scopes[1:]

    print '@@   return None'
    print '@@ }'
    print '@@'

    if self.cls:
      n = len(args)
      print '@@ type pMeth_%d_%s__%s struct { PBase; Rcvr *C_%s }' % (n, self.cls, p.name, self.cls)
      print '@@ func (o *pMeth_%d_%s__%s) Call%d(%s) P {' % (n, self.cls, p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '@@   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
      print '@@ }'
      print '@@'

    else:
      n = len(p.args)
      print '@@ type pFunc_%s struct { PBase }' % p.name
      print '@@ func (o pFunc_%s) Call%d(%s) P {' % (p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '@@   return F_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
      print '@@ }'
      print '@@'
      print '@@ var G_%s = new(pFunc_%s)' % (p.name, p.name)
      print '@@'

    PopPrint()
    code = buf.String()
    self.tail.append(code)

    # The class constructor gets the args of init:
    if self.cls and p.name == '__init__':
      self.args = p.args

  def Vclass(self, p):
    # name, sup, things
    self.cls = p.name
    sup = p.sup if p.sup else 'object'
    self.instvars = {}
    self.meths = {}
    self.args = []

    # Emit all the methods of the class (and possibly other members).
    for x in p.things:
      x.visit(self)
    self.cls = ''

    buf = PushPrint()
    # Emit the struct for the class.
    print '''
@@ type C_%s struct {
@@   C_%s
%s
@@ }
''' % (p.name, sup, '\n'.join(['@@   S_%s   P' % x for x in self.instvars]))

    # The interface for the class.
    print '''
@@ type I_%s interface {
@@   I_%s
@@   C_%s() *C_%s
@@
%s
@@ }
''' % (p.name, sup, p.name, p.name, '/**/')  # TODO: member methods.
    print '''
@@ func (o *C_%s) C_%s() *C_%s {
@@   return o
@@ }
''' % (p.name, p.name, p.name)

    # Make GET and SET interfaces for each instance var and each method.
    print '@@'
    for iv in self.instvars.keys() + self.meths.keys():
      if not IvDone.get(iv):
        print '@@ type I_GET_%s interface { GET_%s() P }' % (iv, iv)
        print '@@ type I_SET_%s interface { SET_%s(P) }' % (iv, iv)
        IvDone[iv] = True

    # For all the instance vars
    print '@@'
    for iv in self.instvars:
      print '@@ func (o *C_%s) GET_%s() P { return o.S_%s }' % (p.name, iv, iv)
      print '@@ func (o *C_%s) SET_%s(x P) { o.S_%s = x }' % (p.name, iv, iv)
      print '@@'
    print '@@'

    # For all the methods
    print '@@'
    for m in self.meths:
      args, = self.meths[m]
      n = len(args)

      formals = ', '.join(['a%d P' for i in range(n)])
      actuals = ', '.join(['a%d' for i in range(n)])

      print '@@ func (o *C_%s) GET_%s() P { return &pMeth_%d_%s__%s { Rcvr: o } }' % (p.name, m, n, p.name, m)

    # The constructor.
    n = len(self.args) - 1 # Subtract 1 because we don't count self.
    print '@@ type pCtor_%d_%s struct { PBase }' % (n, p.name)
    print '@@ func (o pCtor_%d_%s) Call%d(%s) P {' % (n, p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
    print '@@   z := new(C_%s)' % p.name
    for iv in self.instvars:
      print '@@   z.S_%s = None' % iv
    print '@@   z.M_%d___init__(%s)' % (n, (', '.join(['a%d' % i for i in range(n)])))
    print '@@   return z'
    print '@@ }'
    print '@@'
    print '@@ var G_%s = new(pCtor_%d_%s)' % (p.name, n, p.name)
    print '@@'

    self.tail.append(buf.String())
    PopPrint()

  def Vsuite(self, p):
    for x in p.things:
      x.visit(self)

PrintStack= []
def PushPrint():
    global PrintStack
    sys.stdout.flush()
    print >>sys.stderr, 'SAVING'
    PrintStack.append(sys.stdout)
    buf = Buffer()
    sys.stdout = buf
    return buf
def PopPrint():
    global PrintStack
    sys.stdout = PrintStack.pop()
    print >>sys.stderr, 'RESTORED'

class Buffer(object):
  def __init__(self):
    self.b = []
  def write(self, x):
    self.b.append(x)
  def flush(self):
    pass
  def String(self):
    z = ''.join(self.b)
    return z

class Tnode(object):
  def visit(self, a):
    raise Bad('unimplemented visit %s %s', self, type(self))

class Top(Tnode):
  def __init__(self, a, op, b=None):
    self.op = op
    self.a = a
    self.b = b
  def visit(self, a):
    return a.Vop(self)

class Tlit(Tnode):
  def __init__(self, k, v):
    self.k = k
    self.v = v
  def visit(self, a):
    return a.Vlit(self)

class Tvar(Tnode):
  def __init__(self, name):
    self.name = name
  def visit(self, a):
    return a.Vvar(self)

class Ttuple(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, a):
    return a.Vtuple(self)

class Tlist(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, a):
    return a.Vlist(self)

class Tdict(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, a):
    return a.Vdict(self)

class Tsuite(Tnode):
  def __init__(self, things):
    self.things = things
  def visit(self, a):
    return a.Vsuite(self)

class Texpr(Tnode):
  def __init__(self, a):
    self.a = a
  def visit(self, a):
    return a.Vexpr(self)

class Tassign(Tnode):
  def __init__(self, a, b):
    self.a = a
    self.b = b
  def visit(self, a):
    return a.Vassign(self)

class Tprint(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, a):
    return a.Vprint(self)

class Tassert(Tnode):
  def __init__(self, x, y, code):
    self.x = x
    self.y = y
    self.code = code
  def visit(self, a):
    return a.Vassert(self)

class Tif(Tnode):
  def __init__(self, t, yes, no):
    self.t = t
    self.yes = yes
    self.no = no
  def visit(self, a):
    return a.Vif(self)

class Twhile(Tnode):
  def __init__(self, t, yes):
    self.t = t
    self.yes = yes
  def visit(self, a):
    return a.Vwhile(self)

class Treturn(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, a):
    return a.Vreturn(self)

class Tdef(Tnode):
  def __init__(self, name, args, body):
    self.name = name
    self.args = args
    self.body = body
  def visit(self, a):
    return a.Vdef(self)

class Tclass(Tnode):
  def __init__(self, name, sup, things):
    self.name = name
    self.sup = sup
    self.things = things
  def visit(self, a):
    return a.Vclass(self)

class Tcall(Tnode):
  def __init__(self, fn, args):
    self.fn = fn
    self.args = args
  def visit(self, a):
    return a.Vcall(self)

class Tfield(Tnode):
  def __init__(self, p, field):
    self.p = p
    self.field = field
  def visit(self, a):
    return a.Vfield(self)

class Tgetitem(Tnode):
  def __init__(self, a, x):
    self.a = a
    self.x = x
  def visit(self, a):
    return a.Vgetitem(self)

class Tgetitemslice(Tnode):
  def __init__(self, a, x, y, z):
    self.a = a
    self.x = x
    self.y = y
    self.z = z
  def visit(self, a):
    return a.Vgetitemslice(self)

class Parser(object):
  def __init__(self, program, words, p):
    self.program = program  # Source code
    self.words = words      # Lexical output
    self.imports = {}       # path -> name
    self.litInts = {}       # value -> name
    self.litStrs = {}       # value -> name
    self.k = ''
    self.v = ''
    self.p = p
    self.i = 0
    self.Advance()

  def Bad(self, format, *args):
    msg = format % args
    self.Info(msg)
    raise Exception(msg)

  def Info(self, msg):
    sys.stdout.flush()
    print >> sys.stderr, 120 * '#'
    print >> sys.stderr, '   msg = ', msg
    print >> sys.stderr, '   k =', repr(self.k)
    print >> sys.stderr, '   v =', repr(self.v)
    print >> sys.stderr, '   rest =', repr(self.Rest())

  def Advance(self):
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]
    #print '%s GEN: %s' % ( self, self.coder.gen )
    print '<%s|' % repr(self.program[:self.i])
    print '|%s>' % repr(self.program[self.i:])
    print 'Advance(%d, %d) k=<%s> v=<%s>   %s' % (self.p, self.i, self.k, self.v, repr(self.Rest()))

  def Rest(self):
    return self.program[self.i:]

  def VarGlobal(self, id):
    return 'G_%s' % id

  def VarLocal(self, id):
    z = self.coder.lcls.get(id)
    if not z:
        z = 'var_%s' % id
        self.coder.lcls[v] = z
    return z

  def MkTemp(self):
    z = Serial('tmp')
    self.coder.lcls[z] = z
    return z

  def Gen(self, pattern, *args):
    print 'Gen:::', pattern, args
    self.coder.gen.append(pattern % args)

  def Eat(self, v):
    print 'Eating', v
    if self.v != v:
      raise self.Bad('Expected %s, but got %s, at %s', v, self.v, repr(self.Rest()))
    self.Advance()

  def EatK(self, k):
    print 'EatingK', k
    if self.k != k:
      raise self.Bad('Expected Kind %s, but got %s, at %s', k, self.k, repr(self.Rest()))
    self.Advance()

  def Pid(self):
    if self.k != 'A':
      raise self.Bad('Pid expected kind A, but got %s %q', self.k, repr(self.Rest()))
    z = self.v
    self.Advance()
    return z

  def Xprim(self):
    if self.k == 'N':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    elif self.k == 'S':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    elif self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z

    elif self.v == '(':
      self.Eat('(')
      if self.v == ')':
        self.Eat(')')
        # Unit tuple.
        return self.Ttuple([])
      z = self.Xexpr()
      if self.v == ')':
        # Not a tuple, just an Xexpr.
        self.Eat(')')
        return z
      # z is just the first in a tuple.
      self.Eat(',')
      z = [z]
      while self.v != ')':
        x = self.Xexpr()
	z.append(x)
	if self.v == ')':
	  # Omitted trailing ','
	  break
	self.Eat(',')
      self.Eat(')')
      return Tlist(z)

    elif self.v == '[':
      self.Eat('[')
      z = []
      while self.v != ']':
        x = self.Xexpr()
	z.append(x)
	if self.v == ']':
	  # Omitted trailing ','
	  break
	self.Eat(',')
      self.Eat(']')
      return Tlist(z)

    elif self.v == '{':
      self.Eat('{')
      z = []
      while self.v != '}':
        x = self.Xexpr()
	self.Eat(':')
        y = self.Xexpr()
	z.append(x)
	z.append(y)
	if self.v == '}':
	  # Omitted trailing ','
	  break
	self.Eat(',')
      self.Eat('}')
      return Tdict(z)

    else:
      raise self.Bad('Expected Xprim, but got %s, at %s' % (self.v, repr(self.Rest())))

  def Xsuffix(self):
    """Tcall, Tfield, or Tindex"""
    a = self.Xprim()
    while True:
      if self.v == '(':
        self.Eat('(')
        args = []
        while self.v != ')':
	  b = self.Xexpr()
	  args.append(b)
	  if self.v == ',':
	    self.Eat(',')
	  else:
	    break
        self.Eat(')')
        a = Tcall(a, args)

      elif self.v == '.':
        self.Eat('.')
	field = self.v
        self.EatK('A')
	a = Tfield(a, field)

      elif self.v == '[':
        self.Eat('[')
	if self.v != ':':
	  x = self.Xexpr()
	else:
	  x = None

	if self.v == ']':
	  self.Eat(']')
	  if x is None:
	    raise Bad('Index cannot be None')
	  a = Tgetitem(a, x)
	else:
	  self.Eat(':')
	  if self.v != ']':
	    y = self.Xexpr()
	  else:
	    y = None
	  self.Eat(']')
	  a = Tgetitemslice(a, x, y, None)

      else:
        break
    return a

  def Xadd(self):
    a = self.Xsuffix()
    while self.v in ADD_OPS or (self.k == 'N' and self.v[0] in '+-'):
      if self.k == 'N':
        op = self.v[0]  # The sign is the op.
        b = Tlit('N', self.v[1:])  # The rest of the number.
        self.Advance()
      else:
        op = self.v
        self.Eat(op)
        b = self.Xsuffix()
      a = Top(a, ADD_OPS[op], b)
    return a

  def Xrelop(self):
    a = self.Xadd()
    while self.v in REL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xadd()
      a = Top(a, REL_OPS[op], b)
    return a

  def Xexpr(self):
    return self.Xrelop()

  def Csuite(self):
    things = []
    while self.k != 'OUT' and self.k is not None:
      print 'Csuite', self.k, self.v
      if self.v == ';;':
        self.EatK(';;')
      else:
        t = self.Command();
	if t:
	  things.append(t)
    return Tsuite(things)

  def Command(self):
    if self.v == 'print':
      return self.Cprint()
    elif self.v == 'if':
      return self.Cif()
    elif self.v == 'while':
      return self.Cwhile()
    elif self.v == 'return':
      return self.Creturn()
    elif self.v == 'def':
      return self.Cdef('')
    elif self.v == 'class':
      return self.Cclass()
    elif self.v == 'assert':
      return self.Cassert()
    elif self.v == 'pass':
      self.Eat('pass')
      return
    elif self.k == 'A':
      return self.Cother()
    else:
      raise self.Bad('Unknown stmt: %s %s %s', self.k, self.v, repr(self.Rest()))

  def Cother(self):
    print 'Cother...'
    a = self.Xexpr()
    print 'Cother...a', a
    if self.v in ['=', '+=', '*=']:
      self.Eat(self.v)
      b = self.Xexpr()
      print 'Cother...b', b
      return Tassign(a, b)
    else:
      return Texpr(a)

  def Cprint(self):
    self.Eat('print')
    t = self.Xexpr()
    return Tprint([t])

  def Cassert(self):
    i = self.i
    self.Eat('assert')
    x = self.Xexpr()
    y = None
    j = self.i
    if self.v == ',':
      self.Eat(',')
      y = self.Xexpr()
    return Tassert(x, y, self.program[i:j])

  def Cif(self):
    self.Eat('if')
    t = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    yes = self.Csuite()
    self.EatK('OUT')
    no = None
    if self.v == 'else':
      self.Eat('else')
      self.Eat(':')
      self.EatK(';;')
      self.EatK('IN')
      no = self.Csuite()
      self.EatK('OUT')
    return Tif(t, yes, no)

  def Cwhile(self):
    self.Eat('while')
    t = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    yes = self.Csuite()
    self.EatK('OUT')
    return Twhile(t, yes)

  def Creturn(self):
    self.Eat('return')
    t = self.Xexpr()
    return Treturn([t])

  def Cclass(self):
    self.Eat('class')
    name = self.Pid()
    # No superclasses yet
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')

    things = []
    while self.k != 'OUT':
      if self.v == 'def':
        t = self.Cdef(name)
	things.append(t)
      elif self.v == 'pass':
        self.Eat('pass')
      elif self.k == ';;':
        self.EatK(';;')
      else:
        raise self.Bad('Classes may only contain "def" or "pass" commands.')
    self.EatK('OUT')

    return Tclass(name, None, things)

  def Cdef(self, cls):
    self.Eat('def')
    name = self.Pid()
    print 'Cdef -------- %q :: name', cls, name
    self.Eat('(')
    args = []
    while self.k == 'A':
      arg = self.Pid()
      if self.v == ',':
        self.Eat(',')
      args.append(arg)
    print 'Cdef -------- %q :: args', cls, args
    self.Eat(')')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    suite = self.Csuite()
    self.EatK('OUT')
    return Tdef(name, args, suite)

def Dump(x, pre='/'):
  t = type(x)
  print '###', pre, '----', t.__name__, '::', repr(x)
  if str(t)[:4] == '<cla':
    for i in vars(x):
      if i[0] != '_':
        Dump(getattr(x, i), pre + '/' + i)
  elif t == list:
    for i in range(len(x)):
      Dump(x[i], pre + '/!' + str(i))
  elif t == dict:
    for i in x:
      Dump(x[i], pre + '/@' + str(i))


pass
