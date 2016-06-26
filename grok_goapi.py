# python grok_goapi.py < go1.txt > goapi.py

import re
import sys

MATCH_SIMPLE_FUNC = re.compile(
    'pkg ([a-z0-9/]+), func ()([A-Za-z0-9_]+)[(]([^()]*)[)] ?[(]?([^()]*)[)]?$'
    ).match
MATCH_SIMPLE_METH = re.compile(
    'pkg ([a-z0-9/]+), method [(]([^()]*)[)] ([A-Za-z0-9_]+)[(]([^()]*)[)] ?[(]?([^()]*)[)]?$'
    ).match
MATCH_TEXT_METH = re.compile(
    'pkg ([a-z0-9/]+), method [(][^()]*[)] (.*)$'
    ).match

FAVS = set([
    'bool', 'string', 'error',
    'int', 'int8', 'int16', 'int32', 'int64', 'uint', 'uint8', 'uint16', 'uint32', 'uint64', 'uintptr',
    'float32', 'float64',
    #'[]uint8', '[]string',
    ])

MATCH_PUBLIC_TYPENAME = re.compile(
    '[A-Z][A-Za-z0-9_]*$'
    ).match

MATCH_SCOPED_PUBLIC_TYPENAME = re.compile(
    '(?:([a-z][a-z0-9_]*)[.])?([A-Z][A-Za-z0-9_]*)$'
    ).match

FAV_PACKAGES = set(['bytes', 'io', 'os', 'reflect', 'time', 'fmt'])

def GrokType(pkg, s):
  s = s.strip()
  if s in FAVS:
    return s
  if s.startswith('*'):
    t = GrokType(pkg, s[1:])
    return ('*%s' % t) if t else None
  if s.startswith('[]'):
    t = GrokType(pkg, s[2:])
    return ('[]%s' % t) if t else None
  if s.startswith('map[string]'):
    t = GrokType(pkg, s[11:])
    return ('map[string]%s' % t) if t else None
  if MATCH_PUBLIC_TYPENAME(s):
    #return '%s.%s' % (pkg, s)
    if pkg:
      if pkg in FAV_PACKAGES:
        # Go ahead an annotate with package name if it is a FAV.
        return '%s.%s' % (pkg, s)
      else:
        return s  # Same package: use simple name.
    else:
      return None
  m = MATCH_SCOPED_PUBLIC_TYPENAME(s)
  if m:
    scope, name = m.groups()
    if scope in FAV_PACKAGES:
      return '%s.%s' % (scope, name)
  return None


METHS = {}
NONALFA = re.compile('[^A-Za-z0-9]')

def FinishMeths():
  print 'QMeths = {'
  for name, vec in sorted(METHS.items()):
    #ttt, rrr = vec[0]
    tt, rr = vec[0]
    consistent = True
    for t2, r2, in vec:
      if t2 != tt or r2 != rr:
        consistent = False
    if consistent:
      text = '%s (%s) (%s)' % (
          name,
          ', '.join(t for t in tt),
          ', '.join(r for r in rr),
          )
      signature = 'signature_%s__%s_return_%s' % (
          name,
          '_also_'.join([NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s) for s in tt]),
          '_also_'.join([NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s) for s in rr]),
          )
      print '  "%s": QMeth(%s, %s, %s, %s, %s),' % (name, repr(name), repr(tt), repr(rr), repr(signature), repr(text))
    else:
      print '  ## INCONSISTENT ## "/%s": %s' % (name, repr(vec))
  print '  }'

def GrokApiFunc(pkg, receiver, name, takes, rets):
  if receiver and pkg not in FAV_PACKAGES:
    # For methods, only provide pkg if it is a fav.
    # That way local names in non-Fav packages like `Template` that have collisions are inhibited.
    pkg = ''
  tt, rr = [], []
  for t in takes.split(','):
    x = GrokType(pkg, t)
    if not x: return
    tt.append(x)
  for r in rets.split(','):
    x = GrokType(pkg, r)
    if not x: return
    rr.append(x)
  # print '>>>>>', pkg, name, tt, rr
  ttt = ', '.join([repr(t) for t in tt])
  rrr = ', '.join([repr(r) for r in rr])
  if receiver:
    v = METHS.get(name)
    if v is None:
      v = []
      METHS[name] = v

    v.append((tt, rr))
  else:
    print '  "%s.%s": QFunc([%s], [%s]),' % (pkg, name, ttt, rrr)


def GrokApiFile(r):
  print '''# MACHINE GENERATED CODE.  See grok_goapi.py.

class QFunc(object):
  def __init__(self, takes, rets):
    self.takes = takes
    self.rets = rets

class QMeth(object):
  def __init__(self, name, takes, rets, signature, text):
    self.name = name
    self.takes = takes
    self.rets = rets
    self.signature = signature
    self.text = text

QFuncs = {
'''

  for line in r:
    line = line.strip()
    f = MATCH_SIMPLE_FUNC(line)
    m = MATCH_SIMPLE_METH(line)
    if f:
      GrokApiFunc(*tuple(f.groups()))
    if m:
      GrokApiFunc(*tuple(m.groups()))
    else:
      print '  # NOT: %s' % line
      pass
  print '''
  }
'''
  FinishMeths()

if __name__ == '__main__':
  GrokApiFile(sys.stdin)
