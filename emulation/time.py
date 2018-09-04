from go import time as gotime

# time() returns seconds since the Unix Epoch as a float, with subsecond precision.
def time():
  return float(gotime.Now().UnixNano()) / 1000000000.0
