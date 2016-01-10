from go import crypto/md5
from go import fmt

class new:
  def __init__(s):
    .b = md5.Sum(s)

  def hexdigest():
    return fmt.Sprintf('%x', .b)
