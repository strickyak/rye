from go import strings as foo
from go import strconv
from go import fmt as F
from go import net/http
from go import reflect as R

import github.com/strickyak/rye/twice as Doppel

print foo.Contains("Hello", "ell"), Doppel.Twice(21)

print strconv.ParseBool("true")
print strconv.Atoi("1234")

try:
  print strconv.Atoi("23skidoo")
except:
  print 'Expected.'

print strconv.UnquoteChar("\\'foo", 39)
print F.Sprintf("%d %d %d", 111, 222, 333)

plus300 = Doppel.Adder(300)
print plus300.Plus(69)

class Augmentor(Doppel.Adder):
  def __init__(self, some, more):
    super(some + more)

aug123 = Augmentor(100, 23)
print aug123.Plus(4000)
print aug123

class Augmentarian(Augmentor):
  def __init__(self, some, even, more):
    super(some + more, even)

aug246 = Augmentarian(200, 40, 6)
print aug246.Plus(8000)
print aug246

class DoubleAugmentarian(Augmentor):
  def __init__(self, some, even, more):
    super(some + more, even)
  def Plus(self, x):
    return 2 * super.Plus(x)

da246 = DoubleAugmentarian(200, 40, 6)
print da246.Plus(8000)
print da246

assert (-1 >> 3) < 0
assert (-1 >> 3) == -1
assert (-1 >>> 3) > 0
assert (-1 >>> 3) == 2305843009213693951

def AddToCat(n):
  global cat
  cat += n

def Lion(x):
  global cat
  with defer AddToCat(x):
    assert cat == 100
    cat += x
    assert cat == 110
  assert cat == 120

def Tiger():
  global cat
  cat += 100
  Lion(10)
  assert cat == 120

cat = 0
Tiger()

hdr_t = R.TypeOf(go_new(http.Header)).Elem()
say hdr_t.Name()
hdr = R.MakeMap(hdr_t).Interface()
say hdr
say hdr.keys()
say hdr.items()
hdr['color'] = [ 'purple', 'violet' ]
hdr['size'] = [ 'XXL', 'XL' ]
say hdr
say hdr.keys()
say hdr.items()

assert [ 'color', 'size' ] == sorted(hdr.keys()[:])
z = dict(hdr.items()[:])

i = 0
for k in z:
  v = z[k]
  v = v[0:len(v)]
  if k == 'color':
    assert list(v) == [ 'purple', 'violet' ]
    i += 1
  if k == 'size':
    assert list(v) == [ 'XXL', 'XL' ]
    i += 1
assert i == 2

def Tri(n):
  return n if n<2 else n+Tri(n-1)

promises = [go Tri(n) for n in range(10)]
say promises

z = 0
for p in promises:
  x = p.Wait()
  say x
  z += x
print 'Sum', z

s = "하나 둘 셋"  # Hangul "hana tul ses".
print [c for c in s]
print [c for c in byt(s)]
assert len(s) == 14
assert len([c for c in s]) == 6
assert len([c for c in byt(s)]) == 14

assert s[0:6].isalpha()
assert s[0:6].isalnum()
assert not s[0:7].isalpha()
assert s[6:7].isspace()

d = "๒๑"  # Thai "21".
print [c for c in d]
print [c for c in byt(d)]
print len([c for c in d]) == 2
print len([c for c in byt(d)]) == 6
# assert d.isdigit() # Fails.
assert d.isalnum()

b = byt(1)
b[0] = 237
# assert not str(b).isalpha() # fails.

print "401 OKAY."
