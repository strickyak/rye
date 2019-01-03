# rye.py is comparable to the "go" command that comes with go.
#
#  Usage:
#      alias rye='python .../compiler/rye.py'
#
#      rye build filename.py
#      rye run filename.py ...args...
#
#  The result of building "filename.py" is the binary "filename/filename".

import os          # rye_pragma from "github.com/strickyak/rye/emulation"
import re          # rye_pragma from "github.com/strickyak/rye/emulation"
import subprocess  # rye_pragma from "github.com/strickyak/rye/emulation"
import sys         # rye_pragma from "github.com/strickyak/rye/emulation"
import time        # rye_pragma from "github.com/strickyak/rye/emulation"

import lex  # The rye lexical scanner.
import parse  # The rye parser.
import codegen  # The rye compiler.
import linemap  # Maps .go lines to .py lines.

rye_true = False  # Magic variable:  if compiled by rye, rye_true is always True.

GOPATH = os.getenv('GOPATH').split(':')[0]

PATH_MATCH = re.compile('(.*)/src/(.*)').match

PrintSteps = 0

def GoBuild():
  if PrintSteps > 1: return ['go', 'build', '-x']
  return ['go', 'build']

def GoInstall():
  if PrintSteps > 1: return ['go', 'install', '-x']
  return ['go', 'install']

def TranslateModuleAndDependencies(filename, longmod, mod, cwd, twd, did, opts):
  already_compiled, imports = TranslateModule(filename, longmod, mod, cwd, opts)
  did[longmod] = True

  for k, v in imports.items():
    if v.imported[0] == 'go':
      continue # Don't traverse "go" dependencies.

    longpath2 = v.imported
    longmod2 = '/'.join(longpath2)
    mod2 = longpath2[-1]
    filename2 = '/' + twd + '/src/' + longmod2 + '.py'

    if not did.get(longmod2):
      TranslateModuleAndDependencies(filename2, longmod2, mod2, os.path.dirname(longmod2), twd, did, opts)

  # We put this off until the dependencies are built.
  # Installing speeds up everything.
  if not already_compiled:
    Execute(GoInstall() + ['%s/rye_/%s' % (os.path.dirname(longmod), mod)])

def BuildBuiltins(ryefile, gofile, generated_pyfile, opts):
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
  codegen.CodeGen().GenModule('BUILTINS', 'BUILTINS', tree, 'BUILTINS', internal=generated_pyfile, opts=opts)
  sys.stdout.close()
  t3 = time.time()
  print >> sys.stderr, '{{ CodeGen took %.3f seconds }}' % (t3-t2)
  sys.stdout = None
  if not os.getenv("RYE_NOFMT"):
    Execute( ['gofmt', '-w', gofile] )
  finish = time.time()
  print >>sys.stderr, '{{ build_builtins: DURATION %9.3f }}' % (finish-start)

def TranslateModule(filename, longmod, mod, cwp, opts):
  print >>sys.stderr, '=== TranslateModule', [filename, longmod, mod, cwp]
  print >>sys.stderr, '=== TranslateModule opts:', opts
  d = os.path.dirname(filename)
  b = os.path.basename(filename).split('.')[0]
  d2 = os.path.join(d, 'rye_', b)

  try:
    os.makedirs(d2)
  except:
    pass

  gopath = os.path.join(d2, 'rye_module.go')
  print >>sys.stderr, '=== TranslateModule gopath:', gopath

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
    # TODO: Get imports without entire codegen running.
    sys.stdout = open('/dev/null', 'w')
  else:
    sys.stdout = open(gopath, 'w')
  gen = codegen.CodeGen()
  gen.GenModule(mod, longmod, tree, cwp, internal="", opts=opts)
  sys.stdout.flush()
  sys.stdout.close()
  sys.stdout = None
  finish = time.time()
  print >> sys.stderr, '{{ CodeGen took %.3f seconds }}' % (finish-t2)
  print >>sys.stderr, '{{ %s: %s; took %9.3f }}' % (
      longmod, "already_compiled" if already_compiled else "Compiled",
      finish-start)

  if not already_compiled:
    if not os.getenv("RYE_NOFMT"):
      cmd = ['gofmt', '-w', gopath]
      Execute(cmd)

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
    print >>w, 'func init() { RegisterLineInfo("%s/rye_/%s", &lineInfo) }' % (os.path.dirname(longmod), os.path.basename(longmod))
    w.close()

  return already_compiled, gen.imports

