import sys

class X:
  pass

h = {}
for line in sys.stdin:
  x = X()
  for k, v in eval(line)[0].items():
    setattr(x, k, v)

  if x.what=='time' and x.exit==0 and x.cpus==2:
    weeks = float(x.secs) / 86400 / 7
    d = h.get(weeks, [])
    d.append(x.real)
    h[weeks] = d

for weeks, vec in sorted(h.items()):
  mean = 1.0 * sum(vec) / len(vec)
  print weeks, mean, min(vec), max(vec)
pass
