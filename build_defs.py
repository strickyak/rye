# rye.py -- strick
import os
import os.path
import re
import subprocess
import sys
import traceback

import tr

PATH_MATCH = re.compile('(.*)/src/(.*)').match

Stuff = None

def TranslateInternal(filename, wpath, imod):
  global Stuff

  print >>sys.stderr, '*** TranslateInternal', [filename, imod]
  sys.stdout = open(wpath, 'w')
  program = open(filename).read()
  words = tr.Lex(program).tokens
  parser = tr.Parser(program, words, -1)
  try:
    tree = parser.Csuite()
  except:
    print >> sys.stderr, "\n*** ERROR: ", sys.exc_info()[1]
    print >> sys.stderr, "\n*** OCCURRED BEFORE THIS: ", repr(parser.Rest()[:100])
    print >> sys.stderr, "\n*** TRACEBACK:"
    traceback.print_tb(sys.exc_info()[2])
    sys.exit(13)

  gen = tr.CodeGen(None)
  gen.InjectForInternal(Stuff)
  gen.GenModule(imod, "github.com/strickyak/rye", tree, "github.com/strickyak/rye", internal=imod)
  Stuff = gen.ExtractForInternal()
  sys.stdout.close()

  if not os.getenv("RYE_NOFMT"):
    cmd = ['gofmt', '-w', wpath]
    status = Execute(['gofmt', '-w', wpath])
    if status:
      raise Exception('Failure in gofmt: %s' % cmd)

def Execute(cmd):
  return subprocess.call(cmd)

Stuff = dict(), dict(), dict()

if __name__ == '__main__':
  TranslateInternal("def.rye.ry", "gen_rye.go", "rye")
  TranslateInternal("def.go.ry", "gen_go.go", "go")
  TranslateInternal("def.builtins.ry", "gen_builtins.go", "builtins")
