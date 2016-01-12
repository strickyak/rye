## Rye Literals

/0

/1 None, True, False

/2
```
a = None
b = True
c = False
```
/3 As in Python.

/1 int
/2
```
a = 42
b = -100
c = 0xFF00  # Hex.
d = 0755    # Octal.
```
/3 As in Python.

/1 float
/2
```
a = 0.0
neg_pi = -3.14159265359
mole = 6.02214129E+23
```
/3 TODO -- describe exact rules.
```
RE_FLOAT = re.compile('[+-]?[0-9]+[.][-+0-9eE]*')
```
TODO -- How to deal with infinities, negative zero, not a number?

/1 str
/2
```
a = "Hello " + 'world\n'
b = """This string includes '''."""
c = '''Multi
Line
String'''
d = `func Hello() {
  println("Hello", 'world\n')
}
`
```
/3
Single and Triple `"` and `'` quotes are like in Python.
The raw forms with `R"` and `R'` are NOT supported.

Instead, you can use verbatim quotes with \` as in Go.
Backslash escapes are not resolved in  \` quotes;
they are resolved in `"` and `'` quotes, single or triple.

Notice the Python2 meaning of \` (shell evaluation) is not supported.

/1 tuple
/2
```
a = ()
b = (8,)
c = (3, 4, 5)
d = tuple([3, 4, 5])
```
/3 As in Python.

/1 list
/2
```
a = []
b = [8]
c = [3, 4, 5]
d = list((3, 4, 5))
```
/3 As in Python.


/1 dict
/2
```
a = {}
b = {'color': 'red', 'size': 50}
c = dict(color='red', size=50)
assert 'red' == b['color']
assert 'red' == c.get('color')
d = {100: 2, 1000: 3}  # Keys will convert to str.
assert 2 == d[100]     # Keys will convert to str.
assert 2 == d['100']
```
/3 As in Python.

Keys are always strings in Rye dicts.
(In Python they can be many other types.)
If you use other types as keys, they are converted to
strings using str().


/1 set
/2
```
a = set({})
b = {'color', 'size'}
c = {100, 1000}      # Members will convert to str.
assert 100 in c      # Members will convert to str.
assert '100' in c
```
/3 As in Python.

Members of sets are always strings in Rye sets.
(In Python they can be many other types.)
If you use other types as members, they are converted to
strings using str().


