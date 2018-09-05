def plus1(a, b : int) ->int :
  return a + b + len("!")

def times10(c : int, d : int) ->int :
  return c * d * 10

def idouble(a : int) ->int :
  return a+a

x = plus1(times10(5, 6), 100)
print x
y = idouble(33) + 34
print y
