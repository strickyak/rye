"""
The MD5 cryptographic hash algorithm.

Class md5 computes MD5 hashes.
In this minimal implementation, the string to be
hashed must be passed to the construtor.
The object cannot be updated.
The only method on the object is hexdigest().

Globals:
  blocksize = 1     # It works on individual bytes.
  digest_size = 16  # The digest is 16 bytes.
  new = md5         # new is an alias for md5.
"""
from go import "crypto/md5" as go_md5
from go import fmt

blocksize = 1
digest_size = 16

class md5:
  "The MD5 cryptographic hash."
  def __init__(s):
    "Create an md5 object with the hash of the given string."
    .b = str(go_md5.Sum(s))

  def hexdigest():
    "Return the digest as a hexadecimal string."
    return fmt.Sprintf('%x', .b)

  def digest():
    "Return the digest as a binary string."
    return .b

new = md5  # new is an alias for md5.
