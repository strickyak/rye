from go import log
from go import os
from go import "os/exec" as E

def call(vec :list, stdin=None, stdout=None, stderr=None) ->int :
  try:
    cmd = E.Command(vec[0], *vec[1:])
    # See RyeFileDesc in builtins.py, for why .f:
    cmd.Stdin = os.Stdin if stdin is None else stdin.f
    cmd.Stdout = os.Stdout if stdout is None else stdout.f
    cmd.Stderr = os.Stderr if stderr is None else stderr.f
    cmd.Run()
    return 0
  except as ex:
    log.Panicf('emulation/subcommand.call: ERROR %q WHILE CALLING %v', ex, vec)
