// +build prego

 package two
 import . "github.com/strickyak/rye"

// cwp: BUILTINS
// path: BUILTINS
// thispkg: rye__/BUILTINS
// modname: BUILTINS
// internal: <open file 'gen_internals.py', mode 'w' at 0x7f20c05c6810>

 import "fmt"
 import "io"
 import "os"
 import "reflect"
 import "runtime"
 import "time"
 import "bytes"
 import "unsafe"
 import i_bytes "bytes" // <parse.Timport object at 0x7f20c05e33d0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'bytes'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'bytes', 'pkg': 'bytes', 'line': 1, 'where': 0}
 import i_strings "strings" // <parse.Timport object at 0x7f20be1299d0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'strings'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'strings', 'pkg': 'strings', 'line': 1, 'where': 0}
 import i_unicode "unicode" // <parse.Timport object at 0x7f20be129a10>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'unicode'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'unicode', 'pkg': 'unicode', 'line': 1, 'where': 0}
 import i_bufio "bufio" // <parse.Timport object at 0x7f20be129a50>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'bufio'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'bufio', 'pkg': 'bufio', 'line': 2, 'where': 39}
 import i_io "io" // <parse.Timport object at 0x7f20be129a90>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'io'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'io', 'pkg': 'io', 'line': 2, 'where': 39}
 import i_ioutil "io/ioutil" // <parse.Timport object at 0x7f20be129ad0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'io', 'ioutil'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'ioutil', 'pkg': 'io/ioutil', 'line': 2, 'where': 39}
 import i_os "os" // <parse.Timport object at 0x7f20be129b10>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'os'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'os', 'pkg': 'os', 'line': 2, 'where': 39}
var _ = i_bufio.ErrNegativeCount // bufio
var _ = i_bytes.ErrTooLarge // bytes
var _ = i_io.ErrUnexpectedEOF // io
var _ = i_ioutil.Discard // io/ioutil
var _ = i_os.Stdout // os
var _ = i_strings.TrimSpace // strings
var _ = i_unicode.Zs // unicode
 var _ = fmt.Sprintf
 var _ = io.EOF
 var _ = os.Stderr
 var _ = reflect.ValueOf
 var _ = runtime.Stack
 var _ = time.Sleep
 var _ = bytes.Split
 var _ = unsafe.Sizeof(0)
 var _ = MkInt

 func jinner_eval_module () (U,V) {
// @ 0 @ 1 @ 
//Vimport: ['go', 'bytes'] bytes go
// $ 0 $ 1 $
// @ 0 @ 1 @ 
//Vimport: ['go', 'strings'] strings go
// $ 0 $ 1 $
// @ 0 @ 1 @ 
//Vimport: ['go', 'unicode'] unicode go
// $ 0 $ 1 $
// @ 39 @ 2 @ 
//Vimport: ['go', 'bufio'] bufio go
// $ 39 $ 2 $
// @ 39 @ 2 @ 
//Vimport: ['go', 'io'] io go
// $ 39 $ 2 $
// @ 39 @ 2 @ 
//Vimport: ['go', 'io', 'ioutil'] ioutil go
// $ 39 $ 2 $
// @ 39 @ 2 @ 
//Vimport: ['go', 'os'] os go
// $ 39 $ 2 $
// @ 2209 @ 101 @ 
// $ 2209 $ 101 $
// @ 2377 @ 108 @ 
// $ 2377 $ 108 $
// @ 2599 @ 116 @ 
// $ 2599 $ 116 $
// @ 2673 @ 122 @ 
// $ 2673 $ 122 $
// @ 2747 @ 128 @ 
// $ 2747 $ 128 $
// @ 2825 @ 134 @ 
// $ 2825 $ 134 $
// @ 2911 @ 139 @ 
// $ 2911 $ 139 $
// @ 2995 @ 144 @ 
// $ 2995 $ 144 $
// @ 3077 @ 149 @ 
// $ 3077 $ 149 $
// @ 3473 @ 173 @ 
// $ 3473 $ 173 $
// @ 3621 @ 182 @ 
// $ 3621 $ 182 $
// @ 3713 @ 186 @ 
// $ 3713 $ 186 $
// @ 3907 @ 198 @ 
// $ 3907 $ 198 $
// @ 7004 @ 352 @ 
// $ 7004 $ 352 $
// @ 17994 @ 764 @ 
// $ 17994 $ 764 $
// @ 18147 @ 772 @ 
// LitIntern: interning "J_object is a fake class to hold methods for the builtin class object.": jlitS_f78e09fad53439d480c756b091d4516f = MkStrJ( "J_object is a fake class to hold methods for the builtin class object." )
// Yprim: constructing Ylit with {'s': 'jlitS_f78e09fad53439d480c756b091d4516f_1,jlitS_f78e09fad53439d480c756b091d4516f_2', 'flavor': '', 't': '"J_object is a fake class to hold methods for the builtin class object."'}
// ######################### BAD Comma in Yprim: jlitS_f78e09fad53439d480c756b091d4516f_1,jlitS_f78e09fad53439d480c756b091d4516f_2
// $ 18147 $ 772 $
// @ 22100 @ 907 @ 
// $ 22100 $ 907 $
   return NoneJ_1, NoneJ_2
 }

//(begin tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_ord(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 2223 @ 102 @ ord
// { native F
t1 := StringJ(a_x_1, a_x_2)
t2 := int(t1[0])
return macro.MkintJ(t2)
// } native F
// $ 2223 $ 102 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: ord

