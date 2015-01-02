from . import data

assert data.Eval('True') is True
assert data.Eval('False') is False
assert data.Eval('None') is None

assert data.Eval('12345') == 12345
assert data.Eval('-12345') == -12345
assert data.Eval('-1234.5') == -1234.5

assert data.Eval('[ 123, 4.5, False] ') == [123, 4.5, False]
# TODO: DiCT::EQ  # assert data.Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }
assert data.Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }

d = data.Eval('{ "color": "red", "area": 51 } ')
e = { "color": "red", "area": 51 }

assert sorted([(k, d[k]) for k in d]) == sorted([(k, e[k]) for k in e])

for foo in ['foo', dict(abc=range(5), cdef=None, ghi=3.25, xyz="하나 둘 셋")]:
  must foo == data.Eval(repr(foo))
  must foo == data.Eval(data.Eval(repr(repr(foo))))
  must foo == data.Eval(data.Eval(data.Eval(repr(repr(repr(foo))))))
