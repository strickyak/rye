import os
import re
import sys

BUILTINS = set(
    'len repr str int float list dict tuple range sorted type'
    .split())

# RE_WHITE returns 3 groups.
# The first group includes white space or comments, including all newlines, always ending with newline.
# The second group is buried in the first one, to provide any repetition of the alternation of white or comment.
# The third group is the residual white space at the front of the line after the last newline, which is the indentation that matters.
RE_WHITE = re.compile('(([ \t\n]*[#][^\n]*[\n]|[ \t\n]*[\n])*)?([ \t]*)')

RE_KEYWORDS = re.compile(
    '\\b(class|def|if|else|while|True|False|None|print|and|or|try|except|raise|return|break|continue|pass|as|go)\\b')
RE_LONG_OPS = re.compile(
    '[+]=|[-]=|[*]=|/=|//|<<|>>|==|!=|<=|>=|[*][*]|[.][.]')
RE_OPS = re.compile('[-.@~!%^&*+=,|/<>:]')
RE_GROUP = re.compile('[][(){}]')
RE_ALFA = re.compile('[A-Za-z_][A-Za-z0-9_]*')
RE_NUM = re.compile('[+-]?[0-9]+[-+.0-9_e]*')
RE_STR = re.compile('(["](([^"\\\\\n]|[\\\\].)*)["]|[\'](([^\'\\\\\n]|[\\\\].)*)[\'])')

RE_WORDY_REL_OP = re.compile('\\b(not\\s+in|is\\s+not|in|is)\\b')
RE_NOT_IN = re.compile('^not\\s+in$')
RE_IS_NOT = re.compile('^is\\s*not$')

TAB_WIDTH = 8

DETECTERS = [
  [RE_KEYWORDS, 'K'],
  [RE_WORDY_REL_OP, 'W'],
  [RE_ALFA, 'A'],
  [RE_NUM, 'N'],
  [RE_LONG_OPS, 'L'],
  [RE_OPS, 'O'],
  [RE_GROUP, 'G'],
  [RE_STR, 'S'],
]

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

NONALFA = re.compile('[^A-Za-z0-9_]')
def CleanIdentWithSkids(s):
  return NONALFA.sub('_', s)

def CleanQuote(x):
  return re.sub('[^A-Za-z0-9_]', '~', x)[:10]
  return '~~~'
  return re.sub('[^A-Za-z0-9_]', '~', x)
  return re.sub('[`]', '\'', x)
def Bad(format, *args):
  raise Exception(format % args)