var jspecFunc_ord = JCallSpec{Name: "ord", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_ord (a0_1 U, a0_2 V) (U,V) { return H_1_ord(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_chr(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 2391 @ 109 @ chr
// { native F
t1 := macro.TakeIntJ(a_x_1, a_x_2)
s := string([]byte{byte(t1)})
return macro.MkStrJ(s)
// } native F
// $ 2391 $ 109 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: chr

var jspecFunc_chr = JCallSpec{Name: "chr", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_chr (a0_1 U, a0_2 V) (U,V) { return H_1_chr(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('vec', None), ('start', None)]
// typPlus: [None, None]
// self.scope [vec] = vec
// self.scope [start] = start
///////////////////////////////

 func H_2_sum(a_vec_1 U, a_vec_2 V, a_start_1 U, a_start_2 V, ) (U,V) {
// Emit local variables: considering:  start a_start
// Emit local variables: considering more:  start a_start
// Emit local variables: considering:  vec a_vec
// Emit local variables: considering more:  vec a_vec
// Emit local variables: considering:  x /*Yvar.str*/v_x_1, v_x_2
// Emit local variables: considering more:  x /*Yvar.str*/v_x_1, v_x_2
   var v_x_1 U = NoneJ_1; var v_x_2 V = NoneJ_2; _, _ = /*Yvar.str*/v_x_1, v_x_2
// Emit local variables: considering:  z /*Yvar.str*/v_z_1, v_z_2
// Emit local variables: considering more:  z /*Yvar.str*/v_z_1, v_z_2
   var v_z_1 U = NoneJ_1; var v_z_2 V = NoneJ_2; _, _ = /*Yvar.str*/v_z_1, v_z_2
// @ 2624 @ 117 @ sum
// Vvar: local var start -> str: 'a_start'
// Yprim: constructing Yvar with {'s': 'a_start', 'flavor': '', 't': <parse.Tvar object at 0x7f20be129e10>}
// Yprim: constructing Yvar with {'s': 'v_z', 'flavor': '', 't': False}
// @@@@@@ Creating var "z" in scope @@@@@@ -> Yvar: <codegen2.Yvar object at 0x7f20be0d8710>
// Calling DoAssign because isinstance(lhs <class 'codegen2.Yvar'>, Ybase)
/*G.DoAssign*/ /*Yvar.str*/v_z_1, v_z_2 = /*Yvar.str*/a_start_1, a_start_2
// $ 2624 $ 117 $
// @ 2636 @ 118 @ sum
// Vvar: local var vec -> str: 'a_vec'
// Yprim: constructing Yvar with {'s': 'a_vec', 'flavor': '', 't': <parse.Tvar object at 0x7f20be129ed0>}
// Vfor: var=<parse.Tvar object at 0x7f20be129e90> i=__101 ptv=/*Yvar.str*/a_vec_1, a_vec_2

   for_returning__101_1, for_returning__101_2 := func () (U,V) { // around FOR
     var nextee PJ = macro.TakePJ(/*Yvar.str*/a_vec_1, a_vec_2)
     var nexter__101 JNexter = nextee.Iter()
     enougher__101, canEnough__101 := nexter__101.(Enougher)
     if canEnough__101 {
             defer enougher__101.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx___101_1, ndx___101_2, more___101 := nexter__101.NextJ()
       if !more___101 {
         break
       }
       // BEGIN FOR

// Yprim: constructing Yvar with {'s': 'v_x', 'flavor': '', 't': False}
// @@@@@@ Creating var "x" in scope @@@@@@ -> Yvar: <codegen2.Yvar object at 0x7f20be0d8850>
// Calling DoAssign because isinstance(lhs <class 'codegen2.Yvar'>, Ybase)
/*G.DoAssign*/ /*Yvar.str*/v_x_1, v_x_2 = ndx___101_1, ndx___101_2
// @ 2654 @ 119 @ sum
// Vvar: local var z -> Yvar: <codegen2.Yvar object at 0x7f20be0d8710>
// Yprim: constructing Yvar with {'s': '/*Yvar.str*/v_z_1, v_z_2', 'flavor': '', 't': <parse.Tvar object at 0x7f20be129f10>}
// ######################### BAD Comma in Yprim: /*Yvar.str*/v_z_1, v_z_2
// Vvar: local var x -> Yvar: <codegen2.Yvar object at 0x7f20be0d8850>
// Yprim: constructing Yvar with {'s': '/*Yvar.str*/v_x_1, v_x_2', 'flavor': '', 't': <parse.Tvar object at 0x7f20be129f50>}
// ######################### BAD Comma in Yprim: /*Yvar.str*/v_x_1, v_x_2
uu_102_1, uu_102_2 := /*Yvar.str*//*Yvar.str*/v_z_1, v_z_2
 _,_ = uu_102_1,uu_102_2
uu_103_1, uu_103_2 := /*Yvar.str*//*Yvar.str*/v_x_1, v_x_2
 _,_ = uu_103_1,uu_103_2
// Calling DoAssign because isinstance(lhs <class 'codegen2.Yvar'>, Ybase)
/*G.DoAssign*/ /*Yvar.str*/v_z_1, v_z_2 = /*DoAdd*/ macro.AddJ( uu_102_1,uu_102_2 ,  uu_103_1,uu_103_2 )
// $ 2654 $ 119 $

       // END FOR
     }
     return 0,V0
   }() // around FOR
   if for_returning__101_1 != 0 && for_returning__101_2 != V0 { return for_returning__101_1, for_returning__101_2 }

// $ 2636 $ 118 $
// @ 2663 @ 120 @ sum
// Vvar: local var z -> Yvar: <codegen2.Yvar object at 0x7f20be0d8710>
// Yprim: constructing Yvar with {'s': '/*Yvar.str*/v_z_1, v_z_2', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0c60d0>}
// ######################### BAD Comma in Yprim: /*Yvar.str*/v_z_1, v_z_2
   return /*Yvar.str*//*Yvar.str*/v_z_1, v_z_2 
// $ 2663 $ 120 $

   return NoneJ_1, NoneJ_2
 }

// LitIntern: interning 0: jlitI_cfcd208495d565ef66e7dff9f98764da = MkIntJ(0)
// Yprim: constructing Ylit with {'s': 'jlitI_cfcd208495d565ef66e7dff9f98764da_1,jlitI_cfcd208495d565ef66e7dff9f98764da_2', 'flavor': '', 't': '0'}
// ######################### BAD Comma in Yprim: jlitI_cfcd208495d565ef66e7dff9f98764da_1,jlitI_cfcd208495d565ef66e7dff9f98764da_2
// Nando: "vec", "start" W{0,V0}, W{/*Yint.str*/jlitI_cfcd208495d565ef66e7dff9f98764da_1,jlitI_cfcd208495d565ef66e7dff9f98764da_2}
///////////////////////////////
// name: sum

var jspecFunc_sum = JCallSpec{Name: "sum", Args: []string{"vec", "start"}, Defaults: []W{W{0,V0}, W{/*Yint.str*/jlitI_cfcd208495d565ef66e7dff9f98764da_1,jlitI_cfcd208495d565ef66e7dff9f98764da_2}}, Star: "", StarStar: ""}
func jfnFunc_sum (a0_1 U, a0_2 V,a1_1 U, a1_2 V) (U,V) { return H_2_sum(a0_1, a0_2,a1_1, a1_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
// self.scope [vec] = vec
///////////////////////////////

 func H_1_any(a_vec_1 U, a_vec_2 V, ) (U,V) {
// Emit local variables: considering:  e /*Yvar.str*/v_e_1, v_e_2
// Emit local variables: considering more:  e /*Yvar.str*/v_e_1, v_e_2
   var v_e_1 U = NoneJ_1; var v_e_2 V = NoneJ_2; _, _ = /*Yvar.str*/v_e_1, v_e_2
// Emit local variables: considering:  vec a_vec
// Emit local variables: considering more:  vec a_vec
// @ 2689 @ 123 @ any
// Vvar: local var vec -> str: 'a_vec'
// Yprim: constructing Yvar with {'s': 'a_vec', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0c6210>}
// Vfor: var=<parse.Tvar object at 0x7f20be0c61d0> i=__104 ptv=/*Yvar.str*/a_vec_1, a_vec_2

   for_returning__104_1, for_returning__104_2 := func () (U,V) { // around FOR
     var nextee PJ = macro.TakePJ(/*Yvar.str*/a_vec_1, a_vec_2)
     var nexter__104 JNexter = nextee.Iter()
     enougher__104, canEnough__104 := nexter__104.(Enougher)
     if canEnough__104 {
             defer enougher__104.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx___104_1, ndx___104_2, more___104 := nexter__104.NextJ()
       if !more___104 {
         break
       }
       // BEGIN FOR

// Yprim: constructing Yvar with {'s': 'v_e', 'flavor': '', 't': False}
// @@@@@@ Creating var "e" in scope @@@@@@ -> Yvar: <codegen2.Yvar object at 0x7f20be0d8750>
// Calling DoAssign because isinstance(lhs <class 'codegen2.Yvar'>, Ybase)
/*G.DoAssign*/ /*Yvar.str*/v_e_1, v_e_2 = ndx___104_1, ndx___104_2
// @ 2707 @ 124 @ any
// Vvar: local var e -> Yvar: <codegen2.Yvar object at 0x7f20be0d8750>
// Yprim: constructing Yvar with {'s': '/*Yvar.str*/v_e_1, v_e_2', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0c6250>}
// ######################### BAD Comma in Yprim: /*Yvar.str*/v_e_1, v_e_2
   if /*AsBool*/BoolJ(/*Yvar.str*//*Yvar.str*/v_e_1, v_e_2) {
// @ 2719 @ 125 @ any
   return /*Ybool.str*/TrueJ_1, TrueJ_2 
// $ 2719 $ 125 $
   }
// $ 2707 $ 124 $

       // END FOR
     }
     return 0,V0
   }() // around FOR
   if for_returning__104_1 != 0 && for_returning__104_2 != V0 { return for_returning__104_1, for_returning__104_2 }

// $ 2689 $ 123 $
// @ 2733 @ 126 @ any
   return /*Ybool.str*/FalseJ_1, FalseJ_2 
// $ 2733 $ 126 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "vec" W{0,V0}
///////////////////////////////
// name: any

var jspecFunc_any = JCallSpec{Name: "any", Args: []string{"vec"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_any (a0_1 U, a0_2 V) (U,V) { return H_1_any(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
// self.scope [vec] = vec
///////////////////////////////

 func H_1_all(a_vec_1 U, a_vec_2 V, ) (U,V) {
// Emit local variables: considering:  e /*Yvar.str*/v_e_1, v_e_2
// Emit local variables: considering more:  e /*Yvar.str*/v_e_1, v_e_2
   var v_e_1 U = NoneJ_1; var v_e_2 V = NoneJ_2; _, _ = /*Yvar.str*/v_e_1, v_e_2
// Emit local variables: considering:  vec a_vec
// Emit local variables: considering more:  vec a_vec
// @ 2763 @ 129 @ all
// Vvar: local var vec -> str: 'a_vec'
// Yprim: constructing Yvar with {'s': 'a_vec', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0c6550>}
// Vfor: var=<parse.Tvar object at 0x7f20be0c6510> i=__105 ptv=/*Yvar.str*/a_vec_1, a_vec_2

   for_returning__105_1, for_returning__105_2 := func () (U,V) { // around FOR
     var nextee PJ = macro.TakePJ(/*Yvar.str*/a_vec_1, a_vec_2)
     var nexter__105 JNexter = nextee.Iter()
     enougher__105, canEnough__105 := nexter__105.(Enougher)
     if canEnough__105 {
             defer enougher__105.Enough()
     }
     // else case without Enougher will be faster.
     for {
       ndx___105_1, ndx___105_2, more___105 := nexter__105.NextJ()
       if !more___105 {
         break
       }
       // BEGIN FOR

// Yprim: constructing Yvar with {'s': 'v_e', 'flavor': '', 't': False}
// @@@@@@ Creating var "e" in scope @@@@@@ -> Yvar: <codegen2.Yvar object at 0x7f20be0d8990>
// Calling DoAssign because isinstance(lhs <class 'codegen2.Yvar'>, Ybase)
/*G.DoAssign*/ /*Yvar.str*/v_e_1, v_e_2 = ndx___105_1, ndx___105_2
// @ 2781 @ 130 @ all
// Vvar: local var e -> Yvar: <codegen2.Yvar object at 0x7f20be0d8990>
// Yprim: constructing Yvar with {'s': '/*Yvar.str*/v_e_1, v_e_2', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0c6590>}
// ######################### BAD Comma in Yprim: /*Yvar.str*/v_e_1, v_e_2
   if (/*Vboolop*/ !(/*AsBool*/BoolJ(/*Yvar.str*//*Yvar.str*/v_e_1, v_e_2)))  {
// @ 2797 @ 131 @ all
   return /*Ybool.str*/FalseJ_1, FalseJ_2 
// $ 2797 $ 131 $
   }
// $ 2781 $ 130 $

       // END FOR
     }
     return 0,V0
   }() // around FOR
   if for_returning__105_1 != 0 && for_returning__105_2 != V0 { return for_returning__105_1, for_returning__105_2 }

// $ 2763 $ 129 $
// @ 2812 @ 132 @ all
   return /*Ybool.str*/TrueJ_1, TrueJ_2 
// $ 2812 $ 132 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "vec" W{0,V0}
///////////////////////////////
// name: all

var jspecFunc_all = JCallSpec{Name: "all", Args: []string{"vec"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_all (a0_1 U, a0_2 V) (U,V) { return H_1_all(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_len(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 2839 @ 135 @ len
// { native F
__q := LenJ(a_x_1, a_x_2)
return macro.MkintJ(__q)
// } native F
// $ 2839 $ 135 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: len

var jspecFunc_len = JCallSpec{Name: "len", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_len (a0_1 U, a0_2 V) (U,V) { return H_1_len(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_repr(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 2926 @ 140 @ repr
// { native F
z := ReprJ(a_x_1, a_x_2)
return macro.MkStrJ(z)
// } native F
// $ 2926 $ 140 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: repr

var jspecFunc_repr = JCallSpec{Name: "repr", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_repr (a0_1 U, a0_2 V) (U,V) { return H_1_repr(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_str(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 3009 @ 145 @ str
// { native F
z := StrJ(a_x_1, a_x_2)
return macro.MkStrJ(z)
// } native F
// $ 3009 $ 145 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: str

var jspecFunc_str = JCallSpec{Name: "str", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_str (a0_1 U, a0_2 V) (U,V) { return H_1_str(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_int(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 3091 @ 150 @ int
// { native F
z := ForceIntJ(a_x_1, a_x_2)
return macro.MkIntJ(z)
// } native F
// $ 3091 $ 150 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: int

var jspecFunc_int = JCallSpec{Name: "int", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_int (a0_1 U, a0_2 V) (U,V) { return H_1_int(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_list(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 3488 @ 174 @ list
// { native F
z := MkListJ(ListJ(a_x_1, a_x_2))
return inline.MkPJ(z)
// } native F
// $ 3488 $ 174 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: list

var jspecFunc_list = JCallSpec{Name: "list", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_list (a0_1 U, a0_2 V) (U,V) { return H_1_list(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('args', None), ('kv', None)]
// typPlus: [None, None]
// self.scope [args] = args
// self.scope [kv] = kv
///////////////////////////////

 func H_0V_dict(  a_args_1 U, a_args_2 V, a_kv_1 U, a_kv_2 V) (U,V) {
// Emit local variables: considering:  args a_args
// Emit local variables: considering more:  args a_args
// Emit local variables: considering:  kv a_kv
// Emit local variables: considering more:  kv a_kv
// @ 3646 @ 183 @ dict
// { native F
return N2_dict(a_args_1, a_args_2, a_kv_1, a_kv_2)
// } native F
// $ 3646 $ 183 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: dict

var jspecFunc_dict = JCallSpec{Name: "dict", Args: []string{}, Defaults: []W{}, Star: "args", StarStar: "kv"}
 type jpFunc_dict struct { JPNewCallable }
 func (o *jpFunc_dict) Contents() interface{} {
   return macro.TakePJ(H_dict_1, H_dict_2)
 }

 func (o jpFunc_dict) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U,V) {
   argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
   _, _, _ = argv, star, starstar
   su, sv := inline.MkPJ(star)
   ssu, ssv := inline.MkPJ(starstar)
   return H_0V_dict( su, sv, ssu, ssv)
 }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_tuple(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 3729 @ 187 @ tuple
// { native F

    x := inline.TakePJ(a_x_1, a_x_2)
    l := x.List()
    t := MkTupleJ(l)
    return inline.MkPJ(t)
  
// } native F
// $ 3729 $ 187 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: tuple

var jspecFunc_tuple = JCallSpec{Name: "tuple", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_tuple (a0_1 U, a0_2 V) (U,V) { return H_1_tuple(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
// self.scope [x] = x
///////////////////////////////

 func H_1_type(a_x_1 U, a_x_2 V, ) (U,V) {
// Emit local variables: considering:  x a_x
// Emit local variables: considering more:  x a_x
// @ 3922 @ 199 @ type
// { native F

    switch inline.Tag(a_x_1, a_x_2) {
    case Int:
      //a := inline.TakeIntJ(a_x_1, a_x_2)
      return H_int_1, H_int_2
    case Str:
      //a := inline.TakeStrJ(a_x_1, a_x_2)
      return H_str_1, H_str_2
    case Py:
      a := inline.TakePJ(a_x_1, a_x_2)
      return a.PType()
    }
    panic("Bad")
  
// } native F
// $ 3922 $ 199 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "x" W{0,V0}
///////////////////////////////
// name: type

var jspecFunc_type = JCallSpec{Name: "type", Args: []string{"x"}, Defaults: []W{W{0,V0}}, Star: "", StarStar: ""}
func jfnFunc_type (a0_1 U, a0_2 V) (U,V) { return H_1_type(a0_1, a0_2) }


//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func (self *JDict) N_0_clear( ) (U,V) {
// @ 7044 @ 354 @ JDict.clear
// { native F
self.mu.Lock()
self.ppp = make(JScope)
self.mu.Unlock()
// } native F
// $ 7044 $ 354 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: clear

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func (self *JDict) N_0_copy( ) (U,V) {
// @ 7151 @ 360 @ JDict.copy
// { native F
z := make(JScope)
self.mu.Lock()
for k, v := range self.ppp { z[k] = v }
self.mu.Unlock()
p := MkDictJ(z)
return inline.MkPJ(p)
// } native F
// $ 7151 $ 360 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: copy

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func (self *JDict) N_0_items( ) (U,V) {
// @ 7355 @ 369 @ JDict.items
// { native F
z := make([]W, 0, len(self.ppp))
self.mu.Lock()
for k, v := range self.ppp {
  key_1, key_2 := MkStrJ(k)
  tup := MkTupleJ([]W{ W{key_1,key_2}, v})
  u, v := inline.MkPJ(tup)
  z = append(z, W{u,v})
}
self.mu.Unlock()
p := MkListJ(z)
return inline.MkPJ(p)
// } native F
// $ 7355 $ 369 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: items

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func (self *JDict) N_0_iteritems( ) (U,V) {
// @ 7731 @ 383 @ JDict.iteritems
// Vcall: fn: <parse.Tfield object at 0x7f20be0cf4d0>
// Vcall: args: []
// Vcall: names: []
// Vcall: star: None
// Vcall: starstar: None
// Yprim: constructing Yself with {'s': 'self', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cf490>}
uu_106_1, uu_106_2 := /*General*/ /*invoker*/ JF_INVOKE_0_items(/*Yself.str*/inline.MkPJ(&self.JBase), ) 
   return  uu_106_1,uu_106_2  
// $ 7731 $ 383 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: iteritems

//(tail)
// zip(p.argsPlus, typPlus): [('key', None), ('default', None)]
// typPlus: [None, None]
// self.scope [key] = key
// self.scope [default] = default
///////////////////////////////

 func (self *JDict) N_2_get(a_key_1 U, a_key_2 V, a_default_1 U, a_default_2 V, ) (U,V) {
// Emit local variables: considering:  default a_default
// Emit local variables: considering more:  default a_default
// Emit local variables: considering:  key a_key
// Emit local variables: considering more:  key a_key
// @ 8320 @ 408 @ JDict.get
// { native F
k := StringJ(a_key_1, a_key_2)
self.mu.Lock()
z, ok := self.ppp[k]
self.mu.Unlock()
if ok { return z.U, z.V }
return a_default_1, a_default_2
// } native F
// $ 8320 $ 408 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "key", "default" W{0,V0}, W{NoneJ_1, NoneJ_2}
///////////////////////////////
// name: get

//(tail)



var specMeth_0_JDict__clear = JCallSpec{Name: "JDict::clear", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}

func (o *JDict) GET_clear() (U,V) {
  z := (&JCall0{JPNewCallable{JCallSpec:&specMeth_0_JDict__clear}, o.N_0_clear})
  return inline.ForgeUV(z)
}
var specMeth_0_JDict__copy = JCallSpec{Name: "JDict::copy", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}

func (o *JDict) GET_copy() (U,V) {
  z := (&JCall0{JPNewCallable{JCallSpec:&specMeth_0_JDict__copy}, o.N_0_copy})
  return inline.ForgeUV(z)
}
var specMeth_2_JDict__get = JCallSpec{Name: "JDict::get", Args: []string{"key", "default"}, Defaults: []W{W{0,V0}, W{NoneJ_1, NoneJ_2}}, Star: "", StarStar: ""}

func (o *JDict) GET_get() (U,V) {
  z := (&JCall2{JPNewCallable{JCallSpec:&specMeth_2_JDict__get}, o.N_2_get})
  return inline.ForgeUV(z)
}
var specMeth_0_JDict__items = JCallSpec{Name: "JDict::items", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}

func (o *JDict) GET_items() (U,V) {
  z := (&JCall0{JPNewCallable{JCallSpec:&specMeth_0_JDict__items}, o.N_0_items})
  return inline.ForgeUV(z)
}
var specMeth_0_JDict__iteritems = JCallSpec{Name: "JDict::iteritems", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}

func (o *JDict) GET_iteritems() (U,V) {
  z := (&JCall0{JPNewCallable{JCallSpec:&specMeth_0_JDict__iteritems}, o.N_0_iteritems})
  return inline.ForgeUV(z)
}

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func H_0_object( ) (U,V) {
// @ 18010 @ 765 @ object
// LitIntern: interning "object is the construtor for builtin object type.": jlitS_c3e2c5b4de9c3a86f8baadfd17315a39 = MkStrJ( "object is the construtor for builtin object type." )
// Yprim: constructing Ylit with {'s': 'jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_1,jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_2', 'flavor': '', 't': '"object is the construtor for builtin object type."'}
// ######################### BAD Comma in Yprim: jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_1,jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_2
// $ 18010 $ 765 $
// @ 18064 @ 766 @ object
// { native F

    jo := J_object{}
    jb := jo.GetJBase()
    return ForgeUV(jb)
  
// } native F
// $ 18064 $ 766 $

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: object

var jspecFunc_object = JCallSpec{Name: "object", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}
func jfnFunc_object () (U,V) { return H_0_object() }


//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

 func (self *J_object) N_0___init__( ) (U,V) {

   return NoneJ_1, NoneJ_2
 }

// Nando:  
///////////////////////////////
// name: __init__

//(tail)



var specMeth_0_J_object____init__ = JCallSpec{Name: "J_object::__init__", Args: []string{}, Defaults: []W{}, Star: "", StarStar: ""}

func (o *J_object) GET___init__() (U,V) {
  z := (&JCall0{JPNewCallable{JCallSpec:&specMeth_0_J_object____init__}, o.N_0___init__})
  return inline.ForgeUV(z)
}

//(tail)
// zip(p.argsPlus, typPlus): [('f', None), ('a', None), ('b', None), ('c', None), ('d', None)]
// typPlus: [None, None, None, None, None]
// self.scope [f] = f
// self.scope [a] = a
// self.scope [b] = b
// self.scope [c] = c
// self.scope [d] = d
///////////////////////////////

 func H_5__force_generation_of_call_4_(a_f_1 U, a_f_2 V, a_a_1 U, a_a_2 V, a_b_1 U, a_b_2 V, a_c_1 U, a_c_2 V, a_d_1 U, a_d_2 V, ) (U,V) {
// Emit local variables: considering:  a a_a
// Emit local variables: considering more:  a a_a
// Emit local variables: considering:  b a_b
// Emit local variables: considering more:  b a_b
// Emit local variables: considering:  c a_c
// Emit local variables: considering more:  c a_c
// Emit local variables: considering:  d a_d
// Emit local variables: considering more:  d a_d
// Emit local variables: considering:  f a_f
// Emit local variables: considering more:  f a_f
// @ 22151 @ 908 @ _force_generation_of_call_4_
// Vcall: fn: <parse.Tvar object at 0x7f20be0cfa90>
// Vcall: args: [<parse.Tvar object at 0x7f20be0cfb10>, <parse.Tvar object at 0x7f20be0cfb50>, <parse.Tvar object at 0x7f20be0cfb90>, <parse.Tvar object at 0x7f20be0cfbd0>]
// Vcall: names: ['', '', '', '']
// Vcall: star: None
// Vcall: starstar: None
// Vvar: local var f -> str: 'a_f'
// Yprim: constructing Yvar with {'s': 'a_f', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cfa90>}
// Vvar: local var a -> str: 'a_a'
// Yprim: constructing Yvar with {'s': 'a_a', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cfb10>}
uu_107_1, uu_107_2 := /*Yvar.str*/a_a_1, a_a_2
 _,_ = uu_107_1,uu_107_2
// Vvar: local var b -> str: 'a_b'
// Yprim: constructing Yvar with {'s': 'a_b', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cfb50>}
uu_108_1, uu_108_2 := /*Yvar.str*/a_b_1, a_b_2
 _,_ = uu_108_1,uu_108_2
// Vvar: local var c -> str: 'a_c'
// Yprim: constructing Yvar with {'s': 'a_c', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cfb90>}
uu_109_1, uu_109_2 := /*Yvar.str*/a_c_1, a_c_2
 _,_ = uu_109_1,uu_109_2
// Vvar: local var d -> str: 'a_d'
// Yprim: constructing Yvar with {'s': 'a_d', 'flavor': '', 't': <parse.Tvar object at 0x7f20be0cfbd0>}
uu_110_1, uu_110_2 := /*Yvar.str*/a_d_1, a_d_2
 _,_ = uu_110_1,uu_110_2
uu_111_1, uu_111_2 := /*L1572*/JCALL_4( /*Yvar.str*/a_f_1, a_f_2,  uu_107_1,uu_107_2 ,  uu_108_1,uu_108_2 ,  uu_109_1,uu_109_2 ,  uu_110_1,uu_110_2  )
   return  uu_111_1,uu_111_2  
// $ 22151 $ 908 $

   return NoneJ_1, NoneJ_2
 }

// Nando: "f", "a", "b", "c", "d" W{0,V0}, W{0,V0}, W{0,V0}, W{0,V0}, W{0,V0}
///////////////////////////////
// name: _force_generation_of_call_4_

var jspecFunc__force_generation_of_call_4_ = JCallSpec{Name: "_force_generation_of_call_4_", Args: []string{"f", "a", "b", "c", "d"}, Defaults: []W{W{0,V0}, W{0,V0}, W{0,V0}, W{0,V0}, W{0,V0}}, Star: "", StarStar: ""}
 type jpFunc__force_generation_of_call_4_ struct { JPNewCallable }
 func (o *jpFunc__force_generation_of_call_4_) Contents() interface{} {
   return macro.TakePJ(H__force_generation_of_call_4__1, H__force_generation_of_call_4__2)
 }
 func (o jpFunc__force_generation_of_call_4_) Call5(a0_1 U, a0_2 V, a1_1 U, a1_2 V, a2_1 U, a2_2 V, a3_1 U, a3_2 V, a4_1 U, a4_2 V) (U,V) {
   return H_5__force_generation_of_call_4_(a0_1, a0_2, a1_1, a1_2, a2_1, a2_2, a3_1, a3_2, a4_1, a4_2)
 }

 func (o jpFunc__force_generation_of_call_4_) JCallV(a1 []W, a2 []W, kv1 []JKV, kv2 map[string]W) (U,V) {
   argv, star, starstar := JNewSpecCall(o.JCallSpec, a1, a2, kv1, kv2)
   _, _, _ = argv, star, starstar
   return H_5__force_generation_of_call_4_(argv[0].U, argv[0].V, argv[1].U, argv[1].V, argv[2].U, argv[2].V, argv[3].U, argv[3].V, argv[4].U, argv[4].V)
 }


//(end tail)

var H__force_generation_of_call_4__1 U; var H__force_generation_of_call_4__2 V // *jpFunc__force_generation_of_call_4_
var H_all_1 U; var H_all_2 V // *jpFunc_all
var H_any_1 U; var H_any_2 V // *jpFunc_any
var H_chr_1 U; var H_chr_2 V // *jpFunc_chr
var H_dict_1 U; var H_dict_2 V // *jpFunc_dict
var H_int_1 U; var H_int_2 V // *jpFunc_int
var H_len_1 U; var H_len_2 V // *jpFunc_len
var H_list_1 U; var H_list_2 V // *jpFunc_list
var H_object_1 U; var H_object_2 V // *jpFunc_object
var H_ord_1 U; var H_ord_2 V // *jpFunc_ord
var H_repr_1 U; var H_repr_2 V // *jpFunc_repr
var H_str_1 U; var H_str_2 V // *jpFunc_str
var H_sum_1 U; var H_sum_2 V // *jpFunc_sum
var H_tuple_1 U; var H_tuple_2 V // *jpFunc_tuple
var H_type_1 U; var H_type_2 V // *jpFunc_type

 func init /*New_Module*/ () {
   H__force_generation_of_call_4__1, H__force_generation_of_call_4__2 = ForgeUV((&jpFunc__force_generation_of_call_4_{JPNewCallable{JCallSpec:&jspecFunc__force_generation_of_call_4_}})) // t=*jpFunc__force_generation_of_call_4_
   H_all_1, H_all_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_all}, jfnFunc_all})) // t=*jpFunc_all
   H_any_1, H_any_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_any}, jfnFunc_any})) // t=*jpFunc_any
   H_chr_1, H_chr_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_chr}, jfnFunc_chr})) // t=*jpFunc_chr
   H_dict_1, H_dict_2 = ForgeUV((&jpFunc_dict{JPNewCallable{JCallSpec:&jspecFunc_dict}})) // t=*jpFunc_dict
   H_int_1, H_int_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_int}, jfnFunc_int})) // t=*jpFunc_int
   H_len_1, H_len_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_len}, jfnFunc_len})) // t=*jpFunc_len
   H_list_1, H_list_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_list}, jfnFunc_list})) // t=*jpFunc_list
   H_object_1, H_object_2 = ForgeUV((&JCall0{JPNewCallable{JCallSpec:&jspecFunc_object}, jfnFunc_object})) // t=*jpFunc_object
   H_ord_1, H_ord_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_ord}, jfnFunc_ord})) // t=*jpFunc_ord
   H_repr_1, H_repr_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_repr}, jfnFunc_repr})) // t=*jpFunc_repr
   H_str_1, H_str_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_str}, jfnFunc_str})) // t=*jpFunc_str
   H_sum_1, H_sum_2 = ForgeUV((&JCall2{JPNewCallable{JCallSpec:&jspecFunc_sum}, jfnFunc_sum})) // t=*jpFunc_sum
   H_tuple_1, H_tuple_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_tuple}, jfnFunc_tuple})) // t=*jpFunc_tuple
   H_type_1, H_type_2 = ForgeUV((&JCall1{JPNewCallable{JCallSpec:&jspecFunc_type}, jfnFunc_type})) // t=*jpFunc_type
   jinner_eval_module()
 }

var jlitI_cfcd208495d565ef66e7dff9f98764da_1 U; var jlitI_cfcd208495d565ef66e7dff9f98764da_2 V; func init() { jlitI_cfcd208495d565ef66e7dff9f98764da_1, jlitI_cfcd208495d565ef66e7dff9f98764da_2 = MkIntJ(0) }
var jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_1 U; var jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_2 V; func init() { jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_1, jlitS_c3e2c5b4de9c3a86f8baadfd17315a39_2 = MkStrJ( "object is the construtor for builtin object type." ) }
var jlitS_f78e09fad53439d480c756b091d4516f_1 U; var jlitS_f78e09fad53439d480c756b091d4516f_2 V; func init() { jlitS_f78e09fad53439d480c756b091d4516f_1, jlitS_f78e09fad53439d480c756b091d4516f_2 = MkStrJ( "J_object is a fake class to hold methods for the builtin class object." ) }

func JF_INVOKE_0___init__(fn_1 U, fn_2 V, ) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_0___init__:         
    return x.N_0___init__()         
  case JI_GET___init__:         
    tu, tv := x.GET___init__()    
    return JCALL_0(tu, tv, )

  case *JGo:                
    return x.Invoke("__init__", ) 
  }
  panic(fmt.Sprintf("Cannot invoke '__init__' with 0 arguments on %v", fn))
}
type JI_INVOKE_0___init__ interface { N_0___init__() (U,V) }
func JF_INVOKE_0_clear(fn_1 U, fn_2 V, ) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_0_clear:         
    return x.N_0_clear()         
  case JI_GET_clear:         
    tu, tv := x.GET_clear()    
    return JCALL_0(tu, tv, )

  case *JGo:                
    return x.Invoke("clear", ) 
  }
  panic(fmt.Sprintf("Cannot invoke 'clear' with 0 arguments on %v", fn))
}
type JI_INVOKE_0_clear interface { N_0_clear() (U,V) }
func JF_INVOKE_0_copy(fn_1 U, fn_2 V, ) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_0_copy:         
    return x.N_0_copy()         
  case JI_GET_copy:         
    tu, tv := x.GET_copy()    
    return JCALL_0(tu, tv, )

  case *JGo:                
    return x.Invoke("copy", ) 
  }
  panic(fmt.Sprintf("Cannot invoke 'copy' with 0 arguments on %v", fn))
}
type JI_INVOKE_0_copy interface { N_0_copy() (U,V) }
func JF_INVOKE_0_items(fn_1 U, fn_2 V, ) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_0_items:         
    return x.N_0_items()         
  case JI_GET_items:         
    tu, tv := x.GET_items()    
    return JCALL_0(tu, tv, )

  case *JGo:                
    return x.Invoke("items", ) 
  }
  panic(fmt.Sprintf("Cannot invoke 'items' with 0 arguments on %v", fn))
}
type JI_INVOKE_0_items interface { N_0_items() (U,V) }
func JF_INVOKE_0_iteritems(fn_1 U, fn_2 V, ) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_0_iteritems:         
    return x.N_0_iteritems()         
  case JI_GET_iteritems:         
    tu, tv := x.GET_iteritems()    
    return JCALL_0(tu, tv, )

  case *JGo:                
    return x.Invoke("iteritems", ) 
  }
  panic(fmt.Sprintf("Cannot invoke 'iteritems' with 0 arguments on %v", fn))
}
type JI_INVOKE_0_iteritems interface { N_0_iteritems() (U,V) }
func JF_INVOKE_2_get(fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V) (U,V) {
  fn := inline.CheckTakePJ(fn_1, fn_2)
  switch x := fn.(type) {   
  case JI_INVOKE_2_get:         
    return x.N_2_get(a_0_1, a_0_2, a_1_1, a_1_2)         
  case JI_GET_get:         
    tu, tv := x.GET_get()    
    return JCALL_2(tu, tv, a_0_1, a_0_2, a_1_1, a_1_2)

  case *JGo:                
    return x.Invoke("get", W{a_0_1, a_0_2}, W{a_1_1, a_1_2}) 
  }
  panic(fmt.Sprintf("Cannot invoke 'get' with 2 arguments on %v", fn))
}
type JI_INVOKE_2_get interface { N_2_get(a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V) (U,V) }

