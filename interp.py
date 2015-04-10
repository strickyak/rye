import time
from go import bufio, fmt, os
from go import github.com/strickyak/rye/GPL
from . import tr
from . import Eval as EvalLiteral

class Scopes:
  def __init__():
    .y = None  # yielded
    .g = {}  # globals
    .l = []  # vector of locals
    .b = {}
    native:
      `self.M_b= MkDict(BuiltinObj.Dict())`

  def Get(name):
    for d in .l:
      if name in d:
        return d[name]

    if name in .g:
      return .g[name]

    if name in .b:
      return .b[name]

    raise 'Variable %q not found; locals=%v ; globals=%v ; builtins=%v' % (name, .l, .g, .b)

  def Put(name, x):
    # say name, x
    for d in .l:
      if name in d:
        d[name] = x
        # say name, 'PUT_IN', d, .l
        return

    .g[name] = x
    # say name, 'PUT_GLOBAL', .g

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
      say '>>>', code, '>>>', z
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
        .sco.Put(target.name, e)
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
    elif target is None:
      say 'Ignoring weird target None.'
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

  def Vraw(self, p):  # raw
    if p.raw in RAWS:
      return RAWS[p.raw]
    raise 'Raw value Not Implemented', p.raw

  def Vlit(self, p):  # k, v
    switch p.k:
      case 'N':
        return int(p.v)
      case 'F':
        return float(p.v)
      case 'S':
        return EvalLiteral.Eval(p.v)
      default:
        raise 'Weird case', p.k

  def Vvar(self, p):  # name
    return .sco.Get(p.name)

  def Vitems(self, p):  # xx, trailing_comma
    return [x.visit(self) for x in p.xx]

  def Vtuple(self, p):  # xx
    return tuple([x.visit(self) for x in p.xx])

  def Vlist(self, p):  # xx
    return [x.visit(self) for x in p.xx]

  def Vlambda(self, p):  # lvars, lexpr, where
    raise 'Expression Not Implemented'

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
    pairs = [(p.xx[i+i].visit(self), p.xx[i+i+1].visit(self)) for i in range(len(p.xx)/2)]
    return dict(pairs)

  def Vcall(self, p):  # fn, args, names, star, starstar
    fn = p.fn.visit(self)
    args = [a.visit(self) for a in p.args]

    vec = [args[i] for i in range(len(args)) if not p.names[i]]
    d = dict([(p.names[i], args[i]) for i in range(len(args)) if p.names[i]])

    if p.star:
      vec2 = p.star.visit(self)
      try:
        vec2 = list(vec2)
      except:
        raise 'Star argument must be a list'
      vec = vec + list(p.star.visit(self))

    if p.starstar:
      d2 = p.starstar.visit(self)
      try:
        d2 = dict(d2)
      except:
        raise 'StarStar argument must be a dict'

      for k, v in d2.items():
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
    zz = p.xx.visit(self)
    if p.saying:
      print '#..# %s # %s' % (p.code, ' # '.join([repr(x) for x in zz]))
    else:
      if p.xx.trailing_comma:
        print ' '.join([str(x) for x in zz]),
      else:
        print ' '.join([str(x) for x in zz])

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
    if .sco.y is not None:
      # Generated result.
      if p.aa is not None:
        raise 'Cannot return a value from a generator.'
      z = None
    elif p.aa is None:
      z = None
    else:
      vv = [a.visit(self) for a in p.aa]
      if len(vv) == 1:
        z = vv[0]
      else:
        z = vv
    raise ReturnEvent(z)

  def Vyield(self, p):  # Statement.
    if not p.aa:
      raise 'Yield statement is missing a value.'
    if .sco.y is None:
      raise 'Should not happen: Executed a yield when .y is None.'
    vv = [a.visit(self) for a in p.aa]
    if len(vv) == 1:
      z = vv[0]
    else:
      z = vv
    .sco.y.append(z)
    say vv, z, .sco.y

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
    switch type(p.listx):
      case tr.Titems:
        for e in p.listx.items.xx:
          self.Vdel(e)
      case tr.Tgetitem:
        del p.listx.a.visit(self)[p.listx.x.visit(self)]
      case tr.Tgetitemslice:
        del p.listx.a.visit(self)[p.listx.x.visit(self): p.listx.y.visit(self)]
      default:
        raise 'Cannot delete non-getitem non-slice'

  def Vnative(self, p):  # Statement.
    raise 'Statement Not Implemented'

  def Vdef(self, p):  # Statement.  name, args, dflts, star, starstar, body.
    defaults = {}
    for nom, d in zip(p.args, p.dflts):
      if d:
        defaults[nom] = d.visit(self)  # defaults evaluated now, at def time.
    argnames = dict([(arg, True) for arg in p.args])

    # LOOK AHEAD for "yield" and "global" statements.
    finder = tr.YieldGlobalAndLocalFinder()
    finder.Vsuite(p.body)
    yields = finder.yields
    force_globals = finder.force_globals
    assigned = finder.assigned
    say p.name, yields, force_globals, assigned
    lcl_vars = {}
    for x in assigned:
      if x not in force_globals:
        lcl_vars[x] = True

    def InterpFunc(*vec, **kw):
      # Save old Scopes, and start new one for this function.
      saved_sco = .sco
      .sco = Scopes()

      # Share builtins and globals; start new locals.
      .sco.b = saved_sco.b
      .sco.g = saved_sco.g
      lcl = {}
      .sco.l = [lcl]
      if yields:
        .sco.y = []

      # Create slots for all locals, initialized to None.
      for x in lcl_vars:
        lcl[x] = None

      if True or defaults or kw:
        newargs = defaults.copy()
        for nom, val in zip(p.args, vec):
          newargs[nom] = val

        newkw = {}
        for nom, x in kw.items():
          if nom in argnames:
            if nom in newargs:
              raise 'Keyword arg for %q conflicts with positional arg' % nom
            else:
              newargs[nom] = x
          else:
            newkw[nom] = x

        for nom in p.args:
          if nom not in newargs:
            raise 'No value provided for parameter %q' % nom

        lcl.update(newargs)

        if newkw:
          if p.starstar:
            lcl[p.starstar] = newkw
          else:
            raise 'Extra keyword args cannot be used (%d of them)' % len(newkw)

        excess = len(vec) - len(p.args)
        if excess > 0:
          # We have positional args that were not used.
          if p.star:
            lcl[p.star] = vec[excess:]
          else:
            raise 'Extra positional args cannot be used (%d of them)' % excess

      else:
        if p.star:
          lcl[p.star] = []
        if p.starstar:
          lcl[p.starstar] = {}
         
        # Only the simplest case is supported: exact unnamed args.
        must not kw, '**kw args not accepted by this function'
        if p.star:
          if len(p.args) > len(vec):
            raise 'Interp func %q got %d args, expected %d or more' % (p.name, len(vec), len(p.args))
        else:
          if len(p.args) != len(vec):
            raise 'Interp func %q got %d args, expected %d' % (p.name, len(vec), len(p.args))

        # Fill in arguments as locals.
        for nom, val in zip(p.args, vec):
          lcl[nom] = val
        if p.star:
          excess = len(vec) - len(p.args)
          lcl[p.star] = vec[excess:]
        if p.starstar:
          lcl[p.starstar] = {}

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
            case BreakEvent:
              raise 'break statement not inside loop'
            case ContinueEvent:
              raise 'continue statement not inside loop'
            default:
              raise ex
        if .sco.y is not None:
          say 'return', .sco.y
          return .sco.y
        else:
          say 'return', z
          return z

    .sco.Put(p.name, InterpFunc)

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
