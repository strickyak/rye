"""
rye/pye/re is Emulation for a small subset of Python's "re" class.

Use re.compile(s) to compile a regular expression s.
It results in an instance of re_compiled.

On that, you can use match() or search() or sub().
match() and search() result in a re_matched object
which supports two methods, group(i) and groups().
"""
from go import bytes, regexp

def compile(r:str):
  "Compile the given string as a regular expression."
  return re_compiled(r)

class re_compiled:
  "A compiled regular expression."

  def __init__(r):
    "(Internal)"
    .rmatch = regexp.MustCompile('^(?:%s)' % r)  # Non-capturing outer group.
    .rsearch = regexp.MustCompile(r)

  def match(s :str|byt):
    "Match the string, anchored at the beginning."
    m = .rmatch.FindStringSubmatch(str(s))
    return re_matched(m) if m else None
  
  def search(s :str|byt):
    "Match the string, not anchored."
    m = .rsearch.FindStringSubmatch(str(s))
    return re_matched(m) if m else None

  def sub(replacement, s):
    "Substitute the replacement for all occurances of the regular expression in the string."
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
  "The result of match() or search() on a compiled regular expression."
  def __init__(m):
    "(Internal)"
    .m = m

  def group(i):
    "Return the ith group matched."
    return .m[i]

  def groups():
    "Return all groups matched."
    return tuple([str(e) for e in .m[1:]])
