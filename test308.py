s = 0
for i in range(100):
  for j in range(100):
    for k in range(100):
      n = i * 10000 + j * 100 + k + 1
      s = s + n
print s, 'Woot!'
