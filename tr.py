import re
import sys

RE_WHITE = re.compile('([ \t\n]*[\n])?([ \t]*)')

RE_KEYWORDS = re.compile(
    'class|def|if|else|while|True|False|None|print|and|or|try|except|raise|return|break|continue|pass')
RE_LONG_OPS = re.compile(
    '[+]=|[*]=|//|<<|>>|<=|>=|[*][*]')
RE_OPS = re.compile('[@~!%^&*+=,|/<>:]')
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

class Parser(object):
  def __init__(self, program, words, p):
    self.program = program
    self.words = words
    self.lcls = {}
    self.glbls = {}
    self.litInts = {}
    self.litStrs = {}
    self.defs = ''
    self.gen = []
    self.k = ''
    self.v = ''
    self.p = p
    self.i = 0
    self.Advance()

  def Advance(self):
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]
    print 'GEN: ', self.gen
    print 'Advance(%d) < %s >< %s > %s' % (self.p, self.k, self.v, repr(self.Rest()))

  def Rest(self):
    return self.program[self.i:]

  def VarGlobal(self, id):
    z = self.glbls.get(id)
    if not z:
        z = 'G_%s' % id
        self.glbls[v] = z
    return z
    
  def VarLocal(self, id):
    z = self.lcls.get(id)
    if not z:
        z = 'var_%s' % id
        self.lcls[v] = z
    return z

  def MkTemp(self):
    z = Serial('tmp')
    self.lcls[z] = z
    return z

  def Gen(self, pattern, *args):
    print "Gen:::", pattern, args
    self.gen.append(pattern % args)

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
      raise Bad('Expected %s, but got %s, at %s' % (v, self.v, repr(self.Rest())))
    self.Advance()

  def EatK(self, k):
    print "EatingK", k
    if self.k != k:
      raise Bad('Expected Kind %s, but got %s, at %s' % (k, self.k, repr(self.Rest())))
    self.Advance()

  def Pid(self):
    if self.k == 'A':
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
      if self.v in self.Lcls:
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
    raise Bad('Expected Xprim, but got %s, at %s' % (self.v, repr(self.Rest())))

  def Xadd(self):
    a = self.Xprim()
    if self.v in '+-':
      op = self.v
      self.Advance()
      b = self.Xprim()
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
        self.Cstmt();

  def Cstmt(self):
    if self.v == 'print':
      self.Cprint()
    elif self.v == 'return':
      self.Creturn()
    elif self.v == 'def':
      self.Cdef()
    else:
      raise Bad("Unknown stmt: %s %s %s", self.k, self.v, repr(self.Rest()))

  def Cprint(self):
    self.Eat('print')
    t = self.Xexpr()
    self.Gen('println( (%s).String() )', t)

  def Creturn(self):
    self.Eat('return')
    t = self.Xexpr()
    self.Gen('return (%s)', t)

  def Cdef(self):
    self.Eat('def')
    name = self.Pid()
    self.Eat('(')
    arg = self.Pid() # Single argument for now!
    self.Eat(')')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    self.glbls[name] = 'G_%s' % name
    p = ProcParser(self)
    p.Run()
    self.EatK('OUT')

  def Run(self):
    t = self.Csuite()

    print '@@ package main'
    print '@@ import . "github.com/strickyak/rye/runt"'
    for k, v in self.litInts.items():
      print '@@ var %s P = MkInt(%s)' % (v, k)
    for k, v in self.litStrs.items():
      print '@@ var %s P = MkStr(%s)' % (v, k)
    for k, v in self.lcls.items():
      print '@@ var %s P /*%s*/' % (v, k)
    for k, v in self.glbls.items():
      print '@@ Global /* %s %s */' % (v, k)
    print '@@ func ryeMain() P {'
    for x in self.gen:
      print '@@  ', x
    print '@@ }'
    print '@@ func main() { ryeMain() }'
    return t
    

class ProcParser(Parser):
  def __init__(self, outer):
    self.outer = outer
    Parser.__init__(self, outer.Rest(), outer.words, outer.p)

pass
