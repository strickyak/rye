from go import bufio, os
from . import tr
from . import Eval

def main(args):
  scope = dict()
  for a in args:
    say '<<<', a
    z = Interpret(a + '\n', scope)
  say 'OKAY'

def Repl(builtins, glbls):
  scope = {}
  for k, v in builtins.items():
    scope[k] = v
  for k, v in glbls.items():
    scope[k] = v

  input = bufio.NewReader(os.Stdin)
  serial = 0
  while True:
    print >>os.Stderr, "> ",
    try:
      line = input.ReadString(ord('\n'))
    except as ex:
      print >>os.Stderr, "*** ", ex
      return
    line = line.strip(' \t\n\r')
    if not line:
      continue
    try:
      z = Interpret(line + '\n', scope)
      if z is not None:
        serial += 1
        tmp = '_%d' % serial
        scope[tmp] = z
        print >>os.Stderr, "OKAY: %s = %s" % (tmp, repr(z))

    except as ex:
      print >>os.Stderr, "*** ", ex

def Interpret(program, scope):
  words = tr.Lex(program).tokens
  words = list(tr.SimplifyContinuedLines(words))
  parser = tr.Parser(program, words, -1, '<EVAL>')
  #tree = parser.Csuite()
  tree = parser.Command()
  say tree

  walker = ShowExprWalker()
  say tree.visit(walker)

  walker2 = EvalWalker([scope])
  print >>os.Stderr, "------------------"
  z = tree.visit(walker2)
  print >>os.Stderr, "------------------"
  return z

UNOPS = dict(
  UnaryPlus=(lambda a: +a),
  UnaryMinus=(lambda a: -a),
  UnaryInvert=(lambda a: ~a),
  )

