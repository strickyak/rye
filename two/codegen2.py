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

INTLIKE_GO_TYPES = {'int', 'int8', 'int16', 'int32', 'int64', 'uint', 'uint8', 'uint16', 'uint32', 'uint64', 'uintptr'}
FLOATLIKE_GO_TYPES = {'float32', 'float64'}

RYE_SPECIALS = {
    'go_cast', 'go_type', 'go_indirect', 'go_addr', 'go_elem', 'go_new', 'go_make', 'go_append',
    'len', 'str', 'repr', 'int', 'float',
    }

NoTyps = None
NoTyp = None
None2 = 'NoneJ_1,NoneJ_2'

SMALLER = not os.getenv('RYE_BLOAT')
RyeEnv = os.getenv('RYE')
DebugCall = 'C' in RyeEnv if RyeEnv else False

NONALFA = re.compile('[^A-Za-z0-9]')
TROUBLE_CHAR = re.compile('[^]-~ !#-Z[]')
def GoStringLiteral(s):
  if rye_rye:
    return strconv.QuoteToASCII(s)
  else:
    return '"' + TROUBLE_CHAR.sub((lambda m: '\\x%02x' % ord(m.group(0))), s) + '"'


STARTS_WITH_V = re.compile('^ *([/][*].*[*][/] *)v').match

def CleanIdentWithSkids(s):
  if len(s) < 50:
    # Work around lack of callable() for .sub in RYE.
    return md5.new(s).hexdigest()
    # TODO = NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
    # return NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)
  else:
    return md5.new(s).hexdigest()

def Doublet(x):
  return x, x

def W(s):
  return 'W{%s}' % s
def BifurcateUV(s):
  s = str(s)
  if ',' in s:
    return s
  else:
    return '%s_1 U, %s_2 V' % (s, s)
def BifurcateUVSemi(s):
  s = str(s)
  return '%s_1 U; %s_2 V' % (s, s)
def Bifurcate(s):
  s = str(s)
  if ',' in s:
    return s
  else:
    return '%s_1, %s_2' % (s, s)
def Wifurcate(s):
  return W(Bifurcate(s))

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

def InsertRye__(s):
  a = s.split('/')
  b = a[:-1] + ['rye__'] + a[-1:]
  return '/'.join(b)

def TypName(t):
  if type(t) is parse.Tvar:
    return str(t.name)
  elif t:
    return repr(t)

SerialNum = 100
def Serial(s):
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (str(s), SerialNum)

def Pair(x):
  if isinstance (x, Ybase):
    return x  # Y's are lazy; don't force them.
  if type(x) is str:
    serial = Serial('uu')
    print '%s_1, %s_2 := %s' % (serial, serial, x)
    print ' _,_ = %s_1,%s_2' % (serial, serial)
    return ' %s_1,%s_2 ' % (serial, serial)

