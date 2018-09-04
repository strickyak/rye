import re
import sys

# These are marks in the generated .go file:
MATCH_PUSH = re.compile('\\s*// [@] [0-9]+ [@] ([0-9]+) @ *(.*)').match
MATCH_POP  = re.compile('\\s*// [$] [0-9]+ [$] ([0-9]+)').search

def ScanFileForLinemap(filename, srcFilename=None):
  srcLines = None
  if srcFilename:
    srcLines = open(srcFilename).read().split('\n')
  srcDict = {}
  whatDict = {}

  linemap = [0]  # Initial 0 for nonexistant "line 0".
  stack = []
  fd = open(filename)
  try:
    for line in fd.read().split('\n'):
      push = MATCH_PUSH(line)
      pop = MATCH_POP(line)

      if push:
        n = int(push.group(1))
        stack.append(n)
        whatDict[n] = push.group(2)
      if pop:
        stack = stack[:-1]

      x = stack[-1] if stack else 0
      linemap.append(x)
      if srcLines and x:
        srcDict[x] = srcLines[x-1]
  finally:
    fd.close()
  return linemap, srcDict, whatDict

def main(argv):
  for a in argv:
    lm = ScanFileForLinemap(a)
    print a, len(lm), repr(lm)
    print [(i, lm[i]) for i in range(len(lm))]

if __name__ == '__main__':
  main(sys.argv[1:])
