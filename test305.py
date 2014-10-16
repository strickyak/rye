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
  # def __init__(self, x, y, *v, **kw):
  def __init__(self, x, y):
    self.x = x
    self.y = y
    # self.v = v
    # self.kw = kw
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
du = Durian("zero", "one")
foo = du.Foo
print foo()

class Base:
  def __init__(self):
    pass

  def Public(self, x):
    return self.Protected(x)

class Imp(Base):
  def __init__(self):
    pass
  def Protected(self, x):
    return x*x

assert Imp().Public(8) == 64

def ArgCount(*args, **kw):
  return len(args), len(kw)
assert ArgCount(5, 6, 7, 8, abcd=1, defg=2) == (4, 2)
assert ArgCount(5, 6, pqr=1, *[7, 8], **{'xyz':2}) == (4, 2)

##class Formatted:
##  def __init__(self, *args, **kw):
##    self.args = args
##    self.kw = kw
##  def Thus(self, fmt):
##    return fmt % [self.args + self.kw.keys()]
##print Formatted(11, 22, color=888).Thus("eleven %d twenty-two %d color %s")

pass
