rye
===

Rye translates a rigid subset of Python into Go.

Why Go?
   - the Go runtime rules
   - Go libraries rule
   - Go concurrency rules

Why Python?
   - I love Single Inheritance.
   - I love exceptions.
   - I love duck typing.
   - (They have their problems, but for some projects, they're the problems I choose.)

Limitations
   - All ints are int64.  No longs.
   - No monkey-patching.
   - No module reflection (yet).
   - No named arguments.  No * or ** or default arguments (yet).
   - Bitwise operators are missing.
   - Many builtins and methods on builtin types are missing.
   - No python libraries.  Use go libraries.
   - No nested functions or classes.  No nested scopes.
   - Dictionary keys are converted to strings (as in javascript).
   - import statements are a little different.
   - Source files, as in Python (not directories, as in Go), are modules.
   - More limitations.
   - Also bugs.
