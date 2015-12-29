import time
from go import bufio, fmt, os
from . import lex, parse
from lib import data, flags

SerialNum = 10
def Serial(s):  # Borrowed from codegen.
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (s, SerialNum)

class Scopes:
  def __init__():
    .y = None  # yielded
    .g = {}  # globals
    .l = []  # vector of locals
    .b = {}
    # Install builtins as builtins.
    native:
      `self.M_b= MkDict(BuiltinObj.Self.Dict())`

    def isInstanceFn(obj, c):
      say obj, str(c)
      t = obj.cls
      while t:
        say str(t), str(c)
        if t is c:
          return True
        t = t.superclass
      say str(t), False
      return False
    .b['isinstance'] = isInstanceFn

    def isSubclassFn(t, c):
      while t:
        say str(t), str(c)
        if t is c:
          return True
        t = t.superclass
      say str(t), False
      return False
    .b['issubclass'] = isSubclassFn

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

#MAIN = flags.Bool('main', False, 'Whether to execute main()')
FILEIN = flags.String('f', '', 'Source file to interpret')

def main(args):
  args = flags.Munch(args)
  sco = Scopes()

  if FILEIN.X:
    fd = open(FILEIN.X)
    with defer fd.close():
      code = fd.read()
    Interpret(code + '\n', sco)
  else:
    Repl(sco)
  print >>os.Stderr, '[interp.py] OKAY'

def Repl(sco):
  stdin = bufio.NewReader(os.Stdin)
  serial = 0
  while True:
    try:
      line = ''
      while True:
        os.Stderr.Write('rye> ')
        line2, hasMoreInLine = stdin.ReadLine()
        line += line2
        if not hasMoreInLine:
          break
    except as ex:
      print >>os.Stderr, "*** ", ex
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
      for l in sco.l:
        print >>os.Stderr, ' '.join(sorted([k for k in l]))
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
  words = lex.Lex(program).tokens
  words = list(lex.SimplifyContinuedLines(words, program=program))
  parser = parse.Parser(program, words, -1, '<EVAL>')
  suite = parser.Csuite()
  walker2 = Interpreter(sco)
  #say suite
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
    if tt is parse.Tvar:
      if target.name != '_' and target.name != 'rye_rye':
        .sco.Put(target.name, e)
    elif tt is parse.Titems or tt is parse.Ttuple:
      try:
        ee = list(e)
        n = len(ee)
      except:
        raise 'Non-iterable value in Destructuring Assignment'

      if len(target.xx) != n:
        raise 'len(target) == %d should match len(value) == %d' % (len(target.xx), n)

      for vi, ei in zip(target.xx, ee):
        .DestructuringAssign(vi, ei)
    elif tt is parse.Tgetitem:
      obj = target.a.visit(self)
      sub = target.x.visit(self)
      obj[sub] = e
    elif tt is parse.Tfield:
      obj = target.p.visit(self)
      obj.__setattr__(target.field, e)
    elif tt is parse.Traw and target.raw == '_':
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
        return data.Eval(p.v)
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

  def Vlambda(self, p):  # lvars, expr, where, line
    lamb = Serial('__lambda__')
    ret = parse.Treturn([p.expr])
    ret.where, ret.line, ret.gloss = p.where, p.line, 'lambda'
    suite = parse.Tsuite([ret])
    suite.where, suite.line, suite.gloss = p.where, p.line, 'lambda'

    if type(p.lvars) == parse.Titems:
      t = parse.Tdef(lamb, [x.name for x in p.lvars.xx], None, None, [None for x in p.lvars.xx], '', '', suite)
    elif type(p.lvars) == parse.Tvar:
      t = parse.Tdef(lamb, [p.lvars.name], None, None, [None], '', '', suite)
    else:
      raise Exception("Bad p.lvars type: %s" % type(p.lvars))

    t.where, t.line, t.gloss = p.where, p.line, 'lambda'
    t.visit(self)
    return parse.Tvar(lamb).visit(self)

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
    say 'Bird: VCall PREPARE ...', str(p.fn), p.line, (p.args), (p.names), (p.star), (p.starstar)
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
        d[k] = v

    say 'Bird: VCall CALL', str(fn), p.line, str(vec), str(d)
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

  # STATEMENTS

  def Vsuite(self, p):  # Statement.  things
    z = None
    for x in p.things:
      say x.where, x.line, type(x)
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
    if type(q) == parse.Top:
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
      case parse.Titems:
        for e in p.listx.items.xx:
          self.Vdel(e)
      case parse.Tgetitem:
        del p.listx.a.visit(self)[p.listx.x.visit(self)]
      case parse.Tgetitemslice:
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
    finder = parse.YieldGlobalAndLocalFinder()
    finder.Vsuite(p.body)
    lcl_vars = {}
    for x in finder.assigned:
      if x not in finder.force_globals:
        lcl_vars[x] = True
    outer_sco = .sco

    def InterpFunc(*vec, **kw):
      # Save old Scopes, and start new one for this function.
      saved_sco = .sco
      .sco = Scopes()

      # Share builtins and globals; start new locals.
      .sco.b = outer_sco.b
      .sco.g = outer_sco.g
      lcl = {}
      .sco.l = [lcl] + outer_sco.l
      if finder.yields:
        .sco.y = []

      # Create slots for all locals, initialized to None.
      for x in lcl_vars:
        lcl[x] = None

      if True or defaults or kw:
        say p.name, p.line, p.where, p.gloss, vec, kw
        newargs = {}
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

        say sorted(newkw.keys())
        say sorted(newargs.keys())
        say sorted(defaults.keys())

        for nom in p.args:
          if nom not in newargs:
            if nom in defaults:
              newargs[nom] = defaults[nom]
            else:
              raise 'No value provided for parameter %q' % nom

        lcl.update(newargs)

        if newkw:
          if p.starstar:
            lcl[p.starstar] = newkw
          else:
            raise 'Extra keyword args cannot be used (%d of them)' % len(newkw)
        else:
          if p.starstar:
            lcl[p.starstar] = newkw

        excess = len(vec) - len(p.args)
        if excess > 0:
          # We have positional args that were not used.
          if p.star:
            say vec[excess:]
            lcl[p.star] = tuple(vec[len(p.args):])
            must len(lcl[p.star]) == excess
          else:
            raise 'Extra positional args cannot be used (%d of them)' % excess
        else:
          if p.star:
            lcl[p.star] = ()

      else:
        if p.star:
          lcl[p.star] = []
        if p.starstar:
          lcl[p.starstar] = {}

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
          #say 'Vdef: Trying', p.body
          p.body.visit(self)
          #say 'Vdef: Success!'
        except as ex:
          #say 'Vdef: Caught', type(ex), ex
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
          #say 'Vdef: yielded return', .sco.y
          return .sco.y
        else:
          #say 'Vdef: return', z
          return z

    say .sco, .sco.g
    say p.name, InterpFunc
    .sco.Put(p.name, InterpFunc)
    say .sco, .sco.g

  def Vclass(self, p):  # Tclass: name, sup, things.
    cls = Class(p, self)

    def ctor(*args, **kw):
      must not kw  # not yet
      obj = Instance(cls)
      fn = cls.meths.get('__init__')
      if fn:
        fn(obj, *args, **kw)
      return obj

    .sco.Put(p.name, ctor)
    # raise 'Statement Not Implemented'

