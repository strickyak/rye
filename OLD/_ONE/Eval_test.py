from . import Eval as E

assert E.Eval('True') is True
assert E.Eval('False') is False
assert E.Eval('None') is None

assert E.Eval('12345') == 12345
assert E.Eval('-12345') == -12345
assert E.Eval('-1234.5') == -1234.5

assert E.Eval('[ 123, 4.5, False] ') == [123, 4.5, False]
# TODO: DiCT::EQ  # assert E.Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }
assert E.Eval('{ "color": "red", "area": 51 } ') == { "color": "red", "area": 51 }

d = E.Eval('{ "color": "red", "area": 51 } ')
e = { "color": "red", "area": 51 }

assert sorted([(k, d[k]) for k in d]) == sorted([(k, e[k]) for k in e])
