## Rye Data Types

/0

/1 None
/2
None
assert type(None) == None

/3
As in python, there is a single instance of this type, named None.

None may also represent nil values returned from Go functions,
as nil pointers, nil intefaces, and nil slices.

/1 int
/2
```
print 12345 + 0xFF
```
/3 Always int64 as in Go (not bignums as in Python).

/1 str
/2
```
s = 'Counting 하나 둘 셋\n'
print len(s)
print repr(s[7]), repr(s[9])
print [c for c in s]
```
/3
Immutable string, as in Python.  Use UTF-8 if unicode is desired.

If you ask len(), you get the byte length of the string (as in Go).

If you index them, it uses byte offsets, and you get single-byte strings (as in Go).

If you iterate over them, it treats them as unicode runes (as in Go), and you may get a multi-byte char for some elements (as in the four korean runes in the demo).

/1 float
/2
```
print 2.5 * math.Pi
```
/3 Always float64, as in both Python & Go.

/1 list
/2
```
['a', None, [], 5]
```
/3 As in python: Variable length, mutable elements, appendable.

/1 dict
/2
```
{'color': red, 'size': 55}
```
/3
As in python, except the keys should be type str.
If they keys are not str, they are converted to str, using str().

/1 tuple
/2
```
('a', None, (), 5)
```
/3 As in python.  Fixed length, immutable elements.

Go functions that return multiple values will return types to Rye.

/1 byt
/2
```
x = byt('lorem ipsum') + mkbyt(1024)
```
/3
A special kind of python list,
where the elements must be of type int  and must be in the range 0 to 255.

Somewhat like []byte in go.  Mutable elements, but appendable.

Converts easily to and from either str or list.  
Although it is more like a list of int than a python str, it is often used
to hold string data.  
In go, many interfaces use mutable []byte arguments; the byt type this allows compatability.

A special builtin function mkbyt(n) creates a new byt with n zero elements.

/1 set
/2
```
{'nutella', 'tiramisu', 'stracciatella'}
```
/3 Like set in python.

/1 object
/2
```
class Foo(object):
  def __init__(self, x)
    self.x = x
  def bar(self)
    return 1000 + self.x

print Foo(23).bar()
```
/3 

All classes in Rye derive implicitly or explicitly from object.

Only single inheritance is supported.

/1 Go Values
/2
```
b = bytes.NewBuffer(None)
b.append("Hello Go World!")
print b.String()
```
/3
Values created in Go can be usually be used transparently as Rye values.

Public fields and methods can be called on them.

There are some special builtin functions for dealing with the finer points
of Go types and values.  
You can also include native go code to interface to tricker things in Go.

