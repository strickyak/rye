from go import io, io/ioutil
from go import os

stdin = PYE_FileDesc(os.Stdin, writing=False)
stdout = PYE_FileDesc(os.Stdout, writing=True)
stderr = PYE_FileDesc(os.Stderr, writing=True)

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
