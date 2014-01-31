import os
import sys

import tr

if __name__ == '__main__':
  program = open(sys.argv[1]).read()

  words = tr.Lex(program).tokens
  z = tr.Parser(program, words, -1).Run()
