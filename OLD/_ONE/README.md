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

Types
   - None, bool, int, float.  (All ints are int64; there is no long.  Use Go's math/big if you need big ints.)
   - str, byt (byt corresponds to []byte in Go.  It converts to/from str, but also to/from a list of int.).
   - tuple, list, dict.  Iterators on them.
   - object (ultimate superclass of all user-defined classes).
   - Various callable things: named functions, lambda exprs, class constructors, uncalled instance methods.
   - Reflective references to Go objects (passed in and out of Go functions).

Goroutines
   - go fn(x, y, z, ...) is an expression, returning a promise, running in background.  You can Wait() for the result.
   - Generator functions (those using yield) run in their own goroutine, when iterated in a "for" loop.
   - Go libaries (like net/http) also create goroutines, which can call back into Rye functions.

Rye Modules
   - Each source .py file defines 1 module (not each directory, as in go).

Imports
   - Import statements:
    - from . import stuff      # where stuff.py is another Rye module
    - from go import net/http  # Go standard library
    - from go import github.com/russross/blackfriday   # Other Go packages
    - import re                # A small amount of python library emulation is available.

Limitations
   - All ints are int64.  No longs.
   - No monkey-patching.
   - No module reflection (yet).
   - No set type (use dict).
   - Many builtins and methods on builtin types are missing.
   - No python libraries.  Use go libraries.
   - No nested classes.  No class methods.
   - Dictionary keys are converted to strings (as in javascript).
   - No code blocks after ":" on same line with "if", "for", etc., statements.
   - Continued lines (due to unclosed parens, etc.) should indent but not outdent.
   - Lack of meaningful error messages due to unclosed parens, etc.  (Obscured by "continued line" logic.)
   - More limitations.
   - Also bugs.

Compile
   - python .../rye/rye.py build src.py (creates a binary src/src).
   - python .../rye/rye.py run src.py arg1 arg2...

Interpreter
   - python .../rye/rye.py interp src.py
   - Commands /b ; /g ; /l ; (builtins; globals; locals)
   - work in progress.

Demos
   - test*.py in this "rye" directory
   - See also "strickyak/aphid" github for lots of Rye code.
