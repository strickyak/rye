s = 0
for i in range(100):
  for j in range(100):
    for k in range(100):
      n = i * 10000 + j * 100 + k + 1
      s = s + n
print s, 'Woot!'

print ()
print (100,)
print (100, 200, 300)
print (100, 200, 300, 400,)
print "%d plus %s equals %f" % (100, "200", 300.0)
