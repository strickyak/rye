from go import reflect

string_t = reflect.TypeOf('abc')
must string_t.String() == 'string'

chan_t = reflect.ChanOf(3, string_t)
must chan_t.String() == 'chan string'

c = reflect.MakeChan(chan_t, 3)
say c

r = reflect.ValueOf('Dipsy')
say rye_what(r), r
c.Send(reflect.ValueOf('Dipsy'))
val, ok = c.Recv()
must ok

say rye_what(3.1415)
say rye_what('Po')
say rye_what(go_type(reflect.SelectDir))
say rye_what(go_cast(reflect.SelectDir, 1))
say rye_what(val)
say rye_what(val.Interface())
say rye_what(go_cast(string, val.Interface()))

must str(val.Interface()) == 'Dipsy'
# assert val.Interface().String() == 'Dipsy'

print "testreflect OKAY."
