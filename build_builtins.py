# rye.py -- strick
import os
import os.path
import re
import subprocess
import sys
import traceback

import lex
import parse
import codegen

PATH_MATCH = re.compile('(.*)/src/(.*)').match

Stuff = None

def TranslateInternal(filename, wpath, imod):
  global Stuff

  print >>sys.stderr, '<-> TranslateInternal', [filename, imod]
  sys.stdout = open(wpath, 'w')
  program = open(filename).read()
  words = lex.Lex(program).tokens
  parser = parse.Parser(program, words, -1, 'github.com/strickyak/rye')
  try:
    tree = parser.Csuite()
  except:
    print >> sys.stderr, "\n*** ERROR: ", sys.exc_info()[1]
    print >> sys.stderr, "\n*** OCCURRED BEFORE THIS: ", repr(parser.Rest()[:100])
    print >> sys.stderr, "\n*** TRACEBACK:"
    traceback.print_tb(sys.exc_info()[2])
    sys.exit(13)

  gen = codegen.CodeGen()
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

Stuff = dict(), dict(), dict(), dict()

if __name__ == '__main__':
  TranslateInternal("builtins.ry", "gen_builtins.go", "builtins")
