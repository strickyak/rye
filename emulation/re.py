""" `rye/pye/re` is Emulation for a small subset of Python's "re" class.

Use `re.compile(s)` to compile a regular expression s.
It results in an instance of `RyeRegexpCompiled`.

On that, you can use `match()` or `search()` or `sub()`.

`match()` and `search()` result in a `RyeRegexpMatched` object
which supports two methods, `group(i :int)` and `groups()`.
"""

from go import bytes, regexp
import sys

def compile(r :str) ->RyeRegexpCompiled :
  "Compile the given string as a regular expression."
  return RyeRegexpCompiled(r)

class RyeRegexpCompiled:
  "Internal class for a compiled regular expression."

  def __init__(r :str):
    "Internal: r is the regular expression in godoc regexp/syntax"
    .rmatch = regexp.MustCompile('^(?:%s)' % r)  # Non-capturing outer group.
    .rsearch = regexp.MustCompile(r)

  def match(s :str) -> RyeRegexpMatched? :
    "Match the string, anchored at the beginning."
    m = .rmatch.FindStringSubmatch(str(s))
    return RyeRegexpMatched(m) if m else None
  
  def search(s :str) -> RyeRegexpMatched? :
    "Match the string, not anchored."
    m = .rsearch.FindStringSubmatch(str(s))
    return RyeRegexpMatched(m) if m else None

  def sub(replacement :str, s :str) ->str :
    "Substitute the replacement for all occurances of the regular expression in the string."
    if callable(replacement):
      m = .rsearch.FindAllStringIndex(s, -1)
      if m:
        p = 0
        z = go_new(bytes.Buffer)
        for m2 in m:
          say p, m2, m
          i, j = m2[:2]
          z.WriteString(s[p:i])
          m3 = [str(s[m2[k+k]:m2[k+k+1]]) for k in range(int(len(m2)/2))]
          r = str(replacement(RyeRegexpMatched(m3)))
          z.WriteString(r)
          say p, i, j, r, str(z)
          p = j
        z.WriteString(s[p:])
        return str(z)
      else:
        return s
    else:
      return .rsearch.ReplaceAllString(s, replacement)

class RyeRegexpMatched:
  "Internal class for a positive match() or search() result."

  def __init__(m):
    "Internal: m is the []string result from FindStringSubmatch()."
    .m = m

  def group(i :int) ->str :
    "Return the ith group matched."
    return .m[i]

  def groups() ->tuple :
    "Return all groups matched."
    return tuple([str(e) for e in .m[1:]])
