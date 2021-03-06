from go import bytes
from go import strings as foo
from go import strconv
from go import fmt as F
from go import io
from go import net.http
from go import reflect as R

import twice as Doppel

######### Test the "must expect" mechanism.
must except [5][1] == 5
must except {'abc': 1}['def']
must len([888])
must len([]) == 0
must except len(3.1415) == 555

gotExpectedException = False
try:
  must except len('Taking len of string is legal, no exception!')
except:
  gotExpectedException = True
must gotExpectedException
##############################

print foo.Contains("Hello", "ell"), Doppel.Twice(21)

print strconv.ParseBool("true")
print strconv.Atoi("1234")

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

must (-1 >> 3) < 0
must (-1 >> 3) == -1
must (-1 >>> 3) > 0
must (-1 >>> 3) == 2305843009213693951

def AddToCat(n):
  global cat
  cat += n

def Lion(x):
  global cat
  try:
    must cat == 100
    cat += x
    must cat == 110
  finally: AddToCat(x)
  must cat == 120

def Tiger():
  global cat
  cat += 100
  Lion(10)
  must cat == 120

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

must [ 'color', 'size' ] == sorted(hdr.keys()[:])
z = dict(hdr.items()[:])

i = 0
for k in z:
  v = z[k]
  v = v[0:len(v)]
  if k == 'color':
    must list(v) == [ 'purple', 'violet' ]
    i += 1
  if k == 'size':
    must list(v) == [ 'XXL', 'XL' ]
    i += 1
must i == 2

#BUG in OptimizedGoCall?
#def Tri(n):
#  return n if n<2 else n+(go Tri(n-1)).Wait()

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
say s
say len(s)
say [c for c in s]
say [c for c in byt(s)]
must len(s) == 14
must len([c for c in s]) == 6
must len([c for c in byt(s)]) == 14

must list(byt("ABC")) == [65, 66, 67]
must list(byt("ABC")) == [ord(c) for c in "ABC"]

must s[0:6].isalpha()
must s[0:6].isalnum()
must not s[0:7].isalpha()
must s[6:7].isspace()

d = "๒๑"  # Thai "21".
print [c for c in d]
print [c for c in byt(d)]
must len([c for c in d]) == 2
must len([c for c in byt(d)]) == 6
# must d.isdigit() # Fails.
must d.isalnum()

b = mkbyt(1)
say b
b[0] = 237
say b[0]
must b[0] == 237
say 3*b
say byt(3 * [237])
must (3 * b) == byt(3 * [237])
say b
must list(3 * b) == (3 * [237])
say b
must (3*b)[2] == 237
say b
must len(6*mkbyt(7)) == 42
say b

#########
bb = go_new(bytes.Buffer)
print >>bb, 'Hello bytes.Buffer'
print >>bb, 'Bye.'
must bb.String() == 'Hello bytes.Buffer\nBye.\n'
#########
r, w = io.Pipe()
rbuf = mkbyt(5)
promise = go r.Read(rbuf)
print >>w, 'Pipe'
must promise.Wait() == len('Pipe\n')
must str(rbuf) == 'Pipe\n'
#########

x = []
for a in ['red', 'white', 'blue', 404, None]:
  switch a:
    case 'blue':
      x.append('B')
    case 'white':
      x.append('W')
    case 'red':
      x.append('R')
    case None:
      x.append('N')
    default:
      x.append('d')
must x == ['R', 'W', 'B', 'd', 'N']

#########

must {} == dict(None)
must {} == dict(list(None))

#########
abc = 100
try:
  abc = strconv.Atoi("200")
except:
  abc += 1
must 200 == abc

abc = 100
try:
  abc = strconv.Atoi("200")
except:
  abc += 1
finally:
  abc += 10
must 210 == abc

abc = 100
try:
  print strconv.Atoi("23skidoo")
  abc = 200
except:
  abc += 1
must 101 == abc

abc = 100
try:
  print strconv.Atoi("23skidoo")
  abc = 200
except:
  abc += 1
finally:
  abc += 10
must 111 == abc


abc = 100
try:
  abc = strconv.Atoi("200")
finally:
  abc += 10
must 210 == abc

try:
  abc = 100
  try:
    abc = strconv.Atoi("23skidoo")
    abc = 200
  finally:
    abc += 10
except:
  abc += 1000
must 1110 == abc

#########

if a, b, c = 100, 200, 300:
  must (a, b, c == 100, 200, 300)
else:
  must False

if a, b, c = '':
  must False
else:
  pass

if a, (b1, b2), c = 100, [200, 201], 300:
  must (a, (b1, b2), c == 100, [200, 201], 300)
else:
  must False

#########

thing = object()
say thing
def Gen():
  yield 10
  yield 20
  raise thing
def Consume():
  try:
    for x in Gen():
      say x
    must False
  except as ex:
    say ex
    must ex is thing
    return True
must Consume()

#########

print "401 OKAY."
