## Using Go from Rye

/0

/1 Calling Go functions
/2
```
from go import os, strings as S, time
from go import path/filepath

home = os.Getenv('HOME')
homes = filepath.Dir(home)

t = time.Now()  # An instance of Go type time.Time
yyyy_mm_dd = S.Split(str(t), ' ')[0]
```
/3
To call go functions from Rye, import the package
using "from go import".  The final portion of the
name (like "filepath" of "path/filepath") is the
qualifier name to use in Rye, unless a new qualifier name
is specified (like "S" in "strings as S").

Notice that often you do not have to think much
about the types in Go; things work naturally when they can.  
Because `time.Now()` returns a type `time.Time` that has a
`String()` method, we can convert it into a string using `str()`.

/1 Calling Go methods
/2
```
t = time.Now()  # An instance of Go type time.Time
unix_secs = t.Unix()  # Seconds since the epoch. 
iso_now = t.Format('2006-01-02T15:04:05')
email_now = t.Format(time.RFC822)

MATCH_VAR_NAME = regexp.MustCompile(
   `^[_A-Za-z][_A-Za-z0-9]*$`).FindStringIndex

if MATCH_VAR_NAME(varname):
  print "%s is a variable." % varname
```
/3 
Calling methods on Go objects looks just like calling a method
on a Python object.

`MATCH_VAR_NAME` is an uncalled method bound to a receiver
object; it can be called later as many times as you like
(as in Python, and in Go version 1.3 and beyond).

Since `FindStringIndex()` returns a slice that is empty when
there is no match, the `if` predicate will be false when
there is no match (empty slices are false).

/1 Return values (from Go) and Errors

/2
```
from go import os

# In Go, Clearenv returns no result.
a = os.Clearenv()
assert a is None  # In Rye, it returns None.

# In Go, Chdir returns error.
b = os.Chdir('/')
# In Rye the error is removed, so it returns None.
assert b is None

# In Rye, if a final Go return value is of type error
# is not nil, it is automatically thrown.
got_exception = False
try:
  c = os.Chdir('/This/Path/Does/Not/Exist')
except Exception as ex:
  got_exception = True
assert got_exception

# In Go, LookupEnv returns two (non-error) values:
#     func LookupEnv(key string) (string, bool)
d, ok = os.LookEnv('HOME')
assert ok and type(d) is str
```

/3
If the final return value of a Go function or method
is of type `error`, it is removed from the return type
as far as Rye is concerned, and it is automatically tested
and raised as an exception if it is not nil.
This is a convenience to the programmer, and consistent
with Python's philosophy of using Exceptions (rather than
error return values) to indicate errors.

If a Go function or method returns no value,
in Rye it returns None.

If a Go function or method returns more than one value,
in Rye it returns a tuple of values.





