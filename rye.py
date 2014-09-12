# rye.py -- strick
import os
import os.path
import re
import subprocess
import sys
import traceback

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
  parser = tr.Parser(program, words, -1)
  try:
    tree = parser.Csuite()
  except:
    print >> sys.stderr, "\n*** ERROR: ", sys.exc_info()[1]
    print >> sys.stderr, "\n*** OCCURRED BEFORE THIS: ", parser.Rest()[:200], '......'
    print >> sys.stderr, "\n*** TRACEBACK:"
    traceback.print_tb(sys.exc_info()[2])
    sys.exit(13)

  gen = tr.CodeGen(None)
  gen.GenModule(mod, longmod, tree, cwp)
  sys.stdout.close()

  for k, v in gen.imports.items():
    print >> sys.stderr, "####### %s -> %s" % (k, vars(v))

  if not os.getenv("RYE_NOFMT"):
    cmd = ['gofmt', '-w', wpath]
    status = Execute(['gofmt', '-w', wpath])
    if status:
      raise Exception('Failure in gofmt: %s' % cmd)

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
  print >>w, 'import "os"'
  print >>w, 'import "github.com/strickyak/rye/runt"'
  print >>w, 'import MY "%s"' % longmod
  print >>w, 'func main() {'
  print >>w, '  MY.Eval_Module()'
  print >>w, '  MY.M_1_main(runt.MkStrs(os.Args[1:]))'
  print >>w, '}'
  w.close()
  return wpath

# TODO -- fix this, and break out more functions.
def LongMod(*TODO):
  d = os.path.dirname(a)
  mod = os.path.basename(a).split('.')[0]
  if d == '.' or d == "":
    return '%s/%s' % (cwd, mod)
  else:
    return '%s/%s/%s' % (cwd, d, mod)

def BuildRun(to_run, args):
  #print >>sys.stderr, "#+# BUILD", args
  pwd = os.getcwd()
  m = PATH_MATCH(pwd)
  if not m:
    raise Exception('PWD does not contain /src/: %s' % pwd)
  twd, cwd = m.groups()
  #print >>sys.stderr, "#+# TWD=", twd
  #print >>sys.stderr, "#+# CWD=", cwd

  main_filename = None
  main_longmod = None
  main_mod = None
  first = True
  did = {}
  run_args = None
  for a in args:
    if run_args is not None:
      run_args.append(a)
      continue
    if a == '--':
      run_args = []
      continue
    if did.get(a):
      continue
    d = os.path.dirname(a)
    mod = os.path.basename(a).split('.')[0]
    if d == '.' or d == "":
      longmod = '%s/%s' % (cwd, mod)
    else:
      longmod = '%s/%s/%s' % (cwd, d, mod)

    TranslateModule(a, longmod, mod, cwd)
    did[a] = True

    if first:
      main_longmod = longmod
      main_mod = mod
      main_filename = WriteMain(a, longmod, mod)
      first = False

  bindir = os.path.dirname(os.path.dirname(main_filename))
  target = "%s/%s" % (bindir, main_mod)

  cmd = "set -x; go build -o '%s' '%s'" % (target, main_filename)
  cmd = ['go', 'build', '-o', target, main_filename]
  print >> sys.stderr, "+ %s" % repr(cmd)
  status = Execute(cmd)
  if status:
    print >> sys.stderr, "%s: Exited with status %d" % (main_longmod, status)
    status = status if status&255 else 13
    sys.exit(status)

  if to_run:
    cmd = ['%s/%s' % (main_mod, main_mod)] + (run_args if run_args else [])
    print >> sys.stderr, "+ %s" % repr(cmd)
    status = Execute(cmd)
    if status:
      print >> sys.stderr, "%s: Exited with status %d" % (main_longmod, status)
      status = status if status&255 else 13
      sys.exit(status)

def Execute(cmd):
  return subprocess.call(cmd)

def Help(args):
  print >> sys.stderr, """
Usage:
  python rye.py build filename.py otherlibs.py...
  python rye.py run filename.py otherlibs.py... -- flags-and-args...
"""

def main(args):
  cmd = args[0] if len(args) else 'help'

  if cmd[0] == 'b':
    return BuildRun(False, args[1:])
  if cmd[0] == 'r':
    return BuildRun(True, args[1:])
  if cmd[0] == 'h':
    return Help(args[1:])
  return Help(args[1:])


if __name__ == '__main__':
  main(sys.argv[1:])
