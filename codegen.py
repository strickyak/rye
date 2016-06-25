import md5
import os
import re
import sys

rye_rye = False
if rye_rye:
  from rye_lib import data
  from go import strconv
  from . import parse
  from . import samples
  from . import gen_internals
  from . import goapi
else:
  import parse
  import samples
  import gen_internals
  import goapi

OPTIONAL_MODULE_OBJS = True  # Required for interp.

INTLIKE_GO_TYPES = {'int', 'int8', 'int16', 'int32', 'int64', 'uint', 'uint8', 'uint16', 'uint32', 'uint64', 'uintptr'}
FLOATLIKE_GO_TYPES = {'float32', 'float64'}

RYE_SPECIALS = {
    'go_cast', 'go_type', 'go_indirect', 'go_addr', 'go_new', 'go_make', 'go_append',
    'len', 'str', 'repr', 'int', 'float',
    }

NoTyps = None
NoTyp = None

SMALLER = not os.getenv('RYE_BLOAT')

NONALFA = re.compile('[^A-Za-z0-9]')
TROUBLE_CHAR = re.compile('[^]-~ !#-Z[]')
def GoStringLiteral(s):
  if rye_rye:
    return strconv.QuoteToASCII(s)
  else:
    return '"' + TROUBLE_CHAR.sub((lambda m: '\\x%02x' % ord(m.group(0))), s) + '"'

def CleanIdentWithSkids(s):
  if len(s) < 50:
    # Work around lack of callable() for .sub in RYE.
    return md5.new(s).hexdigest()
    # TODO = NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
    # return NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
  else:
    return md5.new(s).hexdigest()

############################################################

# p might be abosolute; but more are always relative.
# returns a list of part names.
def CleanPath(cwd, p, *more):
  if p.startswith('/'):
    # Absolute.
    q = p.split('/')
  else:
    # Relative
    q = cwd.split('/') + p.split('/')
  for m in more:
    q += m.split('/')

  x = []
  for w in q:
    if w == '' or w == '.':
      continue
    elif w == '..':
      if x:
        x = x[:-1]
      else:
        raise Exception('Bad path (too many ".."): "%s" in "%s"' % (p, cwd))
    elif w.startswith('.'):
      raise Exception('Bad word in path (starts with "."): "%s" in "%s"' % (p, cwd))
    else:
      x.append(w)
  return x

