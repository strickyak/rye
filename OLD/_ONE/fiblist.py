def fib(s):
  if len(s) < 2:
    return s
  else:
    return fib(s[1:]) + fib(s[2:])

print len(fib([1,1,1,1,1,1]))
print len(fib([1,1,1,1,1,1,1]))
print len(fib([1,1,1,1,1,1,1,1]))
print len(fib([1,1,1,1,1,1,1,1,1]))
print len(fib([1,1,1,1,1,1,1,1,1,1]))

#print len(fib(36 * 'x'))
#print len(fib('xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'))  # 36
print len(fib([ 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, ]))
