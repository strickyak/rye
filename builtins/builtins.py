from go import bytes, strings, unicode
from go import bufio, io, io/ioutil, log, os, time

# MACRO go_type(t) -- creates a reflective Value of the go type t.

# MACRO go_new(t) -- creates a new zeroed instance of the go type t and returns reflective pointer to it.

# MACRO go_cast(t, x) -- Casts the value x to the go type t.

# Newer Python requires this, but for rye, it is the identity function.
def Exception(x):
  return x

def rye_opts():
  native:
    'return MkStr(CompileOptions)'

def go_deref(x):
  native:
    'return GoDeref(a_x)'

#def go_wrap(x):
#  native:
#    'return MkValue(reflect.ValueOf(JContents(a_x)))'
#
#def go_typeof(x):
#  native:
#    'return MkGo(reflect.ValueOf(JContents(a_x)).Type())'
#
#def go_kindof(x):
#  native:
#    'return MkStr(reflect.ValueOf(JContents(a_x)).Type().Kind().String())'
#
#def go_valueof(x):
#  native:
#    'return MkValue(reflect.ValueOf(reflect.ValueOf(JContents(a_x))))'

def rye_what(x):
  native:
    'return N_rye_what(a_x)'

def callable(x):
  native:
    'return MkBool(JCallable(a_x))'

def setattrs(obj, **kw):
  for k, v in kw.items():
    setattr(obj, k, v)
  return obj

#def globals():
  # TODO -- this is not going to work.
  #z = {}
  #native:
  #  'for k, ptr := range ModuleObj().Map() {'
  #  '  v_z.SetItem(MkStr(k), *ptr)'
  #  '}'
  #return z

def id(x):
  native: `
    if a_x.X == nil {
      if len(a_x.S) == 0 {
        panic("Rye does not support id() on int")
      } else {
        panic("Rye does not support id() on str")
      }
    }
    return MkInt(int64(reflect.ValueOf(a_x.X.B()).Pointer()))
  `

def hash(x):
  native:
    'return MkInt(JHash(a_x))'

def cmp(x, y):
  native:
    'return Mkint(JCompare(a_x, a_y))'

def getattr(x, name, *dflt):
  n = len(dflt)
  if n:
    try:
      native:
        'return a_x.FetchField(JString(a_name))'
    except:
      return dflt[0]
  else:
    native:
      'return a_x.FetchField(JString(a_name))'

def setattr(x, name, val):
  native:
    'a_x.StoreField(JString(a_name), a_val)'


def isinstance(obj, cls):
  native:
    'return MkBool(IsSubclass(JPType(a_obj), a_cls))'

def issubclass(subcls, cls):
  native:
    'return MkBool(IsSubclass(a_subcls, a_cls))'

def ord(x):
  native:
    'return Mkint(int(JString(a_x)[0]))'

def chr(x):
  native:
    'return MkStr(string([]byte{byte(JInt(a_x))}))'

def sum(vec, init=0):
  z = init
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
    'return Mkint(JLen(a_x))'

def repr(x):
  native:
    'return MkStr(JRepr(a_x))'

def str(x):
  native:
    'return MkStr(JString(a_x))'

def int(x):
  native:
    'return MkInt(JForceInt(a_x))'

def float(x):
  native:
    'return MkFloat(JForceFloat(a_x))'

def range(x):
  native:
    'return N_range(a_x)'

def xrange(x):
  i = 0
  while i < x:
    yield i
    i += 1

def sorted(x, cmp=None, key=None, reverse=False):
  native:
    'return N_sorted(a_x, a_cmp, a_key, a_reverse)'

def list(x):
  native:
    'return MkList(JList(a_x))'

def set(a=None):
  native:
    'return N_set(a_a)'

def dict(*vec, **kv):
  native:
    'return N_dict(a_vec, a_kv)'

def tuple(x):
  native:
    'return MkTuple(JList(a_x))'

def bool(x):
  native:
    'return MkBool(JBool(a_x))'

def type(x):
  native:
    'return JPType(a_x)'

def byt(x):
  native:
    'return N_byt(a_x)'

def mkbyt(n):
  native:
    'return N_mkbyt(a_n)'

def rye_stack():
  native:
    'return MkStr(RyeStack())'

def rye_pickle(x):
  native:
    'return MkByt(Pickle(a_x))'

