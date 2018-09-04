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

rye_rye = False  # Magic variable:  if compiled by rye, rye_rye is always True.
if rye_rye:
  from . import lex  # The rye lexical scanner.
  from . import parse  # The rye parser.
  from . import codegen2  # The rye compiler.
  from . import linemap  # Maps .go lines to .py lines.
else:
  import lex  # The rye lexical scanner.
  import parse  # The rye parser.
  import codegen2  # The rye compiler.
  import linemap  # Maps .go lines to .py lines.

GOPATH = os.getenv('GOPATH').split(':')[0]

PATH_MATCH = re.compile('(.*)/src/(.*)').match

def TranslateModuleAndDependencies(filename, longmod, mod, cwd, twd, did):
  already_compiled, imports = TranslateModule(filename, longmod, mod, cwd)
  did[longmod] = True

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
    Execute(['go', 'install', '%s/rye__/%s' % (os.path.dirname(longmod), mod)])

def BuildBuiltins(ryefile, gofile):
  start = time.time()
  program = open(ryefile).read()
  words = lex.Lex(program, filename=ryefile).tokens
  words = list(lex.SimplifyContinuedLines(words, filename=ryefile))
  t1 = time.time()
  print >> sys.stderr, '{{ Lexer took %.3f seconds }}' % (t1-start)
  parser = parse.Parser(program, words, -1, 'BUILTINS')

  try:
    tree = parser.Csuite()
  except Exception as err:
    print lex.AddWhereInProgram(str(err), len(program) - len(parser.Rest()), filename=ryefile)
    sys.exit(13)
  t2 = time.time()
  print >> sys.stderr, '{{ Parser took %.3f seconds }}' % (t2-t1)

  sys.stdout = open(gofile, 'w')
  codegen2.CodeGen().GenModule('BUILTINS', 'BUILTINS', tree, 'BUILTINS', internal=True)
  sys.stdout.close()
  t3 = time.time()
  print >> sys.stderr, '{{ CodeGen took %.3f seconds }}' % (t3-t2)
  sys.stdout = None
  ###if not os.getenv("RYE_NOFMT"):
  ###  Execute( ['gofmt', '-w', gofile] )
  finish = time.time()
  print >>sys.stderr, '{{ build_builtins: DURATION %9.3f }}' % (finish-start)

def TranslateModule(filename, longmod, mod, cwp):
  print >>sys.stderr, '=== TranslateModule', [filename, longmod, mod, cwp]
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, 'rye__', b))
  except:
    pass

  popath = os.path.join(d, 'rye__', b, 'ryemodule.pre.go')
  gopath = os.path.join(d, 'rye__', b, 'ryemodule.go')

  # If we don't recompile one, we may not notice its dirty dependency.
  w_st = None
  w_mtime = 0  # As old as the hills.
  try:
    w_mtime = os.stat(gopath).st_mtime
  except:
    pass
  r_mtime = os.stat(filename).st_mtime
  already_compiled = (w_mtime > r_mtime)
  if already_compiled:
    print >>sys.stderr, 'Already Compiled: ', filename

  start = time.time()
  program = open(filename).read()
  words = lex.Lex(program, filename=filename).tokens
  t1 = time.time()
  print >> sys.stderr, '{{ Lexer took %.3f seconds }}' % (t1-start)
  words = list(lex.SimplifyContinuedLines(words, filename=filename))
  parser = parse.Parser(program, words, -1, cwp)
  try:
    tree = parser.Csuite()
  except Exception as err:
    print >>sys.stderr, lex.AddWhereInProgram(str(err), len(program) - len(parser.Rest()), filename=filename)
    sys.exit(13)
  t2 = time.time()
  print >> sys.stderr, '{{ Parser took %.3f seconds }}' % (t2-t1)

  if already_compiled:
    # TODO: Get imports without entire codegen2 running.
    sys.stdout = open('/dev/null', 'w')
  else:
    sys.stdout = open(popath, 'w')
  gen = codegen2.CodeGen()
  gen.GenModule(mod, longmod, tree, cwp, internal=False)
  sys.stdout.flush()
  sys.stdout.close()
  sys.stdout = None
  finish = time.time()
  print >> sys.stderr, '{{ CodeGen took %.3f seconds }}' % (finish-t2)
  print >>sys.stderr, '{{ %s: %s; took %9.3f }}' % (
      longmod, "already_compiled" if already_compiled else "Compiled",
      finish-start)

  if not already_compiled:
    if NOINLINE:
      cmd = [GOPATH + '/src/github.com/strickyak/prego/main', '--noinline', '--set', 'noinline', '--source', GOPATH + '/src/github.com/strickyak/rye/macros2.pre.go']
    else:
      cmd = [GOPATH + '/src/github.com/strickyak/prego/main', '--source', GOPATH + '/src/github.com/strickyak/rye/macros2.pre.go']
    rfd, wfd = open(popath, 'r'), open(gopath, 'w')
    Execute(cmd, stdin=rfd, stdout=wfd)
    rfd.close()
    wfd.close()

    #cmd = ['gofmt', '-w', gopath]
    #Execute(cmd)

    lm, ld, lw = linemap.ScanFileForLinemap(gopath, filename)
    w = open(gopath, 'a')
    print >>w, 'var lineMap = []int32{', ','.join([str(x) for x in lm]), '}'
    print >>w, 'var srcLines = []IntStringPair{'
    for k5, v5 in ld.items():
      print >>w, '  {N: %d, S: `%s`},' % (int(k5), v5.replace('`', '?'))
    print >>w, '}'
    print >>w, 'var srcWhats = []IntStringPair{'
    for k5, v5 in lw.items():
      if v5:
        print >>w, '  {N: %d, S: `%s`},' % (int(k5), v5.replace('`', '?'))
      else:
        print >>w, '  {N: %d, S: `?`},' % int(k5)
    print >>w, '}'
    print >>w, 'var lineInfo = LineInfo{LookupLineNumber: lineMap, SourceLines: srcLines, SourceWhats: srcWhats}'
    print >>w, 'func init() { RegisterLineInfo("%s/rye__/%s", &lineInfo) }' % (os.path.dirname(longmod), os.path.basename(longmod))
    w.close()

  return already_compiled, gen.imports

