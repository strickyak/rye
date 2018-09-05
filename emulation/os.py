from go import os, os/exec as Exec
from go import path/filepath

def getenv(s : str) -> str:
  "Get the environment variable, or return empty string if missing."
  return os.Getenv(s)

def makedirs(s : str, mode=0777) -> None:
  "Make the dir with the given mode, making dirs as needed along the way."
  os.MkdirAll(s, mode)

def getcwd() -> str:
  "Return the current directory."
  return os.Getwd()

def remove(filename):
  return os.Remove(filename)

def link(old, new):
  return os.Link(old, new)

def system(s : str) -> int:
  cmd = Exec.Command("/bin/sh", "-c", s)
  cmd.Run()
  return 0

class PathCls:
  def __init__():
    pass
  def basename(x : str) -> str:
    return filepath.Base(x)
  def dirname(x : str) -> str:
    return filepath.Dir(x)
  def join(*args) -> str:
    return filepath.Join(*args)

path = PathCls()

class stat:
  "For the given filename, return an object with .st_size & .st_mtime fields."
  def __init__(filename : str):
    fi = os.Stat(filename)
    .st_size = fi.Size()
    .st_mtime = fi.ModTime().Unix()
