from go import github.com/strickyak/rye/test_govar2 as T2

def ChangeXyzzy(a):
  T2.Xyzzy = a

assert int(T2.Xyzzy) == 42, T2.Xyzzy
ChangeXyzzy(23)
assert int(T2.Xyzzy) == 23, T2.Xyzzy
T2.Xyzzy = 888
assert int(T2.Xyzzy) == 888, T2.Xyzzy
