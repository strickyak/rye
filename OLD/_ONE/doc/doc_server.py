"""Example Doc server, written in Rye."""

from go import bytes, log, os, regexp
from go import html/template, io/ioutil, net/http
from go import path/filepath as FP

from go import github.com/microcosm-cc/bluemonday
from go import github.com/russross/blackfriday


SIMPLE_URL = regexp.MustCompile('^/([-A-Za-z0-9_]*)$').FindStringSubmatch

def ExpandMacros(s):
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
  page = page if page else 'Home'
  t = ioutil.ReadFile('%s.md' % page)
  t = RenderMarkdown(t)
  t = ExpandMacros(t)
  z = '''<html><head>
    <META http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <title>%s</title>
    <style type="text/css">
      body {
        font-family: Verdana, Geneva, sans-serif;
        // font-family: "Palatino Linotype", "Book Antiqua", Palatino, serif;
      }
      code {
        font-style: bold;
      }
    </style>
    </head><body>
  ''' % page + TopNav() + t

  if page == 'Home':
    z += TableOfContents()

  return z

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

TITLE = regexp.MustCompile('^##(.*)$').FindStringSubmatch
SECTION = regexp.MustCompile('^/1(.*)$').FindStringSubmatch
Titles, Sections = {}, {}
def ScanForTitles():
  for fname in FP.Glob('*.md'):
    slug = fname[:-3]
    say fname
    g = ioutil.ReadFile(fname)
    for s in str(g).split('\n'):
      mt = TITLE(s)
      ms = SECTION(s)
      switch:
        case mt:
          Titles[slug] = mt[1].strip()
        case ms:
          v = Sections.get(slug, [])
          v.append(ms[1].strip())
          Sections[slug] = v

def TableOfContents():
  w = bytes.NewBuffer(None)
  print >>w, '<table width=100%>'
  print >>w, '<tr>'
  for slug, title in sorted(Titles.items()):
    if slug == 'Home': continue
    print >>w, '<td><b>[<a href="/%s">%s</a>]</b>' % (slug, slug)
  #print >>w, '<tr>'
  #for slug, title in sorted(Titles.items()):
  #  if slug == 'Home': continue
  #  print >>w, '<td><b><a href="/%s">%s</a></b>' % (slug, title)
  print >>w, '<tr valign=top>'
  for slug, title in sorted(Titles.items()):
    if slug == 'Home': continue
    print >>w, '<td>'
    for s in Sections[slug]:
      print >>w, '<small><a href="/%s#%s">%s</a><br></small>' % (slug, s, s)
  print >>w, '</table>'
  return str(w)

def TopNav():
  w = bytes.NewBuffer(None)
  print >>w, '[<a href="Home">Rye Home</a>] &nbsp; &nbsp; '
  for slug, _ in sorted(Sections.items()):
    print >>w, '[<a href="/%s">%s</a>] &nbsp; &nbsp; ' % (slug, slug)
  print >>w, '<hr>'
  return str(w)

def main(args):
  ScanForTitles()
  http.HandleFunc("/", MainHandler)
  log.Fatal(http.ListenAndServe(":8080", None))
