import re
import sys

# These are marks in the generated .go file:
MATCH_PUSH = re.compile('\\s*// [@] [0-9]+ [@] ([0-9]+)').match
MATCH_POP  = re.compile('\\s*// [$] [0-9]+ [$] ([0-9]+)').search

def ScanFileForLinemap(filename):
  linemap = [0]  # Initial 0 for nonexistant "line 0".
  stack = []
  fd = open(filename)
  try:
    for line in fd.read().split('\n'):
      push = MATCH_PUSH(line)
      pop = MATCH_POP(line)

      if push:
        stack.append(int(push.group(1)))
      if pop:
        stack = stack[:-1]

      linemap.append(stack[-1] if stack else 0)
      #print >> sys.stderr, 'linemap: %d %d %d # %d # %s' % ( int(bool(push)), int(bool(pop)), len(linemap)-1, linemap[-1], line )
  finally:
    fd.close()
  return linemap

def main(argv):
  for a in argv:
    lm = ScanFileForLinemap(a)
    print a, len(lm), repr(lm)
    print [(i, lm[i]) for i in range(len(lm))]

if __name__ == '__main__':
  main(sys.argv[1:])
