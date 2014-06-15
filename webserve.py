go import fmt
go import html
go import net/http

class BarClass:
  def __init__(self):
    pass

  def ServeHTTP(self, w, r):
    fmt.Fprintf(w, "Bar bar bar, %q", html.EscapeString(r.URL.Path))

  native:
    'func (o *C_BarClass) ServeHTTP(w i_http.ResponseWriter, r *i_http.Request) {'
    '  o.M_2_ServeHTTP(MkGo(w), MkGo(r))'
    '}'
  pass
http.Handle('/bar', BarClass())

def FooFunc(w, r):
  try:
    fmt.Fprintf(w, "Hello Foo, %q", html.EscapeString(r.URL.Path))
  except as ex:
    print "FooFunc Exception:", ex
http.HandleFunc('/foo', FooFunc)

print 'gotype=', gotype(http.Dir)
print 'gocast=', gocast(http.Dir, "/etc")
d = gocast(http.Dir, "/etc")
fs = http.FileServer(d)
http.Handle("/", fs)

http.ListenAndServe( ':8080' , None )
