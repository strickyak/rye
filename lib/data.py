from go import os, regexp, strconv

RE_WHITE = regexp.MustCompile('^([#][^\n]*[\n]|[ \t\n]+)*')

RE_KEYWORDS = regexp.MustCompile('^\\b(None|null|nil|True|False|true|false|byt)\\b')
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
    tail = self.s[self.p:]
    for r, k in DETECTERS:
      m = r.FindString(tail)
      if m:
        self.p += len(m)
        #say 'Token', k, m
        return k, m
    raise Exception('eval.EvalParser: Cannot Parse: %q' % tail)

  def Parse(k, x):
    if not k:
      raise Exception('eval.EvalParser: Unexpected end of string')
    switch k:
      case 'K':
        switch x[0].lower():
          case 'n':
            return None
          case 't':
            return True
          case 'f':
            return False
        if x == 'byt':
          k2, x2 = self.Token()
          if not k2:
            raise Exception('eval.EvalParser: Unexpected end of string')
          must x2 == '(', 'EvalParser wanted ( after byt'
          k2, x2 = self.Token()
          s2 = .Parse(k2, x2)
          must type(s2) == str, 'EvalParser wanted str after byt('
          k2, x2 = self.Token()
          must x2 == ')', 'EvalParser wanted ) after byt(str'
          return byt(s2)
        raise Exception('eval.EvalParser: Weird Keyword token: %s' % x) 
      case 'N':
        return strconv.ParseInt(x, 10, 64)
      case 'F':
        return strconv.ParseFloat(x, 64)
      case '3':
        return Unquote(x[3:-3])
        #y = RequoteSingleToDouble(x[3:-3])
        #return strconv.Unquote('\"' + y + '\"')
      case 'S':
        if x[0] == "'":
          return strconv.Unquote('"' + RequoteSingleToDouble(x[1:-1]) + '"')
        elif x[0] == "`":
          return strconv.Unquote(x)
        elif x[0] == '"':
          return strconv.Unquote(x)
        else:
          raise Exception('eval.EvalParser: Strange string of type S: %q' % x)

      case 'G':
        switch x:
          case '[':
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
          case '(':
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
          case '{':
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

def Octval(c):
  """Integer value of single octal char."""
  return c - '0'

def Octval3(c, d, e):
  """Integer value of three octal chars."""
  return 64*Octval(c) + 8*Octval(c) + Octval(c)

def Unquote(a):
  """Decode backslash escapes in a."""
  z = byt('')
  # Use byt for transfering bytes from a to z,
  # because octal escapes are for bytes, not for runes.
  # If we used str, we might change bytes to multibyte runes. 
  while a:
    i = a.find('\\')
    if i >= 0:
      z += byt(a[:i])  # Before the quote
      a = a[i:]
      if len(a) < 2:
        raise Exception('Backslash followed by end of string')
      switch a[1]:
        case '\\':
          z += byt('\\')
        case 'n':
          z += byt('\n')
        case 'r':
          z += byt('\r')
        case 'b':
          z += byt('\b')
        case 't':
          z += byt('\t')
        case 'a':
          z += byt('\a')
        case 'v':
          z += byt('\v')
        default:
          if '0' <= a[1] and a[1] <= '7':
            if len(a) >= 4 and '0' <= a[2] and a[2] <= '7' and '0' <= a[3] and a[3] <= '7':
              z += byt(chr(Octval3(a[1], a[2], a[3])))
              a = a[2:]  # Consume two extra octal digits.
            else:
              raise Exception('Bad Octal Chars or Not Enough Chars')
          else:
            raise Exception('Unknown Backslash Escape')
      a = a[2:]  # Consume backlash and following char.
    else:
      return str(z + byt(a))
      a = ''
  return z

def main(args):
  for a in args:
    print a, Eval(a)
