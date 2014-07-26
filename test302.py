print "HELLO"
print "WORLD"
print "HELLO" + "WORLD"

assert [1, 2, 3] == [1, 2, 3]
assert [1, 2, 3] != [1, 2, 0]
assert [1, 2, 3] != [1, 2, 3, 4]
assert [1, 2, 3, 4] != [1, 2, 3]

assert [1, 2, 3] <= [1, 2, 4]
assert [1, 2, 3] <= [1, 2, 3, 0]

assert [1, 2, 3] >= [1, 2, 2]
assert [1, 2, 3, 0] >= [1, 2, 3]

assert False == False
assert False == 0
assert False == 0.0

assert True == True
assert True == 1
assert True == 1.0

assert False < True
assert False < 1
assert False < 1.0
