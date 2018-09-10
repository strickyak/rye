import md5
import os
import re
import sys

rye_rye = False
if rye_rye:
  from rye_lib import data
  from go import strconv

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([ \t\n]*[#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')

RE_KEYWORDS = re.compile(
    '\\b(del|say|from|class|def|native|if|elif|else|while|True|False|None|print|and|or|try|except|raise|yield|return|break|continue|pass|as|go|defer|with|global|assert|must|lambda|switch|finally)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[-]=|[*]=|/=|//|<<|>>>|>>|==|!=|<=|>=|[*][*]|[.][.]|::')
RE_OPS = re.compile('[-.@?~!%^&*+=,|/<>:]')
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

TAB_WIDTH = 8

DETECTERS = [
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
  if len(s) < 30:
    # Nicer name for shorter things.
    return NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
  else:
    # Hex Hash for longer things.
    return md5.new(s).hexdigest()

def AddWhereInProgram(err, pos, filename=None, program=None):
  if filename:
    fd = open(filename)
    try:
      program = fd.read()
    finally:
      fd.close()

  if program:
    numLines = 1
    numBytes = 0
    for line in program.split('\n'):
      ll = len(line)
      if numBytes + ll > pos:
        col = pos - numBytes + 1
        col = 1 if col < 1 else col
        where = '%s:%d:%d [pos=%d]' % ((filename if filename else ''), numLines, col, pos)
        picture = ((col-1) * '-') + '^'
        return '%s\n  %s\n  >%s\n  >%s\n' % (err, where, line, picture)
      numBytes += ll + 1 # 1 for the newline.
      numLines += 1

  return '%s\n\tPOSITION:%d' % (err, pos)


def SimplifyContinuedLines(tokens, filename=None, program=None):
  z = []
  lookOut = 0
  startedAt = 0
  waiting = []
  deep = 0   #  Grouping depth.
  eat_out = 0  # How many OUT marks to ignore.
  for triple in tokens:
    kind, val, pos = triple
    if kind == 'G':
      if val in ['(', '[', '{']:
        deep += 1
        if deep == 1:
          startedAt = pos
      elif val in ['}', ']', ')']:
        deep -= 1

    if deep:
      if kind == 'IN':
        lookOut += 1
      elif kind == 'OUT':
        lookOut -= 1
        if lookOut < 1:
          raise Exception(AddWhereInProgram(
              "Un-indented while parens/brackets/braces are open: deep=%d lookOut=%d triple=%s" % (deep, lookOut, triple),
              pos,
              filename=filename))

    if eat_out:
      if kind != 'OUT':
        raise Exception(AddWhereInProgram(
            'Expected un-indent at position %d' % pos,
            pos,
            filename=filename))
      eat_out -= 1
    elif waiting or deep:
      waiting.append(triple)
    else:
      z.append(triple)

    if waiting and not deep and val == ';;':
      for w_triple in waiting:
        w_kind, _, _ = w_triple

        if w_kind == 'IN':
          eat_out += 1
        elif w_kind == 'OUT':
          eat_out -= 1
        elif w_kind == ';;':
          pass
        else:
          z.append(w_triple)

      waiting = []
      z.append(triple)
  return z

class Lex(object):
  def __init__(self, program, filename=None):
    self.buf = program
    self.filename = filename
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
    raise Exception(AddWhereInProgram('Cannot parse token: %s' % repr(rest[0]), self.i, filename=self.filename))

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
          raise Exception(AddWhereInProgram('Cannot un-indent: New column is %d; previous columns are %s' % (col, repr(self.indents)), i+col, filename=self.filename))
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
      z = int((z+TAB_WIDTH-1) / TAB_WIDTH) * TAB_WIDTH
    else:
      z += 1
  return z

pass
