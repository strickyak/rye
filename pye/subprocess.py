from go import os
from go import os/exec as E

def call(vec):
  try:
    cmd = E.Command(vec[0], *vec[1:])
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
    return 0
  except as ex:
    msg = 'PYE: subcommand.call: ERROR <<<%s>>> CALLING <<<%s>>>' % (ex, repr(vec))
    say msg
    return 13
