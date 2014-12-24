from . import tr

def main(args):
  for a in args:
    print '<<<', a
    z = Interpret(a + '\n')
    print '>>>', z
  print 'OKAY'

def Interpret(program):
  say program
  words = tr.Lex(program).tokens
  say words
  words = list(tr.SimplifyContinuedLines(words))
  say words
  parser = tr.Parser(program, words, -1, '<EVAL>')
  say parser
  tree = parser.Csuite()
  say tree
  walker = ShowExprWalker()
  say walker
  z = tree.visit(walker)
  say z
  return z

class ShowExprWalker:
  def __init__():
    pass

  def Vop(self, p):  # a, op, b=None, returns_bool
    say p
    z = '( OP%s %s %s %s )' % ('-bool' if p.returns_bool else '', p.a.visit(self), p.op, p.b.visit(self) if p.b else None)
    say z
    return z

  def Vboolop(self, p):  # a, op, b=None
    say p
    z = '( BOOLOP %s %s %s )' % (p.a.visit(self), p.op, p.b.visit(self) if p.b else None)
    say z
    return z

  def Vcondop(self, p):  # a, b, c
    say p
    z = '( CONDOP %s IF %s ELSE %s )' % (p.a.visit(self), p.b.visit(self), p.c.visit(self))
    say z
    return z

  def Vgo(self, p):  # fcall
    say p
    z = '( GO-CALL %s )' % (p.fcall.visit(self), )
    say z
    return z

  def Vraw(self, p):  # raw
    say p
    z = '( RAW {{{%s}}} )' % (p.raw, )
    say z
    return z

  def Vlit(self, p):  # k, v
    say p
    z = '( LIT %s {{{%s}}} )' % (p.k, p.v, )
    say z
    return z

  def Vvar(self, p):  # name
    say p
    z = '( VAR %s )' % (p.name, )
    say z
    return z

  def Vitems(self, p):  # xx, trailing_comma
    say p
    z = '( ITEMS[%d] %s )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vtuple(self, p):  # xx
    say p
    z = '( TUPLE[%d] %s )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  def Vlist(self, p):  # xx
    say p
    z = '( LIST[%d] %s )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
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
    z = '( DICT[%d] %s )' % (len(p.xx), (', '.join([x.visit(self) for x in p.xx])))
    say z
    return z

  # STATEMENTS

  def Vsuite(self, p):  # Statement.  things
    z = '( SUITE[%d] %s )' % (len(p.things), (', '.join([x.visit(self) for x in p.things])))
    return z

  def Vexpr(self, p):  # Statement.  a
    z = '( EXPR %s )' % (p.a.visit(self), )
    return z

  def Vassign(self, p):  # Statement.  a, b, pragma.
    z = '( ASSIGN %s %s )' % (p.a.visit(self), p.b.visit(self), )
    return z


  def Vprint(self, p):  # Statement.  w, xx, saying, code
    #z = '( PRINT %s )' % ([x.visit(self) for x in p.xx.xx], )
    z = '( PRINT %s )' % (p.xx.visit(self), )
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
  def Vcall(self, p):  # for defer or go.
    raise 'Statement Not Implemented'
  def Vfield(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vgetitem(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vgetitemslice(self, p):  # Statement.
    raise 'Statement Not Implemented'
  def Vcurlysetter(self, p):  # Statement.
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
#1829 class Tcall(Tnode):
#1839 class Tfield(Tnode):
#1846 class Tgetitem(Tnode):
#1853 class Tgetitemslice(Tnode):
#1862 class Tcurlysetter(Tnode):
