from go import os/exec as E

def call(vec):
  try:
    E.Command(vec[0], vec[1:]).Run()
    return 0
  except as ex:
    say 'PYE: subcommand.call: ERROR <<<%s>>> CALLING <<<%s>>>' % (ex, repr(vec))
    return 13
