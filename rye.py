# rye.py is comparable to the "go" command that comes with go.
#
#  Usage:
#      alias rye='python .../rye.py'
#
#      rye build filename.py
#      rye run filename.py ...args...
#
#  The result of building "filename.py" is the binary "filename/filename".

import os
import re
import subprocess
import sys
import time
import traceback

rye_rye = False  # Magic variable:  if compiled by rye, rye_rye is always True.
if rye_rye:
  from . import lex  # The rye lexical scanner.
  from . import parse  # The rye parser.
  from . import codegen  # The rye compiler.
  from . import linemap  # Maps .go lines to .py lines.
else:
  import lex  # The rye lexical scanner.
  import parse  # The rye parser.
  import codegen  # The rye compiler.
  import linemap  # Maps .go lines to .py lines.

PATH_MATCH = re.compile('(.*)/src/(.*)').match

def TranslateModuleAndDependencies(filename, longmod, mod, cwd, twd, did):
  already_compiled, imports = TranslateModule(filename, longmod, mod, cwd)
  did[longmod] = True

  if os.getenv('JUST_TRANSLATE'):
    return

  for k, v in imports.items():
    if v.imported[0] == 'go':
      continue # Don't traverse "go" dependencies.

    longpath2 = v.imported
    longmod2 = '/'.join(longpath2)
    mod2 = longpath2[-1]
    filename2 = '/' + twd + '/src/' + longmod2 + '.py'

    if not did.get(longmod2):
      TranslateModuleAndDependencies(filename2, longmod2, mod2, os.path.dirname(longmod2), twd, did)

  # We put this off until the dependencies are built.
  # Installing speeds up everything.
  if not already_compiled:
    Execute(['go', 'install', longmod])


def TranslateModule(filename, longmod, mod, cwp):
  print >>sys.stderr, '=== TranslateModule', [filename, longmod, mod, cwp]
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, b))
  except:
    pass

  wpath = os.path.join(d, b, 'ryemodule.go')

  # If we don't recompile one, we may not notice its dirty dependency.
  w_st = None
  w_mtime = 0  # As old as the hills.
  try:
    w_mtime = os.stat(wpath).st_mtime
  except:
    pass
  r_mtime = os.stat(filename).st_mtime
  already_compiled = (w_mtime > r_mtime)
  if already_compiled:
    print >>sys.stderr, 'Alreadyccompiled: ', filename

  start = time.time()
  program = open(filename).read()
  words = lex.Lex(program).tokens
  words = list(lex.SimplifyContinuedLines(words))
  parser = parse.Parser(program, words, -1, cwp)
  try:
    tree = parser.Csuite()
  except:
    print >> sys.stderr, "\n*** TRACEBACK:"
    traceback.print_tb(sys.exc_info()[2])
    print >> sys.stderr, "\n*** OCCURRED BEFORE THIS: ", repr(parser.Rest()[:100])
    print >> sys.stderr, "\n*** ERROR: ", sys.exc_info()[1]
    sys.exit(13)

  if already_compiled:
    print >> sys.stderr, "Already Compiled:", longmod
    sys.stdout = open('/dev/null', 'w')
  else:
    sys.stdout = open(wpath, 'w')
  gen = codegen.CodeGen()
  gen.GenModule(mod, longmod, tree, cwp)
  sys.stdout.flush()
  sys.stdout.close()
  sys.stdout = None
  finish = time.time()
  print >>sys.stderr, '{{ %s: %s DURATION %9.3f }}' % (
      longmod, "already_compiled" if already_compiled else "Compiled",
      finish-start)

  if not already_compiled:
    if not os.getenv("RYE_NOFMT"):
      cmd = ['gofmt', '-w', wpath]
      Execute(cmd)

    lm = linemap.ScanFileForLinemap(wpath)
    w = open(wpath, 'a')
    print >>w, 'var linemap = []int32{', ','.join([str(x) for x in lm]), '}'
    print >>w, 'func init() { RegisterLinemap("%s", linemap) }' % longmod
    w.close()

  return already_compiled, gen.imports


