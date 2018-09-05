from go import os

stdin = os.Stdin
stdout = os.Stdout
stderr = os.Stderr

class PYE_NeverFlushClose:
  def __init__(fd):
    .fd = fd
  def Write(x):
    .fd.Write(x)
  def write(x):
    .fd.Write(x)
  def Flush(x):
    pass
  def flush(x):
    pass
  def Close(x):
    pass
  def close(x):
    pass

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