def WriteMain(filename, longmod, mod, toInterpret, opts):
  d = os.path.dirname(filename)
  # Adding _main so Rye Tests like "something_test.py" will work.
  b = '%s_main' % os.path.basename(filename).split('.')[0]
  m = '%s' % b

  try:
    os.makedirs(os.path.join(d, 'rye_'))
  except:
    pass
  gopath = os.path.join(d, 'rye_', '%s.go' % b)
  w = open(gopath, 'w')

  print >>w, '''// +build rye_main

package main
import "os"
import "runtime/pprof"
import rye "github.com/strickyak/rye/runtime"
import MY "%s/rye_/%s"

var _ = os.Args
func main() {
  rye.RememberRyeCompileOptions(`%s`)

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
      if rye.DebugExcept < 1 { rye.DebugExcept = 1 }
      rye.PrintStackFYIUnlessEOFBecauseExcept(r)
      panic(r)
    }
  }()

  defer rye.Flushem()
  MY.G___name__ = rye.MkStr("__main__")
  MY.Eval_Module()

  //  // This code was for interp:
  //  glbl := rye.MkDict(make(rye.Scope))
  //  for k, ptr := range MY.ModuleMap {
  //    glbl.SetItem(rye.MkStr(k), *ptr)
  //  }
  //  interp.Eval_Module()
  //  sco := interp.G_0_Scopes().(*interp.C_Scopes)
  //  sco.M_g = glbl
  //  interp.G_1_Repl(sco)

  MY.G_1_main(rye.MkStrs(os.Args[1:]))

  rye.Shutdown()
}
''' % (os.path.dirname(longmod), os.path.basename(longmod), opts)

  w.close()
  return gopath


def Build(ryefile, toInterpret, opts):
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

  main_filename = WriteMain(ryefile, longmod, mod, toInterpret, opts)

  TranslateModuleAndDependencies(ryefile, longmod, mod, cwd, twd, did, opts)

  target = "%s/%s" % (d if d else '.', mod)
  cmd = GoBuild() + ['-o', target, main_filename]
  Execute(cmd)

  # Return the binary filename.
  return target


def Execute(cmd, stdin=None, stdout=None, stderr=None):
  pretty = ' '.join([repr(s) for s in cmd])
  if PrintSteps: print >> sys.stderr, "\n++++++ %s" % pretty
  start = time.time()
  status = subprocess.call(cmd, stdin=stdin, stdout=stdout, stderr=stderr)
  end = time.time()
  if PrintSteps: print >> sys.stderr, '       : {{ %s took %.3f seconds }}' % (cmd[0], (end-start))
  if status:
    print >> sys.stderr, "\nFAILURE (exit status %d) IN COMMAND: %s" % (status, pretty)
    sys.exit(status)


def Help():
  print >> sys.stderr, """
Usage:
  python rye.py --opts=ComileOptions build filename.py
  python rye.py --opts=ComileOptions run filename.py args...
"""

MATCH_OPTS = re.compile('[-][-]?opts=(.*)').match

def Main(args):
  global PrintSteps
  start = time.time()
  opts = ''
  PrintSteps = 0

  while args and args[0] and args[0][0]=='-':
    flag = args[0]
    args = args[1:]
    m = MATCH_OPTS(flag)
    if m:
      opts = m.group(1)
    elif flag == '-O':
      opts += 'O'
    elif flag == '-x':
      PrintSteps += 1
    else:
      print >>sys.stderr, "ERROR: unknown flag: %s" % repr(flag)

  opts = ''.join(sorted([c for c in opts]))

  # This opts check is important, because they are inserted into main.
  if not re.compile('^[A-Za-z0-9]*$').match(opts):
    print >> sys.stderr, """ERROR: Illegal compile options.
Only letters A-Z or a-z or digits 0-9 may be used in --opts=
"""
    sys.exit(2)

  cmd = args[0] if args else 'help'
  if cmd == 'build_builtins':
    BuildBuiltins(args[1], args[2], args[3], opts=opts)
  elif cmd == 'build':
    Build(args[1], toInterpret=False, opts=opts)
  elif cmd == 'run':
    binfile = Build(args[1], toInterpret=False, opts=opts)
    Execute ([binfile] + args[2:])
  elif cmd == 'interpret':
    binfile = Build(args[1], toInterpret=True, opts=opts)
    Execute ([binfile] + args[2:])
  else:
    Help()
    sys.exit(2)

  finish = time.time()
  print >>sys.stderr, '{{ Finished in %9.3f: %s }}' % (finish-start, ' '.join(args))

if __name__ == '__main__':
  Main(sys.argv[1:])