type JI_GET___init__ interface { GET___init__() (U,V) }
func JF_GET___init__(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET___init__:         
    return x.GET___init__()    
  }
   return h.FetchField("__init__") 
}

type JI_GET_clear interface { GET_clear() (U,V) }
func JF_GET_clear(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET_clear:         
    return x.GET_clear()    
  }
   return h.FetchField("clear") 
}

type JI_GET_copy interface { GET_copy() (U,V) }
func JF_GET_copy(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET_copy:         
    return x.GET_copy()    
  }
   return h.FetchField("copy") 
}

type JI_GET_get interface { GET_get() (U,V) }
func JF_GET_get(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET_get:         
    return x.GET_get()    
  }
   return h.FetchField("get") 
}

type JI_GET_items interface { GET_items() (U,V) }
func JF_GET_items(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET_items:         
    return x.GET_items()    
  }
   return h.FetchField("items") 
}

type JI_GET_iteritems interface { GET_iteritems() (U,V) }
func JF_GET_iteritems(h_1 U, h_2 V) (U,V) {
  if inline.Tag(h_1, h_2) != Py { panic("cannot GET Field on int or str") }
  h := inline.TakePJ(h_1, h_2)
  switch x := h.(type) { 
  case JI_GET_iteritems:         
    return x.GET_iteritems()    
  }
   return h.FetchField("iteritems") 
}


  type JI_0 interface { JCall0() (U,V) }
  func JCALL_0 (fn_1 U, fn_2 V, ) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_0:
        return f.JCall0()
      case IJCallV:
        return f.JCallV([]W{}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_1 interface { JCall1( U,V) (U,V) }
  func JCALL_1 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_1:
        return f.JCall1(a_0_1, a_0_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_2 interface { JCall2( U,V,  U,V) (U,V) }
  func JCALL_2 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_2:
        return f.JCall2(a_0_1, a_0_2, a_1_1, a_1_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_3 interface { JCall3( U,V,  U,V,  U,V) (U,V) }
  func JCALL_3 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_3:
        return f.JCall3(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_4 interface { JCall4( U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_4 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_4:
        return f.JCall4(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_5 interface { JCall5( U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_5 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_5:
        return f.JCall5(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_6 interface { JCall6( U,V,  U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_6 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V, a_5_1 U, a_5_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_6:
        return f.JCall6(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2, a_5_1, a_5_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}, W{a_5_1, a_5_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_7 interface { JCall7( U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_7 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V, a_5_1 U, a_5_2 V, a_6_1 U, a_6_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_7:
        return f.JCall7(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2, a_5_1, a_5_2, a_6_1, a_6_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}, W{a_5_1, a_5_2}, W{a_6_1, a_6_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_8 interface { JCall8( U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_8 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V, a_5_1 U, a_5_2 V, a_6_1 U, a_6_2 V, a_7_1 U, a_7_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_8:
        return f.JCall8(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2, a_5_1, a_5_2, a_6_1, a_6_2, a_7_1, a_7_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}, W{a_5_1, a_5_2}, W{a_6_1, a_6_2}, W{a_7_1, a_7_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_9 interface { JCall9( U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_9 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V, a_5_1 U, a_5_2 V, a_6_1 U, a_6_2 V, a_7_1 U, a_7_2 V, a_8_1 U, a_8_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_9:
        return f.JCall9(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2, a_5_1, a_5_2, a_6_1, a_6_2, a_7_1, a_7_2, a_8_1, a_8_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}, W{a_5_1, a_5_2}, W{a_6_1, a_6_2}, W{a_7_1, a_7_2}, W{a_8_1, a_8_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

  type JI_10 interface { JCall10( U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V,  U,V) (U,V) }
  func JCALL_10 (fn_1 U, fn_2 V, a_0_1 U, a_0_2 V, a_1_1 U, a_1_2 V, a_2_1 U, a_2_2 V, a_3_1 U, a_3_2 V, a_4_1 U, a_4_2 V, a_5_1 U, a_5_2 V, a_6_1 U, a_6_2 V, a_7_1 U, a_7_2 V, a_8_1 U, a_8_2 V, a_9_1 U, a_9_2 V) (U,V) {
    //// if fn.X == nil { panic("cannot JCALL on int or str") }
    fn := macro.TakePJ(fn_1, fn_2)
    switch f := fn.(type) {
      case JI_10:
        return f.JCall10(a_0_1, a_0_2, a_1_1, a_1_2, a_2_1, a_2_2, a_3_1, a_3_2, a_4_1, a_4_2, a_5_1, a_5_2, a_6_1, a_6_2, a_7_1, a_7_2, a_8_1, a_8_2, a_9_1, a_9_2)
      case IJCallV:
        return f.JCallV([]W{W{a_0_1, a_0_2}, W{a_1_1, a_1_2}, W{a_2_1, a_2_2}, W{a_3_1, a_3_2}, W{a_4_1, a_4_2}, W{a_5_1, a_5_2}, W{a_6_1, a_6_2}, W{a_7_1, a_7_2}, W{a_8_1, a_8_2}, W{a_9_1, a_9_2}}, nil, nil, nil)
    }
    panic(fmt.Sprintf("No way to call: %v", fn))
  }

// self.signatures.items: {}

//ydefs// _force_generation_of_call_4_ => Yfunc:(BUILTINS._force_generation_of_call_4_) [[ <codegen2.Yfunc object at 0x7f20be0d8490> ]]
//
//ydefs// all => Yfunc:(BUILTINS.all) [[ <codegen2.Yfunc object at 0x7f20be0cff90> ]]
//
//ydefs// any => Yfunc:(BUILTINS.any) [[ <codegen2.Yfunc object at 0x7f20be0cff50> ]]
//
//ydefs// chr => Yfunc:(BUILTINS.chr) [[ <codegen2.Yfunc object at 0x7f20be0cfed0> ]]
//
//ydefs// dict => Yfunc:(BUILTINS.dict) [[ <codegen2.Yfunc object at 0x7f20be0d8150> ]]
//
//ydefs// int => Yfunc:(BUILTINS.int) [[ <codegen2.Yfunc object at 0x7f20be0d80d0> ]]
//
//ydefs// len => Yfunc:(BUILTINS.len) [[ <codegen2.Yfunc object at 0x7f20be0cffd0> ]]
//
//ydefs// list => Yfunc:(BUILTINS.list) [[ <codegen2.Yfunc object at 0x7f20be0d8110> ]]
//
//ydefs// object => Yfunc:(BUILTINS.object) [[ <codegen2.Yfunc object at 0x7f20be0d8210> ]]
//
//ydefs// ord => Yfunc:(BUILTINS.ord) [[ <codegen2.Yfunc object at 0x7f20be0cfe90> ]]
//
//ydefs// repr => Yfunc:(BUILTINS.repr) [[ <codegen2.Yfunc object at 0x7f20be0d8050> ]]
//
//ydefs// str => Yfunc:(BUILTINS.str) [[ <codegen2.Yfunc object at 0x7f20be0d8090> ]]
//
//ydefs// sum => Yfunc:(BUILTINS.sum) [[ <codegen2.Yfunc object at 0x7f20be0cff10> ]]
//
//ydefs// tuple => Yfunc:(BUILTINS.tuple) [[ <codegen2.Yfunc object at 0x7f20be0d8190> ]]
//
//ydefs// type => Yfunc:(BUILTINS.type) [[ <codegen2.Yfunc object at 0x7f20be0d81d0> ]]
//
//
//ymeth// JDict => [ ['clear', 'copy', 'get', 'items', 'iteritems'] ]
//ydefs// JDict => clear => Yfunc:(BUILTINS.JDict::clear) [[ <codegen2.Yfunc object at 0x7f20be0d8290> ]]
//ydefs// JDict => copy => Yfunc:(BUILTINS.JDict::copy) [[ <codegen2.Yfunc object at 0x7f20be0d8310> ]]
//ydefs// JDict => get => Yfunc:(BUILTINS.JDict::get) [[ <codegen2.Yfunc object at 0x7f20be0d8450> ]]
//ydefs// JDict => items => Yfunc:(BUILTINS.JDict::items) [[ <codegen2.Yfunc object at 0x7f20be0d8390> ]]
//ydefs// JDict => iteritems => Yfunc:(BUILTINS.JDict::iteritems) [[ <codegen2.Yfunc object at 0x7f20be0d8410> ]]
//
//ymeth// J_object => [ ['__init__'] ]
//ydefs// J_object => __init__ => Yfunc:(BUILTINS.J_object::__init__) [[ <codegen2.Yfunc object at 0x7f20be0d8510> ]]
//
//


