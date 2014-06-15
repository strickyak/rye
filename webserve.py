go import fmt
go import html
go import net/http

class HelloClass:
  def __init__(self):
    pass

  def ServeHTTP(self, w, r):
    fmt.Fprintf(w, "Bar bar bar, %q", html.EscapeString(r.URL.Path))

  native:
    'func (o *C_HelloClass) ServeHTTP(w i_http.ResponseWriter, r *i_http.Request) {'
    '  o.M_2_ServeHTTP(MkGo(w), MkGo(r))'
    '}'
  pass

hc = gocast(http.Handler, HelloClass())
http.Handle('/bar', hc)

def HelloFunc(w, r):
  try:
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
  except as ex:
    print "HelloFunc Exception:", ex
http.HandleFunc('/foo', HelloFunc)

print 'gotype=', gotype(http.Dir)
print 'gocast=', gocast(http.Dir, "/etc")
d = gocast(http.Dir, "/etc")
fs = http.FileServer(d)
http.Handle("/", fs)

http.ListenAndServe( ':8080' , None )
