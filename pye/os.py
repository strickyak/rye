from go import os
from go import path/filepath

def getenv(s):
  return os.Getenv(s)

def makedirs(s, mode=0777):
  os.MkdirAll(s, mode)

def getcwd():
  return os.Getwd()

class PathCls:
  def __init__():
    pass
  def basename(x):
    return filepath.Base(x)
  def dirname(x):
    return filepath.Dir(x)
  def join(*args):
    return filepath.Join(*args)

path = PathCls()

class stat:
  def __init__(filename):
    fi = os.Stat(filename)
    .st_size = fi.Size()
    .st_mtime = fi.ModTime().Unix()