class CodeGen(object):
  def __init__(self):
    self.glbls = {}         # name -> (type, initialValue)
    self.imports = {}       # name -> Vimport
    self.defs = {}          # name -> ArgsDesc()
    self.lits = {}          # key -> name
    self.invokes = {}       # key -> (n, fieldname)
    self.scope = None       # None means we are at module level.
    self.tail = []
    self.cls = None
    self.getNeeded = {}      # keys are getter names.
    self.setNeeded = {}      # keys are setter names.
    self.maxNumCallArgs = -1
    self.SerialNum = 100


  def Serial(self, s):
    self.SerialNum += 1
    return '%s_%d' % (str(s), self.SerialNum)

  def InjectForInternal(self, stuff):
    self.invokes, self.defs, self.getNeeded, self.setNeeded = stuff

  def ExtractForInternal(self):
    stuff = self.invokes, self.defs, self.getNeeded, self.setNeeded
    return stuff

  def GenModule(self, modname, path, tree, cwp=None, internal=""):
    self.cwp = cwp
    self.path = path
    self.modname = modname
    if internal:
      self.internal = open('gen_internals.py', 'w')
    else:
      self.internal = None
      self.glbls['__name__'] = ('M', 'MkStr("%s")' % modname)

    self.func_level = 0
    self.func = None
    self.yields = None
    self.force_globals = {}

    if internal:
      print ' package rye'
    else:
      print ' package %s' % modname.split('/')[-1]
      print ' import . "github.com/strickyak/rye"'

    print ' import "fmt"'
    print ' import "io"'
    print ' import "os"'
    print ' import "reflect"'
    print ' import "runtime"'
    print ' import "time"'  # For FAV packages.
    print ' import "bytes"'  # For FAV packages.

    # LOOK FOR MAIN
    main_def = None
    for th in tree.things:
      if type(th) == parse.Tdef:
        if th.name == 'main':
          main_def = th
    # ADD A MAIN, if there isn't one.
    if not main_def and not internal:
      main_def = parse.Tdef('main', ['argv'], NoTyps, NoTyp, [None], None, None, parse.Tsuite([]))
      main_def.where, main_def.line, main_def.gloss = 0, 0, 'synthetic-def-main'
      tree.things.append(main_def)

    # IMPORTS.
    to_be_sampled = {}
    for th in tree.things:
      imps = []
      if type(th) == parse.Timport: # Simple import
        imps = [th]
      if type(th) == parse.Tif:
        if type(th.t) is parse.Tvar and th.t.name == 'rye_rye':  # Under "if rye_rye:" ...
          for th2 in th.yes.things:
            if type(th2) is parse.Timport:  # ... we find an import.
              imps.append(th2)

      for imp in imps:
        vec = imp.imported
        if vec[0] == 'go':
          vec = vec[1:]  # Trim leading "go" mark, for go paths.
        else:
          vec = vec[:-1] + ['rye__'] + vec[-1:]  # Insert "rye__" as penultimate part.
        pkg = '/'.join(vec)
        if imp.alias == '_':
          print ' import _ "%s"' % pkg
        else:
          alias = 'i_%s' % imp.alias
          print ' import %s "%s"' % (alias, pkg)

          if samples.SAMPLES.get(pkg):
            to_be_sampled[alias] = pkg

          if not self.internal:
            self.glbls[imp.alias] = ('*PModule', 'None')

    for alias, pkg in sorted(to_be_sampled.items()):
      print 'var _ = %s.%s // %s' % (alias, samples.SAMPLES[pkg], pkg)

    # MAKE CONSTRUCTORS FOR CLASSES
    for th in tree.things:
      if type(th) == parse.Tclass and th.sup != "native":
        # name, sup, things
        # Create constructor functions with Tdef().

        # newctor:
        c_init = None
        for thth in list(th.things):
          if type(thth) is parse.Tdef and thth.name == '__init__':
            c_init = thth

        c_name = th.name
        #print ''
        #print '// NEWCTOR:', c_name
        if c_init:
          # When there is a __init__, we just use * and ** to call it.
          c_args =  c_init.args
          c_dflts =  c_init.dflts
          c_star =  c_init.star
          c_starstar =  c_init.starstar
        else:
          # Since there was no __init__, we just use * and ** to call it.
          c_args =  []
          c_dflts =  []
          c_star =  'rye_ctor_vec__'
          c_starstar =  'rye_ctor_kw__'

        if c_args and c_args[0] == 'self':
          c_args = c_args[1:]
          c_dflts = c_dflts[1:]

        x_star = parse.Tvar(c_star) if c_star else None
        x_starstar = parse.Tvar(c_starstar) if c_starstar else None

        natives = [
              '   z := new(C_%s)' % th.name,
              '   z.Self = z',
              '   z.Rye_ClearFields__()',
              ]
        c1 = parse.Tnative(natives)

        c2 = parse.Tassign(parse.Tvar('rye_result__'), parse.Traw('MkX(&z.PBase)'))

        # Tcall: fn, args, names, star, starstar
        call = parse.Tcall(parse.Tfield(parse.Tvar('rye_result__'), '__init__'), [parse.Tvar(a) for a in c_args], c_args, x_star, x_starstar)
        c3 = parse.Tassign(parse.Traw('_'), call)

        c4 = parse.Treturn([parse.Tvar('rye_result__')])

        suite = parse.Tsuite([c1, c2, c3, c4])
        ctor = parse.Tdef(c_name, c_args, NoTyps, NoTyp, c_dflts, c_star, c_starstar, suite, isCtor=True)
        for t in [c1, c2, c3, c4, suite, ctor]:
          t.line = th.line
          t.where = th.where
          t.gloss = 'ctor'

        tree.things.append(ctor)

    # Avoid golang's "import not used" errors in corner cases.
    print ' var _ = fmt.Sprintf'
    print ' var _ = io.EOF'
    print ' var _ = os.Stderr'
    print ' var _ = reflect.ValueOf'
    print ' var _ = runtime.Stack'
    print ' var _ = time.Sleep'  # For FAV packages.
    print ' var _ = bytes.Split'  # For FAV packages.
    print ' var _ = MkInt'  # From rye runtime.
    print ''

    # BEGIN: Eval_Module, innter_eval_module
    if not self.internal:
      print ' var eval_module_once bool'
      print ' func Eval_Module () M {'
      print '   if eval_module_once == false {'
      print '     eval_module_once = true'
      print '     _ = inner_eval_module()'
      print '   }'
      if OPTIONAL_MODULE_OBJS:
        print '   return ModuleObj'
      else:
        print '   return None'
      print ' }'

    print ' func inner_eval_module () M {'

    # ALL THINGS IN MODULE.
    for th in tree.things:
      self.Gloss(th)
      th.visit(self)
      self.Ungloss(th)

    # END: Eval_Module, innter_eval_module
    print '   return None'
    print ' }'
    print ''
    print '//(begin tail)'
    print '\n//(tail)\n'.join(self.tail)
    print '//(end tail)'
    print ''

    # G_rye_rye is defined in runtime.go, so don't put it in other modules.
    if 'rye_rye' in self.glbls:
      del self.glbls['rye_rye']

    for g, (t, v) in sorted(self.glbls.items()):
      print 'var G_%s M // %s' % (g, t)
    print ''
    print ' func init /*New_Module*/ () {'
    for g, (t, v) in sorted(self.glbls.items()):
      print '   G_%s = %s' % (g, v)
    if internal:
      print '   inner_eval_module()'
    print ' }'
    print ''
    if OPTIONAL_MODULE_OBJS:
      print 'var %s = map[string]*M {' % ('BuiltinMap' if self.internal else 'ModuleMap')
      for g, (t, v) in sorted(self.glbls.items()):
        print '  "%s": &G_%s,' % (g, g)
      print '}'
      print ''
      print 'var %s = MakeModuleObject(%s, "%s/%s")' % (
          'BuiltinObj' if internal else 'ModuleObj',
          'BuiltinMap' if internal else 'ModuleMap',
          self.cwp, self.modname)
      print ''

    for key, code in sorted(self.lits.items()):
      print 'var %s = %s' % (key, code)
    print ''

    if self.internal:
        print >> self.internal, '# This file is generated when builtins.py is compiled.'
        print >> self.internal, 'InternalInvokers = ['
    for _, (n, fieldname) in sorted(self.invokes.items()):
      self.getNeeded[fieldname] = True
      formals = ', '.join(['a_%d M' % i for i in range(n)])
      args = ', '.join(['a_%d' % i for i in range(n)])

      if self.internal or (n, fieldname) not in gen_internals.InternalInvokers:
        letterF = 'F' if self.internal else 'f'
        letterI = 'I' if self.internal else 'i'
        letterGet = 'I' if self.internal or fieldname in gen_internals.InternalGetters else 'i'
        print 'func %s_INVOKE_%d_%s(fn M, %s) M {' % (letterF, n, fieldname, formals)
        print '  if fn.X == nil {'
        print '    if len(fn.S) == 0 { panic("cannot INVOKE on int") }'
        print '    fn = M{X: MkBStr(fn.S).Self}'  # TODO: stop using MkBStr.
        print '  }'
        print '  switch x := fn.X.(type) {   '
        print '  case %s_INVOKE_%d_%s:         ' % (letterI, n, fieldname)
        print '    return x.M_%d_%s(%s)         ' % (n, fieldname, args)
        print '  case %s_GET_%s:         ' % (letterGet, fieldname)
        print '    tmp := x.GET_%s()    ' % fieldname
        print '    return %s_%d(tmp, %s)' % (('CALL' if n<11 else 'call'), n, ', '.join(['a_%d' % j for j in range(n)]))
        print ''
        print '  case *PGo:                '
        print '    return x.Invoke("%s", %s) ' % (fieldname, args)
        print '  }'
        print '  panic(fmt.Sprintf("Cannot invoke \'%s\' with %d arguments on %%v", fn))' % (fieldname, n)
        print '}'
        print 'type %s_INVOKE_%d_%s interface { M_%d_%s(%s) M }' % (letterI, n, fieldname, n, fieldname, formals)
        if self.internal:
          print >> self.internal, '  (%d, "%s"),' % (n, fieldname)
    print ''

    if self.internal:
      print >> self.internal, '  ]'
      print >> self.internal, '# FooBarBaz'
      print >> self.internal, 'InternalGetters = ['

    for iv in sorted(self.getNeeded):
      if self.internal or iv not in gen_internals.InternalGetters:
        letterF = 'F' if self.internal else 'f'
        letterI = 'I' if self.internal else 'i'
        print 'type %s_GET_%s interface { GET_%s() M }' % (letterI, iv, iv)
        print 'func %s_GET_%s(h M) M {' % (letterF, iv)
        print '  if h.X == nil { panic("cannot GET Field on int or str") }'
        print '  switch x := h.X.(type) { '
        print '  case %s_GET_%s:         ' % (letterI, iv)
        print '    return x.GET_%s()    ' % iv
        print '  }'
        print '   return h.FetchField("%s") ' % iv
        print '}'
        print ''
        if self.internal:
          print >> self.internal, '  "%s",' % iv

    if self.internal:
      print >> self.internal, '  ]'
      print >> self.internal, 'InternalSetters = ['

    for iv in sorted(self.setNeeded):
      if self.internal or iv not in gen_internals.InternalSetters:
        letterF = 'F' if self.internal else 'f'
        letterI = 'I' if self.internal else 'i'
        print 'type %s_SET_%s interface { SET_%s(M) }' % (letterI, iv, iv)
        print 'func %s_SET_%s(h M, a M) {' % (letterF, iv)
        print '  switch x := h.X.(type) { '
        print '  case %s_SET_%s:         ' % (letterI, iv)
        print '    x.SET_%s(a)    ' % iv
        print '    return'
        print '  }'
        print '    h.StoreField("%s", a)' % iv
        print '}'
        print ''
        if self.internal:
          print >> self.internal, '  "%s",' % iv
    print ''

    if self.internal:
      print >> self.internal, '  ]'

    maxCall = 11 if self.internal else 1+self.maxNumCallArgs
    for i in range(maxCall):
      if (internal and i<11) or (not internal and i>=11):
        whichI = 'I' if i<11 else 'i'
        whichCall = 'CALL' if i<11 else 'call'
        print '  type %s_%d interface { Call%d(%s) M }' % (whichI, i, i, ", ".join(i * ['M']))
        print '  func %s_%d (fn M, %s) M {' % (whichCall, i, ', '.join(['a_%d M' % j for j in range(i)]))
        print '    if fn.X == nil { panic("cannot CALL on int or str") }'
        print '    switch f := fn.X.(type) {'
        print '      case %s_%d:' % (whichI, i)
        print '        return f.Call%d(%s)' % (i, ', '.join(['a_%d' % j for j in range(i)]))
        print '      case ICallV:'
        print '        return f.CallV([]M{%s}, nil, nil, nil)' % ', '.join(['a_%d' % j for j in range(i)])
        print '    }'
        print '    panic(fmt.Sprintf("No way to call: %v", fn))'
        print '  }'
        print ''

    if self.internal:
      self.internal.close()

    # BEGIN Just a comment {
    def describe(t):
      if type(t) is parse.Tvar:
        return t.name
      elif type(t) is parse.Traw:
        return t.raw
      else:
        return '-:' + str(type(t)) + ':' + repr(t)

    for th in tree.things:
      if type(th) == parse.Tdef:
        # name, args, typs, rettyp, dflts, star, starstar, body.
        jtyps = [[describe(e) for e in (t if t else [])] for t in th.typs] if th.typs else None
        j = dict(what='func', name=th.name, star=th.star, starstar=th.starstar, typs=jtyps)
        print '//@@// %s %s' % (th.name, j)
      elif type(th) == parse.Tclass:
        for m in th.things:
          if type(m) == parse.Tdef:
            jtyps = [[describe(e) for e in (t if t else [])] for t in m.typs] if m.typs else None
            j = dict(what='meth', cls=th.name, name=m.name, star=m.star, starstar=m.starstar, typs=jtyps)
            print '//@@// %s.%s %s' % (th.name, m.name, j)
    pass
    # END Just a comment }

  def Gloss(self, th):
    print '// @ %d @ %d @ %s' % (th.where, th.line, self.CurrentFuncNick())

  def Ungloss(self, th):
    print '// $ %d $ %d $' % (th.where, th.line)

  def Vexpr(self, p):
    print ' _ = %s' % p.a.visit(self)

  def AssignFieldAFromRhs(self, a, rhs, pragma):
      lhs = a.p.visit(self)
      if type(lhs) is Zself:  # Special optimization for self.
        self.instvars[a.field] = True
        lhs = 'self.M_%s' % a.field
        print '   %s = %s' % (lhs, rhs)
      elif type(lhs) is Zimport:  # For module variables.
        if lhs.imp.imported[0] == 'go':
          print '  reflect.ValueOf(& %s.%s).Elem().Set( reflect.ValueOf(%s.Contents()).Convert(reflect.TypeOf(%s.%s)))' % (
              lhs, a.field, rhs, lhs, a.field)
        else:
          print '   %s.G_%s = %s' % (lhs, a.field, rhs)
      else:
        self.setNeeded[a.field] = True
        letterF = 'F' if self.internal or a.field in gen_internals.InternalSetters else 'f'
        print '   %s_SET_%s(%s, %s)' % (letterF, a.field, lhs, rhs)

  def AssignItemAFromRhs(self, a, rhs, pragma):
        p = a.a.visit(self)
        q = a.x.visit(self)
        print '   (%s).SetItem(%s, %s)' % (p, q, rhs)

  def AssignTupleAFromB(self, a, b, pragma):
        serial = self.Serial('detuple')
        tmp = parse.Tvar(serial)
        parse.Tassign(tmp, b).visit(self)

        print '   len_%s := %s.Len()' % (serial, tmp.visit(self))
        print '   if len_%s != %d { panic(fmt.Sprintf("Assigning object of length %%d to %%d variables, in destructuring assignment.", len_%s, %d)) }' % (serial, len(a.xx), serial, len(a.xx))

        i = 0
        for x in a.xx:
          if type(x) is parse.Tvar and x.name == '_':
            pass # Tvar named '_' is the bit bucket;  don't Tassign.
          else:
            parse.Tassign(x, parse.Tgetitem(tmp, parse.Tlit('N', i))).visit(self)
          i += 1

  def AssignAFromB(self, a, b, pragma):
    # Resolve rhs first.
    type_a = type(a)

    if type_a is parse.Titems or type_a is parse.Ttuple:
      return self.AssignTupleAFromB(a, b, pragma)

    lhs = '?lhs?'
    rhs = b.visit(self)

    if type_a is parse.Tfield:
      return self.AssignFieldAFromRhs(a, rhs, pragma)

    elif type_a is parse.Tgetitem:  # p[q] = rhs
      return self.AssignItemAFromRhs(a, rhs, pragma)

    elif type_a is parse.Tvar:
      # Are we in a function scope?
      if self.scope is not None and a.name not in self.force_globals:
        # Inside a function.
        if self.scope.get(a.name):
          lhs = self.scope[a.name]
        else:
          lhs = 'v_%s' % a.name
          print '// @@@@@@ Creating var "%s" in scope @@@@@@' % a.name
          self.scope[a.name] = lhs
      else:
        # At the module level.
        lhs = a.visit(self)
        self.glbls[a.name] = ('M', 'None')

    elif type_a is parse.Traw:
      lhs = a.raw

    else:
      raise Exception('Weird Assignment, a class is %s, Left is (%s) (%s) Right is (%s) (%s)' % (type(a).__name__, a, a.visit(self), b, b.visit(self)))

    # Assign the variable, unless it is the magic un-assignable rye_rye.
    if str(lhs) != 'G_rye_rye':
      print '   %s = %s' % (lhs, rhs)

  def Vassign(self, p):
    # a, b, pragma
    self.AssignAFromB(p.a, p.b, p.pragma)

  def Vprint(self, p):
    # w, xx, saying, code
    vv = [a.visit(self) for a in p.xx.xx]
    if p.saying:
      # We are not concerned with the trailing_comma if saying.
      where = '[%s:%d %s.%s]' % (
          self.modname, p.line,
          self.cls.name if self.cls else '',
          self.func.name if self.func else '',
          )
      if self.cls:
        print '   fmt.Fprintln(%s, "## %s %s ", self.ShortPointerHashString(), " # ", %s.Repr())' % (
            '(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStderr()',
            where,
            str(p.code).replace('"', '\\"'),
            '.Repr(), "#", '.join([str(v) for v in vv]))
      else:
        print '   fmt.Fprintln(%s, "## %s %s # ", %s.Repr())' % (
            '(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStderr()',
            where,
            str(p.code).replace('"', '\\"'),
            '.Repr(), "#", '.join([str(v) for v in vv]))
    else:
      if p.xx.trailing_comma:
        printer = self.Serial('printer')
        print '%s := %s' % (
            printer,
            'M(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()')
        for i in range(len(vv)):
          print 'io.WriteString(%s, %s.String()) // i=%d' % (
              printer, str(vv[i]), i)
          print 'io.WriteString(%s, " ")' % printer
      else:
        if vv:
          print '   fmt.Fprintln(%s, %s.String())' % (
              'M(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()',
              '.String(), '.join([str(v) for v in vv]))
        else:
          print '   fmt.Fprintln(%s, "")' % (
              'M(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()')


  def Vimport(self, p):
    # imported, alias, fromWhere
    print '//Vimport: %s %s %s' % (p.imported, p.alias, p.fromWhere)
    if self.glbls.get(p.alias):
      t, v = self.glbls.get(p.alias)
      if t != '*PModule':
        raise Exception("Import alias %s already used", p.alias)

    self.imports[p.alias] = p
    vec = p.imported
    if not vec[0] == 'go':
      # Modules already contain protections against evaling more than once.
      print '   G_%s = i_%s.Eval_Module() ' % (p.alias, p.alias)

  def Vassert(self, p):
    where = '%s:%s %s.%s' % (
        self.modname, str(p.line),
        self.cls.name if self.cls else '',
        self.func.name if self.func else '',
        )
    # TODO: A way to skip 'assert' but still execute 'must'.
    print 'if %s {' % ('true' if p.is_must else 'SkipAssert == 0')

    if p.fails:
      print '''
        func() {
          defer func() {
            r := recover()
            if r == nil {
              panic(fmt.Sprintf("Missing expected Exception:  %%s", %s))
            }
            return
          }()
        _ = M(%s)
        }()
''' % ( GoStringLiteral(p.code), p.x.visit(self))
      # TODO:  Check regexp of exception.

    elif p.y is None and type(p.x) == parse.Top and p.x.op in parse.REL_OPS.values():
      # Since message is empty, print LHS, REL_OP, and RHS, since we can.
      a = p.x.a.visit(self)
      b = p.x.b.visit(self)
      sa = self.Serial('left')
      sb = self.Serial('right')
      print '   %s, %s := %s, %s' % (sa, sb, a, b)
      print '   if ! (%s.%s(%s)) {' % (sa, p.x.op, sb)
      print '     panic(fmt.Sprintf("Assertion Failed [%s]:  (%%s) ;  left: (%%s) ;  op: %%s ;  right: (%%s) ", %s, %s.Repr(), "%s", %s.Repr() ))' % (
          where, GoStringLiteral(p.code), sa, p.x.op, sb, )
      print '   }'
    else:
      print '   if ! (%s) {' % DoBool(p.x.visit(self))
      print '     panic(fmt.Sprintf("Assertion Failed [%s]:  %%s ;  message=%%s", %s, M(%s).String() ))' % (
          where, GoStringLiteral(p.code), "None" if p.y is None else p.y.visit(self) )
      print '   }'
    print '}'

  # try/except/finally:  tr, exvar, ex, fin
  def Vtry(self, p):
    # call, body
    serial = self.Serial('try')

    if p.fin:
      print '  // BEGIN OUTER FINALLY %s' % serial
      print '  fin_func_%s := func() {' % serial
      print '    // BEGIN FINALLY %s' % serial
      print '    fin_save_r := recover()'
      p.fin.visit(self)
      print '    if fin_save_r != nil { panic(fin_save_r) }'
      print '    // END FINALLY %s' % serial
      print '  }'
      print '  fin_ret_%s := func() M { defer fin_func_%s()' % (serial, serial)

    if p.ex:
      print '''
         // BEGIN OUTER EXCEPT %s
         %s_try := func() (%s_z M) {
           defer func() {
             r := recover()
             if r != nil {
               PrintStackFYIUnlessEOFBecauseExcept(r)
               %s_z = func() M {
               // BEGIN EXCEPT
      ''' % (serial, serial, serial, serial)
      # Assign, for the side effect of var creation.
      if p.exvar:
        parse.Tassign(p.exvar, parse.Traw('MkRecovered(r)')).visit(self)

      p.ex.visit(self)

      print '''
                 return MissingM
               // END EXCEPT
               }()
               return
             }
           }()
      '''
    print '// BEGIN TRY %s' % serial
    p.tr.visit(self)
    print '// END TRY %s' % serial

    if p.ex:
      print '''
           return MissingM
         }()
         if %s_try != MissingM { return %s_try }
         // END OUTER EXCEPT %s
      ''' % (serial, serial, serial)

    if p.fin:
      print '    return MissingM'
      print '  }()'
      print '  if fin_ret_%s != MissingM { return fin_ret_%s }' % (serial, serial)
      print '  // END OUTER FINALLY %s' % serial

  def Vlambda(self, p):
    # lvars, expr, where
    lamb = self.Serial('__lambda__')
    ret = parse.Treturn([p.expr])
    ret.where, ret.line, ret.gloss = p.where, p.line, 'lambda'
    suite = parse.Tsuite([ret])
    suite.where, suite.line, suite.gloss = p.where, p.line, 'lambda'

    if type(p.lvars) == parse.Titems:
      t = parse.Tdef(lamb, [x.name for x in p.lvars.xx], NoTyps, NoTyp, [None for x in p.lvars.xx], '', '', suite)
    elif type(p.lvars) == parse.Tvar:
      t = parse.Tdef(lamb, [p.lvars.name], NoTyps, NoTyp, [None], '', '', suite)
    else:
      raise Exception("Bad p.lvars type: %s" % type(p.lvars))

    t.where, t.line, t.gloss = p.where, p.line, 'lambda'
    t.visit(self)
    return parse.Tvar(lamb).visit(self)

  def Vforexpr(self, p):
    # Tforexpr(z, vv, ll, cond)
    i = self.Serial('_')
    ptv = p.ll.visit(self)
    print '''
   forexpr%s := func () M { // around FOR EXPR
     var zz%s []M
     var nexter%s Nexter = %s.Iter()
     enougher%s, canEnough%s := nexter%s.(Enougher)
     if canEnough%s {
             defer enougher%s.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx_%s, more_%s := nexter%s.Next()
       if !more_%s {
         break
       }
       // BEGIN FOR EXPR
''' % (i, i, i, ptv, i, i, i, i, i, i, i, i, i)

    parse.Tassign(p.vv, parse.Traw("ndx_%s" % i)).visit(self)

    if p.cond:
      print '  if %s {' % DoBool(p.cond.visit(self))

    print '  zz%s = append(zz%s, %s)' % (i, i, p.z.visit(self))

    if p.cond:
      print '  }'

    print '''
       // END FOR EXPR
     }
     return MkList(zz%s)
   }() // around FOR EXPR
   _ = forexpr%s  // Workaround a bug in nested forexprs.
''' % (i, i)
    return 'forexpr%s' % i

  def optimized_for_range(self, var, call, b):
    #print "// xxx optimized_for_range:", repr(var), repr(call), repr(b)
    i = self.Serial('for_range')
    if len(call.args) != 1:
      raise Exception('Exactly one arg required in optimized for...range(): ' + repr(call.args))
    if call.names[0]:
      raise Exception('No names allowed in optimized for...range()')
    if call.star:
      raise Exception('No *args allowed in optimized for...range()')
    if call.starstar:
      raise Exception('No **kw allowed in optimized for...range()')

    a0 = call.args[0]
    n = DoInt(a0.visit(self))  # General case.
    if type(a0) == parse.Tlit and a0.k == 'I':
      n = a0.v  # Optimized literal int case.

    print '''
      var i_%s int64
      var n_%s int64 = %s
      for i_%s = int64(0); i_%s < n_%s; i_%s++ {
        var tmp_%s M = MkInt(i_%s)
''' % (i, i, n, i, i, i, i, i, i)
    parse.Tassign(var, parse.Traw("tmp_%s" % i)).visit(self)
    print '   // Begin optimized_for_range Block'
    b.visit(self)
    print '   // End optimized_for_range Block'
    print '}'

  def Vfor(self, p):
    # var, t, b.

    # Optimization: for range(int)
    if type(p.t) == parse.Tcall and type(p.t.fn) == parse.Tvar and (p.t.fn.name == 'range' or p.t.fn.name == 'xrange'):
      return self.optimized_for_range(p.var, p.t, p.b)

    # Else normal case.
    i = self.Serial('_')
    ptv = p.t.visit(self)
    print '''
   for_returning%s := func () M { // around FOR
     var nexter%s Nexter = %s.Iter()
     enougher%s, canEnough%s := nexter%s.(Enougher)
     if canEnough%s {
             defer enougher%s.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx_%s, more_%s := nexter%s.Next()
       if !more_%s {
         break
       }
       // BEGIN FOR
''' % (i, i, ptv, i, i, i, i, i, i, i, i, i)

    parse.Tassign(p.var, parse.Traw("ndx_%s" % i)).visit(self)

    p.b.visit(self)

    print '''
       // END FOR
     }
     return MissingM
   }() // around FOR
   if for_returning%s != MissingM { return for_returning%s }
''' % (i, i)

  # New "with defer".    (call, body)
  def Vwithdefer(self, p):
    # call, body
    var = self.Serial('with_defer_returning')
    immanentized = self.ImmanentizeCall(p.call, 'defer')
    print '  %s := func() M { defer %s' % (var, immanentized.visit(self))
    p.body.visit(self)
    print '    return MissingM'
    print '  }()'
    print '  if %s != MissingM { return %s }' % (var, var)

  def Vglobal(self, p):
    pass

  def Vswitch(self, p):
    # (self, a, cases, clauses, default_clause):
    serial = self.Serial('sw')
    self.Gloss(p)
    if p.a:
      print '   %s := M(%s)' % (serial, p.a.visit(self))
      print '   _ = %s' % serial
    self.Ungloss(p)
    print '   switch true {'
    for ca, cl in zip(p.cases, p.clauses):
      self.Gloss(ca)
      if p.a:
        print '      case %s.EQ(%s): {' % (serial, ca.visit(self))
      else:
        print '      case %s: {' % DoBool(ca.visit(self))
      self.Ungloss(ca)
      cl.visit(self)
      print '      }  // end case'
    if p.default_clause:
      print '      default: {'
      p.default_clause.visit(self)
      print '      }  // end case'

    print '   } // end switch'

  def Vif(self, p):
    # Special case for "if rye_rye":
    if type(p.t) is parse.Tvar and p.t.name == 'rye_rye':  # Under "if rye_rye:" ...
      print '  // { // if rye_rye:'
      p.yes.visit(self)
      print '  // } // endif rye_rye'
      return

    # Normal case.
    if p.varlist:
      print '   if if_tmp := %s ; if_tmp.Bool() {' % p.t.visit(self)
      self.AssignAFromB(p.varlist, parse.Traw('if_tmp'), None)
    else:
      print '   if %s {' % DoBool(p.t.visit(self))
    p.yes.visit(self)
    if p.no:
      print '   } else {'
      p.no.visit(self)
    print '   }'

  def Vwhile(self, p):
    print '   for %s {' % DoBool(p.t.visit(self))
    p.yes.visit(self)
    print '   }'

  def Vreturn(self, p):
    if p.aa is None:
      print '   return None '
    else:
      vv = [a.visit(self) for a in p.aa]
      if len(vv) == 1:
        print '   return %s ' % vv[0]
      else:
        print '   return Entuple( %s )' % ', '.join(vv)

  def Vyield(self, p):
    if p.aa is None:
      print '   gen.Yield( None )'
    else:
      vv = [a.visit(self) for a in p.aa]
      if len(vv) == 1:
        print '   gen.Yield( %s )' % vv[0]
      else:
        print '   gen.Yield( Entuple( %s ) )' % ', '.join(vv)
    print '       { wantMore := gen.Wait()'
    print '         if !wantMore {'
    print '           gen.Finish()'
    print '           return None'
    print '       }}'

  def Vbreak(self, p):
    print '   break'

  def Vcontinue(self, p):
    print '   continue'

  def Vraise(self, p):
    print '   panic( M(%s) )' % p.a.visit(self)

  def Vdel(self, p):
    if type(p.listx) == parse.Titems:
      for e in p.listx.items.xx:
        self.Vdel(e)

    elif type(p.listx) == parse.Tgetitem:
      print "%s.DelItem(%s)" % (p.listx.a.visit(self), p.listx.x.visit(self))

    elif type(p.listx) == parse.Tgetitemslice:
      print "%s.DelItemSlice(%s, %s)" % (
        p.listx.a.visit(self),
        p.listx.x.visit(self),
        p.listx.y.visit(self))

    else:
      raise Exception("not implemented: del %s" % repr(p.listx))

  def LitIntern(self, v, key, code):
    if not self.lits.get(key):
      self.lits[key] = code
    return Zlit(v, key)

  def Vraw(self, p):
    if p.raw == 'False':
      return Ybool('false', 'False')
    if p.raw == 'True':
      return Ybool('true', 'True')
    return p.raw

  def Vlit(self, p):
    if p.k == 'N':
      z = p.v
      key = 'litI_' + CleanIdentWithSkids(str(z))
      code = 'MkInt(%s)' % str(z)
      lit = self.LitIntern(z, key, code)
      return Yint(str(z), lit)
    elif p.k == 'F':
      z = p.v
      key = 'litF_' + CleanIdentWithSkids(str(z))
      code = 'MkFloat(%s)' % str(z)
      lit = self.LitIntern(z, key, code)
      return Yfloat(str(z), lit)
    elif p.k == 'S':
      z = parse.DecodeStringLit(p.v)
      key = 'litS_' + CleanIdentWithSkids(z)
      golit = GoStringLiteral(z)
      code = 'MkStr( %s )' % golit
      lit = self.LitIntern(golit, key, code)
      return Ystr(golit, lit)
    else:
      raise Exception('Unknown Vlit', p.k, p.v)

  def Vop(self, p):
    if p.returns_bool:

      # Optimizations.
      if p.op == 'EQ':
        return DoEQ(p.a.visit(self), p.b.visit(self))
      if p.op == 'NE':
        return DoNE(p.a.visit(self), p.b.visit(self))
      if p.op == 'LT':
        return DoLT(p.a.visit(self), p.b.visit(self))
      if p.op == 'LE':
        return DoLE(p.a.visit(self), p.b.visit(self))
      if p.op == 'GT':
        return DoGT(p.a.visit(self), p.b.visit(self))
      if p.op == 'GE':
        return DoGE(p.a.visit(self), p.b.visit(self))

      return Ybool('(/*Vop returns bool*/%s.%s(%s))' % (p.a.visit(self), p.op, p.b.visit(self)), None)
    if p.b:

      # Optimizations.
      if p.op == 'Add':
        return self.DoAdd(p.a.visit(self), p.b.visit(self))
      if p.op == 'Sub':
        return DoSub(p.a.visit(self), p.b.visit(self))
      if p.op == 'Mul':
        return DoMul(p.a.visit(self), p.b.visit(self))
      if p.op == 'Div':
        return DoDiv(p.a.visit(self), p.b.visit(self))
      #if p.op == 'IDiv':
      #  return DoIDiv(p.a.visit(self), p.b.visit(self))
      if p.op == 'Mod':
        return DoMod(p.a.visit(self), p.b.visit(self))
      return ' %s.%s(%s) ' % (p.a.visit(self), p.op, p.b.visit(self))
    else:
      return ' %s.%s() ' % (p.a.visit(self), p.op)

  def DoAdd(self, a, b):
    if type(a) != str:
      z = a.DoAdd(b)
      if z: return z
    v = self.Serial('doAdd')
    print 'var %s_left M = %s' % (v, str(a))
    # self.Record('%s_left' % v)
    print 'var %s_right M = %s' % (v, str(b))
    # self.Record('%s_right' % v)

    tv = self.TempVar(v)
    #if type(tv) is Fint:
    #  tv.DoAdd3('%s_left' % v, '%s_right' % v)
    #  return str(tv)
    if type(tv) is str:
      z = '(/*DoAdd*/%s_left.Add(%s_right))' % (v, v)
      print '%s = %s' % (tv, z)
      self.RecordOp(tv, '%s_left' % v, '%s_right' % v, 'Add')
      return tv
    else:
      raise Exception('Unknown tv: %s' % tv)

  def TempVar(self, name):
    #if self.recorded:
    #  if True or self.recorded.get('%s/%s/<func int>'):
    #    print 'var %s FInt' % name
    #    print '%s.Fast.Me() = &%s.Fast' % (name, name)
    #    return Fint(name)

    print 'var %s M' % name
    return name

  def Record(self, v):
    pass
    #if self.recording:
    #  print 'if Recording != nil { fmt.Fprintf(Recording, "{\t%s\t%s\t%%s\t}\\n", M(%s).PType().String()) }' % (self.modname, v, v)

  def RecordOp(self, c, a, b, op):
    pass
    #if self.recording:
    #  print 'if Recording != nil { fmt.Fprintf(Recording, "{\t%s\t%s\t%%s\t%%s\t%%s\t%s}\\n", M(%s).PType().String(),M(%s).PType().String(),  M(%s).PType().String(), ) }' % (self.modname, c, op, c, a, b, )

  def Vboolop(self, p):
    if p.b is None:
      return Ybool('(/*Vboolop*/  %s (%s)) ' % (p.op, DoBool(p.a.visit(self))), None)
    elif p.op=='&&':
      s = self.Serial('andand')
      print '%s := func() M {' % s
      print '  var z M = %s' % p.a.visit(self)
      print '  if z.Bool() { z = %s }' % p.b.visit(self)
      print '  return z'
      print '}'
      return '%s()' % s
    elif p.op=='||':
      s = self.Serial('oror')
      print '%s := func() M {' % s
      print '  var z M = %s' % p.a.visit(self)
      print '  if !z.Bool() { z = %s }' % p.b.visit(self)
      print '  return z'
      print '}'
      return '%s()' % s
    else:
      raise Exception('notreached(Vboolop)')
      # This is how we used to do it, but short-circuit values did not work:
      #return Ybool('(/*Vboolop*/ %s %s (%s)) ' % (DoBool(p.a.visit(self)), p.op, DoBool(p.b.visit(self))), None)

  def Vcondop(self, p):  # b if a else c
    s = self.Serial('cond')
    print '%s := func (a bool) M { if a {' % s
    print 'return %s' % p.b.visit(self)
    print '}'
    print 'return %s' % p.c.visit(self)
    print '}'
    return ' %s(%s) ' % (s, DoBool(p.a.visit(self)))

  def Vgetitem(self, p):
    return ' %s.GetItem(%s) ' % (p.a.visit(self), p.x.visit(self))

  def Vgetitemslice(self, p):
    return ' %s.GetItemSlice(%s, %s, %s) ' % (
        p.a.visit(self),
        'None' if p.x is None else p.x.visit(self),
        'None' if p.y is None else p.y.visit(self),
        'None' if p.z is None else p.z.visit(self))

  def Vtuple(self, p):
    return 'MkTupleV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vlist(self, p):
    return 'MkListV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vdict(self, p):
    return 'MkDictV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vset(self, p):
    return 'MkSetV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vvar(self, p):
    if p.name == 'self':
      return Zself(p, 'self')
    if p.name == 'super':
      return Zsuper(p, 'super')
    if p.name in self.force_globals:
      return Zglobal(p, '/*force_globals*/G_%s' % p.name)
    if p.name in self.imports:
      return Zimport(p, 'i_%s' % p.name, self.imports[p.name])
    if self.scope and p.name in self.scope:
      return Zlocal(p, self.scope[p.name])
    if p.name in RYE_SPECIALS:
      return Zspecial(p, 'G_%s' % p.name)
    return Zglobal(p, 'G_%s' % p.name)

  def ImmanentizeCall(self, p, why):
    "Eval all args of Tcall now and return new Tcall, for defer or go."
    s = self.Serial(why)
    print '%s_fn := M( %s )' % (s, p.fn.visit(self))
    n = len(p.args)
    i = 0
    for a in p.args:
      print '%s_a%d := M( %s )' % (s, i, a.visit(self))
      i += 1

    if p.star:
      print '%s_star := M( %s )' % (s, p.star.visit(self))
    if p.starstar:
      print '%s_starstar := M( %s )' % (s, p.starstar.visit(self))

    return parse.Tcall(
        parse.Traw('%s_fn' % s),
        [parse.Traw('%s_a%d' % (s,i)) for i in range(n)],
        p.names,
        parse.Traw('%s_star' % s) if p.star else p.star,
        parse.Traw('%s_starstar' % s) if p.starstar else p.starstar,
    )

  def Vgo(self, p):
    immanentized = self.ImmanentizeCall(p.fcall, 'gox')
    return 'MkPromise(func () M { return %s })' % immanentized.visit(self)

  def OptimizedGoCall(self, ispec, args, qfunc):
    z = []
    z.append( '// BEGIN OptimizedGoCall: %s TAKES %s RETURNS %s' % (ispec, qfunc.takes, qfunc.rets) )
    s = self.Serial('opt_go_call')

    ins = []
    for i in range(len(qfunc.takes)):
      t = qfunc.takes[i]
      if t == 'string':
        v = '%s(%s)' % (t, DoStr(args[i]))
      elif t == '[]string':
        v = 'ListToStrings(%s.List())' % args[i]
      elif t == '[]uint8':
        v = '%s(%s)' % (t, DoByt(args[i]))
      elif t == 'bool':
        v = '%s(%s)' % (t, DoBool(args[i]))
      elif t in INTLIKE_GO_TYPES:
        v = '%s(%s)' % (t, DoInt(args[i]))
      elif t in FLOATLIKE_GO_TYPES:
        v = '%s(%s)' % (t, DoFloat(args[i]))
      else:
        raise Exception("OptimizedGoCall: Not supported yet: takes: %s" % t)

      var = '%s_t_%d' % (s, i)
      z.append( 'var %s %s = %s' % (var, t, v) )
      ins.append(var)

    outs = ['%s_r_%d' % (s, i) for i in range(len(qfunc.rets))]
    if qfunc.rets:
      assigns = '%s :=' % ', '.join(outs)
    else:
      assigns = ''
    z.append( '%s /*OptimizedGoCall OK*/ %s(%s) // OptimizedGoCall OK' % (assigns, ispec, ', '.join(ins)) )

    # Handle final magic error return.
    rets = qfunc.rets
    magic_error = False
    if rets and rets[-1] == 'error':
      magic_error = outs[-1]
      z.append( 'if %s != nil { panic(%s) }' % (magic_error, magic_error) )
      outs = outs[:-1]
      rets = rets[:-1]

    if rets:
      results = []
      for i in range(len(rets)):
        r = rets[i]
        if r == 'bool':
          v = Ybool(outs[i], None)
        elif r == 'string':
          v = Ystr(outs[i], None)
        elif r == '[]string':
          v = 'MkStrs(%s)' % outs[i]
        elif r == '[]uint8':
          v = Ybyt(outs[i], None)
        elif r in INTLIKE_GO_TYPES:
          v = Yint(outs[i], None)
        elif r in FLOATLIKE_GO_TYPES:
          v = Yfloat(outs[i], None)
        else:
          v = 'AdaptForReturn(reflect.ValueOf(%s))' % outs[i]
          #raise Exception("OptimizedGoCall: Not supported yet: returns: %s" % r)
        results.append(v)

      if len(results) > 1:
        z.append( '%s_retval := MkTuple([]M{%s})' % (s, ', '.join([str(r) for r in results])) )
        for line in z: print line # After all exceptions could have occurred in this subroutine.
        return '%s_retval' % s
      else:
        for line in z: print line # After all exceptions could have occurred in this subroutine.
        return results[0]

    else:
      for line in z: print line # After all exceptions could have occurred in this subroutine.
      return 'None'


  def Vcall(self, p):
    # fn, args, names, star, starstar

    def NativeGoTypeName(a):
        if type(a) is parse.Tfield:
          return '%s.%s' % (a.p.visit(self), a.field)
        elif type(a) is parse.Tvar:
          return a.name
        else:
          raise Exception('Strange thing for go_type: ' + a)

    n = len(p.args)
    self.maxNumCallArgs = max(self.maxNumCallArgs, n)

    # arglist_thunk: As a text sequence with commas.
    arglist_thunk = lambda: ', '.join(["%s" % (a.visit(self)) for a in p.args])
    # argvec_thunk: As a python list, which could be joined with commas, or used separately.
    argvec_thunk = lambda: [a.visit(self) for a in p.args]

    #print '// Vcall: fn:', repr(p.fn)
    #print '// Vcall: args:', repr(p.args)
    #print '// Vcall: names:', repr(p.names)
    #print '// Vcall: star:', repr(p.star)
    #print '// Vcall: starstar:', repr(p.starstar)
    if p.star or p.starstar or any(p.names):
      return 'M(%s).ToP().(ICallV).CallV([]M{%s}, %s, []KV{%s}, %s) ' % (

          p.fn.visit(self),

          ', '.join([str(p.args[i].visit(self)) for i in range(len(p.args)) if not p.names[i]]),  # fixed args with no names.

          ('(%s).List()' % p.star.visit(self)) if p.star else 'nil',

          ', '.join(['KV{"%s", %s}' % (p.names[i], p.args[i].visit(self)) for i in range(n) if p.names[i]]),  # named args.

          ('(%s).Dict()' % p.starstar.visit(self)) if p.starstar else 'nil',
      )

    if type(p.fn) is parse.Tfield:  # CASE var.Meth(...)
      if type(p.fn.p) is parse.Tvar:

        if p.fn.p.name == 'super':  # CASE super.Meth(...)
          return 'self.%s.M_%d_%s(%s)' % (self.tailSup(self.sup), n, p.fn.field, arglist_thunk())

        if p.fn.p.name in self.imports:  # CASE import.Func(...)
          imp = self.imports[p.fn.p.name]

          if imp.imported[0] == 'go':  # CASE go.*: import.Func(...)

            # Try Optimization with QFunc.
            ipath = '/'.join(imp.imported[1:])
            iname = '%s.%s' % (ipath, p.fn.field)
            ispec = 'i_%s.%s' % (p.fn.p.name, p.fn.field)
            argvec = argvec_thunk()  # Call it here, so use them once!
            qfunc = goapi.QFuncs.get(iname)
            if qfunc:
              try:  # Try to use the argvec.
                return self.OptimizedGoCall(ispec, argvec, qfunc)
              except Exception as ex:
                print '//OptimizedGoCall NO: << %s.%s(%s) >>: %s' % (p.fn.p.name, p.fn.field, ', '.join([str(a) for a in argvec]), repr(ex))

            # Otherwise use reflection with MkGo().  Use the argvec.
            return 'MkGo(%s).Call(%s) ' % (ispec, ', '.join([str(a) for a in argvec]))
          else:
            return '%s_%d( i_%s.G_%s, %s) ' % (('CALL' if n<11 else 'call'), n, p.fn.p.name, p.fn.field, arglist_thunk())

      # General Method Invocation.
      key = '%d_%s' % (n, p.fn.field)
      self.invokes[key] = (n, p.fn.field)
      letterF = 'F' if self.internal or ((n, p.fn.field) in gen_internals.InternalInvokers) else 'f'

      invoker = '/*invoker*/ %s_INVOKE_%d_%s' % (letterF, n, p.fn.field)
      general = '/*General*/ %s(%s, %s) ' % (invoker, p.fn.p.visit(self), arglist_thunk())

      # GOMAXPROCS=2
      # With OPT=append: 17.37 17.71 19.736 19.686
      # Without: 18.244 18.094 18.628
      #if os.getenv('OPT') == 'append' and p.fn.field == 'append':
      if p.fn.field == 'append':
        print '{ o := %s // optimize_append' % p.fn.p.visit(self)
        print '  if o.X != nil {'
        print '    if p, ok := o.X.(*PList); ok {'
        print '      p.PP = append(p.PP, %s)' % (arglist_thunk())
        print '    } else {'
        print '      %s(o, %s)' % (invoker, arglist_thunk())
        print '    }'
        print '  } else {'
        print '    %s(o, %s)' % (invoker, arglist_thunk())
        print '} }'
        return 'None'

      return general


    zfn = p.fn.visit(self)
    if type(zfn) is Zspecial:
      if p.fn.name == 'go_type':
        assert len(p.args) == 1, 'go_type got %d args, wants 1' % len(p.args)
        return 'GoElemType(new(%s))' % NativeGoTypeName(p.args[0])
      elif p.fn.name == 'go_indirect':
        assert len(p.args) == 1, 'go_addr got %d args, wants 1' % len(p.args)
        return 'MkValue(reflect.Indirect(reflect.ValueOf(%s.Contents())))' % p.args[0].visit(self)
      elif p.fn.name == 'go_addr':
        assert len(p.args) == 1, 'go_addr got %d args, wants 1' % len(p.args)
        return 'MkGo(reflect.ValueOf(%s.Contents()).Addr())' % p.args[0].visit(self)
      elif p.fn.name == 'go_new':
        assert len(p.args) == 1, 'go_new got %d args, wants 1' % len(p.args)
        return 'MkGo(new(%s))' % NativeGoTypeName(p.args[0])
      elif p.fn.name == 'go_make':
        if len(p.args) == 1:
          return 'MkGo(make(%s))' % NativeGoTypeName(p.args[0])
        elif len(p.args) == 2:
          return 'MkGo(make(%s, int(%s.Int())))' % (NativeGoTypeName(p.args[0]), p.args[1].visit(self))
        else:
          raise Exception('go_make got %d args, wants 1 or 2' % len(p.args))
      elif p.fn.name == 'go_cast':
        assert len(p.args) == 2, 'go_cast got %d args, wants 2' % len(p.args)
        return 'GoCast(GoElemType(new(%s)), %s)' % (NativeGoTypeName(p.args[0]), p.args[1].visit(self))
      elif p.fn.name == 'go_append':
        assert len(p.args) == 2, 'go_append got %d args, wants 2' % len(p.args)
        return 'GoAppend(%s, %s)' % (p.args[0].visit(self), p.args[1].visit(self))

      elif p.fn.name == 'len':
        assert len(p.args) == 1, 'len got %d args, wants 1' % len(p.args)
        return Yint('/*Y*/int64(%s.Len())' % p.args[0].visit(self), None)
      elif p.fn.name == 'str':
        assert len(p.args) == 1, 'str got %d args, wants 1' % len(p.args)
        return Ystr('/*Y*/%s.String()' % p.args[0].visit(self), None)
      elif p.fn.name == 'repr':
        assert len(p.args) == 1, 'repr got %d args, wants 1' % len(p.args)
        return Ystr('/*Y*/%s.Repr()' % p.args[0].visit(self), None)
      elif p.fn.name == 'int':
        assert len(p.args) == 1, 'int got %d args, wants 1' % len(p.args)
        return Yint('/*Y*/%s.ForceInt()' % p.args[0].visit(self), None)
      elif p.fn.name == 'float':
        assert len(p.args) == 1, 'float got %d args, wants 1' % len(p.args)
        return Yfloat('/*Y*/%s.ForceFloat()' % p.args[0].visit(self), None)

      else:
        raise Exception('Undefind builtin: %s' % p.fn.name)

    if type(zfn) is Zglobal and zfn.t.name in self.defs:
      fp = self.defs[zfn.t.name]
      if not fp.star and not fp.starstar:
        want = len(fp.args)
        arglist = arglist_thunk()

        missing = want - n # Number of arguments missing.
        if missing and all(fp.dflts[-missing:]):
          # Can provide the missing arguments from dflts.
          arglist += ', ' + ','.join([str(x.visit(self)) for x in fp.dflts[-missing:]])
          n += missing

        if n != want:
          raise Exception('Calling global function "%s", got %d args, wanted %d args' % (zfn.t.name, n, want))
        return 'G_%d_%s(%s) ' % (n, zfn.t.name, arglist)

    if type(zfn) is Zsuper:  # for calling super-constructor.
      return 'self.%s.M_%d___init__(%s) ' % (self.tailSup(self.sup), n, arglist_thunk())

    return '%s_%d( M(%s), %s )' % (('CALL' if n<11 else 'call'), n, p.fn.visit(self), arglist_thunk())

  def Vfield(self, p):
    # p, field
    x = p.p.visit(self)
    if type(x) is Zsuper:
      raise Exception('Special syntax "super" not used with Function Call syntax')
    if type(x) is Zself and not self.cls:
      raise Exception('Using a self field but not in a class definition: field="%s"' % p.field)
    if type(x) is Zself and self.instvars.get(p.field):  # Special optimization for self instvars.
      return 'self.M_%s' % p.field
    elif type(x) is Zimport:
      if x.imp.imported[0] == 'go':
        return ' MkGo(%s.%s) ' % (x, p.field)
      else:
        return ' %s.G_%s ' % (x, p.field)
    else:
      self.getNeeded[p.field] = True
      letterF = 'F' if self.internal or p.field in gen_internals.InternalGetters else 'f'
      return ' %s_GET_%s(%s) ' % (letterF, p.field, x)

  def Vnative(self, p):
    if self.func:
      print '// { native F'
      for s in p.ss:
        print s
      print '// } native F'
    else:
      # Append to tail, or these would go in inner_eval_module().
      self.tail.append('// { native M')
      for s in p.ss:
        self.tail.append(s)
      self.tail.append('// } native M')
    return ''

  def Vdef(self, p):
    # name, args, typs, rettyp, dflts, star, starstar, body.

    # SAVE STATUS BEFORE THIS FUNC.
    save_func = self.func
    save_yields = self.yields
    save_force_globals = self.force_globals
    self.func = p
    self.func_level += 1

    # START A PRINT BUFFER -- but not if Nested.
    nesting = None
    fn_var = None
    if self.func_level >= 2:
      nesting = self.Serial('nesting')
      # Create the local var for the fn, so it's visibile in the fn.
      fn_var = parse.Tvar(p.name)
      # Set it to None now, but it'll be filled in before it can be called.
      self.AssignAFromB(fn_var, parse.Traw('None'), None)
    else:
      buf = PushPrint()

    # LOOK AHEAD for "yield" and "global" statements.
    finder = parse.YieldAndGlobalFinder()
    finder.Vsuite(p.body)
    self.yields = finder.yields
    self.force_globals = finder.force_globals

    # Tweak args.  Record meth, if meth.
    args = p.args    # Will drop the initial 'self' element, if in a cls.
    typs = p.typs    # Will drop the initial 'self' element, if in a cls.
    dflts = p.dflts  # Will drop the initial 'self' element, if in a cls.
    if nesting:
      pass
    elif self.cls and not nesting:
      if len(p.args) > 0 and p.args[0] == 'self':  # User may omit self.
        args = p.args[1:]  # Skip self; it is assumed.
        typs = p.typs[1:]  # Skip self; it is assumed.
        dflts = p.dflts[1:]  # Skip self; it is assumed.
      self.meths[p.name] = ArgDesc(self.modname, self.cls.name, '%s.%s::%s' % (self.modname, self.cls.name, p.name), args, dflts, p.star, p.starstar, isCtor=p.isCtor)
    else:
      self.defs[p.name] = ArgDesc(self.modname, None, '%s.%s' % (self.modname, p.name), args, p.dflts, p.star, p.starstar, isCtor=p.isCtor)

    # Copy scope and add argsPlus to the new one.
    save_scope = self.scope
    if self.scope:
      self.scope = dict([(k, w) for k, w in self.scope.items()]) # Copy it.
    else:
      self.scope = {}
    for a in p.argsPlus:
      self.scope[a] = 'a_%s' % a

    #################
    # Render the body, but hold it in buf2, because we will prepend the vars.
    buf2 = PushPrint()
    if self.yields:
      print '''
        gen := NewGenerator()
        go func() {
          // Recover & repanic, printing FYI.
          defer func() {
            r := recover()
            if r != nil {
              PrintStackFYIUnlessEOFBecauseExcept(r)
              if !gen.Finished {
                gen.YieldException(r)
              } else {
                panic(r)
              }
            }
          }()

           mustBeNone := func() M {
             { wantMore := gen.Wait()
               if !wantMore {
                 gen.Finish()
                 return None
             }}
'''

    p.body.visit(self)

    if self.yields:
      print '''
            return None
          }()
          if mustBeNone != None {
             panic("Return Value in Generator must be None.")
          }
          gen.Finish()
        }()
        return MkX(&gen.PBase)
'''

    PopPrint()
    code2 = str(buf2)

    print '///////////////////////////////'
    print ''

    if self.internal and self.cls:
      # Record a synthetic 'invokes', so it gets generated with builtins.
      ikey = '%d_%s' % (len(args), p.name)
      self.invokes[ikey] = (len(args), p.name)
      self.getNeeded[p.name] = True

    letterV = 'V' if p.star or p.starstar else ''
    emptiesV = (', MkList(nil), MkDict(nil)' if args else 'MkList(nil), MkDict(nil)') if p.star or p.starstar else ''
    stars = ' %s M, %s M' % (AOrSkid(p.star), AOrSkid(p.starstar)) if p.star or p.starstar else ''

    if nesting:
      func_head = 'fn_%s := func' % nesting
      if self.cls:
        func_key = '%s__%s__%s__fn_%s' % (self.modname, self.cls.name, p.name, nesting)
      else:
        func_key = '%s__%s__fn_%s' % (self.modname, p.name, nesting)
    elif self.cls:
      gocls = self.cls.name if self.sup == 'native' else 'C_%s' % self.cls.name
      func_head = 'func (self *%s) M_%d%s_%s' % (gocls, len(args), letterV, p.name)
      func_key = '%s__%s__%s' % (self.modname, self.cls.name, p.name)
    else:
      func_head = 'func G_%d%s_%s' % (len(args), letterV, p.name)
      func_key = '%s__%s' % (self.modname, p.name)

    if not nesting:
      # TODO: Be able to emit this Counter & Init for nested functions, too.
      print 'var counter_%s int64' % func_key
      print 'func init() {CounterMap["%s"]= &counter_%s}' % (func_key, func_key)

    # Start the function.
    print ' %s(%s %s) M {' % (func_head, ' '.join(['a_%s M,' % a for a in args]), stars)
    if not nesting:
      print '  counter_%s++' % func_key

    if typs:
      # Check typs of input arguments.
      for (a, t) in zip(args, typs):
        if t:
          print '    CheckTyp("arg %s in func %s", a_%s, %s)' % (a, func_key, a, ','.join([str(e.visit(self)) for e in t]))

    for v, v2 in sorted(self.scope.items()):
      if save_scope is None or v not in save_scope:
        if v2[0] != 'a':  # Skip args
          print "   var %s M = None; _ = %s" % (v2, v2)

    if p.rettyp:
      # Start inner function for checking all types of return values.
      print '   retval := func() M { // retval func'

    # Main Body.
    print code2
    print '   return None'  # For falling off the bottom.

    if p.rettyp:
      # End inner function for checking all types of return values.
      print '   }() // retval func'
      print '   CheckTyp("return value of function %s", retval, %s)' % (func_key, (','.join([str(e.visit(self)) for e in p.rettyp])))
      print '   return retval'

    # End the function
    print ' }'
    print ''

    # Restore the old scope.
    self.scope = save_scope

    n = len(args)
    argnames = ', '.join(['"%s"' % a for a in p.args])
    defaults = ', '.join([(str(d.visit(self)) if d else 'MissingM') for d in p.dflts])

    if nesting:
      tmp = '  MForge(&pNest_%s{PNewCallable{CallSpec: &specNest_%s}, fn_%s})  ' % (
          nesting, nesting, nesting)

      self.AssignAFromB(fn_var, parse.Traw(tmp), None)


    # Now for the Nested case, START A PRINT BUFFER.
      buf = PushPrint()

    print '///////////////////////////////'
    print '// name:', p.name

    if nesting:
      print 'var specNest_%s = CallSpec{Name: "%s__%s", Args: []string{%s}, Defaults: []M{%s}, Star: "%s", StarStar: "%s"}' % (
          nesting, p.name, nesting, argnames, defaults, p.star, p.starstar)

      print ' type pNest_%s struct { PNewCallable; fn func(%s %s) M }' % (nesting, ' '.join(['a_%s M,' % a for a in args]), stars)
      print ' func (o *pNest_%s) Contents() interface{} {' % nesting
      print '   return o.fn'
      print ' }'
      if p.star or p.starstar:
        pass  # No direct pNest method; use CallV().
      else:
        print ' func (o pNest_%s) Call%d(%s) M {' % (nesting, n, ', '.join(['a%d M' % i for i in range(n)]))
        print '   return o.fn(%s)' % (', '.join(['a%d' % i for i in range(n)]))
        print ' }'
      print ''
      print ' func (o pNest_%s) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {' % nesting
      print '   argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   return o.fn(%s MkX(&star.PBase), MkX(&starstar.PBase))' % (' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return o.fn(%s)' % (', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

    elif self.cls:
      if SMALLER and n < 4 and not p.star and not p.starstar and not p.isCtor:
        # Optimize most functions to use PCall%d instead of defining a new struct.
        pass
      else:
        print ' type pMeth_%d_%s__%s struct { PNewCallable; Rcvr *%s }' % (n, self.cls.name, p.name, gocls)
        print ' func (o *pMeth_%d_%s__%s) Contents() interface{} {' % (n, self.cls.name, p.name)
        print '   return o.Rcvr.M_%d%s_%s' % (n, letterV, p.name)
        print ' }'
        print ' func (o *pMeth_%d_%s__%s) Call%d(%s) M {' % (n, self.cls.name, p.name, n, ', '.join(['a%d M' % i for i in range(n)]))
        print '   return o.Rcvr.M_%d%s_%s(%s%s)' % (n, letterV, p.name, ', '.join(['a%d' % i for i in range(n)]), emptiesV)
        print ' }'
        print ''
        print ' func (o *pMeth_%d_%s__%s) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {' % (n, self.cls.name, p.name)
        print '   argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)'
        print '   _, _, _ = argv, star, starstar'

        if p.star or p.starstar:  # If either, we always pass both.
          print '   return o.Rcvr.M_%dV_%s(%s MkX(&star.PBase), MkX(&starstar.PBase))' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
        else:  # If neither, we never pass either.
          print '   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

        print ' }'
        print ''

    else:
      print ''
      print 'var specFunc_%s = CallSpec{Name: "%s", Args: []string{%s}, Defaults: []M{%s}, Star: "%s", StarStar: "%s"}' % (
          p.name, p.name, argnames, defaults, p.star if p.star else '', p.starstar if p.starstar else '')

      if SMALLER and n < 4 and not p.star and not p.starstar and not p.isCtor:
        # Optimize most functions to use PCall%d instead of defining a new struct.
        formals = ','.join(['a%d M' % i for i in range(n)])
        actuals = ','.join(['a%d' % i for i in range(n)])
        print 'func fnFunc_%s (%s) M { return G_%d_%s(%s) }' % (p.name, formals, n, p.name, actuals)
        print ''

        self.glbls[p.name] = ('*pFunc_%s' % p.name,
                              'MForge(&PCall%d{PNewCallable{CallSpec:&specFunc_%s}, fnFunc_%s, fnFunc_%s})' % (
                                  n, p.name, p.name, p.name))
      else:
        print ' type pFunc_%s struct { PNewCallable }' % p.name
        print ' func (o *pFunc_%s) Contents() interface{} {' % p.name
        print '   return G_%s' % p.name
        print ' }'
        if p.star or p.starstar:
          pass  # No direct pFunc method; use CallV().
        else:
          print ' func (o pFunc_%s) Call%d(%s) M {' % (p.name, n, ', '.join(['a%d M' % i for i in range(n)]))
          print '   return G_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
          print ' }'
        print ''
        print ' func (o pFunc_%s) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {' % p.name
        print '   argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)'
        print '   _, _, _ = argv, star, starstar'

        # TODO: I think this is old, before named params.
        if p.star or p.starstar:  # If either, we always pass both.
          print '   return G_%dV_%s(%s MkX(&star.PBase), MkX(&starstar.PBase))' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
        else:  # If neither, we never pass either.
          print '   return G_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

        print ' }'
        print ''

        self.glbls[p.name] = ('*pFunc_%s' % p.name,
                              'MForge(&pFunc_%s{PNewCallable{CallSpec:&specFunc_%s}})' % (p.name, p.name))

    PopPrint()
    code = str(buf)
    self.tail.append(code)

    # Unsave.
    self.func = save_func
    self.yields = save_yields
    self.force_globals = save_force_globals
    self.func_level -= 1

  def qualifySup(self, sup):
    if type(sup) == parse.Tvar:
      return 'C_%s' % sup.name
    elif type(sup) == parse.Tfield:
      return 'i_%s.C_%s' % (sup.p.name, sup.field)
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def tailSup(self, sup):
    if type(sup) == parse.Tvar:
      return 'C_%s' % sup.name
    elif type(sup) == parse.Tfield:
      return 'C_%s' % sup.field
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def Vclass(self, p):
    # name, sup, things
    if self.cls:
      raise Exception("Nested classes not supported.")

    self.cls = p
    self.sup = p.sup
    self.instvars = {}
    self.meths = {}

    gocls = self.cls.name if self.sup == 'native' else 'C_%s' % self.cls.name
    # Emit all the methods of the class (and possibly other members).
    for x in p.things:
      x.visit(self)
    self.cls = None

    buf = PushPrint()

    # Emit the struct for the class.
    if self.sup != 'native':
      print '''
 type C_%s struct {
   %s
   %s
 }

 func init() {
   if Classes == nil { Classes = make(map[string]reflect.Type) }
   Classes[`%s.C_%s`] = reflect.TypeOf(C_%s{})
 }

''' % (p.name, self.qualifySup(p.sup),
       '\n'.join(['   M_%s   M' % x for x in sorted(self.instvars)]),
       self.modname if self.modname else 'main', p.name, p.name)

    if self.sup != 'native':
      print '''
 func (o *C_%s) PtrC_%s() *C_%s {
   return o
 }
''' % (p.name, p.name, p.name)

    if self.sup != 'native':
      print '''
 func (o *C_%s) PtrC_object() *C_object {
   return &o.C_object
 }
''' % (p.name, )

    # For all the instance vars
    print ''
    for iv in sorted(self.instvars):
      print ' func (o *C_%s) GET_%s() M { return o.M_%s }' % (p.name, iv, iv)
      print ' func (o *C_%s) SET_%s(x M) { o.M_%s = x }' % (p.name, iv, iv)
      print ''
    print ''

    # For all the methods
    print ''
    for m in sorted(self.meths):  # ArgDesc in self.meths[m]
      mp = self.meths[m]
      args = mp.args
      dflts = mp.dflts
      n = len(args)

      argnames = ', '.join(['"%s"' % a for a in args])
      defaults = ', '.join([(str(d.visit(self)) if d else 'MissingM') for d in dflts])

      print 'var specMeth_%d_%s__%s = CallSpec{Name: "%s::%s", Args: []string{%s}, Defaults: []M{%s}, Star: "%s", StarStar: "%s"}' % (
          n, p.name, m, p.name, m, argnames, defaults, mp.star, mp.starstar)

      if SMALLER and n < 4 and not mp.star and not mp.starstar and not mp.isCtor:
        # Optimize most functions to use PCall%d instead of defining a new struct.
        formals = ','.join(['a%d M' % i for i in range(n)])
        actuals = ','.join(['a%d' % i for i in range(n)])
        print ''

        print 'func (o *%s) GET_%s() M {' % (gocls, m)
        print '  return MForge(&PCall%d{PNewCallable{CallSpec:&specMeth_%d_%s__%s}, o.M_%d_%s, o.M_%d_%s})' % (
                                  n, n, p.name, m, n, m, n, m)
        print '}'
      else:
        print 'func (o *%s) GET_%s() M { return MForge(&pMeth_%d_%s__%s {PNewCallable{CallSpec: &specMeth_%d_%s__%s}, o})}' % (
            gocls, m, n, p.name, m, n, p.name, m)

    # Special methods for classes.
    if self.sup != 'native':
      print 'func (o *C_%s) Rye_ClearFields__() {' % p.name
      for iv in sorted(self.instvars):
        print '   o.M_%s = None' % iv
      if p.sup and type(p.sup) is parse.Tvar:
        print '// superclass:', p.sup.visit(self)
        if p.sup.name not in ['native', 'object']:
          print '   o.C_%s.Rye_ClearFields__()' % p.sup.name
      if p.sup and type(p.sup) is parse.Tfield:
        print '// superclass:', p.sup.visit(self)
        print '   o.C_%s.Rye_ClearFields__()' % p.sup.field
      print '}'

      print ''
      print 'func (o *pFunc_%s) Superclass() M {' % (p.name)
      if p.sup and type(p.sup) is parse.Tvar:
        print '  return %s' % p.sup.visit(self)
      else:
        print '  return None'
      print '}'
      print ''
      print 'func (o *C_%s) PType() M { return G_%s }' % (p.name, p.name)
      print 'func (o *pFunc_%s) Repr() string { return "%s" }' % (p.name, p.name)
      print 'func (o *pFunc_%s) String() string { return "<class %s>" }' % (p.name, p.name)
      print ''


    self.tail.append(str(buf))
    PopPrint()

  def CurrentFuncNick(self):  # Just a nickname, for diagnostics.
    cn =  self.cls.name if self.cls else ''
    fn =  self.func.name if self.func else ''
    if cn:
      return '%s.%s' % (cn, fn)
    else:
      return fn

  def Vsuite(self, p):
    for th in p.things:
      print '// @ %d @ %d @ %s' % (th.where, th.line, self.CurrentFuncNick())
      th.visit(self)
      print '// $ %d $ %d $' % (th.where, th.line)

PrinterStack = []
def PushPrint():
    PrinterStack.append(sys.stdout)
    buf = Buffer()
    sys.stdout = buf
    return buf
def PopPrint():
    sys.stdout = PrinterStack.pop()

class Buffer(object):
  def __init__(self):
    self.b = []
  def write(self, x):
    self.b.append(str(x))  # Forces immutable copy.
  def flush(self):
    pass
  def __str__(self):
    z = ''.join(self.b)
    return z

def DoSub(a, b):
  if type(a) != str:
    z = a.DoSub(b)
    if z: return z
  return '(/*DoSub*/%s.Sub(%s))' % (str(a), str(b))
def DoMul(a, b):
  if type(a) != str:
    z = a.DoMul(b)
    if z: return z
  return '(/*DoMul*/%s.Mul(%s))' % (str(a), str(b))
def DoDiv(a, b):
  if type(a) != str:
    z = a.DoDiv(b)
    if z: return z
  return '(/*DoDiv*/%s.Div(%s))' % (str(a), str(b))
#def DoIDiv(a, b):
#  if type(a) != str:
#    z = a.DoIDiv(b)
#    if z: return z
#  return '(/*DoIDiv*/%s.IDiv(%s))' % (str(a), str(b))
def DoMod(a, b):
  if type(a) != str:
    z = a.DoMod(b)
    if z: return z
  return '(/*DoMod*/%s.Mod(%s))' % (str(a), str(b))

def DoEQ(a, b):
  if type(a) != str:
    z = a.DoEQ(b)
    if z: return z
  return Ybool('(/*DoEQ*/%s.EQ(%s))' % (str(a), str(b)), None)
def DoNE(a, b):
  if type(a) != str:
    z = a.DoNE(b)
    if z: return z
  return Ybool('(/*DoNE*/%s.NE(%s))' % (str(a), str(b)), None)
def DoLT(a, b):
  if type(a) != str:
    z = a.DoLT(b)
    if z: return z
  return Ybool('(/*DoLT*/%s.LT(%s))' % (str(a), str(b)), None)
def DoLE(a, b):
  if type(a) != str:
    z = a.DoLE(b)
    if z: return z
  return Ybool('(/*DoLE*/%s.LE(%s))' % (str(a), str(b)), None)
def DoGT(a, b):
  if type(a) != str:
    z = a.DoGT(b)
    if z: return z
  return Ybool('(/*DoGT*/%s.GT(%s))' % (str(a), str(b)), None)
def DoGE(a, b):
  if type(a) != str:
    z = a.DoGE(b)
    if z: return z
  return Ybool('(/*DoGE*/%s.GE(%s))' % (str(a), str(b)), None)

def DoNot(a):
  return '/*DoNot*/!(%s)' % DoBool(a)

def DoBool(a):
  if type(a) != str:
    z = a.DoBool()
    if z: return z
  return '/*DoBool*/%s.Bool()' % a

def DoInt(a):
  if type(a) != str:
    z = a.DoInt()
    if z: return z
  return '/*DoInt*/%s.Int()' % a

def DoFloat(a):
  if type(a) != str:
    z = a.DoFloat()
    if z: return z
  return '/*DoFloat*%s.Float()' % a

def DoByt(a):
  if type(a) != str:
    z = a.DoByt()
    if z: return z
  return '/*DoByt*/%s.Bytes()' % str(a)

def DoStr(a):
  if type(a) != str:
    z = a.DoStr()
    if z: return z
  return '/*DoStr*/%s.Str()' % str(a)

#class Fbase(object):
#  def DoAdd3(self, a, b): return ''
#
#class Fint(Fbase):
#  def __init__(self, name):
#    self.name = name
#  def __str__(self):
#    return '%s.M()' % self.name
#  def Name(self):
#    return self.name
#  def DoAdd3(self, a, b):
#    sa = str(a)  # TODO non-strs.
#    print '%s_Bleft := %s' % (self.name, sa)
#    sb = str(b)  # TODO non-strs.
#    print '%s_Bright := %s' % (self.name, sb)
#    print '// (Fint::DoAdd3)'
#    print 'if %s_Bleft.PType() == G_int && %s_Bright.PType() == G_int {' % (self.name, self.name)
#    print '  %s.Fast.N = %s_Bleft.Int() + %s_Bright.Int()' % (self.name, self.name, self.name)
#    print '} else {'
#    print '  %s.Slow = %s_Bleft.Add(%s_Bright)' % (self.name, self.name, self.name)
#    print '}'
#    print ''

class Ybase(object):
  """Ybase: Future Optimized Typed values."""
  def DoBool(self): return ''
  def DoInt(self): return ''
  def DoFloat(self): return ''
  def DoByt(self): return ''
  def DoStr(self): return ''
  def DoAdd(self, b): return ''
  def DoSub(self, b): return ''
  def DoMul(self, b): return ''
  def DoDiv(self, b): return ''
  #def DoIDiv(self, b): return ''
  def DoMod(self, b): return ''
  def DoEQ(self, b): return ''
  def DoNE(self, b): return ''
  def DoLT(self, b): return ''
  def DoLE(self, b): return ''
  def DoGT(self, b): return ''
  def DoGE(self, b): return ''

class Ybool(Ybase):
  def __init__(self, y, s):
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkBool(%s)' % str(self.y)
    return str(self.s)
  def DoInt(self):
    return '/*Ybool.DoInt*/int64(%s)' % self.y
  def DoBool(self):
    return str(self.y)

class Yint(Ybase):
  def __init__(self, y, s):
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkInt(int64(%s))' % str(self.y)
    return str(self.s)
  def DoInt(self):
    return str(self.y)
  def DoFloat(self):
    return 'float64(%s)' % self.y
  def DoBool(self):
    return '/*Yint.DoBool*/(%s != 0)' % self.y

  def doArith(self, b, op):
    if type(b) is Ybool:
      return Yint('(/*YYint.doArith*/ int64(%s) %s BoolToInt64(%s) )' % (self.y, op, b.y), None)
    if type(b) is Yint:
      return Yint('(/*YYint.doArith*/ int64(%s) %s int64(%s) )' % (self.y, op, b.y), None)
    if type(b) is Yfloat:
      return Yfloat('(/*YYint.doArith*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''
  def DoAdd(self, b): return self.doArith(b, '+')
  def DoSub(self, b): return self.doArith(b, '-')
  def DoMul(self, b): return self.doArith(b, '*')
  def DoDiv(self, b): return self.doArith(b, '/')
  def DoMod(self, b): return self.doArith(b, '%')

  def doRelop(self, b, op):
    if type(b) is Ybool:
      return Ybool('(/*YYint.doRelop*/ int64(%s) %s BoolToInt64(%s) )' % (self.y, op, b.y), None)
    if type(b) is Yint:
      return Ybool('(/*YYint.doRelop*/ int64(%s) %s int64(%s) )' % (self.y, op, b.y), None)
    if type(b) is Yfloat:
      return Ybool('(/*YYint.doRelop*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''

  def DoEQ(self, b): return self.doRelop(b, '==')
  def DoNE(self, b): return self.doRelop(b, '!=')
  def DoLT(self, b): return self.doRelop(b, '<')
  def DoLE(self, b): return self.doRelop(b, '<=')
  def DoGT(self, b): return self.doRelop(b, '>')
  def DoGE(self, b): return self.doRelop(b, '>=')

class Yfloat(Ybase):
  def __init__(self, y, s):
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkFloat(float64(%s))' % str(self.y)
    return str(self.s)
  def DoFloat(self):
    return str(self.y)
  def DoBool(self):
    return '/*Yfloat.DoBool*/(%s != 0)' % self.y

  def doArith(self, b, op):
    if type(b) is Ybool:
      return Yfloat('(/*YYfloat.doArith*/ float64(%s) %s BoolToFloat64(%s) )' % (self.y, op, b.y), None)
    if type(b) in [Yfloat, Yint]:
      return Yfloat('(/*YYfloat.doArith*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''
  def DoAdd(self, b): return self.doArith(b, '+')
  def DoSub(self, b): return self.doArith(b, '-')
  def DoMul(self, b): return self.doArith(b, '*')
  def DoDiv(self, b): return self.doArith(b, '/')
  def DoMod(self, b): return self.doArith(b, '%')

  def doRelop(self, b, op):
    if type(b) is Ybool:
      return Ybool('(/*YYfloat.doRelop*/ float64(%s) %s BoolToFloat64(%s) )' % (self.y, op, b.y), None)
    if type(b) in [Yfloat, Yint]:
      return Ybool('(/*YYfloat.doRelop*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''
  def DoEQ(self, b): return self.doRelop(b, '==')
  def DoNE(self, b): return self.doRelop(b, '!=')
  def DoLT(self, b): return self.doRelop(b, '<')
  def DoLE(self, b): return self.doRelop(b, '<=')
  def DoGT(self, b): return self.doRelop(b, '>')
  def DoGE(self, b): return self.doRelop(b, '>=')

class Ystr(Ybase):
  def __init__(self, y, s):
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkStr(%s)' % str(self.y)
    return str(self.s)
  def DoByt(self):
    return '/*Ystr.DoByt*/[]byte(%s)' % self.y
  def DoStr(self):
    return str(self.y)
  def DoBool(self):
    return '/*Ystr.DoBool*/(%s != "")' % self.y

  def doRelop(self, b, op):
    if type(b) is Ystr:
      return Ybool('(/*YYstr.doRelop*/ string(%s) %s string(%s) )' % (self.y, op, b.y), None)
    if type(b) is Ybyt:
      return Ybool('(/*YYstr.doRelop*/ string(%s) %s string(%s) )' % (self.y, op, b.y), None)
    return ''

  def DoEQ(self, b): return self.doRelop(b, '==')
  def DoNE(self, b): return self.doRelop(b, '!=')
  def DoLT(self, b): return self.doRelop(b, '<')
  def DoLE(self, b): return self.doRelop(b, '<=')
  def DoGT(self, b): return self.doRelop(b, '>')
  def DoGE(self, b): return self.doRelop(b, '>=')

class Ybyt(Ybase):
  def __init__(self, y, s):
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkByt(%s)' % str(self.y)
    return str(self.s)
  def DoByt(self):
    return str(self.y)
  def DoStr(self):
    return '/*Ybyt.DoStr*/string(%s)' % self.y
  def DoBool(self):
    return '/*Ybyt.DoBool*/(%s != "")' % self.y

  def doRelop(self, b, op):
    if type(b) is Ystr:
      return Ybool('(/*YYbyt.doRelop*/ string(%s) %s string(%s) )' % (self.y, op, b.y), None)
    if type(b) is Ybyt:
      return Ybool('(/*YYbyt.doRelop*/ string(%s) %s string(%s) )' % (self.y, op, b.y), None)
    return ''

  def DoEQ(self, b): return self.doRelop(b, '==')
  def DoNE(self, b): return self.doRelop(b, '!=')
  def DoLT(self, b): return self.doRelop(b, '<')
  def DoLE(self, b): return self.doRelop(b, '<=')
  def DoGT(self, b): return self.doRelop(b, '>')
  def DoGE(self, b): return self.doRelop(b, '>=')

class Z(object):  # Returns from visits (emulated runtime value).
  def __init__(self, t, s):
    self.t = t  # T node
    self.s = s  # String for backwards compat
  def __str__(self):
    return self.s
  def DoBool(self): return ''
  def DoInt(self): return ''
  def DoFloat(self): return ''
  def DoByt(self): return ''
  def DoStr(self): return ''
  def DoAdd(self, b): return ''
  def DoSub(self, b): return ''
  def DoMul(self, b): return ''
  def DoDiv(self, b): return ''
  #def DoIDiv(self, b): return ''
  def DoMod(self, b): return ''
  def DoEQ(self, b): return ''
  def DoNE(self, b): return ''
  def DoLT(self, b): return ''
  def DoLE(self, b): return ''
  def DoGT(self, b): return ''
  def DoGE(self, b): return ''

class Zself(Z):
  def __str__(self):
    return 'MkX(&self.PBase)'

class Zsuper(Z):
  pass
class Zlocal(Z):
  pass
class Zglobal(Z):
  pass

class Zimport(Z):
  def __init__(self, t, s, imp):
    if rye_rye:
      super(t, s)
    else:
      Z.__init__(self, t, s)
    self.imp = imp  # imports[] object
  pass

class Zspecial(Z):
  pass
class Zlit(Z):
  pass

pass

class ArgDesc(object):
  def __init__(self, module, cls, name, args, dflts, star, starstar, isCtor):
    self.module = module
    self.cls = cls
    self.name = name
    self.args = args
    self.dflts = dflts
    self.star = star
    self.starstar = starstar
    self.isCtor = isCtor

  def NotYetUsed_CallSpec(self):
    argnames = ', '.join(['"%s"' % a for a in self.args])
    defaults = ', '.join([(str(d.visit(self)) if d else 'MissingM') for d in self.dflts])
    return 'PCallable{Name: "%s", Args: []string{%s}, Defaults: []M{%s}, Star: "%s", StarStar: "%s"}' % (self.name, argnames, defaults, self.star, self.starstar)

def AOrSkid(s):
  if s:
    return 'a_%s' % s
  else:
    return '_'

