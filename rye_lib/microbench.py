from go import fmt
from go import time

def Run(func, seconds=1.0):
  """Return nanoseconds per run, attempting to run at least seconds."""
  native: `
    goal := a_seconds.Self.Float()
    start := i_time.Now()

    j := 1  // How many to run next.
    n := 0  // Total runs.
    for {
      fn, ok := a_func.Self.(ICallV)
      if !ok {
        panic(i_fmt.Sprintf("Argument not callable: %s", a_func.Self.String()))
      }
      for i := 0; i < j; i++ {
        _ = fn.CallV(nil, nil, nil, nil)
      }
      n += j
      dur := float64(i_time.Since(start).Nanoseconds()) / 1000000000.0
      // println("microbench n, j, dur", n, j, dur)
      if dur > goal {
        per := dur / float64(n)
        // println("microbench.Run:", dur, n, per)
        return MkFloat(per)
      }
      j += j
    }
  `
