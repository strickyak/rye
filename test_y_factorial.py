# Borrowed from lua-5.0/test/factorial.lua

# traditional fixed-point operator from functional programming
def Y(g):
  def a(f):
    return f(f)
  return a(lambda f: g(lambda x: f(f)(x)))

# factorial without recursion
def F(f):
  return lambda(n): 1 if n==0 else n*f(n-1)

factorial = Y(F)   # factorial is the fixed point of F

assert factorial(4) == 24

# -- now test it
# function test(x)
# 	io.write(x,"! = ",factorial(x),"\n")
# end
# 
# for n=0,16 do
# 	test(n)
# end