Classes = dict(object=None)

class Class:
  def __init__(tclass, visitor):
    .tclass = tclass
    .name = tclass.name
    .sup = tclass.sup
    .meths = {}
    say .name, .sup, sorted(Classes.keys())
    .superclass = Classes[tclass.sup.name]
    say tclass, .name, .sup, .sup.visit(visitor), .meths

    saved_sco = visitor.sco
    visitor.sco = Scopes()
    visitor.sco.g = .meths  # Methods will get added to .g when visited.

    # Copy meths from our superclass.
    if .superclass:
      for k, meth in .superclass.meths.items():
        .meths[k] = meth

    def restore_sco():
      visitor.sco = saved_sco

    with defer restore_sco():
      for t in tclass.things:
        switch type(t):
          default:
            raise 'Weird type in Tclass.things: %q' % str(type(t))
          case parse.Tdef: # name, args, dflts, star, starstar, body.
            if t.args:
              if t.args[0] != 'self':
                t.args = ['self'] + t.args
                t.dflts = [None] + t.dflts
            else:
              t.args = ['self']
              t.dflts = [None]

            .meths[t.name] = False  # Reserve a slot, for visitor.sco.Put() at end of Vdef().
            t.visit(visitor)
            say t.name, .meths[t.name]

    Classes[.name] = self


class Instance:
  def __init__(cls):
    .cls = cls
    .d = {}
    for methname, fn in .cls.meths.items():
      def wrapper(*args, **kw):
        say "Bird: Calling method (%v@%v).%v ( %v ; %v )" % (str(cls), hash(self), methname, args, kw)
        return fn(self, *args, **kw)
      .d[methname] = wrapper
  def __str__():
    return 'Instance(%s@%d)' % (.cls.name, (id(self) % 997) + 2)
  def __repr__():
    return self.__str__()
    #return 'Instance(%s){%s}' % (.cls.name, ','.join(['%s:%s' % (k, v) for k, v in sorted(.d.items())]))
  def __getattr__(name):
    if name in .d:
      return .d[name]
    raise 'No such attr %q on %q' % (name, .cls.name)
  def __setattr__(name, value):
    .d[name] = value

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
