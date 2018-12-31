import os

assert os.getenv("HOME") != ""

os.makedirs("./tmp1/tmp2/tmp3")
fd = open("./tmp1/tmp2/tmp3/foo4", "w")
print >> fd, "hello foo4"
fd.close()
fd = open("./tmp1/tmp2/tmp3/foo4", "r")
s = fd.read()
fd.close()
assert s == "hello foo4\n"
os.remove("./tmp1/tmp2/tmp3/foo4")
assert 0 == os.system("rmdir ./tmp1/tmp2/tmp3")
assert 0 == os.system("rmdir ./tmp1/tmp2")
assert 0 == os.system("rmdir ./tmp1")

assert os.getcwd().endswith("rye/emulation/tests")


