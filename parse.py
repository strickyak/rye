#import codecs
import md5
import os
import re
import sys

rye_rye = False
if rye_rye:
  from lib import data
  from go import strconv

RYE_FLOW = os.getenv('RYE_FLOW')
BUILTINS = list( 'go_cast go_type go_new go_make go_append'.split())

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([ \t\n]*[#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')
RE_PRAGMA = re.compile('[ \t]*[#][#][A-Za-z:()]+')

RE_KEYWORDS = re.compile(
    '\\b(del|say|from|class|def|native|if|elif|else|while|True|False|None|print|and|or|try|except|raise|yield|return|break|continue|pass|as|go|defer|with|global|assert|must|lambda|switch)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[-]=|[*]=|/=|//|<<|>>>|>>|==|!=|<=|>=|[*][*]|[.][.]')
RE_OPS = re.compile('[-.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_FLOAT = re.compile('[+-]?[0-9]+[.][-+0-9eE]*')
RE_INT = re.compile('(0[Xx][0-9A-Fa-f]+|[+-]?[0-9]+)')

RE_STR = re.compile('(["](([^"\\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\\n]|[\\\\].)*)[\'])')
RE_STR2 = re.compile('(?s)[`]([^`]*)[`]')
RE_STR3 = re.compile('(?s)("""(([^\\\\]|[\\\\].)*?)"""|\'\'\'(([^\\\\]|[\\\\].)*?)\'\'\')')

RE_SEMI = re.compile(';')

RE_WORDY_REL_OP = re.compile('\\b(not\\s+in|is\\s+not|in|is)\\b')
RE_NOT_IN = re.compile('^not\\s+in$')
RE_IS_NOT = re.compile('^is\\s*not$')

RE_NOT_NEWLINE = re.compile('[^\\n]')

### Experimental: For string interpolation, if we do that:
# RE_NEST1 = '[^()]*([(][^()]*[)][^()]*)*[^()]*'
# RE_SUBST = re.compile('(.*)[\\\\][(](' + NEST1 + ')[)](.*)')

TAB_WIDTH = 8

DETECTERS = [
  [RE_PRAGMA, 'P'],
  [RE_KEYWORDS, 'K'],
  [RE_WORDY_REL_OP, 'W'],
  [RE_ALFA, 'A'],
  [RE_FLOAT, 'F'],
  [RE_INT, 'N'],
  [RE_LONG_OPS, 'L'],
  [RE_OPS, 'O'],
  [RE_GROUP, 'G'],
  [RE_STR3, 'S'],
  [RE_STR2, 'S'],
  [RE_STR, 'S'],
  [RE_SEMI, ';;'],
]

UNARY_OPS = {
  '+': 'UnaryPlus',
  '-': 'UnaryMinus',
  '~': 'UnaryInvert',
}
SHIFT_OPS = {
  '<<': 'ShiftLeft',
  '>>': 'ShiftRight',
  '>>>': 'UnsignedShiftRight',
}
ADD_OPS = {
  '+': 'Add',
  '-': 'Sub',
}
MUL_OPS = {
  '*': 'Mul',
  '/': 'Div',
  '//': 'IDiv',
  '%': 'Mod',
}
REL_OPS = {
  '==': 'EQ',
  '!=': 'NE',
  '<': 'LT',
  '<=': 'LE',
  '>': 'GT',
  '>=': 'GE',
}

MaxNumCallArgs = -1

FIRST_WORD = re.compile('^([^\\s]*)').match
def FirstWord(s):
  return FIRST_WORD(s).group(1)

TRIM_PRAGMA = re.compile('\\s*[#][#](\\w*)').match
def TrimPragma(s):
  m = TRIM_PRAGMA(s)
  if m:
    return m.group(1)
  raise Exception('Bad pragma: %s' % repr(s))

NOT_PRINTABLE_ASCII = re.compile('[^!-~]')
NONALFA = re.compile('[^A-Za-z0-9]')
TROUBLE_CHAR = re.compile('[^]-~ !#-Z[]')
def GoStringLiteral(s):
  if rye_rye:
    return strconv.QuoteToASCII(s)
  else:
    return '"' + TROUBLE_CHAR.sub((lambda m: '\\x%02x' % ord(m.group(0))), s) + '"'

def CleanIdentWithSkids(s):
  if len(s) < 50:
    # Work around lack of callable() for .sub in RYE.
    return md5.new(s).hexdigest()
    # TODO = NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
    # return NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
  else:
    return md5.new(s).hexdigest()

def Bad(format, *args):
  raise Exception(format % args)

############################################################

SerialNum = 10
def Serial(s):
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (s, SerialNum)

# p might be abosolute; but more are always relative.
# returns a list of part names.
def CleanPath(cwd, p, *more):
  if p.startswith('/'):
    # Absolute.
    q = p.split('/')
  else:
    # Relative
    q = cwd.split('/') + p.split('/')
  for m in more:
    q += m.split('/')

  x = []
  for w in q:
    if w == '' or w == '.':
      continue
    elif w == '..':
      if x:
        x = x[:-1]
      else:
        raise Exception('Bad path (too many ".."): "%s" in "%s"' % (p, cwd))
    elif w.startswith('.'):
      raise Exception('Bad word in path (starts with "."): "%s" in "%s"' % (p, cwd))
    else:
      x.append(w)
  return x

class Tnode(object):
  def __init__(self):
    self.where = None
    self.line = None
    self.gloss = None
  def visit(self, a):
    raise Bad('unimplemented visit %s %s', self, type(self))

class Top(Tnode):
  def __init__(self, a, op, b=None, returns_bool=False):
    self.a = a
    self.op = op
    self.b = b
    self.returns_bool = returns_bool
  def visit(self, v):
    return v.Vop(self)

class Tboolop(Tnode):
  def __init__(self, a, op, b=None):
    self.a = a
    self.op = op
    self.b = b
  def visit(self, v):
    return v.Vboolop(self)

class Tcondop(Tnode):
  def __init__(self, a, b, c):
    self.a = a
    self.b = b
    self.c = c
  def visit(self, v):
    return v.Vcondop(self)

class Tgo(Tnode):
  def __init__(self, fcall):
    self.fcall = fcall
  def visit(self, v):
    return v.Vgo(self)

class Traw(Tnode):
  def __init__(self, raw):
    self.raw = raw
  def visit(self, v):
    return v.Vraw(self)

class Tlit(Tnode):
  def __init__(self, k, v):
    self.k = k
    self.v = v
  def visit(self, v):
    return v.Vlit(self)

class Tvar(Tnode):
  def __init__(self, name):
    self.name = name
  def visit(self, v):
    return v.Vvar(self)

class Titems(Tnode):
  def __init__(self, xx, trailing_comma=False):
    self.xx = xx
    self.trailing_comma = trailing_comma
  def visit(self, v):
    return v.Vtuple(self)  # By default, make a tuple.

class Ttuple(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vtuple(self)

class Tlist(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vlist(self)

class Tlambda(Tnode):
  def __init__(self, lvars, lexpr, where, line):
    self.lvars = lvars
    self.lexpr = lexpr
    self.where = where
    self.line = line
  def visit(self, v):
    return v.Vlambda(self)

class Tforexpr(Tnode):
  def __init__(self, z, vv, ll, cond):
    self.z = z
    self.vv = vv
    self.ll = ll
    self.cond = cond
  def visit(self, v):
    return v.Vforexpr(self)

class Tdict(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vdict(self)

class Tsuite(Tnode):
  def __init__(self, things):
    self.things = things
  def visit(self, v):
    return v.Vsuite(self)

class Texpr(Tnode):
  def __init__(self, a):
    self.a = a
  def visit(self, v):
    return v.Vexpr(self)

class Tassign(Tnode):
  def __init__(self, a, b, pragma=None):
    self.a = a
    self.b = b
    self.pragma = pragma
  def visit(self, v):
    return v.Vassign(self)

class Tprint(Tnode):
  def __init__(self, w, xx, saying, code):
    self.w = w
    self.xx = xx
    self.saying = saying
    self.code = code
  def visit(self, v):
    return v.Vprint(self)

class Twithdefer(Tnode):
  def __init__(self, call, body):
    self.call = call
    self.body = body
  def visit(self, v):
    return v.Vwithdefer(self)

class Tglobal(Tnode):
  def __init__(self, vars):
    self.vars = vars
  def visit(self, v):
    return v.Vglobal(self)

class Timport(Tnode):
  def __init__(self, imported, alias, fromWhere):
    self.imported = imported
    self.alias = alias
    self.fromWhere = fromWhere
  def visit(self, v):
    return v.Vimport(self)

class Tassert(Tnode):
  def __init__(self, x, y, code, is_must, fails):
    self.x = x
    self.y = y
    self.code = code
    self.is_must = is_must
    self.fails = fails
  def visit(self, v):
    return v.Vassert(self)

class Ttry(Tnode):
  def __init__(self, tr, exvar, ex):
    self.tr = tr
    self.exvar = exvar
    self.ex = ex
  def visit(self, v):
    return v.Vtry(self)

class Tswitch(Tnode):
  def __init__(self, a, cases, clauses, default_clause):
    self.a = a
    self.cases = cases
    self.clauses = clauses
    self.default_clause = default_clause
  def visit(self, v):
    return v.Vswitch(self)

class Tif(Tnode):
  def __init__(self, t, yes, no):
    self.t = t
    self.yes = yes
    self.no = no
  def visit(self, v):
    return v.Vif(self)

class Twhile(Tnode):
  def __init__(self, t, yes):
    self.t = t
    self.yes = yes
  def visit(self, v):
    return v.Vwhile(self)

class Tfor(Tnode):
  def __init__(self, var, t, b):
    self.var = var
    self.t = t
    self.b = b
  def visit(self, v):
    return v.Vfor(self)

class Treturn(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, v):
    return v.Vreturn(self)

class Tyield(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, v):
    return v.Vyield(self)

class Tbreak(Tnode):
  def __init__(self):
    pass
  def visit(self, v):
    return v.Vbreak(self)

class Tcontinue(Tnode):
  def __init__(self):
    pass
  def visit(self, v):
    return v.Vcontinue(self)

class Traise(Tnode):
  def __init__(self, a):
    self.a = a
  def visit(self, v):
    return v.Vraise(self)

class Tdel(Tnode):
  def __init__(self, listx):
    self.listx = listx
  def visit(self, v):
    return v.Vdel(self)

class Tnative(Tnode):
  def __init__(self, ss):
    self.ss = ss
  def visit(self, v):
    return v.Vnative(self)

class Tdef(Tnode):
  def __init__(self, name, args, dflts, star, starstar, body):
    self.name = name
    self.args = args
    self.dflts = dflts
    self.star = star
    self.starstar = starstar
    self.body = body

    self.argsPlus = args[:]
    if star:
      self.argsPlus += [star]
    if starstar:
      self.argsPlus += [starstar]

    if len(args) != len(dflts):
      raise Exception('len args (%s) != len dflts (%s) ::: %s' % (args, dflts, (self)))

  def visit(self, v):
    return v.Vdef(self)

class Tclass(Tnode):
  def __init__(self, name, sup, things):
    self.name = name
    self.sup = sup
    self.things = things
  def visit(self, v):
    return v.Vclass(self)

class Tcall(Tnode):
  def __init__(self, fn, args, names, star, starstar):
    self.fn = fn
    self.args = args
    self.names = names
    self.star = star
    self.starstar = starstar
  def visit(self, v):
    return v.Vcall(self)

class Tfield(Tnode):
  def __init__(self, p, field):
    self.p = p
    self.field = field
  def visit(self, v):
    return v.Vfield(self)

class Tgetitem(Tnode):
  def __init__(self, a, x):
    self.a = a
    self.x = x
  def visit(self, v):
    return v.Vgetitem(self)

class Tgetitemslice(Tnode):
  def __init__(self, a, x, y, z):
    self.a = a
    self.x = x
    self.y = y
    self.z = z
  def visit(self, v):
    return v.Vgetitemslice(self)

class Tcurlysetter(Tnode):
  def __init__(self, obj, vec):
    self.obj = obj  # The object
    self.vec = vec  # Pairs of var, expr
  def visit(self, v):
    return v.Vcurlysetter(self)

class Parser(object):
  def __init__(self, program, words, p, cwp):
    self.program = program  # Source code
    self.words = words      # Lexical output
    self.litInts = {}       # value -> name
    self.litStrs = {}       # value -> name
    self.k = ''
    self.v = ''
    self.p = p
    self.cwp = cwp
    self.i = 0
    self.line = 1
    self.Advance()

  def Bad(self, format, *args):
    msg = format % args
    self.Info(msg)
    raise Exception(msg)

  def Info(self, msg):
    try:
      sys.stdout.flush()
    except:
      print >> sys.stderr, '((( cannot flush )))'
    print >> sys.stderr, 120 * '#'
    print >> sys.stderr, '   msg = ', msg
    print >> sys.stderr, '   k =', repr(self.k)
    print >> sys.stderr, '   v =', repr(self.v)
    print >> sys.stderr, '   rest =', repr(self.Rest())

  def LookAheadV(self):
    if self.p+1 < len(self.words):
      return self.words[self.p+1][1]

  def Advance(self):
    oldi = self.i
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]
    newi = self.i
    newlines = RE_NOT_NEWLINE.sub('', self.program[oldi:newi])
    self.line += len(newlines)

  def Rest(self):
    return self.program[self.i:]

  def Eat(self, v):
    if self.v != v:
      raise self.Bad('Expected %s, but got %s, at %s', v, self.v, repr(self.Rest()))
    self.Advance()

  def EatK(self, k):
    if self.k != k:
      raise self.Bad('Expected Kind %s, but got %s, at %s', k, self.k, repr(self.Rest()))
    self.Advance()

  def Pid(self):
    if self.k != 'A':
      raise self.Bad('Pid expected kind A, but got kind=%s; rest=%s', self.k, repr(self.Rest()))
    z = self.v
    self.Advance()
    return z

  def Xvar(self):
    if self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z
    else:
      raise self.Bad('Xvar expected variable name, but got kind=%s; rest=%s', self.k, repr(self.Rest()))

  def Xqualname(self):  # A possibly-singly-qualifed name (qualifier should be an import).
    if self.k == 'A':
      a = Tvar(self.v)
      self.Advance()
      if self.v == '.':
        self.Advance()
        field = self.v
        self.EatK('A')
        return Tfield(a, field)
      else:
        return a
    else:
      raise self.Bad('Xqualname expected variable name, but got "%s"; rest=%s', self.v, repr(self.Rest()))

  def Xprim(self):
    if self.k == 'N':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.k == 'F':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.k == 'S':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.v == 'go':
      self.Advance()
      fcall = self.Xsuffix()
      if type(fcall) != Tcall:
        raise Exception('Go expression must be func or method call: %s' % fcall)
      z = Tgo(fcall)
      return z

    if self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z

    if self.v == '.': # Experimental self-less dot.
      self.Eat('.')
      field = self.v
      self.EatK('A')
      return Tfield(Tvar('self'), field)

    if self.k == 'K':
      if self.v in ['None', 'True', 'False']:
        v = self.v
        self.Eat(self.v)
        z = Traw(v)
        return z
      raise Exception('Keyword "%s" cannot be used as a variable or expression' % self.v)

    if self.v == '(':
      self.Eat('(')
      if self.v == ')':
        self.Eat(')')
        # Unit tuple.
        return Ttuple([])
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
      return Ttuple(z)

    if self.v == '[':
      has_comma = False
      self.Eat('[')
      z = []
      while self.v != ']':
        x = self.Xexpr()
        z.append(x)
        if self.v == ']':
          # Omitted trailing ','
          break
        if self.v == 'for':
          # For expression (list interpolation)
          if has_comma:
            raise Exception('"for" not allowed after a bare tuple; use parens.')
          self.Eat('for')
          vv = self.Xvars()
          self.Eat('in')
          ll = self.Xor()  # Avoid confusion with 'if' Xcond
          cond = None
          if self.v == 'if':
            self.Eat('if')
            cond = self.Xexpr()
          self.Eat(']')
          assert len(z) == 1
          return Tforexpr(z[0], vv, ll, cond)
        self.Eat(',')
        has_comma = True
      self.Eat(']')
      return Tlist(z)

    if self.v == '{':
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
      raise self.Bad('Expected Xprim, but got %s, at %s', self.v, repr(self.Rest()))

  def Xsuffix(self):
    """Tcall, Tfield, or Tindex"""
    a = self.Xprim()
    while True:
      if self.v == '(':
        self.Eat('(')

        args = []
        names = []
        star = None
        starstar = None
        canBare = True
        while self.v != ')':

          # Look for case with named parameter
          named = ''
          starred = ''
          if self.k == 'A' and self.LookAheadV() == '=':
            named = self.v
            self.EatK('A')
            self.Eat('=')
            canBare = False
          elif self.v in ['*', '**']:
            starred = self.v
            self.Eat(starred)
            canBare = False
            if starred == '*' and star:
              raise Exception('Cannot have two * args')
            if starred == '**' and starstar:
              raise Exception('Cannot have two * args')
          else:
            if not canBare:
              raise Exception(
                  'Cannot have arg without name or * or **, after (names=%s star=%s starstar=%s)' % (
                      names, star is not None, starstar is not None))

          b = self.Xexpr()
          if starred == '*':
            star = b
          elif starred == '**':
            starstar = b
          else:
            args.append(b)
            names.append(named)

          if self.v == ',':
            self.Eat(',')
          else:
            break
        self.Eat(')')
        a = Tcall(a, args, names, star, starstar)

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

      elif self.v == '{':
        # Curly Setter: can follow any Rye or Go object.
        # Looks kind of like the Curly-Brace initializer in Go.
        # But instead of starting with a type name, start with an object.
        #   e.g.
        #     b = gonew(fruit.Banana) {Color: 'yellow'}
        #   or
        #     opts = file.Opts() {How: 'r+', Mode: 438, Buffer: 4096}
        # This both changes the object and returns it.
        self.Eat('{')
        vec = []
        while self.v != '}':
          var = self.Xvar()
          self.Eat(':')
          value = self.Xexpr()
          vec.append((var, value))

          if self.v == ',':
            self.Eat(',')
          elif self.v != '}':
            raise Exception('Expected "," or "}" in Field Constructor, but got "%s"' % self.v)
        self.Eat('}')
        return Tcurlysetter(a, vec)

      else:
        break
    return a

  def Xunary(self):
    if self.v in UNARY_OPS:
      op = self.v
      self.Eat(op)
      a = self.Xunary()
      return Top(a, UNARY_OPS[op], None)
    else:
      return self.Xsuffix()

  def Xmul(self):
    a = self.Xunary()
    while self.v in MUL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xunary()
      a = Top(a, MUL_OPS[op], b)
    return a

  def Xadd(self):
    a = self.Xmul()
    while self.v in ADD_OPS or (self.k == 'N' and self.v[0] in '+-'):
      if self.k == 'N':
        op = self.v[0]  # The sign is the op.
        b = Tlit('N', self.v[1:])  # The rest of the number.
        self.Advance()
      else:
        op = self.v
        self.Eat(op)
        b = self.Xmul()
      a = Top(a, ADD_OPS[op], b)
    return a

  def Xshift(self):
    a = self.Xadd()
    while self.v in SHIFT_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xadd()
      a = Top(a, SHIFT_OPS[op], b)
    return a

  def Xbitand(self):
    a = self.Xshift()
    while self.v == '&':
      op = self.v
      self.Eat(op)
      b = self.Xshift()
      a = Top(a, "BitAnd", b)
    return a

  def Xbitxor(self):
    a = self.Xbitand()
    while self.v == '^':
      op = self.v
      self.Eat(op)
      b = self.Xbitand()
      a = Top(a, "BitXor", b)
    return a

  def Xbitor(self):
    a = self.Xbitxor()
    while self.v == '|':
      op = self.v
      self.Eat(op)
      b = self.Xbitxor()
      a = Top(a, "BitOr", b)
    return a

  def Xrelop(self):
    a = self.Xbitor()
    if self.v in REL_OPS:

      chain = None
      while self.v in REL_OPS:
        op = self.v
        self.Eat(op)
        b = self.Xbitor()
        rel = Top(a, REL_OPS[op], b, True)
        if chain:
          chain = Tboolop(chain, "&&", rel)
        else:
          chain = rel
        a = b # For chaining.
      a = chain # For return value.

    elif RE_WORDY_REL_OP.match(self.v):
      op = self.v
      self.Eat(op)
      b = self.Xbitor()
      if op == 'in':
        a = Top(b, 'Contains', a, True)    # N.B. swap a & b for Contains
      elif RE_NOT_IN.match(op):
        a = Top(b, 'NotContains', a, True)    # N.B. swap a & b for NotContains
      elif op == 'is':
        a = Top(a, 'Is', b, True)
      elif RE_IS_NOT.match(op):
        a = Top(b, 'IsNot', a, True)
      else:
        raise Exception('Weird RE_WORDY_REL_OP: %s' % op)
    return a

  def Xnot(self):
    if self.v == 'not':
      self.Eat('not')
      b = self.Xrelop()
      return Tboolop(b, "!")
    else:
      return self.Xrelop()

  def Xand(self):
    a = self.Xnot()
    while self.v == 'and':
      op = self.v
      self.Eat(op)
      b = self.Xnot()
      a = Tboolop(a, "&&", b)
    return a

  def Xor(self):
    a = self.Xand()
    while self.v == 'or':
      op = self.v
      self.Eat(op)
      b = self.Xand()
      a = Tboolop(a, "||", b)
    return a

  def Xcond(self):
    z = self.Xor()
    if self.v == 'if':
      self.Eat('if')
      a = self.Xcond()
      self.Eat('else')
      c = self.Xcond()
      z = Tcondop(a, z, c)
    return z

  def Xlambda(self):
    where = self.i
    line = self.i
    if self.v != 'lambda':
      return self.Xcond()
    self.Eat('lambda')
    lvars = self.Xvars(allowEmpty=True)
    self.Eat(':')
    lexpr = self.Xitems(allowScalar=False, allowEmpty=True)
    return Tlambda(lvars, lexpr, where, line)

  def Xexpr(self):
    return self.Xlambda()

  def Xlistexpr(self):
    return self.Xitems(allowScalar=True, allowEmpty=False)

  def Xvars(self, allowEmpty=False):
    z = []
    has_comma = False
    while True:
      if self.v == '(':
        self.Eat('(')
        x = self.Xvars(allowEmpty)
        self.Eat(')')
        z.append(x)
      elif self.k == 'A':
        z.append(Tvar(self.v))
        self.Advance()
      else:
        break

      if self.v != ',':
        break
      self.Eat(',')
      has_comma = True

    if not allowEmpty and not z:
      raise Exception('Expected variable name')

    if has_comma or not z:
      return Titems(z)  # List of items.
    else:
      return z[0]

  def Xitems(self, allowScalar, allowEmpty):
    "A list of expressions, possibly empty, or possibly a scalar."
    z = []
    comma_needed = False  # needed before more in the list.
    had_comma = False
    trailing_comma = False
    while self.k != ';;' and self.k != 'P' and self.v not in [')', ']', '}', ':', '=', '+=', '-=', '*=', '/=']:
      if self.v == ',':
        self.Eat(',')
        had_comma = True
        comma_needed = False
        trailing_comma = True
      else:
        if comma_needed:
          raise Exception('Comma required before more items in list')
        x = self.Xexpr()
        z.append(x)
        comma_needed = True
        trailing_comma = False
    if allowScalar and len(z) == 1 and not had_comma:
      return z[0]  # Scalar.
    if not allowEmpty and len(z) == 0:
      raise Exception('Empty expression list not allowed')
    return Titems(z, trailing_comma)  # List of items.

  def Csuite(self):
    things = []
    while self.k != 'OUT' and self.k is not None:
      if self.v == ';;':
        self.EatK(';;')
      else:
        if RYE_FLOW:
          num = 1 + sum([int(ch=='\n') for ch in self.program[ : self.i ]])
          what= '"## LINE ## %d ##"' % num
          things.append(Tprint(Tlist([Tlit('S', what)]), False, None))
        cmd = self.Command()
        if cmd:
          if type(cmd) is list:
            for e in cmd:
              things.append(e)
          else:
            things.append(cmd)
    return Tsuite(things)

  def Command(self):
    where = self.i
    line = self.line
    gloss = FirstWord(self.v)
    cmd = self.Command9()
    if cmd:
      if type(cmd) is list:
        for e in cmd:
          # Tag the cmd node with where it was in source.
          e.where = where
          e.gloss = gloss
          e.line = line
      else:
        # Tag the cmd node with where it was in source.
        cmd.where = where
        cmd.gloss = gloss
        cmd.line = line
    return cmd

  def Command9(self):
    if self.v == 'print':
      return self.Cprint(False)
    if self.v == 'say':
      return self.Cprint(True)
    elif self.v == 'if':
      return self.Cif()
    elif self.v == 'switch':
      return self.Cswitch()
    elif self.v == 'while':
      return self.Cwhile()
    elif self.v == 'for':
      return self.Cfor()
    elif self.v == 'return':
      return self.Creturn()
    elif self.v == 'yield':
      return self.Cyield()
    elif self.v == 'break':
      return self.Cbreak()
    elif self.v == 'continue':
      return self.Ccontinue()
    elif self.v == 'raise':
      return self.Craise()
    elif self.v == 'def':
      return self.Cdef('')
    elif self.v == 'class':
      return self.Cclass()
    elif self.v == 'assert':
      return self.Cassert()
    elif self.v == 'must':
      return self.Cassert(True)
    elif self.v == 'import':
      return self.Cimport()
    elif self.v == 'from':
      return self.Cfrom()
    elif self.v == 'with':
      return self.Cwith()
    elif self.v == 'global':
      return self.Cglobal()
    elif self.v == 'try':
      return self.Ctry()
    elif self.v == 'del':
      return self.Cdel()
    elif self.v == 'native':
      return self.Cnative()
    elif self.v == 'pass':
      self.Eat('pass')
      self.EatK(';;')
      return
    elif self.k == 'A' or self.v == '.' or self.v == '(' or self.v == 'go':
      return self.Cother()
    else:
      return self.Cother()

  def Cother(self):
    # lhs (unless not an assignment; then it's the only thing.)
    a = self.Xitems(allowScalar=True, allowEmpty=False)

    op = self.v

    if op in ['+=', '-=', '*=']:
      self.Eat(op)
      binop = op[:-1]  # Remove the '='
      b = self.Xexpr()
      # TODO: this evals lhs twice.
      if binop in ADD_OPS:
        return Tassign(a, Top(a, ADD_OPS[binop], b))
      elif binop in MUL_OPS:
        return Tassign(a, Top(a, MUL_OPS[binop], b))
      else:
        raise Exception('Unknown op, neither ADD_OPS nor MUL_OPS: ' + binop)

    elif op == '=':
      self.Eat(op)
      b = self.Xlistexpr()
      pragma = None
      if self.k == 'P':
        pragma = TrimPragma(self.v)
        self.EatK('P')
      return Tassign(a, b, pragma)

    else:
      return Tassign(Traw('_'), a)

  def Cprint(self, saying):
    # TODO: distinguish trailing ,
    self.Advance()
    if self.v == '>>':
      self.Eat('>>')
      w = self.Xexpr()
      self.Eat(',')
    else:
      w = None
    begin = self.i
    t = self.Xitems(allowScalar=False, allowEmpty=True)
    end = self.i
    self.EatK(';;')
    return Tprint(w, t, saying, self.program[begin : end])

  # New "with defer"
  def Cwith(self):
    self.Eat('with')
    self.Eat('defer')
    call = self.Xexpr()
    if type(call) != Tcall:
      raise Exception('"with defer" statement must contain function or method call')
    body = self.Block()
    return Twithdefer(call, body)

  def Cglobal(self):
    self.Eat('global')
    vars = {}
    while True:
      var = self.Xvar()
      vars[var.name] = var
      if self.v != ',':
        break
      self.Eat(',')
    self.EatK(';;')
    return Tglobal(vars)

  def Cfrom(self):
    self.Eat('from')
    s = ''
    while self.k in ['K', 'A', 'N'] or self.v in ['.', '..', '-', '/']:

      if self.v == 'import':
        break
      s += self.v
      self.Advance()
    #print >> sys.stderr, 'BREAK Cfrom :', self.k, self.v

    if not s:
      raise Exception('No path followed "from"')

    return self.Cimport(fromWhere=s, relative=s.startswith('.'))

  def Cimport(self, fromWhere=None, relative=None):
    self.Eat('import')
    z = []
    while True:
      s = ''
      while self.k in ['K', 'A', 'N'] or self.v in ['.', '-', '/']:
        if self.v == 'as':
          break
        s += self.v
        self.Advance()

      if not s:
        raise Exception('No path followed "import"')

      if not fromWhere:
        fromWhere = 'github.com/strickyak/rye/pye'
        relative = False

      if fromWhere == 'lib':
        fromWhere = 'github.com/strickyak/rye/lib'
        relative = False

      if relative:
        vec = CleanPath(self.cwp, fromWhere if fromWhere else '.', s)
      else:
        vec = CleanPath('/', fromWhere if fromWhere else '.', s)
      alias = vec[-1]

      if self.v == 'as':
        self.Eat('as')
        alias = self.v
        self.EatK('A')

      z.append(Timport(vec, alias, fromWhere=fromWhere))

      # There may be more after a ','
      if self.k == ';;':
        break
      self.Eat(',')

    self.EatK(';;')
    return z

  def Cassert(self, is_must=False):
    fails = False
    i = self.i
    self.Advance()

    if self.v == 'except':
      fails = True
      self.Advance()

    x = self.Xexpr()
    y = None
    j = self.i
    if self.v == ',':
      self.Eat(',')
      y = self.Xlistexpr()
    return Tassert(x, y, self.program[i:j], is_must, fails)

  def Ctry(self):
    exvar = None
    self.Eat('try')
    tr = self.Block()
    self.Eat('except')
    if self.k == 'A': 
      # Currently we only accept Exception.
      self.Eat('Exception')
    if self.v == 'as':
      self.Eat('as')
      exvar = self.Xvar()
    ex = self.Block()
    return Ttry(tr, exvar, ex)

  def Cswitch(self):
    cases = []
    clauses = []
    default_clause = None

    self.Advance()
    a = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')

    while True:
      if self.v == 'case':
        self.Eat('case')
        i, line = self.i, self.line
        x = self.Xexpr()
        x.where, x.line, x.gloss = i, line, 'case'
        cases.append(x)
        c = self.Block()
        clauses.append(c)

      elif self.v == 'default':
        self.Eat('default')
        default_clause = self.Block()

      elif self.k == 'OUT':
        break

      else:
        raise Exception('Expected "case" or "default", but got "%s"' % self.v)

    self.EatK('OUT')
    return Tswitch(a, cases, clauses, default_clause)

  def Cif(self):
    self.Advance()
    t = self.Xexpr()
    yes = self.Block()
    no = None
    if self.v == 'elif':
      no = self.Cif()
      return Tif(t, yes, no)
    if self.v == 'else':
      self.Eat('else')
      no = self.Block()
    return Tif(t, yes, no)

  def Cwhile(self):
    self.Eat('while')
    t = self.Xexpr()
    yes = self.Block()
    return Twhile(t, yes)

  def Cfor(self):
    self.Eat('for')
    x = self.Xvars()
    self.Eat('in')
    t = self.Xlistexpr()
    suite = self.Block()
    return Tfor(x, t, suite)

  def Creturn(self):
    self.Eat('return')
    if self.v == ';;':  # Missing Xitems means None, not [].
      return Treturn(None)
    t = self.Xlistexpr()
    return Treturn([t])

  def Cyield(self):
    self.Eat('yield')
    if self.v == ';;':  # Missing Xitems means None, not [].
      return Tyield(None)
    t = self.Xlistexpr()
    return Tyield([t])

  def Cbreak(self):
    self.Eat('break')
    self.EatK(';;')
    return Tbreak()

  def Ccontinue(self):
    self.Eat('continue')
    self.EatK(';;')
    return Tcontinue()

  def Craise(self):
    self.Eat('raise')
    t = self.Xlistexpr()
    return Traise(t)

  def Cdel(self):
    self.Eat('del')
    listx = self.Xlistexpr()
    return Tdel(listx)

  def Cclass(self):
    self.Eat('class')
    name = self.Pid()
    sup = Tvar('object')

    if self.v == '(':
      self.Advance()
      if self.v == 'native':  # Special keyword.
        self.Advance()
        sup = 'native'
      else:
        sup = self.Xqualname()
      self.Eat(')')

    # Get the body of the class as a suite.
    suite = self.Block()
    # We only allow methods defs or """comments""" in the suite.
    for t in suite.things:
      tt = type(t)
      if tt is Tdef:
        pass  # Good, it's a def.
      elif tt is Tassign:
        if type(t.a) != Traw or t.a.raw != '_':
          raise Exception('Classes many only contain "def" or "pass" or """comments""", but got an assignment statement.')
        if type(t.b) != Tlit or t.b.k != 'S':
          raise Exception('Classes many only contain "def" or "pass" or """comments""", but got a non-string literal.')
        pass  # GOod, it's a string literal (assigned to _).
      else:
        # Oh no, it is something else!
        raise Exception('Classes many only contain "def" or "pass" or """comments""", but got %s' % tt)
    return Tclass(name, sup, suite.things)

  def Cnative(self):
    self.Eat('native')
    self.Eat(':')

    strings = []
    if self.k == 'S':
      strings.append(DecodeStringLit(self.v))
      self.Advance()
      self.EatK(';;')
    else:
      self.EatK(';;')
      self.EatK('IN')
      while self.k == 'S' or self.k == ';;':
        if self.k == 'S':
          strings.append(DecodeStringLit(self.v))
        self.Advance()
      self.EatK('OUT')
    return Tnative(strings)

  def Cdef(self, cls):
    self.Eat('def')
    name = self.Pid()
    self.Eat('(')
    args = []
    dflts = []
    while self.k == 'A':
      arg = self.Pid()
      dflt = None
      if self.v == '=':
        self.Eat('=')
        dflt = self.Xexpr()
      if self.v == ',':
        self.Eat(',')
      args.append(arg)
      dflts.append(dflt)
    star = ''
    starstar = ''
    while self.v in ('*', '**'):
      which = self.v
      self.Advance()
      if which == '*':
        star = self.v
      elif which == '**':
        starstar = self.v
      self.EatK('A')
      if self.v == ',':
        self.Eat(',')
      pass
    self.Eat(')')
    suite = self.Block()
    return Tdef(name, args, dflts, star, starstar, suite)

  def Block(self):
    self.Eat(':')
    if self.v == ';;':
      self.EatK(';;')
      self.EatK('IN')
      suite = self.Csuite()
      self.EatK('OUT')
    else:
      cmd = self.Command()
      suite = Tsuite([cmd])
    return suite

def ParsePragma(s):
  if s == 'i':
    return Yint()
  elif s == 'f':
    return Yfloat()
  elif s == 's':
    return Ystr()
  else:
    raise Exception('Unknown Pragma: %s' % s)

class Y(object):  # Typed values
  def __init__(self):
    pass
class Yint(Y):
    pass
class Yfloat(Y):
    pass
class Ystr(Y):
    pass

# OPERATOR HIERARCHY OF PYTHON
#lambda        Lambda expression
#if else        Conditional expression
#or        Boolean OR
#and        Boolean AND
#not x        Boolean NOT
#in, not in, is, is not, <, <=, >, >=, <>, !=, ==        Comparisons, including membership tests and identity tests
#|        Bitwise OR
#^        Bitwise XOR
#&        Bitwise AND
#<<, >>        Shifts
#+, -        Addition and subtraction
#*, /, //, %        Multiplication, division, remainder [8]
#+x, -x, ~x        Positive, negative, bitwise NOT
#**        Exponentiation [9]
#x[index], x[index:index], x(arguments...), x.attribute        Subscription, slicing, call, attribute reference
#(expressions...), [expressions...], {key: value...}, `expressions...`        Binding or tuple display, list display, dictionary display, string conversion

class StatementWalker(object):

  def Vexpr(self, p):
    pass

  def Vassign(self, p):
    pass

  def Vprint(self, p):
    pass

  def Vimport(self, p):
    pass

  def Vassert(self, p):
    pass

  def Vtry(self, p):
    p.tr.visit(self)
    p.ex.visit(self)

  def Vfor(self, p):
    p.b.visit(self)

  def Vswitch(self, p):
    # (self, a, cases, clauses, default_clause):
    for x in p.cases:
      x.visit(self)
    for x in p.clauses:
      x.visit(self)
    if p.default_clause:
      p.default_clause.visit(self)

  def Vif(self, p):
    p.yes.visit(self)
    if p.no:
      p.no.visit(self)

  def Vwhile(self, p):
    p.yes.visit(self)

  def Vreturn(self, p):
    pass

  def Vbreak(self, p):
    pass

  def Vcontinue(self, p):
    pass

  def Vraise(self, p):
    pass

  def Vdel(self, p):
    pass

  def Vwithdefer(self, p):
    p.body.visit(self)

  def Vglobal(self, p):
    pass

  def Vlit(self, p):
    pass

  def Vraw(self, p):
    pass

  def Vop(self, p):
    pass

  def Vboolop(self, p):
    pass

  def Vcondop(self, p):
    pass

  def Vgetitem(self, p):
    pass

  def Vgetitemslice(self, p):
    pass

  def Vcurlysetter(self, p):
    pass

  def Vtuple(self, p):
    pass

  def Vlist(self, p):
    pass

  def Vdict(self, p):
    pass

  def Vvar(self, p):
    pass

  def Vcall(self, p):
    pass

  def Vgo(self, p):
    pass

  def Vfield(self, p):
    pass

  def Vnative(self, p):
    pass

  def Vdef(self, p):
    pass

  def Vclass(self, p):
    pass

  def Vsuite(self, p):
    for x in p.things:
      x.visit(self)

class YieldAndGlobalFinder(StatementWalker):
  def __init__(self):
    self.yields = False
    self.force_globals = {}

  def Vyield(self, p):
    self.yields = True

  def Vglobal(self, p):
    for v in p.vars:
      self.force_globals[v] = p.vars[v]

class YieldGlobalAndLocalFinder(YieldAndGlobalFinder):
  def __init__(self):
    if rye_rye:
      super()
    else:
      YieldAndGlobalFinder.__init__(self)
    self.assigned = {}

  def Vassign(self, p):
    self.markAssigned(p.a)

  def markAssigned(self, a):
    type_a = type(a)
    if type_a is Titems or type_a is Ttuple:
      for x in a.xx:
        self.markAssigned(x)
    if type_a is Tvar:
      if a.name != '_':
        self.assigned = a.name

def DecodeStringLit(s):
  if s[0] == '`':
    z = s[1:-1]
  else:
    z = s.replace('\n', '\\n')
    try:
      if rye_rye:
        #print >>sys.stderr, 'data.Eval(s) <<<', repr(s)
        z = data.Eval(s)
        #print >>sys.stderr, 'data.Eval(s) >>>', repr(z)
      else:
        z = eval(s)
    except Exception as ex:
      print 'DecodeStringLit:', ex
      raise 'SORRY, rye currently cannot handle this string literal: ' + repr(s) + ' .... BECAUSE:' + str(ex)
  return z
