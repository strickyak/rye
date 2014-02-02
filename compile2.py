import os
import sys

import tr2

if __name__ == '__main__':
  program = open(sys.argv[1]).read()

  words = tr2.Lex(program).tokens
  tree = tr2.Parser(program, words, -1).Run()
  tr2.GenModule('*NAME*', '*PATH*', tree)
