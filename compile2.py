import os
import sys

import tr2

if __name__ == '__main__':
  program = open(sys.argv[1]).read()

  words = tr2.Lex(program).tokens
  # tree = tr2.Parser(program, words, -1).Run()
  tree = tr2.Parser(program, words, -1).Csuite()

  tr2.Generator(None).GenModule('*NAME*', '*PATH*', tree)
