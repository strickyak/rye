import sys  # rye_pragma from "github.com/strickyak/rye/emulation"
from go import "os/exec" as X
from go import sort, strings

def double(x):
  return x + x
print double("beep")
print double(123)

class Guava:
  def __init__(self):
    self.foo = 100
  def Foo(self):
    print >>sys.stderr, 'rye_stack:\n%s;' % rye_stack()
    return self.foo

class Durian:
  def __init__(self, x, y):
    self.x = x
    self.y = y
  def Foo(self):
    print >>sys.stderr, 'rye_stack:\n%s;' % rye_stack()
    return self.x + self.y

class Jackfruit(Durian):
  def __init__(self, x, y, z):
    super(x, y)
    self.z = z

class Rambutan(Jackfruit):
  def Foo(self):
    return self.x * self.y * self.z

must Rambutan(10, 20, 30).Foo() == 6000

print repr(rye_pickle(Durian(33, 44)))
print rye_unpickle(rye_pickle(Durian(33, 44)))
print repr(rye_pickle(Durian("abc", "xyz")))
print repr(rye_pickle(Jackfruit("abc", "xyz", Guava())))
print rye_unpickle(rye_pickle(Jackfruit("abc", "xyz", Guava())))

g = setattrs(Guava(), foo=404)
must g.foo == 404
j = setattrs(Jackfruit(0, 0, 0), x=11, y=22, z=33)
must (j.x, j.y, j.z) == (11, 22, 33)

cmd = setattrs(go_new(X.Cmd), Path= X.LookPath('expr'), Args= ['expr', '22', '*', '4'] )
out = cmd.Output()
say out
must strings.TrimRight(out, '\r\n') == '88'

ss = go_make(sort.StringSlice, 0)
ss = go_append(ss, 'one')
ss = go_append(ss, 'two')
ss = go_append(ss, 'three')
must len(ss) == 3
must ','.join(ss) == 'one,two,three'

################
d0 = dict()
d1 = dict(a=100, b=200, c=300)
d2 = {'x':600, 'y':700, 'z':800}
s0 = set()
s1 = set(['q', 'u'])
s2 = {'p', 'q', 'r'}

must s0 & s0 == set()
must s0 & s1 == set()
must s0 & s2 == set()
must s1 & s0 == set()
must s1 & s1 == s1
must s1 & s2 == {'q'}
must s2 & s0 == set()
must s2 & s1 == {'q'}
must s2 & s2 == s2

must s0 | s0 == set()
must s0 | s1 == s1
must s0 | s2 == s2
must s1 | s0 == s1
must s1 | s1 == s1
must s1 | s2 == {'u', 'r', 'p', 'q'}
must s2 | s0 == s2
must s2 | s1 == {'u', 'r', 'p', 'q'}
must s2 | s2 == s2

must s0 ^ s0 == set()
must s0 ^ s1 == s1
must s0 ^ s2 == s2
must s1 ^ s0 == s1
must s1 ^ s1 == set()
must s1 ^ s2 == {'u', 'r', 'p'}
must s2 ^ s0 == s2
must s2 ^ s1 == {'u', 'r', 'p'}
must s2 ^ s2 == set()

must s0 - s0 == set()
must s0 - s1 == set()
must s0 - s2 == set()
must s1 - s0 == s1
must s1 - s1 == set()
must s1 - s2 == {'u'}
must s2 - s0 == s2
must s2 - s1 == {'r', 'p'}
must s2 - s2 == set()

must (s0 <= s0) == True
must (s0 <= s1) == True
must (s0 <= s2) == True
must (s1 <= s0) == False
must (s1 <= s1) == True
must (s1 <= s2) == False
must (s2 <= s0) == False
must (s2 <= s1) == False
must (s2 <= s2) == True

must (s0 < s0) == False
must (s0 < s1) == True
must (s0 < s2) == True
must (s1 < s0) == False
must (s1 < s1) == False
must (s1 < s2) == False
must (s2 < s0) == False
must (s2 < s1) == False
must (s2 < s2) == False

must (s0 >= s0) == True
must (s0 >= s1) == False
must (s0 >= s2) == False
must (s1 >= s0) == True
must (s1 >= s1) == True
must (s1 >= s2) == False
must (s2 >= s0) == True
must (s2 >= s1) == False
must (s2 >= s2) == True

must (s0 > s0) == False
must (s0 > s1) == False
must (s0 > s2) == False
must (s1 > s0) == True
must (s1 > s1) == False
must (s1 > s2) == False
must (s2 > s0) == True
must (s2 > s1) == False
must (s2 > s2) == False

say d0
say d1
say d2
say s0
say s1
say s2

must d0 == d0
must d0 != d1
must d0 != d2
must d1 != d0
must d1 == d1
must d1 != d2
must d2 != d0
must d2 != d1
must d2 == d2

must s0 == s0
must s0 != s1
must s0 != s2
must s1 != s0
must s1 == s1
must s1 != s2
must s2 != s0
must s2 != s1
must s2 == s2

must s0 == rye_unpickle(rye_pickle(s0))
must s1 == rye_unpickle(rye_pickle(s1))
must s2 == rye_unpickle(rye_pickle(s2))
must d0 == rye_unpickle(rye_pickle(d0))
must d1 == rye_unpickle(rye_pickle(d1))
must d2 == rye_unpickle(rye_pickle(d2))

t1 = (3, 5, 7)
must t1 < (3, 5, 8)
must t1 < (3, 6, 7)
must t1 < (4, 5, 7)
must (3, 5, 8) > t1
must (3, 5, 7) == rye_unpickle(rye_pickle(t1))

v1 = [3, 5, 7]
must v1 < [3, 5, 8]
must v1 < [3, 6, 7]
must v1 < [4, 5, 7]
must [3, 5, 8] > v1
must [3, 5, 7] == rye_unpickle(rye_pickle(v1))

must ("one", 2, 3.1415) == rye_unpickle(rye_pickle(("one", 2, 3.1415)))
must (False, True, None, (), [], {}, set()) == rye_unpickle(rye_pickle((False, True, None, (), [], {}, set())))
must ("one", byt("two")) == rye_unpickle(rye_pickle(("one", byt("two"))))

must 120 == setattrs(Durian(0,0), x=100, y=20).Foo()
must 880 == setattrs(Durian(800,80)).Foo()
say Durian(0,0)
say Durian(800,80)
say Guava()
say Jackfruit(800,80, 1000)
say Rambutan(800,80, 1000)

################

print "402 OKAY."
