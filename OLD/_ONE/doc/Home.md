## Documentation for the Rye Language.

Rye is a useful subset of Python that compiles to Go.

Rye is compatible with most Go libraries, and gives you
access to Go's modern multi-threaded runtime.

From Python, we get its famous ease of use and readability.
The usual bool, int, float, str, list, set, & dict types are
supported.  Classes are supported with single inheritance.
Exceptions are used in Rye.  There are a few extensions
that help you interface with Go.

Because we use Go's runtime instead of Python's runtime,
we also use Go libraries instead of Python libraries.
Rye provides its version of many of the builtin functions of Python,
and many of the builtin methods on the basic data types.
But for more complicated things, use Go libraries.

The source to Rye is available at
[github.com/strickyak/rye](https://github.com/strickyak/rye).
