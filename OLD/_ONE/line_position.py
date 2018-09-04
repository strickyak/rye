import sys

n = 1
p = 0
for line in sys.stdin.read().split('\n'):
  print '[%4d:%6d]\t%s' % (n, p, line)
  n += 1
  p += len(line) + 1
