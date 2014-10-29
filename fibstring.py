def fib(s):
  if len(s) < 2:
    return s
  else:
    return fib(s[1:]) + fib(s[2:])

print len(fib(35 * 'x'))