def WriteMain(filename, longmod, mod, toInterpret):
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
  if toInterpret:
    print >>w, 'import "github.com/strickyak/rye/interp"'
  print >>w, 'import MY "%s"' % longmod
  print >>w, '''
    var _ = os.Args
    func main() {
  '''
  if PROFILE:
    print >>w, '''
      f, err := os.Create(`%s`)
      if err != nil {
        panic(err)
      }
      pprof.StartCPUProfile(f)
      defer pprof.StopCPUProfile()
    ''' % PROFILE

  if True:
    print >>w, '''
         defer func() {
           // Catch and print FYI for uncaught outer exceptions.
           r := recover()
           if r != nil {
             if rye.DebugExcept < 1 { rye.DebugExcept = 1 }
             rye.PrintStackFYIUnlessEOFBecauseExcept(r)
             panic(r)
           }
         }()
'''
  print >>w, '''
      defer rye.Flushem()
      MY.G___name__ = rye.MkStr("__main__")
      MY.Eval_Module()
  '''
  if toInterpret:
    print >>w, """
           glbl := rye.MkDict(make(rye.Scope))
           for k, ptr := range MY.ModuleMap {
             glbl.SetItem(rye.MkStr(k), *ptr)
           }

           interp.Eval_Module()
           sco := interp.G_0_Scopes().(*interp.C_Scopes)
           sco.M_g = glbl
           interp.G_1_Repl(sco)
     """

  else:
    print >>w, '      MY.G_1_main(rye.MkStrs(os.Args[1:]))'

  if PROFILE:
    print >>w, '      rye.Shutdown()'
  print >>w, '}'

  w.close()
  return wpath


def Build(ryefile, toInterpret):
  print >>sys.stderr, "rye build: %s" % ryefile
  pwd = os.getcwd()
  m = PATH_MATCH(pwd)
  if not m:
    raise Exception('PWD does not contain /src/: %s' % pwd)
  twd, cwd = m.groups()
  cwd_split = cwd.split('/')

  main_filename = None
  did = {}

  d = os.path.dirname(ryefile)
  mod = os.path.basename(ryefile).split('.')[0]  # Part before '.' in basename becomes package name.

  if d == '.' or d == "":
    longmod = '%s/%s' % (cwd, mod)
  else:
    longmod = '%s/%s/%s' % (cwd, d, mod)
  longmod = '/'.join(codegen.CleanPath('/', longmod))

  main_filename = WriteMain(ryefile, longmod, mod, toInterpret)

  TranslateModuleAndDependencies(ryefile, longmod, mod, cwd, twd, did)

  if os.getenv('JUST_TRANSLATE'):
    return

  bindir = os.path.dirname(os.path.dirname(main_filename))
  target = "%s/%s" % (bindir, mod)

  cmd = ['go', 'build', '-o', target, main_filename]
  Execute(cmd)

  # Return the binary filename.
  return target
  #return '%s/%s' % (mod, mod)


def Execute(cmd):
  pretty = ' '.join([repr(s) for s in cmd])
  print >> sys.stderr, "\n++++++ %s" % pretty
  status = subprocess.call(cmd)
  if status:
    print >> sys.stderr, "\nFAILURE (exit status %d) IN COMMAND: %s" % (status, pretty)
    sys.exit(status)


def Help():
  print >> sys.stderr, """
Usage:
  python rye.py ?-pprof=cpu.out? build filename.py
  python rye.py ?-pprof=cpu.out? run filename.py args...
"""


PROFILE = None
MATCH_PROF = re.compile('[-]+pprof=(.*)').match

def main(args):
  global PROFILE
  start = time.time()

  while args and args[0][0]=='-':
    opt = args.pop(0)

    m = MATCH_PROF(opt)
    if m:
      PROFILE = m.group(1)
    else:
      raise Exception("Unknown option: %s" % opt)

  cmd = args[0] if args else 'help'

  if cmd[0] == 'b':
    Build(args[1], toInterpret=False)
  elif cmd[0] == 'r':
    binfile = Build(args[1], toInterpret=False)
    Execute ([binfile] + args[2:])
  elif cmd[0] == 'i':
    binfile = Build(args[1], toInterpret=True)
    Execute ([binfile] + args[2:])
  elif cmd[0] == 'h':
    Help()
  else:
    Help()

  finish = time.time()
  print >>sys.stderr, '{{ Finished in %9.3f }}' % (finish-start)

if not rye_rye:
  if __name__ == '__main__':
    main(sys.argv[1:])