def rye_unpickle(x):
  native:
    'return UnPickle(JBytes(a_x))'

def max(*args):
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

def min(*args):
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

def zip(*args):
  n = min([len(a) for a in args])
  return [tuple([a[i] for a in args]) for i in range(n)]

def rye_zip_padding_with_None(*args):
  m = max([len(a) for a in args])
  return [tuple([a[i] if i < len(a) else None for a in args]) for i in range(m)]

def map(fn, *lists):
  switch len(lists):
    case 0:
      raise 'map called with no lists'
    case 1:
      return [fn(x) for x in lists[0]]
    default:
      # It pads None for short list later items.
      return [fn(*tup) for tup in rye_zip_padding_with_None(*lists)]

def reduce(fn, vec, init=None):
  vec = list(vec)
  if init is None:
    a = vec.pop(0)
  else:
    a = init
  for e in vec:
    a = fn(a, e)
  return a

class PList(native):
  def append(x):
    native:
      'self.PP = append(self.PP, a_x)'

  def extend(x):
    native:
      'self.PP = append(self.PP, JList(a_x)...)'

  def count(x):
    z = 0
    for e in self:
      if e == x:
        z += 1
    return z

  def index(x):
    i = 0
    for e in self:
      if e == x:
        return i
      i += 1
    raise 'ValueError'

  def remove(x):
    del self[self.index(x)]

  def insert(i, x):
    # Tgetitemslice not supported yet ### self[i:i] = [x]
    z = self[:i] + [x] + self[i:]
    native:
      'self.PP = JList(v_z)'

  def pop(i = -1):
    x = self[i]
    del self[i]
    return x

  def reverse():
    n = len(self)
    #say n, self
    for i in range(int(n/2)):
      #say i, n-i-1, self[i], self[n-i-1]
      self[i], self[n-i-1] = self[n-i-1], self[i]

  def sort(cmp=None, key=None, reverse=False):
    native:
      'self.PP = JList(N_sorted(/*inline.*/MkObj(&self.PBase), a_cmp, a_key, a_reverse))'

  def copy():
    native: `
      var zz []M
      for _, e := range self.PP { zz = append(zz, e) }
      return MkList(zz)
    `

class PDict(native):
  def clear():
    native:
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'self.ppp = make(map[string]M)'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'

  def copy():
    native:
      'z := make(map[string]M)'
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'for k, v := range self.ppp { z[k] = v }'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'
      'return MkDict(z)'

  def items():
    native:
      'z := make([]M, 0, len(self.ppp))'
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'for k, v := range self.ppp { z = append(z, MkTuple([]M{MkStr(k), v})) }'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'
      'return MkList(z)'
  def iteritems():
    return .items()

  def keys():
    native:
      'z := make([]M, 0, len(self.ppp))'
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'for k, _ := range self.ppp { z = append(z, MkStr(k)) }'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'
      'return MkList(z)'
  def iterkeys():
    return .keys()
  def iter():
    return .keys()

  def values():
    native:
      'z := make([]M, 0, len(self.ppp))'
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'for _, v := range self.ppp { z = append(z, v) }'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'
      'return MkList(z)'
  def itervalues():
    return .values()

  def get(key, default = None):
    native:
      'k := JString(a_key)'
      "if 'm' {"
      'self.mu.Lock()'
      '}'
      'z, ok := self.ppp[k]'
      "if 'm' {"
      'self.mu.Unlock()'
      '}'
      'if ok { return z }'
      'return a_default'

  def has_key(key):
    return (key in self)

  def setdefault(key, default=None):
    if key in self:
      return self[key]
    else:
      self[key] = default
      return default

  def update(x):
    # TODO -- atomic update.
    stuff = dict(x).items()
    for k, v in stuff:
      self[k] = v


