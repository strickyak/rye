import os
import os.path
import re
import sys

import tr

AT_AT = re.compile('[@][@]')

class ReWrite:
  def __init__(self, fd):
    self.fd = fd
  def write(self, x):
    x = AT_AT.sub('', x)
    self.fd.write(x)
  def close(self):
    self.fd.close()
  def flush(self):
    self.fd.flush()

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

  sys.stdout = ReWrite(open('%s/__%s/%s.go' % (dirname, modname, modname), "w"))
  main = ReWrite(open('%s/__%s/__main/main.go' % (dirname, modname), "w"))
  tr.Generator(None).GenModule('%s/__%s' % (dirname, modname), filename, tree, main)
