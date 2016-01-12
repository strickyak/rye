## Rye Statements

/0

/1 print

say

/2
```
print 'The result is', result

import sys
print >>sys.stderr, 'The result is', result

w = bytes.NewBuffer()
print >>w, 'The result is', result

say result  # Rye extension.
```
/3
As in Python, prints the arguments with spaces as separators.

By default it prints on sys.stdout, but you can redirect to
any io.Writer.

The command "say" differs by printing to sys.stderr,
and printing the module, function or class and method,
line number, and the source code for its arguments, before
printing the repr() of the values of the arguments.
It is a convenience for debugging.

/1 assign
/2
```
_ = DoSomething()  # Value is thrown away.
s = "Hello Variable"
self.field = "save me in a field"

# Destructuring assignments:
a, b, c = 100, 200, 400
d, e, f = FunctionReturning3Values()
x, y = y, x  # Swap values.
a, (b, c), d = 10, range(2), 20
```
/3
As in Python. Assigning to _ does not save the value.

Destructuring assignments are supported.


/1 if
/2
```
if a == 0:
  DoZero()
elif a > 0:
  DoPositive(a)
else:
  DoNegative(a)
```
/3 As in Python.

False values are None, False, zero values 0 and 0.0,
and empty strings, lists, dicts, etc.: '', byt(''), (), [], {}, set({}).

/1 while
/2
```
while p is not Nil:
  print p.item
  p = p.next
```

/3 As in Python.

/1 for in

/2
```
for k, v in sorted(myDict.items()):
  if k is None:
    break  # Stop early on first missing value.
  print k, '->', v
```
/3
As in Python.

/1 try except
/2
```
  s = None
  try:
    s = ioutil.ReadFile(filename)
  except Exception as ex:
    raise Exception(
        'Cannot read file %q because %q'
        % (filename, ex))
```
/3
This syntax looks like Python, but it uses Go's panic() and recover().
In Go, there is no restriction on what types can be raised,
so the word Exception is just sugar, to be syntactially compatable with Python.
"except Exception" catches all types, even strings and tuples.

There is no guarding of the type caught with except.
It is not possible to catch only subtypes of some type,
unless you write your own code in the exception handler.

/1 try finally
/2
```
  fd = os.Open("/etc/termcap")
  try:
    contents = ioutil.ReadAll(fd))
    for line in str(contents).split('\n'):
      ProcessLine(line)
  finally:
    fd.Close()
```
/3 As in Python.

Flow control statements (like return, yield, break, & continue)
cannot be used inside the finally block.



/1 raise
/2
```
# These two throw a string:
raise 'something went wrong: %s' % what
raise Exception('something went wrong: %s' % what)

# This one throws a Go error:
raise errors.New('something went wrong: %s' % what)

# You can even raise a tuple:
raise 'something went wrong', what
```
/3
Rye uses Go's panic() and recover() for throwing and catching exceptions.
In Go, there is are no restrictions on what types can be thrown,
so Rye allows throwing and catching any data type.
The function Exception() is the identity function, for syntactic sugar
to be syntactically compatable with Python.

/1 assert

must

/2
```
assert x == y
assert y in someList
assert z is not None, 'z should not be None'
must x == y
must except 7 / 0
```
/3
As in Python.  If the outermost operation is a relational operator
( == != < <= > >= ) or one of the wordy binary operators "is", "is not",
"in", "not in", then if the assert fails, both the left hand side
and the right hand side arguments are printed, so often you
do not need to add an extra failure message.

The "must" statement is like the "assert" statement,
but it cannot be disabled by the runtime environment flag
that disables asserts;  "must" is always executed.

If "except" follows "assert" or "must", the statement only succeeds
if the expression results in an exception.  This helps you unit test
for failures.

