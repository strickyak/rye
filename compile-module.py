import os
import os.path
import re
import sys

import tr

if __name__ == '__main__':
  filename = sys.argv[1]
  modname = os.path.basename(filename).split('.')[0]
  dirname = os.path.dirname(filename)
  if not dirname:
    dirname = '.'
  program = open(filename).read()

  words = tr.Lex(program).tokens
  tree = tr.Parser(program, words, -1).Csuite()

  try:
    os.makedirs('%s/__%s' % (dirname, modname))
  except OSError:
    pass
  try:
    os.makedirs('%s/__%s/__main' % (dirname, modname))
  except OSError:
    pass

  sys.stdout = open('%s/__%s/%s.go' % (dirname, modname, modname), "w")
  main = open('%s/__%s/__main/main.go' % (dirname, modname), "w")
  tr.CodeGen(None).GenModule('%s/__%s' % (dirname, modname), filename, tree, main)