class Lex(object):
  def __init__(self, program):
    self.buf = program
    self.i = 0
    self.line_num = 1
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
    self.line_num += sum([c == '\n' for c in blank_lines])

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
    self.lits = {}          # key -> name
    self.scopes = []
    self.tail = []
    self.cls = ''
    self.gsNeeded = {}        # keys are getter/setter names.

  def GenModule(self, modname, path, suite, cwp=None, main=None):
    self.cwp = cwp
    if modname is None:
      print ' package main'
      print ' import "os"'
      print ' import "runtime/pprof"'
    else:
      print ' package %s' % os.path.basename(modname)
    print ' import "fmt"'
    print ' import . "github.com/strickyak/rye/runt"'

    for th in suite.things:
      if type(th) == Timport:
        if True or not th.Go:
	  print ' import i_%s "%s"' % (th.alias, '/'.join(th.imported))

    print ' var _ = fmt.Sprintf'
    print ' var _ = MkInt'
    print ''
    print 'var G = New_Module()'

    print ' func Eval_Module() P {'
    for th in suite.things:
      th.visit(self)
    print '   return None'
    print ' }'
    print ''
    print '\n\n'.join(self.tail)
    print ''
    for i in range(MaxNumCallArgs + 1):
      print '  type i_%d interface { Call%d(%s) P }' % (i, i, ", ".join(i * ['P']))
    print ''
    print ' type Module struct {'
    print '    PModule'
    for g, (t, v) in sorted(self.glbls.items()):
      print '    M_%s %s' % (g, t)
    print ' }'
    print ''
    for g, (t, v) in sorted(self.glbls.items()):
      print 'var M_%s %s' % (g, t)
    print ''
    print ' func New_Module() *Module {'
    print '   G := new(Module)'
    print '   G.Self = G'
    print '   G.Init_PModule()'
    for g, (t, v) in sorted(self.glbls.items()):
      print '   M_%s = %s' % (g, v)
      print '   G.M_%s = M_%s' % (g, g)
      if len(v) > 4 and v[:4] == "new(":  #)
        print '   M_%s.SetSelf(M_%s)' % (g, g)
    print '   return G'
    print ' }'
    print ''

    for key, code in sorted(self.lits.items()):
      print 'var %s = %s' % (key, code)
    print ''

    for iv in sorted(self.gsNeeded):
      print ' type i_GET_%s interface { GET_%s() P }' % (iv, iv)
      print ' type i_SET_%s interface { SET_%s(P) }' % (iv, iv)

    if main:
      sys.stdout.close()
      sys.stdout = main
      print '''
 package main
 import "os"
 import "runtime/pprof"
 import "github.com/strickyak/rye/runt"
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

  def Vassign(self, p):
    # a, op, b
    # Resolve rhs first.
    rhs = p.b.visit(self)
    lhs = 'TODO'

    a = p.a
    if a.__class__ is Tfield:
      # p, field
      x = a.p.visit(self)
      if type(x) is ZSelf:  # Special optimization for self.
        self.instvars[a.field] = True
        lhs = 'self.M_%s' % a.field
      else:
        lhs = "%s.M_%s" % (x, a.field)

    elif a.__class__ is Tvar:
      # Are we in a function scope?
      if len(self.scopes):
        # Inside a function.
        scope = self.scopes[0]
        if scope.get(a.name):
          lhs = scope[a.name]
        else:
          lhs = scope[a.name] = 'v_%s' % a.name
      else:
        # At the module level.
        lhs = a.visit(self)
        self.glbls[a.name] = ('P', 'None')
    elif a.__class__ is Tgetitem:
        p = a.a.visit(self)
        q = a.x.visit(self)
        print '   (%s).SetItem(%s, %s)' % (p, q, rhs)
        return
    elif a.__class__ is Traw:
      lhs = a.raw
    else:
      raise Exception('Weird Assignment, a class is %s' % a.__class__.__name__)

    print '   %s = %s' % (lhs, rhs)

  def Vprint(self, p):
    vv = [a.visit(self) for a in p.xx.xx]
    print '   fmt.Println(%s.String())' % '.String(), '.join([str(v) for v in vv])

  def Vimport(self, p):
    im = '/'.join(p.imported)
    if self.glbls.get(p.alias):
      raise Exception("Import alias %s already used", p.alias)
    self.imports[p.alias] = p

    if p.Go:
      self.glbls[p.alias] = ('*PImport', 'GoImport("%s")' % im)
    else:
      self.glbls[p.alias] = ('*PImport', 'RyeImport("%s", i_%s.G)' % (im, p.alias))
      print '   if EvalRyeModuleOnce("%s") { i_%s.Eval_Module() } ' % (im, p.alias)

  def Vassert(self, p):
    print '   if ! P(%s).Bool() {' % p.x.visit(self)
    print '     panic("Assertion Failed:  %s ;  message=" + P(%s).String() )' % (
       p.code.encode('unicode_escape'), "None" if p.y is None else p.y.visit(self) )
    print '   }'

  def Vtry(self, p):
    print '''
   func() {
     defer func() {
       r := recover()
       if r != nil {
         // BEGIN EXCEPT
%s
         // END EXCEPT
         return
       }
     }()
     // BEGIN TRY
%s
     // END TRY
   }()
''' % (p.ex.visit(self), p.tr.visit(self))

  def Vfor(self, p):
    # Assign, for the side effect of var creation.
    Tassign(p.var, Traw('None')).visit(self)
    print '''
   func() {
     var i Nexter = %s.Iter().(Nexter)
     defer func() {
       r := recover()
       if r != nil {
         if r != G_StopIterationSingleton {
           panic(r)
         }
       }
     }()
     for {
       %s = i.Next()
       // BEGIN FOR
''' % (p.t.visit(self), p.var.visit(self))
    p.b.visit(self)
    print '''
       // END FOR
     }
   }()
