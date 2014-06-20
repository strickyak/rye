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
