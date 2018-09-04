# A sample func or var from each api/go1.txt package,
# so we can prevent errors when import is not used in rye.
#
# Run this:
#     python pkg_sample.py < go/api/go1.txt
# Append the output to tr.py, replacing SAMPLES.

import re
import sys

MATCH_API = re.compile(
    '^pkg ([A-Za-z0-9_/]+), (func|var) ([A-Za-z0-9_]+)'
    ).match

d = {}
for line in sys.stdin:
  m = MATCH_API(line)
  if m:
    pkg, kind, name = m.groups()
    d[pkg] = name

print 'SAMPLES = {'
for k, v in sorted(d.items()):
  print '  "%s": "%s",' % (k, v)
print '  }'