class CodeGen(object):
  def __init__(self):
    self.glbls = {}         # name -> (type, initialValue)
    self.imports = {}       # name -> Vimport
    self.ydefs = {}          # name -> yfunc
    self.ymeths = DeepDict() # clsname -> methname -> Yfunc
    self.lits = {}          # key -> name
    self.invokes = {}       # key -> (n, fieldname)
    self.scope = None       # None means we are at module level.
    self.tail = []
    self.cls = None
    self.getNeeded = {}      # keys are getter names.
    self.setNeeded = {}      # keys are setter names.
    self.signatures = {}     # keys are SIGNATURE interface names.
    self.maxNumCallArgs = -1


  def Serial(self, s):
    global SerialNum
    SerialNum += 1
    return '%s_%d' % (str(s), SerialNum)

  def Pair(self, x):
    if isinstance (x, Ybase):
      return x  # Y's are lazy; don't force them.
    if type(x) is str:
      serial = self.Serial('uu')
      print '%s_1, %s_2 := %s' % (serial, serial, x)
      return ' %s_1,%s_2 ' % (serial, serial)

  def InjectForInternal(self, stuff):
    self.invokes, self.ydefs, self.getNeeded, self.setNeeded = stuff

  def ExtractForInternal(self):
    stuff = self.invokes, self.ydefs, self.getNeeded, self.setNeeded
    return stuff

  def GenModule(self, modname, path, tree, cwp=None, internal=""):
    buf = PushPrint()
    try:
      z = self.GenModule2(modname, path, tree, cwp, internal)
    finally:
      PopPrint()
      print str(buf)
    return z

  def GenModule2(self, modname, path, tree, cwp=None, internal=""):
    self.cwp = cwp
    self.path = path
    self.thispkg = InsertRye__(path)
    self.modname = modname
    if internal:
      self.internal = open('gen_internals.py', 'w')
    else:
      self.internal = None
      self.glbls['__name__'] = ('ZZZ', 'macro.MkStrJX("%s")' % modname)

    self.func_level = 0
    self.func = None
    self.yields = None
    self.force_globals = {}

    print '// +build prego'
    print ''
    if internal:
      print ' package two'
      print ' import . "github.com/strickyak/rye"'
    else:
      print ' package %s' % modname.split('/')[-1]
      print ' import . "github.com/strickyak/rye"'
      print ' import . "github.com/strickyak/rye/two"'

    if True:
      print ''
      print '// cwp:', self.cwp
      print '// path:', self.path
      print '// thispkg:', self.thispkg
      print '// modname:', self.modname
      print '// internal:', self.internal
      print ''

    print ' import "fmt"'
    print ' import "io"'
    print ' import "os"'
    print ' import "reflect"'
    print ' import "runtime"'
    print ' import "time"'  # For FAV packages.
    print ' import "bytes"'  # For FAV packages.
    print ' import "unsafe"'

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
        imp.pkg = pkg
        if imp.alias == '_':
          print ' import _ "%s" // %s' % (pkg, repr(imp))  # was vars()
        else:
          alias = 'i_%s' % imp.alias
          print ' import %s "%s" // %s' % (alias, pkg, repr(imp))  # was vars()
          if rye_rye:
            pass
          else:
            print '// DIR %s // VARS %s' % (dir(imp), vars(imp))

          if samples.SAMPLES.get(pkg):
            to_be_sampled[alias] = pkg

          if not self.internal:
            self.glbls[imp.alias] = ('*PModule', None2)

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
              '   z := new(J_%s)' % th.name,
              '   z.Self = z',
              '   z.Rye_ClearFields__()',
              ]
        c1 = parse.Tnative(natives)

        c2 = parse.Tassign(parse.Tvar('rye_result__'), parse.Traw('macro.MkPJ(&z.JBase)'))

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
    print ' var _ = unsafe.Sizeof(0)'
    print ' var _ = MkInt'  # From rye runtime.
    print ''

    # BEGIN: JEval_Module, innter_eval_module
    if not self.internal:
      print ' var jeval_module_once bool'
      print ' func JEval_Module () (U, V) {'
      print '   if jeval_module_once == false {'
      print '     jeval_module_once = true'
      print '     _,_ = jinner_eval_module()'
      print '   }'
      print '   return NoneJ_1, NoneJ_2'
      print ' }'

    print ' func jinner_eval_module () (U,V) {'

    # Look ahead at global functions and classes (and their methods).
    # Build the preVdef for them, so they can be called by forward reference.
    # Populates ydefs & ymeths with Yfuncs.
    for p in tree.things:
      if type(p) is parse.Tdef:
        self.preVdef(p)
      elif type(p) is parse.Tclass:
        for q in p.things:
          if type(q) is parse.Tdef:
            self.cls = p
            self.preVdef(q)
            self.cls = None

    # ALL THINGS IN MODULE.
    for th in tree.things:
      self.Gloss(th)
      th.visit(self)
      self.Ungloss(th)

    # END: JEval_Module, innter_eval_module
    print '   return NoneJ_1, NoneJ_2'
    print ' }'
    print ''
    print '//(begin tail)'
    print '\n//(tail)\n'.join(self.tail)
    print '//(end tail)'
    print ''

    # H_rye_rye is defined in runtime.go, so don't put it in other modules.
    if 'rye_rye' in self.glbls:
      del self.glbls['rye_rye']

    for g, (t, v) in sorted(self.glbls.items()):
      print 'var H_%s_1 U; var H_%s_2 V // %s' % (g, g, t)
    print ''
    print ' func init /*New_Module*/ () {'
    for g, (t, v) in sorted(self.glbls.items()):
      print '   H_%s_1, H_%s_2 = %s // t=%s' % (g, g, v, t)
    if internal:
      print '   jinner_eval_module()'
    print ' }'
    print ''

    for key, code in sorted(self.lits.items()):
      print 'var %s_1 U; var %s_2 V; func init() { %s_1, %s_2 = %s }' % (key, key, key, key, code)
    print ''

    if self.internal:
        print >> self.internal, '# This file is generated when builtins.py is compiled.'
        print >> self.internal, 'InternalInvokers = ['
    for _, (n, fieldname) in sorted(self.invokes.items()):
      self.getNeeded[fieldname] = True
      formals = ', '.join(['a_%d_1 U, a_%d_2 V' % Doublet(i) for i in range(n)])
      args = ', '.join(['a_%d_1, a_%d_2' % Doublet(i) for i in range(n)])

      if self.internal or (n, fieldname) not in gen_internals.InternalInvokers:
        letterF = 'JF' if self.internal else 'jf'
        letterI = 'JI' if self.internal else 'ji'
        letterGet = 'JI' if self.internal or fieldname in gen_internals.InternalGetters else 'ji'
        print 'func %s_INVOKE_%d_%s(fn_1 U, fn_2 V, %s) (U,V) {' % (letterF, n, fieldname, formals)
        print '  fn := inline.CheckTakePJ(fn_1, fn_2)'
        print '  switch x := fn.(type) {   '
        print '  case %s_INVOKE_%d_%s:         ' % (letterI, n, fieldname)
        print '    return x.N_%d_%s(%s)         ' % (n, fieldname, args)
        print '  case %s_GET_%s:         ' % (letterGet, fieldname)
        print '    tu, tv := x.GET_%s()    ' % fieldname
        print '    return %s_%d(tu, tv, %s)' % (('JCALL' if n<11 else 'jcall'), n, ', '.join([Bifurcate('a_%d' % j) for j in range(n)]))
        print ''
        print '  case *JGo:                '
        print '    return x.Invoke("%s", %s) ' % (fieldname, ', '.join(['W{a_%d_1, a_%d_2}' % (j,j) for j in range(n)]))
        print '  }'
        print '  panic(fmt.Sprintf("Cannot invoke \'%s\' with %d arguments on %%v", fn))' % (fieldname, n)
        print '}'
        print 'type %s_INVOKE_%d_%s interface { N_%d_%s(%s) (U,V) }' % (letterI, n, fieldname, n, fieldname, formals)
        if self.internal:
          print >> self.internal, '  (%d, "%s"),' % (n, fieldname)
    print ''

    if self.internal:
      print >> self.internal, '  ]'
      print >> self.internal, '# FooBarBaz'
      print >> self.internal, 'InternalGetters = ['

    for iv in sorted(self.getNeeded):
      if self.internal or iv not in gen_internals.InternalGetters:
        letterF = 'JF' if self.internal else 'jf'
        letterI = 'JI' if self.internal else 'ji'
        print 'type %s_GET_%s interface { GET_%s() (U,V) }' % (letterI, iv, iv)
        print 'func %s_GET_%s(h_1 U, h_2 V) (U,V) {' % (letterF, iv)
        print '  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }'
        print '  h := inline.TakePJ(h_1, h_2)'
        print '  switch x := h.(type) { '
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
        letterF = 'JF' if self.internal else 'jf'
        letterI = 'JI' if self.internal else 'ji'
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
        whichI = 'JI' if i<11 else 'ji'
        whichCall = 'JCALL' if i<11 else 'jcall'
        print '  type %s_%d interface { JCall%d(%s) (U,V) }' % (whichI, i, i, ", ".join(i * [' U,V']))
        print '  func %s_%d (fn_1 U, fn_2 V, %s) (U,V) {' % (whichCall, i, ', '.join(['a_%d_1 U, a_%d_2 V' % Doublet(j) for j in range(i)]))
        print '    //// if fn.X == nil { panic("cannot JCALL on int or str") }'
        print '    fn := macro.TakePJ(fn_1, fn_2)'
        print '    switch f := fn.(type) {'
        print '      case %s_%d:' % (whichI, i)
        print '        return f.JCall%d(%s)' % (i, ', '.join(['a_%d_1, a_%d_2' % Doublet(j) for j in range(i)]))
        print '      case IJCallV:'
        print '        return f.JCallV([]W{%s}, nil, nil, nil)' % ', '.join(['W{a_%d_1, a_%d_2}' % Doublet(j) for j in range(i)])
        print '    }'
        print '    panic(fmt.Sprintf("No way to call: %v", fn))'
        print '  }'
        print ''

    print '// self.signatures.items:', self.signatures
    for name, sig in sorted(self.signatures.items()):
      print 'type %s interface { %s } // self.signatures.items' % (name, sig)

    if self.internal:
      self.internal.close()

    # BEGIN Just a comment {
    #def describe(t):
    #  if type(t) is parse.Tvar:
    #    return t.name
    #  elif type(t) is parse.Traw:
    #    return t.raw
    #  else:
    #    return '-:' + str(type(t)) + ':' + repr(t)

    if True:
      print ''
      for yk, yd in sorted(self.ydefs.items()):
        print '//ydefs// %s => %s [[ %s ]]' % (yk, yd, repr(yd))  # was vars()
        #//for ydk, ydv in sorted(vars(yd).items()):
        #//  print '//ydefs// ... ... ... %s :: %s' % (ydk, ydv)
        print '//'
      print '//'

      for yck, ycd in sorted(self.ymeths.dd.items()):
        print '//ymeth// %s => [ %s ]' % (yck, sorted(ycd.keys()))
        for ymk, ymd in sorted(ycd.items()):
          print '//ydefs// %s => %s => %s [[ %s ]]' % (yck, ymk, ymd, repr(ymd))  # was vars()
        print '//'
      print '//'
      print ''

    #if False:
    #  for th in tree.things:
    #    if type(th) == parse.Tdef:
    #      # name, args, typs, rettyp, dflts, star, starstar, body.
    #      jtyps = [describe(t) for t in th.typs] if th.typs else None
    #      j = dict(what='func', name=th.name, star=th.star, starstar=th.starstar, typs=jtyps)
    #      print '//@@// %s %s' % (th.name, j)
    #    elif type(th) == parse.Tclass:
    #      for m in th.things:
    #        if type(m) == parse.Tdef:
    #          jtyps = [describe(t) for t in m.typs] if m.typs else None
    #          j = dict(what='meth', cls=th.name, name=m.name, star=m.star, starstar=m.starstar, typs=jtyps)
    #          print '//@@// %s.%s %s' % (th.name, m.name, j)
    pass
    # END Just a comment }

  def Gloss(self, th):
    print '// @ %d @ %d @ %s' % (th.where, th.line, self.CurrentFuncNick())

  def Ungloss(self, th):
    print '// $ %d $ %d $' % (th.where, th.line)

  def AssignFieldAFromRhs(self, a, rhs):
      lhs = a.p.visit(self)
      if type(lhs) is Yself:  # Special optimization for self.
        self.instvars[a.field] = True
        lhs = 'self.M_%s' % a.field
        print '   %s = %s' % (Bifurcate(lhs), rhs)
      elif type(lhs) is Yimport:  # For module variables.
        if lhs.imp.imported[0] == 'go':
          print '  reflect.ValueOf(& %s.%s).Elem().Set( reflect.ValueOf(%s.Contents()).Convert(reflect.TypeOf(%s.%s)))' % (
              lhs, a.field, rhs, lhs, a.field)
        else:
          target = '%s.H_%s' % (lhs, a.field)
          print '   %s = %s' % (Bifurcate(target), rhs)
      else:
        self.setNeeded[a.field] = True
        letterF = 'F' if self.internal or a.field in gen_internals.InternalSetters else 'f'
        print '   %s_SET_%s(%s, %s)' % (letterF, a.field, lhs, rhs)

  def AssignItemAFromRhs(self, a, rhs):
        p = Pair(a.a.visit(self))
        q = Pair(a.x.visit(self))
        print '   FnSetItemJ(inline.CheckTakePJ(%s), %s, %s)' % (p, q, rhs)

  def AssignTupleAFromB(self, a, b):
        serial = self.Serial('detuple')
        tmp = parse.Tvar(serial)
        parse.Tassign(tmp, b).visit(self)

        print '   len_%s := FnLenJ(%s)' % (serial, Pair(tmp.visit(self)))
        print '   if len_%s != %d { panic(fmt.Sprintf("Assigning object of length %%d to %%d variables, in destructuring assignment.", len_%s, %d)) }' % (serial, len(a.xx), serial, len(a.xx))

        i = 0
        for x in a.xx:
          if type(x) is parse.Tvar and x.name == '_':
            pass # Tvar named '_' is the bit bucket;  don't Tassign.
          else:
            parse.Tassign(x, parse.Tgetitem(tmp, parse.Tlit('N', i))).visit(self)
          i += 1

  def Assign(self, a, b):
    # Resolve rhs first.
    type_a = type(a)

    # p, q, ... = rhs
    if type_a is parse.Titems or type_a is parse.Ttuple:
      return self.AssignTupleAFromB(a, b)

    lhs = '?lhs?'
    rhs = b.visit(self)

    if type_a is parse.Tfield:      # p.q = rhs
      return self.AssignFieldAFromRhs(a, rhs)

    elif type_a is parse.Tgetitem:  # p[q] = rhs
      return self.AssignItemAFromRhs(a, rhs)

    elif type_a is parse.Tvar:
      # Are we in a function scope?
      if self.scope is not None and a.name not in self.force_globals:
        # Inside a function.
        if self.scope.get(a.name):
          lhs = self.scope[a.name]
        else:
          lhs = Yvar(False, 'v_%s' % a.name)
          lhs.flavor = 'L'
          print '// @@@@@@ Creating var "%s" in scope @@@@@@ -> Yvar: %s' % (a.name, repr(lhs))
          self.scope[a.name] = lhs
      else:
        # At the module level.
        lhs = a.visit(self)
        self.glbls[a.name] = ('U', 'NoneJ_1, NoneJ_2')

    elif type_a is parse.Traw:
      if a.raw != '_':
        raise Exception("When is Traw used? %s" % a.raw)
      lhs = a.raw
      # TODO -- fix skipping _ = Ylit
      if isinstance(rhs, YbaseTyped) and rhs.flavor == 'L':
        return  # Assigning literals to _ is pointless.

    else:
      raise Exception('Weird Assignment, a class is %s, Left is (%s) (%s) Right is (%s) (%s)' % (type(a).__name__, a, a.visit(self), b, b.visit(self)))

    # Assign the variable, unless it is the magic un-assignable rye_rye.
    if isinstance(lhs, Ybase):
      print '// Calling DoAssign because isinstance(lhs %s, Ybase)' % type(lhs)
      ok = DoAssign(lhs, rhs)
      if ok:
        return

    if str(lhs) == 'H_rye_rye':
      print '// Assignment of rye_rye suppressed in rye.'
    elif str(lhs) == '_':
      if type(rhs) == parse.Tlit:
        print '   _,_ = %s // 1. Assign void = Tlit' % rhs
      else:
        print '   _,_ = %s // 2. Assign void: = type: %s repr: %s ' % (rhs, type(rhs), repr(rhs))
    else:
      print '   %s = %s  // 3. Assign %s lhs %s = rhs %s' % (Bifurcate(lhs), rhs, type_a, type(lhs), type(rhs))

  def Vassign(self, p):
    # a, b
    self.Assign(p.a, p.b)

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
          print '   fmt.Fprintln(%s, %s)' % (
              'M(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()',
              ', '.join([ForceString(v).AsStr() for v in vv]))
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
      print '   H_%s = i_%s.JEval_Module() ' % (p.alias, p.alias)

  def Vassert(self, p):
    where = '%s:%s %s.%s' % (
        self.modname, str(p.line),
        self.cls.name if self.cls else '',
        self.func.name if self.func else '',
        )
    print '// Vassert: p.x=', p.x
    print '// Vassert: p.y=', p.y
    print '// Vassert: p.fails=', p.fails
    print '// Vassert: type(p.x)=', type(p.x)
    if type(p.x) == parse.Top:  print '// Vassert: p.x.a=', p.x.a
    if type(p.x) == parse.Top:  print '// Vassert: p.x.op=', p.x.op
    if type(p.x) == parse.Top:  print '// Vassert: p.x.b=', p.x.b

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
''' % GoStringLiteral(p.code)
      print '_ = %s' % p.x.visit(self)
      print '''
        }()
      '''
      # TODO:  Check regexp of exception.


    elif p.y is None and type(p.x) == parse.Top and p.x.op in parse.REL_OPS.values():
      # Since message is empty, print LHS, REL_OP, and RHS, since we can.
      a = p.x.a.visit(self)
      b = p.x.b.visit(self)
      sa = self.Serial('left')
      sb = self.Serial('right')
      print '   %s, %s := %s, %s' % (Bifurcate(sa), Bifurcate(sb), a, b)
      print '   if ! (%s.%s(%s)) {' % (sa, p.x.op, sb)
      print '     panic(fmt.Sprintf("Assertion Failed [%s]:  (%%s) ;  left: (%%s) ;  op: %%s ;  right: (%%s) ", %s, %s.Repr(), "%s", %s.Repr() ))' % (
          where, GoStringLiteral(p.code), sa, p.x.op, sb, )
      print '   }'
    else:
      print '   if ! (%s) {' % AsBool(p.x.visit(self))
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
         %s_try_1, %s_try_2 := func() (%s_z_1, %s_z_2 U) {
           defer func() {
             r := recover()
             if r != nil {
               PrintStackFYIUnlessEOFBecauseExcept(r)
               %s_z_1, %s_z_2 = func() (U,V) {
               // BEGIN EXCEPT
      ''' % (serial, serial, serial, serial, serial, serial, serial)
      # Assign, for the side effect of var creation.
      if p.exvar:
        parse.Tassign(p.exvar, parse.Traw('MkRecovered(r)')).visit(self)

      p.ex.visit(self)

      print '''
                 return 0,V0
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
           return 0,V0
         }()
         if %s_try != 0,V0 { return %s_try }
         // END OUTER EXCEPT %s
      ''' % (serial, serial, serial)

    if p.fin:
      print '    return 0,V0'
      print '  }()'
      print '  if fin_ret_%s != 0,V0 { return fin_ret_%s }' % (serial, serial)
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
     var zz%s []W
     var Nextee = macro.TakePJ(%s)
     var nexter%s JNexter = nextee.Iter()
     enougher%s, canEnough%s := nexter%s.(Enougher)
     if canEnough%s {
             defer enougher%s.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx_%s_1, ndx_%s_2, more_%s := nexter%s.NextJ()
       if !more_%s {
         break
       }
       // BEGIN FOR EXPR
''' % (i, i, Bifurcate(ptv), i, i, i, i, i, i, i, i, i, i, i)

    parse.Tassign(p.vv, parse.Traw(Bifurcate("ndx_%s" % i))).visit(self)

    if p.cond:
      print '  if %s {' % AsBool(p.cond.visit(self))

    print '  zz%s = append(zz%s, %s)' % (i, i, p.z.visit(self))

    if p.cond:
      print '  }'

    print '''
       // END FOR EXPR
     }
     return MkListJ(zz%s)
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
    n = AsInt(a0.visit(self))  # General case.
    #### if type(a0) == parse.Tlit and a0.k == 'I':
    ####   n = a0.v  # Optimized literal int case.

    bt = Bifurcate('tmp_%s' % i)
    print '''
      var i_%s int64
      var n_%s int64 = %s
      for i_%s = int64(0); i_%s < n_%s; i_%s++ {
        %s := MkIntJ(i_%s)
''' % (i, i, n, i, i, i, i, bt, i)
    parse.Tassign(var, parse.Traw(bt)).visit(self)
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
    print '// Vfor: var=%s i=%s ptv=%s' % (p.var, i, ptv)
    print '''
   for_returning%s_1, for_returning%s_2 := func () (U,V) { // around FOR
     var nextee PJ = macro.TakePJ(%s)
     var nexter%s JNexter = nextee.Iter()
     enougher%s, canEnough%s := nexter%s.(Enougher)
     if canEnough%s {
             defer enougher%s.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx_%s_1, ndx_%s_2, more_%s := nexter%s.NextJ()
       if !more_%s {
         break
       }
       // BEGIN FOR
''' % (i, i, Bifurcate(ptv), i, i, i, i, i, i, i, i, i, i, i)

    parse.Tassign(p.var, parse.Traw(Bifurcate("ndx_%s" % i))).visit(self)

    p.b.visit(self)

    print '''
       // END FOR
     }
     return 0,V0
   }() // around FOR
   if for_returning%s_1 != 0 && for_returning%s_2 != V0 { return for_returning%s_1, for_returning%s_2 }
''' % (i, i, i, i)

  # New "with defer".    (call, body)
  def Vwithdefer(self, p):
    # call, body
    var = self.Serial('with_defer_returning')
    immanentized = self.ImmanentizeCall(p.call, 'defer')
    print '  %s := func() (U,V) { defer %s' % (var, immanentized.visit(self))
    p.body.visit(self)
    print '    return 0,V0'
    print '  }()'
    print '  if %s != 0,V0 { return %s }' % (var, var)

  def Vglobal(self, p):
    pass

  def Vswitch(self, p):
    # (self, a, cases, clauses, default_clause):
    serial = self.Serial('sw')
    self.Gloss(p)
    if p.a:
      print '   var %s U' % Bifurcate(serial)
      print '   %s := %s' % (Bifurcate(serial), p.a.visit(self))
      print '   _,_ = %s' % Bifurcate(serial)
    self.Ungloss(p)
    print '   switch true {'
    for ca, cl in zip(p.cases, p.clauses):
      self.Gloss(ca)
      if p.a:
        print '      case %s.EQ(%s): {' % (serial, ca.visit(self))
      else:
        print '      case %s: {' % AsBool(ca.visit(self))
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
      self.Assign(p.varlist, parse.Traw('if_tmp'))
    else:
      print '   if %s {' % AsBool(p.t.visit(self))
    p.yes.visit(self)
    if p.no:
      print '   } else {'
      p.no.visit(self)
    print '   }'

  def Vwhile(self, p):
    # NB don't print the predicate on the 'for' line,
    # or else extra code generated will go before the 'for'.
    print '   for {'
    print '     if !(%s) { break }' % AsBool(p.t.visit(self))
    p.yes.visit(self)
    print '   }'

  def Vreturn(self, p):
    if p.aa is None:
      print '   return None '
    else:
      vv = [a.visit(self) for a in p.aa]
      if len(vv) == 1:
        #print >>sys.stderr, '   //return type=%s ' % type(vv[0])
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
    print '   panic( ContentsJ(%s) )' % p.a.visit(self)

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
      print '// LitIntern: interning %s: %s = %s' % (v, key, code)
      self.lits[key] = code
    ret = Ylit(v, '%s_1,%s_2' % (key, key))
    ret.flavor = 'L'
    return ret

  def Vraw(self, p):
    if p.raw == 'False':
      return Ybool('false', 'FalseJ_1, FalseJ_2')
    if p.raw == 'True':
      return Ybool('true', 'TrueJ_1, TrueJ_2')
    if p.raw == 'None':
      return 'NoneJ_1, NoneJ_2'
    return p.raw

  def Vlit(self, p):
    ret = None
    if p.k == 'N':
      z = p.v
      key = 'jlitI_' + CleanIdentWithSkids(str(z))
      code = 'MkIntJ(%s)' % str(z)
      lit = self.LitIntern(z, key, code)
      ret = Yint(str(z), lit)
    elif p.k == 'F':
      z = p.v
      key = 'jlitF_' + CleanIdentWithSkids(str(z))
      code = 'MkFloatJ(%s)' % str(z)
      lit = self.LitIntern(z, key, code)
      ret = Yfloat(str(z), lit)
    elif p.k == 'S':
      z = parse.DecodeStringLit(p.v)
      key = 'jlitS_' + CleanIdentWithSkids(z)
      golit = GoStringLiteral(z)
      code = 'MkStrJ( %s )' % golit
      lit = self.LitIntern(golit, key, code)
      ret = Ystr(golit, lit)
    else:
      raise Exception('Unknown Vlit', p.k, p.v)
    ret.flavor = 'L'
    return ret

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

      return Ybool('/*Vop returns bool*/Fn%sJ(%s,%s)' % (p.op, Pair(p.a.visit(self)), Pair(p.b.visit(self))), None)
    if p.b:

      # Optimizations.
      if p.op == 'Add':
        #return DoAdd(p.a.visit(self), p.b.visit(self))
        #print '/* ...Add...', '*/'
        #print '/* ...Add...', type(p.a), '*/'
        pav = p.a.visit(self)
        #print '/* ...Add...', '*/'
        #print '/* ...Add...', type(p.b), '*/'
        pbv = p.b.visit(self)
        #print '/* ...Add...', '*/'
        #print '/* ...Add...', '*/'
        #print '/* op == Add:  pav=%s */' % pav
        #print '/* op == Add:    pav=%s */' % pbv
        return DoAdd(pav, pbv)
      if p.op == 'Sub':
        return DoSub(p.a.visit(self), p.b.visit(self))
      if p.op == 'Mul':
        return DoMul(p.a.visit(self), p.b.visit(self))
      if p.op == 'Div':
        return DoDiv(p.a.visit(self), p.b.visit(self))
      if p.op == 'Mod':
        return DoMod(p.a.visit(self), p.b.visit(self))
      return ' Fn%sJ(%s, %s) ' % (p.op, p.a.visit(self), p.b.visit(self))
    else:
      return ' Fn%sJ(%s) ' % (p.op, p.a.visit(self))

  def Record(self, v):
    pass
    #if self.recording:
    #  print 'if Recording != nil { fmt.Fprintf(Recording, "{\t%s\t%s\t%%s\t}\\n", M(%s).PType().String()) }' % (self.modname, v, v)

  def RecordOp(self, c, a, b, op):
    pass
    #if self.recording:
    #  print 'if Recording != nil { fmt.Fprintf(Recording, "{\t%s\t%s\t%%s\t%%s\t%%s\t%s}\\n", M(%s).PType().String(),M(%s).PType().String(),  M(%s).PType().String(), ) }' % (self.modname, c, op, c, a, b, )

  def Vboolop(self, p):
    if p.op == '!':
      return Ybool('(/*Vboolop*/ !(%s)) ' % AsBool(p.a.visit(self)), None)

    elif p.op == '&&':
      s = self.Serial('andand')
      bs = Bifurcate(s)
      # NB these must be on three different lines, or extra generated lines go the wrong place.
      print '  %s := %s' % (bs, Pair(p.a.visit(self)))
      print '  if BoolJ(%s) {' % bs
      print '    %s = %s' % (bs, Pair(p.b.visit(self)))
      print '  }'
      return bs

    elif p.op == '||':
      s = self.Serial('oror')
      # NB these must be on three different lines, or extra generated lines go the wrong place.
      print '  %s := %s' % (bs, Pair(p.a.visit(self)))
      print '  if ! BoolJ(%s) {' % bs
      print '    %s = %s' % (bs, Pair(p.b.visit(self)))
      print '  }'
      return bs

    else:
      raise Exception('notreached(Vboolop)')

  def Vcondop(self, p):  # b if a else c
    fn = self.Serial('cond')
    print '%s := func (a bool) (U,V) { if a {' % fn
    print 'return %s' % p.b.visit(self)
    print '}'
    print 'return %s' % p.c.visit(self)
    print '}'
    return ' %s(%s) ' % (fn, AsBool(p.a.visit(self)))

  def Vgetitem(self, p):
    return ' FnGetItemJ(%s, %s) ' % (Pair(p.a.visit(self)), Pair(p.x.visit(self)))

  def Vgetitemslice(self, p):
    # No more third index; we can never use it.
    return Pair(' FnGetItemSliceJ(%s, %s, %s) ' % (
        p.a.visit(self),
        None2 if p.x is None else Pair(p.x.visit(self)),
        None2 if p.y is None else Pair(p.y.visit(self))))

  def Vtuple(self, p):
    return Pair('MkTupleVUV( %s )' % ', '.join(['W{%s}' % Pair(str(x.visit(self))) for x in p.xx]))

  def Vlist(self, p):
    return Pair('MkListVUV( %s )' % ', '.join(['W{%s}' % Pair(str(x.visit(self))) for x in p.xx]))

  def Vdict(self, p):
    return Pair('MkDictVJ( %s )' % ', '.join(['W{%s}' % Pair(str(x.visit(self))) for x in p.xx]))

  def Vset(self, p):
    return Pair('MkSetVJ( %s )' % ', '.join(['W{%s}' % Pair(str(x.visit(self))) for x in p.xx]))

  def Vvar(self, p):
    if p.name == 'self':
      return Yself(p, 'self')
    if p.name == 'super':
      return Ysuper(p, 'super')
    if p.name in self.force_globals:
      z = Yvar(p, '/*force_globals*/' + 'H_%s' % p.name)
      z.flavor = 'G'  # Global.
      return z
    if p.name in self.imports:
      return Yimport('i_%s' % p.name, self.imports[p.name])
    if self.scope and p.name in self.scope:
      x = self.scope[p.name]
      print '// Vvar: local var %s -> %s: %s' % (p.name, type(x).__name__, repr(x))

      # YAK Start replacing local Yint vars with Yint:
      if isinstance(x, Yint):
        x.flavor = 'L'
        return x
      if isinstance(x, Ystr):
        x.flavor = 'L'
        return x
      if isinstance(x, Ybool):
        x.flavor = 'L'
        return x
      if isinstance(x, Yfloat):
        x.flavor = 'L'
        return x

      z = Yvar(p, str(x))
      z.flavor = 'L'  # Local.
      return z
    if p.name in RYE_SPECIALS:
      return Yspecial(p, Bifurcate('H_%s' % p.name))
    z = Yvar(p, Bifurcate('H_%s' % p.name))
    z.flavor = 'G'  # Global.
    return z

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

  def OptimizedGoCall(self, ispec, args, qfunc, maybeResult):
    if len(args) != len(qfunc.takes):
      raise Exception('wrong num args for shortcut METH: %s GOT %s WANT %s' % (qfunc.name, args, qfunc.takes))

    print( '// ATTEMPT OptimizedGoCall: %s%s TAKES %s RETURNS %s' % (ispec, qfunc.Invoklet(), qfunc.takes, qfunc.rets) )
    buf = PushPrint()
    try:

      s = self.Serial('opt_go_call')
      ins = []
      for i in range(len(qfunc.takes)):
        t = qfunc.takes[i]
        if t == 'string':
          v = '%s(%s)' % (t, AsStr(args[i]))
        elif t == '[]string':
          v = 'ListToStrings(%s.List())' % args[i]
        elif t == '[]uint8':
          v = '%s(%s)' % (t, AsByt(args[i]))
        elif t == 'bool':
          v = '%s(%s)' % (t, AsBool(args[i]))
        elif t in INTLIKE_GO_TYPES:
          v = '%s(%s)' % (t, AsInt(args[i]))
        elif t in FLOATLIKE_GO_TYPES:
          v = '%s(%s)' % (t, AsFloat(args[i]))
        else:
          raise Exception("OptimizedGoCall: Not supported yet: takes: %s" % t)

        var = '%s_t_%d' % (s, i)
        print( '/*OptimizedGoCall*/ var %s %s = %s' % (var, t, v) )
        ins.append(var)

      # Handle final magic error return.
      rets = qfunc.rets
      magic_error = ''
      if rets and rets[-1] == 'error':
        magic_error = rets[-1]

      outs = ['%s_r_%d' % (s, i) for i in range(len(qfunc.rets))]

      types = []
      results = []
      if rets:
        for i in range(len(rets)):
          r = rets[i]
          typ = r
          if r == 'bool':
            v = Ybool(outs[i], None)
          elif r == 'error':
            v = outs[i]
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
            typ = 'interface{}'
            v = 'AdaptForReturn(reflect.ValueOf(%s))' % outs[i]
            #raise Exception("OptimizedGoCall: Not supported yet: returns: %s" % r)

          types.append(typ)
          results.append(v)

      if outs:
        maybeValue = '\n'.join(['var %s_r_%d %s /*result*/' % (s, i, types[i]) for i in range(len(qfunc.rets))])
        assigns = '%s =' % ', '.join(outs)
      else:
        maybeValue = '/*void result*/'
        assigns = ''
      print( '%s /*OptimizedGoCall OK*/ %s%s(%s) // OptimizedGoCall OK' % (assigns, ispec, qfunc.Invoklet(), ', '.join(ins)) )
      for out in outs:
        print('_ = %s' % out)

      if magic_error:
        print 'if %s != nil { panic(%s) /*Check magic_error*/}' % (outs[-1], outs[-1])
        results = results[:-1]

    finally:
        PopPrint()
    print str(buf)

    maybeResult.put(maybeValue, True)

    if results:
      if len(results) > 1:
        print '// OptimizedGoCall: Returning Ytuple', repr(results)
        return Ytuple(results, None)
      else:
        return results[0]
    else:
      return None2

  def Vcall(self, p):
    # fn, args, names, star, starstar
    print '// Vcall: fn:', repr(p.fn)
    print '// Vcall: args:', repr(p.args)
    print '// Vcall: names:', repr(p.names)
    print '// Vcall: star:', repr(p.star)
    print '// Vcall: starstar:', repr(p.starstar)

    def NativeGoTypeName(a):
        """For go types in go_type, go_new, go_make, go_cast."""
        if type(a) is parse.Tfield:
          return '%s.%s' % (a.p.visit(self), a.field)
        elif type(a) is parse.Tvar:
          return a.name
        else:
          raise Exception('Strange thing for go_type: ' + a)

    n = len(p.args)
    self.maxNumCallArgs = max(self.maxNumCallArgs, n)  # Remember max per module.

    # arglist_thunk: As a text sequence with commas.
    arglist_thunk = lambda: ', '.join([Pair(str(a.visit(self))) for a in p.args])
    # argvec_thunk: As a python list, which could be joined with commas, or used separately.
    argvec_thunk = lambda: [a.visit(self) for a in p.args]

    if p.star or p.starstar or any(p.names):
      # Slow road, if `*args` or `**kwargs` or `kw=` gets used.
      called = Pair(p.fn.visit(self))
      if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_star:: %%s}}}\\n", `%s`)}' % called
      return 'inline.TakePJ(%s).(IJCallV).JCallV([]W{%s}, %s, []JKV{%s}, %s) ' % (

          called,

          ', '.join([str(p.args[i].visit(self)) for i in range(len(p.args)) if not p.names[i]]),  # fixed args with no names.

          ('(%s).List()' % p.star.visit(self)) if p.star else 'nil',

          ', '.join(['JKV{"%s", %s}' % (p.names[i], p.args[i].visit(self)) for i in range(n) if p.names[i]]),  # named args.

          ('(%s).Dict()' % p.starstar.visit(self)) if p.starstar else 'nil',
      )

    # Now we're on the fast road.
    if type(p.fn) is parse.Tfield:  # CASE var.Meth(...)
      if type(p.fn.p) is parse.Tvar:

        if p.fn.p.name == 'super':  # CASE super.Meth(...)
          return 'self.%s.M_%d_%s(%s)' % (self.tailSup(self.sup), n, p.fn.field, arglist_thunk())

        if p.fn.p.name in self.imports:  # CASE import.Func(...)
          imp = self.imports[p.fn.p.name]

          if imp.imported[0] == 'go':  # CASE import.Func(...) imported from go.*

            # Try Optimization with QFunc.
            ipath = '/'.join(imp.imported[1:])
            iname = '%s.%s' % (ipath, p.fn.field)
            ispec = 'i_%s.%s' % (p.fn.p.name, p.fn.field)
            argvec = argvec_thunk()  # Call it here, so use them once!
            qfunc = goapi.QFuncs.get(iname)
            if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_i:: %%s %%s %%s %%s}}}\\n", `%s`, `%s`, `%s`, `%s`)}' % (ipath, iname, ispec, qfunc)
            if qfunc:
              maybeResult = Maybe('/*Maybe*/', True)
              sys.stdout.append(maybeResult)
              try:  # Try to use the argvec.
                raise Exception("Disable-OptimizedGoCall")
                return self.OptimizedGoCall(ispec, argvec, qfunc, maybeResult)
              except Exception as ex:
                print '//OptimizedGoCall NO: << %s.%s(%s) >>: %s' % (p.fn.p.name, p.fn.field, ', '.join(['%s_1, %s_2' % Doublet(str(a)) for a in argvec]), repr(ex))

            # Otherwise use reflection with MkGo().  Use the argvec.
            return self.Pair('MkGoJ(%s).JCall(%s) ' % (ispec, ', '.join(['%s_1, %s_2' % Doublet(str(a)) for a in argvec])))
          else:
            # Case impot.func() but not imported from go.
            if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_qG:: %%s.%%s}}}\\n", `%s`, `%s`)}' % (p.fn.p.name, p.fn.field) #YAK
            return self.Pair('%s_%d( i_%s.H_%s, %s) ' % (('JCALL' if n<11 else 'jcall'), n, p.fn.p.name, p.fn.field, arglist_thunk()))

      # General Method Invocation.
      key = '%d_%s' % (n, p.fn.field)
      self.invokes[key] = (n, p.fn.field)
      letterF = 'JF' if self.internal or ((n, p.fn.field) in gen_internals.InternalInvokers) else 'jf'

      invoker = '/*invoker*/ %s_INVOKE_%d_%s' % (letterF, n, p.fn.field)
      invoked = Pair(p.fn.p.visit(self))
      if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_qi:: invoker %%s invoked %%s}}}\\n", `%s`, `%s`)}' % (invoker, invoked)
      general = '/*General*/ %s(%s, %s) ' % (invoker, invoked, arglist_thunk())

      # GOMAXPROCS=2
      # With OPT=append: 17.37 17.71 19.736 19.686
      # Without: 18.244 18.094 18.628
      #if os.getenv('OPT') == 'append' and p.fn.field == 'append':
      if p.fn.field == 'append':
        print '{'
        print '  o1, o2 := %s // optimize_append' % Pair(p.fn.p.visit(self)) 
        print '  var o PJ = inline.CheckTakePJ(o1, o2)'
        print '  if p, ok := o.(*JList); ok {'
        print '    p.PP = append(p.PP, W{%s})' % (arglist_thunk())
        print '  } else {'
        print '    %s(o1, o2, %s)' % (invoker, arglist_thunk())
        print '  }'
        print '}'
        return 'NoneJ_1, NoneJ_2'

      qmeth = goapi.QMeths.get(p.fn.field)
      if qmeth:
        s = self.Serial('sig')
        so = '%s_o' % s
        sob = Bifurcate(so)
        sp = '%s_p' % s
        print 'var %s %s' % (s, qmeth.signature)
        print '_ = %s' % s  # In case of exception.
        print '%s := %s // Optimize QMeth' % (sob, Pair(p.fn.p.visit(self)))
        print 'if !inline.IsEggs(%s) {' % sob
        print '  %s := inline.CheckTakePJ(%s)' % (sp, sob)
        print '  if p, ok := %s.Contents().(%s); ok {' % (sp, qmeth.signature)
        print '    %s = p' % s
        print '  } else {'
        #print '    fmt.Fprintf(os.Stderr, "ZZZ: MISSED: signature: %s type: %%T\\n", %s_o.X.Contents())' % (qmeth.signature, s)
        print '  }'
        print '}'

        print 'var %s_r_1 U = 0' % s
        print 'var %s_r_2 V = V0' % s
        maybeResult = Maybe('/*Maybe*/', True)
        sys.stdout.append(maybeResult)
        print 'if %s != nil {' % s

        argvec = argvec_thunk()  # Call it here, so use them once!

        try:  # Try to use the argvec.
          ####raise Exception("Disable-OptimizedGoCall")
          self.signatures[qmeth.signature] = qmeth.text
          fast = self.OptimizedGoCall('%s' % s, argvec, qmeth, maybeResult)
        except Exception as ex:
          print '//OptimizedGoCallMeth NO: << %s.%s(%s) >>: %s' % (s, p.fn.field, ', '.join(['%s_1, %s_2' % Doublet(str(a)) for a in argvec]), repr(ex))
          #fast = '/*GeneralCallMeth Fast*/ %s(%s, %s) ' % (invoker, sob, ', '.join(['%s_1, %s_2' % Doublet(str(a)) for a in argvec]))
          fast = '/*GeneralCallMeth Fast*/ %s(%s, %s) ' % (invoker, sob, ', '.join([Pair(str(a)) for a in argvec]))
          # moving this line INTO the except:
          ####print '  %s_r = %s' % (s, fast)

        print '} else {'
        #slow = '/*GeneralCallMeth Slow*/ %s(%s, %s) ' % (invoker, sob, ', '.join(['%s_1, %s_2' % Doublet(str(a)) for a in argvec]))
        slow = '/*GeneralCallMeth Slow*/ %s(%s, %s) ' % (invoker, sob, ', '.join([Pair(str(a)) for a in argvec]))
        print '  %s = %s' % (Bifurcate('%s_r' % s), slow)
        print '}'
        print '_,_ = %s' % Bifurcate('%s_r' % s)
        return Yeither(Bifurcate('%s_r' % s), fast, self)

      return self.Pair(general)

    zfn = p.fn.visit(self)
    if type(zfn) is Yspecial:

      if p.fn.name == 'go_type':
        assert len(p.args) == 1, 'go_type got %d args, wants 1' % len(p.args)
        return self.Pair('GoElemType(new(%s))' % NativeGoTypeName(p.args[0]))

      elif p.fn.name == 'go_indirect':
        assert len(p.args) == 1, 'go_indirect got %d args, wants 1' % len(p.args)
        return self.Pair('MkValue(reflect.Indirect(reflect.ValueOf(%s.Contents())))' % p.args[0].visit(self))

      elif p.fn.name == 'go_addr':
        assert len(p.args) == 1, 'go_addr got %d args, wants 1' % len(p.args)
        return self.Pair('MkGo(reflect.ValueOf(%s.Contents()).Addr())' % p.args[0].visit(self))

      elif p.fn.name == 'go_elem':
        assert len(p.args) == 1, 'go_elem got %d args, wants 1' % len(p.args)
        return self.Pair('MkValue(reflect.ValueOf(%s.Contents()).Elem())' % p.args[0].visit(self))

      elif p.fn.name == 'go_new':
        assert len(p.args) == 1, 'go_new got %d args, wants 1' % len(p.args)
        return self.Pair('MkGo(new(%s))' % NativeGoTypeName(p.args[0]))

      elif p.fn.name == 'go_make':
        if len(p.args) == 1:
          return self.Pair('MkGo(make(%s))' % NativeGoTypeName(p.args[0]))
        elif len(p.args) == 2:
          return self.Pair('MkGo(make(%s, int(%s.Int())))' % (NativeGoTypeName(p.args[0]), p.args[1].visit(self)))
        else:
          raise Exception('go_make got %d args, wants 1 or 2' % len(p.args))

      elif p.fn.name == 'go_cast':
        assert len(p.args) == 2, 'go_cast got %d args, wants 2' % len(p.args)
        return self.Pair('GoCast(GoElemType(new(%s)), %s)' % (NativeGoTypeName(p.args[0]), p.args[1].visit(self)))

      elif p.fn.name == 'go_append':
        assert len(p.args) == 2, 'go_append got %d args, wants 2' % len(p.args)
        return self.Pair('GoAppend(%s, %s)' % (p.args[0].visit(self), p.args[1].visit(self)))

      elif p.fn.name == 'len':
        assert len(p.args) == 1, 'len got %d args, wants 1' % len(p.args)
        return self.Pair(DoLen(p.args[0].visit(self)))
        #return Yint('/*Y*/int64(%s.Len())' % p.args[0].visit(self), None)

      elif p.fn.name == 'str':
        assert len(p.args) == 1, 'str got %d args, wants 1' % len(p.args)
        return self.Pair(ForceString(p.args[0].visit(self)))
        #return Ystr('/*Y*/%s.String()' % p.args[0].visit(self), None)

      elif p.fn.name == 'repr':
        assert len(p.args) == 1, 'repr got %d args, wants 1' % len(p.args)
        return self.Pair(Ystr('/*Y*/%s.Repr()' % p.args[0].visit(self), None))

      elif p.fn.name == 'int':
        assert len(p.args) == 1, 'int got %d args, wants 1' % len(p.args)
        return self.Pair(ForceInt(p.args[0].visit(self)))
        #return Yint('/*Y*/%s.ForceInt()' % p.args[0].visit(self), None)

      elif p.fn.name == 'float':
        assert len(p.args) == 1, 'float got %d args, wants 1' % len(p.args)
        return self.Pair(ForceFloat(p.args[0].visit(self)))
        #return Yfloat('/*Y*/%s.ForceFloat()' % p.args[0].visit(self), None)

      else:
        raise Exception('Undefind builtin: %s' % p.fn.name)

    if type(zfn) is not str and zfn.flavor == 'G' and zfn.t.name in self.ydefs:
      fp = self.ydefs[zfn.t.name]
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
        if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_fn::H_%%s_%%s}}}\\n", `%d`, `%s`)}' % (n, zfn.t.name) #YAK
        return self.Pair('H_%d_%s(%s) ' % (n, zfn.t.name, arglist))

    if type(zfn) is Ysuper:  # for calling super-constructor.
      return self.Pair('self.%s.H_%d___init__(%s) ' % (self.tailSup(self.sup), n, arglist_thunk()))

    if DebugCall: print 'if DebugCall>1 { fmt.Fprintf(os.Stderr, "{{{JCALL_str::H_%%s_%%s}}}\\n", `%d`, `%s`)}' % (n, zfn) #YAK
    return self.Pair('/*L1572*/%s_%d( %s, %s )' % (('JCALL' if n<11 else 'jcall'), n, Pair(zfn), arglist_thunk()))

  def Vfield(self, p):
    # p, field
    x = p.p.visit(self)
    if type(x) is Ysuper:
      raise Exception('Special syntax "super" not used with Function Call syntax')
    if type(x) is Yself and not self.cls:
      raise Exception('Using a self field but not in a class definition: field="%s"' % p.field)
    if type(x) is Yself and self.instvars.get(p.field):  # Special optimization for self instvars.
      return Bifurcate('self.M_%s' % p.field)
    elif type(x) is Yimport:
      if x.imp.imported[0] == 'go':
        return ' MkGo(%s.%s) ' % (x, p.field)
      else:
        return Bifurcate('%s.H_%s' % (x, p.field))
    else:
      self.getNeeded[p.field] = True
      letterF = 'JF' if self.internal or p.field in gen_internals.InternalGetters else 'jf'
      return ' %s_GET_%s(%s) ' % (letterF, p.field, Pair(x))

  def Vnative(self, p):
    if self.func:
      print '// { native F'
      for s in p.ss:
        print s
      print '// } native F'
    else:
      # Append to tail, or these would go in jinner_eval_module().
      self.tail.append('// { native M')
      for s in p.ss:
        self.tail.append(s)
      self.tail.append('// } native M')
    return ''

  def preVdef(self, p):
    # Sets:  self.ymeths[] or self.ydefs[].
    #   If method, needs p.cls.name
    args = p.args    # Will drop the initial 'self' element, if in a cls.
    typs = p.typs    # Will drop the initial 'self' element, if in a cls.
    dflts = p.dflts  # Will drop the initial 'self' element, if in a cls.
    if self.cls:
      if len(p.args) > 0 and p.args[0] == 'self':  # User may omit self.
        args = p.args[1:]  # Skip self; it is assumed.
        typs = p.typs[1:]  # Skip self; it is assumed.
        dflts = p.dflts[1:]  # Skip self; it is assumed.
      # module, cls, name, args, dflts, star, starstar, isCtor, typs, rettyp
      nick = '%s.%s::%s' % (self.modname, self.cls.name, p.name)
      if self.cls.name in self.ymeths.dd:
        if p.name in self.ymeths.dd[self.cls.name]:
          raise Exception("Redeclaration of method `%s` not allowed in rye" % nick)
      y = Yfunc(self.modname, self.cls.name, nick, args, dflts, p.star, p.starstar, isCtor=p.isCtor, typs=typs, rettyp=p.rettyp)
      self.ymeths.Put(y, self.cls.name, p.name)
    else:
      # module, cls, name, args, dflts, star, starstar, isCtor, typs, rettyp
      nick = '%s.%s' % (self.modname, p.name)
      if p.name in self.ydefs:
        raise Exception("Redeclaration of class or function `%s` not allowed in rye" % nick)
      y = Yfunc(self.modname, None, nick, args, p.dflts, p.star, p.starstar, isCtor=p.isCtor, typs=typs, rettyp=p.rettyp)
      self.ydefs[p.name] = y

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
      self.Assign(fn_var, parse.Traw(None2))
    else:
      buf = PushPrint()

    # LOOK AHEAD for "yield" and "global" statements.
    finder = parse.YieldAndGlobalFinder()
    finder.Vsuite(p.body)
    self.yields = finder.yields
    self.force_globals = finder.force_globals

    if nesting:
      args, typs, dflts = p.args, p.typs, p.dflts
    elif self.cls:
      y = self.ymeths.Get(self.cls.name, p.name)
      args, typs, dflts = y.args, y.typs, y.dflts
    else:
      y = self.ydefs.get(p.name)
      args, typs, dflts = p.args, p.typs, p.dflts
      if y:
        #if (args, typs, dflts) != (y.args, y.typs, y.dflts):
        #  raise Exception ((args, typs, dflts) , (y.args, y.typs, y.dflts))
        assert (args, typs, dflts) == (y.args, y.typs, y.dflts) , ((args, typs, dflts) , (y.args, y.typs, y.dflts), p.name)

    # Copy scope and add argsPlus to the new one.
    save_scope = self.scope
    if self.scope:
      self.scope = dict([(k, w) for k, w in self.scope.items()]) # Copy it.
    else:
      self.scope = {}

    # Create typPlus the same length as p.argsPlus
    typPlus = [t for t in typs] if typs else []
    if len(typPlus) < len(p.argsPlus):
      typPlus.extend((len(p.argsPlus)-len(typPlus)) * [None])

    print '// zip(p.argsPlus, typPlus): %s' % zip(p.argsPlus, typPlus)
    print '// typPlus:', repr([TypName(e) for e in typPlus])
    for a, t in zip(p.argsPlus, typPlus):
      if t:
        if TypName(t)=='int':
          self.scope[a] = Yint('ai_%s' % a)
          self.scope[a].flavor = 'L'
        elif TypName(t)=='str':
          self.scope[a] = Ystr('as_%s' % a)
          self.scope[a].flavor = 'L'
        else:
          self.scope[a] = 'a_%s' % a
      else:
        self.scope[a] = 'a_%s' % a
      print '// self.scope [%s] = %s' % (a, a if type(a)==str else var(a))

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

           mustBeNone_1, mustBeNone2 := func() (U,V) {
             { wantMore := gen.Wait()
               if !wantMore {
                 gen.Finish()
                 return None
             }}
'''

    p.body.visit(self)

    if self.yields:
      print '''
            return NoneJ_1, NoneJ_2
          }()
          if mustBeNone_1 != NoneJ_1 || mustBeNone_2 != NoneJ_2 {
             panic("Return Value in Generator must be None.")
          }
          gen.Finish()
        }()
        return macro.MkPJ(&gen.JBase)
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
    emptiesV = (', MkListJ(nil), MkDictJ(nil)' if args else 'MkListJ(nil), MkDictJ(nil)') if p.star or p.starstar else ''
    stars = ' %s_1 U, %s_2 V, %s_1 U, %s_2 V' % (Doublet(AOrSkid(p.star)) + Doublet(AOrSkid(p.starstar))) if p.star or p.starstar else ''

    if nesting:
      func_head = 'fn_%s := func' % nesting
      if self.cls:
        func_key = '%s__%s__%s__fn_%s' % (self.modname, self.cls.name, p.name, nesting)
      else:
        func_key = '%s__%s__fn_%s' % (self.modname, p.name, nesting)
    elif self.cls:
      gocls = self.cls.name if self.sup == 'native' else 'J_%s' % self.cls.name
      func_head = 'func (self *%s) N_%d%s_%s' % (gocls, len(args), letterV, p.name)
      func_key = '%s__%s__%s' % (self.modname, self.cls.name, p.name)
    else:
      func_head = 'func H_%d%s_%s' % (len(args), letterV, p.name)
      func_key = '%s__%s' % (self.modname, p.name)

    #if not nesting:
    #  # TODO: Be able to emit this Counter & Init for nested functions, too.
    #  print 'var counter_%s int64' % func_key
    #  print 'func init() {CounterMap["%s"]= &counter_%s}' % (func_key, func_key)

    # Start the function.
    print ' %s(%s %s) (U,V) {' % (func_head, ' '.join(['a_%s_1 U, a_%s_2 V,' % Doublet(a) for a in args]), stars)
    #if not nesting:
    #  print '  counter_%s++' % func_key

    if typs:
      # Check typs of input arguments.
      for (a, t) in zip(args, typs):
        if t:
          print '    CheckTyp("arg %s in func %s", a_%s, %s)' % (a, func_key, a, str(t.visit(self)))

    # Begin Typed Functions
    SUPPORTED_TYPES = {'int': 'int64', 'str': 'string'}

    if typs and not p.star and not p.starstar and not nesting and not self.cls and any(typs) and all([(not t or t.name in SUPPORTED_TYPES) for t in typs]):
        print '    return TH_%d%s_%s(' % (len(args), letterV, p.name)
        for (a, t) in zip(args, typs):
          if t and t.name == 'int':
            print '        a_%s.Int(),' % a,
          elif t and t.name == 'str':
            print '        a_%s.Str(),' % a,
          else:
            print '        a_%s_1, a_%s_2,' % Doublet(a),
        print '    )'
        print '}'

        print 'func TH_%d%s_%s(' % (len(args), letterV, p.name)
        for (a, t) in zip(args, typs):
          if t and t.name == 'int':
            print '        ai_%s %s,' % (a, SUPPORTED_TYPES[t.name])
          elif t and t.name == 'str':
            print '        as_%s %s,' % (a, SUPPORTED_TYPES[t.name])
          else:
            print '        a_%s_1, a_%s_2 U,' % Doublet(a)
        print '    ) (U,V) {'

    else:
      for a, t in zip(p.argsPlus, typPlus):
        if t:
          if TypName(t)=='int':
            print 'var ai_%s int64 = a_%s.Int()' % (a, a)
            print '_ = ai_%s' % a
          elif TypName(t)=='str':
            print 'var as_%s string = a_%s.Str()' % (a, a)
            print '_ = as_%s' % a
          else:
            pass
        else:
          pass

    # End Typed Functions

    # Emit local variables.
    for v, v2 in sorted(self.scope.items()):
      print '// Emit local variables: considering: ', v, v2
      if save_scope is None or v not in save_scope:
        print '// Emit local variables: considering more: ', v, v2
        if STARTS_WITH_V(str(v2)):  # Skip args
          short_var = str(v2).split()[-1][:-2]
          print '   var %s_1 U = NoneJ_1; var %s_2 V = NoneJ_2; _, _ = %s' % (short_var, short_var, Bifurcate(v2))

    if p.rettyp:
      # Start inner function for checking all types of return values.
      print '   retval_1, retval_2 := func() (U,V) { // retval func'

    # Main Body.
    print code2
    print '   return NoneJ_1, NoneJ_2'  # For falling off the bottom.

    if p.rettyp:
      # End inner function for checking all types of return values.
      print '   }() // retval func'
      print '   CheckTyp("return value of function %s", retval, %s)' % (func_key, str(p.rettyp.visit(self)))
      print '   return retval_1, retval_2'

    # End the function
    print ' }'
    print ''

    # Restore the old scope.
    self.scope = save_scope

    n = len(args)
    argnames = ', '.join(['"%s"' % a for a in p.args])
    defaults = ', '.join([(Wifurcate(str(d.visit(self))) if d else 'W{0,V0}') for d in p.dflts])
    print '// Nando:', argnames, defaults

    if nesting:
      tmp = '  ForgeJ(&pNest_%s{JPNewCallable{JCallSpec: &specNest_%s}, fn_%s})  ' % (
          nesting, nesting, nesting)

      self.Assign(fn_var, parse.Traw(tmp))


    # Now for the Nested case, START A PRINT BUFFER.
      buf = PushPrint()

    print '///////////////////////////////'
    print '// name:', p.name

    if nesting:
      print 'var specNest_%s = JCallSpec{Name: "%s__%s", Args: []string{%s}, Defaults: []W{%s}, Star: "%s", StarStar: "%s"}' % (
          nesting, p.name, nesting, argnames, defaults, p.star, p.starstar)

      print ' type pNest_%s struct { JPNewCallable; fn func(%s %s) (U,V) }' % (nesting, ' '.join(['a_%s M,' % a for a in args]), stars)
      print ' func (o *pNest_%s) Contents() interface{} {' % nesting
      print '   return o.fn'
      print ' }'
      if p.star or p.starstar:
        pass  # No direct pNest method; use JCallV().
      else:
        print ' func (o pNest_%s) Call%d(%s) (U,V) {' % (nesting, n, ', '.join([BifurcateUV('a%d' % i) for i in range(n)]))
        print '   return o.fn(%s)' % (', '.join([Bifurcate('a%d' % i) for i in range(n)]))
        print ' }'
      print ''
      print ' func (o pNest_%s) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U,V) {' % nesting
      print '   argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   su, sv := inline.MkPJ(star)'
        print '   ssu, ssv := inline.MkPJ(starstar)'
        print '   return o.fn(%s su, sv, ssu, ssv)' % (' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return o.fn(%s)' % (', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

    elif self.cls:
      if SMALLER and n < 4 and not p.star and not p.starstar and not p.isCtor:
        # Optimize most functions to use JCall%d instead of defining a new struct.
        pass
      else:
        print ' type pMeth_%d_%s__%s struct { JPNewCallable; Rcvr *%s }' % (n, self.cls.name, p.name, gocls)
        print ' func (o *pMeth_%d_%s__%s) Contents() interface{} {' % (n, self.cls.name, p.name)
        print '   return o.Rcvr.M_%d%s_%s' % (n, letterV, p.name)
        print ' }'
        print ' func (o *pMeth_%d_%s__%s) Call%d(%s) (U,V) {' % (n, self.cls.name, p.name, n, ', '.join([BifurcateUV('a%d' % i) for i in range(n)]))
        print '   return o.Rcvr.M_%d%s_%s(%s%s)' % (n, letterV, p.name, ', '.join([BifurcateUV('a%d' % i) for i in range(n)]), emptiesV)
        print ' }'
        print ''
        print ' func (o *pMeth_%d_%s__%s) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U,V) {' % (n, self.cls.name, p.name)
        print '   argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)'
        print '   _, _, _ = argv, star, starstar'

        if p.star or p.starstar:  # If either, we always pass both.
          print '   su, sv := inline.MkPJ(star)'
          print '   ssu, ssv := inline.MkPJ(starstar)'
          print '   return o.Rcvr.M_%dV_%s(%s su, sv, ssu, ssv)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
        else:  # If neither, we never pass either.
          print '   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

        print ' }'
        print ''

    else:
      print ''
      print 'var jspecFunc_%s = JCallSpec{Name: "%s", Args: []string{%s}, Defaults: []W{%s}, Star: "%s", StarStar: "%s"}' % (
          p.name, p.name, argnames, defaults, p.star if p.star else '', p.starstar if p.starstar else '')

      if SMALLER and n < 4 and not p.star and not p.starstar and not p.isCtor:
        # Optimize most functions to use JCall%d instead of defining a new struct.
        formals = ','.join(['a%d_1 U, a%d_2 V' % (i,i) for i in range(n)])
        actuals = ','.join(['a%d_1, a%d_2' % (i,i) for i in range(n)])
        print 'func jfnFunc_%s (%s) (U,V) { return H_%d_%s(%s) }' % (p.name, formals, n, p.name, actuals)
        print ''

        self.glbls[p.name] = ('*jpFunc_%s' % p.name,
                              'ForgeUV((&JCall%d{JPNewCallable{JCallSpec:&jspecFunc_%s}, jfnFunc_%s}))' % (
                                  n, p.name, p.name))
      else:
        print ' type jpFunc_%s struct { JPNewCallable }' % p.name
        print ' func (o *jpFunc_%s) Contents() interface{} {' % p.name
        print '   return macro.TakePJ(%s)' % Bifurcate('H_%s' % p.name)
        print ' }'
        if p.star or p.starstar:
          pass  # No direct jpFunc method; use JCallV().
        else:
          print ' func (o jpFunc_%s) Call%d(%s) (U,V) {' % (p.name, n, ', '.join([BifurcateUV('a%d' % i) for i in range(n)]))
          print '   return H_%d_%s(%s)' % (n, p.name, ', '.join([Bifurcate('a%d' % i) for i in range(n)]))
          print ' }'
        print ''
        print ' func (o jpFunc_%s) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U,V) {' % p.name
        print '   argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)'
        print '   _, _, _ = argv, star, starstar'

        # TODO: I think this is old, before named params.
        if p.star or p.starstar:  # If either, we always pass both.
          print '   su, sv := inline.MkPJ(star)'
          print '   ssu, ssv := inline.MkPJ(starstar)'
          print '   return H_%dV_%s(%s su, sv, ssu, ssv)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
        else:  # If neither, we never pass either.
          print '   return H_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d].U, argv[%d].V' % (i, i) for i in range(n)]))

        print ' }'
        print ''

        self.glbls[p.name] = ('*jpFunc_%s' % p.name,
                              'ForgeUV((&jpFunc_%s{JPNewCallable{JCallSpec:&jspecFunc_%s}}))' % (p.name, p.name))

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
      return 'J_%s' % sup.name
    elif type(sup) == parse.Tfield:
      return 'i_%s.J_%s' % (sup.p.name, sup.field)
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def tailSup(self, sup):
    if type(sup) == parse.Tvar:
      return 'J_%s' % sup.name
    elif type(sup) == parse.Tfield:
      return 'J_%s' % sup.field
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def Vclass(self, p):
    # name, sup, things
    if self.cls:
      raise Exception("Nested classes not supported.")

    self.cls = p
    self.sup = p.sup
    self.instvars = {}

    gocls = self.cls.name if self.sup == 'native' else 'J_%s' % self.cls.name
    # Emit all the methods of the class (and possibly other members).
    for x in p.things:
      x.visit(self)
    self.cls = None

    buf = PushPrint()

    # Emit the struct for the class.
    if self.sup != 'native':
      print '''
 type J_%s struct {
   %s
   %s
 }

 func init() {
   if Classes == nil { Classes = make(map[string]reflect.Type) }
   Classes[`%s.J_%s`] = reflect.TypeOf(J_%s{})
 }

''' % (p.name, self.qualifySup(p.sup),
       '\n'.join([BifurcateUVSemi('M_%s' % x) for x in sorted(self.instvars)]),
       self.modname if self.modname else 'main', p.name, p.name)

    if self.sup != 'native':
      print '''
 func (o *J_%s) PtrJ_%s() *J_%s {
   return o
 }
''' % (p.name, p.name, p.name)

    if self.sup != 'native':
      print '''
 func (o *J_%s) PtrJ_object() *J_object {
   return &o.J_object
 }
''' % (p.name, )

    # For all the instance vars
    print ''
    for iv in sorted(self.instvars):
      slots = Bifurcate('o.M_%s' % iv)
      print ' func (o *J_%s) GET_%s() (U,V) { return %s }' % (p.name, iv, slots)
      print ' func (o *J_%s) SET_%s(x_1 U, x_2 V) { %s = x_1, x_2 }' % (p.name, iv, slots)
      print ''
    print ''

    # For all the methods
    print ''
    if self.ymeths.dd.get(p.name):
      for m, mp in sorted(self.ymeths.dd[p.name].items()):  # Yfuncs for methods.
        args = mp.args
        dflts = mp.dflts
        n = len(args)

        argnames = ', '.join(['"%s"' % a for a in args])
        defaults = ', '.join([(Wifurcate(str(d.visit(self))) if d else 'W{0,V0}') for d in dflts])

        print 'var specMeth_%d_%s__%s = JCallSpec{Name: "%s::%s", Args: []string{%s}, Defaults: []W{%s}, Star: "%s", StarStar: "%s"}' % (
            n, p.name, m, p.name, m, argnames, defaults, mp.star, mp.starstar)

        if SMALLER and n < 4 and not mp.star and not mp.starstar and not mp.isCtor:
          # Optimize most functions to use JCall%d instead of defining a new struct.
          formals = ','.join([BifurcateUV('a%d' % i) for i in range(n)])
          actuals = ','.join([Bifurcate('a%d' % i) for i in range(n)])
          print ''

          print 'func (o *%s) GET_%s() (U,V) {' % (gocls, m)
          print '  z := (&JCall%d{JPNewCallable{JCallSpec:&specMeth_%d_%s__%s}, o.N_%d_%s})' % (n, n, p.name, m, n, m)
          print '  return inline.ForgeUV(z)'
          print '}'
        else:
          print 'func (o *%s) GET_%s() (U,V) {' % (gocls, m)
          print '  z := (&jpMeth_%d_%s__%s {JPNewCallable{JCallSpec: &specMeth_%d_%s__%s}, o})}' % (n, p.name, m, n, p.name, m)
          print '  return inline.ForgeUV(z)'
          print '}'

    # Special methods for classes.
    if self.sup != 'native':
      print 'func (o *J_%s) Rye_ClearFields__() {' % p.name
      for iv in sorted(self.instvars):
        print '   %s = NoneJ_1, NoneJ_2' % Bifurcate('o.M_%s' % iv)
      if p.sup and type(p.sup) is parse.Tvar:
        print '// superclass: %s' % p.sup.visit(self)
        if p.sup.name not in ['native', 'object']:
          print '   o.J_%s.Rye_ClearFields__()' % p.sup.name
      if p.sup and type(p.sup) is parse.Tfield:
        print '// superclass: %s' % p.sup.visit(self)
        print '   o.J_%s.Rye_ClearFields__()' % p.sup.field
      print '}'

      print ''
      print 'func (o *jpFunc_%s) Superclass() (U,V) {' % (p.name)
      if p.sup and type(p.sup) is parse.Tvar:
        print '  return %s' % Pair(p.sup.visit(self))
      else:
        print '  return NoneJ_1, NoneJ_2'
      print '}'
      print ''
      print 'func (o *J_%s) PType() (U,V) { return H_%s_1, H_%s_2 }' % (p.name, p.name, p.name)
      print 'func (o *jpFunc_%s) Repr() string { return "%s" }' % (p.name, p.name)
      print 'func (o *jpFunc_%s) String() string { return "<class %s>" }' % (p.name, p.name)
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

