# rye.py -- strick
import os
import os.path
import re
import subprocess
import sys
import traceback

import tr

PROFILE = os.getenv("RYE_PROFILE")

PATH_MATCH = re.compile('(.*)/src/(.*)').match

def TranslateModule(filename, longmod, mod, cwp):
  print >>sys.stderr, '*** TranslateModule', [filename, longmod, mod, cwp]
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, b))
  except:
    pass

  wpath = os.path.join(d, b, 'ryemodule.go') 

  # BUG: If we don't recompile one, we may not notice its dirty dependency.
  #try:
  #  w_st = os.stat(wpath)
  #  w_mtime = w_st.st_mtime
  #except:
  #  w_mtime = 0
  #r_st = os.stat(filename)
  #r_mtime = r_st.st_mtime
  #if w_mtime > r_mtime:
  #  print >> sys.stderr, "*** ALREADY COMPILED: %s" % filename
  #  return {}

  sys.stdout = open(wpath, 'w')
  program = open(filename).read()
  words = tr.Lex(program).tokens

  #print >> sys.stderr, "\n\n(TOKENS)<<<", words, "\n\n"
  words = list(tr.SimplifyContinuedLines(words))
  #print >> sys.stderr, "\n\n(TOKENS)>>>", words, "\n\n"

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

  if not os.getenv("RYE_NOFMT"):
    cmd = ['gofmt', '-w', wpath]
    status = Execute(['gofmt', '-w', wpath])
    if status:
      raise Exception('Failure in gofmt: %s' % cmd)

  return gen.imports

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
  if PROFILE:
    print >>w, 'import "runtime/pprof"'
  print >>w, 'import "github.com/strickyak/rye"'
  print >>w, 'import MY "%s"' % longmod
  print >>w, '''
    func main() {
  '''
  if PROFILE:
    print >>w, '''
      f, err := os.Create("/tmp/rye.cpu")
      if err != nil {
        panic(err)
      }
      pprof.StartCPUProfile(f)
      defer pprof.StopCPUProfile()
    '''

  print >>w, '''
      MY.Eval_Module()
      MY.G_1_main(rye.MkStrs(os.Args[1:]))
  }
  '''
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
  print >>sys.stderr, "*** BUILD", to_run, args
  pwd = os.getcwd()
  m = PATH_MATCH(pwd)
  if not m:
    raise Exception('PWD does not contain /src/: %s' % pwd)
  twd, cwd = m.groups()
  cwd_split = cwd.split('/')

  main_filename = None
  main_longmod = None
  main_mod = None
  first = True
  did = {}
  full_run_args = []
  todo = [ args ]
  while todo:
    run_args = None
    chunk = todo.pop(0)
    for a in chunk:
      if run_args is not None:  # After --, just collect run_args.
        run_args.append(a)
        continue

      if a == '--':  # On --, switch to run_args mode.
        run_args = []
        continue

      if did.get(a):  # Don't do any twice.
        continue

      d = os.path.dirname(a)

      mod = os.path.basename(a).split('.')[0]  # Part before '.' in basename becomes package name.

      if d == '.' or d == "":
        longmod = '%s/%s' % (cwd, mod)
      else:
        longmod = '%s/%s/%s' % (cwd, d, mod)
      longdir = os.path.dirname(longmod)

      imports = TranslateModule(a, longmod, mod, cwd)
      did[a] = True

      if first:
        main_longmod = longmod
        main_mod = mod
        main_filename = WriteMain(a, longmod, mod)
        first = False

      for k, v in imports.items():
        # print >> sys.stderr, "####### %s -> %s" % (k, vars(v))
        if v.fromWhere is None:  # Todo, handle more fromWhere.

          if v.imported[:len(cwd_split)] != cwd_split:
            raise Exception("Cannot handle this import yet: %s (vs %s)", v.imported, cwd_split)

          impfile = '%s.py' % ('/'.join(v.imported[len(cwd_split):]))
          # print >> sys.stderr, "IMPFILE %s" % impfile

          if not did.get(impfile):
            todo.append([impfile])
            # print >> sys.stderr, "ADDED %s" % impfile

    full_run_args += (run_args if run_args else [])


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
    cmd = ['%s/%s' % (main_mod, main_mod)] + full_run_args
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
