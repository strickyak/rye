"""Test the time implementation from emulation.

This should work with either python or rye.
"""
import time  # rye_pragma from "github.com/strickyak/rye/emulation"

a = time.time()
print sum([x for x in range(1000)])  # Faster than 1 second.
b = time.time()

assert b > a
assert b - a < 1.0  # less than a second.