def WriteMain(filename, longmod, mod, toInterpret):
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]

  try:
    os.makedirs(os.path.join(d, 'rye__', b, b))
  except:
    pass
  popath = os.path.join(d, 'rye__', b, b, 'ryemain.pre.go')
  gopath = os.path.join(d, 'rye__', b, b, 'ryemain.go')
  w = open(popath, 'w')

  print >>w, '''// +build prego
// +build ignore_main

package main
import "os"
import "unsafe"
import "reflect"
import "runtime/pprof"
import . "github.com/strickyak/rye"
// import "github.com/strickyak/rye/interp"
import MY "%s/rye__/%s"

var _ = unsafe.Sizeof(0)
var _ = os.Args
func main() {

  ppfile := os.Getenv("RYE_PPROF")
  if ppfile != "" {
    f, err := os.Create(ppfile)
    if err != nil {
      panic(err)
    }
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
  }

  defer func() {
    // Catch and print FYI for uncaught outer exceptions.
    r := recover()
    if r != nil {
      if DebugExcept < 1 { DebugExcept = 1 }
      PrintStackFYIUnlessEOFBecauseExcept(r)
      panic(r)
    }
  }()

  defer Flushem()

  var MAIN string = "__main__"
  MY.H___name___1,  MY.H___name___2  = macro.MkStrJ(MAIN)

  _,_ = MY.JEval_Module()

  //  // This code was for interp:
  //  glbl := MkDict(make(Scope))
  //  for k, ptr := range MY.ModuleMap {
  //    glbl.SetItem(MkStr(k), *ptr)
  //  }
  //  interp.Eval_Module()
  //  sco := interp.H_0_Scopes().(*interp.C_Scopes)
  //  sco.M_g = glbl
  //  interp.H_1_Repl(sco)

  args := MkStrsJ(os.Args[1:])
  MY.H_1_main(inline.MkPJ(args))

  Shutdown()
}
''' % (os.path.dirname(longmod), os.path.basename(longmod))

  w.close()
  cmd = [GOPATH + '/src/github.com/strickyak/prego/main', '--source', GOPATH + '/src/github.com/strickyak/rye/macros2.pre.go']
  rfd, wfd = open(popath, 'r'), open(gopath, 'w')
  Execute(cmd, stdin=rfd, stdout=wfd)
  rfd.close()
  wfd.close()
  return gopath


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
  longmod = '/'.join(codegen2.CleanPath('/', longmod))

  main_filename = WriteMain(ryefile, longmod, mod, toInterpret)

  TranslateModuleAndDependencies(ryefile, longmod, mod, cwd, twd, did)

  target = "%s/%s.bin" % (d if d else '.', mod)
  cmd = ['go', 'build', '-o', target, main_filename]
  Execute(cmd)

  # Return the binary filename.
  return target


def Execute(cmd, stdin=None, stdout=None, stderr=None):
  pretty = ' '.join([repr(s) for s in cmd])
  print >> sys.stderr, "\n++++++ %s" % pretty
  start = time.time()
  status = subprocess.call(cmd, stdin=stdin, stdout=stdout, stderr=stderr)
  end = time.time()
  print >> sys.stderr, '{{ %s took %.3f seconds }}' % (cmd[0], (end-start))
  if status:
    print >> sys.stderr, "\nFAILURE (exit status %d) IN COMMAND: %s" % (status, pretty)
    sys.exit(status)


def Help():
  print >> sys.stderr, """
Usage:
  python rye.py ?-pprof=cpu.out? build filename.py
  python rye.py ?-pprof=cpu.out? run filename.py args...
"""

NOINLINE = False

def Main(args):
  global NOINLINE
  start = time.time()


  cmd = args[0] if args else 'help'
  if cmd == '--noinline':
    NOINLINE = True
    args = args[1:]
    cmd = args[0] if args else 'help'

  if cmd == 'build_builtins':
    BuildBuiltins(args[1], args[2])
  elif cmd == 'build':
    Build(args[1], toInterpret=False)
  elif cmd == 'run':
    binfile = Build(args[1], toInterpret=False)
    Execute ([binfile] + args[2:])
  elif cmd == 'interpret':
    binfile = Build(args[1], toInterpret=True)
    Execute ([binfile] + args[2:])
  else:
    Help()
    sys.exit(2)

  finish = time.time()
  print >>sys.stderr, '{{ Finished in %9.3f: %s }}' % (finish-start, ' '.join(args))

if __name__ == '__main__':
  Main(sys.argv[1:])
