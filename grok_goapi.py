# python grok_goapi.py < /opt/yak/go/api/go1.txt > goapi.py

import re
import sys

MATCH_SIMPLE_FUNC = re.compile(
  '^pkg ([a-z0-9/]+), func ([A-Za-z0-9_]+)[(]([^()]*)[)] ?[(]?([^()]*)[)]?$'
  ).match

FAVS = set([
    'bool', 'string',
    'int', 'int8', 'int16', 'int32', 'int64', 'uint', 'uint8', 'uint16', 'uint32', 'uint64', 'uintptr',
    'float32', 'float64',
    '[]uint8', '[]string',
    ])

def GrokApiFunc(pkg, name, takes, rets):
  tt, rr = [], []
  for t in takes.split(','):
    t = t.strip()
    if t in FAVS:
      tt.append(t)
    else: return
  for r in rets.split(','):
    r = r.strip()
    if r in FAVS:
      rr.append(r)
    else: return
  # print '>>>>>', pkg, name, tt, rr
  ttt = ', '.join([repr(t) for t in tt])
  rrr = ', '.join([repr(r) for r in rr])
  print '  "%s.%s": QFunc([%s], [%s]),' % (pkg, name, ttt, rrr)


def GrokApiFile(r):
  print '''# MACHINE GENERATED CODE.  See grok_goapi.py.

class QFunc(object):
  def __init__(self, takes, rets):
    self.takes = takes
    self.rets = rets

QFUNCS = {
'''

  for line in r:
    line = line.strip()
    m = MATCH_SIMPLE_FUNC(line)
    if m:
      print '# PACKAGE %s NAME %s TAKES %s RETURNS %s' % tuple(m.groups())
      GrokApiFunc(*tuple(m.groups()))
    else:
      print '# NOT: %s' % line
  print '''
  }
'''

if __name__ == '__main__':
  GrokApiFile(sys.stdin)