class PStr(native):
  "PStr is a fake class to hold methods for the builtin type str."

  def split(x = None, n = -1):
    "Split self with delimiter x at most n times.  If x is None, split on white space."

    if x is None:
      native:
        `
        s := self.S
        n := int(JInt(a_n))
        var v []string
        for n < 0 || len(v) < n {
          if len(s) == 0 { break }
          i := i_strings.IndexAny(s, " \t\n\r")
          if i >= 0 { if i>0 {v = append(v, s[:i])}; s = s[i+1:]
          } else { if len(s) > 0 { v = append(v, s); break }}
        }
        return MkStrs(v)
        `
    if n < 0:
      native:
        'return MkStrs(i_strings.Split(self.S, JString(a_x)))'
    else:
      native:
        'return MkStrs(i_strings.SplitN(self.S, JString(a_x), 1 + int(JInt(a_n))))'

  def join(vec):
    "Join the elements of vec adding self between the elements."
    native:
      'ss := make([]string, JLen(a_vec))'
      'for i, p := range JList(a_vec) {'
      '  ss[i] = JString(p)'
      '}'
      'return MkStr(i_strings.Join(ss, self.S))'

  def lower():
    "Return self converted to lower case."
    native:
      'return MkStr(i_strings.ToLower(self.S))'

  def title():
    "Return self converted to title case."
    native:
      'return MkStr(i_strings.ToTitle(self.S))'

  def upper():
    "Return self converted to upper case."
    native:
      'return MkStr(i_strings.ToUpper(self.S))'

  def endswith(x):
    "Does self end with string x?"
    native:
      'return MkBool(i_strings.HasSuffix(self.S, JString(a_x)))'

  def startswith(x):
    "Does self start with string x?"
    native:
      'return MkBool(i_strings.HasPrefix(self.S, JString(a_x)))'

  def strip(x=' \t\n\r'):
    "Return self with chars in x stripped away from front and end."
    native:
      'return MkStr(i_strings.Trim(self.S, JStr(a_x)))'

  def lstrip(x=' \t\n\r'):
    "Return self with chars in x stripped away from front."
    native:
      'return MkStr(i_strings.TrimLeft(self.S, JStr(a_x)))'

  def rstrip(x=' \t\n\r'):
    "Return self with chars in x stripped away from end."
    native:
      'return MkStr(i_strings.TrimRight(self.S, JStr(a_x)))'

  def replace(old, new, count = -1):
    "Return self with nonoverlapping occurances of old replaced with new at most count times."
    native:
      'return MkStr(i_strings.Replace(self.S, JString(a_old), JString(a_new), int(JInt(a_count))))'

  def find(x):
    "Return the index of the first occurance of x in self, or -1 if not found."
    native:
      'return Mkint(i_strings.Index(self.S, JString(a_x)))'

  def rfind(x):
    "Return the index of the last occurance of x in self, or -1 if not found."
    native:
      'return Mkint(i_strings.LastIndex(self.S, JString(a_x)))'

  def index(x):
    "Return the index of the first occurance of x in self, or throw an exception."
    z = self.find(x)
    if z < 0:
      raise 'ValueError'
    return z

  def rindex(x):
    "Return the index of the last occurance of x in self, or throw an exception."
    z = self.rfind(x)
    if z < 0:
      raise 'ValueError'
    return z

  def isalpha():
    "Are all runes in self unicode letters?"
    if self:
      for c in self:
        if not unicode.IsLetter(ord(c)):
          return False
      return True
    else:
      return False

  def isdigit():
    "Are all runes in self unicode digits?"
    if self:
      for c in self:
        if not unicode.IsDigit(ord(c)):
          return False
      return True
    else:
      return False

  def isalnum():
    "Are all runes in self unicode letters or digits?"
    if self:
      for c in self:
        if (not unicode.IsDigit(ord(c))) and (not unicode.IsLetter(ord(c))):
          return False
      return True
    else:
      return False

  def islower():
    "Are all runes in self unicode lower case letters?"
    if self:
      for c in self:
        if not unicode.IsLower(ord(c)):
          return False
      return True
    else:
      return False

  def isupper():
    "Are all runes in self unicode upper case letters?"
    if self:
      for c in self:
        if not unicode.IsUpper(ord(c)):
          return False
      return True
    else:
      return False

  def isspace():
    "Are all runes in self unicode spaces?"
    if self:
      for c in self:
        if not unicode.IsSpace(ord(c)):
          return False
      return True
    else:
      return False


