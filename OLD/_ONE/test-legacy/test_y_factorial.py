# Borrowed from lua-5.0/test/factorial.lua (pasted below)

# traditional fixed-point operator from functional programming
def Y(g):
  def a(f):
    return f(f)
  return a(lambda f: g(lambda x: f(f)(x)))

# factorial without recursion
def FACTORIAL(f):
  def inner(n ::int) ::int :
     return 1 if n==0 else n*f(n-1)
  return inner

factorial = Y(FACTORIAL)   # factorial is the fixed point of FACTORIAL

assert factorial(4) == 24
assert factorial(5) == 120

for i in range(10):
  say i, factorial(i)

"""
  -- function closures are powerful

  -- traditional fixed-point operator from functional programming
  Y = function (g)
        local a = function (f) return f(f) end
        return a(function (f)
                   return g(function (x)
                               local c=f(f)
                               return c(x)
                             end)
                 end)
  end


  -- factorial without recursion
  F = function (f)
        return function (n)
                 if n == 0 then return 1
                 else return n*f(n-1) end
               end
      end

  factorial = Y(F)   -- factorial is the fixed point of F

  -- now test it
  function test(x)
      io.write(x,"! = ",factorial(x),"\n")
  end

  for n=0,16 do
      test(n)
  end
"""
