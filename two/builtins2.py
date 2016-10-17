from go import bytes, strings, unicode
from go import bufio, io, io/ioutil, os

## MACRO go_type(t) -- creates a reflective Value of the go type t.
#
## MACRO go_new(t) -- creates a new zeroed instance of the go type t and returns reflective pointer to it.
#
## MACRO go_cast(t, x) -- Casts the value x to the go type t.
#
## Newer Python requires this, but for rye, it is the identity function.
#def Exception(x):
#  return x
#
#def go_deref(x):
#  native:
#    'return GoDeref(a_x)'
#
#def go_wrap(x):
#  native:
#    'return MkValue(reflect.ValueOf(a_x.Contents()))'
#
#def go_typeof(x):
#  native:
#    'return MkGo(reflect.ValueOf(a_x.Contents()).Type())'
#
#def go_kindof(x):
#  native:
#    'return MkStr(reflect.ValueOf(a_x.Contents()).Type().Kind().String())'
#
#def go_valueof(x):
#  native:
#    'return MkValue(reflect.ValueOf(reflect.ValueOf(a_x.Contents())))'
#
#def rye_what(x):
#  native:
#    'return N_rye_what(a_x)'
#
#def callable(x):
#  native:
#    'return MkBool(a_x.Callable())'
#
#def setattrs(obj, **kw):
#  for k, v in kw.items():
#    setattr(obj, k, v)
#  return obj
#
##def globals():
#  # TODO -- this is not going to work.
#  #z = {}
#  #native:
#  #  'for k, ptr := range ModuleObj().Map() {'
#  #  '  v_z.SetItem(MkStr(k), *ptr)'
#  #  '}'
#  #return z
#
#def id(x):
#  native: `
#    if a_x.X == nil {
#      if len(a_x.S) == 0 {
#        panic("Rye does not support id() on int")
#      } else {
#        panic("Rye does not support id() on str")
#      }
#    }
#    return MkInt(int64(reflect.ValueOf(a_x.X.B()).Pointer()))
#  `
#
#def hash(x):
#  native:
#    'return MkInt(a_x.Hash())'
#
#def cmp(x, y):
#  native:
#    'return Mkint(a_x.Compare(a_y))'
#
#def getattr(x, name, *dflt):
#  n = len(dflt)
#  if n:
#    try:
#      native:
#        'return a_x.FetchField(a_name.String())'
#    except:
#      return dflt[0]
#  else:
#    native:
#      'return a_x.FetchField(a_name.String())'
#
#def setattr(x, name, val):
#  native:
#    'a_x.StoreField(a_name.String(), a_val)'
#
#
#def isinstance(obj, cls):
#  native:
#    'return MkBool(IsSubclass(a_obj.PType(), a_cls))'
#
#def issubclass(subcls, cls):
#  native:
#    'return MkBool(IsSubclass(a_subcls, a_cls))'

def ord(x):
  native:
    #'return macro.MkintJ(int(StringJ(a_x_1, a_x_2)[0]))'
    `t1 := StringJ(a_x_1, a_x_2)`
    `t2 := int(t1[0])`
    `return macro.MkintJ(t2)`

def chr(x):
  native:
    #'s := string([]byte{byte(macro.TakeIntJ(a_x_1, a_x_2))})'
    #'return macro.MkStrJ(s)'
    `t1 := macro.TakeIntJ(a_x_1, a_x_2)`
    `s := string([]byte{byte(t1)})`
    `return macro.MkStrJ(s)`

def sum(vec, start=0):
  z = start
  for x in vec:
    z += x
  return z

def any(vec):
  for e in vec:
    if e:
      return True
  return False

def all(vec):
  for e in vec:
    if not e:
      return False
  return True

def len(x):
  native:
    '__q := LenJ(a_x_1, a_x_2)'
    'return macro.MkintJ(__q)'

def repr(x):
  native:
    'z := ReprJ(a_x_1, a_x_2)'
    'return macro.MkStrJ(z)'

def str(x):
  native:
    'z := StrJ(a_x_1, a_x_2)'
    'return macro.MkStrJ(z)'

def int(x):
  native:
    'z := ForceIntJ(a_x_1, a_x_2)'
    'return macro.MkIntJ(z)'

