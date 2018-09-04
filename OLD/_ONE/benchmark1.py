import sys

rye_rye = False
if rye_rye:
  from . import lex, parse, codegen
else:
  import lex, parse, codegen

def ExercizeTranslater(n):
  Stuff = dict(), dict(), dict(), dict()
  program = SOURCE
  words = lex.Lex(program).tokens
  parser = parse.Parser(program, words, -1, 'bogus')
  tree = parser.Csuite()
  for i in range(n):
    gen = codegen.CodeGen()
    gen.InjectForInternal(Stuff)
    gen.GenModule('bogus', 'bogus', tree, 'bogus', internal='bogus')

def run():
  n = int(sys.argv[1]) if len(sys.argv) > 1 else 500
  ExercizeTranslater(n)

SOURCE = """
from go import strings, unicode
from go import bufio, io, io/ioutil, os

def XException(x):
  return x

def Xgo_deref(x):
  native:
    'return GoDeref(a_x)'

def Xgo_wrap(x):
  native:
    'return MkValue(reflect.ValueOf(a_x.Self.Contents()))'

def Xgo_typeof(x):
  native:
    'return MkGo(reflect.ValueOf(a_x.Self.Contents()).Type())'

def Xgo_kindof(x):
  native:
    'return MkStr(reflect.ValueOf(a_x.Self.Contents()).Type().Kind().String())'

def Xgo_valueof(x):
  native:
    'return MkValue(reflect.ValueOf(reflect.ValueOf(a_x.Self.Contents())))'

def Xrye_what(x):
  native:
    'return N_rye_what(a_x)'

def Xcallable(x):
  native:
    'return MkBool(a_x.Self.Callable())'

#def Xglobals():
  # TODO -- this is not going to work.
  #z = {}
  #native:
  #  'for k, ptr := range ModuleObj().Map() {'
  #  '  v_z.SetItem(MkStr(k), *ptr)'
  #  '}'
  #return z

def Xid(x):
  native:
    'return MkInt(int64(reflect.ValueOf(a_x).Pointer()))'

def Xhash(x):
  native:
    'return MkInt(a_x.Self.Hash())'

def Xgetattr(x, name, *dflt):
  n = len(dflt)
  if n:
    try:
      native:
        'return a_x.Self.FetchField(a_name.Self.String())'
    except:
      return dflt[0]
  else:
    native:
      'return a_x.Self.FetchField(a_name.Self.String())'

def Xsetattr(x, name, val):
  native:
    'a_x.Self.StoreField(a_name.Self.String(), a_val)'


def Xisinstance(obj, cls):
  native:
    'return MkBool(IsSubclass(a_obj.Self.PType(), a_cls))'

def Xissubclass(subcls, cls):
  native:
    'return MkBool(IsSubclass(a_subcls, a_cls))'

def Xord(x):
  native:
    'return Mkint(int(a_x.Self.String()[0]))'

def Xchr(x):
  native:
    'return MkStr(string([]byte{byte(a_x.Self.Int())}))'

def Xsum(vec, start=0):
  z = start
  for x in vec:
    z += x
  return z

def Xany(vec):
  for e in vec:
    if e:
      return True
  return False

def Xall(vec):
  for e in vec:
    if not e:
      return False
  return True

def Xlen(x):
  native:
    'return Mkint(a_x.Self.Len())'

def Xrepr(x):
  native:
    'return MkStr(a_x.Self.Repr())'

def Xstr(x):
  native:
    'return MkStr(a_x.Self.String())'

def Xint(x):
  native:
    'return MkInt(a_x.Self.ForceInt())'

def Xfloat(x):
  native:
    'return MkFloat(a_x.Self.ForceFloat())'

def Xrange(x):
  native:
    'return N_range(a_x)'

def Xxrange(x):
  i = 0
  while i < x:
    yield i
    i += 1

def Xsorted(x, cmp=None, key=None, reverse=False):
  native:
    'return N_sorted(a_x, a_cmp, a_key, a_reverse)'

def Xlist(x):
  native:
    'return MkList(a_x.Self.List())'

def Xdict(*vec, **kv):
  native:
    'return N_dict(a_vec, a_kv)'

def Xtuple(x):
  native:
    'return MkTuple(a_x.Self.List())'

def Xbool(x):
  native:
    'return MkBool(a_x.Self.Bool())'

def Xtype(x):
  native:
    'return a_x.Self.PType()'

def Xbyt(x):
  native:
    'return N_byt(a_x)'

def Xmkbyt(n):
  native:
    'return N_mkbyt(a_n)'

def Xrye_pickle(x):
  native:
    'return MkByt(Pickle(a_x))'

def Xrye_unpickle(x):
  native:
    'return UnPickle(a_x.Self.Bytes())'

def Xmax(*args):
  if len(args) == 0:
    raise 'no args to max()'
  if len(args) == 1:
    v = args[0]
    z = v[0]
    for e in v[1:]:
      if e > z:
        z = e
    return z
  else:
    z = args[0]
    for e in args[1:]:
      if e > z:
        z = e
    return z

def Xmin(*args):
  if len(args) == 0:
    raise 'no args to min()'
  if len(args) == 1:
    v = args[0]
    z = v[0]
    for e in v[1:]:
      if e < z:
        z = e
    return z
  else:
    z = args[0]
    for e in args[1:]:
      if e < z:
        z = e
    return z

def Xzip(*args):
  n = min([len(a) for a in args])
  return [tuple([a[i] for a in args]) for i in range(n)]

def Xzip_padding_with_None(*args):
  m = max([len(a) for a in args])
  return [tuple([a[i] if i < len(a) else None for a in args]) for i in range(m)]

def Xmap(fn, *lists):
  switch len(lists):
    case 0:
      raise 'map called with no lists'
    case 1:
      return [fn(x) for x in lists[0]]
    default:
      # N.B. Behaves like zip, truncating longer lists.
      return [fn(*tuple) for tuple in zip_padding_with_None(*lists)]

def Xreduce(fn, vec, init=None):
  vec = list(vec)
  if init is None:
    a = vec.pop(0)
  else:
    a = init
  for e in vec:
    a = fn(a, e)
  return a

class XPList(native):
  def Xappend(x):
    native:
      'self.PP = append(self.PP, a_x)'

  def Xextend(x):
    native:
      'self.PP = append(self.PP, a_x.Self.List()...)'

  def Xcount(x):
    z = 0
    for e in self:
      if e == x:
        z += 1
    return z

  def Xindex(x):
    i = 0
    for e in self:
      if e == x:
        return i
      i += 1
    raise 'ValueError'

  def Xremove(x):
    del self[self.index(x)]

  def Xinsert(i, x):
    # Tgetitemslice not supported yet ### self[i:i] = [x]
    z = self[:i] + [x] + self[i:]
    native:
      'self.PP = v_z.Self.List()'

  def Xpop(i = -1):
    x = self[i]
    del self[i]
    return x

  def Xreverse():
    n = len(self)
    say n, self
    for i in range(int(n/2)):
      say i, n-i-1, self[i], self[n-i-1]
      self[i], self[n-i-1] = self[n-i-1], self[i]

  def Xsort(cmp=None, key=None, reverse=False):
    native:
      'self.PP = N_sorted(&self.PBase, a_cmp, a_key, a_reverse).Self.List()'

  def Xcopy():
    native: '''
      var zz []B
      for _, e := range self.PP { zz = append(zz, e) }
      return MkList(zz)
    '''

class XPDict(native):
  def Xclear():
    native:
      'self.mu.Lock()'
      'self.ppp = make(map[string]B)'
      'self.mu.Unlock()'

  def Xcopy():
    native:
      'z := make(map[string]B)'
      'self.mu.Lock()'
      'for k, v := range self.ppp { z[k] = v }'
      'self.mu.Unlock()'
      'return MkDict(z)'

  def Xitems():
    native:
      'z := make([]B, 0, len(self.ppp))'
      'self.mu.Lock()'
      'for k, v := range self.ppp { z = append(z, MkTuple([]B{MkStr(k), v})) }'
      'self.mu.Unlock()'
      'return MkList(z)'
  def Xiteritems():
    return .items()

  def Xkeys():
    native:
      'z := make([]B, 0, len(self.ppp))'
      'self.mu.Lock()'
      'for k, _ := range self.ppp { z = append(z, MkStr(k)) }'
      'self.mu.Unlock()'
      'return MkList(z)'
  def Xiterkeys():
    return .keys()
  def Xiter():
    return .keys()

  def Xvalues():
    native:
      'z := make([]B, 0, len(self.ppp))'
      'self.mu.Lock()'
      'for _, v := range self.ppp { z = append(z, v) }'
      'self.mu.Unlock()'
      'return MkList(z)'
  def Xitervalues():
    return .values()

  def Xget(key, default = None):
    native:
      'k := a_key.Self.String()'
      'self.mu.Lock()'
      'z, ok := self.ppp[k]'
      'self.mu.Unlock()'
      'if ok { return z }'
      'return a_default'

  def Xhas_key(key):
    return (key in self)

  def Xsetdefault(key, default=None):
    if key in self:
      return self[key]
    else:
      self[key] = default
      return default

  def Xupdate(x):
    # TODO -- atomic update.
    stuff = dict(x).items()
    for k, v in stuff:
      self[k] = v


class XPStr(native):
  def Xsplit(x = None, n = -1):
    if x is None:
      native:
        '''
        s := self.S
        n := int(a_n.Self.Int())
        var v []string
        for n < 0 || len(v) < n {
          if len(s) == 0 { break }
          i := i_strings.IndexAny(s, " \\t\\n\\r")
          //println("n", n, "i", i, "s", s, "v", v)
          if i >= 0 { if i>0 {v = append(v, s[:i])}; s = s[i+1:]
          } else { if len(s) > 0 { v = append(v, s); break }}
        }
        return MkStrs(v)
        '''
    if n < 0:
      native:
        'return MkStrs(i_strings.Split(self.S, a_x.Self.String()))'
    else:
      native:
        'return MkStrs(i_strings.SplitN(self.S, a_x.Self.String(), 1 + int(a_n.Self.Int())))'

  def Xjoin(vec):
    native:
      'ss := make([]string, a_vec.Self.Len())'
      'for i, p := range a_vec.Self.List() {'
      '  ss[i] = p.Self.String()'
      '}'
      'return MkStr(i_strings.Join(ss, self.S))'

  def Xlower():
    native:
      'return MkStr(i_strings.ToLower(self.S))'

  def Xtitle():
    native:
      'return MkStr(i_strings.ToTitle(self.S))'

  def Xupper():
    native:
      'return MkStr(i_strings.ToUpper(self.S))'

  def Xendswith(x):
    native:
      'return MkBool(i_strings.HasSuffix(self.S, a_x.Self.String()))'

  def Xstartswith(x):
    native:
      'return MkBool(i_strings.HasPrefix(self.S, a_x.Self.String()))'

  def Xstrip(x=' \\t\\n\\r'):
    native:
      'return MkStr(i_strings.Trim(self.S, a_x.Self.String()))'

  def Xlstrip(x=' \\t\\n\\r'):
    native:
      'return MkStr(i_strings.TrimLeft(self.S, a_x.Self.String()))'

  def Xrstrip(x=' \\t\\n\\r'):
    native:
      'return MkStr(i_strings.TrimRight(self.S, a_x.Self.String()))'

  def Xreplace(old, new, count = -1):
    native:
      'return MkStr(i_strings.Replace(self.S, a_old.Self.String(), a_new.Self.String(), int(a_count.Self.Int())))'

  def Xfind(x):
    native:
      'return Mkint(i_strings.Index(self.S, a_x.Self.String()))'

  def Xrfind(x):
    native:
      'return Mkint(i_strings.LastIndex(self.S, a_x.Self.String()))'

  def Xindex(x):
    z = self.find(x)
    if z < 0:
      raise 'ValueError'
    return z

  def Xrindex(x):
    z = self.rfind(x)
    if z < 0:
      raise 'ValueError'
    return z

  def Xisalpha():
    if self:
      for c in self:
        if not unicode.IsLetter(ord(c)):
          return False
      return True
    else:
      return False

  def Xisdigit():
    if self:
      for c in self:
        if not unicode.IsDigit(ord(c)):
          return False
      return True
    else:
      return False

  def Xisalnum():
    if self:
      for c in self:
        if (not unicode.IsDigit(ord(c))) and (not unicode.IsLetter(ord(c))):
          return False
      return True
    else:
      return False

  def Xislower():
    if self:
      for c in self:
        if not unicode.IsLower(ord(c)):
          return False
      return True
    else:
      return False

  def Xisupper():
    if self:
      for c in self:
        if not unicode.IsUpper(ord(c)):
          return False
      return True
    else:
      return False

  def Xisspace():
    if self:
      for c in self:
        if not unicode.IsSpace(ord(c)):
          return False
      return True
    else:
      return False

def Xobject():
  native: '''
    z := &C_object{}
    z.Self = z
    return &z.PBase
  '''

class XC_object(native):
  # Defining __init__ at C_object lets you avoid it elsewhere.
  def X__init__():
    pass
  def X__getattr__(field):
    native:
      '''return FetchFieldByNameForObject(reflect.ValueOf(self.Self), a_field.Self.String())'''
  def X__setattr__(field, value):
    native:
      '''StoreFieldByNameForObject(reflect.ValueOf(self.Self), a_field.Self.String(), a_value)'''

class XC_promise(native):

  def XWait():
    native:
      'return self.Wait()'

def Xrye_chan(size, revSize=-1):
  native:
    'return make_rye_chan(a_size.Self.Int(), a_revSize.Self.Int())'

class XC_rye_chan(native):

  def XThrow(e):
    native:
      'self.Chan <- Either{Left: a_e, Right: nil}'

  def XSend(a):
    native:
      'self.Chan <- Either{Left: nil, Right: a_a}'

  def XRecv():
    native:
      '''
        z := <-self.Chan
        if z.Left != nil {
          panic(z.Left)
        } else if z.Right != nil {
          return z.Right
        }
      '''

  def XTryRecv():
    native:
      '''
        var z Either
        select {
          case z = <-self.Chan:
          default:
        }
        if z.Left != nil {
          panic(z.Left)
        } else if z.Right != nil {
          return z.Right
        } else {
          return None  // if channel was closed.
        }
      '''

  def XClose():
    native:
      'close(self.Chan)'

def Xopen(filename, mode='r'):
  if mode == 'r':
    return PYE_FileDesc(os.Open(filename), False)
  elif mode == 'w':
    return PYE_FileDesc(os.Create(filename), True)
  else:
    raise 'open: Unknow mode', mode

class XPYE_FileDesc:
  def X__init__(fd, writing):
    .writing = writing
    if writing:
      .f = fd
      .b = bufio.NewWriter(fd)
    else:
      .f = fd
      .b = bufio.NewReader(fd)

  def Xread():
    return str(ioutil.ReadAll(.b))

  def Xwrite(x):
    .b.Write(str(x))

  def Xflush():
    .b.Flush()
  def XFlush():
    .b.Flush()

  def Xclose():
    if .writing:
      .b.Flush()
    .f.Close()
  def XClose():
    if .writing:
      .b.Flush()
    .f.Close()

native:
  'func (self *C_PYE_FileDesc) Write(p []byte) (n int, err error) {'
  '  return self.M_b.Self.Contents().(io.Writer).Write(p)'
  '}'
  'func (self *C_PYE_FileDesc) Flush() error {'
  '  self.M_0_Flush()'
  '  return nil'
  '}'

pass
"""

if __name__ == '__main__':
  run()