'''

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
        print '   return Enlist( %s )' % ', '.join(vv)

  def Vbreak(self, p):
    print '   break'

  def Vcontinue(self, p):
    print '   continue'

  def Vraise(self, p):
    print '   panic( %s.String() )' % p.a.visit(self)

  def LitIntern(self, v, key, code):
    if not self.lits.get(key):
      self.lits[key] = code
    return ZLit(v, key)

  def Vlit(self, p):
    if p.k == 'N':
      v = p.v
      key = 'litNum_%s' % CleanIdentWithSkids(repr(v))
      code = 'MkInt(%s)' % v
    elif p.k == 'S':
      v = eval(p.v)
      key = 'litStr_%s' % p.v.encode('hex')
      code = 'MkStr("%s")' % v.encode('unicode_escape')
    else:
      Bad('Unknown Vlit', p.k, p.v)
    return self.LitIntern(v, key, code)

  def Vop(self, p):
    if p.returns_bool:
      return ' MkBool(%s.%s(%s)) ' % (p.a.visit(self), p.op, p.b.visit(self))
    if p.b:
      return ' (%s).%s(%s) ' % (p.a.visit(self), p.op, p.b.visit(self))
    else:
      raise Bad('Monadic %d not imp' % p.op)

  def Vboolop(self, p):
    if p.b is None:
      return ' MkBool( %s (%s).Bool()) ' % (p.op, p.a.visit(self))
    else:
      return ' MkBool(%s.Bool() %s %s.Bool()) ' % (p.a.visit(self), p.op, p.b.visit(self))

  def Vgetitem(self, p):
    return ' (%s).GetItem(%s) ' % (p.a.visit(self), p.x.visit(self))

  def Vgetitemslice(self, p):
    return ' (%s).GetItemSlice(%s, %s, %s) ' % (
        p.a.visit(self),
        None if p.x is None else p.x.visit(self),
        None if p.y is None else p.y.visit(self),
        None if p.z is None else p.z.visit(self))

  def Vtuple(self, p):
    return 'MkTupleV( %s )' % ', '.join([x.visit(self) for x in p.xx])

  def Vlist(self, p):
    return 'MkListV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vdict(self, p):
    return 'MkDictV( %s )' % ', '.join([str(x.visit(self)) for x in p.xx])

  def Vvar(self, p):
    if p.name == 'self':
      return ZSelf(p, 'self')
    for s in self.scopes:
      if p.name in s:
        return ZLocal(p, s[p.name])
    if p.name in BUILTINS:
      return ZBuiltin(p, 'B_%s' % p.name)
    return ZGlobal(p, 'M_%s' % p.name)

  def Vcall(self, p):
    # fn, args
    global MaxNumCallArgs
    n = len(p.args)
    MaxNumCallArgs = max(MaxNumCallArgs, n)
    args = ''
    for a in p.args:
      args += ' %s, ' % (a.visit(self))

    if type(p.fn) is Tfield and type(p.fn.p) is Tvar: 
      #todo# zholder = p.fn.p.visit(self)

      if p.fn.p.name in self.imports:

        imp = self.imports[p.fn.p.name]
        if imp.Go:
          return ' MkGo(i_%s.%s).Call(%s) ' % (p.fn.p.name, p.fn.field, args)
        else:
          return ' i_%s.M_%d_%s(%s) ' % (p.fn.p.name, n, p.fn.field, args)
        #return ' ((%s).FieldForCall("%s")).Call(%s) ' % (p.fn.p.visit(self), p.fn.field, args)

    zfn = p.fn.visit(self)
    if type(zfn) is ZBuiltin:
      return ' B_%d_%s(%s) ' % (n, zfn.t.name, args)

    if type(zfn) is ZGlobal:
      return ' M_%d_%s(%s) ' % (n, zfn.t.name, args)

    arglist = ', '.join(["(%s)" % (a.visit(self)) for a in p.args])
    return '/*NANDO*/ P(%s).(i_%d).Call%d(%s) ' % (p.fn.visit(self), n, n, arglist)

  def Vfield(self, p):
    # p, field
    x = p.p.visit(self)
    if type(x) is ZSelf and self.instvars.get(p.field):  # Special optimization for self instvars.
      return '%s.M_%s' % (x, p.field)
    else:
      self.gsNeeded[p.field] = True
      return ' P(%s).(i_GET_%s).GET_%s() ' % (x, p.field, p.field)

  def Vdef(self, p):
    # name, args, body.
    buf = PushPrint()

    # Tweak args.  Record meth, if meth.
    args = p.args
    if self.cls and len(p.args):
      if p.args[0] != 'self':
        Bad('first arg to method %s is %s; should be self', p.name, p.args[0])
      args = p.args[1:]  # Skip self.
      self.meths[p.name] = [ args, ]  # Could add more, but args will do.

    # prepend new scope dictionary, containing just the args, so far.
    self.scopes = [ dict([(a, 'a_%s' % a) for a in args]) ] + self.scopes

    #################
    # Render the body, but hold it in buf2, because we will prepend the vars.
    buf2 = PushPrint()
    p.body.visit(self)
    PopPrint()
    code2 = buf2.String()
    #################

    if self.cls:
      func = 'func (self *C_%s) M_%d_%s' % (self.cls, len(p.args)-1, p.name)
    else:
      func = 'func M_%d_%s' % (len(p.args), p.name)

    print ''
    print ' %s(%s) P {' % (func, ', '.join(['a_%s P' % a for a in args]))

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

    if self.cls:
      n = len(args)
      print ' type PMeth_%d_%s__%s struct { PBase; Rcvr *C_%s }' % (n, self.cls, p.name, self.cls)
      print ' func (o *PMeth_%d_%s__%s) Call%d(%s) P {' % (n, self.cls, p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '   return o.Rcvr.M_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
      print ' }'
      print ''

    else:
      n = len(p.args)
      print ' type pFunc_%s struct { PBase }' % p.name
      print ' func (o pFunc_%s) Call%d(%s) P {' % (p.name, n, ', '.join(['a%d P' % i for i in range(n)]))
      print '   return M_%d_%s(%s)' % (n, p.name, ', '.join(['a%d' % i for i in range(n)]))
      print ' }'
      print ''
      self.glbls[p.name] = ('*pFunc_%s' % p.name, 'new(pFunc_%s)' % p.name)

    PopPrint()
    code = buf.String()
    self.tail.append(code)

    # The class constructor gets the args of init:
    if self.cls and p.name == '__init__':
      self.args = p.args

  def Vclass(self, p):
    # name, sup, things
    self.cls = p.name
    sup = p.sup if p.sup else 'object'
    self.instvars = {}
    self.meths = {}
    self.args = []

    # Emit all the methods of the class (and possibly other members).
    for x in p.things:
      x.visit(self)
    self.cls = ''

    buf = PushPrint()

    # Emit the struct for the class.
    print '''
 type C_%s struct {
   C_%s
%s
 }
