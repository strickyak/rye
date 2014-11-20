from go import fmt, os, runtime

def print_tb(ignore):
  b = mkbyt(10000)
  n = runtime.Stack(b)
  b = b[:n]
  fmt.Sprintf(os.Stderr, "%s\n", str(b))
