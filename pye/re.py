"""
rye/pye/re is Emulation for a small subset of Python's "re" class.
"""
from go import regexp

class compile:
  def __init__(r):
    .rmatch = regexp.MustCompile('^(%s)' % r)
    .rsearch = regexp.MustCompile(r)

  def match(s):
    m = .rmatch.FindStringSubmatch(s)
    return re_matched(m[1:]) if m else None
  
  def search(s):
    m = .rsearch.FindStringSubmatch(s)
    return re_matched(m) if m else None

class re_matched:
  def __init__(m):
    .m = m

  def group(i):
    return .m[i]

  def groups():
    return .m[1:]