class PByt(native):
  "PByt is a fake class to hold methods for the builtin type byt."

  def split(x = None, n = -1):
    "Split self with delimiter x at most n times.  If x is None, split on white space."

    if x is None:
      native:
        `
        s := self.YY
        n := int(JInt(a_n))
        var v [][]byte
        for n < 0 || len(v) < n {
          if len(s) == 0 { break }
          i := i_bytes.IndexAny(s, " \t\n\r")
          if i >= 0 { if i>0 {v = append(v, s[:i])}; s = s[i+1:]
          } else { if len(s) > 0 { v = append(v, s); break }}
        }
        return MkByts(v)
        `
    if n < 0:
      native:
        'return MkByts(i_bytes.Split(self.YY, JBytes(a_x)))'
    else:
      native:
        'return MkByts(i_bytes.SplitN(self.YY, JBytes(a_x), 1 + int(JInt(a_n))))'

  def join(vec):
    "Join the elements of vec adding self between the elements."
    native:
      'ss := make([][]byte, JLen(a_vec))'
      'for i, p := range JList(a_vec) {'
      '  ss[i] = JBytes(p)'
      '}'
      'return MkByt(i_bytes.Join(ss, self.YY))'

  def lower():
    "Return self converted to lower case."
    native:
      'return MkByt(i_bytes.ToLower(self.YY))'

  def title():
    "Return self converted to title case."
    native:
      'return MkByt(i_bytes.ToTitle(self.YY))'

  def upper():
    "Return self converted to upper case."
    native:
      'return MkByt(i_bytes.ToUpper(self.YY))'

  def endswith(x):
    "Does self end with string x?"
    native:
      'return MkBool(i_bytes.HasSuffix(self.YY, JBytes(a_x)))'

  def startswith(x):
    "Does self start with string x?"
    native:
      'return MkBool(i_bytes.HasPrefix(self.YY, JBytes(a_x)))'

  def strip(x=' \t\n\r'):
    "Return self with chars in x stripped away from front and end."
    native:
      'return MkByt(i_bytes.Trim(self.YY, JStr(a_x)))'

  def lstrip(x=' \t\n\r'):
    "Return self with chars in x stripped away from front."
    native:
      'return MkByt(i_bytes.TrimLeft(self.YY, JStr(a_x)))'

  def rstrip(x=' \t\n\r'):
    "Return self with chars in x stripped away from end."
    native:
      'return MkByt(i_bytes.TrimRight(self.YY, JStr(a_x)))'

  def replace(old, new, count = -1):
    "Return self with nonoverlapping occurances of old replaced with new at most count times."
    native:
      'return MkByt(i_bytes.Replace(self.YY, JBytes(a_old), JBytes(a_new), int(JInt(a_count))))'

  def find(x):
    "Return the index of the first occurance of x in self, or -1 if not found."
    native:
      'return Mkint(i_bytes.Index(self.YY, JBytes(a_x)))'

  def rfind(x):
    "Return the index of the last occurance of x in self, or -1 if not found."
    native:
      'return Mkint(i_bytes.LastIndex(self.YY, JBytes(a_x)))'

  def index(x):
    "Return the index of the first occurance of x in self, or throw an exception."
    z = self.find(x)
    if z < 0:
      raise 'ValueError'
    return z

  def rindex(x):
    "Return the index of the last occurance of x in self, or throw an exception."
    z = self.rfind(x)
    if z < 0:
      raise 'ValueError'
    return z

  def isalpha():
    "Are all runes in self unicode letters?"
    if self:
      for c in self:
        if not unicode.IsLetter(ord(c)):
          return False
      return True
    else:
      return False

  def isdigit():
    "Are all runes in self unicode digits?"
    if self:
      for c in self:
        if not unicode.IsDigit(ord(c)):
          return False
      return True
    else:
      return False

  def isalnum():
    "Are all runes in self unicode letters or digits?"
    if self:
      for c in self:
        if (not unicode.IsDigit(ord(c))) and (not unicode.IsLetter(ord(c))):
          return False
      return True
    else:
      return False

  def islower():
    "Are all runes in self unicode lower case letters?"
    if self:
      for c in self:
        if not unicode.IsLower(ord(c)):
          return False
      return True
    else:
      return False

  def isupper():
    "Are all runes in self unicode upper case letters?"
    if self:
      for c in self:
        if not unicode.IsUpper(ord(c)):
          return False
      return True
    else:
      return False

  def isspace():
    "Are all runes in self unicode spaces?"
    if self:
      for c in self:
        if not unicode.IsSpace(ord(c)):
          return False
      return True
    else:
      return False

