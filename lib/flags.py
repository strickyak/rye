"""
Package flags parses command line flags, similar to golang's flag package.

Flags have one or two dashes and equals, like "--flag_test_b" or "--flag_test_s=foo".
You can omit "=value" and the value will be "1" or 1 or True.
The same flag should not be repeated (or last value wins).

Call   args = flags.Munch(args)    inside main() to process flags.

The value of a flag is in the .X field of the object.
"""
from go import regexp

Flags = {}

FLAG_RE = regexp.MustCompile('^[-][-]?(.+?)([=](.*))?$')

def Munch(args):
  # Consume flags beginning with '-' or '--'.
  while args:
    m = FLAG_RE.FindStringSubmatch(args[0])
    if not m:
      break
    _, name, valued, value = m
    f = Flags.get(name)
    if not f:
      raise 'No such flag', name
    if valued:
      f.Set(value)
    else:
      f.Set()
    args = args[1:]

  # Consume '--' as a flags-stopper.
  if args and args[0] == '--':
    args = args[1:]

  # Return the remaining args.
  return args

class String:
  def __init__(name, dflt, doc):
    .name = name
    .dflt = dflt
    .X = str(dflt)
    .doc = doc
    must name not in Flags
    Flags[name] = self

  def Value():
    return .X

  def Set(x="1"):
    .X = str(x)

class Int:
  def __init__(name, dflt, doc):
    .name = name
    .dflt = dflt
    .X = int(dflt)
    .doc = doc
    must name not in Flags
    Flags[name] = self

  def Value():
    return .X

  def Set(x=1):
    .X = int(x)

class Bool:
  def __init__(name, dflt, doc):
    .name = name
    .dflt = dflt
    .X = bool(dflt)
    .doc = doc
    must name not in Flags
    Flags[name] = self

  def Value():
    return .X

  def Set(x=True):
    .X = bool(x)

pass