''' % (p.name, sup, '\n'.join(['   M_%s   P' % x for x in self.instvars]))

    print '''
 func (o *C_%s) PtrC_%s() *C_%s {
   return o
 }
''' % (p.name, p.name, p.name)

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
    for m in sorted(self.meths):
      args, = self.meths[m]
      n = len(args)
      print ' func (o *C_%s) GET_%s() P { z := &PMeth_%d_%s__%s { Rcvr: o }; z.SetSelf(z); return z }' % (p.name, m, n, p.name, m)

    # The constructor.
    n = len(self.args) - 1 # Subtract 1 because we don't count self.
    arglist = ', '.join(['a%d P' % i for i in range(n)])
    argpass = ', '.join(['a%d' % i for i in range(n)])
    print ' type pCtor_%d_%s struct { PBase }' % (n, p.name)
    print ''
    print ' func (o pCtor_%d_%s) Call%d(%s) P {' % (n, p.name, n, arglist)
    print '   return M_%d_%s(%s)' % (n, p.name, argpass)
    print ' }'
    print ''
    print ' func M_%d_%s(%s) P {' % (n, p.name, arglist)
    print '   z := new(C_%s)' % p.name
    print '   z.Self = z'
    for iv in self.instvars:
      print '   z.M_%s = None' % iv
    print '   z.M_%d___init__(%s)' % (n, argpass)
    print '   return z'
    print ' }'
    print ''
    print ''
    self.glbls[p.name] = ('*pCtor_%d_%s' % (n, p.name), 'new(pCtor_%d_%s)' % (n, p.name))

    self.tail.append(buf.String())
    PopPrint()

  def Vsuite(self, p):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
    for x in p.things:
      x.visit(self)

  def Vseq(self, p):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
    for x in p.things:
      x.visit(self)

PrintStack= []
def PushPrint():
    global PrintStack
    sys.stdout.flush()
    PrintStack.append(sys.stdout)
    buf = Buffer()
    sys.stdout = buf
    return buf
def PopPrint():
    global PrintStack
    sys.stdout = PrintStack.pop()

class Buffer(object):
  def __init__(self):
    self.b = []
  def write(self, x):
    self.b.append(x)
  def flush(self):
    pass
  def String(self):
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
    return v.Vlist(self)  # By default, make a list.

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

class Tdict(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vdict(self)

class Tsuite(Tnode):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
  def __init__(self, things):
    self.things = things
  def visit(self, v):
    return v.Vsuite(self)

class Tseq(Tnode):  # So far, Tsuite and Tseq and Vsuite and Vseq are the same.
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
  def __init__(self, a, b):
    self.a = a
    self.b = b
  def visit(self, v):
    return v.Vassign(self)

class Tprint(Tnode):
  def __init__(self, xx):
    self.xx = xx
  def visit(self, v):
    return v.Vprint(self)

class Timport(Tnode):
  def __init__(self, imported, alias, go):
    self.imported = imported
    self.alias = alias
    self.Go = go
  def visit(self, v):
    return v.Vimport(self)

class Tassert(Tnode):
  def __init__(self, x, y, code):
    self.x = x
    self.y = y
    self.code = code
  def visit(self, v):
    return v.Vassert(self)

class Ttry(Tnode):
  def __init__(self, tr, ex):
    self.tr = tr
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
  def __init__(self, var, t, b):
    self.var = var
    self.t = t
    self.b = b
  def visit(self, v):
    return v.Vfor(self)

class Treturn(Tnode):
  def __init__(self, aa):
    self.aa = aa
  def visit(self, v):
    return v.Vreturn(self)

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

class Tdef(Tnode):
  def __init__(self, name, args, body):
    self.name = name
    self.args = args
    self.body = body
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
  def __init__(self, fn, args):
    self.fn = fn
    self.args = args
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

  def Advance(self):
    self.p += 1
    if self.p >= len(self.words):
      self.k, self.v, self.i = None, None, len(self.program)
    else:
      self.k, self.v, self.i = self.words[self.p]

  def Rest(self):
    return self.program[self.i:]

  def MkTemp(self):
    z = Tvar(Serial('tmp'))
    return z

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

  def Xprim(self):
    #print '//<------ Nando: Xprim: ENTER at ( %s, %s )' % (repr(self.k), repr(self.v))
    if self.k == 'N':
      z = Tlit(self.k, self.v)
      self.Advance()
      #print '//------> Nando: Xprim: RECOGNIZED N: ', z
      return z

    if self.k == 'S':
      z = Tlit(self.k, self.v)
      self.Advance()
      #print '//------> Nando: Xprim: RECOGNIZED S: ', z
      return z

    if self.k == 'A':
      z = Tvar(self.v)
      self.Advance()
      #print '//------> Nando: Xprim: RECOGNIZED A: ', z
      return z

    if self.k == 'K':
      if self.v in ['None', 'True', 'False']:
        v = self.v
        self.Eat(self.v)
	z = Traw(v)
        #print '//------> Nando: Xprim: RECOGNIZED K: ', z
        return z
      raise Exception('Keyword "%s" is not an expression' % self.v)

    if self.v == '(':
      self.Eat('(')
      if self.v == ')':
        self.Eat(')')
        # Unit tuple.
        return self.Ttuple([])
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
      #print '//------> Nando: Xprim: RECOGNIZED Tuple Tlist of: ', z
      return Tlist(z)

    if self.v == '[':
      self.Eat('[')
      z = []
      while self.v != ']':
        x = self.Xexpr()
        z.append(x)
        if self.v == ']':
          # Omitted trailing ','
          break
        self.Eat(',')
      self.Eat(']')
      #print '//------> Nando: Xprim: RECOGNIZED List Tlist of: ', z
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
      #print '//------> Nando: Xprim: RECOGNIZED Tdict of: ', z
      return Tdict(z)

    else:
      #print '//------> Nando: Xprim: RECOGNIZED BAD'
      raise self.Bad('Expected Xprim, but got %s, at %s' % (self.v, repr(self.Rest())))

  def Xsuffix(self):
    """Tcall, Tfield, or Tindex"""
    #print '//------> Nando: Xsuffix: ENTER ' + repr((self.k, self.v))
    a = self.Xprim()
    #print '//------> Nando: Xsuffix: a= ' + repr(a)
    while True:
      if self.v == '(':
        self.Eat('(')
        args = []
        while self.v != ')':
          b = self.Xexpr()
          args.append(b)
          if self.v == ',':
            self.Eat(',')
          else:
            break
        self.Eat(')')
        a = Tcall(a, args)

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

      else:
        #print '//------> Nando: Xsuffix: break'
        break
    #print '//------> Nando: Xsuffix: return ' + repr(a)
    return a

  def Xmul(self):
    a = self.Xsuffix()
    while self.v in MUL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xsuffix()
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

  def Xrelop(self):
    a = self.Xadd()
    if self.v in REL_OPS:
      op = self.v
      self.Eat(op)
      b = self.Xadd()
      a = Top(a, REL_OPS[op], b, True)
    elif RE_WORDY_REL_OP.match(self.v):
      op = self.v
      self.Eat(op)
      b = self.Xadd()
      if op == 'in':
        a = Top(b, "Contains", a, True)    # N.B. swap a & b for Contains
      elif RE_NOT_IN.match(op):
        a = Top(b, "NotContains", a, True)    # N.B. swap a & b for NotContains
      elif op == 'is':
        a = Top(a, "Is", b, True)
      elif RE_IS_NOT.match(op):
        a = Top(b, "IsNot", a, True)
      else:
        raise Exception("Weird RE_WORDY_REL_OP: %s" % op)
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
    while self.v == 'and':
      op = self.v
      self.Eat(op)
      b = self.Xand()
      a = Tboolop(a, "||", b)
    return a

  def Xexpr(self):
    return self.Xor()

  def Xlistexpr(self):
    z = self.Xitems(allowScalar=True, allowEmpty=False)
    return z

  def Xitems(self, allowScalar, allowEmpty):
    "A list of expressions, possibly empty, or possibly a scalar."
    z = []
    had_comma = False
    while self.k != ';;' and self.v not in [')', ']', '}', ':', '=', '+=', '-=', '*=', '/=']:
      if self.v == ',':
        self.Eat(',')
        had_comma = True
      else:
        x = self.Xexpr()
        z.append(x)
    if allowScalar and len(z) == 1 and not had_comma:
      return z[0]  # Scalar.
    if not allowEmpty and len(z) == 0:
      raise Exception('Empty expression list not allowed')
    return Titems(z)  # List of items.

  def Csuite(self):
    things = []
    while self.k != 'OUT' and self.k is not None:
      #print '//Csuite', self.k, self.v
      if self.v == ';;':
        self.EatK(';;')
      else:
        t = self.Command();
        if t:
          things.append(t)
    return Tsuite(things)

  def Command(self):
    if self.v == 'print':
      return self.Cprint()
    elif self.v == 'if':
      return self.Cif()
    elif self.v == 'while':
      return self.Cwhile()
    elif self.v == 'for':
      return self.Cfor()
    elif self.v == 'return':
      return self.Creturn()
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
    elif self.v == 'import':
      return self.Cimport()
    elif self.v == 'go':
      return self.Cgo()
    elif self.v == 'try':
      return self.Ctry()
    elif self.v == 'pass':
      self.Eat('pass')
      return
    elif self.k == 'A':
      return self.Cother()
    else:
      raise self.Bad('Unknown stmt: %s %s %s', self.k, self.v, repr(self.Rest()))

  def Cother(self):
    #print '//Nando: Cother: enter ' + repr((self.k, self.v))
    a = self.Xitems(allowScalar=True, allowEmpty=False)  # lhs (unless not an assignment; then it's the only thing.)
    #print '//Nando: Cother: a= ' + repr(a)
    #print '//Nando: Cother: a= class ' + repr(a.__class__)

    if a.__class__ == Titems:
      xx = a.xx
      self.Eat('=')
      b = self.Xlistexpr()  # rhs
      tmp = self.MkTemp()
      things = [Tassign(tmp, b)]
      i = 0
      for x in xx:
        if x.__class__ is not Tvar or x.name != '_': 
          things.append(Tassign(x, Tgetitem(tmp, Tlit('N', i))))
        i += 1
      #print '//Nando: Cother: Titems -> Tseq ' + repr(things)
      return Tseq(things)

    #print '//Nando: Cother...a', repr(a)
    op = self.v
    #print '//Nando: Cother...op?', repr(op)

    if op in ['+=', '-=', '*=']:
      self.Eat(op)
      binop = op[:-1]  # Remove the '='
      b = self.Xexpr()
      #print '//Cother...op...b', op, b
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
      #print '//Cother...=...b', b
      return Tassign(a, b)
    else:
      # TODO: error if this is not a function or method call.
      #print '//Nando: Cother... ELSE --> ', repr(a)
      return Tassign(Traw('_'), a)

  def Cprint(self):
    # TODO: distinguish trailing ,
    self.Eat('print')
    t = self.Xitems(allowScalar=False, allowEmpty=True)
    self.EatK(';;')
    return Tprint(t)

  def Cgo(self):
    self.Eat('go')
    if self.v == 'import':
      return self.Cimport(go=True)
    raise Exception('go command: not yet implemented (except: go import...)')

  def Cimport(self, go=False):
    self.Eat('import')
    alias = self.v
    self.EatK('A')
    while self.v == '.':
      self.Eat('.')
      alias = '%s.%s' % (alias, self.v)
      self.EatK('A')
    imported = [ alias ]
    while self.v == '/':
      self.Eat('/')
      imported.append(self.v)
      alias = self.v
      self.EatK('A')
    if self.v == 'as':
      self.Eat('as')
      alias = self.v
      self.EatK('A')
    self.EatK(';;')

    return Timport(imported, alias, go)

  def Cassert(self):
    i = self.i
    self.Eat('assert')
    x = self.Xexpr()
    y = None
    j = self.i
    if self.v == ',':
      self.Eat(',')
      y = self.Xexpr()
    return Tassert(x, y, self.program[i:j])

  def Ctry(self):
    self.Eat('try')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    tr = self.Csuite()
    self.EatK('OUT')
    self.Eat('except')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    ex = self.Csuite()
    self.EatK('OUT')
    return Ttry(tr, ex)

  def Cif(self):
    self.Eat('if')
    t = self.Xexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    yes = self.Csuite()
    self.EatK('OUT')
    no = None
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
    if self.k != 'A':
      raise Exception('Got "%s" after for; expected varname', self.v)
    var = self.Xvar()
    self.Eat('in')
    t = self.Xlistexpr()
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    b = self.Csuite()
    self.EatK('OUT')
    return Tfor(var, t, b)

  def Creturn(self):
    self.Eat('return')
    if self.v == ';;':  # Missing Xitems means None, not [].
      return Treturn(None)
    t = self.Xlistexpr()
    return Treturn([t])

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
    t = self.Xexpr()
    return Traise(t)

  def Cclass(self):
    self.Eat('class')
    name = self.Pid()
    # No superclasses yet
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')

    things = []
    while self.k != 'OUT':
      if self.v == 'def':
        t = self.Cdef(name)
        things.append(t)
      elif self.v == 'pass':
        self.Eat('pass')
      elif self.k == ';;':
        self.EatK(';;')
      else:
        raise self.Bad('Classes may only contain "def" or "pass" commands.')
    self.EatK('OUT')

    return Tclass(name, None, things)

  def Cdef(self, cls):
    self.Eat('def')
    name = self.Pid()
    self.Eat('(')
    args = []
    while self.k == 'A':
      arg = self.Pid()
      if self.v == ',':
        self.Eat(',')
      args.append(arg)
    self.Eat(')')
    self.Eat(':')
    self.EatK(';;')
    self.EatK('IN')
    suite = self.Csuite()
    self.EatK('OUT')
    return Tdef(name, args, suite)

class Z(object):  # Returns from visits (emulated runtime value).
  def __init__(self, t, s):
    self.t = t  # T node
    self.s = s  # String for backwards compat
  def __str__(self):
    return self.s
class ZSelf(Z):
  pass
class ZLocal(Z):
  pass
class ZGlobal(Z):
  pass
class ZBuiltin(Z):
  pass
class ZLit(Z):
  pass

pass


# OPERATOR HIERARCHY OF PYTHON
#lambda	Lambda expression
#if else	Conditional expression
#or	Boolean OR
#and	Boolean AND
#not x	Boolean NOT
#in, not in, is, is not, <, <=, >, >=, <>, !=, ==	Comparisons, including membership tests and identity tests
#|	Bitwise OR
#^	Bitwise XOR
#&	Bitwise AND
#<<, >>	Shifts
#+, -	Addition and subtraction
#*, /, //, %	Multiplication, division, remainder [8]
#+x, -x, ~x	Positive, negative, bitwise NOT
#**	Exponentiation [9]
#x[index], x[index:index], x(arguments...), x.attribute	Subscription, slicing, call, attribute reference
#(expressions...), [expressions...], {key: value...}, `expressions...`	Binding or tuple display, list display, dictionary display, string conversion
