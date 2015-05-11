#import codecs
import md5
import os
import re
import sys

rye_rye = False
if rye_rye:
  from lib import data
  from go import strconv
  from . import parse
  from . import samples
else:
  import parse
  import samples

RYE_FLOW = os.getenv('RYE_FLOW')
BUILTINS = list( 'go_cast go_type go_new go_make go_append'.split())

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([ \t\n]*[#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')
RE_PRAGMA = re.compile('[ \t]*[#][#][A-Za-z:()]+')

RE_KEYWORDS = re.compile(
    '\\b(del|say|from|class|def|native|if|elif|else|while|True|False|None|print|and|or|try|except|raise|yield|return|break|continue|pass|as|go|defer|with|global|assert|must|lambda|switch)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[-]=|[*]=|/=|//|<<|>>>|>>|==|!=|<=|>=|[*][*]|[.][.]')
RE_OPS = re.compile('[-.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_FLOAT = re.compile('[+-]?[0-9]+[.][-+0-9eE]*')
RE_INT = re.compile('(0[Xx][0-9A-Fa-f]+|[+-]?[0-9]+)')

RE_STR = re.compile('(["](([^"\\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\\n]|[\\\\].)*)[\'])')
RE_STR2 = re.compile('(?s)[`]([^`]*)[`]')
RE_STR3 = re.compile('(?s)("""(([^\\\\]|[\\\\].)*?)"""|\'\'\'(([^\\\\]|[\\\\].)*?)\'\'\')')

RE_SEMI = re.compile(';')

RE_WORDY_REL_OP = re.compile('\\b(not\\s+in|is\\s+not|in|is)\\b')
RE_NOT_IN = re.compile('^not\\s+in$')
RE_IS_NOT = re.compile('^is\\s*not$')

RE_NOT_NEWLINE = re.compile('[^\\n]')

### Experimental: For string interpolation, if we do that:
# RE_NEST1 = '[^()]*([(][^()]*[)][^()]*)*[^()]*'
# RE_SUBST = re.compile('(.*)[\\\\][(](' + NEST1 + ')[)](.*)')

TAB_WIDTH = 8

DETECTERS = [
  [RE_PRAGMA, 'P'],
  [RE_KEYWORDS, 'K'],
  [RE_WORDY_REL_OP, 'W'],
  [RE_ALFA, 'A'],
  [RE_FLOAT, 'F'],
  [RE_INT, 'N'],
  [RE_LONG_OPS, 'L'],
  [RE_OPS, 'O'],
  [RE_GROUP, 'G'],
  [RE_STR3, 'S'],
  [RE_STR2, 'S'],
  [RE_STR, 'S'],
  [RE_SEMI, ';;'],
]

UNARY_OPS = {
  '+': 'UnaryPlus',
  '-': 'UnaryMinus',
  '~': 'UnaryInvert',
}
SHIFT_OPS = {
  '<<': 'ShiftLeft',
  '>>': 'ShiftRight',
  '>>>': 'UnsignedShiftRight',
}
ADD_OPS = {
  '+': 'Add',
  '-': 'Sub',
}
MUL_OPS = {
  '*': 'Mul',
  '/': 'Div',
  '//': 'IDiv',
  '%': 'Mod',
}
REL_OPS = {
  '==': 'EQ',
  '!=': 'NE',
  '<': 'LT',
  '<=': 'LE',
  '>': 'GT',
  '>=': 'GE',
}

MaxNumCallArgs = -1

FIRST_WORD = re.compile('^([^\\s]*)').match
def FirstWord(s):
  return FIRST_WORD(s).group(1)

TRIM_PRAGMA = re.compile('\\s*[#][#](\\w*)').match
def TrimPragma(s):
  m = TRIM_PRAGMA(s)
  if m:
    return m.group(1)
  raise Exception('Bad pragma: %s' % repr(s))

NOT_PRINTABLE_ASCII = re.compile('[^!-~]')
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

def Bad(format, *args):
  raise Exception(format % args)

############################################################

SerialNum = 10
def Serial(s):
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (s, SerialNum)

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
    self.gsNeeded = {}      # keys are getter/setter names.

  def InjectForInternal(self, stuff):
    self.invokes, self.defs, self.gsNeeded = stuff

  def ExtractForInternal(self):
    stuff = self.invokes, self.defs, self.gsNeeded
    return stuff

  def GenModule(self, modname, path, tree, cwp=None, internal=""):
    self.cwp = cwp
    self.path = path
    self.modname = modname
    if not internal:
      self.glbls['__name__'] = ('P', 'MkStr("%s")' % modname)

    self.func_level = 0
    self.func = None
    self.yields = None
    self.force_globals = {}

    self.internal = internal
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

    # LOOK FOR MAIN
    main_def = None
    for th in tree.things:
      if type(th) == parse.Tdef:
        if th.name == 'main':
          main_def = th
    # ADD A MAIN, if there isn't one.
    if not main_def and not internal:
      main_def = parse.Tdef('main', ['argv'], [None], None, None, parse.Tsuite([]))
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
        pkg = '/'.join(vec)
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
        print ''
        print '// NEWCTOR:', c_name
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

        print '//   NEWCTOR:', repr(c_args)
        print '//   NEWCTOR:', repr(c_dflts)
        print '//   NEWCTOR:', repr(c_star)
        print '//   NEWCTOR:', repr(c_starstar)
        print '//'

        natives = [
              '   z := new(C_%s)' % th.name,
              '   z.Self = z',
              '   z.Rye_ClearFields__()',
        ]
        c1 = parse.Tnative(natives)

        c2 = parse.Tassign(parse.Tvar('rye_result__'), parse.Traw('z'))

        # Tcall: fn, args, names, star, starstar
        call = parse.Tcall(parse.Tfield(parse.Tvar('rye_result__'), '__init__'), [parse.Tvar(a) for a in c_args], c_args, x_star, x_starstar)
        c3 = parse.Tassign(parse.Traw('_'), call)

        c4 = parse.Treturn([parse.Tvar('rye_result__')])

        suite = parse.Tsuite([c1, c2, c3, c4])
        ctor = parse.Tdef(c_name, c_args, c_dflts, c_star, c_starstar, suite)
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
    print ' var _ = MkInt'  # From rye runtime.
    print ''

    # BEGIN: Eval_Module, innter_eval_module
    if self.internal:
      print ' func eval_module_internal_%s () P {' % self.internal
    else:
      print ' var eval_module_once bool'
      print ' func Eval_Module () P {'
      print '   if eval_module_once == false {'
      print '     _ = inner_eval_module()'
      print '     eval_module_once = true'
      print '   }'
      print '   return ModuleObj'
      print ' }'
      print ' func inner_eval_module () P {'

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
      print 'var G_%s P // %s' % (g, t)
    print ''
    print ' func init /*New_Module*/ () {'
    for g, (t, v) in sorted(self.glbls.items()):
      print '   G_%s = %s' % (g, v)
      if v != "None":
        print '   G_%s.SetSelf(G_%s)' % (g, g)

        if False:  # These were never used.
          if self.internal:
            print '   Globals["%s.%s"] = G_%s' % (self.modname, g, g)
          else:
            print '   Globals["%s/%s.%s"] = G_%s' % (self.cwp, self.modname, g, g)
    print ' }'
    print ''
    print 'var %s = map[string]*P {' % ('BuiltinMap' if self.internal else 'ModuleMap')
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

    if self.internal and self.internal != "builtins":
      return

    for key, (n, fieldname) in sorted(self.invokes.items()):
      self.gsNeeded[fieldname] = True
      formals = ', '.join(['a_%d P' % i for i in range(n)])
      args = ', '.join(['a_%d' % i for i in range(n)])
      print 'func f_INVOKE_%d_%s(fn P, %s) P {' % (n, fieldname, formals)
      print '    fn = fn.GetSelf()'
      print '  switch x := fn.(type) {   '
      print '  case i_INVOKE_%d_%s:         ' % (n, fieldname)
      print '    return x.M_%d_%s(%s)         ' % (n, fieldname, args)
      print '  case i_GET_%s:         ' % fieldname
      print '    tmp := x.GET_%s()    ' % fieldname
      print '    return call_%d(tmp, %s)' % (n, ', '.join(['a_%d' % j for j in range(n)]))
      print '    '

      print '  case *PGo:                '
      print '    return x.Invoke("%s", %s) ' % (fieldname, args)
      print '  }'
      print '  panic(fmt.Sprintf("Cannot invoke \'%s\' with %d arguments on %%v", fn))' % (fieldname, n)
      print '}'
      print 'type i_INVOKE_%d_%s interface { M_%d_%s(%s) P }' % (n, fieldname, n, fieldname, formals)
    print ''

    for iv in sorted(self.gsNeeded):
      print ' type i_GET_%s interface { GET_%s() P }' % (iv, iv)
      print ' type i_SET_%s interface { SET_%s(P) }' % (iv, iv)

      print 'func f_GET_%s(h P) P {' % iv
      print '  switch x := h.(type) { '
      print '  case i_GET_%s:         ' % iv
      print '    return x.GET_%s()    ' % iv
      print '  }'
      print '   return h.FetchField("%s") ' % iv
      print '}'
      print ''

      print 'func f_SET_%s(h P, a P) {' % iv
      print '  switch x := h.(type) { '
      print '  case i_SET_%s:         ' % iv
      print '    x.SET_%s(a)    ' % iv
      print '    return'
      print '  }'
      print '    h.StoreField("%s", a)' % iv
      print '}'
      print ''
    print ''

    maxCall = 1 + (4 if self.internal == "builtins" else MaxNumCallArgs)
    for i in range(maxCall):
      print '  type i_%d interface { Call%d(%s) P }' % (i, i, ", ".join(i * ['P']))
      print '  func call_%d (fn P, %s) P {' % (i, ', '.join(['a_%d P' % j for j in range(i)]))
      print '    switch f := fn.(type) {'
      print '      case i_%d:' % i
      print '        return f.Call%d(%s)' % (i, ', '.join(['a_%d' % j for j in range(i)]))
      print '      case ICallV:'
      print '        return f.CallV([]P{%s}, nil, nil, nil)' % ', '.join(['a_%d' % j for j in range(i)])
      print '    }'
      print '    panic(fmt.Sprintf("No way to call: %v", fn))'
      print '  }'
    print ''

  def Gloss(self, th):
    print '// @ %d @ %d @ %s' % (th.where, th.line, th.gloss)

  def Ungloss(self, th):
    print '// $ %d $ %d $ %s' % (th.where, th.line, th.gloss)

  def Vexpr(self, p):
    print ' _ = %s' % p.a.visit(self)

  def AssignFieldAFromRhs(self, a, rhs, pragma):
      lhs = a.p.visit(self)
      if type(lhs) is Zself:  # Special optimization for self.
        self.instvars[a.field] = True
        lhs = 'self.M_%s' % a.field
        print '   %s = %s' % (lhs, rhs)
      elif type(lhs) is Zimport:  # For module variables.
        lhs = '%s.G_%s' % (lhs, a.field)
        print '   %s = %s' % (lhs, rhs)
      else:
        self.gsNeeded[a.field] = True
        print '   f_SET_%s(%s, %s)' % (a.field, lhs, rhs)

  def AssignItemAFromRhs(self, a, rhs, pragma):
        p = a.a.visit(self)
        q = a.x.visit(self)
        print '   (%s).SetItem(%s, %s)' % (p, q, rhs)

  def AssignTupleAFromB(self, a, b, pragma):
        serial = Serial('detuple')
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
    print '// @@@@@@ AssignAFromB: %s %s %s' % (type(a), a, self.scope)
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
        self.glbls[a.name] = ('P', 'None')

    elif type_a is parse.Traw:
      lhs = a.raw

    else:
      raise Exception('Weird Assignment, a class is %s, A IS (%s) (%s) B IS (%s) (%s)' % (type(a).__name__, a, a.visit(self), b, b.visit(self)))

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
        print '   fmt.Fprintln(%s, "#%s %s ", self.ShortPointerHashString(), " # ", %s.Repr())' % (
            'P(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStderr()',
            where,
            str(p.code).replace('"', '\\"'),
            '.Repr(), "#", '.join([str(v) for v in vv]))
      else:
        print '   fmt.Fprintln(%s, "#%s %s # ", %s.Repr())' % (
            'P(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStderr()',
            where,
            str(p.code).replace('"', '\\"'),
            '.Repr(), "#", '.join([str(v) for v in vv]))
    else:
      if p.xx.trailing_comma:
        printer = Serial('printer')
        print '%s := %s' % (
            printer,
            'P(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()')
        for i in range(len(vv)):
          print 'io.WriteString(%s, %s.String()) // i=%d' % (
              printer, str(vv[i]), i)
          print 'io.WriteString(%s, " ")' % printer
      else:
        if vv:
          print '   fmt.Fprintln(%s, %s.String())' % (
              'P(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()',
              '.String(), '.join([str(v) for v in vv]))
        else:
          print '   fmt.Fprintln(%s, "")' % (
              'P(%s).Contents().(io.Writer)' % p.w.visit(self) if p.w else 'CurrentStdout()')
          

  def Vimport(self, p):
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
        self.modname, p.line,
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
        _ = P(%s)
        }()
''' % ( GoStringLiteral(p.code), p.x.visit(self))
      # TODO:  Check regexp of exception.

    elif p.y is None and type(p.x) == parse.Top and p.x.op in REL_OPS.values():
      # Since message is empty, print LHS, REL_OP, and RHS, since we can.
      a = p.x.a.visit(self)
      b = p.x.b.visit(self)
      sa = Serial('left')
      sb = Serial('right')
      print '   %s, %s := %s, %s' % (sa, sb, a, b)
      print '   if ! (%s.%s(%s)) {' % (sa, p.x.op, sb)
      print '     panic(fmt.Sprintf("Assertion Failed [%s]:  (%%s) ;  left: (%%s) ;  op: %%s ;  right: (%%s) ", %s, %s.Repr(), "%s", %s.Repr() ))' % (
          where, GoStringLiteral(p.code), sa, p.x.op, sb, )
      print '   }'
    else:
      print '   if ! P(%s).Bool() {' % p.x.visit(self)
      print '     panic(fmt.Sprintf("Assertion Failed [%s]:  %%s ;  message=%%s", %s, P(%s).String() ))' % (
          where, GoStringLiteral(p.code), "None" if p.y is None else p.y.visit(self) )
      print '   }'
    print '}'

  def Vtry(self, p):
    serial = Serial('except')
    print '''
   %s_try := func() (%s_z P) {
     defer func() {
       r := recover()
       if r != nil {
         PrintStackUnlessEOF(r)
         %s_z = func() P {
         // BEGIN EXCEPT
''' % (serial, serial, serial)
    # Assign, for the side effect of var creation.
    if p.exvar:
      # OLD -- str
      #Tassign(p.exvar, Traw('MkStr(fmt.Sprintf("%s", r))')).visit(self)
      # NEW -- P
      parse.Tassign(p.exvar, parse.Traw('MkRecovered(r)')).visit(self)

    p.ex.visit(self)

    print '''
           return nil
         // END EXCEPT
         }()
         return
       }
     }()
     // BEGIN TRY
'''
    p.tr.visit(self)

    print '''
     // END TRY
     return nil
   }()
   if %s_try != nil { return %s_try }
''' % (serial, serial)

  def Vlambda(self, p):
    # lvars, lexpr, where
    lamb = Serial('__lambda__')
    ret = parse.Treturn(p.lexpr.xx)
    ret.where, ret.line, ret.gloss = p.where, p.line, 'lambda'
    suite = parse.Tsuite([ret])
    suite.where, suite.line, suite.gloss = p.where, p.line, 'lambda'

    if type(p.lvars) == parse.Titems:
      t = parse.Tdef(lamb, [x.name for x in p.lvars.xx], [None for x in p.lvars.xx], '', '', suite)
    elif type(p.lvars) == parse.Tvar:
      t = parse.Tdef(lamb, [p.lvars.name], [None], '', '', suite)
    else:
      raise Exception("Bad p.lvars type: %s" % type(p.lvars))

    t.where, t.line, t.gloss = p.where, p.line, 'lambda'
    t.visit(self)
    return parse.Tvar(lamb).visit(self)

  def Vforexpr(self, p):
    # Tforexpr(z, vv, ll, cond)
    i = Serial('_')
    ptv = p.ll.visit(self)
    print '''
   forexpr%s := func () P { // around FOR EXPR
     var zz%s []P
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
      print '  if (%s).Bool() {' % p.cond.visit(self)

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

  def Vfor(self, p):
    # var, t, b.
    i = Serial('_')
    ptv = p.t.visit(self)
    print '''
   for_returning%s := func () P { // around FOR
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
     return nil
   }() // around FOR
   if for_returning%s != nil { return for_returning%s }
''' % (i, i)

  # New "with defer"
  def Vwithdefer(self, p):
    # call, body
    var = Serial('with_defer_returning')
    immanentized = self.ImmanentizeCall(p.call, 'defer')
    print '  %s := func() P { defer %s' % (var, immanentized.visit(self))
    p.body.visit(self)
    print '    return nil'
    print '  }()'
    print '  if %s != nil { return %s }' % (var, var)

  def Vglobal(self, p):
    print '  //// GLOBAL: %s' % p ## repr(p.vars.keys())

  def Vswitch(self, p):
    # (self, a, cases, clauses, default_clause):
    serial = Serial('sw')
    self.Gloss(p)
    print '   %s := P(%s)' % (serial, p.a.visit(self))
    print '   _ = %s' % serial
    self.Ungloss(p)
    print '   switch true {'
    for ca, cl in zip(p.cases, p.clauses):
      self.Gloss(ca)
      print '      case %s.EQ(%s): {' % (serial, ca.visit(self))
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
    print '   if %s.Bool() {' % p.t.visit(self)
    p.yes.visit(self)
    if p.no:
      print '   } else {'
      p.no.visit(self)
      print '   }'
    else:
      print '   }'

  def Vwhile(self, p):
    print '   for %s.Bool() {' % p.t.visit(self)
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
    print '   panic( P(%s) )' % p.a.visit(self)

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
    return p.raw

  def Vlit(self, p):
    if p.k == 'N':
      z = p.v
      key = 'litI_' + CleanIdentWithSkids(str(z))
      code = 'MkInt(%s)' % str(z)
    elif p.k == 'F':
      z = p.v
      key = 'litF_' + CleanIdentWithSkids(str(z))
      code = 'MkFloat(%s)' % str(z)
    elif p.k == 'S':
      # TODO --  Don't use eval.  Actually parse it.
      z = parse.DecodeStringLit(p.v)

      key = 'litS_' + CleanIdentWithSkids(z)
      golit = GoStringLiteral(z)
      code = 'MkStr( %s )' % golit
    else:
      Bad('Unknown Vlit', p.k, p.v)
    return self.LitIntern(z, key, code)

  def Vop(self, p):
    if p.returns_bool:
      return ' MkBool(%s.%s(%s)) ' % (p.a.visit(self), p.op, p.b.visit(self))
    if p.b:
      return ' (%s).%s(%s) ' % (p.a.visit(self), p.op, p.b.visit(self))
    else:
      return ' (%s).%s() ' % (p.a.visit(self), p.op)

  def Vboolop(self, p):
    if p.b is None:
      return ' MkBool( %s (%s).Bool()) ' % (p.op, p.a.visit(self))
    else:
      return ' MkBool(%s.Bool() %s %s.Bool()) ' % (p.a.visit(self), p.op, p.b.visit(self))

  def Vcondop(self, p):
    s = Serial('cond')
    print '%s := func (a bool) P { if a { return %s } ; return %s }' % (
        s, p.b.visit(self), p.c.visit(self))
    print '_ = %s' % s  # For some reason, it called us twice?
    return ' %s(%s.Bool()) ' % (s, p.a.visit(self))

  def Vgetitem(self, p):
    return ' (%s).GetItem(%s) ' % (p.a.visit(self), p.x.visit(self))

  def Vgetitemslice(self, p):
    return ' (%s).GetItemSlice(%s, %s, %s) ' % (
        p.a.visit(self),
        'None' if p.x is None else p.x.visit(self),
        'None' if p.y is None else p.y.visit(self),
        'None' if p.z is None else p.z.visit(self))

  def Vcurlysetter(self, p):
    # obj, vec of (var, expr)
    serial = Serial('cs')
    tmp = parse.Tvar(serial)
    parse.Tassign(tmp, p.obj).visit(self)

    for var, x in p.vec:
      self.gsNeeded[var.name] = True
      print '    f_SET_%s(%s, %s)' % (var.name, tmp.visit(self), x.visit(self))
    return tmp.visit(self)

  def Vtuple(self, p):
    return 'MkTupleV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vlist(self, p):
    return 'MkListV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vdict(self, p):
    return 'MkDictV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

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
    if p.name in BUILTINS:
      return Zbuiltin(p, 'G_%s' % p.name)
    return Zglobal(p, 'G_%s' % p.name)

  def ImmanentizeCall(self, p, why):
    "Eval all args of Tcall now and return new Tcall, for defer or go."
    s = Serial(why)
    print '%s_fn := P( %s )' % (s, p.fn.visit(self))
    n = len(p.args)
    i = 0
    for a in p.args:
      print '%s_a%d := P( %s )' % (s, i, a.visit(self))
      i += 1

    if p.star:
      print '%s_star := P( %s )' % (s, p.star.visit(self))
    if p.starstar:
      print '%s_starstar := P( %s )' % (s, p.starstar.visit(self))

    return parse.Tcall(
        parse.Traw('%s_fn' % s),
        [parse.Traw('%s_a%d' % (s,i)) for i in range(n)],
        p.names,
        parse.Traw('%s_star' % s) if p.star else p.star,
        parse.Traw('%s_starstar' % s) if p.starstar else p.starstar,
    )

  def Vgo(self, p):
    immanentized = self.ImmanentizeCall(p.fcall, 'gox')
    return 'MkPromise(func () P { return %s })' % immanentized.visit(self)

  def Vcall(self, p):
    # fn, args, names, star, starstar
    global MaxNumCallArgs
    n = len(p.args)
    MaxNumCallArgs = max(MaxNumCallArgs, n)

    arglist = ', '.join(["%s" % (a.visit(self)) for a in p.args])

    #print '// Vcall: fn:', repr(p.fn)
    #print '// Vcall: args:', repr(p.args)
    #print '// Vcall: names:', repr(p.names)
    #print '// Vcall: star:', repr(p.star)
    #print '// Vcall: starstar:', repr(p.starstar)
    if p.star or p.starstar or any(p.names):
      return 'P(%s).(ICallV).CallV([]P{%s}, %s, []KV{%s}, %s) ' % (

          p.fn.visit(self),

          ', '.join([str(p.args[i].visit(self)) for i in range(len(p.args)) if not p.names[i]]),  # fixed args with no names.

          ('(%s).List()' % p.star.visit(self)) if p.star else 'nil',

          ', '.join(['KV{"%s", %s}' % (p.names[i], p.args[i].visit(self)) for i in range(len(p.args)) if p.names[i]]),  # named args.

          ('(%s).Dict()' % p.starstar.visit(self)) if p.starstar else 'nil',
      )

    if type(p.fn) is parse.Tfield:
      if type(p.fn.p) is parse.Tvar:

        if p.fn.p.name == 'super':
          return 'self.%s.M_%d_%s(%s)' % (self.tailSup(self.sup), n, p.fn.field, arglist)

        if p.fn.p.name in self.imports:
          imp = self.imports[p.fn.p.name]
          print '//', p.fn.p.name, imp, imp.imported, p.fn.field

          if imp.imported == ['github.com', 'strickyak', 'rye', 'pye', 'sys'] and p.fn.field == 'exc_info':

            serial = Serial('exc_info')
            print '%s0 := fmt.Sprintf("%%s", r)' % serial
            print '%s1 := fmt.Sprintf("%%v", r)' % serial
            print '%s2 := make([]byte, 9999)' % serial
            print '%s2len := runtime.Stack(%s2, false)' % (serial, serial)
            return 'MkList([]P{ MkStr(%s0), MkStr(%s1), MkStr(string(%s2[:%s2len]))})' % (serial, serial, serial, serial)

          if imp.imported[0] == 'go':
            return 'MkGo(i_%s.%s).Call(%s) ' % (p.fn.p.name, p.fn.field, arglist)
          else:
            return 'call_%d( i_%s.G_%s, %s) ' % (n, p.fn.p.name, p.fn.field, arglist)

      # General Method Invocation.
      key = '%d_%s' % (n, p.fn.field)
      self.invokes[key] = (n, p.fn.field)
      return '/**/ f_INVOKE_%d_%s(%s, %s) ' % (n, p.fn.field, p.fn.p.visit(self), arglist)

    def NativeGoTypeName(a):
        if type(a) is parse.Tfield:
          return '%s.%s' % (a.p.visit(self), a.field)
        elif type(a) is parse.Tvar:
          return a.name
        else:
          raise Exception('Strange thing for go_type: ' + a)

    zfn = p.fn.visit(self)
    if type(zfn) is Zbuiltin:
      if p.fn.name == 'go_type':
        assert len(p.args) == 1, 'go_type got %d args, wants 1' % len(p.args)
        return 'GoElemType(new(%s))' % NativeGoTypeName(p.args[0])
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
      else:
        raise Exception('Undefind builtin: %s' % p.fn.name)

    if type(zfn) is Zglobal and zfn.t.name in self.defs:
      fp = self.defs[zfn.t.name]
      if not fp.star and not fp.starstar:
        want = len(fp.args)
        if n != want:
          raise Exception('Calling global function "%s", got %d args, wanted %d args' % (zfn.t.name, n, want))
        return 'G_%d_%s(%s) ' % (n, zfn.t.name, arglist)

    if type(zfn) is Zsuper:  # for calling super-constructor.
      return 'self.%s.M_%d___init__(%s) ' % (self.tailSup(self.sup), n, arglist)

    return 'call_%d( P(%s), %s )' % (n, p.fn.visit(self), arglist)

  def Vfield(self, p):
    # p, field
    x = p.p.visit(self)
    if type(x) is Zsuper:
      raise Exception('Special syntax "super" not used with Function Call syntax')
    if type(x) is Zself and not self.cls:
      raise Exception('Using a self field but not in a class definition: field="%s"' % p.field)
    if type(x) is Zself and self.instvars.get(p.field):  # Special optimization for self instvars.
      return '%s.M_%s' % (x, p.field)
    elif type(x) is Zimport:
      if x.imp.imported[0] == 'go':
        return ' MkGo(%s.%s) ' % (x, p.field)
      else:
        return ' %s.G_%s ' % (x, p.field)
    else:
      self.gsNeeded[p.field] = True
      return ' f_GET_%s(P(%s)) ' % (p.field, x)

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
    # name, args, dflts, star, starstar, body.

    # SAVE STATUS BEFORE THIS FUNC.
    save_func = self.func
    save_yields = self.yields
    save_force_globals = self.force_globals
    self.func = p
    self.func_level += 1

    # START A PRINT BUFFER -- but not if Nested.
    nesting = None
    if self.func_level >= 2:
      nesting = Serial('nesting')
    else:
      buf = PushPrint()

    # LOOK AHEAD for "yield" and "global" statements.
    finder = parse.YieldAndGlobalFinder()
    finder.Vsuite(p.body)
    self.yields = finder.yields
    self.force_globals = finder.force_globals

    # Tweak args.  Record meth, if meth.
    args = p.args  # Will drop the initial 'self' element, if in a cls.
    dflts = p.dflts  # Will drop the initial 'self' element, if in a cls.
    if nesting:
      pass
    elif self.cls and not nesting:
      if len(p.args) > 0 and p.args[0] == 'self':  # User may omit self.
        args = p.args[1:]  # Skip self; it is assumed.
        dflts = p.dflts[1:]  # Skip self; it is assumed.
      self.meths[p.name] = ArgDesc(self.modname, self.cls.name, '%s.%s::%s' % (self.modname, self.cls.name, p.name), args, dflts, p.star, p.starstar)
    else:
      self.defs[p.name] = ArgDesc(self.modname, None, '%s.%s' % (self.modname, p.name), args, p.dflts, p.star, p.starstar)

    # Copy scoe and add argsPlus to the new one.
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
           mustBeNone := func() P {
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
        return gen
'''

    PopPrint()
    code2 = str(buf2)

    print '///////////////////////////////'
    print ''

    letterV = 'V' if p.star or p.starstar else ''
    emptiesV = (', MkList(nil), MkDict(nil)' if args else 'MkList(nil), MkDict(nil)') if p.star or p.starstar else ''
    stars = ' %s P, %s P' % (AOrSkid(p.star), AOrSkid(p.starstar)) if p.star or p.starstar else ''

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

    print ' %s(%s %s) P {' % (func_head, ' '.join(['a_%s P,' % a for a in args]), stars)
    print '  FuncCounter["%s"]++' % func_key

    for v, v2 in sorted(self.scope.items()):
      if save_scope is None or v not in save_scope:
        if v2[0] != 'a':  # Skip args
          print "   var %s P = None; _ = %s" % (v2, v2)
    print code2

    print '   return None'
    print ' }'
    print ''

    # Restore the old scope.
    self.scope = save_scope

    n = len(args)
    argnames = ', '.join(['"%s"' % a for a in p.args])
    defaults = ', '.join([(str(d.visit(self)) if d else 'nil') for d in p.dflts])

    if nesting:
      fn_var = parse.Tvar(p.name)

      tmp = '''
        &pNest_%s{PCallable: PCallable{
                    Name: "%s__%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"},
                    fn: fn_%s}''' % (
          nesting, p.name, nesting, argnames, defaults, p.star, p.starstar, nesting)

      self.AssignAFromB(fn_var, parse.Traw(tmp), None)


    # Now for the Nested case, START A PRINT BUFFER.
      buf = PushPrint()

    print '///////////////////////////////'
    print '// name:', p.name
    print '// args:', p.args
    print '// dflts:', p.dflts
    print '// star:', p.star
    print '// starstar:', p.starstar
    print '// body:', p.body
    try:
      sys.stdout.flush()
    except:
      print >> sys.stderr, '((( cannot flush )))'

    if nesting:
      print ' type pNest_%s struct { PCallable; fn func(%s %s) P }' % (nesting, ' '.join(['a_%s P,' % a for a in args]), stars)
      print ' func (o *pNest_%s) Contents() interface{} {' % nesting
      print '   return o.fn'
      print ' }'
      if p.star or p.starstar:
        pass  # No direct pNest method; use CallV().
      else:
        print ' func (o pNest_%s) Call%d(%s) P {' % (nesting, n, ', '.join(['a%d P' % i for i in range(n)]))
        print '   return o.fn(%s)' % (', '.join(['a%d' % i for i in range(n)]))
        print ' }'
      print ''
      print ' func (o pNest_%s) CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P {' % nesting
      print '   argv, star, starstar := SpecCall(&o.PCallable, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   return o.fn(%s star, starstar)' % (' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return o.fn(%s)' % (', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

    elif self.cls:
      print ' type pMeth_%d_%s__%s struct { PCallable; Rcvr *%s }' % (n, self.cls.name, p.name, gocls)
      print ' func (o *pMeth_%d_%s__%s) Contents() interface{} {' % (n, self.cls.name, p.name)
      print '   return o.Rcvr.M_%d%s_%s' % (n, letterV, p.name)
      print ' }'
      print ' func (o *pMeth_%d_%s__%s) Call%d(%s) P {' % (n, self.cls.name, p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '   return o.Rcvr.M_%d%s_%s(%s%s)' % (n, letterV, p.name, ', '.join(['a%d' % i for i in range(n)]), emptiesV)
      print ' }'
      print ''
      print ' func (o *pMeth_%d_%s__%s) CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P {' % (n, self.cls.name, p.name)
      print '   argv, star, starstar := SpecCall(&o.PCallable, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   return o.Rcvr.M_%dV_%s(%s star, starstar)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

    else:
      print ' type pFunc_%s struct { PCallable }' % p.name
      print ' func (o *pFunc_%s) Contents() interface{} {' % p.name
      print '   return G_%s' % p.name
      print ' }'
      if p.star or p.starstar:
        pass  # No direct pFunc method; use CallV().
      else:
        print ' func (o pFunc_%s) Call%d(%s) P {' % (p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
        print '   return G_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
        print ' }'
      print ''
      print ' func (o pFunc_%s) CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P {' % p.name
      print '   argv, star, starstar := SpecCall(&o.PCallable, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      # TODO: I think this is old, before named params.
      if p.star or p.starstar:  # If either, we always pass both.
        print '   return G_%dV_%s(%s star, starstar)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return G_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

      self.glbls[p.name] = ('*pFunc_%s' % p.name,
                            '&pFunc_%s{PCallable: PCallable{Name: "%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}}' % (
                                p.name, p.name, argnames, defaults, p.star, p.starstar))

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
       '\n'.join(['   M_%s   P' % x for x in self.instvars]),
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

    # Make GET and SET interfaces for each instance var and each method.
    print ''
    for iv in self.instvars.keys() + self.meths.keys():
      self.gsNeeded[iv] = True

    # For all the instance vars
    print ''
    for iv in sorted(self.instvars):
      print ' func (o *C_%s) GET_%s() P { return o.M_%s }' % (p.name, iv, iv)
      print ' func (o *C_%s) SET_%s(x P) { o.M_%s = x }' % (p.name, iv, iv)
      print ''
    print ''

    # For all the methods
    print ''
    for m in sorted(self.meths):  # ArgDesc in self.meths[m]
      args = self.meths[m].args
      dflts = self.meths[m].dflts
      n = len(args)

      argnames = ', '.join(['"%s"' % a for a in args])
      defaults = ', '.join([(str(d.visit(self)) if d else 'nil') for d in dflts])

      spec = 'PCallable: PCallable{Name: "%s::%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}' % (p.name, m, argnames, defaults, self.meths[m].star, self.meths[m].starstar)

      print ' func (o *%s) GET_%s() P { z := &pMeth_%d_%s__%s { %s, Rcvr: o }; z.SetSelf(z); return z }' % (gocls, m, n, p.name, m, spec)

    # Special methods for classes.
    if self.sup != 'native':
      print 'func (o *C_%s) Rye_ClearFields__() {' % p.name
      for iv in self.instvars:
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
      print 'func (o *C_%s) PType() P { return G_%s }' % (p.name, p.name)
      print 'func (o *pFunc_%s) Repr() string { return "%s" }' % (p.name, p.name)
      print 'func (o *pFunc_%s) String() string { return "<class %s>" }' % (p.name, p.name)
      print ''


    self.tail.append(str(buf))
    PopPrint()

  def Vsuite(self, p):
    for x in p.things:
      print '// @ %d @ %d @ %s' % (x.where, x.line, x.gloss)
      x.visit(self)
      print '// $ %d $ %d $ %s' % (x.where, x.line, x.gloss)

PrinterStack= []
def PushPrint():
    global PrinterStack

    try:
      sys.stdout.flush()
    except:
      print >> sys.stderr, '((( cannot flush )))'

    PrinterStack.append(sys.stdout)
    buf = Buffer()
    sys.stdout = buf
    return buf
def PopPrint():
    global PrinterStack
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

class Z(object):  # Returns from visits (emulated runtime value).
  def __init__(self, t, s):
    self.t = t  # T node
    self.s = s  # String for backwards compat
  def __str__(self):
    return self.s
class Zself(Z):
  pass
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
class Zbuiltin(Z):
  pass
class Zlit(Z):
  pass

pass

class ArgDesc(object):
  def __init__(self, module, cls, name, args, dflts, star, starstar):
    self.module = module
    self.cls = cls
    self.name = name
    self.args = args
    self.dflts = dflts
    self.star = star
    self.starstar = starstar

  def CallSpec(self):
    argnames = ', '.join(['"%s"' % a for a in self.args])
    defaults = ', '.join([(str(d.visit(self)) if d else 'nil') for d in self.dflts])
    return 'PCallable{Name: "%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}' % (self.name, argnames, defaults, self.star, self.starstar)

def AOrSkid(s):
  if s:
    return 'a_%s' % s
  else:
    return '_'