class Maybe(object):
  def __init__(self, s, enable):
    self.s = s
    self.enable = enable
  def put(self, s, enable):
    self.s = s
    self.enable = enable
  def __str__(self):
    return str(self.s)+'\n' if self.enable else '\n'

class Buffer(object):
  def __init__(self):
    self.b = []
  def append(self, x):
    self.b.append(x)
  def write(self, x):
    self.b.append(str(x))  # Forces immutable copy.
  def flush(self):
    pass
  def __str__(self):
    z = ''.join([str(e) for e in self.b])
    return z

def DoLen(a):
  if type(a) != str:
    z = a.DoLen()
    if z: return z
  return Yint('/*G.DoLen else*/ int64(/*global DoLen Yint*/ macro.LenJ(%s))' % a, None)

def DoAdd(a, b):
  if type(a) != str:
    z = a.DoAdd(b)
    if z: return z
  return '/*DoAdd*/ macro.AddJ(%s, %s)' % (Pair(str(a)), Pair(str(b)))

def DoSub(a, b):
  if type(a) != str:
    z = a.DoSub(b)
    if z: return z
  return Pair('/*DoSub*/macro.SubJ(%s, %s)' % ((Pair(str(a))), (Pair(str(b)))))
def DoMul(a, b):
  if type(a) != str:
    z = a.DoMul(b)
    if z: return z
  return Pair('/*DoMul*/macro.MulJ(%s, %s)' % (((Pair(str(a))), (Pair(str(b))))))
