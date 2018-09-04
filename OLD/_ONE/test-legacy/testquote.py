import re

Cannot = "'([\"](([^\"\\\\\\\\\\\\n]|[\\\\\\\\].)*)[\"]|[\\'](([^\\'\\\\\\\\\\\\n]|[\\\\\\\\].)*)[\\'])'"
print Cannot
print '(["](([^"\\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\\n]|[\\\\].)*)[\'])'
print '''<<<(["](([^"\\\n]|[\\].)*)["]|['](([^'\\\n]|[\\].)*)['])>>>'''

S1 = '(["](([^"\\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\\n]|[\\\\].)*)[\'])'
R1 = re.compile(S1)

S2 = '(?s)[`]([^`]*)[`]'
R2 = re.compile(S2)

S3 = '(?s)("""(([^\\\\]|[\\\\].)*?)"""|\'\'\'(([^\\\\]|[\\\\].)*?)\'\'\')'
R3 = re.compile(S3)

print '<<<', S1, '>>>'
print '<<<', S2, '>>>'
print '<<<', S3, '>>>'

x1 = repr(S1)
x2 = repr(S2)
x3 = repr(S3)

rnum=0
for r in [R1, R2, R3]:
  rnum += 1
  snum = 0
  for s in [x1, x2, x3]:
    snum += 1
    m = r.match(s)
    g0 = m.group(0) if m else None
    print (rnum, snum), 'len(s)=', len(s), 'len(g0)', len(g0) if g0 else None, g0
    print

#say str(open('testquote.py').read())
#say str(open('testquote.py').read()).split('\n')

for line in str(open('testquote.py').read()).split('\n'):
  print '<<<', line, '<<<'
  m = R1.search(line)
  if m:
    print '======', m.group(0)
