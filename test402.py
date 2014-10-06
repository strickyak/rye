from go import os/exec as X
from go import strings

def double(x):
  return x + x
print double("beep")
print double(123)

class Guava:
  def __init__(self):
    self.foo = 100
  def Foo(self):
    return self.foo

print pickle(Guava())

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

print repr(pickle(Durian(33, 44)))
print unpickle(pickle(Durian(33, 44)))
print repr(pickle(Durian("abc", "xyz")))
print repr(pickle(Jackfruit("abc", "xyz", Guava())))
print unpickle(pickle(Jackfruit("abc", "xyz", Guava())))

g = Guava() {foo: 404}
must g.foo == 404
j = Jackfruit(0, 0, 0) {x: 11, y:22, z:33}
must (j.x, j.y, j.z) == (11, 22, 33)

cmd = gonew(X.Cmd) { Path: X.LookPath('expr'), Args: ['expr', '22', '*', '4'] }
must strings.TrimRight(cmd.Output(), '\r\n') == '88'
