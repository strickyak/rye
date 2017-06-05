from go import strings

class ZZZ(object):
  pass

class AAA:
  pass

class BBB(AAA):
  pass

  def Bar (a ::int, b ::float, c ::str, aaa ::AAA, ret) ::list:
    return ret

def Foo (a ::int, b ::float, c ::str, aaa ::AAA, ret) ::list:
  return ret

def Foo1 (a :int, b :float, c :str, aaa :AAA, ret) ->int :
  return a*a

must Foo(23, 3.14, "foo", AAA(), [5,6,7]) == [5,6,7]
if 't' in rye_opts():
  must except Foo(100, 3.14, byt('foo'), BBB(), [5])
  must except Foo(23, 3.14, None, AAA(), [5])

  must except Foo(3.14, 3.14, "foo", BBB(), [5])
  must except Foo(23, 3.14, "foo", ZZZ(), [5])
  must except Foo(23, 3, "foo", AAA(), [5])
  must except Foo(23, 3.14, (), AAA(), [5])
  must except Foo(23, 3.14, [], AAA(), [5])
  must except Foo(23, 3.14, {}, AAA(), [5])

must Foo(23, 3.14, "foo", AAA(), []) == []
if 't' in rye_opts():
  must except Foo(23, 3.14, "foo", AAA(), (5, 5))
  must except Foo(23, 3.14, "foo", AAA(), {5: 5})
  must except BBB().Bar(23, 3.14, "foo", AAA(), {5: 5})

print "OKAY test_gradtype.py"


#####################################

def Chop(s):
  return strings.Split(s, ',')

def Head(s ::str) ::str :
  return strings.Split(s, ',')[0]

def H(s ::str) ::str :
  return Head(s)

def HH(s ::str) ::str :
  return Head(s) + Head(s)

def Plus(a ::str, b ::str) :
  return a + b

def PlusS(a ::str, b ::str) ::str :
  return a + b

def PlusSI(a ::str, b ::str) ::str :
  def add(a ::str, b ::str) ::str :
    return a + b
  return add(a, b)

def PlusSS(a, b) ::str :
  return a + b

def PlusSSI(a, b) ::str :
  def add(a, b) ::str :
    return a + b
  return add(a, b)

def Len1(x ::list) ::int :
  if x == []:
    return 0
  else:
    return 1 + len(x[1:])

def Len2(x ::list) ::int :
  return 0 if x == [] else 1 + len(x[1:])

assert Plus('abc', 'xyz') == 'abcxyz'
assert PlusS('abc', 'xyz') == 'abcxyz'
assert PlusS(PlusS('a', 'b'), PlusS('c', 'd')) == 'abcd'

assert list(Chop('abc,qrs,xyz')) == ['abc', 'qrs', 'xyz']
assert [x for x in Chop('abc,qrs,xyz')] == ['abc', 'qrs', 'xyz']
assert Head('abc,qrs,xyz') == 'abc'
assert H('abc,qrs,xyz') == 'abc'
assert HH('abc,qrs,xyz') == 'abcabc'
assert PlusS(HH('abc,qrs,xyz'), 'def') == 'abcabcdef'

assert Len1([]) == 0
assert Len1([8]) == 1
assert Len1([8,8]) == 2

assert Len2([]) == 0
assert Len2([8]) == 1
assert Len2([8,8]) == 2

if 't' in rye_opts():
  assert except Len1('hello')
  assert except PlusS(None, None)
  assert except PlusS('abc', 123)
  assert except PlusSI('abc', 123)
  assert except PlusSS(123, 890)
  assert except PlusSSI(123, 890)
assert PlusSS('', '') == ''

def None1(x):
  return x

def None2(x) ::None :
  return x

def None3() ::str :
  pass

assert None1(None) is None
assert None1(5) is 5
assert except None2(5)
assert except None3()

def PlusZ1(a ::str|byt, b ::str|byt) ::str :
  return str(a) + str(b)

def PlusZ2(a ::str|byt?, b ::str|byt?) ::str? :
  if a is None and b is None: return None
  return (str(a) if a is not None else '') + (str(b) if b is not None else '')

assert 'xy' == PlusZ1('x', 'y')
assert 'xy' == PlusZ1(byt('x'), 'y')
assert 'xy' == PlusZ1('x', byt('y'))
assert 'xy' == PlusZ1(byt('x'), byt('y'))
if 't' in rye_opts(): assert except PlusZ1('x', 4)

assert 'xy' == PlusZ2('x', 'y')
assert 'xy' == PlusZ2(byt('x'), 'y')
assert 'xy' == PlusZ2('x', byt('y'))
assert 'xy' == PlusZ2(byt('x'), byt('y'))
if 't' in rye_opts(): assert except PlusZ2('x', 4)
assert 'x' == PlusZ2(byt('x'), None)
assert 'y' == PlusZ2(None, byt('y'))
assert None == PlusZ2(None, None)

print 'z5 OKAY.'
