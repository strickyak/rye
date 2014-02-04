import re
import sys

RE_WHITE = re.compile('([ \t\n]*[\n])?([ \t]*)')

RE_KEYWORDS = re.compile(
    '\\b(class|def|if|else|while|True|False|None|print|and|or|try|except|raise|return|break|continue|pass)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[*]=|//|<<|>>|<=|>=|[*][*]')
RE_OPS = re.compile('[.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_NUM = re.compile('[+-]?[0-9]+[-+.0-9_e]*')
RE_STR = re.compile('(["](([^"\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\n]|[\\\\].)*)[\'])')
RE_REM = re.compile('[#]')

TAB_WIDTH = 8

DETECTERS = [
  [RE_KEYWORDS, 'K'],
  [RE_LONG_OPS, 'L'],
  [RE_OPS, 'O'],
  [RE_GROUP, 'G'],
  [RE_ALFA, 'A'],
  [RE_NUM, 'N'],
  [RE_STR, 'S'],
  [RE_REM, 'R'],
]

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
    blank_lines, white = m.groups()
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
    self.scope = []
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
    print '@@ func main() { Rye_Module(MkStr("__main__")) }'

  def Vexpr(self, p):
    print '@@ _ = %s' % p.a.visit(self)

  def Vassign(self, p):
    a = p.a
    if a.__class__ is Tfield:
      x = a.p.visit(self)
      if x == 'self':
        self.instvars[a.field] = True
	lhs = 'self.S_%s' % a.field
      else:
        Bad('Can only assign to self.Field: %s %s', x, a.field)
    elif a.__class__ is Tvar:
      lhs = a.visit(self)
    print '@@ %s = %s' % (lhs, p.b.visit(self))

  def Vprint(self, p):
    print 'p.aa', p.aa
    vv = [a.visit(self) for a in p.aa]
    print 'vv', vv
    print '@@   println(%s.String())' % ', '.join(vv)

  def Vreturn(self, p):
    vv = [a.visit(self) for a in p.aa]
    print '@@   return %s ' % ', '.join(vv)

  def Vlit(self, p):
    if p.k == 'N':
      return 'MkInt(%s)' % p.v
    elif p.k == 'S':
      return 'MkStr(%s)' % p.v
    else:
      Bad('Unknown Vlit', p.k, p.v)

  def Vop(self, p):
    if p.b:
      return ' ( %s.%s(%s) ) ' % (p.a.visit(self), p.op, p.b.visit(self))

  def Vvar(self, p):
    if p.name == 'self':
      return 'self'
    for s in self.scope:
      if p.name in s:
        return s[p.name]
    return 'G_%s' % p.name

  def Vcall(self, p):
    s = '[]P{ %s }' % ', '.join([a.visit(self) for a in p.args])
    return '(%s).Call(%s ...)' % (p.fn.visit(self), s)

  def Vdef(self, p):
    # name, args, body.
    sys.stdout.flush()
    print >>sys.stderr, 'SAVING'
    save = sys.stdout
    buf = Buffer()
    sys.stdout = buf

    if self.cls:
      func = 'func (self *C_%s) M_%s' % (self.cls, p.name)
    else:
      func = 'func F_%s' % p.name

    print '@@'
    print '@@ %s(%s) P {' % (func, ', '.join(['a_%s P' % a for a in p.args]))

    # prepend new scope dictionary, containing just the args, so far.
    self.scope = [ dict([(a, 'a_%s' % a) for a in p.args]) ] + self.scope

    p.body.visit(self)

    # Pop that scope.
    self.scope = self.scope[1:]

    print '@@   return None'
    print '@@ }'
    print '@@'

    if self.cls:
      print '@@ var G_%s__%s = &PMeth{ Meth: func(rcvr P, args []P) P {' % (self.cls, p.name)
      print '@@   if len(args) != %d {' % len(p.args)
      print '@@     panic(fmt.Sprintf("method %s::%s got %%d args, wanted %d args", len(args)))' % (self.cls, p.name, len(p.args))
      print '@@   }'
      print '@@   return rcvr.(*PObj).Obj.(*C_%s).M_%s(%s)' % (self.cls, p.name, ', '.join(['args[%d]' % i for i in range(len(p.args))]))
      print '@@ }}'
    else:
      print '@@ var G_%s = &PFunc{ Fn: func(args []P) P {' % p.name
      print '@@   if len(args) != %d {' % len(p.args)
      print '@@     panic(fmt.Sprintf("func %s got %%d args, wanted %d args", len(args)))' % (p.name, len(p.args))
      print '@@   }'
      print '@@   return F_%s(%s)' % (p.name, ', '.join(['args[%d]' % i for i in range(len(p.args))]))
      print '@@ }}'

    code = buf.String()
    self.tail.append(code)
    sys.stdout = save
    print >>sys.stderr, 'RESTORED'

  def Vclass(self, p):
    # name, sup, things
    self.cls = p.name
    self.instvars = {}
    for x in p.things:
      x.visit(self)
    self.cls = ''
    self.tail.append("@@ type C_%s struct {\n%s\n@@ }" % (
      p.name, '\n'.join(['@@   S_%s   P' % x for x in self.instvars])))

  def Vsuite(self, p):
    for x in p.things:
      x.visit(self)


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

class Parser(object):
  def __init__(self, program, words, p):
    self.program = program  # Source code
    self.words = words      # Lexical output
    self.imports = {}       # path -> name
    self.glbls = {}         # name -> goName
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

  def LitInt(self, v):
    z = self.litInts.get(v)
    if not z:
        z = Serial('lit_int')
        self.litInts[v] = z
    return z

  def LitStr(self, v):
    z = self.litStrs.get(v)
    if not z:
        z = Serial('lit_str')
        self.litStrs[v] = z
    return z

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
      #z = self.LitInt(self.v)
      z = Tlit(self.k, self.v)
      self.Advance()
      return z
    elif self.k == 'S':
      #z = self.LitStr(self.v)
      z = Tlit(self.k, self.v)
      self.Advance()
      return z
    elif self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z
    elif self.v == '(':
        self.Advance()
        z = self.Xexpr()
        self.Eat(')')
        return z
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

      # TODO: Tindex

      else:
        break
    return a

  def Xadd(self):
    a = self.Xsuffix()
    while self.v in '+-':
      op = self.v
      self.Advance()
      b = self.Xsuffix()
      fns = {'+': 'Add', '-': 'Sub'}
      a = Top(a, fns[op], b)
    return a

  def Xexpr(self):
    return self.Xadd()

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
    elif self.v == 'return':
      return self.Creturn()
    elif self.v == 'def':
      return self.Cdef('')
    elif self.v == 'class':
      return self.Cclass()
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