def DoDiv(a, b):
  if type(a) != str:
    z = a.DoDiv(b)
    if z: return z
  return Pair('/*DoDiv*/macro.DivJ(%s, %s)' % ((Pair(str(a))), (Pair(str(b)))))
def DoMod(a, b):
  if type(a) != str:
    z = a.DoMod(b)
    if z: return z
  return Pair('/*DoMod*/macro.ModJ(%s, %s)' % ((Pair(str(a))), (Pair(str(b)))))

def DoEQ(a, b):
  if type(a) != str:
    z = a.DoEQ(b)
    if z: return z
  return Ybool('(macro.EQJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)
def DoNE(a, b):
  if type(a) != str:
    z = a.DoNE(b)
    if z: return z
  return Ybool('(macro.NEJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)
def DoLT(a, b):
  if type(a) != str:
    z = a.DoLT(b)
    if z: return z
  return Ybool('(macro.LTJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)
def DoLE(a, b):
  if type(a) != str:
    z = a.DoLE(b)
    if z: return z
  return Ybool('(macro.LEJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)
def DoGT(a, b):
  if type(a) != str:
    z = a.DoGT(b)
    if z: return z
  return Ybool('(macro.GTJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)
def DoGE(a, b):
  if type(a) != str:
    z = a.DoGE(b)
    if z: return z
  return Ybool('(macro.GEJ(%s,%s))' % (Pair(str(a)), Pair(str(b))), None)

def DoNot(a):
  return '/*DoNot*/!(%s)' % AsBool(a)

def ForceInt(a):
  if type(a) != str:
    z = a.ForceInt()
    if z: return z
  return Yint('/*ForceInt*/ForceIntJ(%s)' % a)

def ForceFloat(a):
  if type(a) != str:
    z = a.ForceFloat()
    if z: return z
  return Yfloat('/*ForceFloat*/ForceFloatJ(%s)' % a)

def ForceString(a):
  if type(a) != str:
    z = a.ForceString()
    if z: return z
  return Ystr('/*ForceString*/StringJ(%s)' % a)

def AsBool(a):
  if type(a) != str:
    z = a.AsBool()
    if z: return z
  return '/*AsBool*/BoolJ(%s)' % a

def AsInt(a):
  if type(a) != str:
    z = a.AsInt()
    if z: return z
  return '/*AsInt*/IntJ(%s)' % a

def AsFloat(a):
  if type(a) != str:
    z = a.AsFloat()
    if z: return z
  return '/*AsFloat*/FloatJ(%s)' % a

def AsByt(a):
  if type(a) != str:
    z = a.AsByt()
    if z: return z
  return '/*AsByt*/BytesJ(%s)' % str(a)

def AsStr(a):
  if type(a) != str:
    z = a.AsStr()
    if z: return z
  return '/*AsStr*/StrJ(%s)' % str(a)

def DoAssign(a, b):
  if type(a) != str:
    z = a.DoAssign(b)
    if z: return z
  print '/*G.DoAssign*/ %s = %s' % (a, b)
  return True

class Ybase(object):
  """Ybase: Future Optimized Typed values."""
  def DoLen(self): return ''
  def AsBool(self): return ''
  def ForceInt(self): return ''
  def ForceFloat(self): return ''
  def AsInt(self): return ''
  def AsFloat(self): return ''
  def AsByt(self): return ''
  def AsStr(self): return ''
  def ForceString(self): return ''
  def DoAdd(self, b): return ''
  def DoSub(self, b): return ''
  def DoMul(self, b): return ''
  def DoDiv(self, b): return ''
  def DoMod(self, b): return ''
  def DoEQ(self, b): return ''
  def DoNE(self, b): return ''
  def DoLT(self, b): return ''
  def DoLE(self, b): return ''
  def DoGT(self, b): return ''
  def DoGE(self, b): return ''
  def DoAssign(self, b): return False

class YbaseTyped(Ybase):
  pass

class Ybool(YbaseTyped):
  def __init__(self, y, s):
    self.flavor = ''
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'macro.MkBoolJ(%s)' % self.y
    return '/*Ybool.str*/%s' % self.s
  def ForceInt(self):
    return Yint('/*Ybool.ForceInt*/ BoolToInt64(%s)' % self.y)
  def ForceFloat(self):
    return Yfloat('/*Ybool.ForceFloat*/ BoolToFloat64(%s)' % self.y)
  def AsInt(self):
    return '/*Ybool.AsInt*/ BoolToInt64(%s)' % self.y
  def AsFloat(self):
    return '/*Ybool.AsFloat*/ BoolToFloat64(%s)' % self.y
  def AsBool(self):
    return str(self.y)
  def DoAssign(self, b):
    if self.y:
      print '/*Ybool.DoAssign*/%s = %s' % (self.y, AsBool(b))
      return True

class Yint(YbaseTyped):
  def __init__(self, y, s=None):
    self.flavor = ''
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'macro.MkIntJ(int64(%s))' % str(self.y)
    return '/*Yint.str*/%s' % self.s
  def ForceInt(self):
    return self
  def ForceFloat(self):
    if self.y: return Yfloat('float64(%s)' % self.y)
  def AsInt(self):
    return str(self.y)
  def AsFloat(self):
    if self.y: return 'float64(%s)' % self.y
  def AsBool(self):
    if self.y: return '/*Yint.AsBool*/((%s) != 0)' % self.y
  def doArith(self, b, op):
    if self.y:
      if type(b) is Ybool:
        if b.y: return Yint('(/*Yint.doArith*/ int64(%s) %s BoolToInt64(%s) )' % (self.y, op, b.y), None)
      if type(b) is Yint:
        if b.y: return Yint('(/*Yint.doArith*/ int64(%s) %s int64(%s) )' % (self.y, op, b.y), None)
      if type(b) is Yfloat:
        if b.y: return Yfloat('(/*Yint.doArith*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''
  def DoAdd(self, b): return self.doArith(b, '+')
  def DoSub(self, b): return self.doArith(b, '-')
  def DoMul(self, b): return self.doArith(b, '*')
  def DoDiv(self, b): return self.doArith(b, '/')
  def DoMod(self, b): return self.doArith(b, '%')

  def doRelop(self, b, op):
    if self.y:
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
  def DoAssign(self, b):
    if self.y:
      print '/*Yint.DoAssign*/%s = %s' % (self.y, AsInt(b))
      return True


class Yfloat(YbaseTyped):
  def __init__(self, y, s=None):
    self.flavor = ''
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'macro.MkFloatJ(float64(%s))' % str(self.y)
    return '/*Yfloat.str*/%s' % self.s
  def ForceInt(self):
    if self.y: return 'int64(%s)' % str(self.y)
  def ForceFloat(self):
    return self
  def AsFloat(self):
    if self.y: return str(self.y)
  def AsBool(self):
    if self.y: return '/*Yfloat.AsBool*/((%s) != 0.0)' % self.y

  def doArith(self, b, op):
    if self.y:
      if type(b) is Ybool:
        if b.y: return Yfloat('(/*YYfloat.doArith*/ float64(%s) %s BoolToFloat64(%s) )' % (self.y, op, b.y), None)
      if type(b) in [Yfloat, Yint]:
        if b.y: return Yfloat('(/*YYfloat.doArith*/ float64(%s) %s float64(%s) )' % (self.y, op, b.y), None)
    return ''
  def DoAdd(self, b): return self.doArith(b, '+')
  def DoSub(self, b): return self.doArith(b, '-')
  def DoMul(self, b): return self.doArith(b, '*')
  def DoDiv(self, b): return self.doArith(b, '/')
  def DoMod(self, b): return self.doArith(b, '%')

  def doRelop(self, b, op):
    if self.y:
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
  def DoAssign(self, b):
    if self.y:
      print '/*Yfloat.DoAssign*/%s = %s' % (self.y, AsFloat(b))
      return True

class Ystr(YbaseTyped):
  def __init__(self, y, s=None):
    self.flavor = ''
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'macro.MkStrJ(%s)' % str(self.y)
    return '/*Ystr.str*/%s' % self.s
  def DoLen(self):
    if self.y: return Yint('int64(len(%s))' % self.y, None)
  def AsByt(self):
    if self.y: return '/*Ystr.AsByt*/[]byte(%s)' % self.y
  def AsStr(self):
    return str(self.y)
  def ForceString(self):
    return self
  def AsBool(self):
    if self.y: return '/*Ystr.AsBool*/(%s != "")' % self.y

  def doRelop(self, b, op):
    if self.y:
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
  def DoAssign(self, b):
    if self.y:
      print '/*YStr.DoAssign*/%s = %s' % (self.y, AsStr(b))
      return True

class Ybyt(YbaseTyped):
  def __init__(self, y, s):
    self.flavor = ''
    self.y = y
    self.s = s
  def __str__(self):
    if not self.s:
      self.s = 'MkBytUV(%s)' % str(self.y)
    return '/*Ybyt.str*/%s' % self.s
  def DoLen(self):
    if self.y: return Yint('int64(len(%s))' % self.y, None)
  def AsByt(self):
    if self.y: return str(self.y)
  def AsStr(self):
    if self.y: return '/*Ybyt.AsStr*/string(%s)' % self.y
  def ForceString(self):
    if self.y: return Ystr('/*Ybyt.ForceString*/string(%s)' % self.y)
  def AsBool(self):
    if self.y: return '/*Ybyt.AsBool*/(%s != "")' % self.y

  def doRelop(self, b, op):
    if self.y:
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
  def DoAssign(self, b):
    if self.y:
      print '/*YByt.DoAssign*/%s = %s' % (self.y, AsByt(b))
      return True

class Yeither(Ybase):
  def __init__(self, a, b, codegen):
    self.flavor = ''
    self.a = a  # a must be an M, possibly 0,V0.
    self.b = b  # b can be an optimized Y type.
    self.codegen = codegen
    self.s = ''
    assert self.a
    assert self.b
    print "// (* Yeither: %s :: %s ; %s :: %s *)" % (repr(a), type(a), repr(b), type(b))
  def __str__(self):
    if not self.s:
      # set a to b if a is nil
      print '/*Yeither.str*/ if inline.IsEggs(%s) {' % self.a
      print '  %s = %s' % (self.a, Pair(self.b))
      print '}'
      self.s = self.a
    return '/*Yeither.str*/%s' % self.s
  def DoLen(self):
    ser = self.codegen.Serial('dolen')
    print "// (* Yeither: %s :: %s ; %s :: %s *)" % (repr(self.a), type(self.a), repr(self.b), type(self.b))
    print 'var %s int64' % ser
    print 'if %s == 0,V0 {' % self.a
    #print '    fmt.Fprintf(os.Stderr, "ZZZ: FAST: DoLen: %s\\n")' % type(self.b)
    print '    /*Yeither DoLen b (fast)*/ %s = %s' % (ser, AsInt(DoLen(self.b)))
    print '} else {'
    #print '    fmt.Fprintf(os.Stderr, "ZZZ: SLOW: DoLen: %s\\n")' % type(self.a)
    print '    /*Yeither DoLen a (slow)*/ %s = %s' % (ser, AsInt(DoLen(self.a)))
    print '}'
    return Yint(ser, None)

class Ytuple(YbaseTyped):
  def __init__(self, y, s):
    self.flavor = ''
    self.y = y
    self.s = s
    assert self.y is not None
    assert type(self.y) is list
  def __str__(self):
    if not self.s:
      self.s = '   /*Ytuple.str*/MkTuple([]M{%s})   ' % ', '.join([str(e) for e in self.y])
    return str(self.s)
  def DoLen(self):
    return Yint(str(len(self.y)), None)

  def AsBool(self):
    # Bool of a tuple is size greater than zero.
    return 'true' if self.y else 'false'

class Yimport(Ybase):
  def __init__(self, s, imp):
    self.flavor = ''
    self.s = s
    self.imp = imp  # imports[] object
  def __str__(self):
    return '/*Yimport.str*/%s' % self.s

class Yprim(Ybase):
  def __init__(self, t, s):
    self.flavor = ''
    self.t = t  # T node
    self.s = s  # String for backwards compat
    print '// Yprim: constructing %s with %s' % (type(self).__name__, vars(self))
    if ',' in s:
      #raise Exception("Comma in Yprim: s=%s", s)
      print '// ######################### BAD Comma in Yprim:', s
  def __str__(self):
    return self.s

class Yself(Yprim):
  def __str__(self):
    return '/*Yself.str*/inline.MkPJ(&self.JBase)'

class Ysuper(Yprim):
  def __str__(self):
    raise Exception("cannot str(Ysuper)")

class Yvar(Yprim):  # Local or global
  def __str__(self):
    #return '/*Yvar.str*/%s' % self.s
    return '/*Yvar.str*/%s' % Bifurcate(self.s)

class Yspecial(Yprim):
  pass
class Ylit(Yprim):  # TODO: stop using Ylit
  pass

pass

class Yfunc(object):
  def __init__(self, module, cls, nick, args, dflts, star, starstar, isCtor, typs, rettyp):
    self.module = module
    self.cls = cls
    self.nick = nick
    self.args = args  # self already removed, if method.
    self.dflts = dflts  # self already removed, if method.
    self.star = star
    self.starstar = starstar
    self.isCtor = isCtor
    self.typs = typs  # self already removed, if method.
    self.rettyp = rettyp
  def __str__(self):
    return 'Yfunc:(%s)' % self.nick

  def NotYetUsed_CallSpec(self):
    argnames = ', '.join(['"%s"' % a for a in self.args])
    defaults = ', '.join([(Wifurcate(str(d.visit(self))) if d else 'W{0,V0}') for d in self.dflts])
    return 'PCallable{Nick: "%s", Args: []string{%s}, Defaults: []W{%s}, Star: "%s", StarStar: "%s"}' % (
        self.nick, argnames, defaults, self.star, self.starstar)

def AOrSkid(s):
  if s:
    return 'a_%s' % s
  else:
    return '_'

class DeepDict(object):
  def __init__(self):
    self.dd = {}
  def Get(self, *kk):
    d = self.dd
    for k in kk:
      if type(d) is not dict:
        raise Exception("got %s instead of dict in DeepDict" % type(d))
      d = d.get(k)
      if d is None:
        return None
    return d
  def Put(self, a, *kk):
    d = self.dd
    for k in kk[:-1]:
      n = d.get(k)
      if n is None:
        n = {}
        d[k] = n
      d = n
      if type(d) is not dict:
        raise Exception("got %s instead of dict in DeepDict" % type(d))
    d[kk[-1]] = a
  def __str__(self):
    return repr(self.dd)
  def __repr__(self):
    return repr(self.dd)

pass
