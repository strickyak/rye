go import fmt
go import html
go import net/http

#class HelloClass:
#  def __init__(self):
#    pass
#
#  def ServeHTTP(self, w, r):
#    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
#http.Handle('/foo', HelloClass())

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
sp = http.StripPrefix("/fs", fs)
# http.Handle("/fs", sp) # Did not add /fs to links.
http.Handle("/", fs)

http.ListenAndServe( ':8080' , None )
