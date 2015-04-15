from go import os/exec as X
from go import sort, strings

def double(x):
  return x + x
print double("beep")
print double(123)

class Guava:
  def __init__(self):
    self.foo = 100
  def Foo(self):
    return self.foo

print rye_pickle(Guava())

class Durian:
  def __init__(self, x, y):
    self.x = x
    self.y = y
  def Foo(self):
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

g = Guava() {foo: 404}
must g.foo == 404
j = Jackfruit(0, 0, 0) {x: 11, y:22, z:33}
must (j.x, j.y, j.z) == (11, 22, 33)

cmd = go_new(X.Cmd) { Path: X.LookPath('expr'), Args: ['expr', '22', '*', '4'] }
out = cmd.Output()
say out
must strings.TrimRight(out, '\r\n') == '88'

ss = go_make(sort.StringSlice, 0)
ss = go_append(ss, 'one')
ss = go_append(ss, 'two')
ss = go_append(ss, 'three')
must len(ss) == 3
must ','.join(ss) == 'one,two,three'

print "402 OKAY."
