def TriangleWithRecursion(n):
  if n < 1:
    return 0
  return n + TriangleWithRecursion(n-1)

print TriangleWithRecursion(6)

def TriangleWithWhile(n):
  z = 0
  while n > 0:
    z = z + n
    n = n - 1
  return z

print TriangleWithWhile(6)
