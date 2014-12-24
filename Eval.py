from go import os, regexp, strconv

RE_WHITE = regexp.MustCompile('^([#][^\n]*[\n]|[ \t\n]+)*')

RE_KEYWORDS = regexp.MustCompile('^\\b(None|null|nil|True|False)\\b')
RE_PUNCT = regexp.MustCompile('^[][(){}:,]')
RE_ALFA = regexp.MustCompile('^[A-Za-z_][A-Za-z0-9_]*')
RE_FLOAT = regexp.MustCompile('^[+-]?[0-9][-+0-9eE]*[.eE][-+0-9eE]*')
RE_INT = regexp.MustCompile('^[+-]?[0-9]+')

RE_STR = regexp.MustCompile('^(["](([^"\\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\\n]|[\\\\].)*)[\'])')
RE_STR2 = regexp.MustCompile('^(?s)[`]([^`]*)[`]')
RE_STR3 = regexp.MustCompile('^(?s)("""(([^\\\\]|[\\\\].)*?)"""|\'\'\'(([^\\\\]|[\\\\].)*?)\'\'\')')

DETECTERS = [ (RE_KEYWORDS, 'K'), (RE_FLOAT, 'F'), (RE_INT, 'N'), (RE_PUNCT, 'G'), (RE_STR3, '3'), (RE_STR, 'S'), (RE_STR2, 'S')]

RequoteRE = regexp.MustCompile('"|[\\\\].')
RequoteDict = { '\\"': '\\"', '\\\\': '\\\\', '\\\'': '\'', '"': '\\"' }
def RequoteSingleToDouble(s):
  return RequoteRE.ReplaceAllStringFunc(s, lambda s: RequoteDict.get(s, s))

def Eval(s):
  ep = EvalParser(s)
  k, x = ep.Token()
  z = ep.Parse(k, x)
  ep.Skip()
  if ep.p != ep.n:
    raise Exception(('Eval: Leftover chars: ', ep.p, ep.n, repr(s[ep.p:ep.n])))
  return z

class EvalParser:
  def __init__(s):
    self.s = s
    self.n = len(s)
    self.p = 0

  def Skip():
    w = RE_WHITE.FindString(self.s[self.p:])
    self.p += len(w)

  def Token():
    self.Skip()
    if self.p == self.n:
      #say 'Token', None, None
      return None, None
    for r, k in DETECTERS:
      m = r.FindString(self.s[self.p:])
      if m:
        self.p += len(m)
        print >>os.Stderr, 'Token', k, m
        #say 'Token', k, m
        return k, m
    raise Exception('eval.EvalParser: Cannot Parse')

  def Parse(k, x):
    if not k:
      raise Exception('eval.EvalParser: Unexpected end of string')
    if k == 'K':
      if x[0] in ['n', 'N']:
        return None
      if x[0] in ['t', 'T']:
        return True
      if x[0] in ['f', 'F']:
        return False
      raise Exception('eval.EvalParser: Weird Keyword token: %s' % x) 
    if k == 'N':
      return strconv.ParseInt(x, 10, 64)
    if k == 'F':
      return strconv.ParseFloat(x, 64)
    if k == '3':
      y = RequoteSingleToDouble(x[3:-3])
      return strconv.Unquote('\"' + y + '\"')
    if k == 'S':
      if x[0] == "'":  # TODO, get the escaping right.
        return strconv.Unquote('"' + RequoteSingleToDouble(x[1:-1]) + '"')
      elif x[0] == "`":
        return strconv.Unquote(x)
      elif x[0] == '"':
        return strconv.Unquote(x)
      else:
        raise Exception('eval.EvalParser: Strange string of type S: %q' % x)

    if x == '[':
      v = []
      while True:
        k2, x2 = self.Token()
        if not k2:
          raise Exception('eval.EvalParser: Unexpected end of string')
        if x2 == ']':
          break
        if x2 == ',':
          continue
        a = self.Parse(k2, x2)
        v.append(a)
      return v
    if x == '(':
      v = []
      while True:
        k2, x2 = self.Token()
        if not k2:
          raise Exception('eval.EvalParser: Unexpected end of string')
        if x2 == ')':
          break
        if x2 == ',':
          continue
        a = self.Parse(k2, x2)
        v.append(a)
      return tuple(v)
    if x == '{':
      d = {}
      while True:
        k2, x2 = self.Token()
        if not k2:
          raise Exception('eval.EvalParser: Unexpected end of string')
        if x2 == '}':
          break
        if x2 == ',':
          continue

        a = self.Parse(k2, x2)

        k2, x2 = self.Token()
        if x2 != ':':
          raise Exception('eval.EvalParser: expected ":" after key')

        k2, x2 = self.Token()
        b = self.Parse(k2, x2)

        d[a] = b
      return d
    raise Exception('eval.EvalParser: Weird token: %q' % x)

assert Eval('True') is True
assert Eval('False') is False
assert Eval('None') is None

assert Eval('12345') == 12345
assert Eval('-12345') == -12345
assert Eval('-1234.5') == -1234.5

assert Eval('[ 123, 4.5, False] ') == [123, 4.5, False]
# TODO: DiCT::EQ  # assert Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }
assert Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }

d = Eval('{ "color": "red", "area": 51 } ')
e = { "color": "red", "area": 51 }

assert sorted([(k, d[k]) for k in d]) == sorted([(k, e[k]) for k in e])

def main(argv):
  for a in argv:
    print a, Eval(a)