#def float(x):
#  native:
#    'return MkFloat(a_x.ForceFloat())'
#
#def range(x):
#  native:
#    'return N_range(a_x)'
#
#def xrange(x):
#  i = 0
#  while i < x:
#    yield i
#    i += 1
#
#def sorted(x, cmp=None, key=None, reverse=False):
#  native:
#    'return N_sorted(a_x, a_cmp, a_key, a_reverse)'
#

def list(x):
  native:
    'z := MkListJ(ListJ(a_x_1, a_x_2))'
    'return inline.MkPJ(z)'

#def set(a=None):
#  native:
#    'return N_set(a_a)'
#
def dict(*args, **kv):
  native:
    'return N2_dict(a_args_1, a_args_2, a_kv_1, a_kv_2)'
#
def tuple(x):
  native: `
    x := inline.TakePJ(a_x_1, a_x_2)
    l := x.List()
    t := MkTupleJ(l)
    return inline.MkPJ(t)
  `
#
#def bool(x):
#  native:
#    'return MkBool(a_x.Bool())'
#
def type(x):
  native: `
    switch inline.Tag(a_x_1, a_x_2) {
    case Int:
      //a := inline.TakeIntJ(a_x_1, a_x_2)
      return H_int_1, H_int_2
    case Str:
      //a := inline.TakeStrJ(a_x_1, a_x_2)
      return H_str_1, H_str_2
    case Py:
      a := inline.TakePJ(a_x_1, a_x_2)
      return a.PType()
    }
    panic("Bad")
  `

#def byt(x):
#  native:
#    'return N_byt(a_x)'
#
#def mkbyt(n):
#  native:
#    'return N_mkbyt(a_n)'
#
#def rye_stack():
#  native:
#    'return MkStr(RyeStack())'
#
#def rye_pickle(x):
#  native:
#    'return MkByt(Pickle(a_x))'
#
#def rye_unpickle(x):
#  native:
#    'return UnPickle(a_x.Bytes())'
#
#def max(*args):
#  if len(args) == 0:
#    raise 'no args to max()'
#  if len(args) == 1:
#    v = args[0]
#    z = v[0]
#    for e in v[1:]:
#      if e > z:
#        z = e
#    return z
#  else:
#    z = args[0]
#    for e in args[1:]:
#      if e > z:
#        z = e
#    return z
#
#def min(*args):
#  if len(args) == 0:
#    raise 'no args to min()'
#  if len(args) == 1:
#    v = args[0]
#    z = v[0]
#    for e in v[1:]:
#      if e < z:
#        z = e
#    return z
#  else:
#    z = args[0]
#    for e in args[1:]:
#      if e < z:
#        z = e
#    return z
#
#def zip(*args):
#  n = min([len(a) for a in args])
#  return [tuple([a[i] for a in args]) for i in range(n)]
#
#def zip_padding_with_None(*args):
#  m = max([len(a) for a in args])
#  return [tuple([a[i] if i < len(a) else None for a in args]) for i in range(m)]
#
#def map(fn, *lists):
#  switch len(lists):
#    case 0:
#      raise 'map called with no lists'
#    case 1:
#      return [fn(x) for x in lists[0]]
#    default:
#      # N.B. Behaves like zip, truncating longer lists.  TODO: Really?
#      return [fn(*tuple) for tuple in zip_padding_with_None(*lists)]
#
#def reduce(fn, vec, init=None):
#  vec = list(vec)
#  if init is None:
#    a = vec.pop(0)
#  else:
#    a = init
#  for e in vec:
#    a = fn(a, e)
#  return a
#
#class PList(native):
#  def append(x):
#    native:
#      'self.PP = append(self.PP, a_x)'
#
#  def extend(x):
#    native:
#      'self.PP = append(self.PP, a_x.List()...)'
#
#  def count(x):
#    z = 0
#    for e in self:
#      if e == x:
#        z += 1
#    return z
#
#  def index(x):
#    i = 0
#    for e in self:
#      if e == x:
#        return i
#      i += 1
#    raise 'ValueError'
#
#  def remove(x):
#    del self[self.index(x)]
#
#  def insert(i, x):
#    # Tgetitemslice not supported yet ### self[i:i] = [x]
#    z = self[:i] + [x] + self[i:]
#    native:
#      'self.PP = v_z.List()'
#
#  def pop(i = -1):
#    x = self[i]
#    del self[i]
#    return x
#
#  def reverse():
#    n = len(self)
#    #say n, self
#    for i in range(int(n/2)):
#      #say i, n-i-1, self[i], self[n-i-1]
#      self[i], self[n-i-1] = self[n-i-1], self[i]
#
#  def sort(cmp=None, key=None, reverse=False):
#    native:
#      'self.PP = N_sorted(MkX(&self.PBase), a_cmp, a_key, a_reverse).List()'
#
#  def copy():
#    native: `
#      var zz []M
#      for _, e := range self.PP { zz = append(zz, e) }
#      return MkList(zz)
#    `

