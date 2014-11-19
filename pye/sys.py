from go import io, io/ioutil
from go import os

stdout = FileDesc(os.Stdout)

def os_open(filename, mode):
  if mode == 'r':
    return FileDesc(os.Open(filename))
  elif mode == 'w':
    return FileDesc(os.Create(filename))
  else:
    raise 'Unsupported Mode', mode

class FileDesc:
  def __init__(fd):
    .fd = fd
  def read():
    return str(ioutil.ReadAll(.fd))
  def write(s):
    return io.WriteString(.fd, s)
  def flush():
    .fd.Flush()
  def close():
    .fd.Close()

native:
  'func (self *C_FileDesc) Write(p []byte) (n int, err error) {'
  '  return self.M_fd.(*PGo).V.Interface().(io.Writer).Write(p)'
  '}'
