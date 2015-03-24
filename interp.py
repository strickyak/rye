import time
from go import bufio, fmt, os
from go import github.com/strickyak/rye/GPL
from . import tr
from . import Eval as EvalLiteral

class Scopes:
  def __init__():
    .g = {}  # globals
    .l = []  # vector of locals
    .b = {}
    native:
      `self.M_b= MkDict(BuiltinObj.Dict())`
  def Get(var):
    nom = var.name
    say var, nom
    if .l and nom in .l[0]:
      # Use global if it exists.
      z = .l[0][nom]
    elif nom in .g:
      # Use global if it exists.
      z = .g[nom]
    elif nom in .b:
      # Fall back to builtins.
      z = .b[nom]
    else:
      raise 'Variable %q not found; locals=%v ; globals=%v ; builtins=%v' % (nom, .l, .g, .b)
    return z

  def Put(var, x):
    say var.name, var, x
    nom = var.name
    if .l and nom in .l[0]:
      .l[0][nom] = x
    else:
      .g[var.name] = x

def main(args):
  sco = Scopes()
  if args:
    for a in args:
      # If it looks like a filename, open the file.
      if a.startswith('.') or a.startswith('/'):
        fd = open(a)
        with defer fd.close():
          code = fd.read()
      else:
        code = a
      say '<<<', code
      z = Interpret(code + '\n', sco)
  else:
    Repl(sco)
  say 'OKAY'

def Repl(sco):
  GPL.ReadHistoryFile('.rye.interp.history')
  serial = 0
  while True:
    try:
      line = GPL.ReadLine("rye> ")
    except as ex:
      print >>os.Stderr, "*** ", ex
      GPL.WriteHistoryFile('.rye.interp.history')
      return
    line = line.strip(' \t\n\r')
    if not line:
      continue

    if line.startswith('/h') or line.startswith('/H') or line.startswith('?'):
      print >>os.Stderr, '''
/b -- print builtins
/g -- print globals
/l -- print locals
/q -- quit
'''
      continue
    if line == '/q':
      return
    if line == '/b':
      print >>os.Stderr, ' '.join(sorted([k for k in sco.b]))
      continue
    if line == '/g':
      print >>os.Stderr, ' '.join(sorted([k for k in sco.g]))
      continue
    if line == '/l':
      print >>os.Stderr, ' '.join(sorted([k for k in sco.l]))
      continue

    try:
      z = Interpret(line + '\n', sco)
      if z is not None:
        serial += 1
        tmp = '_%d' % serial
        sco.g[tmp] = z
        print >>os.Stderr, "    %s = %s" % (tmp, repr(z))
    except as ex:
      print >>os.Stderr, "*** ", ex

def Interpret(program, sco):
  words = tr.Lex(program).tokens
  words = list(tr.SimplifyContinuedLines(words))
  parser = tr.Parser(program, words, -1, '<EVAL>')
  suite = parser.Csuite()
  walker2 = Interpreter(sco)
  say suite
  print >>os.Stderr, "------------------"
  start = time.time()
  z = suite.visit(walker2)
  finish = time.time()
  print >>os.Stderr, fmt.Sprintf("------------------ %.6f sec", finish - start)
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
    Is=(lambda a, b: id(a) == id(b)),
    IsNot=(lambda a, b: id(a) != id(b)),
    Contains=(lambda a, b: b in a),
    NotContains=(lambda a, b: b not in a),
    )

RAWS = {
    'True': True,
    'False': False,
    'None': None,
    }