class JDict(native):
  def clear():
    native:
      'self.mu.Lock()'
      'self.ppp = make(JScope)'
      'self.mu.Unlock()'

  def copy():
    native:
      'z := make(JScope)'
      'self.mu.Lock()'
      'for k, v := range self.ppp { z[k] = v }'
      'self.mu.Unlock()'
      'p := MkDictJ(z)'
      'return inline.MkPJ(p)'

  def items():
    native:
      'z := make([]W, 0, len(self.ppp))'
      'self.mu.Lock()'
      'for k, v := range self.ppp {'
      '  key_1, key_2 := MkStrJ(k)'
      '  tup := MkTupleJ([]W{ W{key_1,key_2}, v})'
      '  u, v := inline.MkPJ(tup)'
      '  z = append(z, W{u,v})'
      '}'
      'self.mu.Unlock()'
      'p := MkListJ(z)'
      'return inline.MkPJ(p)'

  def iteritems():
    return .items()

#  def keys():
#    native:
#      'z := make([]M, 0, len(self.ppp))'
#      'self.mu.Lock()'
#      'for k, _ := range self.ppp { z = append(z, MkStr(k)) }'
#      'self.mu.Unlock()'
#      'return MkList(z)'
#  def iterkeys():
#    return .keys()
#  def iter():
#    return .keys()
#
#  def values():
#    native:
#      'z := make([]M, 0, len(self.ppp))'
#      'self.mu.Lock()'
#      'for _, v := range self.ppp { z = append(z, v) }'
#      'self.mu.Unlock()'
#      'return MkList(z)'
#  def itervalues():
#    return .values()
#
  def get(key, default = None):
    native:
      'k := StringJ(a_key_1, a_key_2)'
      'self.mu.Lock()'
      'z, ok := self.ppp[k]'
      'self.mu.Unlock()'
      'if ok { return z.U, z.V }'
      'return a_default_1, a_default_2'

