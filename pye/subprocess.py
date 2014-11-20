from go import os/exec as E

def call(vec):
  err = E.Command(vec[0], vec[1:]).Run()
  return 3 if err else 0
