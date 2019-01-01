from go import "math/big", "crypto/rand"

TRIALS = 10000
N = 6  # Six digit numbers
histo2 = {}  # Frequency of N-duplications
m = big.NewInt(int(N * '9'))
for i in xrange(TRIALS):
  s = rand.Int(rand.Reader, m).String()
  s = (6 - len(s)) * '0' + s
  histo1 = {}  # Frequency of digits in s
  for c in s:
    histo1[c] = histo1.get(c, 0) + 1
  z = max(histo1.values())
  histo2[z] = histo2.get(z, 0) + 1

count = reduce((lambda a, b: a+b), histo2.values(), 0)
for k, v in sorted(histo2.items()):
  print k, v, count, '%12.3f' % (100.0 * v / count)
