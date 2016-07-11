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

must except Foo(100, 3.14, byt('foo'), BBB(), [5])
must Foo(23, 3.14, "foo", AAA(), [5])
must except Foo(23, 3.14, None, AAA(), [5])

must except Foo(3.14, 3.14, "foo", BBB(), [5])
must except Foo(23, 3.14, "foo", ZZZ(), [5])
must except Foo(23, 3, "foo", AAA(), [5])
must except Foo(23, 3.14, (), AAA(), [5])
must except Foo(23, 3.14, [], AAA(), [5])
must except Foo(23, 3.14, {}, AAA(), [5])

must Foo(23, 3.14, "foo", AAA(), []) == []

must except Foo(23, 3.14, "foo", AAA(), (5, 5))
must except Foo(23, 3.14, "foo", AAA(), {5: 5})
must except BBB().Bar(23, 3.14, "foo", AAA(), {5: 5})

print "OKAY test_gradtype.py"
