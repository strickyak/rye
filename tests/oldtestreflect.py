from go import reflect
from go.html import template
from go import html.template as T

string_t = reflect.TypeOf('abc')
must string_t.String() == 'string'
html_t = reflect.TypeOf(go_cast(template.HTML, 'abc'))
must html_t.String() == 'template.HTML'
t_html_t = reflect.TypeOf(go_cast(T.HTML, 'abc'))
must t_html_t.String() == 'template.HTML'

string_chan_t = reflect.ChanOf(3, string_t)
must string_chan_t.String() == 'chan string'
html_chan_t = reflect.ChanOf(3, html_t)
must html_chan_t.String() == 'chan template.HTML'

c = reflect.MakeChan(string_chan_t, 3)
say c

r = reflect.ValueOf('Dipsy')
c.Send(reflect.ValueOf('Dipsy'))
val, ok = c.Recv()
must ok

must str(val.Interface()) == 'Dipsy'
# assert val.Interface().String() == 'Dipsy'

print "testreflect OKAY."

int64_t = go_type(int64)
zero = int(reflect.Zero(int64_t).Interface())
must 3 == 3 + zero
must 0 == 3 * zero
must 3 == 3 - zero
must 0 == zero
must not zero
