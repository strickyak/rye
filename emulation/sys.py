from go import os

stdin = os.Stdin
stdout = os.Stdout
stderr = os.Stderr

native:
  'func init() {'
  '  // Rye Runtime needs pointers:'
  '  PtrSysStdin = &G_stdin'
  '  PtrSysStdout = &G_stdout'
  '  PtrSysStderr = &G_stderr'
  '}'

argv = os.Args

def exit(status):
  os.Exit(status)
