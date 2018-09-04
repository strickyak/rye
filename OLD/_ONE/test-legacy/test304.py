def double(x):
  return x + x
print double(21)
print double("beep")

def f(x):
  def g(y):
    return x + y
  return g

assert f(30)(15) == 45
