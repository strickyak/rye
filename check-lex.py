import os
import sys

import tr

if __name__ == '__main__':
  program = open(sys.argv[1]).read()
  golden = open(sys.argv[2]).read()

  exp = eval(golden)
  got = tr.Lex(program).tokens
  if got != exp:
    print >> sys.stderr, " Program:", repr(program)
    print >> sys.stderr, "  Golden:", repr(golden)
    print >> sys.stderr, "Expected:", repr(exp)
    print >> sys.stderr, "     Got:", repr(got)
    raise Exception("FAIL")
