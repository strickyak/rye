import codecs
import os
import re
import sys

RYE_FLOW = os.getenv('RYE_FLOW')

# TODO: move 'unpickle pickle goreify goderef gocast gotype gonew' into 'rye' space.   Also byt?
# 'unpickle pickle goreify goderef gocast gotype gonew len repr str int bool float list dict tuple range sorted type byt'
BUILTINS = set(
    'rye_unpickle rye_pickle unpickle pickle go_reify go_deref go_cast go_type go_new'
    .split())

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([ \t\n]*[#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')
RE_PRAGMA = re.compile('[ \t]*[#][#][A-Za-z:()]+')

RE_KEYWORDS = re.compile(
    '\\b(del|say|from|class|def|native|if|elif|else|while|True|False|None|print|and|or|try|except|raise|yield|return|break|continue|pass|as|go|defer|with|global|assert|must)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[-]=|[*]=|/=|//|<<|>>>|>>|==|!=|<=|>=|[*][*]|[.][.]')
RE_OPS = re.compile('[-.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_FLOAT = re.compile('[+-]?[0-9]+[.][-+0-9eE]*')
RE_INT = re.compile('[+-]?[0-9]+')
RE_STR3 = re.compile('((?s)"""(.*?)"""|\'\'\'(.*?)\'\'\')')
RE_STR = re.compile('(["](([^"\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\n]|[\\\\].)*)[\'])')

RE_WORDY_REL_OP = re.compile('\\b(not\\s+in|is\\s+not|in|is)\\b')
RE_NOT_IN = re.compile('^not\\s+in$')
RE_IS_NOT = re.compile('^is\\s*not$')

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
  [RE_STR, 'S'],
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
  return '"' + TROUBLE_CHAR.sub((lambda m: '\\x%02x' % ord(m.group(0))), s) + '"'

def CleanIdentWithSkids(s):
  return NONALFA.sub((lambda m: '_%02x' % ord(m.group(0))), s)

def CleanQuote(x):
  return re.sub('[^A-Za-z0-9_]', '~', x)[:10]
  return '~~~'
  return re.sub('[^A-Za-z0-9_]', '~', x)
  return re.sub('[`]', '\'', x)
def Bad(format, *args):
  raise Exception(format % args)

def SimplifyContinuedLines(tokens):
  # Try to throw no exceptions from this func.
  w = []  # Waiting.
  deep = 0   #  Grouping depth.
  eat_out = 0  # How many OUT marks to ignore.
  for triple in tokens:
    kind, val, pos = triple
    if kind == 'G':
      if val in ['(', '[', '{']:
        deep += 1
      elif val in ['}', ']', ')']:
        deep -= 1

    if eat_out:
      if kind != 'OUT':
        raise Exception('Expected un-indent at position ', pos)
      eat_out -= 1
    elif w or deep:
      #print >> sys.stderr, 'w.append(triple)', triple, w
      w.append(triple)
    else:
      #print >> sys.stderr, 'yield triple', triple
      yield triple

    if w and not deep and val == ';;':
      # Try to throw no exceptions from here.
      #if eat_out:
      #  raise Exception('eat_out:' + eat_out)
      for w_triple in w:
        w_kind, _, _ = w_triple

        if w_kind == 'IN':
          eat_out += 1
        elif w_kind == 'OUT':
          eat_out -= 1
        elif w_kind == ';;':
          pass
        else:
          #print >> sys.stderr, 'yield w_triple from w', w_triple, w
          yield w_triple
      w = []
      yield triple  # The newline.
  pass

class Lex(object):
  def __init__(self, program):
    self.buf = program
    self.i = 0
    self.indents = [1]
    self.tokens = []
    n = len(self.buf)
    while self.i < n:
      self.DoWhite()
      if self.i < n:
        self.DoBlack()

  def Add(self, x):
    self.tokens.append(x)

  def DoBlack(self):
    rest = self.buf[self.i:]
    for reg, kind in DETECTERS:
      m = reg.match(rest)
      if m:
        got = m.group(0)
        self.Add((kind, got, self.i))
        self.i += len(got)
        return
    raise Bad("Cannot parse (at %d): %s", self.i, repr(rest))

  def DoWhite(self):
    # pragma looks like a comment, but is considered Black.
    if RE_PRAGMA.match(self.buf[self.i:]):
      return

    m = RE_WHITE.match(self.buf[self.i:])
    # blank_lines includes all the newlines;
    #   if blank_lines is empty, we're not on a new line.
    # white is the remnant at the front of partial line;
    #   white is the new indentation level.
    blank_lines, _, white = m.groups()
    # both is the entire match.
    both = m.group(0)
    i = self.i
    self.i += len(both)

    if not blank_lines:
      return # White space is inconsequential if not after \n

    self.Add((';;', ';;', i))

    col = 1 + TabWidth(white)  # Conventionally, columns start at 1.
    if col < self.indents[-1]:
      # outdent (i.e. un-indent).
      j = len(self.indents) - 1  # For iterating backwards thru indents.
      outage = both  # For recording the white space.
      while col < self.indents[j]:
        # Not back far enough yet.
        self.Add(('OUT', outage, i))
        outage = ''  # We put all the white space in the first OUT.
        j -= 1
        if j < 0 or self.indents[j] < col:
          raise Bad('Cannot un-indent: New column is %d; previous columns are %s', col, repr(self.indents))
        if self.indents[j] == col:
            self.indents[j+1:] = []  # Trim tail to index j.

    elif col > self.indents[-1]:
        # indent
        self.Add(('IN', both, i))
        self.indents.append(col)

def TabWidth(s):
  z = 0
  for c in s:
    if c == '\t':
      z = ((z+TAB_WIDTH-1) // TAB_WIDTH) * TAB_WIDTH
    else:
      z += 1
  return z

SerialNum = 10
def Serial(s):
  global SerialNum
  SerialNum += 1
  return '%s_%d' % (s, SerialNum)

class CodeGen(object):
  def __init__(self, up):
    self.up = up
    self.glbls = {}         # name -> (type, initialValue)
    self.imports = {}       # name -> Vimport
    self.defs = {}          # name -> ArgsDesc()
    self.lits = {}          # key -> name
    self.invokes = {}       # key -> (n, fieldname)
    self.scopes = []
    self.force_globals = {}
    self.tail = []
    self.cls = ''
    self.gsNeeded = {}      # keys are getter/setter names.

  def InjectForInternal(self, stuff):
    self.invokes, self.defs, self.gsNeeded = stuff

  def ExtractForInternal(self):
    stuff = self.invokes, self.defs, self.gsNeeded
    return stuff

  def GenModule(self, modname, path, suite, cwp=None, main=None, internal=""):
    self.cwp = cwp
    self.path = path
    self.modname = modname
    self.internal = internal
    if modname is None:
      print ' package main'
      print ' import "runtime/pprof"'
      print ' import . "github.com/strickyak/rye"'
    elif internal:
      print ' package rye'
    else:
      print ' package %s' % os.path.basename(modname)
      print ' import . "github.com/strickyak/rye"'
    print ' import "fmt"'
    print ' import "os"'
    print ' import "reflect"'

    # Look for main
    main_def = None
    for th in suite.things:
      if type(th) == Tdef:
        if th.name == 'main':
          main_def = th
    # Add a main, if there isn't one.
    if not main_def and not internal:
      main_def = Tdef('main', ['argv'], [None], None, None, Tsuite([]))
      suite.things.append(main_def)

    for th in suite.things:
      if type(th) == Timport:
          if th.fromWhere == '.':
            th.imported = self.cwp.split('/') + th.imported
            th.fromWhere = None
          print ' import i_%s "%s"' % (th.alias, '/'.join(th.imported))

    for th in suite.things:
      if type(th) == Tclass and th.sup != "native":
        # name, sup, things
        # Create constructor functions with Tdef().

        # newctor:
        c_init = None
        for thth in list(th.things):
          if type(thth) is Tdef and thth.name == '__init__':
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

        x_star = Tvar(c_star) if c_star else None
        x_starstar = Tvar(c_starstar) if c_starstar else None

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
        c1 = Tnative(natives)

        c2 = Tassign(Tvar('rye_result__'), Traw('z'))

        # Tcall: fn, args, names, star, starstar
        call = Tcall(Tfield(Tvar('rye_result__'), '__init__'), [Tvar(a) for a in c_args], c_args, x_star, x_starstar)
        c3 = Tassign(Traw('_'), call)

        c4 = Treturn([Tvar('rye_result__')])

        seq = Tseq([c1, c2, c3, c4])
        ctor = Tdef(c_name, c_args, c_dflts, c_star, c_starstar, seq)
        for t in [c1, c2, c3, c4, seq, ctor]:
          t.where = th.where
          t.gloss = 'ctor'

        suite.things.append(ctor)

    # Avoid golang's "import not used" errors in corner cases.
    print ' var _ = fmt.Sprintf'
    print ' var _ = os.Stderr'
    print ' var _ = reflect.ValueOf'
    print ' var _ = MkInt'  # From rye runtime.
    print ''

    if self.internal:
      print ' func eval_module_internal_%s () P {' % self.internal
    else:
      print ' var eval_module_once P'
      print ' func Eval_Module () P {'
      print '   if (eval_module_once == nil) {'
      print '     eval_module_once = inner_eval_module()'
      print '   }'
      print '   return eval_module_once'
      print ' }'
      print ' func inner_eval_module () P {'

    for th in suite.things:
      try:
        print '// @ %d @ %s' % (th.where, th.gloss)
      except:
        pass
      th.visit(self)
      try:
        print '// $ %d $ %s' % (th.where, th.gloss)
      except:
        pass
    print '   return None'
    print ' }'
    print ''
    print '\n\n'.join(self.tail)
    print ''

    for g, (t, v) in sorted(self.glbls.items()):
      print 'var G_%s P // %s' % (g, t)
    print ''
    print ' func init /*New_Module*/ () {'
    for g, (t, v) in sorted(self.glbls.items()):
      print '   G_%s = %s' % (g, v)
      if v != "None":
        print '   G_%s.SetSelf(G_%s)' % (g, g)

        if self.internal:
          print '   Globals["%s.%s"] = G_%s' % (self.modname, g, g)
        else:
          print '   Globals["%s/%s.%s"] = G_%s' % (self.cwp, self.modname, g, g)
    print ' }'
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
      print '    return x.M_%d_%s(%s)         ' % (n, fieldname, args)    ###### ZZZZZZZZZZZZz added %d
      print '  case i_GET_%s:         ' % fieldname
      print '    tmp := x.GET_%s()    ' % fieldname
      print '    return call_%d(tmp, %s)' % (n, ', '.join(['a_%d' % j for j in range(n)]))
      print '    '

      print '  case *PGo:                '
      print '    return x.Invoke("%s", %s) ' % (fieldname, args)
      print '  }'
      print '  panic(fmt.Sprintf("Cannot invoke \'%s\' with %d arguments on %%v", fn))' % (fieldname, n)
      print '}'
      print 'type i_INVOKE_%d_%s interface { M_%d_%s(%s) P }' % (n, fieldname, n, fieldname, formals) ########## ZZZZZZZZZZZZZz added %d inside after M
    print ''

    for iv in sorted(self.gsNeeded):
      print ' type i_GET_%s interface { GET_%s() P }' % (iv, iv)
      print ' type i_SET_%s interface { SET_%s(P) }' % (iv, iv)

      print 'func f_GET_%s(h P) P {' % iv
      print '  switch x := h.(type) { '
      print '  case i_GET_%s:         ' % iv
      print '    return x.GET_%s()    ' % iv
      print '  case *PGo:             '
      print '    return FetchFieldByName(x.V, "%s") ' % iv
      print '  }'
      print '  panic(fmt.Sprintf("Cannot GET \'%s\' on %%T: %%v", h, h))' % iv
      print '}'
      print ''

      print 'func f_SET_%s(h P, a P) {' % iv
      print '  switch x := h.(type) { '
      print '  case i_SET_%s:         ' % iv
      print '    x.SET_%s(a)    ' % iv
      print '    return'
      print '  case *PGo:             '
      print '    StoreFieldByName(x.V, "%s", a)' % iv
      print '    return'
      print '  }'
      print '  panic(fmt.Sprintf("Cannot SET \'%s\' on %%T: %%v", h, h))' % iv
      print '}'
      print ''
    print ''

    #if self.internal:
    #  return
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


    if main:
      sys.stdout.close()
      sys.stdout = main
      print '''
 package main
 import "os"
 import "runtime/pprof"
 import "github.com/strickyak/rye"
 import MY "%s"
 func main() {
        f, err := os.Create("zzz.cpu")
        if err != nil {
            panic(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()

        MY.Eval_Module()
 }
''' % modname
      sys.stdout.close()

    elif not path:
      print '''
 func main() {
        f, err := os.Create("zzz.cpu")
        if err != nil {
            panic(err)
        }
        pprof.StartCPUProfile(f)
        defer pprof.StopCPUProfile()

        Eval_Module()
 }
'''

  def Vexpr(self, p):
    print ' _ = %s' % p.a.visit(self)

  def AssignFieldAFromRhs(self, a, rhs, pragma):
      lhs = a.p.visit(self)
      if type(lhs) is Zself:  # Special optimization for self.
        self.instvars[a.field] = True
        lhs = 'self.M_%s /*Apragma:%s*/' % (a.field, pragma)
        print '   %s /*ZFpragma:%s*/ = %s' % (lhs, pragma, rhs)
      else:
        self.gsNeeded[a.field] = True
        print '   f_SET_%s(%s, %s)' % (a.field, lhs, rhs)

  def AssignItemAFromRhs(self, a, rhs, pragma):
        p = a.a.visit(self)
        q = a.x.visit(self)
        print '   (%s).SetItem(%s, %s)' % (p, q, rhs)

  def AssignTupleAFromB(self, a, b, pragma):
        serial = Serial('detuple')
        tmp = Tvar(serial)
        Tassign(tmp, b).visit(self)

        ### print '   tuple_%s, ok_%s := %s.(*PTuple)' % (serial, serial, tmp.visit(self))
        ### print '   if ! ok_%s { panic("Not a tuple, in destructuring assignment: " + (%s.Show())) }' % (serial, tmp.visit(self))

        print '   len_%s := %s.Len()' % (serial, tmp.visit(self))
        print '   if len_%s != %d { panic(fmt.Sprintf("Assigning object of length %%d to %%d variables, in destructuring assignment.", len_%s, %d)) }' % (serial, len(a.xx), serial, len(a.xx))

        i = 0
        for x in a.xx:
          if type(x) is Tvar and x.name == '_':
            pass # Tvar named '_' is the bit bucket;  don't Tassign.
          else:
            Tassign(x, Tgetitem(tmp, Tlit('N', i))).visit(self)
            # Tassign(x, Traw('tuple_%s.PP[%d]' % (serial, i))).visit(self)
          i += 1

  def AssignAFromB(self, a, b, pragma):
    # Resolve rhs first.
    type_a = type(a)

    if type_a is Titems or type_a is Ttuple:
      return self.AssignTupleAFromB(a, b, pragma)

    lhs = '?lhs?'
    rhs = b.visit(self)

    if type_a is Tfield:
      return self.AssignFieldAFromRhs(a, rhs, pragma)

    elif type_a is Tgetitem:  # p[q] = rhs
      return self.AssignItemAFromRhs(a, rhs, pragma)

    elif type_a is Tvar:
      # Are we in a function scope?
      if len(self.scopes) and a.name not in self.force_globals:
        # Inside a function.
        scope = self.scopes[0]
        if scope.get(a.name):
          lhs = scope[a.name]
        else:
          lhs = scope[a.name] = 'v_%s /*Bpragma:%s*/' % (a.name, pragma)
      else:
        # At the module level.
        lhs = a.visit(self)
        self.glbls[a.name] = ('P', 'None')

    elif type_a is Traw:
      lhs = a.raw

    else:
      raise Exception('Weird Assignment, a class is %s, A IS (%s) (%s) B IS (%s) (%s)' % (type(a).__name__, a, a.visit(self), b, b.visit(self)))

    print '   %s /*Zpragma:%s*/ = %s' % (lhs, pragma, rhs)

  def Vassign(self, p):
    # a, b, pragma
    self.AssignAFromB(p.a, p.b, p.pragma)

  def Vprint(self, p):
    vv = [a.visit(self) for a in p.xx.xx]
    if p.saying:
      print '   fmt.Fprintln(os.Stderr, "# %s # ", %s.Repr())' % (
          codecs.encode(p.code, 'unicode_escape').replace('"', '\\"'),
                '.Repr(), "#", '.join([str(v) for v in vv]))
    else:
      print '   fmt.Println(%s.String())' % '.String(), '.join([str(v) for v in vv])

  def Vimport(self, p):
    ##im = '/'.join(p.imported)
    if self.glbls.get(p.alias):
      raise Exception("Import alias %s already used", p.alias)
    print '//## p.fromWhere:', p.fromWhere
    if p.fromWhere == '.':
      p.imported = self.cwp.split('/') + p.imported
      p.fromWhere = None
    print '//## p:', p.fromWhere, p.imported, p.alias

    self.imports[p.alias] = p

    if not p.fromWhere:
      # Modules already contain protections against evaling more than once.
      print '   i_%s.Eval_Module() ' % p.alias

  def Vassert(self, p):
    is_must = p.is_must
    # TODO: A way to skip 'assert' but still execute 'must'.
    if p.y is None and type(p.x) == Top and p.x.op in REL_OPS.values():
      # Since message is empty, print LHS, REL_OP, and RHS, since we can.
      a = p.x.a.visit(self)
      b = p.x.b.visit(self)
      sa = Serial('left')
      sb = Serial('right')
      print '   %s, %s := %s, %s' % (sa, sb, a, b)
      print '   if ! (%s.%s(%s)) {' % (sa, p.x.op, sb)
      print '     panic(fmt.Sprintf("Assertion Failed:  (%%s) ;  left: (%%s) ;  op: %%s ;  right: (%%s) ", %s, %s.Repr(), "%s", %s.Repr() ))' % (
          GoStringLiteral(p.code), sa, p.x.op, sb, )
      print '   }'
    else:
      print '   if ! P(%s).Bool() {' % p.x.visit(self)
      print '     panic(fmt.Sprintf("Assertion Failed:  %%s ;  message=%%s", %s, P(%s).String() ))' % (
          GoStringLiteral(p.code), "None" if p.y is None else p.y.visit(self) )
      print '   }'

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
      Tassign(p.exvar, Traw('MkStr(fmt.Sprintf("%s", r))')).visit(self)

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

  def Vforexpr(self, p):
    # Tforexpr(z, vv, ll, cond, has_comma)
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

    Tassign(p.vv, Traw("ndx_%s" % i)).visit(self)

    if p.cond:
      print '  if (%s).Bool() {' % p.cond.visit(self)

    if len(p.z) == 1 and not p.has_comma:
      # scalar case.
      print '  zz%s = append(zz%s, %s)' % (i, i, p.z[0].visit(self))
    else:
      print '  zz%s = append(zz%s, %s)' % (i, i, Ttuple(p.z).visit(self))

    if p.cond:
      print '  }'

    print '''
       // END FOR EXPR
     }
     return MkList(zz%s)
   }() // around FOR EXPR
''' % i
    return 'forexpr%s' % i

  def Vfor(self, p):
    i = Serial('_')
    ptv = p.t.visit(self)
    print '''
   maybe%s := func () P { // around FOR
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

    Tassign(p.var, Traw("ndx_%s" % i)).visit(self)

    p.b.visit(self)

    print '''
       // END FOR
     }
     return nil
   }() // around FOR
   if maybe%s != nil { return maybe%s }
''' % (i, i)

  # New "with defer"
  def Vwithdefer(self, p):
    # call, body
    var = Serial('withdefer')
    immanentized = self.ImmanentizeCall(p.call, 'defer')
    print '  %s := func() P { defer %s' % (var, immanentized.visit(self))
    p.body.visit(self)
    print '    return nil'
    print '  }()'
    print '  if %s != nil { return %s }' % (var, var)

  # Old "defer"
  def Vdefer(self, p):
    # Note, p.cmd is a Tassign, and lhs is '_'
    immanentized = self.ImmanentizeCall(p.cmd.b, 'defer')
    print 'defer', immanentized.visit(self)

  def Vglobal(self, p):
    print '  //// GLOBAL: %s' % repr(p.vars.keys())

  def Vif(self, p):
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
    if type(p.listx) == Titems:
      for e in p.listx.items.xx:
        self.Vdel(e)

    elif type(p.listx) == Tgetitem:
      print "%s.DelItem(%s)" % (p.listx.a.visit(self), p.listx.x.visit(self))

    elif type(p.listx) == Tgetitemslice:
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

  def Vlit(self, p):
    if p.k == 'N':
      v = p.v
      key = 'litI_' + CleanIdentWithSkids(str(v))
      code = 'MkInt(%s)' % v
    elif p.k == 'F':
      v = p.v
      key = 'litF_' + CleanIdentWithSkids(str(v))
      code = 'MkFloat(%s)' % v
    elif p.k == 'S':
      # TODO --  Don't use eval.  Actually parse it.
      v = p.v
      v = v.replace('\n', '\\n')
      try:
        v = eval(v)
      except:
        raise "Sorry, rye currently cannot handle this string literal: " + repr(v)
      key = 'litS_' + CleanIdentWithSkids(v)
      golit = GoStringLiteral(v)
      code = 'MkStr( %s )' % golit
    else:
      Bad('Unknown Vlit', p.k, p.v)
    return self.LitIntern(v, key, code)

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
    return ' %s(%s.Bool()) ' % (s, p.a.visit(self))

  def Vgetitem(self, p):
    return ' (%s).GetItem(%s) ' % (p.a.visit(self), p.x.visit(self))

  def Vgetitemslice(self, p):
    return ' (%s).GetItemSlice(%s, %s, %s) ' % (
        p.a.visit(self),
        None if p.x is None else p.x.visit(self),
        None if p.y is None else p.y.visit(self),
        None if p.z is None else p.z.visit(self))

  def Vcurlysetter(self, p):
    # obj, vec of (var, expr)
    serial = Serial('cs')
    tmp = Tvar(serial)
    Tassign(tmp, p.obj).visit(self)

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
    if p.name == 'rye':
      return Zrye(p, 'rye')
    if p.name == 'self':
      return Zself(p, 'self')
    if p.name == 'super':
      return Zsuper(p, 'super')
    if p.name in self.force_globals:
      return Zglobal(p, '/*force_globals*/G_%s' % p.name)
    if p.name in self.imports:
      return Zimport(p, 'i_%s' % p.name, self.imports[p.name])
    for s in self.scopes:
      if p.name in s:
        return Zlocal(p, s[p.name])
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

    return Tcall(
        Traw('%s_fn' % s),
        [Traw('%s_a%d' % (s,i)) for i in range(n)],
        p.names,
        Traw('%s_star' % s) if p.star else p.star,
        Traw('%s_starstar' % s) if p.starstar else p.starstar,
    )

  def Vgox(self, p):
    immanentized = self.ImmanentizeCall(p.fcall, 'gox')
    return 'MkPromise(func () P { return %s })' % immanentized.visit(self)

  def Vcall(self, p):
    # fn, args, names, star, starstar
    global MaxNumCallArgs
    n = len(p.args)
    MaxNumCallArgs = max(MaxNumCallArgs, n)

    arglist = ', '.join(["%s" % (a.visit(self)) for a in p.args])

    # We are somewhat limited in what we can call with stars or names:
    #print >> sys.stderr, '// Vcall: fn:', repr(p.fn)
    #print >> sys.stderr, '// Vcall: args:', repr(p.args)
    #print >> sys.stderr, '// Vcall: names:', repr(p.names)
    #print >> sys.stderr, '// Vcall: star:', repr(p.star)
    #print >> sys.stderr, '// Vcall: starstar:', repr(p.starstar)
    print '// Vcall: fn:', repr(p.fn)
    print '// Vcall: args:', repr(p.args)
    print '// Vcall: names:', repr(p.names)
    print '// Vcall: star:', repr(p.star)
    print '// Vcall: starstar:', repr(p.starstar)
    if p.star or p.starstar or any(p.names):
      return 'P(%s).(ICallV).CallV([]P{%s}, %s, []KV{%s}, %s) ' % (

          p.fn.visit(self),

          ', '.join([str(p.args[i].visit(self)) for i in range(len(p.args)) if not p.names[i]]),  # fixed args with no names.

          ('(%s).List()' % p.star.visit(self)) if p.star else 'nil',

          ', '.join(['KV{"%s", %s}' % (p.names[i], p.args[i].visit(self)) for i in range(len(p.args)) if p.names[i]]),  # named args.

          ('(%s).Dict()' % p.starstar.visit(self)) if p.starstar else 'nil',
      )

    if type(p.fn) is Tfield:
      if type(p.fn.p) is Tvar:

        if p.fn.p.name == 'super':
          return 'self.%s.M_%d_%s(%s)' % (self.tailSup(self.sup), n, p.fn.field, arglist)

        if p.fn.p.name in self.imports:
          imp = self.imports[p.fn.p.name]
          if imp.fromWhere:
            return 'MkGo(i_%s.%s).Call(%s) ' % (p.fn.p.name, p.fn.field, arglist)
          else:
            return 'call_%d( i_%s.G_%s, %s) ' % (n, p.fn.p.name, p.fn.field, arglist)

            # return 'call_%d( P(%s), %s )' % (n, p.fn.visit(self), arglist)
            # return 'i_%s.M_%d_%s(%s) ' % (p.fn.p.name, n, p.fn.field, arglist)

      # General Method Invocation.
      key = '%d_%s' % (n, p.fn.field)
      self.invokes[key] = (n, p.fn.field)
      return '/**/ f_INVOKE_%d_%s(%s, %s) ' % (n, p.fn.field, p.fn.p.visit(self), arglist)

    zfn = p.fn.visit(self)
    if type(zfn) is Zbuiltin:
      if p.fn.name == 'go_deref':
        return 'GoDeref(%s)' % p.args[0].visit(self)
      if p.fn.name == 'go_reify':
        return 'GoReify(%s)' % p.args[0].visit(self)
      if p.fn.name == 'go_type':
        return 'GoElemType(new(%s.%s))' % (p.args[0].p.visit(self), p.args[0].field)
      elif p.fn.name == 'go_new':
        return 'MkGo(new(%s.%s))' % (p.args[0].p.visit(self), p.args[0].field)
      elif p.fn.name == 'go_cast':
        return 'GoCast(GoElemType(new(%s.%s)), %s)' % (p.args[0].p.visit(self), p.args[0].field, p.args[1].visit(self))
      if p.fn.name == 'rye_pickle':
        return 'MkStr(string(Pickle(%s))) ' % p.args[0].visit(self)
      elif p.fn.name == 'rye_unpickle':
        return 'UnPickle(%s.String()) ' % p.args[0].visit(self)
      else:
        return 'G_%d_%s(%s) ' % (n, zfn.t.name, arglist)

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
    if type(x) is Zrye:
      raise Excpetion('Special Rye syntax "rye.%s" not used with Function Call syntax' % p.field)
    if type(x) is Zsuper:
      raise Excpetion('Special syntax "super" not used with Function Call syntax')
    if type(x) is Zself and not self.cls:
      raise Exception('Using a self field but not in a class definition: field="%s"' % p.field)
    if type(x) is Zself and self.instvars.get(p.field):  # Special optimization for self instvars.
      return '%s.M_%s' % (x, p.field)
    elif type(x) is Zimport:
      if x.fromWhere:
        return '/**/ MkGo(%s.%s) ' % (x, p.field)
      else:
        return '/**/ %s.G_%s' % (x, p.field)
    else:
      self.gsNeeded[p.field] = True
      return ' f_GET_%s(P(%s)) ' % (p.field, x)

  def Vnative(self, p):
    for s in p.ss:
      print s
    return

    # OLD:
    buf = PushPrint()

    print '/*NATIVE{*/'
    for s in p.strings:
      print '/**/ %s' % s
    print '/*}NATIVE*/'

    PopPrint()
    code = str(buf)
    self.tail.append(code)

  def Vdef(self, p):
    # name, args, dflts, star, starstar, body.
    buf = PushPrint()

    finder = YieldAndGlobalFinder()
    finder.Vsuite(p.body)
    self.yields = finder.yields
    self.force_globals = finder.force_globals
    # TODO: stack force_globals, if nested.

    # Tweak args.  Record meth, if meth.
    args = p.args  # Will drop the initial 'self' element, if in a cls.
    dflts = p.dflts  # Will drop the initial 'self' element, if in a cls.
    if self.cls:
      if len(p.args) > 0 and p.args[0] == 'self':  # User may omit self.
        args = p.args[1:]  # Skip self; it is assumed.
        dflts = p.dflts[1:]  # Skip self; it is assumed.
      self.meths[p.name] = ArgDesc(self.modname, self.cls, '%s.%s::%s' % (self.modname, self.cls, p.name), args, dflts, p.star, p.starstar)
    else:
      self.defs[p.name] = ArgDesc(self.modname, None, '%s.%s' % (self.modname, p.name), args, p.dflts, p.star, p.starstar)

    # prepend new scope dictionary, containing just the args, so far.
    self.scopes = [ dict([(a, 'a_%s' % a) for a in p.argsPlus]) ] + self.scopes

    #################
    # Render the :ody, but hold it in buf2, because we will prepend the vars.
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
    #################
    print '///////////////////////////////'

    letterV = 'V' if p.star or p.starstar else ''
    emptiesV = (', MkList(nil), MkDict(nil)' if args else 'MkList(nil), MkDict(nil)') if p.star or p.starstar else ''
    stars = ' %s P, %s P' % (AOrSkid(p.star), AOrSkid(p.starstar)) if p.star or p.starstar else ''

    if self.cls:
      gocls = self.cls if self.sup == 'native' else 'C_%s' % self.cls
      func = 'func (self *%s) M_%d%s_%s' % (gocls, len(args), letterV, p.name)
    else:
      func = 'func G_%d%s_%s' % (len(args), letterV, p.name)

    print ''
    print ' %s(%s %s) P {' % (func, ' '.join(['a_%s P,' % a for a in args]), stars)
    if self.force_globals:
      print '  //// self.force_globals:', self.force_globals

    scope = self.scopes[0]
    for v, v2 in scope.items():
      if v2[0] != 'a':  # Skip args
        print "   var %s P = None; _ = %s" % (v2, v2)
    print code2

    # Pop that scope.
    self.scopes = self.scopes[1:]

    print '   return None'
    print ' }'
    print ''

    print '///////////////////////////////'
    print '// name:', p.name
    print '// args:', p.args
    print '// dflts:', p.dflts
    print '// star:', p.star
    print '// starstar:', p.starstar
    print '// body:', p.body
    sys.stdout.flush()

    n = len(args)
    argnames = ', '.join(['"%s"' % a for a in p.args])
    defaults = ', '.join([(str(d.visit(self)) if d else 'nil') for d in p.dflts])

    pFuncMaker = '&pFunc_%s{PCallSpec: PCallSpec{Name: "%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}}' % (p.name, p.name, argnames, defaults, p.star, p.starstar)

    if self.cls:
      print ' type pMeth_%d_%s__%s struct { PCallSpec; Rcvr *%s }' % (n, self.cls, p.name, gocls)
      print ' func (o *pMeth_%d_%s__%s) Contents() interface{} {' % (n, self.cls, p.name)
      print '   return o.Rcvr.M_%d%s_%s' % (n, letterV, p.name)
      print ' }'
      print ' func (o *pMeth_%d_%s__%s) Call%d(%s) P {' % (n, self.cls, p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '   return o.Rcvr.M_%d%s_%s(%s%s)' % (n, letterV, p.name, ', '.join(['a%d' % i for i in range(n)]), emptiesV)
      print ' }'
      print ''
      print ' func (o *pMeth_%d_%s__%s) CallV(a1 []P, a2 []P, kv1 []KV, kv2 map[string]P) P {' % (n, self.cls, p.name)
      print '   argv, star, starstar := SpecCall(&o.PCallSpec, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   return o.Rcvr.M_%dV_%s(%s star, starstar)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''

    else:
      print ' type pFunc_%s struct { PCallSpec }' % p.name
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
      print '   argv, star, starstar := SpecCall(&o.PCallSpec, a1, a2, kv1, kv2)'
      print '   _, _, _ = argv, star, starstar'

      if p.star or p.starstar:  # If either, we always pass both.
        print '   return G_%dV_%s(%s star, starstar)' % (n, p.name, ' '.join(['argv[%d],' % i for i in range(n)]))
      else:  # If neither, we never pass either.
        print '   return G_%d_%s(%s)' % (n, p.name, ', '.join(['argv[%d]' % i for i in range(n)]))

      print ' }'
      print ''
      self.glbls[p.name] = ('*pFunc_%s' % p.name, pFuncMaker)

    PopPrint()
    code = str(buf)
    self.tail.append(code)
    self.force_globals = dict()  # TODO: unstack, if nested.

    ## The class constructor gets the args of init:
    #if self.cls and p.name == '__init__':
    #  self.args = args

  def qualifySup(self, sup):
    if type(sup) == Tvar:
      return 'C_%s' % sup.name
    elif type(sup) == Tfield:
      return 'i_%s.C_%s' % (sup.p.name, sup.field)
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def tailSup(self, sup):
    if type(sup) == Tvar:
      return 'C_%s' % sup.name
    elif type(sup) == Tfield:
      return 'C_%s' % sup.field
    else:
      raise Exception('qualifySup: Strange sup: %s' % sup)

  def Vclass(self, p):
    # name, sup, things
    self.cls = p.name
    self.sup = p.sup
    self.instvars = {}
    self.meths = {}

    #self.args = None  # If no __init__, we need to fix it later, but we can't (we don't know the arity of foreign __init__s) so we need CallV.

    gocls = self.cls if self.sup == 'native' else 'C_%s' % self.cls
    # Emit all the methods of the class (and possibly other members).
    for x in p.things:
      x.visit(self)
    self.cls = ''

    buf = PushPrint()

    # Emit the struct for the class.
    if self.sup != 'native':
      print '''
 type C_%s struct {
   %s
%s
 }

 func init() { Classes[`%s.C_%s`] = reflect.TypeOf(C_%s{}) }

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

      spec = 'PCallSpec: PCallSpec{Name: "%s::%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}' % (p.name, m, argnames, defaults, self.meths[m].star, self.meths[m].starstar)

      print ' func (o *%s) GET_%s() P { z := &pMeth_%d_%s__%s { %s, Rcvr: o }; z.SetSelf(z); return z }' % (gocls, m, n, p.name, m, spec)

    # Special methods for classes.
    if self.sup != 'native':
      #desc = self.meths['__init__']
      #if desc is None:
      #  # Until we can pass constructor args by * and **, we might not know how to call it if it came from another package.
      #  # So until then, we require an __init__ in every class.
      #  raise Exception('Method __init__ must be defined for class %s' + self.name)

      print 'func (o *C_%s) Rye_ClearFields__() {' % p.name
      for iv in self.instvars:
        print '   o.M_%s = None' % iv
      if p.sup and type(p.sup) is Tvar:
        print '// superclass:', p.sup.visit(self)
        if p.sup.name not in ['native', 'object']:
          print '   o.C_%s.Rye_ClearFields__()' % p.sup.name
      if p.sup and type(p.sup) is Tfield:
        print '// superclass:', p.sup.visit(self)
        print '   o.C_%s.Rye_ClearFields__()' % p.sup.field
      print '}'

      print ''
      print 'func (o *C_%s) Type() P { return G_%s }' % (p.name, p.name)
      print 'func (o *pFunc_%s) Repr() string { return "%s" }' % (p.name, p.name)
      print 'func (o *pFunc_%s) String() string { return "<class %s>" }' % (p.name, p.name)
      print ''


    self.tail.append(str(buf))
    PopPrint()

  def Vsuite(self, p):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
    self.Vseq(p)  # Because they are the same.

  def Vseq(self, p):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
    for x in p.things:
      print '// @ %d @ %s' % (x.where, x.gloss)
      x.visit(self)
      print '// $ %d $ %s' % (x.where, x.gloss)

PrinterStack= []
def PushPrint():
    global PrinterStack
    sys.stdout.flush()
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
    self.b.append(x)
  def flush(self):
    pass
  def __str__(self):
    z = ''.join(self.b)
    return z

class Tnode(object):
  def visit(self, a):
    raise Bad('unimplemented visit %s %s', self, type(self))

class Top(Tnode):
  def __init__(self, a, op, b=None, returns_bool=False):
    self.a = a
    self.op = op
    self.b = b
    self.returns_bool = returns_bool
  def visit(self, v):
    return v.Vop(self)

class Tboolop(Tnode):
  def __init__(self, a, op, b=None):
    self.a = a
    self.op = op
    self.b = b
  def visit(self, v):
    return v.Vboolop(self)

class Tcondop(Tnode):
  def __init__(self, a, b, c):
    self.a = a
    self.b = b
    self.c = c
  def visit(self, v):
    return v.Vcondop(self)

class Tgox(Tnode):
  def __init__(self, fcall):
    self.fcall = fcall
  def visit(self, v):
    return v.Vgox(self)

class Traw(Tnode):
  def __init__(self, raw):
    self.raw = raw
  def visit(self, v):
    return self.raw

class Tlit(Tnode):
  def __init__(self, k, v):
    self.k = k
    self.v = v
  def visit(self, v):
    return v.Vlit(self)

class Tvar(Tnode):
  def __init__(self, name):
    self.name = name
  def visit(self, v):
    return v.Vvar(self)

class Titems(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vtuple(self)  # By default, make a tuple.

class Ttuple(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vtuple(self)

class Tlist(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vlist(self)

class Tforexpr(Tnode):
  def __init__(self, z, vv, ll, cond, has_comma):
    self.z = z
    self.vv = vv
    self.ll = ll
    self.cond = cond
    self.has_comma = has_comma
  def visit(self, v):
    return v.Vforexpr(self)

class Tdict(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vdict(self)

class Tsuite(Tnode):  # So far, Tsuite and Tseq, and Vsuite and Vseq, are the same.
  def __init__(self, things):
    self.things = things
  def visit(self, v):
    return v.Vsuite(self)

class Tseq(Tnode):  # So far, Tsuite and Tseq, and Vsuite and Vseq, are the same.
  def __init__(self, things):
    self.things = things
  def visit(self, v):
    return v.Vseq(self)

class Texpr(Tnode):
  def __init__(self, a):
    self.a = a
  def visit(self, v):
    return v.Vexpr(self)

class Tassign(Tnode):
  def __init__(self, a, b, pragma=None):
    self.a = a
    self.b = b
    self.pragma = pragma
  def visit(self, v):
    return v.Vassign(self)

class Tprint(Tnode):
  def __init__(self, xx, saying, code):
    self.xx = xx
    self.saying = saying
    self.code = code
  def visit(self, v):
    return v.Vprint(self)

class Tdefer(Tnode):
  def __init__(self, cmd):
    self.cmd = cmd
  def visit(self, v):
    return v.Vdefer(self)

class Twithdefer(Tnode):
  def __init__(self, call, body):
    self.call = call
    self.body = body
  def visit(self, v):
    return v.Vwithdefer(self)

class Tglobal(Tnode):
  def __init__(self, vars):
    self.vars = vars
  def visit(self, v):
    return v.Vglobal(self)

class Timport(Tnode):
  def __init__(self, imported, alias, fromWhere):
    self.imported = imported
    self.alias = alias
    self.fromWhere = fromWhere
  def visit(self, v):
    return v.Vimport(self)

class Tassert(Tnode):
  def __init__(self, x, y, code, is_must):
    self.x = x
    self.y = y
    self.code = code
    self.is_must = is_must
  def visit(self, v):
    return v.Vassert(self)

class Ttry(Tnode):
  def __init__(self, tr, exvar, ex):
    self.tr = tr
    self.exvar = exvar
    self.ex = ex
  def visit(self, v):
    return v.Vtry(self)

class Tif(Tnode):
  def __init__(self, t, yes, no):
    self.t = t
    self.yes = yes
    self.no = no
  def visit(self, v):
    return v.Vif(self)

class Twhile(Tnode):
  def __init__(self, t, yes):
    self.t = t
    self.yes = yes
  def visit(self, v):
    return v.Vwhile(self)

class Tfor(Tnode):
  def __init__(self, var, t, b, foo):
    self.var = var
    self.t = t
    self.b = b
    self.foo = foo
  def visit(self, v):
    return v.Vfor(self)

class Treturn(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, v):
    return v.Vreturn(self)

class Tyield(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, v):
    return v.Vyield(self)

class Tbreak(Tnode):
  def __init__(self):
    pass
  def visit(self, v):
    return v.Vbreak(self)

class Tcontinue(Tnode):
  def __init__(self):
    pass
  def visit(self, v):
    return v.Vcontinue(self)

class Traise(Tnode):
  def __init__(self, a):
    self.a = a
  def visit(self, v):
    return v.Vraise(self)

class Tdel(Tnode):
  def __init__(self, listx):
    self.listx = listx
  def visit(self, v):
    return v.Vdel(self)

class Tnative(Tnode):
  def __init__(self, ss):
    self.ss = ss
  def visit(self, v):
    return v.Vnative(self)

class Tdef(Tnode):
  def __init__(self, name, args, dflts, star, starstar, body):
    self.name = name
    self.args = args
    self.dflts = dflts
    self.star = star
    self.starstar = starstar
    self.body = body

    self.argsPlus = args[:]
    if star:
      self.argsPlus += [star]
    if starstar:
      self.argsPlus += [starstar]

    if len(args) != len(dflts):
      raise Exception('len args (%s) != len dflts (%s) ::: %s' % (args, dflts, vars(self)))

  def visit(self, v):
    return v.Vdef(self)

class Tclass(Tnode):
  def __init__(self, name, sup, things):
    self.name = name
    self.sup = sup
    self.things = things
  def visit(self, v):
    return v.Vclass(self)

class Tcall(Tnode):
  def __init__(self, fn, args, names, star, starstar):
    self.fn = fn
    self.args = args
    self.names = names
    self.star = star
    self.starstar = starstar
  def visit(self, v):
    return v.Vcall(self)

class Tfield(Tnode):
  def __init__(self, p, field):
    self.p = p
    self.field = field
  def visit(self, v):
    return v.Vfield(self)

class Tgetitem(Tnode):
  def __init__(self, a, x):
    self.a = a
    self.x = x
  def visit(self, v):
    return v.Vgetitem(self)

class Tgetitemslice(Tnode):
  def __init__(self, a, x, y, z):
    self.a = a
    self.x = x
    self.y = y
    self.z = z
  def visit(self, v):
    return v.Vgetitemslice(self)

class Tcurlysetter(Tnode):
  def __init__(self, obj, vec):
    self.obj = obj  # The object
    self.vec = vec  # Pairs of var, expr
  def visit(self, v):
    return v.Vcurlysetter(self)

class Parser(object):
  def __init__(self, program, words, p):
    self.program = program  # Source code
    self.words = words      # Lexical output
    self.litInts = {}       # value -> name
    self.litStrs = {}       # value -> name
    self.k = ''
    self.v = ''
    self.p = p
    self.i = 0
    self.Advance()

  def Bad(self, format, *args):
    msg = format % args
    self.Info(msg)
    raise Exception(msg)

  def Info(self, msg):
    sys.stdout.flush()
    print >> sys.stderr, 120 * '#'
    print >> sys.stderr, '   msg = ', msg
    print >> sys.stderr, '   k =', repr(self.k)
    print >> sys.stderr, '   v =', repr(self.v)
    print >> sys.stderr, '   rest =', repr(self.Rest())

  def LookAheadV(self):
    if self.p+1 < len(self.words):
      return self.words[self.p+1][1]

  def Advance(self):
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]

  def Rest(self):
    return self.program[self.i:]

  def Eat(self, v):
    if self.v != v:
      raise self.Bad('Expected %s, but got %s, at %s', v, self.v, repr(self.Rest()))
    self.Advance()

  def EatK(self, k):
    if self.k != k:
      raise self.Bad('Expected Kind %s, but got %s, at %s', k, self.k, repr(self.Rest()))
    self.Advance()

  def Pid(self):
    if self.k != 'A':
      raise self.Bad('Pid expected kind A, but got kind=%s; rest=%s', self.k, repr(self.Rest()))
    z = self.v
    self.Advance()
    return z

  def Xvar(self):
    if self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z
    else:
      raise self.Bad('Xvar expected variable name, but got kind=%s; rest=%s', self.k, repr(self.Rest()))

  def Xqualname(self):  # A possibly-singly-qualifed name (qualifier should be an import).
    if self.k == 'A':
      a = Tvar(self.v)
      self.Advance()
      if self.v == '.':
        self.Advance()
        field = self.v
        self.EatK('A')
        return Tfield(a, field)
      else:
        return a
    else:
      raise self.Bad('Xqualname expected variable name, but got "%s"; rest=%s', self.v, repr(self.Rest()))

  def Xprim(self):
    if self.k == 'N':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.k == 'F':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.k == 'S':
      z = Tlit(self.k, self.v)
      self.Advance()
      return z

    if self.v == 'go':
      self.Advance()
      fcall = self.Xsuffix()
      if type(fcall) != Tcall:
        raise Exception('Go expression must be func or method call: %s' % fcall)
      z = Tgox(fcall)
      return z

    if self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      return z

    if self.v == '.': # Experimental self-less dot.
      self.Eat('.')
      field = self.v
      self.EatK('A')
      return Tfield(Tvar('self'), field)

    if self.k == 'K':
      if self.v in ['None', 'True', 'False']:
        v = self.v
        self.Eat(self.v)
        z = Traw(v)
        return z
      raise Exception('Keyword "%s" cannot be used as a variable or expression' % self.v)

    if self.v == '(':
      self.Eat('(')
      if self.v == ')':
        self.Eat(')')
        # Unit tuple.
        return Ttuple([])
      z = self.Xexpr()
      if self.v == ')':
        # Not a tuple, just an Xexpr.
        self.Eat(')')
        return z
      # z is just the first in a tuple.
      self.Eat(',')
      z = [z]
      while self.v != ')':
        x = self.Xexpr()
        z.append(x)
        if self.v == ')':
          # Omitted trailing ','
          break
        self.Eat(',')
      self.Eat(')')
      return Ttuple(z)

    if self.v == '[':
      has_comma = False
      self.Eat('[')
      z = []
      while self.v != ']':
        x = self.Xexpr()
        z.append(x)
        if self.v == ']':
          # Omitted trailing ','
          break
        if self.v == 'for':
          # For expression (list interpolation)
          if has_comma:
            raise Exception('"for" not allowed after a bare tuple; use parens.')
          self.Eat('for')
          vv = self.Xvars()
          self.Eat('in')
          ll = self.Xor()  # Avoid confusion with 'if' Xcond
          cond = None
          if self.v == 'if':
            self.Eat('if')
            cond = self.Xexpr()
          self.Eat(']')
          return Tforexpr(z, vv, ll, cond, has_comma)
        self.Eat(',')
        has_comma = True
      self.Eat(']')
      return Tlist(z)

    if self.v == '{':
      self.Eat('{')
      z = []
      while self.v != '}':
        x = self.Xexpr()
        self.Eat(':')
        y = self.Xexpr()
        z.append(x)
        z.append(y)
        if self.v == '}':
          # Omitted trailing ','
          break
        self.Eat(',')
      self.Eat('}')
      return Tdict(z)

    else:
      raise self.Bad('Expected Xprim, but got %s, at %s', self.v, repr(self.Rest()))

  def Xsuffix(self):
    """Tcall, Tfield, or Tindex"""
    a = self.Xprim()
    while True:
      if self.v == '(':
        self.Eat('(')

        args = []
        names = []
        star = None
        starstar = None
        while self.v != ')':

          # Look for case with named parameter
          named = ''
          starred = ''
          if self.k == 'A' and self.LookAheadV() == '=':
            named = self.v
            self.EatK('A')
            self.Eat('=')
          elif self.v in ['*', '**']:
            starred = self.v
            self.Eat(starred)

          b = self.Xexpr()
          if starred == '*':
            star = b
          elif starred == '**':
            starstar = b
          else:
            args.append(b)
            names.append(named)

          if self.v == ',':
            self.Eat(',')
          else:
            break
        self.Eat(')')
        a = Tcall(a, args, names, star, starstar)

      elif self.v == '.':
        self.Eat('.')
        field = self.v
        self.EatK('A')
        a = Tfield(a, field)

      elif self.v == '[':
        self.Eat('[')
        if self.v != ':':
          x = self.Xexpr()
        else:
          x = None

        if self.v == ']':
          self.Eat(']')
          if x is None:
            raise Bad('Index cannot be None')
          a = Tgetitem(a, x)
        else:
          self.Eat(':')
          if self.v != ']':
            y = self.Xexpr()
          else:
            y = None
          self.Eat(']')
          a = Tgetitemslice(a, x, y, None)

      elif self.v == '{':
        # Curly Setter: can follow any Rye or Go object.
        # Looks kind of like the Curly-Brace initializer in Go.
        # But instead of starting with a type name, start with an object.
        #   e.g.
        #     b = gonew(fruit.Banana) {Color: 'yellow'}
        #   or
        #     opts = file.Opts() {How: 'r+', Mode: 438, Buffer: 4096}
        # This both changes the object and returns it.
        self.Eat('{')
        vec = []
        while self.v != '}':
          var = self.Xvar()
          self.Eat(':')
          value = self.Xexpr()
          vec.append((var, value))

          if self.v == ',':
            self.Eat(',')
          elif self.v != '}':
            raise Exception('Expected "," or "}" in Field Constructor, but got "%s"' % self.v)
        self.Eat('}')
        return Tcurlysetter(a, vec)

      else:
        break
    return a

  def Xunary(self):
    if self.v in UNARY_OPS:
      op = self.v
      self.Eat(op)
      a = self.Xunary()
      return Top(a, UNARY_OPS[op], None)
    else:
      return self.Xsuffix()

  def Xmul(self):
    a = self.Xunary()
    while self.v in MUL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xunary()
      a = Top(a, MUL_OPS[op], b)
    return a

  def Xadd(self):
    a = self.Xmul()
    while self.v in ADD_OPS or (self.k == 'N' and self.v[0] in '+-'):
      if self.k == 'N':
        op = self.v[0]  # The sign is the op.
        b = Tlit('N', self.v[1:])  # The rest of the number.
        self.Advance()
      else:
        op = self.v
        self.Eat(op)
        b = self.Xmul()
      a = Top(a, ADD_OPS[op], b)
    return a

  def Xshift(self):
    a = self.Xadd()
    while self.v in SHIFT_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xadd()
      a = Top(a, SHIFT_OPS[op], b)
    return a

  def Xbitand(self):
    a = self.Xshift()
    while self.v == '&':
      op = self.v
      self.Eat(op)
      b = self.Xshift()
      a = Top(a, "BitAnd", b)
    return a

  def Xbitxor(self):
    a = self.Xbitand()
    while self.v == '^':
      op = self.v
      self.Eat(op)
      b = self.Xbitand()
      a = Top(a, "BitXor", b)
    return a

  def Xbitor(self):
    a = self.Xbitxor()
    while self.v == '|':
      op = self.v
      self.Eat(op)
      b = self.Xbitxor()
      a = Top(a, "BitOr", b)
    return a

  def Xrelop(self):
    a = self.Xbitor()
    if self.v in REL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xbitor()
      a = Top(a, REL_OPS[op], b, True)
    elif RE_WORDY_REL_OP.match(self.v):
      op = self.v
      self.Eat(op)
      b = self.Xbitor()
      if op == 'in':
        a = Top(b, 'Contains', a, True)    # N.B. swap a & b for Contains
      elif RE_NOT_IN.match(op):
        a = Top(b, 'NotContains', a, True)    # N.B. swap a & b for NotContains
      elif op == 'is':
        a = Top(a, 'Is', b, True)
      elif RE_IS_NOT.match(op):
        a = Top(b, 'IsNot', a, True)
      else:
        raise Exception('Weird RE_WORDY_REL_OP: %s' % op)
    return a

  def Xnot(self):
    if self.v == 'not':
      self.Eat('not')
      b = self.Xrelop()
      return Tboolop(b, "!")
    else:
      return self.Xrelop()

  def Xand(self):
    a = self.Xnot()
    while self.v == 'and':
      op = self.v
      self.Eat(op)
      b = self.Xnot()
      a = Tboolop(a, "&&", b)
    return a

  def Xor(self):
    a = self.Xand()
    while self.v == 'or':
      op = self.v
      self.Eat(op)
      b = self.Xand()
      a = Tboolop(a, "||", b)
    return a

  def Xcond(self):
    z = self.Xor()
    if self.v == 'if':
      self.Eat('if')
      a = self.Xcond()
      self.Eat('else')
      c = self.Xcond()
      z = Tcondop(a, z, c)
    return z

  def Xexpr(self):
    return self.Xcond()

  def Xlistexpr(self):
    return self.Xitems(allowScalar=True, allowEmpty=False)

  def Xvars(self):
    z = []
    has_comma = False
    while True:
      if self.v == '(':
        self.Eat('(')
        x = self.Xvars()
        self.Eat(')')
        z.append(x)
      elif self.k == 'A':
        z.append(Tvar(self.v))
        self.Advance()
      else:
        break

      if self.v != ',':
        break
      self.Eat(',')
      has_comma = True

    if not z:
      raise Exception('Expected variable name after "for"')

    if has_comma:
      return Titems(z)  # List of items.
    else:
      return z[0]

  def Xitems(self, allowScalar, allowEmpty):
    "A list of expressions, possibly empty, or possibly a scalar."
    z = []
    comma_needed = False  # needed before more in the list.
    had_comma = False
    while self.k != ';;' and self.k != 'P' and self.v not in [')', ']', '}', ':', '=', '+=', '-=', '*=', '/=']:
      if self.v == ',':
        self.Eat(',')
        had_comma = True
        comma_needed = False
      else:
        if comma_needed:
          raise Exception('Comma required before more items in list')
        x = self.Xexpr()
        z.append(x)
        comma_needed = True
    if allowScalar and len(z) == 1 and not had_comma:
      return z[0]  # Scalar.
    if not allowEmpty and len(z) == 0:
      raise Exception('Empty expression list not allowed')
    return Titems(z)  # List of items.

#  def ParseTarget(self):
#    "Parse a variable name (return str) or vector of targets (return list)."
#    if self.v == '(':
#      self.Advance()
#      x = self.ParseTarget()
#      self.Eat(')')
#      return x
#
#    elif self.k == 'A':
#      first = Target(self.v)
#      self.Advance()
#      vec = []
#      while self.v == ',':
#        self.Advance()
#        vec.append(self.ParseTarget())
#      if vec:
#        return [first] + vec
#      else:
#        return first
#
#    else:
#      raise 'Expected "(" or Identifier, got %s: %s' % (self.k, self.v) # ')'

  def Csuite(self):
    things = []
    while self.k != 'OUT' and self.k is not None:
      #print '//Csuite', self.k, self.v
      if self.v == ';;':
        self.EatK(';;')
      else:
        if RYE_FLOW:
          num = 1 + sum([int(ch=='\n') for ch in self.program[ : self.i ]])
          what= '"## LINE ## %d ##"' % num
          things.append(Tprint(Tlist([Tlit('S', what)]), False, None))
        t = self.Command();
        if t:
          things.append(t)
    return Tsuite(things)

  def Command(self):
    where = self.i
    gloss = self.v
    cmd = self.Command9()
    if cmd:
      # Tag the cmd node with where it was in source.
      cmd.where = where
      cmd.gloss = gloss
    return cmd

  def Command9(self):
    if self.v == 'print':
      return self.Cprint(False)
    if self.v == 'say':
      return self.Cprint(True)
    elif self.v == 'if':
      return self.Cif()
    elif self.v == 'while':
      return self.Cwhile()
    elif self.v == 'for':
      return self.Cfor()
    elif self.v == 'return':
      return self.Creturn()
    elif self.v == 'yield':
      return self.Cyield()
    elif self.v == 'break':
      return self.Cbreak()
    elif self.v == 'continue':
      return self.Ccontinue()
    elif self.v == 'raise':
      return self.Craise()
    elif self.v == 'def':
      return self.Cdef('')
    elif self.v == 'class':
      return self.Cclass()
    elif self.v == 'assert':
      return self.Cassert()
    elif self.v == 'must':
      return self.Cassert(True)
    elif self.v == 'import':
      return self.Cimport()
    elif self.v == 'from':
      return self.Cfrom()
    elif self.v == 'defer':
      return self.Cdefer()
    elif self.v == 'with':
      return self.Cwith()
    elif self.v == 'global':
      return self.Cglobal()
    elif self.v == 'try':
      return self.Ctry()
    elif self.v == 'del':
      return self.Cdel()
    elif self.v == 'native':
      return self.Cnative()
    elif self.v == 'pass':
      self.Eat('pass')
      self.EatK(';;')
      return
    elif self.k == 'S':  # String as comment.
      self.EatK('S')
      self.EatK(';;')
      return
    elif self.k == 'A' or self.v == '.' or self.v == '(' or self.v == 'go':
      return self.Cother()
    else:
      raise self.Bad('Unknown stmt: %s %s %s', self.k, self.v, repr(self.Rest()))

  def Cother(self):
    a = self.Xitems(allowScalar=True, allowEmpty=False)  # lhs (unless not an assignment; then it's the only thing.)

    op = self.v

    if op in ['+=', '-=', '*=']:
      self.Eat(op)
      binop = op[:-1]  # Remove the '='
      b = self.Xexpr()
      # TODO: this evals lhs twice.
      if binop in ADD_OPS:
        return Tassign(a, Top(a, ADD_OPS[binop], b))
      elif binop in MUL_OPS:
        return Tassign(a, Top(a, MUL_OPS[binop], b))
      else:
        raise Exception('Unknown op, neither ADD_OPS nor MUL_OPS: ' + binop)

    elif op == '=':
      self.Eat(op)
      b = self.Xlistexpr()
      pragma = None
      if self.k == 'P':
        pragma = TrimPragma(self.v)
        self.EatK('P')
      return Tassign(a, b, pragma)

    else:
      # TODO: error if this is not a function or method call.
      return Tassign(Traw('_'), a)

  def Cprint(self, saying):
    # TODO: distinguish trailing ,
    self.Advance()
    begin = self.i
    t = self.Xitems(allowScalar=False, allowEmpty=True)
    end = self.i
    self.EatK(';;')
    return Tprint(t, saying, self.program[begin : end])

  # Old "defer"
  def Cdefer(self):
    self.Eat('defer')
    cmd = self.Cother()
    assert type(cmd) == Tassign
    if cmd.a.visit(self) != '_':
      raise Exception('"defer" statement cannot assign to variables')
    if type(cmd.b) != Tcall:
      raise Exception('"defer" statement must contain function or method call')
    return Tdefer(cmd)

  # New "with defer"
  def Cwith(self):
    self.Eat('with')
    self.Eat('defer')
    call = self.Xexpr()
    if type(call) != Tcall:
      raise Exception('"with defer" statement must contain function or method call')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    body = self.Csuite()
    self.EatK('OUT')
    return Twithdefer(call, body)

  def Cglobal(self):
    self.Eat('global')
    vars = {}
    while True:
      var = self.Xvar()
      vars[var.name] = var
      if self.v != ',':
        break
      self.Eat(',')
    self.EatK(';;')
    return Tglobal(vars)

  def Cfrom(self):
    self.Eat('from')
    if self.v == 'go':
      self.Eat('go')
      return self.Cimport(fromWhere='go')
    elif self.v == '.':
      self.Eat('.')
      return self.Cimport(fromWhere='.')
    raise Exception('from command: must be "from go import ..." or "from . import ..."')

  def Cimport(self, fromWhere=None):
    #self.Eat('import')
    #alias = self.v
    #self.EatK('A')
    #while self.v == '.':
    #  self.Eat('.')
    #  alias = '%s.%s' % (alias, self.v)
    #  self.EatK('A')
    #imported = [ alias ]
    #while self.v == '/':
    #  self.Eat('/')
    #  imported.append(self.v)
    #  alias = self.v
    #  self.EatK('A')

    self.Eat('import')
    s = ''
    while self.k in ['A', 'N'] or self.v in ['.', '-', '/']:
      s += self.v
      self.Advance()

    imported = s.split('/')
    alias = imported[-1]

    if self.v == 'as':
      self.Eat('as')
      alias = self.v
      self.EatK('A')
    self.EatK(';;')

    return Timport(imported, alias, fromWhere)

  def Cassert(self, is_must=False):
    i = self.i
    self.Advance()
    x = self.Xexpr()
    y = None
    j = self.i
    if self.v == ',':
      self.Eat(',')
      y = self.Xlistexpr()
    return Tassert(x, y, self.program[i:j], is_must)

  def Ctry(self):
    exvar = None
    self.Eat('try')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    tr = self.Csuite()
    self.EatK('OUT')
    self.Eat('except')
    if self.v == 'as':
      self.Eat('as')
      exvar = self.Xvar()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    ex = self.Csuite()
    self.EatK('OUT')
    return Ttry(tr, exvar, ex)

  def Cif(self):
    self.Advance()
    t = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    yes = self.Csuite()
    self.EatK('OUT')
    no = None
    if self.v == 'elif':
      no = self.Cif()
      return Tif(t, yes, no)
    if self.v == 'else':
      self.Eat('else')
      self.Eat(':')
      self.EatK(';;')
      self.EatK('IN')
      no = self.Csuite()
      self.EatK('OUT')
    return Tif(t, yes, no)

  def Cwhile(self):
    self.Eat('while')
    t = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    yes = self.Csuite()
    self.EatK('OUT')
    return Twhile(t, yes)

  def Cfor(self):
    self.Eat('for')
    x = self.Xvars()
    self.Eat('in')
    t = self.Xlistexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    suite = self.Csuite()
    self.EatK('OUT')
    return Tfor(x, t, suite, 2)

  def Creturn(self):
    self.Eat('return')
    if self.v == ';;':  # Missing Xitems means None, not [].
      return Treturn(None)
    t = self.Xlistexpr()
    return Treturn([t])

  def Cyield(self):
    self.Eat('yield')
    if self.v == ';;':  # Missing Xitems means None, not [].
      return Tyield(None)
    t = self.Xlistexpr()
    return Tyield([t])

  def Cbreak(self):
    self.Eat('break')
    self.EatK(';;')
    return Tbreak()

  def Ccontinue(self):
    self.Eat('continue')
    self.EatK(';;')
    return Tcontinue()

  def Craise(self):
    self.Eat('raise')
    t = self.Xlistexpr()
    return Traise(t)

  def Cdel(self):
    self.Eat('del')
    listx = self.Xlistexpr()
    return Tdel(listx)

  def Cnative(self):
    self.Eat('native')
    code = self.Xexpr()
    if type(code) is not Xlit or code.k != 'S':
      raise Exception('native expects a string literal, got %s' % code)
    return Tnative(code.v)

  def Cclass(self):
    self.Eat('class')
    name = self.Pid()
    sup = Tvar('object')

    if self.v == '(':
      self.Advance()
      if self.v == 'native':  # Special keyword.
        self.Advance()
        sup = 'native'
      else:
        sup = self.Xqualname()
      self.Eat(')')

    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')

    things = []
    while self.k != 'OUT':
      if self.v == 'def':
        t = self.Cdef(name)
        things.append(t)
      #elif self.v == 'native':
      #  t = self.Cnative()
      #  things.append(t)
      elif self.v == 'pass':
        self.Eat('pass')
      elif self.k == ';;':
        self.EatK(';;')
      else:
        raise self.Bad('Classes may only contain "def" or "pass" commands.')
    self.EatK('OUT')

    return Tclass(name, sup, things)

  def Cnative(self):
    self.Eat('native')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    strings = []
    while self.k == 'S' or self.k == ';;':
      if self.k == 'S':
        strings.append(self.v[1:-1])
      self.Advance()
    self.EatK('OUT')
    return Tnative(strings)

  def Cdef(self, cls):
    self.Eat('def')
    name = self.Pid()
    self.Eat('(')
    args = []
    dflts = []
    while self.k == 'A':
      arg = self.Pid()
      dflt = None
      if self.v == '=':
        self.Eat('=')
        dflt = self.Xexpr()
      if self.v == ',':
        self.Eat(',')
      args.append(arg)
      dflts.append(dflt)
    star = ''
    starstar = ''
    while self.v in ('*', '**'):
      which = self.v
      self.Advance()
      if which == '*':
        star = self.v
      elif which == '**':
        starstar = self.v
      self.EatK('A')
      if self.v == ',':
        self.Eat(',')
      pass
    self.Eat(')')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    suite = self.Csuite()
    self.EatK('OUT')
    return Tdef(name, args, dflts, star, starstar, suite)


def ParsePragma(s):
  if s == 'i':
    return Yint()
  elif s == 'f':
    return Yfloat()
  elif s == 's':
    return Ystr()
  else:
    raise Exception('Unknown Pragma: %s' % s)

class Y(object):  # Typed values
  def __init__(self):
    pass
class Yint(Y):
    pass
class Yfloat(Y):
    pass
class Ystr(Y):
    pass


class Z(object):  # Returns from visits (emulated runtime value).
  def __init__(self, t, s):
    self.t = t  # T node
    self.s = s  # String for backwards compat
  def __str__(self):
    return self.s
class Zrye(Z):
  pass
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
    Z.__init__(self, t, s)
    self.imp = imp  # imports[] object
    self.fromWhere = imp.fromWhere  # imports[] object
  pass
class Zbuiltin(Z):
  pass
class Zlit(Z):
  pass

pass


# OPERATOR HIERARCHY OF PYTHON
#lambda        Lambda expression
#if else        Conditional expression
#or        Boolean OR
#and        Boolean AND
#not x        Boolean NOT
#in, not in, is, is not, <, <=, >, >=, <>, !=, ==        Comparisons, including membership tests and identity tests
#|        Bitwise OR
#^        Bitwise XOR
#&        Bitwise AND
#<<, >>        Shifts
#+, -        Addition and subtraction
#*, /, //, %        Multiplication, division, remainder [8]
#+x, -x, ~x        Positive, negative, bitwise NOT
#**        Exponentiation [9]
#x[index], x[index:index], x(arguments...), x.attribute        Subscription, slicing, call, attribute reference
#(expressions...), [expressions...], {key: value...}, `expressions...`        Binding or tuple display, list display, dictionary display, string conversion



class StatementWalker(object):

  def Vexpr(self, p):
    pass

  def Vassign(self, p):
    pass

  def Vprint(self, p):
    pass

  def Vimport(self, p):
    pass

  def Vassert(self, p):
    pass

  def Vtry(self, p):
    p.tr.visit(self)
    p.ex.visit(self)

  def Vfor(self, p):
    p.b.visit(self)

  def Vif(self, p):
    p.yes.visit(self)
    if p.no:
      p.no.visit(self)

  def Vwhile(self, p):
    p.yes.visit(self)

  def Vreturn(self, p):
    pass

  def Vbreak(self, p):
    pass

  def Vcontinue(self, p):
    pass

  def Vraise(self, p):
    pass

  def Vdel(self, p):
    pass

  def Vdefer(self, p):
    pass

  def Vwithdefer(self, p):
    p.body.visit(self)

  def Vglobal(self, p):
    pass

  def Vlit(self, p):
    pass

  def Vop(self, p):
    pass

  def Vboolop(self, p):
    pass

  def Vcondop(self, p):
    pass

  def Vgetitem(self, p):
    pass

  def Vgetitemslice(self, p):
    pass

  def Vcurlysetter(self, p):
    pass

  def Vtuple(self, p):
    pass

  def Vlist(self, p):
    pass

  def Vdict(self, p):
    pass

  def Vvar(self, p):
    pass

  def Vcall(self, p):
    pass

  def Vgox(self, p):
    pass

  def Vfield(self, p):
    pass

  def Vnative(self, p):
    pass

  def Vdef(self, p):
    pass

  def Vclass(self, p):
    pass

  def Vsuite(self, p):
    for x in p.things:
      x.visit(self)

  def Vseq(self, p):
    for x in p.things:
      x.visit(self)

class YieldAndGlobalFinder(StatementWalker):
  def __init__(self):
    self.yields = False
    self.force_globals = {}

  def Vyield(self, p):
    self.yields = True

  def Vglobal(self, p):
    for v in p.vars:
      self.force_globals[v] = p.vars[v]

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
    return 'PCallSpec{Name: "%s", Args: []string{%s}, Defaults: []P{%s}, Star: "%s", StarStar: "%s"}' % (self.name, argnames, defaults, self.star, self.starstar)

def AOrSkid(s):
  if s:
    return 'a_%s' % s
  else:
    return '_'
