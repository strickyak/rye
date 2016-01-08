print "HELLO"
print "WORLD"
print "HELLO" + "WORLD"
print 1000
print 33
print 967 + 33
print 'a', 'nando', 'a' in 'nando'
print 'a', 'frodo', 'a' in 'frodo'

assert "HELLO" == "HEL" + "LO"
assert "WORLD" == "W%sD" % "ORL"
assert "HELLO" + " WORLD" == "%s %s" % ("HELLO", "WORLD")
assert 1000 == 10 * 10 * 5 * 2
assert 33 == (1 << 5) + 1
assert 967 + 33 == int('1' + 3 * '0')

assert ('a' in 'nando') == True
assert ('a' in 'frodo') == False

assert ~0 == -1
assert ~1 == -2
assert ~2 == -3
assert ~-3 == +2
assert -2 == 4 - 6
assert +23 == 30 - 7

assert (-1 >> 3) < 0
assert (-1 >> 3) == -1
