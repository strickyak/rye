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

class Coder(object):
  def __init__(self, outer, cls, name, args):
    self.outer = outer
    self.cls = cls
    self.name = name
    self.args = args
    self.lcls = {}          # name -> goName 
    self.defs = []          # start of function.
    self.gen = []           # body of function.
    for a in args:
      self.lcls[a] = 'a_%s' % a

  def Finish(self):
    print '@@'
    hdr = '@@ func G_%s(' % self.name
    for arg in self.args:
      hdr += 'a_%s P, ' % arg
    hdr += ') P {'
    print hdr

    for k, v in self.lcls.items():
      if v[0] != 'a':
        print '@@   var %s P /*%s*/' % (v, k)

    for x in self.defs:
      print '@@   %s' % x
      
    for x in self.gen:
      print '@@   %s' % x
      
    print '@@   return MkStr("") // Extra Return'
    print '@@ }'
    print '@@'
    

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
    self.coder = Coder(None, None, 'Rye_Module', ['__name__'])
    self.Advance()

  def Bad(self, format, *args):
    msg = format % args
    self.Info(msg)
    raise Exception(msg)

  def Info(self, msg):
    sys.stdout.flush()
    print >> sys.stderr, 120 * '#'
    print >> sys.stderr, "   msg = ", msg
    print >> sys.stderr, "   k =", repr(self.k)
    print >> sys.stderr, "   v =", repr(self.v)
    print >> sys.stderr, "   rest =", repr(self.Rest())

  def Advance(self):
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]
    print '%s GEN: %s' % ( self, self.coder.gen )
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
    print "Gen:::", pattern, args
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
    print "Eating", v
    if self.v != v:
      raise self.Bad('Expected %s, but got %s, at %s', v, self.v, repr(self.Rest()))
    self.Advance()

  def EatK(self, k):
    print "EatingK", k
    if self.k != k:
      raise self.Bad('Expected Kind %s, but got %s, at %s', k, self.k, repr(self.Rest()))
    self.Advance()

  def Pid(self):
    if self.k != 'A':
      raise self.Bad("Pid expected kind A, but got %s %q", self.k, repr(self.Rest()))
    z = self.v
    self.Advance()
    return z

  def Xprim(self):
    if self.k == 'N':
      z = self.LitInt(self.v)
      self.Advance()
      return z
    if self.k == 'S':
      z = self.LitStr(self.v)
      self.Advance()
      return z
    if self.k == 'A':
      if self.v in self.coder.lcls:
        z = self.VarLocal(self.v)
        self.Advance()
        return z
      else:
        z = 'G_%s' % self.v
        self.Advance()
        return z
    if self.v == '(':
        self.Advance()
        z = self.Xparen(self.v)
        self.Eat(')')
        return z
    raise self.Bad('Expected Xprim, but got %s, at %s' % (self.v, repr(self.Rest())))

  def Xcall(self):
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
        t = self.MkTemp()
        self.Gen('%s = %s ( %s )', t, a, ', '.join(args))
        a = t

      elif self.v == '.':
        self.Eat('.')
	field = self.v
        self.EatK('A')
        t = self.MkTemp()
        self.Gen('%s = %s.Field(%s)', t, a, field)
	a = t

      else:
        break
    return a

  def Xadd(self):
    a = self.Xcall()
    if self.v in '+-':
      op = self.v
      self.Advance()
      b = self.Xcall()
      t = self.MkTemp()
      fns = {'+': 'Add', '-': 'Sub'}
      self.Gen('%s = %s.%s(%s)', t, a, fns[op], b)
      a = t
    return a

  def Xexpr(self):
    return self.Xadd()

  def Csuite(self):
    while self.k != 'OUT' and self.k is not None:
      print "Csuite", self.k, self.v
      if self.v == ';;':
        self.EatK(';;')
      else:
        self.Command();

  def Command(self):
    if self.v == 'print':
      self.Cprint()
    elif self.v == 'return':
      self.Creturn()
    elif self.v == 'def':
      self.Cdef('')
    elif self.v == 'class':
      self.Cclass()
    elif self.v == 'pass':
      self.Eat('pass')
    else:
      raise self.Bad("Unknown stmt: %s %s %s", self.k, self.v, repr(self.Rest()))

  def Cprint(self):
    self.Eat('print')
    t = self.Xexpr()
    self.Gen('println( (%s).String() )', t)

  def Creturn(self):
    self.Eat('return')
    t = self.Xexpr()
    self.Gen('return (%s)', t)

  def Cclass(self):
    self.Eat('class')
    name = self.Pid()
    # No superclasses yet
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')

    while self.k != 'OUT':
      if self.v == 'def':
        self.Cdef(name)
      elif self.v == 'pass':
        self.Eat('pass')
      elif self.k == ';;':
        self.EatK(';;')
      else:
        raise self.Bad('Classes may only contain "def" or "pass" commands.')
    self.EatK('OUT')

  def Cdef(self, cls):
    self.Eat('def')
    name = self.Pid()
    print "Cdef -------- %q :: name", cls, name
    self.Eat('(')
    args = []
    while self.k == 'A':
      arg = self.Pid() # Single argument for now!
      if self.v == ',':
        self.Eat(',')
      args.append(arg)
    print "Cdef -------- %q :: args", cls, args
    self.Eat(')')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    self.glbls[name] = 'func'

    save = self.coder
    self.coder = Coder(save, cls, name, args)
    self.Csuite()
    self.coder.Finish()
    self.coder = save

    self.EatK('OUT')

  def Run(self):
    #try:
    print '@@ package main'
    print '@@ import . "github.com/strickyak/rye/runt"'
    t = self.Csuite()
    self.coder.Finish()
    #    self.Info()
    #except:
    #    print sys.exc_info()
    #    Bad("RETHROW")

    for k, v in self.litInts.items():
      print '@@ var %s P = MkInt(%s)' % (v, k)

    for k, v in self.litStrs.items():
      print '@@ var %s P = MkStr(%s)' % (v, k)

    for k, v in self.glbls.items():
      if v != 'func':
        print '@@ var %s P // %s' % (v, k)

    print '@@ func main() { G_Rye_Module(MkStr("__main__")) }'
    return t

pass