#  def has_key(key):
#    return (key in self)
#
#  def setdefault(key, default=None):
#    if key in self:
#      return self[key]
#    else:
#      self[key] = default
#      return default
#
#  def update(x):
#    # TODO -- atomic update.
#    stuff = dict(x).items()
#    for k, v in stuff:
#      self[k] = v
#
#
#class PStr(native):
#  "PStr is a fake class to hold methods for the builtin type str."
#
#  def split(x = None, n = -1):
#    "Split self with delimiter x at most n times.  If x is None, split on white space."
#
#    if x is None:
#      native:
#        `
#        s := self.S
#        n := int(a_n.Int())
#        var v []string
#        for n < 0 || len(v) < n {
#          if len(s) == 0 { break }
#          i := i_strings.IndexAny(s, " \t\n\r")
#          if i >= 0 { if i>0 {v = append(v, s[:i])}; s = s[i+1:]
#          } else { if len(s) > 0 { v = append(v, s); break }}
#        }
#        return MkStrs(v)
#        `
#    if n < 0:
#      native:
#        'return MkStrs(i_strings.Split(self.S, a_x.String()))'
#    else:
#      native:
#        'return MkStrs(i_strings.SplitN(self.S, a_x.String(), 1 + int(a_n.Int())))'
#
#  def join(vec):
#    "Join the elements of vec adding self between the elements."
#    native:
#      'ss := make([]string, a_vec.Len())'
#      'for i, p := range a_vec.List() {'
#      '  ss[i] = p.String()'
#      '}'
#      'return MkStr(i_strings.Join(ss, self.S))'
#
#  def lower():
#    "Return self converted to lower case."
#    native:
#      'return MkStr(i_strings.ToLower(self.S))'
#
#  def title():
#    "Return self converted to title case."
#    native:
#      'return MkStr(i_strings.ToTitle(self.S))'
#
#  def upper():
#    "Return self converted to upper case."
#    native:
#      'return MkStr(i_strings.ToUpper(self.S))'
#
#  def endswith(x):
#    "Does self end with string x?"
#    native:
#      'return MkBool(i_strings.HasSuffix(self.S, a_x.String()))'
#
#  def startswith(x):
#    "Does self start with string x?"
#    native:
#      'return MkBool(i_strings.HasPrefix(self.S, a_x.String()))'
#
#  def strip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from front and end."
#    native:
#      'return MkStr(i_strings.Trim(self.S, a_x.Str()))'
#
#  def lstrip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from front."
#    native:
#      'return MkStr(i_strings.TrimLeft(self.S, a_x.Str()))'
#
#  def rstrip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from end."
#    native:
#      'return MkStr(i_strings.TrimRight(self.S, a_x.Str()))'
#
#  def replace(old, new, count = -1):
#    "Return self with nonoverlapping occurances of old replaced with new at most count times."
#    native:
#      'return MkStr(i_strings.Replace(self.S, a_old.String(), a_new.String(), int(a_count.Int())))'
#
#  def find(x):
#    "Return the index of the first occurance of x in self, or -1 if not found."
#    native:
#      'return Mkint(i_strings.Index(self.S, a_x.String()))'
#
#  def rfind(x):
#    "Return the index of the last occurance of x in self, or -1 if not found."
#    native:
#      'return Mkint(i_strings.LastIndex(self.S, a_x.String()))'
#
#  def index(x):
#    "Return the index of the first occurance of x in self, or throw an exception."
#    z = self.find(x)
#    if z < 0:
#      raise 'ValueError'
#    return z
#
#  def rindex(x):
#    "Return the index of the last occurance of x in self, or throw an exception."
#    z = self.rfind(x)
#    if z < 0:
#      raise 'ValueError'
#    return z
#
#  def isalpha():
#    "Are all runes in self unicode letters?"
#    if self:
#      for c in self:
#        if not unicode.IsLetter(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isdigit():
#    "Are all runes in self unicode digits?"
#    if self:
#      for c in self:
#        if not unicode.IsDigit(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isalnum():
#    "Are all runes in self unicode letters or digits?"
#    if self:
#      for c in self:
#        if (not unicode.IsDigit(ord(c))) and (not unicode.IsLetter(ord(c))):
#          return False
#      return True
#    else:
#      return False
#
#  def islower():
#    "Are all runes in self unicode lower case letters?"
#    if self:
#      for c in self:
#        if not unicode.IsLower(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isupper():
#    "Are all runes in self unicode upper case letters?"
#    if self:
#      for c in self:
#        if not unicode.IsUpper(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isspace():
#    "Are all runes in self unicode spaces?"
#    if self:
#      for c in self:
#        if not unicode.IsSpace(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#
#class PByt(native):
#  "PByt is a fake class to hold methods for the builtin type byt."
#
#  def split(x = None, n = -1):
#    "Split self with delimiter x at most n times.  If x is None, split on white space."
#
#    if x is None:
#      native:
#        `
#        s := self.YY
#        n := int(a_n.Int())
#        var v [][]byte
#        for n < 0 || len(v) < n {
#          if len(s) == 0 { break }
#          i := i_bytes.IndexAny(s, " \t\n\r")
#          if i >= 0 { if i>0 {v = append(v, s[:i])}; s = s[i+1:]
#          } else { if len(s) > 0 { v = append(v, s); break }}
#        }
#        return MkByts(v)
#        `
#    if n < 0:
#      native:
#        'return MkByts(i_bytes.Split(self.YY, a_x.Bytes()))'
#    else:
#      native:
#        'return MkByts(i_bytes.SplitN(self.YY, a_x.Bytes(), 1 + int(a_n.Int())))'
#
#  def join(vec):
#    "Join the elements of vec adding self between the elements."
#    native:
#      'ss := make([][]byte, a_vec.Len())'
#      'for i, p := range a_vec.List() {'
#      '  ss[i] = p.Bytes()'
#      '}'
#      'return MkByt(i_bytes.Join(ss, self.YY))'
#
#  def lower():
#    "Return self converted to lower case."
#    native:
#      'return MkByt(i_bytes.ToLower(self.YY))'
#
#  def title():
#    "Return self converted to title case."
#    native:
#      'return MkByt(i_bytes.ToTitle(self.YY))'
#
#  def upper():
#    "Return self converted to upper case."
#    native:
#      'return MkByt(i_bytes.ToUpper(self.YY))'
#
#  def endswith(x):
#    "Does self end with string x?"
#    native:
#      'return MkBool(i_bytes.HasSuffix(self.YY, a_x.Bytes()))'
#
#  def startswith(x):
#    "Does self start with string x?"
#    native:
#      'return MkBool(i_bytes.HasPrefix(self.YY, a_x.Bytes()))'
#
#  def strip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from front and end."
#    native:
#      'return MkByt(i_bytes.Trim(self.YY, a_x.Str()))'
#
#  def lstrip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from front."
#    native:
#      'return MkByt(i_bytes.TrimLeft(self.YY, a_x.Str()))'
#
#  def rstrip(x=' \t\n\r'):
#    "Return self with chars in x stripped away from end."
#    native:
#      'return MkByt(i_bytes.TrimRight(self.YY, a_x.Str()))'
#
#  def replace(old, new, count = -1):
#    "Return self with nonoverlapping occurances of old replaced with new at most count times."
#    native:
#      'return MkByt(i_bytes.Replace(self.YY, a_old.Bytes(), a_new.Bytes(), int(a_count.Int())))'
#
#  def find(x):
#    "Return the index of the first occurance of x in self, or -1 if not found."
#    native:
#      'return Mkint(i_bytes.Index(self.YY, a_x.Bytes()))'
#
#  def rfind(x):
#    "Return the index of the last occurance of x in self, or -1 if not found."
#    native:
#      'return Mkint(i_bytes.LastIndex(self.YY, a_x.Bytes()))'
#
#  def index(x):
#    "Return the index of the first occurance of x in self, or throw an exception."
#    z = self.find(x)
#    if z < 0:
#      raise 'ValueError'
#    return z
#
#  def rindex(x):
#    "Return the index of the last occurance of x in self, or throw an exception."
#    z = self.rfind(x)
#    if z < 0:
#      raise 'ValueError'
#    return z
#
#  def isalpha():
#    "Are all runes in self unicode letters?"
#    if self:
#      for c in self:
#        if not unicode.IsLetter(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isdigit():
#    "Are all runes in self unicode digits?"
#    if self:
#      for c in self:
#        if not unicode.IsDigit(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isalnum():
#    "Are all runes in self unicode letters or digits?"
#    if self:
#      for c in self:
#        if (not unicode.IsDigit(ord(c))) and (not unicode.IsLetter(ord(c))):
#          return False
#      return True
#    else:
#      return False
#
#  def islower():
#    "Are all runes in self unicode lower case letters?"
#    if self:
#      for c in self:
#        if not unicode.IsLower(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isupper():
#    "Are all runes in self unicode upper case letters?"
#    if self:
#      for c in self:
#        if not unicode.IsUpper(ord(c)):
#          return False
#      return True
#    else:
#      return False
#
#  def isspace():
#    "Are all runes in self unicode spaces?"
#    if self:
#      for c in self:
#        if not unicode.IsSpace(ord(c)):
#          return False
#      return True
#    else:
#      return False