/1 from import
/2
```
# For Rye module handlers.py in the same directory.
from . import handlers

# Special builtin rye libraries.
from rye_lib import data

# Go libraries.
from go import os, regexp
fron go import bytes as BYTES, math/big
from go import github.com/microcosm-cc/bluemonday

# Python emulation of a subset of a few basic
# modules, for bootstrapping Rye using Python.
import sys
```
/3
In Rye, you only import modules (not functions or classes in a module).
The final name in the path becomes the module variable name, unless you use "as" to give a new name.
That's different from Python:  In Python, if you say `import os.path`, you would call
`os.path.basename()`.  But in Rye, if you say `from go import math/big`,
you would call `big.NewInt()`, not `math.big.NewInt()`.

Use "from . import" for Rye modules from the same directory or from subdirectories.

Use "from rye_lib import" for a few special Rye libraries.

Use "from go import" to import Go packages.

The final style without the word "from" is only for certain
Python Emulation modules, so the Rye compiler can be bootstrapped
using Python.   However "sys.stdout" and "sys.stderr" might also be needed
for redirecting the output of the "print" and "say" statements.

/1 def
/2
```
def Twice(x): return x + x

def IncrementBy(n):
  "Returns a function that adds n to its arg.'
  def f(x):
    return x + n
  return f

def GeneratesSquares():
  i = 0
  while True:
    yield i * i
    i += 1
```
/3
As in Python.  No decorators.

You may used named arguments, default arguments,
*args, and **kw, all as in Python.
(All of these impose a performance hit; simple fixed number
of arguments is the fastest.)

/1 class
/2
```
class Apple(object):
  def __init__(self, color):
    self.color = color
  def __str__(self):
    return 'A %s apple.' % self.color

class Banana(Apple):
  def __str__(self):
    return super.__str__().replace(
        'apple', 'banana')

assert type(Banana("yellow")) is Banana
assert str(Banana("yellow")) == 'A yellow banana.'
```

/3
As in Python.  But:

Only single inheritance.  Classes are not old-style or new-style,
but simpler Rye-style.  All classes inherit from "object".

Only "def" statements (or "pass") can be used in the class body.
All methods must include the initial "self" argument (or it will
be automatically inserted for you!).  No other variables
or nested classes can be inserted in the class object.

There are no class or static methods; only instance methods.

No decorators.  No metaclasses.  No monkey-patching.

/1 switch case default
/2
```
switch color:
  case 'green':
    print 'You can go.'
  case 'yellow':
    print 'Slow down!'
  default:
    print 'You better stop.'

switch:
  case color.startswith('g'):
    print 'You can go.'
  case color == 'yellow':
    print 'Slow down!'
  default:
    print 'You better stop.'

```
/3 Just like the switch statement in Go.  (There is no equivalent in Python.)

Any expression can appear after the word "case", and operator == is used to compare
the guard value with the switch value.

But if the value after "switch" is omitted, the first case whose guard
evaluates to a true value is executed.

Cases are tried in the order listed.  If none match, and there is a "default"
clause, it is executed.

/1 native
/2
```
native: `
  type Page struct { Title, Body string }
`

def NewPage(title, body):
  z = None
  native: `
    v_z = MkGo(&Page{
        Title: a_title.Self.Str(),
        Body: a_body.Self.Str(),
    })
  `
  return z
```
/3 Embed go code directly in a Rye program.

In some cases, Rye may not be able to do exactly
what you need to do, to satisfy some other Go library.

How to do this is not documented and is subject to change,
but you can find some examples in the Rye distribution.
You can observe 4 things in the example: Arguments have
`a_` prepended, other local variables have `v_` prepended,
you can gently get the string value of a str-like Rye object
with `.Self.Str()`,
and MkGo(x) converts any Go value into a reflective
Rye value referencing that Go value.

/1 del
/2
```
del d[k]
```
/3 As in Python.

/1 return
/2
```
return     # Returns None.
return x+y
return first, second, third
```
/3 As in Python.

/1 yield
/2
```
yield     # Yields None.
yield x+y
yield first, second, third
```
/3 As in Python.

/1 break
/2
```
break
```
/3 As in Python.


/1 continue
/2
```
continue
```
/3 As in Python.

/1 pass
/2
```
pass
```
/3 As in Python.

