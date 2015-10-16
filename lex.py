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

TROUBLE_CHAR = re.compile('[^]-~ !#-Z[]')
def GoStringLiteral(s):
  if rye_rye:
    return strconv.QuoteToASCII(s)
  else:
    return '"' + TROUBLE_CHAR.sub((lambda m: '\\x%02x' % ord(m.group(0))), s) + '"'

NONALFA = re.compile('[^A-Za-z0-9]')
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

def SimplifyContinuedLines(tokens):
  # Try to throw no exceptions from this func.
  w = []  # Waiting.
  deep = 0   #  Grouping depth.
  eat_out = 0  # How many OUT marks to ignore.
  for triple in tokens:
    kind, val, pos = triple
    if kind == 'G':
      if val in ['(', '[', '{']:
        deep += 1
      elif val in ['}', ']', ')']:
        deep -= 1

    if eat_out:
      if kind != 'OUT':
        raise Exception('Expected un-indent at position ', pos)
      eat_out -= 1
    elif w or deep:
      #print >> sys.stderr, 'w.append(triple)', triple, w
      w.append(triple)
    else:
      #print >> sys.stderr, 'yield triple', triple
      yield triple

    if w and not deep and val == ';;':
      # Try to throw no exceptions from here.
      #if eat_out:
      #  raise Exception('eat_out:' + eat_out)
      for w_triple in w:
        w_kind, _, _ = w_triple

        if w_kind == 'IN':
          eat_out += 1
        elif w_kind == 'OUT':
          eat_out -= 1
        elif w_kind == ';;':
          pass
        else:
          #print >> sys.stderr, 'yield w_triple from w', w_triple, w
          yield w_triple
      w = []
      yield triple  # The newline.
  pass

class Lex(object):
  def __init__(self, program):
    self.buf = program
    self.i = 0
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
    # pragma looks like a comment, but is considered Black.
    if RE_PRAGMA.match(self.buf[self.i:]):
      return

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
          #self.indents[j+1:] = []  # Trim tail to index j.
          self.indents = self.indents[:j+1]  # Trim tail to index j.

    elif col > self.indents[-1]:
        # indent
        self.Add(('IN', both, i))
        self.indents.append(col)

def TabWidth(s):
  z = 0
  for c in s:
    if c == '\t':
      z = ((z+TAB_WIDTH-1) // TAB_WIDTH) * TAB_WIDTH
    else:
      z += 1
  return z

pass
