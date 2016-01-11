## Rye Expressions

/0

/1 Literals
/2
/3 See [Literals](/literals).

/1 self & super
/2
```
self.field += 1

self.method(a, b, c)

super.method(a, b, c)

class X:
  def __init__(self, x, y):
    super.__init__(x)
    self.y = y
```
/3
`self` is like in Python.

`super` is like in Java, and is simpler than Python's `super`.
A method can use `super` as the receiver to call a method at the parent level of the hierarchy.

/1 Variables
/2
```
A
a
a0
CamelCase
lower_case
__name__
```
/3 As in Python.

/1 Qualified Names
/2
```
from go import os, math/big, html/template
from go import github.com/russross/blackfriday

fd = os.Open('/dev/null')
x = big.NewInt(42)
h = blackfriday.MarkdownCommon(s)
return go_cast(template.HTML, h)
```
/3
The final portions of the imported paths become qualifer names,
like `os`, `big`, `template`, and `blackfriday` in the example.

Qualfier names cannot be used on their own;  there are no variables
`os`, `big`, etc, in the example.

Instead, they must be followed by a dot and another name,
like `os.Open`, `big.NewInt`, `blackfriday.MarkdownCommon`,
and `template.HTML` (which names a type in Go).

/1 Calls
/2
```
def Twice(x): x + x
class Grape:
  def Print(self):
    print "I am a Grape."

print Twice('foo')
g = Grape()
g.Print()
```
/3 As in Python.

/1 Comprehensions
/2
```
squares = [x*x for x in range(5)]
assert squares == [0, 1, 4, 9, 16]
```
/3 As in Python.

Only the list comprehension inside [square brackets] is supported.
It always evaluates to a list.

Dictionary comprehensions & generating comprehensions are NOT supported.

/1 Indexing & Slicing
/2
```
w = a[i]
a[i] = w + 1
x = b[i:j]
y = b[i:]
z = b[:-1]
```
/3
As in Python, you can index a `list`, `dict`, `tuple`, or a `str`.
In Rye you can also index a `byt` (because it acts like a list of int).
You can access a value, or change a value with assignment, using an index.

As in Python, you can slice a `list` or a `str` (and in Rye, a `byt`).
But you cannot assign to a slice, and you cannot use a slice with two colons
(specifying a stride).

Negative indices work just as in Python.


/1 go
/2
```
go f(a, b, c)
```
/3 As in Go.

A go expression is the keyword `go` followed by a function or
method call.  The receiver (if it is a method call) and the arguments
are all evaluated first synchronously.   Then a background goroutine
is created to execute the call.

TODO -- document the promise in Rye.

/1 OPERATORS
/2
/3
The following are the layers of operators
and their binding strength, from binding the strongest
to the weakest.

/1 unary operators
/2
```
x = +a
y = -b  # Negate.
z = ~c  # Invert.
```
/3 As in Python.

/1 multiplicative operators
/2
```
a = b * c
a = b / c
a = b // c
a = b % c
```
/3 As in Python.

/1 additive operators
/2
```
a = b + c
a = b - c
```
/3 As in Python.

/1 shift operators
/2
```
a = b << c
a = b >> c
a = b >>> c
```
/3
The bit shift operators operate on int, which are internally 64-bit signed twos-complement integers.

The `<<` is an signed shift left.

The `>>` is an signed shift right (as in Java).

The `>>>` is an unsigned shift right (as in Java).

/1 Bitwise and
/2
```
a = b & c
```
/3 As in Python (using int64).

/1 Bitwise xor
/2
```
a = b ^ c
```
/3 As in Python (using int64).

/1 Bitwise or
/2
```
a = b | c
```
/3 As in Python (using int64).

/1 Relational operators
/2
```
a == b
a != b
a < b
a >= b
a > b
a <= b
0 <= a < x
x in y
x not in y
z is None
z is not None
```
/3 As in Python.

/1 not
/2
```
not p
```
/3 As in Python.

/1 and
/2
```
p and q and r
```
/3 As in Python:
returns the first false value, or else True.

Evaluation stops when it finds a false value.


/1 or
/2
```
p or q or r
```
/3 As in Python:
returns the first true value, or else False.

Evaluation stops when it finds a true value.

/1 conditional
/2
```
x = a if p else b
```
/3 As in Python.

/1 lambda
/2
```
fn = lambda x, y: y if y < x else x
```
/3 As in Python.

