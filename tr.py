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
    raise Exception("Cannot parse (at %d): %s", self.i, repr(rest))

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

    self.Add((';', ';', i))
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
          raise Exception('Cannot un-indent: New column is %d; previous columns are %s', col, repr(self.indents))
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
  def __init__(self, program, words):
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
    self.p = -1
    self.i = 0
    self.Advance()

  def Advance(self):
    self.p += 1
    self.k, self.v, self.i = self.words[self.p]

  def Rest(self):
    return self.program[self.i:]

  def VarGlobal(self, id):
    z = self.glbls.get(id)
    if not z:
        z = 'var_%s' % id
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
    self.gen.append(pattern % args)

  def LitInt(self, v):
    z = self.lcls.get(v)
    if not z:
        z = Serial('lit_int')
        self.lcls[v] = z
    return z

  def LitStr(self, v):
    z = self.lcls.get(v)
    if not z:
        z = Serial('lit_str')
        self.lcls[v] = z
    return z

  def Eat(self, v):
    if s.v != v:
      raise Exception('Expected %q, but got %q, at %q' % (v, s.v, s.rest()))

  def Xprim(self):
    if s.k == 'N':
      z = self.LitInt(s.v)
      self.advance()
      return z
    if s.k == 'S':
      z = self.LitStr(s.v)
      self.advance()
      return z
    if s.k == 'A':
      if s.v in self.Lcls:
        z = self.VarLocal(s.v)
        self.advance()
        return z
      else:
        z = 'G_%s' % s.v
        self.advance()
        return z
    if s.v == '(':
        self.advance()
        z = self.Xparen(s.v)
        self.Eat(')')
        return z
    raise Exception('Expected Xprim, but got %q, at %q' % (v, s.v, s.rest()))

  def Xadd(self):
    a = self.Xprim()
    if s.v in '+-':
      b = self.Xprim()
      t = self.MkTemp()
      self.Gen('%s = ((%s)+(%s))', t, a, b)

pass