def object():
  "object is the construtor for builtin object type."
  native: `return MForge(&C_object{})`

class C_object(native):
  "C_object is a fake class to hold methods for the builtin class object."
  # Defining __init__ at C_object lets you avoid it elsewhere.
  def __init__():
    pass
  def __getattr__(field):
    "Return the value of the named field on self."
    native:
      `return FetchFieldByNameForObject(reflect.ValueOf(self.Self), JString(a_field))`
  def __setattr__(field, value):
    "Set the value of the named field on self."
    native:
      `StoreFieldByNameForObject(reflect.ValueOf(self.Self), JString(a_field), a_value)`

def rye_chan(size):
  "rye_chan creates a Started rye channel of the given buffer size."
  native:
    'return make_rye_chan(int(JInt(a_size)))'

class C_channel(native):
  "C_channel is a fake class to hold methods for the builtin rye_chan type."

  def Warm():
    native:
      `
        Log.Printf("... C_channel::Warm waiting: %v ...", self)
        b := <-self.Back
        if b != FeedbackStart {
          Log.Panicf("C_channel::Warm got feedback %d, expected FeedbackStart", b)
        }
        Log.Printf("... C_channel::Warm got Start: %v", self)
      `

  def Start(size):
    native:
      'self.Start(int(JInt(a_size)))'

  def Raise(e):
    "Send a special record Causing the receiver to throw the given exception when it is read."
    native:
      'self.Raise(a_e)'

  def Send(a):
    "Send a rye value on the chan."
    native:
      'self.Send(a_a)'

  def Recv(timeout=None):
    "Receive a rye value from the chan, blocking until one is ready.  ok if false if the channel was closed."
    native:
      `
        if (a_timeout == None) {
          z, ok := self.Recv()
          return MkTuple([]M{ z, MkBool(ok) })
	} else {
          z, ok := self.RecvWithTimeout(i_time.Duration((JFloat(a_timeout)*1000000000.0)) * i_time.Nanosecond)
          return MkTuple([]M{ z, MkBool(ok) })
	}
      `

  def Wait():
    """Receive a rye value from the chan, blocking until one is ready.  Throw "EOF" if the channel was closed."""
    native:
      `
        z, ok := self.Recv()
	if !ok {
		panic("EOF")
	}
        return z
      `

  def Close():
    "Close the channel, signaling that nothing more will be Sent on it."
    native:
      'self.Close()'

def open(filename, mode='r'):
  "Open the named file, with given mode, as in Python.  Returns an instance of RyeFileDesc."
  if mode == 'r':
    return RyeFileDesc(os.Open(filename), False)
  elif mode == 'w':
    return RyeFileDesc(os.Create(filename), True)
  elif mode == 'a':
    return RyeFileDesc(os.OpenFile(filename, int(os.O_WRONLY)|int(os.O_APPEND), 0644), True)
  else:
    raise 'open: Unknown mode: %q' % mode

class RyeFileDesc:
  "The internal type returned from the builtin open() function.  Go's io.Writer protocol is also supported."
  def __init__(fd, writing):
    "Internal."
    .writing = writing
    if writing:
      .f = fd
      .b = bufio.NewWriter(fd)
    else:
      .f = fd
      .b = bufio.NewReader(fd)

  def read():
    "Read the rest of the file as a string."
    return str(ioutil.ReadAll(.b))

  def write(x):
    "Write tye bytes in x, which can be str or byt."
    .b.Write(str(x))

  def flush():
    "Flush buffered written bytes to the file."
    .b.Flush()
  def Flush():
    "Flush buffered written bytes to the file."
    .b.Flush()

  def close():
    "Close the file."
    if .writing:
      .b.Flush()
    .f.Close()
  def Close():
    "Close the file."
    if .writing:
      .b.Flush()
    .f.Close()

native: `
  // io.Writer protocol for writing:
  func (self *C_RyeFileDesc) Write(p []byte) (n int, err error) {
    return JContents(self.M_b).(io.Writer).Write(p)
  }
  func (self *C_RyeFileDesc) Flush() error {
    self.M_0_Flush()
    return nil
  }
`

def _rye__force_generation_of_call_4_(f, a, b, c, d):
  return f(a, b, c, d)

pass