BINOPS = dict(
    Add=(lambda a, b: a + b),
    Sub=(lambda a, b: a - b),
    Mul=(lambda a, b: a * b),
    Div=(lambda a, b: a / b),
    IDiv=(lambda a, b: a // b),
    Mod=(lambda a, b: a % b),
    EQ=(lambda a, b: a == b),
    NE=(lambda a, b: a != b),
    LT=(lambda a, b: a < b),
    LE=(lambda a, b: a <= b),
    GT=(lambda a, b: a > b),
    GE=(lambda a, b: a >= b),
    ShiftLeft=(lambda a, b: a << b),
    ShiftRight=(lambda a, b: a >> b),
    UnsignedShiftRight=(lambda a, b: a >>> b),
    BitAnd=(lambda a, b: a & b),
    BitOr=(lambda a, b: a | b),
    BitXor=(lambda a, b: a ^ b),
    )

RAWS = {
    'True': True,
    'False': False,
    'None': None,
    }

def DestructuringAssign(scope, target, e):
  tt = type(target)
  if tt is tr.Tvar:
    if target.name != '_':
      scope[target.name] = e
  elif tt is tr.Titems or tt is tr.Ttuple:
    try:
      ee = list(e)
      n = len(ee)
    except:
      raise 'Non-iterable value in Destructuring Assignment'

    if len(target.xx) != n:
      raise 'len(target) == %d should match len(value) == %d' % (len(target.xx), n)

    for vi, ei in zip(target.xx, ee):
      DestructuringAssign(scope, vi, ei)
  elif tt is tr.Traw and target.raw == '_':
    pass
  else:
    raise 'Weird target', target

class EvalWalker:
  def __init__(scopes):
    .scopes = scopes

  def Vop(self, p):  # a, op, b=None, returns_bool
    a = p.a.visit(self)
    op = p.op
    b = p.b.visit(self) if p.b else None
    if b:
      fn = BINOPS.get(op)
      if not fn:
        raise 'Unknown BINOP', op
      return fn(a, b)
    else:
      fn = UNOPS.get(op)
      if not fn:
        raise 'Unknown UNOP', op
      return fn(a)

  def Vboolop(self, p):  # a, op, b=None
    a = p.a.visit(self)
    op = p.op
    b = p.b.visit(self) if p.b else None
    if op == '!':
      return False if a else True
    if op == '&&':
      if not a:
        return a
      return b
    if op == '||':
      if a:
        return a
      return b
    raise 'Unknown boolop: ', op

  def Vcondop(self, p):  # a, b, c # b if a else c
    say p.a, p.b, p.c
    a = p.a.visit(self)
    say a, not a
    if a:
      zt = p.b.visit(self)
      return zt
    else:
      zf = p.c.visit(self)
      return zf

  def Vgo(self, p):  # fcall
    raise 'Expression Not Implemented'
    say p
    z = '( GO-CALL %v )' % (p.fcall.visit(self), )
    say z
    return z

  def Vraw(self, p):  # raw
    if p.raw in RAWS:
      return RAWS[p.raw]
    raise 'Raw value Not Implemented', p.raw

  def Vlit(self, p):  # k, v
    if p.k == 'N':
      return int(p.v)
    elif p.k == 'F':
      return float(p.v)
    elif p.k == 'S':
      return Eval.Eval(p.v)
    else:
      raise 'Weird case', p.k

  def Vvar(self, p):  # name
    for scope in .scopes:
      if p.name in scope:
        return scope[p.name]
    raise 'Unknown variable: %q' % p.name

  def Vitems(self, p):  # xx, trailing_comma
    return [x.visit(self) for x in p.xx]

  def Vtuple(self, p):  # xx
    return tuple([x.visit(self) for x in p.xx])

  def Vlist(self, p):  # xx
    return [x.visit(self) for x in p.xx]

  def Vlambda(self, p):  # lvars, lexpr, where
    raise 'Expression Not Implemented'
    say p
    z = '( LAMBDA ... )'
    say z
    return z

  def Vforexpr(self, p):  # z, vv, ll, cond, has_comma
    zz = []
    for e in p.ll.visit(self):
      DestructuringAssign(.scopes[0], p.vv, e)
      zz.append(p.z.visit(self))
    return zz

  def Vdict(self, p):  # xx
    say p.xx
    pairs = [(p.xx[i+i].visit(self), p.xx[i+i+1].visit(self)) for i in range(len(p.xx)/2)]
    return dict(pairs)

  def Vcall(self, p):  # fn, args, names, star, starstar
    fn = p.fn.visit(self)
    args = [a.visit(self) for a in p.args]

    vec = [args[i] for i in range(len(args)) if not p.names[i]]
    d = dict([(p.names[i], args[i]) for i in range(len(args)) if p.names[i]])

    if p.star:
      vec = vec + [x.visit(self) for x in p.star]
    if p.starstar:
      for k, v in p.starstar.items():
        d[k] = v.visit(self)

    return fn(*vec, **d)

  def Vfield(self, p):  # p, field
    return getattr(p.p.visit(self), p.field)

  def Vgetitem(self, p):  # a, x
    a = p.a.visit(self)
    x = p.x.visit(self)
    return a[x]

  def Vgetitemslice(self, p):  # a, x, y, (z)
    a = p.a.visit(self)
    x = p.x.visit(self)
    y = p.y.visit(self)
    return a[x:y]

  def Vcurlysetter(self, p):  # # obj, vec of (var, expr)
    raise 'CurlySetter Not Implemented'

  # STATEMENTS

  def Vsuite(self, p):  # Statement.  things
    z = '( SUITE[%d] %v )' % (len(p.things), (', '.join([x.visit(self) for x in p.things])))
    return z

  def Vexpr(self, p):  # Statement.  a
    z = '( EXPR %v )' % (p.a.visit(self), )
    return z

  def Vassign(self, p):  # Statement.  a, b, pragma.
    z = p.b.visit(self)
    DestructuringAssign(.scopes[0], p.a, z)
    return z

  def Vprint(self, p):  # Statement.  w, xx, saying, code
    # TODO: p.xx.trailing_comma
    # TODO: w
    say p.xx
    zz = p.xx.visit(self)
    say zz
    if p.saying:
      print '#..# %s # %s' % (p.code, ' # '.join([repr(x) for x in zz]))
    else:
      print ' '.join([str(x) for x in zz])

  def Vdefer(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vwithdefer(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vglobal(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vimport(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vassert(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vtry(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vif(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vwhile(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vfor(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vreturn(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vyield(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vbreak(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vcontinue(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vraise(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vdel(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vnative(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vdef(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vclass(self, p):  # Statement.
    raise 'Statement Not Implemented'

class ShowExprWalker:
  def __init__():
    pass

  def Vop(self, p):  # a, op, b=None, returns_bool
    say p
    z = '( OP%v %v %v %v )' % ('-bool' if p.returns_bool else '', p.a.visit(self), p.op, p.b.visit(self) if p.b else None)
    say z
    return z

  def Vboolop(self, p):  # a, op, b=None
    say p
    z = '( BOOLOP %v %v %v )' % (p.a.visit(self), p.op, p.b.visit(self) if p.b else None)
    say z
    return z

  def Vcondop(self, p):  # a, b, c
    say p
    z = '( CONDOP %v IF %v ELSE %v )' % (p.a.visit(self), p.b.visit(self), p.c.visit(self))
    say z
    return z

  def Vgo(self, p):  # fcall
    say p
    z = '( GO-CALL %v )' % (p.fcall.visit(self), )
    say z
    return z

  def Vraw(self, p):  # raw
    say p
    z = '( RAW {{{%v}}} )' % (p.raw, )
    say z
    return z

  def Vlit(self, p):  # k, v
    say p
    z = '( LIT %v {{{%v}}} )' % (p.k, p.v, )
    say z
    return z

  def Vvar(self, p):  # name
    say p
    z = '( VAR %v )' % (p.name, )
    say z
    return z

  def Vitems(self, p):  # xx, trailing_comma
    say p
    z = '( ITEMS[%d] %v )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vtuple(self, p):  # xx
    say p
    z = '( TUPLE[%d] %v )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vlist(self, p):  # xx
    say p
    z = '( LIST[%d] %v )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vlambda(self, p):  # lvars, lexpr, where
    say p
    z = '( LAMBDA ... )'
    say z
    return z

  def Vforexpr(self, p):  # z, vv, ll, cond, has_comma
    say p
    z = '( FOREXPR ... )'
    say z
    return z

  def Vdict(self, p):  # xx
    say p
    z = '( DICT[%d] %v )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vcall(self, p):  # for defer or go.
    return "( CALL... )"
  def Vfield(self, p):  # 
    return "( FIELD... )"
  def Vgetitem(self, p):  # 
    return "( GETITEM... )"
  def Vgetitemslice(self, p):  # 
    return "( GETITEMSLICE... )"
  def Vcurlysetter(self, p):  # 
    return "( VCURLYSETTER... )"

  # STATEMENTS

  def Vsuite(self, p):  # Statement.  things
    z = '( SUITE[%d] %v )' % (len(p.things), (', '.join([x.visit(self) for x in p.things])))
    return z

  def Vexpr(self, p):  # Statement.  a
    z = '( EXPR %v )' % (p.a.visit(self), )
    return z

  def Vassign(self, p):  # Statement.  a, b, pragma.
    z = '( ASSIGN %v %v )' % (p.a.visit(self), p.b.visit(self), )
    return z


  def Vprint(self, p):  # Statement.  w, xx, saying, code
    z = '( PRINT %v )' % (p.xx.visit(self), )
    return z

  def Vdefer(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vwithdefer(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vglobal(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vimport(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vassert(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vtry(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vif(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vwhile(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vfor(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vreturn(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vyield(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vbreak(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vcontinue(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vraise(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vdel(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vnative(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vdef(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vclass(self, p):  # Statement.
    raise 'Statement Not Implemented'

# EXPRESSIONS
#1568 class Top(Tnode): a, op, b=None, returns_bool
#1577 class Tboolop(Tnode): a, op, b=None
#1585 class Tcondop(Tnode): a, b,c
#1593 class Tgo(Tnode): fcall
#1599 class Traw(Tnode): raw
#1605 class Tlit(Tnode): k, v
#1612 class Tvar(Tnode): name
#1618 class Titems(Tnode): xx, trailing_comma
#1625 class Ttuple(Tnode): xx
#1631 class Tlist(Tnode): xx
#1637 class Tlambda(Tnode): lvars, lexpr, where
#1645 class Tforexpr(Tnode): z, vv, ll, cond, has_comma
#1655 class Tdict(Tnode): xx
#1829 class Tcall(Tnode):
#1839 class Tfield(Tnode):
#1846 class Tgetitem(Tnode):
#1853 class Tgetitemslice(Tnode):
#1862 class Tcurlysetter(Tnode):

# STATEMENTS
#1661 class Tsuite(Tnode):
#1667 class Texpr(Tnode):
#1673 class Tassign(Tnode):
#1681 class Tprint(Tnode):
#1690 class Tdefer(Tnode):
#1696 class Twithdefer(Tnode):
#1703 class Tglobal(Tnode):
#1709 class Timport(Tnode):
#1717 class Tassert(Tnode):
#1727 class Ttry(Tnode):
#1735 class Tif(Tnode):
#1743 class Twhile(Tnode):
#1750 class Tfor(Tnode):
#1758 class Treturn(Tnode):
#1764 class Tyield(Tnode):
#1770 class Tbreak(Tnode):
#1776 class Tcontinue(Tnode):
#1782 class Traise(Tnode):
#1788 class Tdel(Tnode):
#1794 class Tnative(Tnode):
#1800 class Tdef(Tnode):
#1821 class Tclass(Tnode):
