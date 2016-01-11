"""Example Doc server, written in Rye."""

from go import log, os, regexp
from go import html/template, io/ioutil, net/http

from go import github.com/microcosm-cc/bluemonday
from go import github.com/russross/blackfriday


SIMPLE_URL = regexp.MustCompile('^/([-A-Za-z0-9_]*)$').FindStringSubmatch

#def XPreProcess(s):
#  lines = str(s).split('\n')
#  z = []
#  for line in lines:
#    if line.startswith('/title'):
#      z.append('# %s' % line[6:])
#    elif line.startswith('/1'):
#      z.append('</table cellpadding="0" border="1"><table><tr><td cellpadding="50"> %s' % line[2:])
#    elif line.startswith('/2'):
#      z.append('<td cellpadding="50"> &nbsp; %s' % line[2:])
#    elif line.startswith('/3'):
#      z.append('<td cellpadding="50"> &nbsp; %s' % line[2:])
#    else:
#      z.append(line)
#  return '\n'.join(z)

def PreProcess(s):
  s = str(s)
  s = s.replace('/0', '<table cellpadding=10 cellspacing=5 border=0>') 
  s = s.replace('/1', '<tr><td colspan=3 cellpadding=20><hr> <tr><td cellpadding=20><b>') 
  s = s.replace('/2', '</b><td cellpadding=20>') 
  s = s.replace('/3', '<td cellpadding=20>') 
  return go_cast(template.HTML, s)

def RenderMarkdown(s):
  t = blackfriday.MarkdownCommon(s)
  html = bluemonday.UGCPolicy().SanitizeBytes(t)
  return go_cast(template.HTML, html)

def RenderPage(page):
  page = page if page else 'home'
  t = ioutil.ReadFile('%s.md' % page)
  t = RenderMarkdown(t)
  t = PreProcess(t)
  return '''<html><head>
<style type="text/css">
  body {
    font-family: Verdana, Geneva, sans-serif;
  }
</style>
</head><body>
  ''' + t


def XXRenderPage(page):
  page = page if page else 'home'
  z = []
  lines = str(ioutil.ReadFile('%s.m' % page)).split('\n')
  in_code = False
  title = 'Rye Documentation'
  for line in lines:
    line = line.rstrip()
    if line.startswith('    '):
      # Preformatted Code.
      if not in_code:
        z.append('<pre>')
        in_code = True
    else:
      # Not Preformatted Code.
      if in_code:
        z.append('</pre>')
        in_code = False

      if line.startswith('/'):
        if line.startswith('/title'):
          title = line[len('/title'):]
        elif line.startswith('/1'):
          z.append('<tr><td>')
          z.append(line)
        elif line.startswith('/2'):
          z.append('<td>')
          z.append(line)
        elif line.startswith('/3'):
          z.append('<td>')
          z.append(line)
        else:
          raise Exception("Unknown tag: %q" % line)
      elif line.strip() == '':
        z.append('<p>')
      else:
        z.append(line)
    
  return '''<html>
<head><title>%s</title></head><body>
%s
''' % (title, '\n'.join(z))
  

def MainHandler(w, r):
  m = SIMPLE_URL(r.URL.Path)
  if m:
    try:
      ht = RenderPage(m[1])
      w.Header().Set('Content-Type', 'text/html; charset=UTF-8')
      print >>w, ht
    except Exception as ex:
      w.Header().Set('Content-Type', 'text/plain; charset=UTF-8')
      print >>w, 'Caught Exception:', ex
      
  else:
    w.Header().Set('Content-Type', 'text/plain; charset=UTF-8')
    w.WriteHeader(http.StatusNotFound)
    print >>w, '404 NOT FOUND'

def main(args):
  http.HandleFunc("/", MainHandler)
  log.Fatal(http.ListenAndServe(":8080", None))
