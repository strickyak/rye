"""Visitor to extract Doumentation."""
import re, sys

rye_rye = False
if rye_rye:
  from . import lex, parse
  from rye_lib import data
else:
  import lex, parse

PRIVATE = re.compile('_[^_]').match

class DocVisitor:
  """Extract documentation from Parse Tree."""
  def __init__(self):
    self.funcs = {}
    self.classes = {}

  def Vop(self, p):  # a, op, b=None, returns_bool
    pass

  def Vboolop(self, p):  # a, op, b=None
    pass

  def Vcondop(self, p):  # a, b, c # b if a else c
    pass

  def Vgo(self, p):  # fcall
    pass

  def Vraw(self, p):  # raw
    if p.raw == '_':
      return p.raw

  def Vlit(self, p):  # k, v
    if p.k == 'N':
      return int(p.v)
    elif p.k == 'F':
      return float(p.v)
    elif p.k == 'S':
      return p.v
      if rye_rye:
        return data.Eval(p.v)
      else:
        return eval(p.v)
    else:
      raise 'Weird case', p.k

  def Vvar(self, p):  # name
    return p.name

  def Vitems(self, p):  # xx, trailing_comma
    pass

  def Vtuple(self, p):  # xx
    pass

  def Vlist(self, p):  # xx
    pass

  def Vlambda(self, p):  # lvars, expr, where, line
    pass

  def Vforexpr(self, p):  # z(*body*), vv(*vars*), ll(*list*), cond, has_comma
    pass

  def Vdict(self, p):  # xx
    pass

  def Vcall(self, p):  # fn, args, names, star, starstar
    pass

  def Vfield(self, p):  # p, field
    pass

  def Vgetitem(self, p):  # a, x
    pass

  def Vgetitemslice(self, p):  # a, x, y, (z)
    pass

  # STATEMENTS

  def Vsuite(self, p):  # Statement.  things
    if p.things:
      return p.things[0].visit(self)

  def Vexpr(self, p):  # Statement.  a
    z = p.a.visit(self)

  def Vassign(self, p):  # Statement.  a, b, pragma.
    # A bare literal comment is parsed as an assignment to _.
    # For our purpose, we can return the literal if assigned to _,
    # or just return None.
    if p.a.visit(self) == '_':
      return p.b.visit(self)

  def Vprint(self, p):  # Statement.  w, xx, saying, code
    pass

  def Vwithdefer(self, p):  # Statement.
    pass

  def Vglobal(self, p):  # Statement.
    pass

  def Vimport(self, p):  # Statement.
    pass

  def Vassert(self, p):  # Statement.
    pass

  def Vtry(self, p):  # Statement. tr exvar ex
    pass

  def Vif(self, p):  # Statement.  t, yes, no.
    pass

  def Vwhile(self, p):  # Statement. t, yes.
    pass

  def Vfor(self, p):  # Statement. var, t, b.
    pass

  def Vreturn(self, p):  # Statement.
    pass

  def Vyield(self, p):  # Statement.
    pass

  def Vbreak(self, p):  # Statement.
    pass

  def Vcontinue(self, p):  # Statement.
    pass

  def Vraise(self, p):  # Statement.
    pass

  def Vdel(self, p):  # Statement.
    pass

  def Vnative(self, p):  # Statement.
    pass

  def Vdef(self, p):  # Statement.  name, args, dflts, star, starstar, body.
    # Body should be Tsuite.
    if PRIVATE(p.name): return
    dflts = [('=' + str(x.visit(self)) if x else '') for x in p.dflts]

    siglist = ', '.join(['%s%s' % (k, v) for k, v in zip(p.args, dflts)])
    if p.star:
      siglist += '%s*%s' % ((', ' if siglist else ''), p.star)
    if p.starstar:
      siglist += '%s**%s' % ((', ' if siglist else ''), p.starstar)
    signature = 'def %s(%s)' % (p.name, siglist)

    self.funcs[p.name] = dict(
        #type='func',
        Name=p.name,
        #args=p.args, dflts=dflts, star=p.star, starstar=p.starstar,
        Remark=self.TopRemark(p.body.things),
        #siglist=siglist,
        Signature=signature,
        )

  def Vclass(self, p):  # Tclass: name, sup, things.
    if PRIVATE(p.name): return

    stash = self.funcs  # Stash
    self.funcs = {}

    for th in p.things:
      th.visit(self)

    self.classes[p.name] = dict(
        ClassName=p.name,
        Super=p.sup,
        Methods=self.funcs,
        Remark=self.TopRemark(p.things),
        )

    self.funcs = stash  # Restore

  def TopRemark(self, things):
    "Given a list of things, if the first thing is a remark, return it; else None."
    if things:
      t0 = things[0]
      if type(t0) == parse.Tassign:
        return t0.visit(self)


def _ExtractDocumetationFromTree(tree):
  "Given a parse tree, return all the remarks and prototypes (as jsonic dicts)."
  dv = DocVisitor()
  # Build dv.funcs & dv.classes
  for th in tree.things:
    th.visit(dv)

  return dict(Remark=dv.TopRemark(tree.things), Funcs=dv.funcs, Classes=dv.classes)

def ExtractDocumetationFromFile(ryefile):
  program = open(ryefile).read()
  words = lex.Lex(program, filename=ryefile).tokens
  words = list(lex.SimplifyContinuedLines(words, filename=ryefile))
  parser = parse.Parser(program, words, -1, 'BUILTINS')

  try:
    tree = parser.Csuite()
  except Exception as err:
    print lex.AddWhereInProgram(str(err), len(program) - len(parser.Rest()), filename=ryefile)
    sys.exit(13)

  return _ExtractDocumetationFromTree(tree)

def Esc(s):
  'Escape string for HTML.  Also handles None.'
  if not s: return ''
  s = str(s)
  s = s.replace('&', '&amp;')
  s = s.replace('<', '&lt;')
  s = s.replace('>', '&gt;')
  s = s.replace('"', '&quot;')
  return s

def SuperName(p):
  if type(p) is parse.Tvar:
    return p.name
  return str(p)

def GenerateHtmlForFunctions(funcs):
  print '<table width="100%" border="1" cellpadding="10">'
  for fn in sorted(funcs.values()):
    print '''
      <tr>
        <td><b>%s</b>
        <td><tt>%s</tt>
        <td>%s
      ''' % (fn['Name'], Esc(fn.get('Signature')), Esc(fn.get('Remark')))
  print '</table>'


def GenerateHtmlFromFile(ryefile):
  d = ExtractDocumetationFromFile(ryefile)
  print '<html><head><title>%s</title></head><body>' % ryefile
  print '<h2>%s</h2>' % ryefile
  if d.get('Remark'):
    print '<pre>%s</pre>' % d['Remark']
  if d.get('Funcs'):
    print '<h3>Functions</h3>'
    GenerateHtmlForFunctions(d['Funcs'])
  if d.get('Classes'):
    for cl in sorted(d['Classes'].values()):
      print '<h3>class %s(%s)</h3>' % (cl['ClassName'], Esc(SuperName(cl.get('Super'))))
      GenerateHtmlForFunctions(cl['Methods'])


def main(args):
  z = dict([(ryefile, ExtractDocumetationFromFile(ryefile)) for ryefile in args])

  if rye_rye:
    print >> sys.stderr, data.PrettyPrint(z)
  else:
    print >> sys.stderr, z

  GenerateHtmlFromFile(args[0])

if not rye_rye:
  if __name__ == '__main__':
    main(sys.argv[1:])
