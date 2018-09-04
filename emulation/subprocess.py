from go import os
from go import os/exec as E

def call(vec, stdin=None, stdout=None, stderr=None):
  try:
    cmd = E.Command(vec[0], *vec[1:])
    # See PYE_FileDesc in builtins.py, for why .f:
    cmd.Stdin = os.Stdin if stdin is None else stdin.f
    cmd.Stdout = os.Stdout if stdout is None else stdout.f
    cmd.Stderr = os.Stderr if stderr is None else stderr.f
    cmd.Run()
    return 0
  except as ex:
    msg = 'PYE: subcommand.call: ERROR <<<%s>>> CALLING <<<%s>>>' % (ex, repr(vec))
    say msg
    return 13
