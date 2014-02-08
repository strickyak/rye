def double(x):
  return x + x
print double("beep")
pass

class Guava:
  def __init__(self):
    self.foo = 100
  def Foo(self):
    return self.foo
print Guava().foo
print Guava().Foo()
pass

class Durian:
  def __init__(self, x, y):
    self.x = x
    self.y = y
  def Foo(self):
    return self.x + self.y
    pass
print Durian(3, 4).x
print Durian(3, 4).y
print Durian(3, 4).Foo()
print Durian(33, 44).x
print Durian(33, 44).y
print Durian(33, 44).Foo() + 1
print Durian("foo", "bar").x
print Durian("foo", "bar").y
print Durian("foo", "bar").Foo() + "1"
pass
