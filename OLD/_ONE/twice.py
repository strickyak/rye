def Twice(x):
  return x + x

class Adder:
  def __init__(self, augend):
    self.augend = augend

  def Plus(self, x):
    return self.augend + x

print 'Twice 100 is', Twice(100)
print 'Twice "Foo" is', Twice("Foo")

print 'Add 66 to 1000 is', Adder(1000).Plus(66)