class Interpreter:
  """Interpreter is a visitor to the Parse Tree that interprets."""
  def __init__(sco):
    .sco = sco

  def DestructuringAssign(target, e):
    tt = type(target)
    if tt is tr.Tvar:
      if target.name != '_':
        .sco.Put(target, e)
    elif tt is tr.Titems or tt is tr.Ttuple:
      try:
        ee = list(e)
        n = len(ee)
      except:
        raise 'Non-iterable value in Destructuring Assignment'

      if len(target.xx) != n:
        raise 'len(target) == %d should match len(value) == %d' % (len(target.xx), n)

      for vi, ei in zip(target.xx, ee):
        .DestructuringAssign(vi, ei)
    elif tt is tr.Traw and target.raw == '_':
      pass
    else:
      raise 'Weird target', target

  def Vop(self, p):  # a, op, b=None, returns_bool
    a = p.a.visit(self)
    op = p.op
    if p.b is None:
      fn = UNOPS.get(op)
      say a, op, fn
      if not fn:
        raise 'Unknown UNOP', op
      return fn(a)
    else:
      b = p.b.visit(self)
      fn = BINOPS.get(op)
      say a, op, b, fn
      if not fn:
        raise 'Unknown BINOP', op
      return fn(a, b)

  def Vboolop(self, p):  # a, op, b=None
    a = p.a.visit(self)
    op = p.op
    if op == '!':
      return False if a else True
    if op == '&&':
      if not a:
        return a
      return p.b.visit(self)
    if op == '||':
      if a:
        return a
      return p.b.visit(self)
    raise 'Unknown boolop: ', op

  def Vcondop(self, p):  # a, b, c # b if a else c
    a = p.a.visit(self)
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
      return EvalLiteral.Eval(p.v)
    else:
      raise 'Weird case', p.k

  def Vvar(self, p):  # name
    return .sco.Get(p)

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

  def Vforexpr(self, p):  # z(*body*), vv(*vars*), ll(*list*), cond, has_comma
    zz = []
    for e in p.ll.visit(self):
      .DestructuringAssign(p.vv, e)
      if p.cond:
        if not p.cond.visit(self):
          continue
      zz.append(p.z.visit(self))
    return zz

  def Vdict(self, p):  # xx
    #say p.xx
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
    z = None
    for x in p.things:
      z = x.visit(self)
    # Expressions eval as a Vsuite of Vassign, so we return the final value.
    return z

  def Vexpr(self, p):  # Statement.  a
    z = p.a.visit(self)
    raise "TODO: When is this used? (%s) -> (%s)" % (p, z)

  def Vassign(self, p):  # Statement.  a, b, pragma.
    z = p.b.visit(self)
    .DestructuringAssign(p.a, z)
    # Expressions eval as a Vassign, so we return the value.
    return z

  def Vprint(self, p):  # Statement.  w, xx, saying, code
    # TODO: w
    #say p.xx
    zz = p.xx.visit(self)
    #say zz
    if p.saying:
      print '#..# %s # %s' % (p.code, ' # '.join([repr(x) for x in zz]))
    else:
      if p.xx.trailing_comma:
        print ' '.join([str(x) for x in zz]),
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
    q = p.x  # The thing to assert.
    if type(q) == tr.Top:
      # Verbosely interpret the topmost op.
      a = q.a.visit(self)
      op = q.op
      if q.b is None:
        fn = UNOPS.get(op)
        say a, op, fn
        if not fn:
          raise 'Unknown UNOP', op
        z = fn(a)
      else:
        b = q.b.visit(self)
        fn = BINOPS.get(op)
        say a, op, b, fn
        if not fn:
          raise 'Unknown BINOP', op
        z = fn(a, b)
      extra = None
      if p.y:
        try:
          extra = p.y.visit(self)
        except:
          pass
      if not z:
        raise 'Assertion Failed:  left: (%v) ; op: %v ; right: (%v) ; extra: (%v) ; line: %v; code: %v' % (
          a, op, b, extra, p.line, p.code)
    else:
      z = q.visit(self)
      if not z:
        raise 'Assertion Failed: line: %v; code: %v' % (p.line, p.code)


  def Vtry(self, p):  # Statement. tr exvar ex
    try:
      p.tr.visit(self)
    except as ex:
      .DestructuringAssign(p.exvar, ex)
      p.ex.visit(self)

  def Vif(self, p):  # Statement.  t, yes, no.
    if p.t.visit(self):
      p.yes.visit(self)
    elif p.no:
      p.no.visit(self)

  def Vwhile(self, p):  # Statement. t, yes.
    done = False
    while not done and p.t.visit(self):
      try:
        p.yes.visit(self)
      except as ex:
        switch type(ex):
          case BreakEvent:
            done = True
          case ContinueEvent:
            pass
          default:
            raise ex

  def Vfor(self, p):  # Statement. var, t, b.
    done = False
    for e in p.t.visit(self):
      .DestructuringAssign(p.var, e)
      try:
        p.b.visit(self)
      except as ex:
        switch type(ex):
          case BreakEvent:
            done = True
          case ContinueEvent:
            pass
          default:
            raise ex
      if done:
        break


  def Vreturn(self, p):  # Statement.
    say p
    if p.aa is None:
      z = None
    else:
      vv = [a.visit(self) for a in p.aa]
      if len(vv) == 1:
        z = vv[0]
      else:
        z = vv
    say z
    raise ReturnEvent(z)

  def Vyield(self, p):  # Statement.
    raise 'Statement Not Implemented'

  def Vbreak(self, p):  # Statement.
    raise BreakEvent()

  def Vcontinue(self, p):  # Statement.
    raise ContinueEvent()

  def Vraise(self, p):  # Statement.
    vv = [a.visit(self) for a in p.aa]
    if len(vv) == 1:
      z = vv[0]
    else:
      z = vv
    raise z

  def Vdel(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vnative(self, p):  # Statement.
    raise 'Statement Not Implemented'

  def Vdef(self, p):  # Statement.  name, args, dflts, star, starstar, body.
    must not p.star
    must not p.starstar
    for x in p.dflts:
      must not x, x, p

    # LOOK AHEAD for "yield" and "global" statements.
    finder = tr.YieldGlobalAndLocalFinder()
    finder.Vsuite(p.body)
    yields = finder.yields
    force_globals = finder.force_globals
    assigned = finder.assigned
    lcl_vars = {}
    for x in assigned:
      if x not in force_globals:
        lcl_vars[x] = True

    def Bridge(*vec, **kw):
      vec = vec
      kw = kw

      sco = Scopes()
      saved_sco = .sco
      .sco = sco

      sco.b = .sco.b
      sco.g = .sco.g
      lcl = {}
      sco.l = [lcl]

      for x in lcl_vars:
        lcl[x] = None

      must not kw, kw
      if len(p.args) != len(vec):
        raise 'Interp func %q got %d args, expected %d' % (p.name, len(vec), len(p.args))

      for nom, val in zip(p.args, vec):
        lcl[nom] = val

      def restore_sco():
        .sco = saved_sco

      with defer restore_sco():
        z = None
        try:
          p.body.visit(self)
        except as ex:
          switch type(ex):
            case ReturnEvent:
              z = ex.x
            default:
              raise ex
        return z

    # TODO -- nested func names.
    .sco.g[p.name] = Bridge

  def Vclass(self, p):  # Statement.
    raise 'Statement Not Implemented'

class Event:
  pass

class ReturnEvent(Event):
  def __init__(x):
    .x = x

class BreakEvent(Event):
    pass

class ContinueEvent(Event):
    pass

must type(BreakEvent()) == BreakEvent

pass
