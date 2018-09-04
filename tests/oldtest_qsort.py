from go import math/rand as mathrand

def qsort(vec):
  n = len(vec)
  if n < 2:
    return
  i, j = 0, n - 1
  pivot = vec[n/2]

  while i <= j:
    while vec[i] < pivot and i <= j:
      i += 1
    while pivot < vec[j] and i <= j:
      j -= 1
    if i >= j:
      break
    vec[i], vec[j] = vec[j], vec[i]
    i += 1
    j -= 1

  left, right = vec[:i], vec[i:]
  qsort(left), qsort(right)

for trial in range(100):
  v = [mathrand.Intn(10) for i in range(30)]
  e = sorted(v)
  qsort(v)
  print trial, v
  assert v == e
