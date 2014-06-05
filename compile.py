import os
import sys
import traceback

import tr


if __name__ == '__main__':
  #  if len(sys.argv) > 2
  program = open(sys.argv[1]).read()

  words = tr.Lex(program).tokens
  tree = None
  parser = tr.Parser(program, words, -1)
  try:
    tree = parser.Csuite()
  except:
    info = sys.exc_info()
    traceback.print_tb(info[2])

    print ''
    print 'Before:', repr('|'.join([repr(x) for x in parser.words[:parser.p]]))
    print 'After:', repr('|'.join([repr(x) for x in parser.words[parser.p:]]))

    print ''
    print 'Before:', repr(program[:parser.words[parser.p][2]])
    print ''
    print 'After:', repr(program[parser.words[parser.p][2]:])
    print ''
    print 'SYNTAX ERROR:', info[1]

    sys.exit(13)

  cg = tr.CodeGen(None)
  try:
    cg.GenModule(None, None, tree)
  except:
    info = sys.exc_info()
    traceback.print_tb(info[2])
    print ''
    print 'ERROR', info[1]
    sys.exit(23)
