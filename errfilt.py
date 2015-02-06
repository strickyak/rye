from go import bufio
from go import fmt
from go import io/ioutil
from go import os
from go import regexp
from go import strings

# This requires absolute path starting with [/], but that seems to be what golang prints.
GO_LOC = regexp.MustCompile('^(.*?)([/][-A-Za-z0-9._/]+[/]ryemodule[.]go)[:]([0-9]+)(.*)$')

# These are marks in the generated .go file:
PUSH = regexp.MustCompile('// [@] ([0-9]+) [@]')
POP  = regexp.MustCompile('// [$] ([0-9]+) [$]')

NL = 10  # Newline, as a byte value.

Cache = {}

def TrimRight(s):
  return strings.TrimRight(s, '\n\r')

class LineReader:
  def __init__(r):
    .r = bufio.NewReader(r)
  def ReadLine():
    try:
      return TrimRight(.r.ReadString(NL))
    except as ex:
      if str(ex) == 'EOF':
        return None
      else:
        raise ex

def LookupLocation(file, line):
  tup = (file, line)
  z = Cache.get(tup)
  if z:
    return z

  fd = os.Open(file)
  defer fd.Close()

  linemap = []
  stack = []
  r = LineReader(fd)
  while True:
    s = r.ReadLine()
    if s is None:
      break

    push = PUSH.FindStringSubmatch(s)
    pop = POP.FindStringSubmatch(s)
    if push:
      stack.append(int(push[1]))
    if pop:
      stack = stack[:-1]
    linemap.append(stack[-1] if stack else None)

  i = int(line) - 1  # Use zero-based array index.
  offset = linemap[i]
  if offset is None:  # If no offset, return the original question.
    z = "%s:%s" % (file, line)
    Cache[tup] = z
    return z

  ryefile = file[ : 0 - len('/ryemodule.go') ] + ".py"
  ryebody = ioutil.ReadFile(ryefile)

  ryeline = 1
  for ch in ryebody[:offset]:
    if ch == NL:
      ryeline += 1

  z = "%s:%d" % (ryefile, ryeline)
  Cache[tup] = z
  return z

def AlterLine(s):
  m = GO_LOC.FindStringSubmatch(s)
  if m:
    _, hd, file, line, tl = m
    return hd + LookupLocation(file, line) + tl
  else:
    return s

def AlterStream(w, r):
  lr = LineReader(r)
  while True:
    s = lr.ReadLine()
    if s is None:
      break
    fmt.Fprintf(w, '%s\n', AlterLine(s))

def main(argv):
  AlterStream(os.Stdout, os.Stdin)
