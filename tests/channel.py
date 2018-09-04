from go import io, time
#######################################
# First Send and Close, then Recv.
c = rye_chan(3)
ex = None
say 6, c
try:
  _, _ = c.Recv(timeout=0.001)
except as ex:
  pass
must ex == 'RYE_TIMEOUT'
say 12, c

say 'warming...'
c.Warm()
say 'warmed.'
c.Send(111)
say 15, c
c.Send(222)
c.Send(333)
c.Close()
say 19
x, ok = c.Recv(timeout=0)
say 20, x, ok
must (x, ok) == (111, True)
x, ok = c.Recv(timeout=None)
must (x, ok) == (222, True)
x, ok = c.Recv(timeout=123456789)
must (x, ok) == (333, True)
x, ok = c.Recv()
must (x, ok) == (None, False)
x, ok = c.Recv(timeout=1234567689)
must (x, ok) == (None, False)

#########################################
# A single goroutine as a promise.
foo = 4
def Twelve():
  time.Sleep(10 * time.Millisecond)
  return foo + 7

promise = go Twelve()
foo = 5
must promise.Wait() == 12

ex = None
try:
  promise.Wait()
except as ex:
  pass

must ex == 'EOF'  # io.EOF is String Kind.

#########################################

def Square(n):
  return n * n

promises = []
for i in range(1000):
  p = go Square(i)
  promises.append(p)

total = 0
for p in promises:
  x = p.Wait()
  total += x
must total == sum([x*x for x in range(1000)])
must total == 332833500

#########################################
print 'Testing ENOUGH'
print 'TODO: These are sticking on Send, not getting GCed'

def NaturalNumbers():
  i = 1
  while True:
    yield i
    i += 1

total = 0
for i in range(1000): 
  if i % 100 == 1: say i
  for j in NaturalNumbers():
    if j > 10:
      break
    total += j
must total == 55000
