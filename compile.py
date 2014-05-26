import os
import sys

import tr


if __name__ == '__main__':
  #  if len(sys.argv) > 2
  program = open(sys.argv[1]).read()

  words = tr.Lex(program).tokens
  tree = tr.Parser(program, words, -1).Csuite()

  tr.CodeGen(None).GenModule(None, None, tree)