def object():
  "object is the construtor for builtin object type."
  native: `
    jo := J_object{}
    jb := jo.GetJBase()
    return ForgeUV(jb)
  `

class J_object(native):
  "J_object is a fake class to hold methods for the builtin class object."
  # Defining __init__ at C_object lets you avoid it elsewhere.
  def __init__():
    pass
#  def __getattr__(field):
#    "Return the value of the named field on self."
#    native:
#      `return FetchFieldByNameForObjectJ(reflect.ValueOf(self.Self), a_field.String())`
#  def __setattr__(field, value):
#    "Set the value of the named field on self."
#    native:
#      `StoreFieldByNameForObjectJ(reflect.ValueOf(self.Self), a_field.String(), a_value)`

#class C_promise(native):
#  "C_promise is a fake class to hold methods for the builtin promise type."
#
#  def Wait():
#    "Wait on the promise to be completed, or throw the error if it is broken."
#    native:
#      'return self.Wait()'
#
#def rye_chan(size, revSize=-1):
#  "rye_chan is the construtor for a Go chan of Rye values."
#  native:
#    'return make_rye_chan(a_size.Int(), a_revSize.Int())'
#
#class C_rye_chan(native):
#  "C_rye_chan is a fake class to hold methods for the builtin rye_chan type."
#
#  def Throw(e):
#    "Send a special record Causing the receiver to throw the given exception when it is read."
#    native:
#      'self.Chan <- Either{Left: a_e, Valid: true, Right: MissingM}'
#
#  def Send(a):
#    "Send a rye value on the chan."
#    native:
#      'self.Chan <- Either{Left: nil, Valid: true, Right: a_a}'
#
#  def Recv():
#    "Receive a rye value from the chan, blocking until one is ready.  Or throw the exceptoin, if THrow() was called.  Or None, if closed."
#    native:
#      `
#        z := <-self.Chan
#        if z.Left != nil {
#          panic(z.Left)
#        } else if z.Valid {
#          return z.Right
#        }
#      `
#
#  def TryRecv():
#    "Receive a rye value from the chan, or return None if no value is ready or the channel is closed..  Or throw the exceptoin, if THrow() was called."
#    native:
#      `
#        var z Either
#        select {
#          case z = <-self.Chan:
#          default:
#        }
#        if z.Left != nil {
#          panic(z.Left)
#        } else if z.Right != MissingM {
#          return z.Right
#        } else {
#          return None  // if channel was closed.
#        }
#      `
#
#  def Close():
#    "Close the channel, signaling that nothing more will be Sent on it."
#    native:
#      'close(self.Chan)'
#
#def open(filename, mode='r'):
#  "Open the named file, with given mode, as in Python.  Returns an instance of PYE_FileDesc."
#  if mode == 'r':
#    return PYE_FileDesc(os.Open(filename), False)
#  elif mode == 'w':
#    return PYE_FileDesc(os.Create(filename), True)
#  elif mode == 'a':
#    return PYE_FileDesc(os.OpenFile(filename, int(os.O_WRONLY)|int(os.O_APPEND), 0644), True)
#  else:
#    raise 'open: Unknown mode: %q' % mode
#
#class PYE_FileDesc:
#  "The internal type returned from the builtin open() function.  Go's io.Writer protocol is also supported."
#  def __init__(fd, writing):
#    "Internal."
#    .writing = writing
#    if writing:
#      .f = fd
#      .b = bufio.NewWriter(fd)
#    else:
#      .f = fd
#      .b = bufio.NewReader(fd)
#
#  def read():
#    "Read the rest of the file as a string."
#    return str(ioutil.ReadAll(.b))
#
#  def write(x):
#    "Write tye bytes in x, which can be str or byt."
#    .b.Write(str(x))
#
#  def flush():
#    "Flush buffered written bytes to the file."
#    .b.Flush()
#  def Flush():
#    "Flush buffered written bytes to the file."
#    .b.Flush()
#
#  def close():
#    "Close the file."
#    if .writing:
#      .b.Flush()
#    .f.Close()
#  def Close():
#    "Close the file."
#    if .writing:
#      .b.Flush()
#    .f.Close()
#
#native: `
#  // io.Writer protocol for writing:
#  func (self *C_PYE_FileDesc) Write(p []byte) (n int, err error) {
#    return self.M_b.Contents().(io.Writer).Write(p)
#  }
#  func (self *C_PYE_FileDesc) Flush() error {
#    self.M_0_Flush()
#    return nil
#  }
#`

def _force_generation_of_call_4_(f, a, b, c, d):
  return f(a, b, c, d)

pass
