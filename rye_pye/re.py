"""
rye/pye/re is Emulation for a small subset of Python's "re" class.
"""
from go import bytes, regexp

def compile(r):
  return PYE_RE(r)
class PYE_RE:
  def __init__(r):
    .rmatch = regexp.MustCompile('^(?:%s)' % r)  # Non-capturing outer group.
    .rsearch = regexp.MustCompile(r)

  def match(s :str|byt):
    m = .rmatch.FindStringSubmatch(str(s))
    return re_matched(m) if m else None
  
  def search(s :str|byt):
    m = .rsearch.FindStringSubmatch(str(s))
    return re_matched(m) if m else None

  def sub(replacement, s):
    # TODO: we need builtin callable() to detect if replacemnt is function.
    if callable(replacement):
      m = .rsearch.FindAllStringIndex(s, -1)
      if m:
        #say go_value(m), m, s, .rsearch, replacement
        p = 0
        z = go_new(bytes.Buffer)
        for m2 in m:
          say p, m2, m
          i, j = m2[:2]
          z.WriteString(s[p:i])
          m3 = [str(s[m2[k+k]:m2[k+k+1]]) for k in range(len(m2)//2)]
          r = str(replacement(re_matched(m3)))
          z.WriteString(r)
          say p, i, j, r, str(z)
          p = j
        z.WriteString(s[p:])
        return str(z)
      else:
        return s
    else:
      return .rsearch.ReplaceAllString(s, replacement)

class re_matched:
  def __init__(m):
    .m = m

  def group(i):
    return .m[i]

  def groups():
    return .m[1:]
