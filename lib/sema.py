"""Data Structures safe for use across goroutines."""
from go import sync

class Shared:
  """A shared variable, you can Get() and Set(x)."""
  def __init__(x=None):
    .x = x
    .mu = go_new(sync.Mutex)
  def Get():
    .mu.Lock()
    z = .x
    .mu.Unlock()
    return z
  def Set(x):
    .mu.Lock()
    .x = x
    .mu.Unlock()

class Serial:
  """A serial number generator, starting at i=1."""
  def __init__(i=1):
    .i = i
    .mu = go_new(sync.Mutex)
  def Take():
    return .Next()
  def Next():
    .mu.Lock()
    z = .i
    .i += 1
    .mu.Unlock()
    return z

pass
