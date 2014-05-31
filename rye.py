# rye.py -- strick
import os
import os.path
import re
import sys

import tr

PATH_MATCH = re.compile('(.*)/src/(.*)').match

def TranslateModule(filename, longmod, mod, cwp):
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, b))
  except:
    pass
  wpath = os.path.join(d, b, 'ryemodule.go') 
  sys.stdout = open(wpath, 'w')

  program = open(filename).read()
  words = tr.Lex(program).tokens
  tree = tr.Parser(program, words, -1).Csuite()

  tr.CodeGen(None).GenModule(mod, longmod, tree, cwp)
  sys.stdout.close()
  status = os.system('gofmt -w "%s"' % wpath)
  return wpath

def WriteMain(filename, longmod, mod):
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, b, 'main'))
  except:
    pass
  wpath = os.path.join(d, b, 'main', 'ryemain.go') 
  w = open(wpath, 'w')

  print >>w, 'package main'
  print >>w, 'import MY "%s"' % longmod
  print >>w, 'func main() {'
  print >>w, '  MY.Eval_Module();'
  print >>w, '}'
  w.close()
  return wpath


def BuildRun(to_run, args):
  print >>sys.stderr, "#+# BUILD", args
  pwd = os.getcwd()
  m = PATH_MATCH(pwd)
  if not m:
    raise Exception('PWD does not contain /src/: %s' % pwd)
  twd, cwd = m.groups()
  print >>sys.stderr, "#+# TWD=", twd
  print >>sys.stderr, "#+# CWD=", cwd

  main_filename = None
  main_longmod = None
  main_mod = None
  first = True
  did = {}
  for filename in args:
    if did.get(filename):
      print >>sys.stderr, "#+# ALREADY DID FILENAME", filename
      continue
    print >>sys.stderr, "#+# FOR FILENAME", filename
    d = os.path.dirname(filename)
    mod = os.path.basename(filename).split('.')[0]
    if d == '.' or d == "":
      longmod = '%s/%s' % (cwd, mod)
    else:
      longmod = '%s/%s/%s' % (cwd, d, mod)

    TranslateModule(filename, longmod, mod, cwd)
    did[filename] = True

    if first:
      main_longmod = longmod
      main_mod = mod
      main_filename = WriteMain(filename, longmod, mod)
      first = False

  bindir = os.path.dirname(os.path.dirname(main_filename))
  target = "%s/%s" % (bindir, main_mod)

  cmd = "go build -o '%s' '%s'" % (target, main_filename)
  print >> sys.stderr, "+ %s" % cmd
  status = os.system(cmd)
  if status:
    print >> sys.stderr, "%s: Exited with status %d" % (main_longmod, status)

def Help(args):
  print >> sys.stderr, """
Usage:
  python rye.py build filename.py
"""

def main(args):
  cmd = args[0] if len(args) else 'help'

  if cmd == 'build':
    return BuildRun(False, args[1:])
  if cmd == 'run':
    return BuildRun(True, args[1:])
  if cmd == 'help':
    return Help(args[1:])
  return Help(args[1:])


if __name__ == '__main__':
  main(sys.argv[1:])
