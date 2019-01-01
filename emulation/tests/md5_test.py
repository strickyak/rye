"""Test the md5 implementation from rye_pye.

This should work with either python or rye.
"""
import md5  # rye_pragma from "github.com/strickyak/rye/emulation"

x = md5.md5("hello rye")
assert x.digest() == '\xa0\x9e\x97$jBHD\xd1(\x85w\x86\xe9\xe3n'
assert x.hexdigest() == 'a09e97246a424844d128857786e9e36e'

y = md5.new("")
assert y.digest() == '\xd4\x1d\x8c\xd9\x8f\x00\xb2\x04\xe9\x80\t\x98\xec\xf8B~'
assert y.hexdigest() == 'd41d8cd98f00b204e9800998ecf8427e'

assert md5.blocksize == 1
assert md5.digest_size == 16
