import sys  # rye_pragma from "github.com/strickyak/rye/emulation"

class Buffer(object):
  def __init__(self):
    self.b = []
  def write(self, x):
    self.b.append(str(x))
  def flush(self):
    pass
  def __str__(self):
    z = ''.join(self.b)
    return z

print "ONE"
print >>sys.stderr, "TWO"
save_out = sys.stdout
sys.stdout = Buffer()
for i in range(10):
  print i, i*i, i*i*i
  say i, i*i, i*i*i
  say sys.stdout.b

text = str(sys.stdout)
print >>save_out, text
say len(text)
say text
