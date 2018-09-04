package runtime

// cwp: 'BUILTINS'
// path: 'BUILTINS'
// thispkg: 'rye_/BUILTINS'
// modname: 'BUILTINS'
// internal: '../compiler/generated_builtins.tmp'

import "fmt"
import "io"
import "log"
import "os"
import "reflect"
import goruntime "runtime"
import "time"
import "bytes"
import i_bytes "bytes" // <parse.Timport object at 0x7f963ac1ea90>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'bytes'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'bytes', 'pkg': 'bytes', 'line': 1, 'where': 0}
import i_strings "strings" // <parse.Timport object at 0x7f963ba47d90>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'strings'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'strings', 'pkg': 'strings', 'line': 1, 'where': 0}
import i_unicode "unicode" // <parse.Timport object at 0x7f963ba47dd0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'unicode'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'unicode', 'pkg': 'unicode', 'line': 1, 'where': 0}
import i_bufio "bufio" // <parse.Timport object at 0x7f963ba47e10>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'bufio'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'bufio', 'pkg': 'bufio', 'line': 2, 'where': 39}
import i_io "io" // <parse.Timport object at 0x7f963ba47e50>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'io'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'io', 'pkg': 'io', 'line': 2, 'where': 39}
import i_ioutil "io/ioutil" // <parse.Timport object at 0x7f963ba47e90>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'io', 'ioutil'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'ioutil', 'pkg': 'io/ioutil', 'line': 2, 'where': 39}
import i_log "log" // <parse.Timport object at 0x7f963ba47ed0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'log'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'log', 'pkg': 'log', 'line': 2, 'where': 39}
import i_os "os" // <parse.Timport object at 0x7f963ba47f10>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'os'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'os', 'pkg': 'os', 'line': 2, 'where': 39}
import i_time "time" // <parse.Timport object at 0x7f963ba47f50>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'time'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'time', 'pkg': 'time', 'line': 2, 'where': 39}
var _ = i_bufio.ErrNegativeCount // bufio
var _ = i_bytes.ErrTooLarge      // bytes
var _ = i_io.ErrUnexpectedEOF    // io
var _ = i_ioutil.Discard         // io/ioutil
var _ = i_log.SetPrefix          // log
var _ = i_os.Stdout              // os
var _ = i_strings.TrimSpace      // strings
var _ = i_time.UTC               // time
var _ = i_unicode.Zs             // unicode
var _ = log.Printf
var _ = fmt.Sprintf
var _ = io.EOF
var _ = os.Stderr
var _ = reflect.ValueOf
var _ = goruntime.Stack
var _ = time.Sleep
var _ = bytes.Split
var _ = MkInt

func inner_eval_module() M {
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
	//Vimport: ['go', 'log'] log go
	// $ 39 $ 2 $
	// @ 39 @ 2 @
	//Vimport: ['go', 'os'] os go
	// $ 39 $ 2 $
	// @ 39 @ 2 @
	//Vimport: ['go', 'time'] time go
	// $ 39 $ 2 $
	// @ 400 @ 11 @
	// $ 400 $ 11 $
	// @ 430 @ 14 @
	// $ 430 $ 14 $
	// @ 492 @ 18 @
	// $ 492 $ 18 $
	// @ 934 @ 38 @
	// $ 934 $ 38 $
	// @ 991 @ 42 @
	// $ 991 $ 42 $
	// @ 1055 @ 46 @
	// $ 1055 $ 46 $
	// @ 1325 @ 60 @
	// $ 1325 $ 60 $
	// @ 1593 @ 72 @
	// $ 1593 $ 72 $
	// @ 1648 @ 76 @
	// $ 1648 $ 76 $
	// @ 1713 @ 80 @
	// $ 1713 $ 80 $
	// @ 1939 @ 92 @
	// $ 1939 $ 92 $
	// @ 2023 @ 97 @
	// $ 2023 $ 97 $
	// @ 2114 @ 101 @
	// $ 2114 $ 101 $
	// @ 2203 @ 105 @
	// $ 2203 $ 105 $
	// @ 2267 @ 109 @
	// $ 2267 $ 109 $
	// @ 2342 @ 113 @
	// $ 2342 $ 113 $
	// @ 2414 @ 119 @
	// $ 2414 $ 119 $
	// @ 2488 @ 125 @
	// $ 2488 $ 125 $
	// @ 2566 @ 131 @
	// $ 2566 $ 131 $
	// @ 2619 @ 135 @
	// $ 2619 $ 135 $
	// @ 2674 @ 139 @
	// $ 2674 $ 139 $
	// @ 2730 @ 143 @
	// $ 2730 $ 143 $
	// @ 2788 @ 147 @
	// $ 2788 $ 147 $
	// @ 2852 @ 151 @
	// $ 2852 $ 151 $
	// @ 2903 @ 155 @
	// $ 2903 $ 155 $
	// @ 2965 @ 161 @
	// $ 2965 $ 161 $
	// @ 3078 @ 165 @
	// $ 3078 $ 165 $
	// @ 3134 @ 169 @
	// $ 3134 $ 169 $
	// @ 3186 @ 173 @
	// $ 3186 $ 173 $
	// @ 3252 @ 177 @
	// $ 3252 $ 177 $
	// @ 3310 @ 181 @
	// $ 3310 $ 181 $
	// @ 3366 @ 185 @
	// $ 3366 $ 185 $
	// @ 3415 @ 189 @
	// $ 3415 $ 189 $
	// @ 3462 @ 193 @
	// $ 3462 $ 193 $
	// @ 3513 @ 197 @
	// $ 3513 $ 197 $
	// @ 3572 @ 201 @
	// $ 3572 $ 201 $
	// @ 3634 @ 205 @
	// $ 3634 $ 205 $
	// @ 3701 @ 209 @
	// $ 3701 $ 209 $
	// @ 3971 @ 226 @
	// $ 3971 $ 226 $
	// @ 4241 @ 243 @
	// $ 4241 $ 243 $
	// @ 4349 @ 247 @
	// $ 4349 $ 247 $
	// @ 4503 @ 251 @
	// $ 4503 $ 251 $
	// @ 4779 @ 261 @
	// $ 4779 $ 261 $
	// @ 4933 @ 271 @
	// $ 4933 $ 271 $
	// @ 6030 @ 327 @
	// $ 6030 $ 327 $
	// @ 7947 @ 424 @
	// $ 7947 $ 424 $
	// @ 12373 @ 590 @
	// $ 12373 $ 590 $
	// @ 16774 @ 755 @
	// $ 16774 $ 755 $
	// @ 16882 @ 759 @
	// $ 16882 $ 759 $
	// @ 17673 @ 781 @
	// $ 17673 $ 781 $
	// @ 17819 @ 786 @
	// $ 17819 $ 786 $
	// @ 19346 @ 843 @
	// $ 19346 $ 843 $
	// @ 19774 @ 854 @
	// $ 19774 $ 854 $
	// @ 20604 @ 892 @
	// $ 20604 $ 892 $
	// @ 20863 @ 903 @
	// $ 20863 $ 903 $
	// @ 19774 @ 854 @
	// $ 19774 $ 854 $
	return None
}

//(begin tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_Exception(a_x M) M {
	// @ 420 @ 12 @ Exception
	// Vvar: local var x -> 'a_x'
	return /*Yvar.str*/ a_x
	// $ 420 $ 12 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Exception

var specFunc_Exception = CallSpec{Name: "Exception", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_Exception(a0 M) M { return G_1_Exception(a0) }

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func G_0_rye_opts() M {
	// @ 448 @ 15 @ rye_opts
	// { native F
	return MkStr(CompileOptions)
	// } native F
	// $ 448 $ 15 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_opts

var specFunc_rye_opts = CallSpec{Name: "rye_opts", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func fnFunc_rye_opts() M { return G_0_rye_opts() }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_go_deref(a_x M) M {
	// @ 511 @ 19 @ go_deref
	// { native F
	return GoDeref(a_x)
	// } native F
	// $ 511 $ 19 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: go_deref

var specFunc_go_deref = CallSpec{Name: "go_deref", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_go_deref(a0 M) M { return G_1_go_deref(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_rye_what(a_x M) M {
	// @ 953 @ 39 @ rye_what
	// { native F
	return N_rye_what(a_x)
	// } native F
	// $ 953 $ 39 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_what

var specFunc_rye_what = CallSpec{Name: "rye_what", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_rye_what(a0 M) M { return G_1_rye_what(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_callable(a_x M) M {
	// @ 1010 @ 43 @ callable
	// { native F
	return MkBool(JCallable(a_x))
	// } native F
	// $ 1010 $ 43 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: callable

var specFunc_callable = CallSpec{Name: "callable", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_callable(a0 M) M { return G_1_callable(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('obj', None), ('kw', None)]
// typPlus: [None, None]
///////////////////////////////

func G_1V_setattrs(a_obj M, _ M, a_kw M) M {
	var v_detuple_102 M = None
	_ = v_detuple_102
	var v_k M = None
	_ = v_k
	var v_v M = None
	_ = v_v
	// @ 1082 @ 47 @ setattrs
	// Vcall: fn: <parse.Tfield object at 0x7f963aceb510>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var kw -> 'a_kw'

	for_returning__101 := func() M { // around FOR
		var receiver__101 Receiver = JIter( /*General*/ /*invoker*/ F_INVOKE_0_items( /*Yvar.str*/ a_kw))
		enougher__101, canEnough__101 := receiver__101.(StartEnougher)
		if canEnough__101 {
			defer enougher__101.Enough()
			println("STARTING FOR")
			enougher__101.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___101, more___101 := receiver__101.Recv()
			if !more___101 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "detuple_102" in scope @@@@@@
			v_detuple_102 = item___101 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vvar: local var detuple_102 -> 'v_detuple_102'
			len_detuple_102 := JLen( /*Yvar.str*/ v_detuple_102)
			if len_detuple_102 != 2 {
				panic(fmt.Sprintf("Assigning object of length %d to %d variables, in destructuring assignment.", len_detuple_102, 2))
			}
			// Vvar: local var detuple_102 -> 'v_detuple_102'
			// @@@@@@ Creating var "k" in scope @@@@@@
			v_k = /*Yvar.str*/ v_detuple_102.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vvar: local var detuple_102 -> 'v_detuple_102'
			// @@@@@@ Creating var "v" in scope @@@@@@
			v_v = /*Yvar.str*/ v_detuple_102.GetItem( /*Yint.str*/ litI__1) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 1110 @ 48 @ setattrs
			// Vcall: fn: <parse.Tvar object at 0x7f963aceb590>
			// Vcall: args: [<parse.Tvar object at 0x7f963aceb5d0>, <parse.Tvar object at 0x7f963aceb610>, <parse.Tvar object at 0x7f963aceb650>]
			// Vcall: names: ['', '', '']
			// Vcall: star: None
			// Vcall: starstar: None
			// Making Global Yvar from 'setattr'
			// Vvar: local var obj -> 'a_obj'
			// Vvar: local var k -> 'v_k'
			// Vvar: local var v -> 'v_v'
			_ = G_3_setattr( /*nando1*/ /*Yvar.str*/ a_obj /*Yvar.str*/, v_k /*Yvar.str*/, v_v) // Assign void: = type: <type 'str'> repr: 'G_3_setattr(/*nando1*/ /*Yvar.str*/a_obj, /*Yvar.str*/v_k, /*Yvar.str*/v_v) '
			// $ 1110 $ 48 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__101 != MissingM {
		return for_returning__101
	}

	// $ 1082 $ 47 $
	// @ 1131 @ 49 @ setattrs
	// Vvar: local var obj -> 'a_obj'
	return /*Yvar.str*/ a_obj
	// $ 1131 $ 49 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: setattrs

var specFunc_setattrs = CallSpec{Name: "setattrs", Args: []string{"obj"}, Defaults: []M{MissingM}, Star: "", StarStar: "kw"}

type pFunc_setattrs struct{ PNewCallable }

func (o *pFunc_setattrs) Contents() interface{} {
	return G_setattrs
}

func (o pFunc_setattrs) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_1V_setattrs(argv[0], MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_id(a_x M) M {
	// @ 1338 @ 61 @ id
	// { native F

	if a_x.X == nil {
		if len(a_x.S) == 0 {
			panic("Rye does not support id() on int")
		} else {
			panic("Rye does not support id() on str")
		}
	}
	return MkInt(int64(reflect.ValueOf(a_x.X.B()).Pointer()))

	// } native F
	// $ 1338 $ 61 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: id

var specFunc_id = CallSpec{Name: "id", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_id(a0 M) M { return G_1_id(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_hash(a_x M) M {
	// @ 1608 @ 73 @ hash
	// { native F
	return MkInt(JHash(a_x))
	// } native F
	// $ 1608 $ 73 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: hash

var specFunc_hash = CallSpec{Name: "hash", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_hash(a0 M) M { return G_1_hash(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('y', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_cmp(a_x M, a_y M) M {
	// @ 1665 @ 77 @ cmp
	// { native F
	return Mkint(JCompare(a_x, a_y))
	// } native F
	// $ 1665 $ 77 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: cmp

var specFunc_cmp = CallSpec{Name: "cmp", Args: []string{"x", "y"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func fnFunc_cmp(a0 M, a1 M) M { return G_2_cmp(a0, a1) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('name', None), ('dflt', None)]
// typPlus: [None, None, None]
///////////////////////////////

func G_2V_getattr(a_x M, a_name M, a_dflt M, _ M) M {
	var v_n M = None
	_ = v_n
	// @ 1744 @ 81 @ getattr
	// Vcall: fn: <parse.Tvar object at 0x7f963acebb90>
	// Vcall: args: [<parse.Tvar object at 0x7f963acebbd0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var dflt -> 'a_dflt'
	// @@@@@@ Creating var "n" in scope @@@@@@
	v_n = /*Yint.str*/ MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_dflt)))) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yint'>
	// $ 1744 $ 81 $
	// @ 1760 @ 82 @ getattr
	// Vvar: local var n -> 'v_n'
	if /*AsBool*/ JBool( /*Yvar.str*/ v_n) {
		// @ 1770 @ 83 @ getattr

		// BEGIN OUTER EXCEPT try_103
		try_103_try := func() (try_103_z M) {
			defer func() {
				r := recover()
				if r != nil {
					PrintStackFYIUnlessEOFBecauseExcept(r)
					try_103_z = func() M {
						// BEGIN EXCEPT

						// @ 1856 @ 87 @ getattr
						// Vvar: local var dflt -> 'a_dflt'
						return /*Yvar.str*/ a_dflt.GetItem( /*Yint.str*/ litI__0)
						// $ 1856 $ 87 $

						return MissingM
						// END EXCEPT
					}()
					return
				}
			}()

			// BEGIN TRY try_103
			// @ 1781 @ 84 @ getattr
			// { native F
			return a_x.FetchField(JString(a_name))
			// } native F
			// $ 1781 $ 84 $
			// END TRY try_103

			return MissingM
		}()
		if try_103_try != MissingM {
			return try_103_try
		}
		// END OUTER EXCEPT try_103

		// $ 1770 $ 83 $
	} else {
		// @ 1883 @ 89 @ getattr
		// { native F
		return a_x.FetchField(JString(a_name))
		// } native F
		// $ 1883 $ 89 $
	}
	// $ 1760 $ 82 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: getattr

var specFunc_getattr = CallSpec{Name: "getattr", Args: []string{"x", "name"}, Defaults: []M{MissingM, MissingM}, Star: "dflt", StarStar: ""}

type pFunc_getattr struct{ PNewCallable }

func (o *pFunc_getattr) Contents() interface{} {
	return G_getattr
}

func (o pFunc_getattr) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_2V_getattr(argv[0], argv[1], MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('name', None), ('val', None)]
// typPlus: [None, None, None]
///////////////////////////////

func G_3_setattr(a_x M, a_name M, a_val M) M {
	// @ 1968 @ 93 @ setattr
	// { native F
	a_x.StoreField(JString(a_name), a_val)
	// } native F
	// $ 1968 $ 93 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: setattr

var specFunc_setattr = CallSpec{Name: "setattr", Args: []string{"x", "name", "val"}, Defaults: []M{MissingM, MissingM, MissingM}, Star: "", StarStar: ""}

func fnFunc_setattr(a0 M, a1 M, a2 M) M { return G_3_setattr(a0, a1, a2) }

//(tail)
// zip(p.argsPlus, typPlus): [('obj', None), ('cls', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_isinstance(a_obj M, a_cls M) M {
	// @ 2051 @ 98 @ isinstance
	// { native F
	return MkBool(IsSubclass(JPType(a_obj), a_cls))
	// } native F
	// $ 2051 $ 98 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isinstance

var specFunc_isinstance = CallSpec{Name: "isinstance", Args: []string{"obj", "cls"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func fnFunc_isinstance(a0 M, a1 M) M { return G_2_isinstance(a0, a1) }

//(tail)
// zip(p.argsPlus, typPlus): [('subcls', None), ('cls', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_issubclass(a_subcls M, a_cls M) M {
	// @ 2145 @ 102 @ issubclass
	// { native F
	return MkBool(IsSubclass(a_subcls, a_cls))
	// } native F
	// $ 2145 $ 102 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: issubclass

var specFunc_issubclass = CallSpec{Name: "issubclass", Args: []string{"subcls", "cls"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func fnFunc_issubclass(a0 M, a1 M) M { return G_2_issubclass(a0, a1) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_ord(a_x M) M {
	// @ 2217 @ 106 @ ord
	// { native F
	return Mkint(int(JString(a_x)[0]))
	// } native F
	// $ 2217 $ 106 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: ord

var specFunc_ord = CallSpec{Name: "ord", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_ord(a0 M) M { return G_1_ord(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_chr(a_x M) M {
	// @ 2281 @ 110 @ chr
	// { native F
	return MkStr(string([]byte{byte(JInt(a_x))}))
	// } native F
	// $ 2281 $ 110 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: chr

var specFunc_chr = CallSpec{Name: "chr", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_chr(a0 M) M { return G_1_chr(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None), ('init', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_sum(a_vec M, a_init M) M {
	var v_x M = None
	_ = v_x
	var v_z M = None
	_ = v_z
	// @ 2366 @ 114 @ sum
	// Vvar: local var init -> 'a_init'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*Yvar.str*/ a_init // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
	// $ 2366 $ 114 $
	// @ 2377 @ 115 @ sum
	// Vvar: local var vec -> 'a_vec'

	for_returning__104 := func() M { // around FOR
		var receiver__104 Receiver = JIter( /*Yvar.str*/ a_vec)
		enougher__104, canEnough__104 := receiver__104.(StartEnougher)
		if canEnough__104 {
			defer enougher__104.Enough()
			println("STARTING FOR")
			enougher__104.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___104, more___104 := receiver__104.Recv()
			if !more___104 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "x" in scope @@@@@@
			v_x = item___104 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 2395 @ 116 @ sum
			// Vvar: local var z -> 'v_z'
			// Vvar: local var x -> 'v_x'
			v_z = /*DoAdd*/ JAdd( /*Yvar.str*/ v_z /*Yvar.str*/, v_x) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 2395 $ 116 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__104 != MissingM {
		return for_returning__104
	}

	// $ 2377 $ 115 $
	// @ 2404 @ 117 @ sum
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 2404 $ 117 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: sum

var specFunc_sum = CallSpec{Name: "sum", Args: []string{"vec", "init"}, Defaults: []M{MissingM /*Yint.str*/, litI__0}, Star: "", StarStar: ""}

func fnFunc_sum(a0 M, a1 M) M { return G_2_sum(a0, a1) }

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
///////////////////////////////

func G_1_any(a_vec M) M {
	var v_e M = None
	_ = v_e
	// @ 2430 @ 120 @ any
	// Vvar: local var vec -> 'a_vec'

	for_returning__105 := func() M { // around FOR
		var receiver__105 Receiver = JIter( /*Yvar.str*/ a_vec)
		enougher__105, canEnough__105 := receiver__105.(StartEnougher)
		if canEnough__105 {
			defer enougher__105.Enough()
			println("STARTING FOR")
			enougher__105.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___105, more___105 := receiver__105.Recv()
			if !more___105 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "e" in scope @@@@@@
			v_e = item___105 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 2448 @ 121 @ any
			// Vvar: local var e -> 'v_e'
			if /*AsBool*/ JBool( /*Yvar.str*/ v_e) {
				// @ 2460 @ 122 @ any
				return /*Ybool.str*/ True
				// $ 2460 $ 122 $
			}
			// $ 2448 $ 121 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__105 != MissingM {
		return for_returning__105
	}

	// $ 2430 $ 120 $
	// @ 2474 @ 123 @ any
	return /*Ybool.str*/ False
	// $ 2474 $ 123 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: any

var specFunc_any = CallSpec{Name: "any", Args: []string{"vec"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_any(a0 M) M { return G_1_any(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
///////////////////////////////

func G_1_all(a_vec M) M {
	var v_e M = None
	_ = v_e
	// @ 2504 @ 126 @ all
	// Vvar: local var vec -> 'a_vec'

	for_returning__106 := func() M { // around FOR
		var receiver__106 Receiver = JIter( /*Yvar.str*/ a_vec)
		enougher__106, canEnough__106 := receiver__106.(StartEnougher)
		if canEnough__106 {
			defer enougher__106.Enough()
			println("STARTING FOR")
			enougher__106.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___106, more___106 := receiver__106.Recv()
			if !more___106 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "e" in scope @@@@@@
			v_e = item___106 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 2522 @ 127 @ all
			// Vvar: local var e -> 'v_e'
			if /*Vboolop*/ !( /*AsBool*/ JBool( /*Yvar.str*/ v_e)) {
				// @ 2538 @ 128 @ all
				return /*Ybool.str*/ False
				// $ 2538 $ 128 $
			}
			// $ 2522 $ 127 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__106 != MissingM {
		return for_returning__106
	}

	// $ 2504 $ 126 $
	// @ 2553 @ 129 @ all
	return /*Ybool.str*/ True
	// $ 2553 $ 129 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: all

var specFunc_all = CallSpec{Name: "all", Args: []string{"vec"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_all(a0 M) M { return G_1_all(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_len(a_x M) M {
	// @ 2580 @ 132 @ len
	// { native F
	return Mkint(JLen(a_x))
	// } native F
	// $ 2580 $ 132 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: len

var specFunc_len = CallSpec{Name: "len", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_len(a0 M) M { return G_1_len(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_repr(a_x M) M {
	// @ 2634 @ 136 @ repr
	// { native F
	return MkStr(JRepr(a_x))
	// } native F
	// $ 2634 $ 136 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: repr

var specFunc_repr = CallSpec{Name: "repr", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_repr(a0 M) M { return G_1_repr(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_str(a_x M) M {
	// @ 2688 @ 140 @ str
	// { native F
	return MkStr(JString(a_x))
	// } native F
	// $ 2688 $ 140 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: str

var specFunc_str = CallSpec{Name: "str", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_str(a0 M) M { return G_1_str(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_int(a_x M) M {
	// @ 2744 @ 144 @ int
	// { native F
	return MkInt(JForceInt(a_x))
	// } native F
	// $ 2744 $ 144 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: int

var specFunc_int = CallSpec{Name: "int", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_int(a0 M) M { return G_1_int(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_float(a_x M) M {
	// @ 2804 @ 148 @ float
	// { native F
	return MkFloat(JForceFloat(a_x))
	// } native F
	// $ 2804 $ 148 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: float

var specFunc_float = CallSpec{Name: "float", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_float(a0 M) M { return G_1_float(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_range(a_x M) M {
	// @ 2868 @ 152 @ range
	// { native F
	return N_range(a_x)
	// } native F
	// $ 2868 $ 152 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: range

var specFunc_range = CallSpec{Name: "range", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_range(a0 M) M { return G_1_range(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_xrange(a_x M) M {
	var v_i M = None
	_ = v_i

	gen := NewGenerator()

	go func() {
		// Recover & repanic, printing FYI.
		defer func() {
			r := recover()
			if r == "RYE_ENOUGH" {
				return
			}
			if r != nil {
				PrintStackFYIUnlessEOFBecauseExcept(r)
				gen.Raise(Mk(r))
			}
		}()

		// WAIT TO BE STARTED
		Log.Printf("generator waiting for Start: %v ...", gen)
		b := <-gen.Back
		if b != FeedbackStart {
			Log.Panicf("got feedback %d, expected FeedbackStart", b)
		}
		Log.Printf("... generator got Start: %v", gen)

		mustBeNone := func() M {

			// @ 2920 @ 156 @ xrange
			// @@@@@@ Creating var "i" in scope @@@@@@
			v_i = /*Yint.str*/ litI__0 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yint'>
			// $ 2920 $ 156 $
			// @ 2928 @ 157 @ xrange
			for {
				// Vvar: local var i -> 'v_i'
				// Vvar: local var x -> 'a_x'
				if !( /*DoLT*/ JLT( /*Yvar.str*/ v_i /*Yvar.str*/, a_x)) {
					break
				}
				// @ 2945 @ 158 @ xrange
				// Vvar: local var i -> 'v_i'
				gen.Send( /*Yvar.str*/ v_i)
				// $ 2945 $ 158 $
				// @ 2957 @ 159 @ xrange
				// Vvar: local var i -> 'v_i'
				v_i = /*DoAdd*/ JAdd( /*Yvar.str*/ v_i /*Yint.str*/, litI__1) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// $ 2957 $ 159 $
			}
			// $ 2928 $ 157 $
			// bottom out: return None
			return None

			return None
		}()
		gen.Close()
		if mustBeNone != None {
			Log.Panicf("Return Value in Generator must be None.")
		}
	}()
	return MkObj(&gen.PBase)

}

///////////////////////////////
// name: xrange

var specFunc_xrange = CallSpec{Name: "xrange", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_xrange(a0 M) M { return G_1_xrange(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('cmp', None), ('key', None), ('reverse', None)]
// typPlus: [None, None, None, None]
///////////////////////////////

func G_4_sorted(a_x M, a_cmp M, a_key M, a_reverse M) M {
	// @ 3017 @ 162 @ sorted
	// { native F
	return N_sorted(a_x, a_cmp, a_key, a_reverse)
	// } native F
	// $ 3017 $ 162 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: sorted

var specFunc_sorted = CallSpec{Name: "sorted", Args: []string{"x", "cmp", "key", "reverse"}, Defaults: []M{MissingM, None, None /*Ybool.str*/, False}, Star: "", StarStar: ""}

type pFunc_sorted struct{ PNewCallable }

func (o *pFunc_sorted) Contents() interface{} {
	return G_sorted
}
func (o pFunc_sorted) Call4(a0 M, a1 M, a2 M, a3 M) M {
	return G_4_sorted(a0, a1, a2, a3)
}

func (o pFunc_sorted) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_4_sorted(argv[0], argv[1], argv[2], argv[3])
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_list(a_x M) M {
	// @ 3093 @ 166 @ list
	// { native F
	return MkList(JList(a_x))
	// } native F
	// $ 3093 $ 166 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: list

var specFunc_list = CallSpec{Name: "list", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_list(a0 M) M { return G_1_list(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('a', None)]
// typPlus: [None]
///////////////////////////////

func G_1_set(a_a M) M {
	// @ 3153 @ 170 @ set
	// { native F
	return N_set(a_a)
	// } native F
	// $ 3153 $ 170 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: set

var specFunc_set = CallSpec{Name: "set", Args: []string{"a"}, Defaults: []M{None}, Star: "", StarStar: ""}

func fnFunc_set(a0 M) M { return G_1_set(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None), ('kv', None)]
// typPlus: [None, None]
///////////////////////////////

func G_0V_dict(a_vec M, a_kv M) M {
	// @ 3210 @ 174 @ dict
	// { native F
	return N_dict(a_vec, a_kv)
	// } native F
	// $ 3210 $ 174 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: dict

var specFunc_dict = CallSpec{Name: "dict", Args: []string{}, Defaults: []M{}, Star: "vec", StarStar: "kv"}

type pFunc_dict struct{ PNewCallable }

func (o *pFunc_dict) Contents() interface{} {
	return G_dict
}

func (o pFunc_dict) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_0V_dict(MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_tuple(a_x M) M {
	// @ 3268 @ 178 @ tuple
	// { native F
	return MkTuple(JList(a_x))
	// } native F
	// $ 3268 $ 178 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: tuple

var specFunc_tuple = CallSpec{Name: "tuple", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_tuple(a0 M) M { return G_1_tuple(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_bool(a_x M) M {
	// @ 3325 @ 182 @ bool
	// { native F
	return MkBool(JBool(a_x))
	// } native F
	// $ 3325 $ 182 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: bool

var specFunc_bool = CallSpec{Name: "bool", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_bool(a0 M) M { return G_1_bool(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_type(a_x M) M {
	// @ 3381 @ 186 @ type
	// { native F
	return JPType(a_x)
	// } native F
	// $ 3381 $ 186 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: type

var specFunc_type = CallSpec{Name: "type", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_type(a0 M) M { return G_1_type(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_byt(a_x M) M {
	// @ 3429 @ 190 @ byt
	// { native F
	return N_byt(a_x)
	// } native F
	// $ 3429 $ 190 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: byt

var specFunc_byt = CallSpec{Name: "byt", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_byt(a0 M) M { return G_1_byt(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('n', None)]
// typPlus: [None]
///////////////////////////////

func G_1_mkbyt(a_n M) M {
	// @ 3478 @ 194 @ mkbyt
	// { native F
	return N_mkbyt(a_n)
	// } native F
	// $ 3478 $ 194 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: mkbyt

var specFunc_mkbyt = CallSpec{Name: "mkbyt", Args: []string{"n"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_mkbyt(a0 M) M { return G_1_mkbyt(a0) }

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func G_0_rye_stack() M {
	// @ 3532 @ 198 @ rye_stack
	// { native F
	return MkStr(RyeStack())
	// } native F
	// $ 3532 $ 198 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_stack

var specFunc_rye_stack = CallSpec{Name: "rye_stack", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func fnFunc_rye_stack() M { return G_0_rye_stack() }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_rye_pickle(a_x M) M {
	// @ 3593 @ 202 @ rye_pickle
	// { native F
	return MkByt(Pickle(a_x))
	// } native F
	// $ 3593 $ 202 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_pickle

var specFunc_rye_pickle = CallSpec{Name: "rye_pickle", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_rye_pickle(a0 M) M { return G_1_rye_pickle(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func G_1_rye_unpickle(a_x M) M {
	// @ 3657 @ 206 @ rye_unpickle
	// { native F
	return UnPickle(JBytes(a_x))
	// } native F
	// $ 3657 $ 206 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_unpickle

var specFunc_rye_unpickle = CallSpec{Name: "rye_unpickle", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_rye_unpickle(a0 M) M { return G_1_rye_unpickle(a0) }

//(tail)
// zip(p.argsPlus, typPlus): [('args', None)]
// typPlus: [None]
///////////////////////////////

func G_0V_max(a_args M, _ M) M {
	var v_e M = None
	_ = v_e
	var v_v M = None
	_ = v_v
	var v_z M = None
	_ = v_z
	// @ 3719 @ 210 @ max
	// Vcall: fn: <parse.Tvar object at 0x7f963ad03450>
	// Vcall: args: [<parse.Tvar object at 0x7f963ad03490>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var args -> 'a_args'
	if /*YYint.doRelop*/ int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_args))) == int64(0) {
		// @ 3742 @ 211 @ max
		panic(M( /*Ystr.str*/ litS_070c4397445592bd39b479ef9a7724b4))
		// $ 3742 $ 211 $
	}
	// $ 3719 $ 210 $
	// @ 3769 @ 212 @ max
	// Vcall: fn: <parse.Tvar object at 0x7f963ad03690>
	// Vcall: args: [<parse.Tvar object at 0x7f963ad036d0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var args -> 'a_args'
	if /*YYint.doRelop*/ int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_args))) == int64(1) {
		// @ 3792 @ 213 @ max
		// Vvar: local var args -> 'a_args'
		// @@@@@@ Creating var "v" in scope @@@@@@
		v_v = /*Yvar.str*/ a_args.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 3792 $ 213 $
		// @ 3808 @ 214 @ max
		// Vvar: local var v -> 'v_v'
		// @@@@@@ Creating var "z" in scope @@@@@@
		v_z = /*Yvar.str*/ v_v.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 3808 $ 214 $
		// @ 3821 @ 215 @ max
		// Vvar: local var v -> 'v_v'

		for_returning__107 := func() M { // around FOR
			var receiver__107 Receiver = JIter( /*Yvar.str*/ v_v.GetItemSlice( /*Yint.str*/ litI__1, None, None))
			enougher__107, canEnough__107 := receiver__107.(StartEnougher)
			if canEnough__107 {
				defer enougher__107.Enough()
				println("STARTING FOR")
				enougher__107.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___107, more___107 := receiver__107.Recv()
				if !more___107 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "e" in scope @@@@@@
				v_e = item___107 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 3843 @ 216 @ max
				// Vvar: local var e -> 'v_e'
				// Vvar: local var z -> 'v_z'
				if /*DoGT*/ JGT( /*Yvar.str*/ v_e /*Yvar.str*/, v_z) {
					// @ 3861 @ 217 @ max
					// Vvar: local var e -> 'v_e'
					v_z = /*Yvar.str*/ v_e // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
					// $ 3861 $ 217 $
				}
				// $ 3843 $ 216 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__107 != MissingM {
			return for_returning__107
		}

		// $ 3821 $ 215 $
		// @ 3871 @ 218 @ max
		// Vvar: local var z -> 'v_z'
		return /*Yvar.str*/ v_z
		// $ 3871 $ 218 $
	} else {
		// @ 3892 @ 220 @ max
		// Vvar: local var args -> 'a_args'
		v_z = /*Yvar.str*/ a_args.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 3892 $ 220 $
		// @ 3908 @ 221 @ max
		// Vvar: local var args -> 'a_args'

		for_returning__108 := func() M { // around FOR
			var receiver__108 Receiver = JIter( /*Yvar.str*/ a_args.GetItemSlice( /*Yint.str*/ litI__1, None, None))
			enougher__108, canEnough__108 := receiver__108.(StartEnougher)
			if canEnough__108 {
				defer enougher__108.Enough()
				println("STARTING FOR")
				enougher__108.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___108, more___108 := receiver__108.Recv()
				if !more___108 {
					break
				}
				// BEGIN FOR

				v_e = item___108 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 3933 @ 222 @ max
				// Vvar: local var e -> 'v_e'
				// Vvar: local var z -> 'v_z'
				if /*DoGT*/ JGT( /*Yvar.str*/ v_e /*Yvar.str*/, v_z) {
					// @ 3951 @ 223 @ max
					// Vvar: local var e -> 'v_e'
					v_z = /*Yvar.str*/ v_e // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
					// $ 3951 $ 223 $
				}
				// $ 3933 $ 222 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__108 != MissingM {
			return for_returning__108
		}

		// $ 3908 $ 221 $
		// @ 3961 @ 224 @ max
		// Vvar: local var z -> 'v_z'
		return /*Yvar.str*/ v_z
		// $ 3961 $ 224 $
	}
	// $ 3769 $ 212 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: max

var specFunc_max = CallSpec{Name: "max", Args: []string{}, Defaults: []M{}, Star: "args", StarStar: ""}

type pFunc_max struct{ PNewCallable }

func (o *pFunc_max) Contents() interface{} {
	return G_max
}

func (o pFunc_max) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_0V_max(MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('args', None)]
// typPlus: [None]
///////////////////////////////

func G_0V_min(a_args M, _ M) M {
	var v_e M = None
	_ = v_e
	var v_v M = None
	_ = v_v
	var v_z M = None
	_ = v_z
	// @ 3989 @ 227 @ min
	// Vcall: fn: <parse.Tvar object at 0x7f963ac8b510>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac8b550>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var args -> 'a_args'
	if /*YYint.doRelop*/ int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_args))) == int64(0) {
		// @ 4012 @ 228 @ min
		panic(M( /*Ystr.str*/ litS_868beb4fbac148f9a5d1e20101485c96))
		// $ 4012 $ 228 $
	}
	// $ 3989 $ 227 $
	// @ 4039 @ 229 @ min
	// Vcall: fn: <parse.Tvar object at 0x7f963ac8b750>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac8b790>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var args -> 'a_args'
	if /*YYint.doRelop*/ int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_args))) == int64(1) {
		// @ 4062 @ 230 @ min
		// Vvar: local var args -> 'a_args'
		// @@@@@@ Creating var "v" in scope @@@@@@
		v_v = /*Yvar.str*/ a_args.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 4062 $ 230 $
		// @ 4078 @ 231 @ min
		// Vvar: local var v -> 'v_v'
		// @@@@@@ Creating var "z" in scope @@@@@@
		v_z = /*Yvar.str*/ v_v.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 4078 $ 231 $
		// @ 4091 @ 232 @ min
		// Vvar: local var v -> 'v_v'

		for_returning__109 := func() M { // around FOR
			var receiver__109 Receiver = JIter( /*Yvar.str*/ v_v.GetItemSlice( /*Yint.str*/ litI__1, None, None))
			enougher__109, canEnough__109 := receiver__109.(StartEnougher)
			if canEnough__109 {
				defer enougher__109.Enough()
				println("STARTING FOR")
				enougher__109.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___109, more___109 := receiver__109.Recv()
				if !more___109 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "e" in scope @@@@@@
				v_e = item___109 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 4113 @ 233 @ min
				// Vvar: local var e -> 'v_e'
				// Vvar: local var z -> 'v_z'
				if /*DoLT*/ JLT( /*Yvar.str*/ v_e /*Yvar.str*/, v_z) {
					// @ 4131 @ 234 @ min
					// Vvar: local var e -> 'v_e'
					v_z = /*Yvar.str*/ v_e // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
					// $ 4131 $ 234 $
				}
				// $ 4113 $ 233 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__109 != MissingM {
			return for_returning__109
		}

		// $ 4091 $ 232 $
		// @ 4141 @ 235 @ min
		// Vvar: local var z -> 'v_z'
		return /*Yvar.str*/ v_z
		// $ 4141 $ 235 $
	} else {
		// @ 4162 @ 237 @ min
		// Vvar: local var args -> 'a_args'
		v_z = /*Yvar.str*/ a_args.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 4162 $ 237 $
		// @ 4178 @ 238 @ min
		// Vvar: local var args -> 'a_args'

		for_returning__110 := func() M { // around FOR
			var receiver__110 Receiver = JIter( /*Yvar.str*/ a_args.GetItemSlice( /*Yint.str*/ litI__1, None, None))
			enougher__110, canEnough__110 := receiver__110.(StartEnougher)
			if canEnough__110 {
				defer enougher__110.Enough()
				println("STARTING FOR")
				enougher__110.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___110, more___110 := receiver__110.Recv()
				if !more___110 {
					break
				}
				// BEGIN FOR

				v_e = item___110 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 4203 @ 239 @ min
				// Vvar: local var e -> 'v_e'
				// Vvar: local var z -> 'v_z'
				if /*DoLT*/ JLT( /*Yvar.str*/ v_e /*Yvar.str*/, v_z) {
					// @ 4221 @ 240 @ min
					// Vvar: local var e -> 'v_e'
					v_z = /*Yvar.str*/ v_e // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
					// $ 4221 $ 240 $
				}
				// $ 4203 $ 239 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__110 != MissingM {
			return for_returning__110
		}

		// $ 4178 $ 238 $
		// @ 4231 @ 241 @ min
		// Vvar: local var z -> 'v_z'
		return /*Yvar.str*/ v_z
		// $ 4231 $ 241 $
	}
	// $ 4039 $ 229 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: min

var specFunc_min = CallSpec{Name: "min", Args: []string{}, Defaults: []M{}, Star: "args", StarStar: ""}

type pFunc_min struct{ PNewCallable }

func (o *pFunc_min) Contents() interface{} {
	return G_min
}

func (o pFunc_min) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_0V_min(MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('args', None)]
// typPlus: [None]
///////////////////////////////

func G_0V_zip(a_args M, _ M) M {
	var v_a M = None
	_ = v_a
	var v_i M = None
	_ = v_i
	var v_n M = None
	_ = v_n
	// @ 4259 @ 244 @ zip
	// Vcall: fn: <parse.Tvar object at 0x7f963ac91610>
	// Vcall: args: [<parse.Tforexpr object at 0x7f963ac91790>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'min'
	// Vvar: local var args -> 'a_args'

	forexpr__111 := func() M { // around FOR EXPR
		var zz__111 []M
		var receiver__111 Receiver = JIter( /*Yvar.str*/ a_args)
		enougher__111, canEnough__111 := receiver__111.(StartEnougher)
		if canEnough__111 {
			defer enougher__111.Enough()
			enougher__111.Start(8)
		}
		// else case without StartEnougher will be faster.
		for {
			item___111, more___111 := receiver__111.Recv()
			if !more___111 {
				break
			}
			// BEGIN FOR EXPR

			// @@@@@@ Creating var "a" in scope @@@@@@
			v_a = item___111 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vcall: fn: <parse.Tvar object at 0x7f963ac91650>
			// Vcall: args: [<parse.Tvar object at 0x7f963ac91690>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var a -> 'v_a'
			zz__111 = append(zz__111 /*Yint.str*/, MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ v_a)))))

			// END FOR EXPR
		}
		return MkList(zz__111)
	}() // around FOR EXPR
	_ = forexpr__111 // Workaround a bug in nested forexprs.

	// @@@@@@ Creating var "n" in scope @@@@@@
	v_n = CALL_1( /*nando2*/ M( /*Yvar.str*/ G_min), forexpr__111) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 4259 $ 244 $
	// @ 4293 @ 245 @ zip
	// Vcall: fn: <parse.Tvar object at 0x7f963ac91a90>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac91ad0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'range'
	// Vvar: local var n -> 'v_n'

	forexpr__112 := func() M { // around FOR EXPR
		var zz__112 []M
		var receiver__112 Receiver = JIter(G_1_range( /*nando1*/ /*Yvar.str*/ v_n))
		enougher__112, canEnough__112 := receiver__112.(StartEnougher)
		if canEnough__112 {
			defer enougher__112.Enough()
			enougher__112.Start(8)
		}
		// else case without StartEnougher will be faster.
		for {
			item___112, more___112 := receiver__112.Recv()
			if !more___112 {
				break
			}
			// BEGIN FOR EXPR

			// @@@@@@ Creating var "i" in scope @@@@@@
			v_i = item___112 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vcall: fn: <parse.Tvar object at 0x7f963ac91850>
			// Vcall: args: [<parse.Tforexpr object at 0x7f963ac919d0>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Making Global Yvar from 'tuple'
			// Vvar: local var args -> 'a_args'

			forexpr__113 := func() M { // around FOR EXPR
				var zz__113 []M
				var receiver__113 Receiver = JIter( /*Yvar.str*/ a_args)
				enougher__113, canEnough__113 := receiver__113.(StartEnougher)
				if canEnough__113 {
					defer enougher__113.Enough()
					enougher__113.Start(8)
				}
				// else case without StartEnougher will be faster.
				for {
					item___113, more___113 := receiver__113.Recv()
					if !more___113 {
						break
					}
					// BEGIN FOR EXPR

					v_a = item___113 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
					// Vvar: local var a -> 'v_a'
					// Vvar: local var i -> 'v_i'
					zz__113 = append(zz__113 /*Yvar.str*/, v_a.GetItem( /*Yvar.str*/ v_i))

					// END FOR EXPR
				}
				return MkList(zz__113)
			}() // around FOR EXPR
			_ = forexpr__113 // Workaround a bug in nested forexprs.

			zz__112 = append(zz__112, G_1_tuple( /*nando1*/ forexpr__113))

			// END FOR EXPR
		}
		return MkList(zz__112)
	}() // around FOR EXPR
	_ = forexpr__112 // Workaround a bug in nested forexprs.

	return forexpr__112
	// $ 4293 $ 245 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: zip

var specFunc_zip = CallSpec{Name: "zip", Args: []string{}, Defaults: []M{}, Star: "args", StarStar: ""}

type pFunc_zip struct{ PNewCallable }

func (o *pFunc_zip) Contents() interface{} {
	return G_zip
}

func (o pFunc_zip) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_0V_zip(MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('args', None)]
// typPlus: [None]
///////////////////////////////

func G_0V_rye_zip_padding_with_None(a_args M, _ M) M {
	var v_a M = None
	_ = v_a
	var v_i M = None
	_ = v_i
	var v_m M = None
	_ = v_m
	// @ 4389 @ 248 @ rye_zip_padding_with_None
	// Vcall: fn: <parse.Tvar object at 0x7f963ac91c90>
	// Vcall: args: [<parse.Tforexpr object at 0x7f963ac91e10>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'max'
	// Vvar: local var args -> 'a_args'

	forexpr__114 := func() M { // around FOR EXPR
		var zz__114 []M
		var receiver__114 Receiver = JIter( /*Yvar.str*/ a_args)
		enougher__114, canEnough__114 := receiver__114.(StartEnougher)
		if canEnough__114 {
			defer enougher__114.Enough()
			enougher__114.Start(8)
		}
		// else case without StartEnougher will be faster.
		for {
			item___114, more___114 := receiver__114.Recv()
			if !more___114 {
				break
			}
			// BEGIN FOR EXPR

			// @@@@@@ Creating var "a" in scope @@@@@@
			v_a = item___114 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vcall: fn: <parse.Tvar object at 0x7f963ac91cd0>
			// Vcall: args: [<parse.Tvar object at 0x7f963ac91d10>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var a -> 'v_a'
			zz__114 = append(zz__114 /*Yint.str*/, MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ v_a)))))

			// END FOR EXPR
		}
		return MkList(zz__114)
	}() // around FOR EXPR
	_ = forexpr__114 // Workaround a bug in nested forexprs.

	// @@@@@@ Creating var "m" in scope @@@@@@
	v_m = CALL_1( /*nando2*/ M( /*Yvar.str*/ G_max), forexpr__114) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 4389 $ 248 $
	// @ 4423 @ 249 @ rye_zip_padding_with_None
	// Vcall: fn: <parse.Tvar object at 0x7f963ac9c310>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac9c350>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'range'
	// Vvar: local var m -> 'v_m'

	forexpr__115 := func() M { // around FOR EXPR
		var zz__115 []M
		var receiver__115 Receiver = JIter(G_1_range( /*nando1*/ /*Yvar.str*/ v_m))
		enougher__115, canEnough__115 := receiver__115.(StartEnougher)
		if canEnough__115 {
			defer enougher__115.Enough()
			enougher__115.Start(8)
		}
		// else case without StartEnougher will be faster.
		for {
			item___115, more___115 := receiver__115.Recv()
			if !more___115 {
				break
			}
			// BEGIN FOR EXPR

			// @@@@@@ Creating var "i" in scope @@@@@@
			v_i = item___115 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vcall: fn: <parse.Tvar object at 0x7f963ac91ed0>
			// Vcall: args: [<parse.Tforexpr object at 0x7f963ac9c250>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Making Global Yvar from 'tuple'
			// Vvar: local var args -> 'a_args'

			forexpr__116 := func() M { // around FOR EXPR
				var zz__116 []M
				var receiver__116 Receiver = JIter( /*Yvar.str*/ a_args)
				enougher__116, canEnough__116 := receiver__116.(StartEnougher)
				if canEnough__116 {
					defer enougher__116.Enough()
					enougher__116.Start(8)
				}
				// else case without StartEnougher will be faster.
				for {
					item___116, more___116 := receiver__116.Recv()
					if !more___116 {
						break
					}
					// BEGIN FOR EXPR

					v_a = item___116 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
					cond_117 := func(a bool) M {
						if a {
							// Vvar: local var a -> 'v_a'
							// Vvar: local var i -> 'v_i'
							return /*Yvar.str*/ v_a.GetItem( /*Yvar.str*/ v_i)
						}
						return None
					}
					// Vvar: local var i -> 'v_i'
					// Vcall: fn: <parse.Tvar object at 0x7f963ac9c050>
					// Vcall: args: [<parse.Tvar object at 0x7f963ac9c090>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Vvar: local var a -> 'v_a'
					zz__116 = append(zz__116, cond_117(( /*DoLT*/ JLT( /*Yvar.str*/ v_i /*Yint.str*/, MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ v_a))))))))

					// END FOR EXPR
				}
				return MkList(zz__116)
			}() // around FOR EXPR
			_ = forexpr__116 // Workaround a bug in nested forexprs.

			zz__115 = append(zz__115, G_1_tuple( /*nando1*/ forexpr__116))

			// END FOR EXPR
		}
		return MkList(zz__115)
	}() // around FOR EXPR
	_ = forexpr__115 // Workaround a bug in nested forexprs.

	return forexpr__115
	// $ 4423 $ 249 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_zip_padding_with_None

var specFunc_rye_zip_padding_with_None = CallSpec{Name: "rye_zip_padding_with_None", Args: []string{}, Defaults: []M{}, Star: "args", StarStar: ""}

type pFunc_rye_zip_padding_with_None struct{ PNewCallable }

func (o *pFunc_rye_zip_padding_with_None) Contents() interface{} {
	return G_rye_zip_padding_with_None
}

func (o pFunc_rye_zip_padding_with_None) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_0V_rye_zip_padding_with_None(MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('fn', None), ('lists', None)]
// typPlus: [None, None]
///////////////////////////////

func G_1V_map(a_fn M, a_lists M, _ M) M {
	var v_tup M = None
	_ = v_tup
	var v_x M = None
	_ = v_x
	// @ 4526 @ 252 @ map
	// @ 4526 @ 252 @ map
	// Vcall: fn: <parse.Tvar object at 0x7f963ac9c4d0>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac9c510>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var lists -> 'a_lists'
	sw_118 := M( /*Yint.str*/ MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yvar.str*/ a_lists)))))
	_ = sw_118
	// $ 4526 $ 252 $
	switch true {
	// @ 4554 @ 253 @ map
	case /*L979*/ JEQ(sw_118 /*Yint.str*/, litI__0):
		{
			// $ 4554 $ 253 $
			// @ 4563 @ 254 @ map
			panic(M( /*Ystr.str*/ litS_6fe87cae7f4cfc67d4ffb571068ceb39))
			// $ 4563 $ 254 $
		} // end case
		// @ 4605 @ 255 @ map
	case /*L979*/ JEQ(sw_118 /*Yint.str*/, litI__1):
		{
			// $ 4605 $ 255 $
			// @ 4614 @ 256 @ map
			// Vvar: local var lists -> 'a_lists'

			forexpr__119 := func() M { // around FOR EXPR
				var zz__119 []M
				var receiver__119 Receiver = JIter( /*Yvar.str*/ a_lists.GetItem( /*Yint.str*/ litI__0))
				enougher__119, canEnough__119 := receiver__119.(StartEnougher)
				if canEnough__119 {
					defer enougher__119.Enough()
					enougher__119.Start(8)
				}
				// else case without StartEnougher will be faster.
				for {
					item___119, more___119 := receiver__119.Recv()
					if !more___119 {
						break
					}
					// BEGIN FOR EXPR

					// @@@@@@ Creating var "x" in scope @@@@@@
					v_x = item___119 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
					// Vcall: fn: <parse.Tvar object at 0x7f963ac9c6d0>
					// Vcall: args: [<parse.Tvar object at 0x7f963ac9c710>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Vvar: local var fn -> 'a_fn'
					// Vvar: local var x -> 'v_x'
					zz__119 = append(zz__119, CALL_1( /*nando2*/ M( /*Yvar.str*/ a_fn) /*Yvar.str*/, v_x))

					// END FOR EXPR
				}
				return MkList(zz__119)
			}() // around FOR EXPR
			_ = forexpr__119 // Workaround a bug in nested forexprs.

			return forexpr__119
			// $ 4614 $ 256 $
		} // end case
	default:
		{
			// @ 4715 @ 259 @ map
			// Vcall: fn: <parse.Tvar object at 0x7f963ac9ca50>
			// Vcall: args: []
			// Vcall: names: []
			// Vcall: star: <parse.Tvar object at 0x7f963ac9ca90>
			// Vcall: starstar: None
			// Making Global Yvar from 'rye_zip_padding_with_None'
			// Vvar: local var lists -> 'a_lists'

			forexpr__120 := func() M { // around FOR EXPR
				var zz__120 []M
				var receiver__120 Receiver = JIter(M( /*nando3*/ /*Yvar.str*/ G_rye_zip_padding_with_None).X.(ICallV).CallV([]M{}, JList( /*Yvar.str*/ a_lists), []KV{}, nil))
				enougher__120, canEnough__120 := receiver__120.(StartEnougher)
				if canEnough__120 {
					defer enougher__120.Enough()
					enougher__120.Start(8)
				}
				// else case without StartEnougher will be faster.
				for {
					item___120, more___120 := receiver__120.Recv()
					if !more___120 {
						break
					}
					// BEGIN FOR EXPR

					// @@@@@@ Creating var "tup" in scope @@@@@@
					v_tup = item___120 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
					// Vcall: fn: <parse.Tvar object at 0x7f963ac9c950>
					// Vcall: args: []
					// Vcall: names: []
					// Vcall: star: <parse.Tvar object at 0x7f963ac9c990>
					// Vcall: starstar: None
					// Vvar: local var fn -> 'a_fn'
					// Vvar: local var tup -> 'v_tup'
					zz__120 = append(zz__120, M( /*nando3*/ /*Yvar.str*/ a_fn).X.(ICallV).CallV([]M{}, JList( /*Yvar.str*/ v_tup), []KV{}, nil))

					// END FOR EXPR
				}
				return MkList(zz__120)
			}() // around FOR EXPR
			_ = forexpr__120 // Workaround a bug in nested forexprs.

			return forexpr__120
			// $ 4715 $ 259 $
		} // end case
	} // end switch
	// $ 4526 $ 252 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: map

var specFunc_map = CallSpec{Name: "map", Args: []string{"fn"}, Defaults: []M{MissingM}, Star: "lists", StarStar: ""}

type pFunc_map struct{ PNewCallable }

func (o *pFunc_map) Contents() interface{} {
	return G_map
}

func (o pFunc_map) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_1V_map(argv[0], MkObj(&star.PBase), MkObj(&starstar.PBase))
}

//(tail)
// zip(p.argsPlus, typPlus): [('fn', None), ('vec', None), ('init', None)]
// typPlus: [None, None, None]
///////////////////////////////

func G_3_reduce(a_fn M, a_vec M, a_init M) M {
	var v_a M = None
	_ = v_a
	var v_e M = None
	_ = v_e
	// @ 4813 @ 262 @ reduce
	// Vcall: fn: <parse.Tvar object at 0x7f963ac9cd10>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac9cd50>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'list'
	// Vvar: local var vec -> 'a_vec'
	a_vec = G_1_list( /*nando1*/ /*Yvar.str*/ a_vec) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 4813 $ 262 $
	// @ 4831 @ 263 @ reduce
	// Vvar: local var init -> 'a_init'
	if /*Vop returns bool*/ JIs( /*L1148*/ /*Yvar.str*/ a_init, None) {
		// @ 4852 @ 264 @ reduce
		// Vcall: fn: <parse.Tfield object at 0x7f963ac9cf50>
		// Vcall: args: [<parse.Tlit object at 0x7f963ac9cf90>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var vec -> 'a_vec'
		// @@@@@@ Creating var "a" in scope @@@@@@
		v_a = /*General*/ /*invoker*/ F_INVOKE_1_pop( /*Yvar.str*/ a_vec /*Yint.str*/, litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 4852 $ 264 $
	} else {
		// @ 4879 @ 266 @ reduce
		// Vvar: local var init -> 'a_init'
		v_a = /*Yvar.str*/ a_init // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
		// $ 4879 $ 266 $
	}
	// $ 4831 $ 263 $
	// @ 4890 @ 267 @ reduce
	// Vvar: local var vec -> 'a_vec'

	for_returning__121 := func() M { // around FOR
		var receiver__121 Receiver = JIter( /*Yvar.str*/ a_vec)
		enougher__121, canEnough__121 := receiver__121.(StartEnougher)
		if canEnough__121 {
			defer enougher__121.Enough()
			println("STARTING FOR")
			enougher__121.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___121, more___121 := receiver__121.Recv()
			if !more___121 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "e" in scope @@@@@@
			v_e = item___121 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 4908 @ 268 @ reduce
			// Vcall: fn: <parse.Tvar object at 0x7f963aca12d0>
			// Vcall: args: [<parse.Tvar object at 0x7f963aca1310>, <parse.Tvar object at 0x7f963aca1350>]
			// Vcall: names: ['', '']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var fn -> 'a_fn'
			// Vvar: local var a -> 'v_a'
			// Vvar: local var e -> 'v_e'
			v_a = CALL_2( /*nando2*/ M( /*Yvar.str*/ a_fn) /*Yvar.str*/, v_a /*Yvar.str*/, v_e) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 4908 $ 268 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__121 != MissingM {
		return for_returning__121
	}

	// $ 4890 $ 267 $
	// @ 4923 @ 269 @ reduce
	// Vvar: local var a -> 'v_a'
	return /*Yvar.str*/ v_a
	// $ 4923 $ 269 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: reduce

var specFunc_reduce = CallSpec{Name: "reduce", Args: []string{"fn", "vec", "init"}, Defaults: []M{MissingM, MissingM, None}, Star: "", StarStar: ""}

func fnFunc_reduce(a0 M, a1 M, a2 M) M { return G_3_reduce(a0, a1, a2) }

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_append(a_x M) M {
	// @ 4975 @ 273 @ PList.append
	// { native F
	self.PP = append(self.PP, a_x)
	// } native F
	// $ 4975 $ 273 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: append

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_extend(a_x M) M {
	// @ 5044 @ 277 @ PList.extend
	// { native F
	self.PP = append(self.PP, JList(a_x)...)
	// } native F
	// $ 5044 $ 277 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: extend

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_count(a_x M) M {
	var v_e M = None
	_ = v_e
	var v_z M = None
	_ = v_z
	// @ 5122 @ 281 @ PList.count
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*Yint.str*/ litI__0 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yint'>
	// $ 5122 $ 281 $
	// @ 5132 @ 282 @ PList.count

	for_returning__122 := func() M { // around FOR
		var receiver__122 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
		enougher__122, canEnough__122 := receiver__122.(StartEnougher)
		if canEnough__122 {
			defer enougher__122.Enough()
			println("STARTING FOR")
			enougher__122.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___122, more___122 := receiver__122.Recv()
			if !more___122 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "e" in scope @@@@@@
			v_e = item___122 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 5153 @ 283 @ PList.count
			// Vvar: local var e -> 'v_e'
			// Vvar: local var x -> 'a_x'
			if /*DoEQ*/ JEQ( /*Yvar.str*/ v_e /*Yvar.str*/, a_x) {
				// @ 5172 @ 284 @ PList.count
				// Vvar: local var z -> 'v_z'
				v_z = /*DoAdd*/ JAdd( /*Yvar.str*/ v_z /*Yint.str*/, litI__1) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// $ 5172 $ 284 $
			}
			// $ 5153 $ 283 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__122 != MissingM {
		return for_returning__122
	}

	// $ 5132 $ 282 $
	// @ 5183 @ 285 @ PList.count
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 5183 $ 285 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: count

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_index(a_x M) M {
	var v_e M = None
	_ = v_e
	var v_i M = None
	_ = v_i
	// @ 5213 @ 288 @ PList.index
	// @@@@@@ Creating var "i" in scope @@@@@@
	v_i = /*Yint.str*/ litI__0 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yint'>
	// $ 5213 $ 288 $
	// @ 5223 @ 289 @ PList.index

	for_returning__123 := func() M { // around FOR
		var receiver__123 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
		enougher__123, canEnough__123 := receiver__123.(StartEnougher)
		if canEnough__123 {
			defer enougher__123.Enough()
			println("STARTING FOR")
			enougher__123.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___123, more___123 := receiver__123.Recv()
			if !more___123 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "e" in scope @@@@@@
			v_e = item___123 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 5244 @ 290 @ PList.index
			// Vvar: local var e -> 'v_e'
			// Vvar: local var x -> 'a_x'
			if /*DoEQ*/ JEQ( /*Yvar.str*/ v_e /*Yvar.str*/, a_x) {
				// @ 5263 @ 291 @ PList.index
				// Vvar: local var i -> 'v_i'
				return /*Yvar.str*/ v_i
				// $ 5263 $ 291 $
			}
			// $ 5244 $ 290 $
			// @ 5278 @ 292 @ PList.index
			// Vvar: local var i -> 'v_i'
			v_i = /*DoAdd*/ JAdd( /*Yvar.str*/ v_i /*Yint.str*/, litI__1) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 5278 $ 292 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__123 != MissingM {
		return for_returning__123
	}

	// $ 5223 $ 289 $
	// @ 5289 @ 293 @ PList.index
	panic(M( /*Ystr.str*/ litS__ValueError))
	// $ 5289 $ 293 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: index

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_remove(a_x M) M {
	// @ 5330 @ 296 @ PList.remove
	// Vcall: fn: <parse.Tfield object at 0x7f963aca8250>
	// Vcall: args: [<parse.Tvar object at 0x7f963aca8290>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	/*Yself.str*/
	MkObj(&self.PBase).DelItem( /*General*/ /*invoker*/ F_INVOKE_1_index( /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_x))
	// $ 5330 $ 296 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: remove

//(tail)
// zip(p.argsPlus, typPlus): [('i', None), ('x', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *PList) M_2_insert(a_i M, a_x M) M {
	var v_z M = None
	_ = v_z
	// @ 5437 @ 300 @ PList.insert
	// Vvar: local var i -> 'a_i'
	// Vvar: local var x -> 'a_x'
	// Vvar: local var i -> 'a_i'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*DoAdd*/ JAdd( /*DoAdd*/ JAdd( /*Yself.str*/ MkObj(&self.PBase).GetItemSlice(None /*Yvar.str*/, a_i, None), MkListV( /*Yvar.str*/ a_x)) /*Yself.str*/, MkObj(&self.PBase).GetItemSlice( /*Yvar.str*/ a_i, None, None)) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 5437 $ 300 $
	// @ 5471 @ 301 @ PList.insert
	// { native F
	self.PP = JList(v_z)
	// } native F
	// $ 5471 $ 301 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: insert

//(tail)
// zip(p.argsPlus, typPlus): [('i', None)]
// typPlus: [None]
///////////////////////////////

func (self *PList) M_1_pop(a_i M) M {
	var v_x M = None
	_ = v_x
	// @ 5532 @ 305 @ PList.pop
	// Vvar: local var i -> 'a_i'
	// @@@@@@ Creating var "x" in scope @@@@@@
	v_x = /*Yself.str*/ MkObj(&self.PBase).GetItem( /*Yvar.str*/ a_i) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 5532 $ 305 $
	// @ 5548 @ 306 @ PList.pop
	// Vvar: local var i -> 'a_i'
	/*Yself.str*/
	MkObj(&self.PBase).DelItem( /*Yvar.str*/ a_i)
	// $ 5548 $ 306 $
	// @ 5564 @ 307 @ PList.pop
	// Vvar: local var x -> 'v_x'
	return /*Yvar.str*/ v_x
	// $ 5564 $ 307 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: pop

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PList) M_0_reverse() M {
	var v_detuple_125 M = None
	_ = v_detuple_125
	var v_i M = None
	_ = v_i
	var v_n M = None
	_ = v_n
	// @ 5595 @ 310 @ PList.reverse
	// Vcall: fn: <parse.Tvar object at 0x7f963aca8bd0>
	// Vcall: args: [<parse.Tvar object at 0x7f963aca8c10>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// @@@@@@ Creating var "n" in scope @@@@@@
	v_n = /*Yint.str*/ MkInt(int64( /*G.DoLen else*/ int64( /*global DoLen Yint*/ JLen( /*Yself.str*/ MkObj(&self.PBase))))) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yint'>
	// $ 5595 $ 310 $
	// @ 5630 @ 312 @ PList.reverse
	// Vcall: fn: <parse.Tvar object at 0x7f963aca8d50>
	// Vcall: args: [<parse.Top object at 0x7f963aca8e10>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var n -> 'v_n'

	var i_for_range_124 int64
	var n_for_range_124 int64 = /*ForceInt*/ JForceInt(( /*DoDiv*/ JDiv( /*Yvar.str*/ v_n /*Yint.str*/, litI__2)))
	for i_for_range_124 = int64(0); i_for_range_124 < n_for_range_124; i_for_range_124++ {
		var tmp_for_range_124 M = MkInt(i_for_range_124)

		// @@@@@@ Creating var "i" in scope @@@@@@
		v_i = tmp_for_range_124 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// Begin optimized_for_range Block
		// @ 5704 @ 314 @ PList.reverse
		// Vvar: local var n -> 'v_n'
		// Vvar: local var i -> 'v_i'
		// Vvar: local var i -> 'v_i'
		// @@@@@@ Creating var "detuple_125" in scope @@@@@@
		v_detuple_125 = MkTupleV( /*Yself.str*/ MkObj(&self.PBase).GetItem(( /*DoSub*/ JSub(( /*DoSub*/ JSub( /*Yvar.str*/ v_n /*Yvar.str*/, v_i)) /*Yint.str*/, litI__1))) /*Yself.str*/, MkObj(&self.PBase).GetItem( /*Yvar.str*/ v_i)) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// Vvar: local var detuple_125 -> 'v_detuple_125'
		len_detuple_125 := JLen( /*Yvar.str*/ v_detuple_125)
		if len_detuple_125 != 2 {
			panic(fmt.Sprintf("Assigning object of length %d to %d variables, in destructuring assignment.", len_detuple_125, 2))
		}
		// Vvar: local var detuple_125 -> 'v_detuple_125'
		// Vvar: local var i -> 'v_i'
		( /*Yself.str*/ MkObj(&self.PBase)).SetItem( /*Yvar.str*/ v_i /*Yvar.str*/, v_detuple_125.GetItem( /*Yint.str*/ litI__0))
		// Vvar: local var detuple_125 -> 'v_detuple_125'
		// Vvar: local var n -> 'v_n'
		// Vvar: local var i -> 'v_i'
		( /*Yself.str*/ MkObj(&self.PBase)).SetItem(( /*DoSub*/ JSub(( /*DoSub*/ JSub( /*Yvar.str*/ v_n /*Yvar.str*/, v_i)) /*Yint.str*/, litI__1)) /*Yvar.str*/, v_detuple_125.GetItem( /*Yint.str*/ litI__1))
		// $ 5704 $ 314 $
		// End optimized_for_range Block
	}
	// $ 5630 $ 312 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: reverse

//(tail)
// zip(p.argsPlus, typPlus): [('cmp', None), ('key', None), ('reverse', None)]
// typPlus: [None, None, None]
///////////////////////////////

func (self *PList) M_3_sort(a_cmp M, a_key M, a_reverse M) M {
	// @ 5800 @ 317 @ PList.sort
	// { native F
	self.PP = JList(N_sorted( /*inline.*/ MkObj(&self.PBase), a_cmp, a_key, a_reverse))
	// } native F
	// $ 5800 $ 317 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: sort

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PList) M_0_copy() M {
	// @ 5917 @ 321 @ PList.copy
	// { native F

	var zz []M
	for _, e := range self.PP {
		zz = append(zz, e)
	}
	return MkList(zz)

	// } native F
	// $ 5917 $ 321 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: copy

//(tail)

var specMeth_1_PList__append = CallSpec{Name: "PList::append", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_append() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__append}, o.M_1_append})
}

var specMeth_0_PList__copy = CallSpec{Name: "PList::copy", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PList) GET_copy() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PList__copy}, o.M_0_copy})
}

var specMeth_1_PList__count = CallSpec{Name: "PList::count", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_count() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__count}, o.M_1_count})
}

var specMeth_1_PList__extend = CallSpec{Name: "PList::extend", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_extend() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__extend}, o.M_1_extend})
}

var specMeth_1_PList__index = CallSpec{Name: "PList::index", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_index() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__index}, o.M_1_index})
}

var specMeth_2_PList__insert = CallSpec{Name: "PList::insert", Args: []string{"i", "x"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_insert() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PList__insert}, o.M_2_insert})
}

var specMeth_1_PList__pop = CallSpec{Name: "PList::pop", Args: []string{"i"}, Defaults: []M{ /*Yint.str*/ litI___XminusX_1}, Star: "", StarStar: ""}

func (o *PList) GET_pop() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__pop}, o.M_1_pop})
}

var specMeth_1_PList__remove = CallSpec{Name: "PList::remove", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PList) GET_remove() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PList__remove}, o.M_1_remove})
}

var specMeth_0_PList__reverse = CallSpec{Name: "PList::reverse", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PList) GET_reverse() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PList__reverse}, o.M_0_reverse})
}

var specMeth_3_PList__sort = CallSpec{Name: "PList::sort", Args: []string{"cmp", "key", "reverse"}, Defaults: []M{None, None /*Ybool.str*/, False}, Star: "", StarStar: ""}

func (o *PList) GET_sort() M {
	return MForge(&PCall3{PNewCallable{CallSpec: &specMeth_3_PList__sort}, o.M_3_sort})
}

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_clear() M {
	// @ 6070 @ 329 @ PDict.clear
	// { native F
	if 'm' {
		self.mu.Lock()
	}
	self.ppp = make(map[string]M)
	if 'm' {
		self.mu.Unlock()
	}
	// } native F
	// $ 6070 $ 329 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: clear

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_copy() M {
	// @ 6237 @ 339 @ PDict.copy
	// { native F
	z := make(map[string]M)
	if 'm' {
		self.mu.Lock()
	}
	for k, v := range self.ppp {
		z[k] = v
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkDict(z)
	// } native F
	// $ 6237 $ 339 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: copy

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_items() M {
	// @ 6472 @ 351 @ PDict.items
	// { native F
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for k, v := range self.ppp {
		z = append(z, MkTuple([]M{MkStr(k), v}))
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
	// } native F
	// $ 6472 $ 351 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: items

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_iteritems() M {
	// @ 6751 @ 362 @ PDict.iteritems
	// Vcall: fn: <parse.Tfield object at 0x7f963acaeb10>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	return /*General*/ /*invoker*/ F_INVOKE_0_items( /*Yself.str*/ MkObj(&self.PBase))
	// $ 6751 $ 362 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: iteritems

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_keys() M {
	// @ 6786 @ 365 @ PDict.keys
	// { native F
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for k, _ := range self.ppp {
		z = append(z, MkStr(k))
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
	// } native F
	// $ 6786 $ 365 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: keys

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_iterkeys() M {
	// @ 7047 @ 376 @ PDict.iterkeys
	// Vcall: fn: <parse.Tfield object at 0x7f963acaed50>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	return /*General*/ /*invoker*/ F_INVOKE_0_keys( /*Yself.str*/ MkObj(&self.PBase))
	// $ 7047 $ 376 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: iterkeys

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_iter() M {
	// @ 7080 @ 378 @ PDict.iter
	// Vcall: fn: <parse.Tfield object at 0x7f963acaeed0>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	return /*General*/ /*invoker*/ F_INVOKE_0_keys( /*Yself.str*/ MkObj(&self.PBase))
	// $ 7080 $ 378 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: iter

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_values() M {
	// @ 7116 @ 381 @ PDict.values
	// { native F
	z := make([]M, 0, len(self.ppp))
	if 'm' {
		self.mu.Lock()
	}
	for _, v := range self.ppp {
		z = append(z, v)
	}
	if 'm' {
		self.mu.Unlock()
	}
	return MkList(z)
	// } native F
	// $ 7116 $ 381 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: values

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PDict) M_0_itervalues() M {
	// @ 7372 @ 392 @ PDict.itervalues
	// Vcall: fn: <parse.Tfield object at 0x7f963ac95150>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	return /*General*/ /*invoker*/ F_INVOKE_0_values( /*Yself.str*/ MkObj(&self.PBase))
	// $ 7372 $ 392 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: itervalues

//(tail)
// zip(p.argsPlus, typPlus): [('key', None), ('default', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *PDict) M_2_get(a_key M, a_default M) M {
	// @ 7426 @ 395 @ PDict.get
	// { native F
	k := JString(a_key)
	if 'm' {
		self.mu.Lock()
	}
	z, ok := self.ppp[k]
	if 'm' {
		self.mu.Unlock()
	}
	if ok {
		return z
	}
	return a_default
	// } native F
	// $ 7426 $ 395 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: get

//(tail)
// zip(p.argsPlus, typPlus): [('key', None)]
// typPlus: [None]
///////////////////////////////

func (self *PDict) M_1_has_key(a_key M) M {
	// @ 7670 @ 408 @ PDict.has_key
	// Vvar: local var key -> 'a_key'
	return /*Ybool.str*/ MkBool(( /*Vop returns bool*/ JContains( /*L1148*/ /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_key)))
	// $ 7670 $ 408 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: has_key

//(tail)
// zip(p.argsPlus, typPlus): [('key', None), ('default', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *PDict) M_2_setdefault(a_key M, a_default M) M {
	// @ 7733 @ 411 @ PDict.setdefault
	// Vvar: local var key -> 'a_key'
	if /*Vop returns bool*/ JContains( /*L1148*/ /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_key) {
		// @ 7755 @ 412 @ PDict.setdefault
		// Vvar: local var key -> 'a_key'
		return /*Yself.str*/ MkObj(&self.PBase).GetItem( /*Yvar.str*/ a_key)
		// $ 7755 $ 412 $
	} else {
		// @ 7788 @ 414 @ PDict.setdefault
		// Vvar: local var default -> 'a_default'
		// Vvar: local var key -> 'a_key'
		( /*Yself.str*/ MkObj(&self.PBase)).SetItem( /*Yvar.str*/ a_key /*Yvar.str*/, a_default)
		// $ 7788 $ 414 $
		// @ 7814 @ 415 @ PDict.setdefault
		// Vvar: local var default -> 'a_default'
		return /*Yvar.str*/ a_default
		// $ 7814 $ 415 $
	}
	// $ 7733 $ 411 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: setdefault

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PDict) M_1_update(a_x M) M {
	var v_detuple_127 M = None
	_ = v_detuple_127
	var v_k M = None
	_ = v_k
	var v_stuff M = None
	_ = v_stuff
	var v_v M = None
	_ = v_v
	// @ 7880 @ 419 @ PDict.update
	// Vcall: fn: <parse.Tfield object at 0x7f963ac95b50>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tvar object at 0x7f963ac95a90>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac95ad0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'dict'
	// Vvar: local var x -> 'a_x'
	// @@@@@@ Creating var "stuff" in scope @@@@@@
	v_stuff = /*General*/ /*invoker*/ F_INVOKE_0_items(CALL_1( /*nando2*/ M( /*Yvar.str*/ G_dict) /*Yvar.str*/, a_x)) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 7880 $ 419 $
	// @ 7908 @ 420 @ PDict.update
	// Vvar: local var stuff -> 'v_stuff'

	for_returning__126 := func() M { // around FOR
		var receiver__126 Receiver = JIter( /*Yvar.str*/ v_stuff)
		enougher__126, canEnough__126 := receiver__126.(StartEnougher)
		if canEnough__126 {
			defer enougher__126.Enough()
			println("STARTING FOR")
			enougher__126.Start(8)
		} else {
			// log.Println("Not a StartEnougher")
		}
		// else case without StartEnougher will be faster.
		for {
			item___126, more___126 := receiver__126.Recv()
			if !more___126 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "detuple_127" in scope @@@@@@
			v_detuple_127 = item___126 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vvar: local var detuple_127 -> 'v_detuple_127'
			len_detuple_127 := JLen( /*Yvar.str*/ v_detuple_127)
			if len_detuple_127 != 2 {
				panic(fmt.Sprintf("Assigning object of length %d to %d variables, in destructuring assignment.", len_detuple_127, 2))
			}
			// Vvar: local var detuple_127 -> 'v_detuple_127'
			// @@@@@@ Creating var "k" in scope @@@@@@
			v_k = /*Yvar.str*/ v_detuple_127.GetItem( /*Yint.str*/ litI__0) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// Vvar: local var detuple_127 -> 'v_detuple_127'
			// @@@@@@ Creating var "v" in scope @@@@@@
			v_v = /*Yvar.str*/ v_detuple_127.GetItem( /*Yint.str*/ litI__1) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 7933 @ 421 @ PDict.update
			// Vvar: local var v -> 'v_v'
			// Vvar: local var k -> 'v_k'
			( /*Yself.str*/ MkObj(&self.PBase)).SetItem( /*Yvar.str*/ v_k /*Yvar.str*/, v_v)
			// $ 7933 $ 421 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__126 != MissingM {
		return for_returning__126
	}

	// $ 7908 $ 420 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: update

//(tail)

var specMeth_0_PDict__clear = CallSpec{Name: "PDict::clear", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_clear() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__clear}, o.M_0_clear})
}

var specMeth_0_PDict__copy = CallSpec{Name: "PDict::copy", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_copy() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__copy}, o.M_0_copy})
}

var specMeth_2_PDict__get = CallSpec{Name: "PDict::get", Args: []string{"key", "default"}, Defaults: []M{MissingM, None}, Star: "", StarStar: ""}

func (o *PDict) GET_get() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PDict__get}, o.M_2_get})
}

var specMeth_1_PDict__has_key = CallSpec{Name: "PDict::has_key", Args: []string{"key"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PDict) GET_has_key() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PDict__has_key}, o.M_1_has_key})
}

var specMeth_0_PDict__items = CallSpec{Name: "PDict::items", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_items() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__items}, o.M_0_items})
}

var specMeth_0_PDict__iter = CallSpec{Name: "PDict::iter", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_iter() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__iter}, o.M_0_iter})
}

var specMeth_0_PDict__iteritems = CallSpec{Name: "PDict::iteritems", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_iteritems() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__iteritems}, o.M_0_iteritems})
}

var specMeth_0_PDict__iterkeys = CallSpec{Name: "PDict::iterkeys", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_iterkeys() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__iterkeys}, o.M_0_iterkeys})
}

var specMeth_0_PDict__itervalues = CallSpec{Name: "PDict::itervalues", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_itervalues() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__itervalues}, o.M_0_itervalues})
}

var specMeth_0_PDict__keys = CallSpec{Name: "PDict::keys", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_keys() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__keys}, o.M_0_keys})
}

var specMeth_2_PDict__setdefault = CallSpec{Name: "PDict::setdefault", Args: []string{"key", "default"}, Defaults: []M{MissingM, None}, Star: "", StarStar: ""}

func (o *PDict) GET_setdefault() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PDict__setdefault}, o.M_2_setdefault})
}

var specMeth_1_PDict__update = CallSpec{Name: "PDict::update", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PDict) GET_update() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PDict__update}, o.M_1_update})
}

var specMeth_0_PDict__values = CallSpec{Name: "PDict::values", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PDict) GET_values() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PDict__values}, o.M_0_values})
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('n', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *PStr) M_2_split(a_x M, a_n M) M {
	// @ 8070 @ 428 @ PStr.split
	// $ 8070 $ 428 $
	// @ 8159 @ 430 @ PStr.split
	// Vvar: local var x -> 'a_x'
	if /*Vop returns bool*/ JIs( /*L1148*/ /*Yvar.str*/ a_x, None) {
		// @ 8179 @ 431 @ PStr.split
		// { native F

		s := self.S
		n := int(JInt(a_n))
		var v []string
		for n < 0 || len(v) < n {
			if len(s) == 0 {
				break
			}
			i := i_strings.IndexAny(s, " \t\n\r")
			if i >= 0 {
				if i > 0 {
					v = append(v, s[:i])
				}
				s = s[i+1:]
			} else {
				if len(s) > 0 {
					v = append(v, s)
					break
				}
			}
		}
		return MkStrs(v)

		// } native F
		// $ 8179 $ 431 $
	}
	// $ 8159 $ 430 $
	// @ 8561 @ 444 @ PStr.split
	// Vvar: local var n -> 'a_n'
	if /*DoLT*/ JLT( /*Yvar.str*/ a_n /*Yint.str*/, litI__0) {
		// @ 8577 @ 445 @ PStr.split
		// { native F
		return MkStrs(i_strings.Split(self.S, JString(a_x)))
		// } native F
		// $ 8577 $ 445 $
	} else {
		// @ 8664 @ 448 @ PStr.split
		// { native F
		return MkStrs(i_strings.SplitN(self.S, JString(a_x), 1+int(JInt(a_n))))
		// } native F
		// $ 8664 $ 448 $
	}
	// $ 8561 $ 444 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: split

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_join(a_vec M) M {
	// @ 8778 @ 452 @ PStr.join
	// $ 8778 $ 452 $
	// @ 8843 @ 453 @ PStr.join
	// { native F
	ss := make([]string, JLen(a_vec))
	for i, p := range JList(a_vec) {
		ss[i] = JString(p)
	}
	return MkStr(i_strings.Join(ss, self.S))
	// } native F
	// $ 8843 $ 453 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: join

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_lower() M {
	// @ 9042 @ 461 @ PStr.lower
	// $ 9042 $ 461 $
	// @ 9085 @ 462 @ PStr.lower
	// { native F
	return MkStr(i_strings.ToLower(self.S))
	// } native F
	// $ 9085 $ 462 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: lower

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_title() M {
	// @ 9161 @ 466 @ PStr.title
	// $ 9161 $ 466 $
	// @ 9204 @ 467 @ PStr.title
	// { native F
	return MkStr(i_strings.ToTitle(self.S))
	// } native F
	// $ 9204 $ 467 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: title

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_upper() M {
	// @ 9280 @ 471 @ PStr.upper
	// $ 9280 $ 471 $
	// @ 9323 @ 472 @ PStr.upper
	// { native F
	return MkStr(i_strings.ToUpper(self.S))
	// } native F
	// $ 9323 $ 472 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: upper

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_endswith(a_x M) M {
	// @ 9403 @ 476 @ PStr.endswith
	// $ 9403 $ 476 $
	// @ 9438 @ 477 @ PStr.endswith
	// { native F
	return MkBool(i_strings.HasSuffix(self.S, JString(a_x)))
	// } native F
	// $ 9438 $ 477 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: endswith

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_startswith(a_x M) M {
	// @ 9537 @ 481 @ PStr.startswith
	// $ 9537 $ 481 $
	// @ 9574 @ 482 @ PStr.startswith
	// { native F
	return MkBool(i_strings.HasPrefix(self.S, JString(a_x)))
	// } native F
	// $ 9574 $ 482 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: startswith

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_strip(a_x M) M {
	// @ 9678 @ 486 @ PStr.strip
	// $ 9678 $ 486 $
	// @ 9746 @ 487 @ PStr.strip
	// { native F
	return MkStr(i_strings.Trim(self.S, JStr(a_x)))
	// } native F
	// $ 9746 $ 487 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: strip

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_lstrip(a_x M) M {
	// @ 9842 @ 491 @ PStr.lstrip
	// $ 9842 $ 491 $
	// @ 9902 @ 492 @ PStr.lstrip
	// { native F
	return MkStr(i_strings.TrimLeft(self.S, JStr(a_x)))
	// } native F
	// $ 9902 $ 492 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: lstrip

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_rstrip(a_x M) M {
	// @ 10002 @ 496 @ PStr.rstrip
	// $ 10002 $ 496 $
	// @ 10060 @ 497 @ PStr.rstrip
	// { native F
	return MkStr(i_strings.TrimRight(self.S, JStr(a_x)))
	// } native F
	// $ 10060 $ 497 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rstrip

//(tail)
// zip(p.argsPlus, typPlus): [('old', None), ('new', None), ('count', None)]
// typPlus: [None, None, None]
///////////////////////////////

func (self *PStr) M_3_replace(a_old M, a_new M, a_count M) M {
	// @ 10171 @ 501 @ PStr.replace
	// $ 10171 $ 501 $
	// @ 10266 @ 502 @ PStr.replace
	// { native F
	return MkStr(i_strings.Replace(self.S, JString(a_old), JString(a_new), int(JInt(a_count))))
	// } native F
	// $ 10266 $ 502 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: replace

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_find(a_x M) M {
	// @ 10394 @ 506 @ PStr.find
	// $ 10394 $ 506 $
	// @ 10474 @ 507 @ PStr.find
	// { native F
	return Mkint(i_strings.Index(self.S, JString(a_x)))
	// } native F
	// $ 10474 $ 507 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: find

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_rfind(a_x M) M {
	// @ 10563 @ 511 @ PStr.rfind
	// $ 10563 $ 511 $
	// @ 10642 @ 512 @ PStr.rfind
	// { native F
	return Mkint(i_strings.LastIndex(self.S, JString(a_x)))
	// } native F
	// $ 10642 $ 512 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rfind

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_index(a_x M) M {
	var v_z M = None
	_ = v_z
	// @ 10735 @ 516 @ PStr.index
	// $ 10735 $ 516 $
	// @ 10818 @ 517 @ PStr.index
	// Vcall: fn: <parse.Tfield object at 0x7f963acc0a90>
	// Vcall: args: [<parse.Tvar object at 0x7f963acc0ad0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*General*/ /*invoker*/ F_INVOKE_1_find( /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_x) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 10818 $ 517 $
	// @ 10839 @ 518 @ PStr.index
	// Vvar: local var z -> 'v_z'
	if /*DoLT*/ JLT( /*Yvar.str*/ v_z /*Yint.str*/, litI__0) {
		// @ 10855 @ 519 @ PStr.index
		panic(M( /*Ystr.str*/ litS__ValueError))
		// $ 10855 $ 519 $
	}
	// $ 10839 $ 518 $
	// @ 10878 @ 520 @ PStr.index
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 10878 $ 520 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: index

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PStr) M_1_rindex(a_x M) M {
	var v_z M = None
	_ = v_z
	// @ 10909 @ 523 @ PStr.rindex
	// $ 10909 $ 523 $
	// @ 10991 @ 524 @ PStr.rindex
	// Vcall: fn: <parse.Tfield object at 0x7f963acc0f90>
	// Vcall: args: [<parse.Tvar object at 0x7f963acc0fd0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*General*/ /*invoker*/ F_INVOKE_1_rfind( /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_x) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 10991 $ 524 $
	// @ 11013 @ 525 @ PStr.rindex
	// Vvar: local var z -> 'v_z'
	if /*DoLT*/ JLT( /*Yvar.str*/ v_z /*Yint.str*/, litI__0) {
		// @ 11029 @ 526 @ PStr.rindex
		panic(M( /*Ystr.str*/ litS__ValueError))
		// $ 11029 $ 526 $
	}
	// $ 11013 $ 525 $
	// @ 11052 @ 527 @ PStr.rindex
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 11052 $ 527 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rindex

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_isalpha() M {
	var v_c M = None
	_ = v_c
	// @ 11083 @ 530 @ PStr.isalpha
	// $ 11083 $ 530 $
	// @ 11128 @ 531 @ PStr.isalpha
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 11143 @ 532 @ PStr.isalpha

		for_returning__128 := func() M { // around FOR
			var receiver__128 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__128, canEnough__128 := receiver__128.(StartEnougher)
			if canEnough__128 {
				defer enougher__128.Enough()
				println("STARTING FOR")
				enougher__128.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___128, more___128 := receiver__128.Recv()
				if !more___128 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___128 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 11166 @ 533 @ PStr.isalpha
				// Vcall: fn: <parse.Tfield object at 0x7f963acc6550>
				// Vcall: args: [<parse.Tcall object at 0x7f963acc6610>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963acc6590>
				// Vcall: args: [<parse.Tvar object at 0x7f963acc65d0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_129_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsLetter TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_129_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_129_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLetter(opt_go_call_129_t_0) // OptimizedGoCall OK
				_ = opt_go_call_129_r_0

				if /*Vboolop*/ !(opt_go_call_129_r_0) {
					// @ 11209 @ 534 @ PStr.isalpha
					return /*Ybool.str*/ False
					// $ 11209 $ 534 $
				}
				// $ 11166 $ 533 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__128 != MissingM {
			return for_returning__128
		}

		// $ 11143 $ 532 $
		// @ 11228 @ 535 @ PStr.isalpha
		return /*Ybool.str*/ True
		// $ 11228 $ 535 $
	} else {
		// @ 11256 @ 537 @ PStr.isalpha
		return /*Ybool.str*/ False
		// $ 11256 $ 537 $
	}
	// $ 11128 $ 531 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isalpha

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_isdigit() M {
	var v_c M = None
	_ = v_c
	// @ 11291 @ 540 @ PStr.isdigit
	// $ 11291 $ 540 $
	// @ 11335 @ 541 @ PStr.isdigit
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 11350 @ 542 @ PStr.isdigit

		for_returning__130 := func() M { // around FOR
			var receiver__130 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__130, canEnough__130 := receiver__130.(StartEnougher)
			if canEnough__130 {
				defer enougher__130.Enough()
				println("STARTING FOR")
				enougher__130.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___130, more___130 := receiver__130.Recv()
				if !more___130 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___130 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 11373 @ 543 @ PStr.isdigit
				// Vcall: fn: <parse.Tfield object at 0x7f963acc6c50>
				// Vcall: args: [<parse.Tcall object at 0x7f963acc6d10>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963acc6c90>
				// Vcall: args: [<parse.Tvar object at 0x7f963acc6cd0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_131_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsDigit TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_131_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_131_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsDigit(opt_go_call_131_t_0) // OptimizedGoCall OK
				_ = opt_go_call_131_r_0

				if /*Vboolop*/ !(opt_go_call_131_r_0) {
					// @ 11415 @ 544 @ PStr.isdigit
					return /*Ybool.str*/ False
					// $ 11415 $ 544 $
				}
				// $ 11373 $ 543 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__130 != MissingM {
			return for_returning__130
		}

		// $ 11350 $ 542 $
		// @ 11434 @ 545 @ PStr.isdigit
		return /*Ybool.str*/ True
		// $ 11434 $ 545 $
	} else {
		// @ 11462 @ 547 @ PStr.isdigit
		return /*Ybool.str*/ False
		// $ 11462 $ 547 $
	}
	// $ 11335 $ 541 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isdigit

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_isalnum() M {
	var v_c M = None
	_ = v_c
	// @ 11497 @ 550 @ PStr.isalnum
	// $ 11497 $ 550 $
	// @ 11552 @ 551 @ PStr.isalnum
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 11567 @ 552 @ PStr.isalnum

		for_returning__132 := func() M { // around FOR
			var receiver__132 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__132, canEnough__132 := receiver__132.(StartEnougher)
			if canEnough__132 {
				defer enougher__132.Enough()
				println("STARTING FOR")
				enougher__132.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___132, more___132 := receiver__132.Recv()
				if !more___132 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___132 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 11590 @ 553 @ PStr.isalnum
				// Vcall: fn: <parse.Tfield object at 0x7f963ac4f390>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac4f450>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac4f3d0>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac4f410>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_134_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsDigit TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_134_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_134_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsDigit(opt_go_call_134_t_0) // OptimizedGoCall OK
				_ = opt_go_call_134_r_0

				var andand_133 M = /*Ybool.str*/ MkBool(( /*Vboolop*/ !(opt_go_call_134_r_0)))
				if JBool(andand_133) {
					// Vcall: fn: <parse.Tfield object at 0x7f963ac4f550>
					// Vcall: args: [<parse.Tcall object at 0x7f963ac4f610>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Vcall: fn: <parse.Tvar object at 0x7f963ac4f590>
					// Vcall: args: [<parse.Tvar object at 0x7f963ac4f5d0>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Making Global Yvar from 'ord'
					// Vvar: local var c -> 'v_c'
					var opt_go_call_135_r_0 bool /*result*/
					// ATTEMPT OptimizedGoCall: i_unicode.IsLetter TAKES ['int32'] RETURNS ['bool']
					var opt_go_call_135_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
					opt_go_call_135_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLetter(opt_go_call_135_t_0) // OptimizedGoCall OK
					_ = opt_go_call_135_r_0

					andand_133 = /*Ybool.str*/ MkBool(( /*Vboolop*/ !(opt_go_call_135_r_0)))
				}
				if /*AsBool*/ JBool(andand_133) {
					// @ 11669 @ 554 @ PStr.isalnum
					return /*Ybool.str*/ False
					// $ 11669 $ 554 $
				}
				// $ 11590 $ 553 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__132 != MissingM {
			return for_returning__132
		}

		// $ 11567 $ 552 $
		// @ 11688 @ 555 @ PStr.isalnum
		return /*Ybool.str*/ True
		// $ 11688 $ 555 $
	} else {
		// @ 11716 @ 557 @ PStr.isalnum
		return /*Ybool.str*/ False
		// $ 11716 $ 557 $
	}
	// $ 11552 $ 551 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isalnum

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_islower() M {
	var v_c M = None
	_ = v_c
	// @ 11751 @ 560 @ PStr.islower
	// $ 11751 $ 560 $
	// @ 11807 @ 561 @ PStr.islower
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 11822 @ 562 @ PStr.islower

		for_returning__136 := func() M { // around FOR
			var receiver__136 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__136, canEnough__136 := receiver__136.(StartEnougher)
			if canEnough__136 {
				defer enougher__136.Enough()
				println("STARTING FOR")
				enougher__136.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___136, more___136 := receiver__136.Recv()
				if !more___136 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___136 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 11845 @ 563 @ PStr.islower
				// Vcall: fn: <parse.Tfield object at 0x7f963ac4fc90>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac4fd50>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac4fcd0>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac4fd10>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_137_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsLower TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_137_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_137_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLower(opt_go_call_137_t_0) // OptimizedGoCall OK
				_ = opt_go_call_137_r_0

				if /*Vboolop*/ !(opt_go_call_137_r_0) {
					// @ 11887 @ 564 @ PStr.islower
					return /*Ybool.str*/ False
					// $ 11887 $ 564 $
				}
				// $ 11845 $ 563 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__136 != MissingM {
			return for_returning__136
		}

		// $ 11822 $ 562 $
		// @ 11906 @ 565 @ PStr.islower
		return /*Ybool.str*/ True
		// $ 11906 $ 565 $
	} else {
		// @ 11934 @ 567 @ PStr.islower
		return /*Ybool.str*/ False
		// $ 11934 $ 567 $
	}
	// $ 11807 $ 561 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: islower

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_isupper() M {
	var v_c M = None
	_ = v_c
	// @ 11969 @ 570 @ PStr.isupper
	// $ 11969 $ 570 $
	// @ 12025 @ 571 @ PStr.isupper
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 12040 @ 572 @ PStr.isupper

		for_returning__138 := func() M { // around FOR
			var receiver__138 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__138, canEnough__138 := receiver__138.(StartEnougher)
			if canEnough__138 {
				defer enougher__138.Enough()
				println("STARTING FOR")
				enougher__138.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___138, more___138 := receiver__138.Recv()
				if !more___138 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___138 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 12063 @ 573 @ PStr.isupper
				// Vcall: fn: <parse.Tfield object at 0x7f963ac553d0>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac55490>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac55410>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac55450>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_139_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsUpper TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_139_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_139_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsUpper(opt_go_call_139_t_0) // OptimizedGoCall OK
				_ = opt_go_call_139_r_0

				if /*Vboolop*/ !(opt_go_call_139_r_0) {
					// @ 12105 @ 574 @ PStr.isupper
					return /*Ybool.str*/ False
					// $ 12105 $ 574 $
				}
				// $ 12063 $ 573 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__138 != MissingM {
			return for_returning__138
		}

		// $ 12040 $ 572 $
		// @ 12124 @ 575 @ PStr.isupper
		return /*Ybool.str*/ True
		// $ 12124 $ 575 $
	} else {
		// @ 12152 @ 577 @ PStr.isupper
		return /*Ybool.str*/ False
		// $ 12152 $ 577 $
	}
	// $ 12025 $ 571 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isupper

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PStr) M_0_isspace() M {
	var v_c M = None
	_ = v_c
	// @ 12187 @ 580 @ PStr.isspace
	// $ 12187 $ 580 $
	// @ 12231 @ 581 @ PStr.isspace
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 12246 @ 582 @ PStr.isspace

		for_returning__140 := func() M { // around FOR
			var receiver__140 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__140, canEnough__140 := receiver__140.(StartEnougher)
			if canEnough__140 {
				defer enougher__140.Enough()
				println("STARTING FOR")
				enougher__140.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___140, more___140 := receiver__140.Recv()
				if !more___140 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___140 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 12269 @ 583 @ PStr.isspace
				// Vcall: fn: <parse.Tfield object at 0x7f963ac55ad0>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac55b90>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac55b10>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac55b50>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_141_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsSpace TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_141_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_141_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsSpace(opt_go_call_141_t_0) // OptimizedGoCall OK
				_ = opt_go_call_141_r_0

				if /*Vboolop*/ !(opt_go_call_141_r_0) {
					// @ 12311 @ 584 @ PStr.isspace
					return /*Ybool.str*/ False
					// $ 12311 $ 584 $
				}
				// $ 12269 $ 583 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__140 != MissingM {
			return for_returning__140
		}

		// $ 12246 $ 582 $
		// @ 12330 @ 585 @ PStr.isspace
		return /*Ybool.str*/ True
		// $ 12330 $ 585 $
	} else {
		// @ 12358 @ 587 @ PStr.isspace
		return /*Ybool.str*/ False
		// $ 12358 $ 587 $
	}
	// $ 12231 $ 581 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isspace

//(tail)

var specMeth_1_PStr__endswith = CallSpec{Name: "PStr::endswith", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_endswith() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__endswith}, o.M_1_endswith})
}

var specMeth_1_PStr__find = CallSpec{Name: "PStr::find", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_find() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__find}, o.M_1_find})
}

var specMeth_1_PStr__index = CallSpec{Name: "PStr::index", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_index() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__index}, o.M_1_index})
}

var specMeth_0_PStr__isalnum = CallSpec{Name: "PStr::isalnum", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_isalnum() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__isalnum}, o.M_0_isalnum})
}

var specMeth_0_PStr__isalpha = CallSpec{Name: "PStr::isalpha", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_isalpha() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__isalpha}, o.M_0_isalpha})
}

var specMeth_0_PStr__isdigit = CallSpec{Name: "PStr::isdigit", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_isdigit() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__isdigit}, o.M_0_isdigit})
}

var specMeth_0_PStr__islower = CallSpec{Name: "PStr::islower", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_islower() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__islower}, o.M_0_islower})
}

var specMeth_0_PStr__isspace = CallSpec{Name: "PStr::isspace", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_isspace() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__isspace}, o.M_0_isspace})
}

var specMeth_0_PStr__isupper = CallSpec{Name: "PStr::isupper", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_isupper() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__isupper}, o.M_0_isupper})
}

var specMeth_1_PStr__join = CallSpec{Name: "PStr::join", Args: []string{"vec"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_join() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__join}, o.M_1_join})
}

var specMeth_0_PStr__lower = CallSpec{Name: "PStr::lower", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_lower() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__lower}, o.M_0_lower})
}

var specMeth_1_PStr__lstrip = CallSpec{Name: "PStr::lstrip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PStr) GET_lstrip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__lstrip}, o.M_1_lstrip})
}

var specMeth_3_PStr__replace = CallSpec{Name: "PStr::replace", Args: []string{"old", "new", "count"}, Defaults: []M{MissingM, MissingM /*Yint.str*/, litI___XminusX_1}, Star: "", StarStar: ""}

func (o *PStr) GET_replace() M {
	return MForge(&PCall3{PNewCallable{CallSpec: &specMeth_3_PStr__replace}, o.M_3_replace})
}

var specMeth_1_PStr__rfind = CallSpec{Name: "PStr::rfind", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_rfind() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__rfind}, o.M_1_rfind})
}

var specMeth_1_PStr__rindex = CallSpec{Name: "PStr::rindex", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_rindex() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__rindex}, o.M_1_rindex})
}

var specMeth_1_PStr__rstrip = CallSpec{Name: "PStr::rstrip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PStr) GET_rstrip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__rstrip}, o.M_1_rstrip})
}

var specMeth_2_PStr__split = CallSpec{Name: "PStr::split", Args: []string{"x", "n"}, Defaults: []M{None /*Yint.str*/, litI___XminusX_1}, Star: "", StarStar: ""}

func (o *PStr) GET_split() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PStr__split}, o.M_2_split})
}

var specMeth_1_PStr__startswith = CallSpec{Name: "PStr::startswith", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PStr) GET_startswith() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__startswith}, o.M_1_startswith})
}

var specMeth_1_PStr__strip = CallSpec{Name: "PStr::strip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PStr) GET_strip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PStr__strip}, o.M_1_strip})
}

var specMeth_0_PStr__title = CallSpec{Name: "PStr::title", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_title() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__title}, o.M_0_title})
}

var specMeth_0_PStr__upper = CallSpec{Name: "PStr::upper", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PStr) GET_upper() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PStr__upper}, o.M_0_upper})
}

//(tail)
// zip(p.argsPlus, typPlus): [('x', None), ('n', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *PByt) M_2_split(a_x M, a_n M) M {
	// @ 12496 @ 594 @ PByt.split
	// $ 12496 $ 594 $
	// @ 12585 @ 596 @ PByt.split
	// Vvar: local var x -> 'a_x'
	if /*Vop returns bool*/ JIs( /*L1148*/ /*Yvar.str*/ a_x, None) {
		// @ 12605 @ 597 @ PByt.split
		// { native F

		s := self.YY
		n := int(JInt(a_n))
		var v [][]byte
		for n < 0 || len(v) < n {
			if len(s) == 0 {
				break
			}
			i := i_bytes.IndexAny(s, " \t\n\r")
			if i >= 0 {
				if i > 0 {
					v = append(v, s[:i])
				}
				s = s[i+1:]
			} else {
				if len(s) > 0 {
					v = append(v, s)
					break
				}
			}
		}
		return MkByts(v)

		// } native F
		// $ 12605 $ 597 $
	}
	// $ 12585 $ 596 $
	// @ 12986 @ 610 @ PByt.split
	// Vvar: local var n -> 'a_n'
	if /*DoLT*/ JLT( /*Yvar.str*/ a_n /*Yint.str*/, litI__0) {
		// @ 13002 @ 611 @ PByt.split
		// { native F
		return MkByts(i_bytes.Split(self.YY, JBytes(a_x)))
		// } native F
		// $ 13002 $ 611 $
	} else {
		// @ 13087 @ 614 @ PByt.split
		// { native F
		return MkByts(i_bytes.SplitN(self.YY, JBytes(a_x), 1+int(JInt(a_n))))
		// } native F
		// $ 13087 $ 614 $
	}
	// $ 12986 $ 610 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: split

//(tail)
// zip(p.argsPlus, typPlus): [('vec', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_join(a_vec M) M {
	// @ 13199 @ 618 @ PByt.join
	// $ 13199 $ 618 $
	// @ 13264 @ 619 @ PByt.join
	// { native F
	ss := make([][]byte, JLen(a_vec))
	for i, p := range JList(a_vec) {
		ss[i] = JBytes(p)
	}
	return MkByt(i_bytes.Join(ss, self.YY))
	// } native F
	// $ 13264 $ 619 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: join

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_lower() M {
	// @ 13461 @ 627 @ PByt.lower
	// $ 13461 $ 627 $
	// @ 13504 @ 628 @ PByt.lower
	// { native F
	return MkByt(i_bytes.ToLower(self.YY))
	// } native F
	// $ 13504 $ 628 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: lower

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_title() M {
	// @ 13579 @ 632 @ PByt.title
	// $ 13579 $ 632 $
	// @ 13622 @ 633 @ PByt.title
	// { native F
	return MkByt(i_bytes.ToTitle(self.YY))
	// } native F
	// $ 13622 $ 633 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: title

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_upper() M {
	// @ 13697 @ 637 @ PByt.upper
	// $ 13697 $ 637 $
	// @ 13740 @ 638 @ PByt.upper
	// { native F
	return MkByt(i_bytes.ToUpper(self.YY))
	// } native F
	// $ 13740 $ 638 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: upper

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_endswith(a_x M) M {
	// @ 13819 @ 642 @ PByt.endswith
	// $ 13819 $ 642 $
	// @ 13854 @ 643 @ PByt.endswith
	// { native F
	return MkBool(i_bytes.HasSuffix(self.YY, JBytes(a_x)))
	// } native F
	// $ 13854 $ 643 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: endswith

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_startswith(a_x M) M {
	// @ 13951 @ 647 @ PByt.startswith
	// $ 13951 $ 647 $
	// @ 13988 @ 648 @ PByt.startswith
	// { native F
	return MkBool(i_bytes.HasPrefix(self.YY, JBytes(a_x)))
	// } native F
	// $ 13988 $ 648 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: startswith

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_strip(a_x M) M {
	// @ 14090 @ 652 @ PByt.strip
	// $ 14090 $ 652 $
	// @ 14158 @ 653 @ PByt.strip
	// { native F
	return MkByt(i_bytes.Trim(self.YY, JStr(a_x)))
	// } native F
	// $ 14158 $ 653 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: strip

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_lstrip(a_x M) M {
	// @ 14253 @ 657 @ PByt.lstrip
	// $ 14253 $ 657 $
	// @ 14313 @ 658 @ PByt.lstrip
	// { native F
	return MkByt(i_bytes.TrimLeft(self.YY, JStr(a_x)))
	// } native F
	// $ 14313 $ 658 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: lstrip

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_rstrip(a_x M) M {
	// @ 14412 @ 662 @ PByt.rstrip
	// $ 14412 $ 662 $
	// @ 14470 @ 663 @ PByt.rstrip
	// { native F
	return MkByt(i_bytes.TrimRight(self.YY, JStr(a_x)))
	// } native F
	// $ 14470 $ 663 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rstrip

//(tail)
// zip(p.argsPlus, typPlus): [('old', None), ('new', None), ('count', None)]
// typPlus: [None, None, None]
///////////////////////////////

func (self *PByt) M_3_replace(a_old M, a_new M, a_count M) M {
	// @ 14580 @ 667 @ PByt.replace
	// $ 14580 $ 667 $
	// @ 14675 @ 668 @ PByt.replace
	// { native F
	return MkByt(i_bytes.Replace(self.YY, JBytes(a_old), JBytes(a_new), int(JInt(a_count))))
	// } native F
	// $ 14675 $ 668 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: replace

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_find(a_x M) M {
	// @ 14800 @ 672 @ PByt.find
	// $ 14800 $ 672 $
	// @ 14880 @ 673 @ PByt.find
	// { native F
	return Mkint(i_bytes.Index(self.YY, JBytes(a_x)))
	// } native F
	// $ 14880 $ 673 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: find

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_rfind(a_x M) M {
	// @ 14967 @ 677 @ PByt.rfind
	// $ 14967 $ 677 $
	// @ 15046 @ 678 @ PByt.rfind
	// { native F
	return Mkint(i_bytes.LastIndex(self.YY, JBytes(a_x)))
	// } native F
	// $ 15046 $ 678 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rfind

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_index(a_x M) M {
	var v_z M = None
	_ = v_z
	// @ 15137 @ 682 @ PByt.index
	// $ 15137 $ 682 $
	// @ 15220 @ 683 @ PByt.index
	// Vcall: fn: <parse.Tfield object at 0x7f963ac63b10>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac63b50>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*General*/ /*invoker*/ F_INVOKE_1_find( /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_x) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 15220 $ 683 $
	// @ 15241 @ 684 @ PByt.index
	// Vvar: local var z -> 'v_z'
	if /*DoLT*/ JLT( /*Yvar.str*/ v_z /*Yint.str*/, litI__0) {
		// @ 15257 @ 685 @ PByt.index
		panic(M( /*Ystr.str*/ litS__ValueError))
		// $ 15257 $ 685 $
	}
	// $ 15241 $ 684 $
	// @ 15280 @ 686 @ PByt.index
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 15280 $ 686 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: index

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *PByt) M_1_rindex(a_x M) M {
	var v_z M = None
	_ = v_z
	// @ 15311 @ 689 @ PByt.rindex
	// $ 15311 $ 689 $
	// @ 15393 @ 690 @ PByt.rindex
	// Vcall: fn: <parse.Tfield object at 0x7f963ac6a050>
	// Vcall: args: [<parse.Tvar object at 0x7f963ac6a090>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	// @@@@@@ Creating var "z" in scope @@@@@@
	v_z = /*General*/ /*invoker*/ F_INVOKE_1_rfind( /*Yself.str*/ MkObj(&self.PBase) /*Yvar.str*/, a_x) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 15393 $ 690 $
	// @ 15415 @ 691 @ PByt.rindex
	// Vvar: local var z -> 'v_z'
	if /*DoLT*/ JLT( /*Yvar.str*/ v_z /*Yint.str*/, litI__0) {
		// @ 15431 @ 692 @ PByt.rindex
		panic(M( /*Ystr.str*/ litS__ValueError))
		// $ 15431 $ 692 $
	}
	// $ 15415 $ 691 $
	// @ 15454 @ 693 @ PByt.rindex
	// Vvar: local var z -> 'v_z'
	return /*Yvar.str*/ v_z
	// $ 15454 $ 693 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rindex

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_isalpha() M {
	var v_c M = None
	_ = v_c
	// @ 15485 @ 696 @ PByt.isalpha
	// $ 15485 $ 696 $
	// @ 15530 @ 697 @ PByt.isalpha
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 15545 @ 698 @ PByt.isalpha

		for_returning__142 := func() M { // around FOR
			var receiver__142 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__142, canEnough__142 := receiver__142.(StartEnougher)
			if canEnough__142 {
				defer enougher__142.Enough()
				println("STARTING FOR")
				enougher__142.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___142, more___142 := receiver__142.Recv()
				if !more___142 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___142 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 15568 @ 699 @ PByt.isalpha
				// Vcall: fn: <parse.Tfield object at 0x7f963ac6a5d0>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac6a690>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac6a610>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac6a650>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_143_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsLetter TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_143_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_143_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLetter(opt_go_call_143_t_0) // OptimizedGoCall OK
				_ = opt_go_call_143_r_0

				if /*Vboolop*/ !(opt_go_call_143_r_0) {
					// @ 15611 @ 700 @ PByt.isalpha
					return /*Ybool.str*/ False
					// $ 15611 $ 700 $
				}
				// $ 15568 $ 699 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__142 != MissingM {
			return for_returning__142
		}

		// $ 15545 $ 698 $
		// @ 15630 @ 701 @ PByt.isalpha
		return /*Ybool.str*/ True
		// $ 15630 $ 701 $
	} else {
		// @ 15658 @ 703 @ PByt.isalpha
		return /*Ybool.str*/ False
		// $ 15658 $ 703 $
	}
	// $ 15530 $ 697 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isalpha

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_isdigit() M {
	var v_c M = None
	_ = v_c
	// @ 15693 @ 706 @ PByt.isdigit
	// $ 15693 $ 706 $
	// @ 15737 @ 707 @ PByt.isdigit
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 15752 @ 708 @ PByt.isdigit

		for_returning__144 := func() M { // around FOR
			var receiver__144 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__144, canEnough__144 := receiver__144.(StartEnougher)
			if canEnough__144 {
				defer enougher__144.Enough()
				println("STARTING FOR")
				enougher__144.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___144, more___144 := receiver__144.Recv()
				if !more___144 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___144 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 15775 @ 709 @ PByt.isdigit
				// Vcall: fn: <parse.Tfield object at 0x7f963ac6acd0>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac6ad90>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac6ad10>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac6ad50>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_145_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsDigit TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_145_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_145_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsDigit(opt_go_call_145_t_0) // OptimizedGoCall OK
				_ = opt_go_call_145_r_0

				if /*Vboolop*/ !(opt_go_call_145_r_0) {
					// @ 15817 @ 710 @ PByt.isdigit
					return /*Ybool.str*/ False
					// $ 15817 $ 710 $
				}
				// $ 15775 $ 709 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__144 != MissingM {
			return for_returning__144
		}

		// $ 15752 $ 708 $
		// @ 15836 @ 711 @ PByt.isdigit
		return /*Ybool.str*/ True
		// $ 15836 $ 711 $
	} else {
		// @ 15864 @ 713 @ PByt.isdigit
		return /*Ybool.str*/ False
		// $ 15864 $ 713 $
	}
	// $ 15737 $ 707 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isdigit

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_isalnum() M {
	var v_c M = None
	_ = v_c
	// @ 15899 @ 716 @ PByt.isalnum
	// $ 15899 $ 716 $
	// @ 15954 @ 717 @ PByt.isalnum
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 15969 @ 718 @ PByt.isalnum

		for_returning__146 := func() M { // around FOR
			var receiver__146 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__146, canEnough__146 := receiver__146.(StartEnougher)
			if canEnough__146 {
				defer enougher__146.Enough()
				println("STARTING FOR")
				enougher__146.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___146, more___146 := receiver__146.Recv()
				if !more___146 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___146 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 15992 @ 719 @ PByt.isalnum
				// Vcall: fn: <parse.Tfield object at 0x7f963ac70410>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac704d0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac70450>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac70490>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_148_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsDigit TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_148_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_148_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsDigit(opt_go_call_148_t_0) // OptimizedGoCall OK
				_ = opt_go_call_148_r_0

				var andand_147 M = /*Ybool.str*/ MkBool(( /*Vboolop*/ !(opt_go_call_148_r_0)))
				if JBool(andand_147) {
					// Vcall: fn: <parse.Tfield object at 0x7f963ac705d0>
					// Vcall: args: [<parse.Tcall object at 0x7f963ac70690>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Vcall: fn: <parse.Tvar object at 0x7f963ac70610>
					// Vcall: args: [<parse.Tvar object at 0x7f963ac70650>]
					// Vcall: names: ['']
					// Vcall: star: None
					// Vcall: starstar: None
					// Making Global Yvar from 'ord'
					// Vvar: local var c -> 'v_c'
					var opt_go_call_149_r_0 bool /*result*/
					// ATTEMPT OptimizedGoCall: i_unicode.IsLetter TAKES ['int32'] RETURNS ['bool']
					var opt_go_call_149_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
					opt_go_call_149_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLetter(opt_go_call_149_t_0) // OptimizedGoCall OK
					_ = opt_go_call_149_r_0

					andand_147 = /*Ybool.str*/ MkBool(( /*Vboolop*/ !(opt_go_call_149_r_0)))
				}
				if /*AsBool*/ JBool(andand_147) {
					// @ 16071 @ 720 @ PByt.isalnum
					return /*Ybool.str*/ False
					// $ 16071 $ 720 $
				}
				// $ 15992 $ 719 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__146 != MissingM {
			return for_returning__146
		}

		// $ 15969 $ 718 $
		// @ 16090 @ 721 @ PByt.isalnum
		return /*Ybool.str*/ True
		// $ 16090 $ 721 $
	} else {
		// @ 16118 @ 723 @ PByt.isalnum
		return /*Ybool.str*/ False
		// $ 16118 $ 723 $
	}
	// $ 15954 $ 717 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isalnum

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_islower() M {
	var v_c M = None
	_ = v_c
	// @ 16153 @ 726 @ PByt.islower
	// $ 16153 $ 726 $
	// @ 16209 @ 727 @ PByt.islower
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 16224 @ 728 @ PByt.islower

		for_returning__150 := func() M { // around FOR
			var receiver__150 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__150, canEnough__150 := receiver__150.(StartEnougher)
			if canEnough__150 {
				defer enougher__150.Enough()
				println("STARTING FOR")
				enougher__150.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___150, more___150 := receiver__150.Recv()
				if !more___150 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___150 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 16247 @ 729 @ PByt.islower
				// Vcall: fn: <parse.Tfield object at 0x7f963ac70d10>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac70dd0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac70d50>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac70d90>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_151_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsLower TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_151_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_151_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsLower(opt_go_call_151_t_0) // OptimizedGoCall OK
				_ = opt_go_call_151_r_0

				if /*Vboolop*/ !(opt_go_call_151_r_0) {
					// @ 16289 @ 730 @ PByt.islower
					return /*Ybool.str*/ False
					// $ 16289 $ 730 $
				}
				// $ 16247 $ 729 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__150 != MissingM {
			return for_returning__150
		}

		// $ 16224 $ 728 $
		// @ 16308 @ 731 @ PByt.islower
		return /*Ybool.str*/ True
		// $ 16308 $ 731 $
	} else {
		// @ 16336 @ 733 @ PByt.islower
		return /*Ybool.str*/ False
		// $ 16336 $ 733 $
	}
	// $ 16209 $ 727 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: islower

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_isupper() M {
	var v_c M = None
	_ = v_c
	// @ 16371 @ 736 @ PByt.isupper
	// $ 16371 $ 736 $
	// @ 16427 @ 737 @ PByt.isupper
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 16442 @ 738 @ PByt.isupper

		for_returning__152 := func() M { // around FOR
			var receiver__152 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__152, canEnough__152 := receiver__152.(StartEnougher)
			if canEnough__152 {
				defer enougher__152.Enough()
				println("STARTING FOR")
				enougher__152.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___152, more___152 := receiver__152.Recv()
				if !more___152 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___152 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 16465 @ 739 @ PByt.isupper
				// Vcall: fn: <parse.Tfield object at 0x7f963ac75450>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac75510>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac75490>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac754d0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_153_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsUpper TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_153_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_153_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsUpper(opt_go_call_153_t_0) // OptimizedGoCall OK
				_ = opt_go_call_153_r_0

				if /*Vboolop*/ !(opt_go_call_153_r_0) {
					// @ 16507 @ 740 @ PByt.isupper
					return /*Ybool.str*/ False
					// $ 16507 $ 740 $
				}
				// $ 16465 $ 739 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__152 != MissingM {
			return for_returning__152
		}

		// $ 16442 $ 738 $
		// @ 16526 @ 741 @ PByt.isupper
		return /*Ybool.str*/ True
		// $ 16526 $ 741 $
	} else {
		// @ 16554 @ 743 @ PByt.isupper
		return /*Ybool.str*/ False
		// $ 16554 $ 743 $
	}
	// $ 16427 $ 737 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isupper

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *PByt) M_0_isspace() M {
	var v_c M = None
	_ = v_c
	// @ 16589 @ 746 @ PByt.isspace
	// $ 16589 $ 746 $
	// @ 16633 @ 747 @ PByt.isspace
	if /*AsBool*/ JBool( /*Yself.str*/ MkObj(&self.PBase)) {
		// @ 16648 @ 748 @ PByt.isspace

		for_returning__154 := func() M { // around FOR
			var receiver__154 Receiver = JIter( /*Yself.str*/ MkObj(&self.PBase))
			enougher__154, canEnough__154 := receiver__154.(StartEnougher)
			if canEnough__154 {
				defer enougher__154.Enough()
				println("STARTING FOR")
				enougher__154.Start(8)
			} else {
				// log.Println("Not a StartEnougher")
			}
			// else case without StartEnougher will be faster.
			for {
				item___154, more___154 := receiver__154.Recv()
				if !more___154 {
					break
				}
				// BEGIN FOR

				// @@@@@@ Creating var "c" in scope @@@@@@
				v_c = item___154 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// @ 16671 @ 749 @ PByt.isspace
				// Vcall: fn: <parse.Tfield object at 0x7f963ac75b50>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac75c10>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac75b90>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac75bd0>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'ord'
				// Vvar: local var c -> 'v_c'
				var opt_go_call_155_r_0 bool /*result*/
				// ATTEMPT OptimizedGoCall: i_unicode.IsSpace TAKES ['int32'] RETURNS ['bool']
				var opt_go_call_155_t_0 int32 = int32( /*AsInt*/ JInt(G_1_ord( /*nando1*/ /*Yvar.str*/ v_c)))
				opt_go_call_155_r_0 = /*OptimizedGoCall OK*/ i_unicode.IsSpace(opt_go_call_155_t_0) // OptimizedGoCall OK
				_ = opt_go_call_155_r_0

				if /*Vboolop*/ !(opt_go_call_155_r_0) {
					// @ 16713 @ 750 @ PByt.isspace
					return /*Ybool.str*/ False
					// $ 16713 $ 750 $
				}
				// $ 16671 $ 749 $

				// END FOR
			}
			return MissingM
		}() // around FOR
		if for_returning__154 != MissingM {
			return for_returning__154
		}

		// $ 16648 $ 748 $
		// @ 16732 @ 751 @ PByt.isspace
		return /*Ybool.str*/ True
		// $ 16732 $ 751 $
	} else {
		// @ 16760 @ 753 @ PByt.isspace
		return /*Ybool.str*/ False
		// $ 16760 $ 753 $
	}
	// $ 16633 $ 747 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: isspace

//(tail)

var specMeth_1_PByt__endswith = CallSpec{Name: "PByt::endswith", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_endswith() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__endswith}, o.M_1_endswith})
}

var specMeth_1_PByt__find = CallSpec{Name: "PByt::find", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_find() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__find}, o.M_1_find})
}

var specMeth_1_PByt__index = CallSpec{Name: "PByt::index", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_index() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__index}, o.M_1_index})
}

var specMeth_0_PByt__isalnum = CallSpec{Name: "PByt::isalnum", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_isalnum() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__isalnum}, o.M_0_isalnum})
}

var specMeth_0_PByt__isalpha = CallSpec{Name: "PByt::isalpha", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_isalpha() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__isalpha}, o.M_0_isalpha})
}

var specMeth_0_PByt__isdigit = CallSpec{Name: "PByt::isdigit", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_isdigit() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__isdigit}, o.M_0_isdigit})
}

var specMeth_0_PByt__islower = CallSpec{Name: "PByt::islower", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_islower() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__islower}, o.M_0_islower})
}

var specMeth_0_PByt__isspace = CallSpec{Name: "PByt::isspace", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_isspace() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__isspace}, o.M_0_isspace})
}

var specMeth_0_PByt__isupper = CallSpec{Name: "PByt::isupper", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_isupper() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__isupper}, o.M_0_isupper})
}

var specMeth_1_PByt__join = CallSpec{Name: "PByt::join", Args: []string{"vec"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_join() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__join}, o.M_1_join})
}

var specMeth_0_PByt__lower = CallSpec{Name: "PByt::lower", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_lower() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__lower}, o.M_0_lower})
}

var specMeth_1_PByt__lstrip = CallSpec{Name: "PByt::lstrip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PByt) GET_lstrip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__lstrip}, o.M_1_lstrip})
}

var specMeth_3_PByt__replace = CallSpec{Name: "PByt::replace", Args: []string{"old", "new", "count"}, Defaults: []M{MissingM, MissingM /*Yint.str*/, litI___XminusX_1}, Star: "", StarStar: ""}

func (o *PByt) GET_replace() M {
	return MForge(&PCall3{PNewCallable{CallSpec: &specMeth_3_PByt__replace}, o.M_3_replace})
}

var specMeth_1_PByt__rfind = CallSpec{Name: "PByt::rfind", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_rfind() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__rfind}, o.M_1_rfind})
}

var specMeth_1_PByt__rindex = CallSpec{Name: "PByt::rindex", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_rindex() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__rindex}, o.M_1_rindex})
}

var specMeth_1_PByt__rstrip = CallSpec{Name: "PByt::rstrip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PByt) GET_rstrip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__rstrip}, o.M_1_rstrip})
}

var specMeth_2_PByt__split = CallSpec{Name: "PByt::split", Args: []string{"x", "n"}, Defaults: []M{None /*Yint.str*/, litI___XminusX_1}, Star: "", StarStar: ""}

func (o *PByt) GET_split() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PByt__split}, o.M_2_split})
}

var specMeth_1_PByt__startswith = CallSpec{Name: "PByt::startswith", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *PByt) GET_startswith() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__startswith}, o.M_1_startswith})
}

var specMeth_1_PByt__strip = CallSpec{Name: "PByt::strip", Args: []string{"x"}, Defaults: []M{ /*Ystr.str*/ litS_d6221c0c57599569fa8b67dc1d259878}, Star: "", StarStar: ""}

func (o *PByt) GET_strip() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PByt__strip}, o.M_1_strip})
}

var specMeth_0_PByt__title = CallSpec{Name: "PByt::title", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_title() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__title}, o.M_0_title})
}

var specMeth_0_PByt__upper = CallSpec{Name: "PByt::upper", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *PByt) GET_upper() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PByt__upper}, o.M_0_upper})
}

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func G_0_object() M {
	// @ 16790 @ 756 @ object
	// $ 16790 $ 756 $
	// @ 16844 @ 757 @ object
	// { native F
	return MForge(&C_object{})
	// } native F
	// $ 16844 $ 757 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: object

var specFunc_object = CallSpec{Name: "object", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func fnFunc_object() M { return G_0_object() }

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_object) M_0___init__() M {
	// bottom out: return None
	return None

}

///////////////////////////////
// name: __init__

//(tail)
// zip(p.argsPlus, typPlus): [('field', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_object) M_1___getattr__(a_field M) M {
	// @ 17101 @ 765 @ C_object.__getattr__
	// $ 17101 $ 765 $
	// @ 17152 @ 766 @ C_object.__getattr__
	// { native F
	return FetchFieldByNameForObject(reflect.ValueOf(self.Self), JString(a_field))
	// } native F
	// $ 17152 $ 766 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: __getattr__

//(tail)
// zip(p.argsPlus, typPlus): [('field', None), ('value', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *C_object) M_2___setattr__(a_field M, a_value M) M {
	// @ 17284 @ 769 @ C_object.__setattr__
	// $ 17284 $ 769 $
	// @ 17332 @ 770 @ C_object.__setattr__
	// { native F
	StoreFieldByNameForObject(reflect.ValueOf(self.Self), JString(a_field), a_value)
	// } native F
	// $ 17332 $ 770 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: __setattr__

//(tail)

var specMeth_1_C_object____getattr__ = CallSpec{Name: "C_object::__getattr__", Args: []string{"field"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *C_object) GET___getattr__() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_C_object____getattr__}, o.M_1___getattr__})
}

var specMeth_0_C_object____init__ = CallSpec{Name: "C_object::__init__", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_object) GET___init__() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_C_object____init__}, o.M_0___init__})
}

var specMeth_2_C_object____setattr__ = CallSpec{Name: "C_object::__setattr__", Args: []string{"field", "value"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func (o *C_object) GET___setattr__() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_C_object____setattr__}, o.M_2___setattr__})
}

//(tail)
// zip(p.argsPlus, typPlus): [('size', None)]
// typPlus: [None]
///////////////////////////////

func G_1_rye_chan(a_size M) M {
	// @ 17695 @ 782 @ rye_chan
	// $ 17695 $ 782 $
	// @ 17764 @ 783 @ rye_chan
	// { native F
	return make_rye_chan(int(JInt(a_size)))
	// } native F
	// $ 17764 $ 783 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: rye_chan

var specFunc_rye_chan = CallSpec{Name: "rye_chan", Args: []string{"size"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func fnFunc_rye_chan(a0 M) M { return G_1_rye_chan(a0) }

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_channel) M_0_Warm() M {
	// @ 17940 @ 790 @ C_channel.Warm
	// { native F

	Log.Printf("... C_channel::Warm waiting: %v ...", self)
	b := <-self.Back
	if b != FeedbackStart {
		Log.Panicf("C_channel::Warm got feedback %d, expected FeedbackStart", b)
	}
	Log.Printf("... C_channel::Warm got Start: %v", self)

	// } native F
	// $ 17940 $ 790 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Warm

//(tail)
// zip(p.argsPlus, typPlus): [('size', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_channel) M_1_Start(a_size M) M {
	// @ 18264 @ 801 @ C_channel.Start
	// { native F
	self.Start(int(JInt(a_size)))
	// } native F
	// $ 18264 $ 801 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Start

//(tail)
// zip(p.argsPlus, typPlus): [('e', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_channel) M_1_Raise(a_e M) M {
	// @ 18331 @ 805 @ C_channel.Raise
	// $ 18331 $ 805 $
	// @ 18426 @ 806 @ C_channel.Raise
	// { native F
	self.Raise(a_e)
	// } native F
	// $ 18426 $ 806 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Raise

//(tail)
// zip(p.argsPlus, typPlus): [('a', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_channel) M_1_Send(a_a M) M {
	// @ 18478 @ 810 @ C_channel.Send
	// $ 18478 $ 810 $
	// @ 18514 @ 811 @ C_channel.Send
	// { native F
	self.Send(a_a)
	// } native F
	// $ 18514 $ 811 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Send

//(tail)
// zip(p.argsPlus, typPlus): [('timeout', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_channel) M_1_Recv(a_timeout M) M {
	// @ 18576 @ 815 @ C_channel.Recv
	// $ 18576 $ 815 $
	// @ 18686 @ 816 @ C_channel.Recv
	// { native F

	if a_timeout == None {
		z, ok := self.Recv()
		return MkTuple([]M{z, MkBool(ok)})
	} else {
		z, ok := self.RecvWithTimeout(i_time.Duration((JFloat(a_timeout) * 1000000000.0)) * i_time.Nanosecond)
		return MkTuple([]M{z, MkBool(ok)})
	}

	// } native F
	// $ 18686 $ 816 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Recv

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_channel) M_0_Wait() M {
	// @ 19011 @ 828 @ C_channel.Wait
	// $ 19011 $ 828 $
	// @ 19125 @ 829 @ C_channel.Wait
	// { native F

	z, ok := self.Recv()
	if !ok {
		panic("EOF")
	}
	return z

	// } native F
	// $ 19125 $ 829 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Wait

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_channel) M_0_Close() M {
	// @ 19243 @ 839 @ C_channel.Close
	// $ 19243 $ 839 $
	// @ 19316 @ 840 @ C_channel.Close
	// { native F
	self.Close()
	// } native F
	// $ 19316 $ 840 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Close

//(tail)

var specMeth_0_C_channel__Close = CallSpec{Name: "C_channel::Close", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_channel) GET_Close() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_C_channel__Close}, o.M_0_Close})
}

var specMeth_1_C_channel__Raise = CallSpec{Name: "C_channel::Raise", Args: []string{"e"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *C_channel) GET_Raise() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_C_channel__Raise}, o.M_1_Raise})
}

var specMeth_1_C_channel__Recv = CallSpec{Name: "C_channel::Recv", Args: []string{"timeout"}, Defaults: []M{None}, Star: "", StarStar: ""}

func (o *C_channel) GET_Recv() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_C_channel__Recv}, o.M_1_Recv})
}

var specMeth_1_C_channel__Send = CallSpec{Name: "C_channel::Send", Args: []string{"a"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *C_channel) GET_Send() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_C_channel__Send}, o.M_1_Send})
}

var specMeth_1_C_channel__Start = CallSpec{Name: "C_channel::Start", Args: []string{"size"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *C_channel) GET_Start() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_C_channel__Start}, o.M_1_Start})
}

var specMeth_0_C_channel__Wait = CallSpec{Name: "C_channel::Wait", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_channel) GET_Wait() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_C_channel__Wait}, o.M_0_Wait})
}

var specMeth_0_C_channel__Warm = CallSpec{Name: "C_channel::Warm", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_channel) GET_Warm() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_C_channel__Warm}, o.M_0_Warm})
}

//(tail)
// zip(p.argsPlus, typPlus): [('filename', None), ('mode', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_open(a_filename M, a_mode M) M {
	// @ 19378 @ 844 @ open
	// $ 19378 $ 844 $
	// @ 19472 @ 845 @ open
	// Vvar: local var mode -> 'a_mode'
	if /*DoEQ*/ JEQ( /*Yvar.str*/ a_mode /*Ystr.str*/, litS__r) {
		// @ 19492 @ 846 @ open
		// Vcall: fn: <parse.Tvar object at 0x7f963ac81510>
		// Vcall: args: [<parse.Tcall object at 0x7f963ac81610>, <parse.Traw object at 0x7f963ac81650>]
		// Vcall: names: ['', '']
		// Vcall: star: None
		// Vcall: starstar: None
		// Making Global Yvar from 'PYE_FileDesc'
		// Vcall: fn: <parse.Tfield object at 0x7f963ac81590>
		// Vcall: args: [<parse.Tvar object at 0x7f963ac815d0>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var filename -> 'a_filename'
		var opt_go_call_156_r_0 interface{} /*result*/
		var opt_go_call_156_r_1 error       /*result*/
		// ATTEMPT OptimizedGoCall: i_os.Open TAKES ['string'] RETURNS ['*os.File', 'error']
		var opt_go_call_156_t_0 string = string( /*AsStr*/ JStr( /*Yvar.str*/ a_filename))
		opt_go_call_156_r_0, opt_go_call_156_r_1 = /*OptimizedGoCall OK*/ i_os.Open(opt_go_call_156_t_0) // OptimizedGoCall OK
		_ = opt_go_call_156_r_0
		_ = opt_go_call_156_r_1
		if opt_go_call_156_r_1 != nil {
			panic(opt_go_call_156_r_1) /*Check magic_error*/
		}

		return G_2_PYE_FileDesc( /*nando1*/ AdaptToRye(reflect.ValueOf(opt_go_call_156_r_0)) /*Ybool.str*/, False)
		// $ 19492 $ 846 $
	} else {
		// Vvar: local var mode -> 'a_mode'
		if /*DoEQ*/ JEQ( /*Yvar.str*/ a_mode /*Ystr.str*/, litS__w) {
			// @ 19562 @ 848 @ open
			// Vcall: fn: <parse.Tvar object at 0x7f963ac81810>
			// Vcall: args: [<parse.Tcall object at 0x7f963ac81910>, <parse.Traw object at 0x7f963ac81950>]
			// Vcall: names: ['', '']
			// Vcall: star: None
			// Vcall: starstar: None
			// Making Global Yvar from 'PYE_FileDesc'
			// Vcall: fn: <parse.Tfield object at 0x7f963ac81890>
			// Vcall: args: [<parse.Tvar object at 0x7f963ac818d0>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var filename -> 'a_filename'
			var opt_go_call_157_r_0 interface{} /*result*/
			var opt_go_call_157_r_1 error       /*result*/
			// ATTEMPT OptimizedGoCall: i_os.Create TAKES ['string'] RETURNS ['*os.File', 'error']
			var opt_go_call_157_t_0 string = string( /*AsStr*/ JStr( /*Yvar.str*/ a_filename))
			opt_go_call_157_r_0, opt_go_call_157_r_1 = /*OptimizedGoCall OK*/ i_os.Create(opt_go_call_157_t_0) // OptimizedGoCall OK
			_ = opt_go_call_157_r_0
			_ = opt_go_call_157_r_1
			if opt_go_call_157_r_1 != nil {
				panic(opt_go_call_157_r_1) /*Check magic_error*/
			}

			return G_2_PYE_FileDesc( /*nando1*/ AdaptToRye(reflect.ValueOf(opt_go_call_157_r_0)) /*Ybool.str*/, True)
			// $ 19562 $ 848 $
		} else {
			// Vvar: local var mode -> 'a_mode'
			if /*DoEQ*/ JEQ( /*Yvar.str*/ a_mode /*Ystr.str*/, litS__a) {
				// @ 19633 @ 850 @ open
				// Vcall: fn: <parse.Tvar object at 0x7f963ac81b10>
				// Vcall: args: [<parse.Tcall object at 0x7f963ac81e90>, <parse.Traw object at 0x7f963ac81ed0>]
				// Vcall: names: ['', '']
				// Vcall: star: None
				// Vcall: starstar: None
				// Making Global Yvar from 'PYE_FileDesc'
				// Vcall: fn: <parse.Tfield object at 0x7f963ac81b90>
				// Vcall: args: [<parse.Tvar object at 0x7f963ac81bd0>, <parse.Top object at 0x7f963ac81e10>, <parse.Tlit object at 0x7f963ac81e50>]
				// Vcall: names: ['', '', '']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vvar: local var filename -> 'a_filename'
				// Vcall: fn: <parse.Tvar object at 0x7f963ac81c10>
				// Vcall: args: [<parse.Tfield object at 0x7f963ac81c90>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				// Vcall: fn: <parse.Tvar object at 0x7f963ac81d10>
				// Vcall: args: [<parse.Tfield object at 0x7f963ac81d90>]
				// Vcall: names: ['']
				// Vcall: star: None
				// Vcall: starstar: None
				/*Maybe*/
				// ATTEMPT OptimizedGoCall: i_os.OpenFile TAKES ['string', 'int', 'os.FileMode'] RETURNS ['*os.File', 'error']
				//OptimizedGoCall NO: << os.OpenFile(/*Yvar.str*/a_filename,  JBitOr(/*L1165*//*Yint.str*/MkInt(int64(/*ForceInt*/JForceInt( MkGo(/*Yimport.str*/i_os.O_WRONLY) ))), /*Yint.str*/MkInt(int64(/*ForceInt*/JForceInt( MkGo(/*Yimport.str*/i_os.O_APPEND) )))) , /*Yint.str*/litI__0644) >>: Exception('OptimizedGoCall: Not supported yet: takes: os.FileMode',)
				return G_2_PYE_FileDesc( /*nando1*/ MkGo(i_os.OpenFile).Call( /*Yvar.str*/ a_filename, JBitOr( /*L1165*/ /*Yint.str*/ MkInt(int64( /*ForceInt*/ JForceInt(MkGo( /*Yimport.str*/ i_os.O_WRONLY)))) /*Yint.str*/, MkInt(int64( /*ForceInt*/ JForceInt(MkGo( /*Yimport.str*/ i_os.O_APPEND))))) /*Yint.str*/, litI__0644) /*Ybool.str*/, True)
				// $ 19633 $ 850 $
			} else {
				// @ 19735 @ 852 @ open
				// Vvar: local var mode -> 'a_mode'
				panic(M(( /*DoMod*/ JMod( /*Ystr.str*/ litS_35c93bbf31a763834dd39be445422ce5 /*Yvar.str*/, a_mode))))
				// $ 19735 $ 852 $
			}
		}
	}
	// $ 19472 $ 845 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: open

var specFunc_open = CallSpec{Name: "open", Args: []string{"filename", "mode"}, Defaults: []M{MissingM /*Ystr.str*/, litS__r}, Star: "", StarStar: ""}

func fnFunc_open(a0 M, a1 M) M { return G_2_open(a0, a1) }

//(tail)
// zip(p.argsPlus, typPlus): [('fd', None), ('writing', None)]
// typPlus: [None, None]
///////////////////////////////

func (self *C_PYE_FileDesc) M_2___init__(a_fd M, a_writing M) M {
	// @ 19936 @ 857 @ PYE_FileDesc.__init__
	// $ 19936 $ 857 $
	// @ 19952 @ 858 @ PYE_FileDesc.__init__
	// Vvar: local var writing -> 'a_writing'
	self.M_writing = /*Yvar.str*/ a_writing
	// $ 19952 $ 858 $
	// @ 19975 @ 859 @ PYE_FileDesc.__init__
	// Vvar: local var writing -> 'a_writing'
	if /*AsBool*/ JBool( /*Yvar.str*/ a_writing) {
		// @ 19993 @ 860 @ PYE_FileDesc.__init__
		// Vvar: local var fd -> 'a_fd'
		self.M_f = /*Yvar.str*/ a_fd
		// $ 19993 $ 860 $
		// @ 20007 @ 861 @ PYE_FileDesc.__init__
		// Vcall: fn: <parse.Tfield object at 0x7f963ac88750>
		// Vcall: args: [<parse.Tvar object at 0x7f963ac88790>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var fd -> 'a_fd'
		/*Maybe*/
		// ATTEMPT OptimizedGoCall: i_bufio.NewWriter TAKES ['io.Writer'] RETURNS ['*Writer']
		//OptimizedGoCall NO: << bufio.NewWriter(/*Yvar.str*/a_fd) >>: Exception('OptimizedGoCall: Not supported yet: takes: io.Writer',)
		self.M_b = MkGo(i_bufio.NewWriter).Call( /*Yvar.str*/ a_fd)
		// $ 20007 $ 861 $
	} else {
		// @ 20048 @ 863 @ PYE_FileDesc.__init__
		// Vvar: local var fd -> 'a_fd'
		self.M_f = /*Yvar.str*/ a_fd
		// $ 20048 $ 863 $
		// @ 20062 @ 864 @ PYE_FileDesc.__init__
		// Vcall: fn: <parse.Tfield object at 0x7f963ac88a50>
		// Vcall: args: [<parse.Tvar object at 0x7f963ac88a90>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var fd -> 'a_fd'
		/*Maybe*/
		// ATTEMPT OptimizedGoCall: i_bufio.NewReader TAKES ['io.Reader'] RETURNS ['*Reader']
		//OptimizedGoCall NO: << bufio.NewReader(/*Yvar.str*/a_fd) >>: Exception('OptimizedGoCall: Not supported yet: takes: io.Reader',)
		self.M_b = MkGo(i_bufio.NewReader).Call( /*Yvar.str*/ a_fd)
		// $ 20062 $ 864 $
	}
	// $ 19975 $ 859 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: __init__

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_PYE_FileDesc) M_0_read() M {
	// @ 20106 @ 867 @ PYE_FileDesc.read
	// $ 20106 $ 867 $
	// @ 20151 @ 868 @ PYE_FileDesc.read
	// Vcall: fn: <parse.Tvar object at 0x7f963ac88d10>
	// Vcall: args: [<parse.Tcall object at 0x7f963ac88e50>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x7f963ac88d90>
	// Vcall: args: [<parse.Tfield object at 0x7f963ac88e10>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	/*Maybe*/
	// ATTEMPT OptimizedGoCall: i_ioutil.ReadAll TAKES ['io.Reader'] RETURNS ['[]uint8', 'error']
	//OptimizedGoCall NO: << ioutil.ReadAll(self.M_b) >>: Exception('OptimizedGoCall: Not supported yet: takes: io.Reader',)
	return /*Ystr.str*/ MkStr( /*ForceString*/ JString(MkGo(i_ioutil.ReadAll).Call(self.M_b)))
	// $ 20151 $ 868 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: read

//(tail)
// zip(p.argsPlus, typPlus): [('x', None)]
// typPlus: [None]
///////////////////////////////

func (self *C_PYE_FileDesc) M_1_write(a_x M) M {
	// @ 20203 @ 871 @ PYE_FileDesc.write
	// $ 20203 $ 871 $
	// @ 20256 @ 872 @ PYE_FileDesc.write
	// Vcall: fn: <parse.Tfield object at 0x7f963cb09110>
	// Vcall: args: [<parse.Tcall object at 0x7f963cb091d0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tvar object at 0x7f963cb09150>
	// Vcall: args: [<parse.Tvar object at 0x7f963cb09190>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var x -> 'a_x'
	var sig_162 signature_Write___5b_5duint8_return_int_also_error
	_ = sig_162
	sig_162_o := self.M_b // Optimize QMeth
	if sig_162_o.X != nil {
		if p, ok := sig_162_o.X.Contents().(signature_Write___5b_5duint8_return_int_also_error); ok {
			sig_162 = p
		} else {
		}
	}
	var sig_162_r M = MissingM
	var opt_go_call_163_r_0 int   /*result*/
	var opt_go_call_163_r_1 error /*result*/
	if sig_162 != nil {
		// Vcall: fn: <parse.Tvar object at 0x7f963cb09150>
		// Vcall: args: [<parse.Tvar object at 0x7f963cb09190>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var x -> 'a_x'
		// ATTEMPT OptimizedGoCall: sig_162.Write TAKES ('[]uint8',) RETURNS ('int', 'error')
		var opt_go_call_163_t_0 []uint8 = []uint8( /*Ystr.AsByt*/ []byte( /*ForceString*/ JString( /*Yvar.str*/ a_x)))
		opt_go_call_163_r_0, opt_go_call_163_r_1 = /*OptimizedGoCall OK*/ sig_162.Write(opt_go_call_163_t_0) // OptimizedGoCall OK
		_ = opt_go_call_163_r_0
		_ = opt_go_call_163_r_1
		if opt_go_call_163_r_1 != nil {
			panic(opt_go_call_163_r_1) /*Check magic_error*/
		}

	} else {
		sig_162_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_1_Write(sig_162_o /*Ystr.str*/, MkStr( /*ForceString*/ JString( /*Yvar.str*/ a_x)))
	}
	_ = sig_162_r
	// (* Yeither: 'sig_162_r' :: <type 'str'> ; <codegen.Yint object at 0x7f963cb28410> :: <class 'codegen.Yint'> *)
	if sig_162_r == MissingM {
		sig_162_r = /*Yint.str*/ MkInt(int64(opt_go_call_163_r_0))
	}
	_ = /*Yeither.str*/ sig_162_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28210>
	// $ 20256 $ 872 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: write

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_PYE_FileDesc) M_0_flush() M {
	// @ 20293 @ 875 @ PYE_FileDesc.flush
	// $ 20293 $ 875 $
	// @ 20341 @ 876 @ PYE_FileDesc.flush
	// Vcall: fn: <parse.Tfield object at 0x7f963cb09490>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	var sig_164 signature_Flush___return_error
	_ = sig_164
	sig_164_o := self.M_b // Optimize QMeth
	if sig_164_o.X != nil {
		if p, ok := sig_164_o.X.Contents().(signature_Flush___return_error); ok {
			sig_164 = p
		} else {
		}
	}
	var sig_164_r M = MissingM
	var opt_go_call_165_r_0 error /*result*/
	if sig_164 != nil {
		// ATTEMPT OptimizedGoCall: sig_164.Flush TAKES () RETURNS ('error',)
		opt_go_call_165_r_0 = /*OptimizedGoCall OK*/ sig_164.Flush() // OptimizedGoCall OK
		_ = opt_go_call_165_r_0
		if opt_go_call_165_r_0 != nil {
			panic(opt_go_call_165_r_0) /*Check magic_error*/
		}

	} else {
		sig_164_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Flush(sig_164_o)
	}
	_ = sig_164_r
	// (* Yeither: 'sig_164_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
	if sig_164_r == MissingM {
		sig_164_r = None
	}
	_ = /*Yeither.str*/ sig_164_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28410>
	// $ 20341 $ 876 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: flush

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_PYE_FileDesc) M_0_Flush() M {
	// @ 20371 @ 878 @ PYE_FileDesc.Flush
	// $ 20371 $ 878 $
	// @ 20419 @ 879 @ PYE_FileDesc.Flush
	// Vcall: fn: <parse.Tfield object at 0x7f963cb09750>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	var sig_166 signature_Flush___return_error
	_ = sig_166
	sig_166_o := self.M_b // Optimize QMeth
	if sig_166_o.X != nil {
		if p, ok := sig_166_o.X.Contents().(signature_Flush___return_error); ok {
			sig_166 = p
		} else {
		}
	}
	var sig_166_r M = MissingM
	var opt_go_call_167_r_0 error /*result*/
	if sig_166 != nil {
		// ATTEMPT OptimizedGoCall: sig_166.Flush TAKES () RETURNS ('error',)
		opt_go_call_167_r_0 = /*OptimizedGoCall OK*/ sig_166.Flush() // OptimizedGoCall OK
		_ = opt_go_call_167_r_0
		if opt_go_call_167_r_0 != nil {
			panic(opt_go_call_167_r_0) /*Check magic_error*/
		}

	} else {
		sig_166_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Flush(sig_166_o)
	}
	_ = sig_166_r
	// (* Yeither: 'sig_166_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
	if sig_166_r == MissingM {
		sig_166_r = None
	}
	_ = /*Yeither.str*/ sig_166_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28050>
	// $ 20419 $ 879 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Flush

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_PYE_FileDesc) M_0_close() M {
	// @ 20450 @ 882 @ PYE_FileDesc.close
	// $ 20450 $ 882 $
	// @ 20472 @ 883 @ PYE_FileDesc.close
	if /*AsBool*/ JBool(self.M_writing) {
		// @ 20491 @ 884 @ PYE_FileDesc.close
		// Vcall: fn: <parse.Tfield object at 0x7f963cb09a90>
		// Vcall: args: []
		// Vcall: names: []
		// Vcall: star: None
		// Vcall: starstar: None
		var sig_168 signature_Flush___return_error
		_ = sig_168
		sig_168_o := self.M_b // Optimize QMeth
		if sig_168_o.X != nil {
			if p, ok := sig_168_o.X.Contents().(signature_Flush___return_error); ok {
				sig_168 = p
			} else {
			}
		}
		var sig_168_r M = MissingM
		var opt_go_call_169_r_0 error /*result*/
		if sig_168 != nil {
			// ATTEMPT OptimizedGoCall: sig_168.Flush TAKES () RETURNS ('error',)
			opt_go_call_169_r_0 = /*OptimizedGoCall OK*/ sig_168.Flush() // OptimizedGoCall OK
			_ = opt_go_call_169_r_0
			if opt_go_call_169_r_0 != nil {
				panic(opt_go_call_169_r_0) /*Check magic_error*/
			}

		} else {
			sig_168_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Flush(sig_168_o)
		}
		_ = sig_168_r
		// (* Yeither: 'sig_168_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
		if sig_168_r == MissingM {
			sig_168_r = None
		}
		_ = /*Yeither.str*/ sig_168_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28390>
		// $ 20491 $ 884 $
	}
	// $ 20472 $ 883 $
	// @ 20506 @ 885 @ PYE_FileDesc.close
	// Vcall: fn: <parse.Tfield object at 0x7f963cb09c90>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	var sig_170 signature_Close___return_error
	_ = sig_170
	sig_170_o := self.M_f // Optimize QMeth
	if sig_170_o.X != nil {
		if p, ok := sig_170_o.X.Contents().(signature_Close___return_error); ok {
			sig_170 = p
		} else {
		}
	}
	var sig_170_r M = MissingM
	var opt_go_call_171_r_0 error /*result*/
	if sig_170 != nil {
		// ATTEMPT OptimizedGoCall: sig_170.Close TAKES () RETURNS ('error',)
		opt_go_call_171_r_0 = /*OptimizedGoCall OK*/ sig_170.Close() // OptimizedGoCall OK
		_ = opt_go_call_171_r_0
		if opt_go_call_171_r_0 != nil {
			panic(opt_go_call_171_r_0) /*Check magic_error*/
		}

	} else {
		sig_170_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Close(sig_170_o)
	}
	_ = sig_170_r
	// (* Yeither: 'sig_170_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
	if sig_170_r == MissingM {
		sig_170_r = None
	}
	_ = /*Yeither.str*/ sig_170_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb283d0>
	// $ 20506 $ 885 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: close

//(tail)
// zip(p.argsPlus, typPlus): []
// typPlus: []
///////////////////////////////

func (self *C_PYE_FileDesc) M_0_Close() M {
	// @ 20536 @ 887 @ PYE_FileDesc.Close
	// $ 20536 $ 887 $
	// @ 20558 @ 888 @ PYE_FileDesc.Close
	if /*AsBool*/ JBool(self.M_writing) {
		// @ 20577 @ 889 @ PYE_FileDesc.Close
		// Vcall: fn: <parse.Tfield object at 0x7f963cb09fd0>
		// Vcall: args: []
		// Vcall: names: []
		// Vcall: star: None
		// Vcall: starstar: None
		var sig_172 signature_Flush___return_error
		_ = sig_172
		sig_172_o := self.M_b // Optimize QMeth
		if sig_172_o.X != nil {
			if p, ok := sig_172_o.X.Contents().(signature_Flush___return_error); ok {
				sig_172 = p
			} else {
			}
		}
		var sig_172_r M = MissingM
		var opt_go_call_173_r_0 error /*result*/
		if sig_172 != nil {
			// ATTEMPT OptimizedGoCall: sig_172.Flush TAKES () RETURNS ('error',)
			opt_go_call_173_r_0 = /*OptimizedGoCall OK*/ sig_172.Flush() // OptimizedGoCall OK
			_ = opt_go_call_173_r_0
			if opt_go_call_173_r_0 != nil {
				panic(opt_go_call_173_r_0) /*Check magic_error*/
			}

		} else {
			sig_172_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Flush(sig_172_o)
		}
		_ = sig_172_r
		// (* Yeither: 'sig_172_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
		if sig_172_r == MissingM {
			sig_172_r = None
		}
		_ = /*Yeither.str*/ sig_172_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28390>
		// $ 20577 $ 889 $
	}
	// $ 20558 $ 888 $
	// @ 20592 @ 890 @ PYE_FileDesc.Close
	// Vcall: fn: <parse.Tfield object at 0x7f963cb10210>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	var sig_174 signature_Close___return_error
	_ = sig_174
	sig_174_o := self.M_f // Optimize QMeth
	if sig_174_o.X != nil {
		if p, ok := sig_174_o.X.Contents().(signature_Close___return_error); ok {
			sig_174 = p
		} else {
		}
	}
	var sig_174_r M = MissingM
	var opt_go_call_175_r_0 error /*result*/
	if sig_174 != nil {
		// ATTEMPT OptimizedGoCall: sig_174.Close TAKES () RETURNS ('error',)
		opt_go_call_175_r_0 = /*OptimizedGoCall OK*/ sig_174.Close() // OptimizedGoCall OK
		_ = opt_go_call_175_r_0
		if opt_go_call_175_r_0 != nil {
			panic(opt_go_call_175_r_0) /*Check magic_error*/
		}

	} else {
		sig_174_r = /*GeneralCallMeth Slow*/ /*invoker*/ F_INVOKE_0_Close(sig_174_o)
	}
	_ = sig_174_r
	// (* Yeither: 'sig_174_r' :: <type 'str'> ; 'None' :: <type 'str'> *)
	if sig_174_r == MissingM {
		sig_174_r = None
	}
	_ = /*Yeither.str*/ sig_174_r // Assign void: = type: <class 'codegen.Yeither'> repr: <codegen.Yeither object at 0x7f963cb28110>
	// $ 20592 $ 890 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: Close

//(tail)

type C_PYE_FileDesc struct {
	C_object
	M_b       M
	M_f       M
	M_writing M
}

func init() {
	if Classes == nil {
		Classes = make(map[string]reflect.Type)
	}
	Classes[`BUILTINS.C_PYE_FileDesc`] = reflect.TypeOf(C_PYE_FileDesc{})
}

func (o *C_PYE_FileDesc) PtrC_PYE_FileDesc() *C_PYE_FileDesc {
	return o
}

func (o *C_PYE_FileDesc) PtrC_object() *C_object {
	return &o.C_object
}

func (o *C_PYE_FileDesc) GET_b() M  { return o.M_b }
func (o *C_PYE_FileDesc) SET_b(x M) { o.M_b = x }

func (o *C_PYE_FileDesc) GET_f() M  { return o.M_f }
func (o *C_PYE_FileDesc) SET_f(x M) { o.M_f = x }

func (o *C_PYE_FileDesc) GET_writing() M  { return o.M_writing }
func (o *C_PYE_FileDesc) SET_writing(x M) { o.M_writing = x }

var specMeth_0_PYE_FileDesc__Close = CallSpec{Name: "PYE_FileDesc::Close", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_Close() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PYE_FileDesc__Close}, o.M_0_Close})
}

var specMeth_0_PYE_FileDesc__Flush = CallSpec{Name: "PYE_FileDesc::Flush", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_Flush() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PYE_FileDesc__Flush}, o.M_0_Flush})
}

var specMeth_2_PYE_FileDesc____init__ = CallSpec{Name: "PYE_FileDesc::__init__", Args: []string{"fd", "writing"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET___init__() M {
	return MForge(&PCall2{PNewCallable{CallSpec: &specMeth_2_PYE_FileDesc____init__}, o.M_2___init__})
}

var specMeth_0_PYE_FileDesc__close = CallSpec{Name: "PYE_FileDesc::close", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_close() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PYE_FileDesc__close}, o.M_0_close})
}

var specMeth_0_PYE_FileDesc__flush = CallSpec{Name: "PYE_FileDesc::flush", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_flush() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PYE_FileDesc__flush}, o.M_0_flush})
}

var specMeth_0_PYE_FileDesc__read = CallSpec{Name: "PYE_FileDesc::read", Args: []string{}, Defaults: []M{}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_read() M {
	return MForge(&PCall0{PNewCallable{CallSpec: &specMeth_0_PYE_FileDesc__read}, o.M_0_read})
}

var specMeth_1_PYE_FileDesc__write = CallSpec{Name: "PYE_FileDesc::write", Args: []string{"x"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

func (o *C_PYE_FileDesc) GET_write() M {
	return MForge(&PCall1{PNewCallable{CallSpec: &specMeth_1_PYE_FileDesc__write}, o.M_1_write})
}
func (o *C_PYE_FileDesc) Rye_ClearFields__() {
	o.M_b = None
	o.M_f = None
	o.M_writing = None
	// Making Global Yvar from 'object'
	// superclass: /*Yvar.str*/G_object
}

func (o *pFunc_PYE_FileDesc) Superclass() M {
	// Making Global Yvar from 'object'
	return /*Yvar.str*/ G_object
}

func (o *C_PYE_FileDesc) PType() M           { return G_PYE_FileDesc }
func (o *pFunc_PYE_FileDesc) Repr() string   { return "PYE_FileDesc" }
func (o *pFunc_PYE_FileDesc) String() string { return "<class PYE_FileDesc>" }

//(tail)
// { native M
//(tail)

// io.Writer protocol for writing:
func (self *C_PYE_FileDesc) Write(p []byte) (n int, err error) {
	return JContents(self.M_b).(io.Writer).Write(p)
}
func (self *C_PYE_FileDesc) Flush() error {
	self.M_0_Flush()
	return nil
}

//(tail)
// } native M
//(tail)
// zip(p.argsPlus, typPlus): [('f', None), ('a', None), ('b', None), ('c', None), ('d', None)]
// typPlus: [None, None, None, None, None]
///////////////////////////////

func G_5__rye__force_generation_of_call_4_(a_f M, a_a M, a_b M, a_c M, a_d M) M {
	// @ 20919 @ 904 @ _rye__force_generation_of_call_4_
	// Vcall: fn: <parse.Tvar object at 0x7f963cb10410>
	// Vcall: args: [<parse.Tvar object at 0x7f963cb10450>, <parse.Tvar object at 0x7f963cb10490>, <parse.Tvar object at 0x7f963cb104d0>, <parse.Tvar object at 0x7f963cb10510>]
	// Vcall: names: ['', '', '', '']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var f -> 'a_f'
	// Vvar: local var a -> 'a_a'
	// Vvar: local var b -> 'a_b'
	// Vvar: local var c -> 'a_c'
	// Vvar: local var d -> 'a_d'
	return CALL_4( /*nando2*/ M( /*Yvar.str*/ a_f) /*Yvar.str*/, a_a /*Yvar.str*/, a_b /*Yvar.str*/, a_c /*Yvar.str*/, a_d)
	// $ 20919 $ 904 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: _rye__force_generation_of_call_4_

var specFunc__rye__force_generation_of_call_4_ = CallSpec{Name: "_rye__force_generation_of_call_4_", Args: []string{"f", "a", "b", "c", "d"}, Defaults: []M{MissingM, MissingM, MissingM, MissingM, MissingM}, Star: "", StarStar: ""}

type pFunc__rye__force_generation_of_call_4_ struct{ PNewCallable }

func (o *pFunc__rye__force_generation_of_call_4_) Contents() interface{} {
	return G__rye__force_generation_of_call_4_
}
func (o pFunc__rye__force_generation_of_call_4_) Call5(a0 M, a1 M, a2 M, a3 M, a4 M) M {
	return G_5__rye__force_generation_of_call_4_(a0, a1, a2, a3, a4)
}

func (o pFunc__rye__force_generation_of_call_4_) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_5__rye__force_generation_of_call_4_(argv[0], argv[1], argv[2], argv[3], argv[4])
}

//(tail)
// zip(p.argsPlus, typPlus): [('fd', None), ('writing', None)]
// typPlus: [None, None]
///////////////////////////////

func G_2_PYE_FileDesc(a_fd M, a_writing M) M {
	var v_rye_result__ M = None
	_ = v_rye_result__
	// @ 19774 @ 854 @ PYE_FileDesc
	// { native F
	z := new(C_PYE_FileDesc)
	z.Self = z
	z.Rye_ClearFields__()
	// } native F
	// $ 19774 $ 854 $
	// @ 19774 @ 854 @ PYE_FileDesc
	// @@@@@@ Creating var "rye_result__" in scope @@@@@@
	v_rye_result__ = MkObj(&z.PBase) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 19774 $ 854 $
	// @ 19774 @ 854 @ PYE_FileDesc
	// Vcall: fn: <parse.Tfield object at 0x7f963cb10990>
	// Vcall: args: [<parse.Tvar object at 0x7f963cb10a10>, <parse.Tvar object at 0x7f963cb10a50>]
	// Vcall: names: ['fd', 'writing']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var rye_result__ -> 'v_rye_result__'
	// Vvar: local var fd -> 'a_fd'
	// Vvar: local var writing -> 'a_writing'
	_ = M( /*nando3*/ F_GET___init__( /*Yvar.str*/ v_rye_result__)).X.(ICallV).CallV([]M{}, nil, []KV{KV{"fd" /*Yvar.str*/, a_fd}, KV{"writing" /*Yvar.str*/, a_writing}}, nil) // Assign void: = type: <type 'str'> repr: 'M(/*nando3*/  F_GET___init__(/*Yvar.str*/v_rye_result__) ).X.(ICallV).CallV([]M{}, nil, []KV{KV{"fd", /*Yvar.str*/a_fd}, KV{"writing", /*Yvar.str*/a_writing}}, nil) '
	// $ 19774 $ 854 $
	// @ 19774 @ 854 @ PYE_FileDesc
	// Vvar: local var rye_result__ -> 'v_rye_result__'
	return /*Yvar.str*/ v_rye_result__
	// $ 19774 $ 854 $
	// bottom out: return None
	return None

}

///////////////////////////////
// name: PYE_FileDesc

var specFunc_PYE_FileDesc = CallSpec{Name: "PYE_FileDesc", Args: []string{"fd", "writing"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

type pFunc_PYE_FileDesc struct{ PNewCallable }

func (o *pFunc_PYE_FileDesc) Contents() interface{} {
	return G_PYE_FileDesc
}
func (o pFunc_PYE_FileDesc) Call2(a0 M, a1 M) M {
	return G_2_PYE_FileDesc(a0, a1)
}

func (o pFunc_PYE_FileDesc) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_2_PYE_FileDesc(argv[0], argv[1])
}

//(end tail)

var G_Exception M                         // *pFunc_Exception
var G_PYE_FileDesc M                      // *pFunc_PYE_FileDesc
var G__rye__force_generation_of_call_4_ M // *pFunc__rye__force_generation_of_call_4_
var G_all M                               // *pFunc_all
var G_any M                               // *pFunc_any
var G_bool M                              // *pFunc_bool
var G_byt M                               // *pFunc_byt
var G_callable M                          // *pFunc_callable
var G_chr M                               // *pFunc_chr
var G_cmp M                               // *pFunc_cmp
var G_dict M                              // *pFunc_dict
var G_float M                             // *pFunc_float
var G_getattr M                           // *pFunc_getattr
var G_go_deref M                          // *pFunc_go_deref
var G_hash M                              // *pFunc_hash
var G_id M                                // *pFunc_id
var G_int M                               // *pFunc_int
var G_isinstance M                        // *pFunc_isinstance
var G_issubclass M                        // *pFunc_issubclass
var G_len M                               // *pFunc_len
var G_list M                              // *pFunc_list
var G_map M                               // *pFunc_map
var G_max M                               // *pFunc_max
var G_min M                               // *pFunc_min
var G_mkbyt M                             // *pFunc_mkbyt
var G_object M                            // *pFunc_object
var G_open M                              // *pFunc_open
var G_ord M                               // *pFunc_ord
var G_range M                             // *pFunc_range
var G_reduce M                            // *pFunc_reduce
var G_repr M                              // *pFunc_repr
var G_rye_chan M                          // *pFunc_rye_chan
var G_rye_opts M                          // *pFunc_rye_opts
var G_rye_pickle M                        // *pFunc_rye_pickle
var G_rye_stack M                         // *pFunc_rye_stack
var G_rye_unpickle M                      // *pFunc_rye_unpickle
var G_rye_what M                          // *pFunc_rye_what
var G_rye_zip_padding_with_None M         // *pFunc_rye_zip_padding_with_None
var G_set M                               // *pFunc_set
var G_setattr M                           // *pFunc_setattr
var G_setattrs M                          // *pFunc_setattrs
var G_sorted M                            // *pFunc_sorted
var G_str M                               // *pFunc_str
var G_sum M                               // *pFunc_sum
var G_tuple M                             // *pFunc_tuple
var G_type M                              // *pFunc_type
var G_xrange M                            // *pFunc_xrange
var G_zip M                               // *pFunc_zip

func init /*New_Module*/ () {
	G_Exception = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_Exception}, fnFunc_Exception})
	G_PYE_FileDesc = MForge(&pFunc_PYE_FileDesc{PNewCallable{CallSpec: &specFunc_PYE_FileDesc}})
	G__rye__force_generation_of_call_4_ = MForge(&pFunc__rye__force_generation_of_call_4_{PNewCallable{CallSpec: &specFunc__rye__force_generation_of_call_4_}})
	G_all = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_all}, fnFunc_all})
	G_any = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_any}, fnFunc_any})
	G_bool = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_bool}, fnFunc_bool})
	G_byt = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_byt}, fnFunc_byt})
	G_callable = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_callable}, fnFunc_callable})
	G_chr = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_chr}, fnFunc_chr})
	G_cmp = MForge(&PCall2{PNewCallable{CallSpec: &specFunc_cmp}, fnFunc_cmp})
	G_dict = MForge(&pFunc_dict{PNewCallable{CallSpec: &specFunc_dict}})
	G_float = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_float}, fnFunc_float})
	G_getattr = MForge(&pFunc_getattr{PNewCallable{CallSpec: &specFunc_getattr}})
	G_go_deref = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_go_deref}, fnFunc_go_deref})
	G_hash = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_hash}, fnFunc_hash})
	G_id = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_id}, fnFunc_id})
	G_int = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_int}, fnFunc_int})
	G_isinstance = MForge(&PCall2{PNewCallable{CallSpec: &specFunc_isinstance}, fnFunc_isinstance})
	G_issubclass = MForge(&PCall2{PNewCallable{CallSpec: &specFunc_issubclass}, fnFunc_issubclass})
	G_len = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_len}, fnFunc_len})
	G_list = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_list}, fnFunc_list})
	G_map = MForge(&pFunc_map{PNewCallable{CallSpec: &specFunc_map}})
	G_max = MForge(&pFunc_max{PNewCallable{CallSpec: &specFunc_max}})
	G_min = MForge(&pFunc_min{PNewCallable{CallSpec: &specFunc_min}})
	G_mkbyt = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_mkbyt}, fnFunc_mkbyt})
	G_object = MForge(&PCall0{PNewCallable{CallSpec: &specFunc_object}, fnFunc_object})
	G_open = MForge(&PCall2{PNewCallable{CallSpec: &specFunc_open}, fnFunc_open})
	G_ord = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_ord}, fnFunc_ord})
	G_range = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_range}, fnFunc_range})
	G_reduce = MForge(&PCall3{PNewCallable{CallSpec: &specFunc_reduce}, fnFunc_reduce})
	G_repr = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_repr}, fnFunc_repr})
	G_rye_chan = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_rye_chan}, fnFunc_rye_chan})
	G_rye_opts = MForge(&PCall0{PNewCallable{CallSpec: &specFunc_rye_opts}, fnFunc_rye_opts})
	G_rye_pickle = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_rye_pickle}, fnFunc_rye_pickle})
	G_rye_stack = MForge(&PCall0{PNewCallable{CallSpec: &specFunc_rye_stack}, fnFunc_rye_stack})
	G_rye_unpickle = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_rye_unpickle}, fnFunc_rye_unpickle})
	G_rye_what = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_rye_what}, fnFunc_rye_what})
	G_rye_zip_padding_with_None = MForge(&pFunc_rye_zip_padding_with_None{PNewCallable{CallSpec: &specFunc_rye_zip_padding_with_None}})
	G_set = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_set}, fnFunc_set})
	G_setattr = MForge(&PCall3{PNewCallable{CallSpec: &specFunc_setattr}, fnFunc_setattr})
	G_setattrs = MForge(&pFunc_setattrs{PNewCallable{CallSpec: &specFunc_setattrs}})
	G_sorted = MForge(&pFunc_sorted{PNewCallable{CallSpec: &specFunc_sorted}})
	G_str = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_str}, fnFunc_str})
	G_sum = MForge(&PCall2{PNewCallable{CallSpec: &specFunc_sum}, fnFunc_sum})
	G_tuple = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_tuple}, fnFunc_tuple})
	G_type = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_type}, fnFunc_type})
	G_xrange = MForge(&PCall1{PNewCallable{CallSpec: &specFunc_xrange}, fnFunc_xrange})
	G_zip = MForge(&pFunc_zip{PNewCallable{CallSpec: &specFunc_zip}})
	inner_eval_module()
}

var BuiltinMap = map[string]*M{
	"Exception":                         &G_Exception,
	"PYE_FileDesc":                      &G_PYE_FileDesc,
	"_rye__force_generation_of_call_4_": &G__rye__force_generation_of_call_4_,
	"all":                       &G_all,
	"any":                       &G_any,
	"bool":                      &G_bool,
	"byt":                       &G_byt,
	"callable":                  &G_callable,
	"chr":                       &G_chr,
	"cmp":                       &G_cmp,
	"dict":                      &G_dict,
	"float":                     &G_float,
	"getattr":                   &G_getattr,
	"go_deref":                  &G_go_deref,
	"hash":                      &G_hash,
	"id":                        &G_id,
	"int":                       &G_int,
	"isinstance":                &G_isinstance,
	"issubclass":                &G_issubclass,
	"len":                       &G_len,
	"list":                      &G_list,
	"map":                       &G_map,
	"max":                       &G_max,
	"min":                       &G_min,
	"mkbyt":                     &G_mkbyt,
	"object":                    &G_object,
	"open":                      &G_open,
	"ord":                       &G_ord,
	"range":                     &G_range,
	"reduce":                    &G_reduce,
	"repr":                      &G_repr,
	"rye_chan":                  &G_rye_chan,
	"rye_opts":                  &G_rye_opts,
	"rye_pickle":                &G_rye_pickle,
	"rye_stack":                 &G_rye_stack,
	"rye_unpickle":              &G_rye_unpickle,
	"rye_what":                  &G_rye_what,
	"rye_zip_padding_with_None": &G_rye_zip_padding_with_None,
	"set":      &G_set,
	"setattr":  &G_setattr,
	"setattrs": &G_setattrs,
	"sorted":   &G_sorted,
	"str":      &G_str,
	"sum":      &G_sum,
	"tuple":    &G_tuple,
	"type":     &G_type,
	"xrange":   &G_xrange,
	"zip":      &G_zip,
}

var BuiltinObj = MakeModuleObject(BuiltinMap, "BUILTINS/BUILTINS")

var litI__0 = MkInt(0)
var litI__0644 = MkInt(0644)
var litI__1 = MkInt(1)
var litI__2 = MkInt(2)
var litI___XminusX_1 = MkInt(-1)
var litS_070c4397445592bd39b479ef9a7724b4 = MkStr("no args to max()")
var litS_0bf6f656326a86f95c011a3c50b7f4e2 = MkStr("Does self start with string x?")
var litS_0f001beb5b015cf6114fc39bddb454ea = MkStr("Return self converted to title case.")
var litS_1dc4ee90911bbfd78620f5b12d4ed868 = MkStr("Return the value of the named field on self.")
var litS_1eef9fe0783015967d67ed994ddb207d = MkStr("Does self end with string x?")
var litS_279ac9f834a4472f993fff33803f8da5 = MkStr("Split self with delimiter x at most n times.  If x is None, split on white space.")
var litS_35c93bbf31a763834dd39be445422ce5 = MkStr("open: Unknown mode: %q")
var litS_3647ef5f9a6f72c3a0e004ff3cef5941 = MkStr("Return the index of the last occurance of x in self, or -1 if not found.")
var litS_368f620e2e4727f35d535457db99a5f7 = MkStr("Return the index of the last occurance of x in self, or throw an exception.")
var litS_40bf9d96be0abd122a4f38fc0a015340 = MkStr("Join the elements of vec adding self between the elements.")
var litS_4aa730649437fff2d869a67c054d8d7d = MkStr("Return self with chars in x stripped away from front.")
var litS_4b59f91f2cd4703181a9d3c6bc4fc68b = MkStr("The internal type returned from the builtin open() function.  Go's io.Writer protocol is also supported.")
var litS_543d5f51c6a79a74ebcdfbb50447ae5f = MkStr("Close the file.")
var litS_54722135a5796e011e8582f72f64a77a = MkStr("Set the value of the named field on self.")
var litS_5477c6e46d473902bbd5a31230a4e07c = MkStr("Receive a rye value from the chan, blocking until one is ready.  Throw \x22EOF\x22 if the channel was closed.")
var litS_56ec67083aab08bd4f70164fed9d28cf = MkStr("Return the index of the first occurance of x in self, or -1 if not found.")
var litS_60914cc954768078ed76132cf1f792cd = MkStr("PStr is a fake class to hold methods for the builtin type str.")
var litS_66663bef6fb2e823dbdc75c792c74ce5 = MkStr("Open the named file, with given mode, as in Python.  Returns an instance of PYE_FileDesc.")
var litS_6dfc06aa7418ffcee1111985ca9a2965 = MkStr("Flush buffered written bytes to the file.")
var litS_6edb409879889cb135cf9aee54d23b66 = MkStr("Return self with chars in x stripped away from end.")
var litS_6fe87cae7f4cfc67d4ffb571068ceb39 = MkStr("map called with no lists")
var litS_70e19ea725c8f340ba7e1cdce8dbfa4f = MkStr("Read the rest of the file as a string.")
var litS_713028014dbf53462a9b38306c6cbe75 = MkStr("Return self with nonoverlapping occurances of old replaced with new at most count times.")
var litS_79af4b5efd05f745db48dbede682ad6e = MkStr("Send a special record Causing the receiver to throw the given exception when it is read.")
var litS_7ddcb4ed851122cfda86a9b8a9380e5c = MkStr("PByt is a fake class to hold methods for the builtin type byt.")
var litS_7e29bea0cb35948013d1faa4907aca1b = MkStr("Return self with chars in x stripped away from front and end.")
var litS_80c0c708b72198e1f8c033f6eebe9716 = MkStr("Are all runes in self unicode letters?")
var litS_868beb4fbac148f9a5d1e20101485c96 = MkStr("no args to min()")
var litS_946db78d35c01cf8bf719facef908525 = MkStr("Close the channel, signaling that nothing more will be Sent on it.")
var litS__Internal_XpointX_ = MkStr("Internal.")
var litS__ValueError = MkStr("ValueError")
var litS__a = MkStr("a")
var litS__r = MkStr("r")
var litS__w = MkStr("w")
var litS_ac2725fccc943619fc1334bff030aa13 = MkStr("Return self converted to lower case.")
var litS_b0ae81038217c242e9a30240e373d4e8 = MkStr("rye_chan creates a Started rye channel of the given buffer size.")
var litS_b8cc29323b546da610323995bed5c525 = MkStr("Return self converted to upper case.")
var litS_c3e2c5b4de9c3a86f8baadfd17315a39 = MkStr("object is the construtor for builtin object type.")
var litS_c520965675f28d01f68d665c4adbd254 = MkStr("Return the index of the first occurance of x in self, or throw an exception.")
var litS_cc4f56b2f5251c2ca6b0b4409518e018 = MkStr("Are all runes in self unicode letters or digits?")
var litS_d6221c0c57599569fa8b67dc1d259878 = MkStr(" \x09\x0a\x0d")
var litS_d9f3ab28efe6036aea13bd5e1546f50c = MkStr("Are all runes in self unicode spaces?")
var litS_dfe37f9f86c697b4d4b0e2345835ae88 = MkStr("C_object is a fake class to hold methods for the builtin class object.")
var litS_e2cf2794c521c80f6b48678a6b077050 = MkStr("Are all runes in self unicode digits?")
var litS_e571f4c1483837c764f8000ebb49884f = MkStr("Are all runes in self unicode upper case letters?")
var litS_ea7f2382bfdce7e530b7ecf99e357632 = MkStr("Send a rye value on the chan.")
var litS_ecf688d5037bbeed8af7f28316500cd8 = MkStr("Receive a rye value from the chan, blocking until one is ready.  ok if false if the channel was closed.")
var litS_edc53bbee7f5513020128ff92fa921f5 = MkStr("Are all runes in self unicode lower case letters?")
var litS_f8fe50c7918f2458ba4fed7b3eee3a77 = MkStr("C_channel is a fake class to hold methods for the builtin rye_chan type.")
var litS_fabf74a920cc6f122a46ee78a40e9b38 = MkStr("Write tye bytes in x, which can be str or byt.")

func F_INVOKE_0_Close(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_Close:
		return x.M_0_Close()
	case I_GET_Close:
		tmp := x.GET_Close()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("Close")
	}
	panic(fmt.Sprintf("Cannot invoke 'Close' with 0 arguments on %v", fn))
}

type I_INVOKE_0_Close interface{ M_0_Close() M }

func F_INVOKE_0_Flush(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_Flush:
		return x.M_0_Flush()
	case I_GET_Flush:
		tmp := x.GET_Flush()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("Flush")
	}
	panic(fmt.Sprintf("Cannot invoke 'Flush' with 0 arguments on %v", fn))
}

type I_INVOKE_0_Flush interface{ M_0_Flush() M }

func F_INVOKE_0_Wait(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_Wait:
		return x.M_0_Wait()
	case I_GET_Wait:
		tmp := x.GET_Wait()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("Wait")
	}
	panic(fmt.Sprintf("Cannot invoke 'Wait' with 0 arguments on %v", fn))
}

type I_INVOKE_0_Wait interface{ M_0_Wait() M }

func F_INVOKE_0_Warm(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_Warm:
		return x.M_0_Warm()
	case I_GET_Warm:
		tmp := x.GET_Warm()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("Warm")
	}
	panic(fmt.Sprintf("Cannot invoke 'Warm' with 0 arguments on %v", fn))
}

type I_INVOKE_0_Warm interface{ M_0_Warm() M }

func F_INVOKE_0___init__(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0___init__:
		return x.M_0___init__()
	case I_GET___init__:
		tmp := x.GET___init__()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("__init__")
	}
	panic(fmt.Sprintf("Cannot invoke '__init__' with 0 arguments on %v", fn))
}

type I_INVOKE_0___init__ interface{ M_0___init__() M }

func F_INVOKE_0_clear(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_clear:
		return x.M_0_clear()
	case I_GET_clear:
		tmp := x.GET_clear()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("clear")
	}
	panic(fmt.Sprintf("Cannot invoke 'clear' with 0 arguments on %v", fn))
}

type I_INVOKE_0_clear interface{ M_0_clear() M }

func F_INVOKE_0_close(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_close:
		return x.M_0_close()
	case I_GET_close:
		tmp := x.GET_close()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("close")
	}
	panic(fmt.Sprintf("Cannot invoke 'close' with 0 arguments on %v", fn))
}

type I_INVOKE_0_close interface{ M_0_close() M }

func F_INVOKE_0_copy(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_copy:
		return x.M_0_copy()
	case I_GET_copy:
		tmp := x.GET_copy()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("copy")
	}
	panic(fmt.Sprintf("Cannot invoke 'copy' with 0 arguments on %v", fn))
}

type I_INVOKE_0_copy interface{ M_0_copy() M }

func F_INVOKE_0_flush(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_flush:
		return x.M_0_flush()
	case I_GET_flush:
		tmp := x.GET_flush()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("flush")
	}
	panic(fmt.Sprintf("Cannot invoke 'flush' with 0 arguments on %v", fn))
}

type I_INVOKE_0_flush interface{ M_0_flush() M }

func F_INVOKE_0_isalnum(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_isalnum:
		return x.M_0_isalnum()
	case I_GET_isalnum:
		tmp := x.GET_isalnum()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("isalnum")
	}
	panic(fmt.Sprintf("Cannot invoke 'isalnum' with 0 arguments on %v", fn))
}

type I_INVOKE_0_isalnum interface{ M_0_isalnum() M }

func F_INVOKE_0_isalpha(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_isalpha:
		return x.M_0_isalpha()
	case I_GET_isalpha:
		tmp := x.GET_isalpha()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("isalpha")
	}
	panic(fmt.Sprintf("Cannot invoke 'isalpha' with 0 arguments on %v", fn))
}

type I_INVOKE_0_isalpha interface{ M_0_isalpha() M }

func F_INVOKE_0_isdigit(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_isdigit:
		return x.M_0_isdigit()
	case I_GET_isdigit:
		tmp := x.GET_isdigit()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("isdigit")
	}
	panic(fmt.Sprintf("Cannot invoke 'isdigit' with 0 arguments on %v", fn))
}

type I_INVOKE_0_isdigit interface{ M_0_isdigit() M }

func F_INVOKE_0_islower(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_islower:
		return x.M_0_islower()
	case I_GET_islower:
		tmp := x.GET_islower()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("islower")
	}
	panic(fmt.Sprintf("Cannot invoke 'islower' with 0 arguments on %v", fn))
}

type I_INVOKE_0_islower interface{ M_0_islower() M }

func F_INVOKE_0_isspace(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_isspace:
		return x.M_0_isspace()
	case I_GET_isspace:
		tmp := x.GET_isspace()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("isspace")
	}
	panic(fmt.Sprintf("Cannot invoke 'isspace' with 0 arguments on %v", fn))
}

type I_INVOKE_0_isspace interface{ M_0_isspace() M }

func F_INVOKE_0_isupper(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_isupper:
		return x.M_0_isupper()
	case I_GET_isupper:
		tmp := x.GET_isupper()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("isupper")
	}
	panic(fmt.Sprintf("Cannot invoke 'isupper' with 0 arguments on %v", fn))
}

type I_INVOKE_0_isupper interface{ M_0_isupper() M }

func F_INVOKE_0_items(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_items:
		return x.M_0_items()
	case I_GET_items:
		tmp := x.GET_items()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("items")
	}
	panic(fmt.Sprintf("Cannot invoke 'items' with 0 arguments on %v", fn))
}

type I_INVOKE_0_items interface{ M_0_items() M }

func F_INVOKE_0_iter(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_iter:
		return x.M_0_iter()
	case I_GET_iter:
		tmp := x.GET_iter()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("iter")
	}
	panic(fmt.Sprintf("Cannot invoke 'iter' with 0 arguments on %v", fn))
}

type I_INVOKE_0_iter interface{ M_0_iter() M }

func F_INVOKE_0_iteritems(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_iteritems:
		return x.M_0_iteritems()
	case I_GET_iteritems:
		tmp := x.GET_iteritems()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("iteritems")
	}
	panic(fmt.Sprintf("Cannot invoke 'iteritems' with 0 arguments on %v", fn))
}

type I_INVOKE_0_iteritems interface{ M_0_iteritems() M }

func F_INVOKE_0_iterkeys(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_iterkeys:
		return x.M_0_iterkeys()
	case I_GET_iterkeys:
		tmp := x.GET_iterkeys()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("iterkeys")
	}
	panic(fmt.Sprintf("Cannot invoke 'iterkeys' with 0 arguments on %v", fn))
}

type I_INVOKE_0_iterkeys interface{ M_0_iterkeys() M }

func F_INVOKE_0_itervalues(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_itervalues:
		return x.M_0_itervalues()
	case I_GET_itervalues:
		tmp := x.GET_itervalues()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("itervalues")
	}
	panic(fmt.Sprintf("Cannot invoke 'itervalues' with 0 arguments on %v", fn))
}

type I_INVOKE_0_itervalues interface{ M_0_itervalues() M }

func F_INVOKE_0_keys(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_keys:
		return x.M_0_keys()
	case I_GET_keys:
		tmp := x.GET_keys()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("keys")
	}
	panic(fmt.Sprintf("Cannot invoke 'keys' with 0 arguments on %v", fn))
}

type I_INVOKE_0_keys interface{ M_0_keys() M }

func F_INVOKE_0_lower(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_lower:
		return x.M_0_lower()
	case I_GET_lower:
		tmp := x.GET_lower()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("lower")
	}
	panic(fmt.Sprintf("Cannot invoke 'lower' with 0 arguments on %v", fn))
}

type I_INVOKE_0_lower interface{ M_0_lower() M }

func F_INVOKE_0_read(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_read:
		return x.M_0_read()
	case I_GET_read:
		tmp := x.GET_read()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("read")
	}
	panic(fmt.Sprintf("Cannot invoke 'read' with 0 arguments on %v", fn))
}

type I_INVOKE_0_read interface{ M_0_read() M }

func F_INVOKE_0_reverse(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_reverse:
		return x.M_0_reverse()
	case I_GET_reverse:
		tmp := x.GET_reverse()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("reverse")
	}
	panic(fmt.Sprintf("Cannot invoke 'reverse' with 0 arguments on %v", fn))
}

type I_INVOKE_0_reverse interface{ M_0_reverse() M }

func F_INVOKE_0_title(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_title:
		return x.M_0_title()
	case I_GET_title:
		tmp := x.GET_title()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("title")
	}
	panic(fmt.Sprintf("Cannot invoke 'title' with 0 arguments on %v", fn))
}

type I_INVOKE_0_title interface{ M_0_title() M }

func F_INVOKE_0_upper(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_upper:
		return x.M_0_upper()
	case I_GET_upper:
		tmp := x.GET_upper()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("upper")
	}
	panic(fmt.Sprintf("Cannot invoke 'upper' with 0 arguments on %v", fn))
}

type I_INVOKE_0_upper interface{ M_0_upper() M }

func F_INVOKE_0_values(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_0_values:
		return x.M_0_values()
	case I_GET_values:
		tmp := x.GET_values()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("values")
	}
	panic(fmt.Sprintf("Cannot invoke 'values' with 0 arguments on %v", fn))
}

type I_INVOKE_0_values interface{ M_0_values() M }

func F_INVOKE_1_Raise(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_Raise:
		return x.M_1_Raise(a_0)
	case I_GET_Raise:
		tmp := x.GET_Raise()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Raise", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Raise' with 1 arguments on %v", fn))
}

type I_INVOKE_1_Raise interface{ M_1_Raise(a_0 M) M }

func F_INVOKE_1_Recv(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_Recv:
		return x.M_1_Recv(a_0)
	case I_GET_Recv:
		tmp := x.GET_Recv()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Recv", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Recv' with 1 arguments on %v", fn))
}

type I_INVOKE_1_Recv interface{ M_1_Recv(a_0 M) M }

func F_INVOKE_1_Send(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_Send:
		return x.M_1_Send(a_0)
	case I_GET_Send:
		tmp := x.GET_Send()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Send", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Send' with 1 arguments on %v", fn))
}

type I_INVOKE_1_Send interface{ M_1_Send(a_0 M) M }

func F_INVOKE_1_Start(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_Start:
		return x.M_1_Start(a_0)
	case I_GET_Start:
		tmp := x.GET_Start()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Start", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Start' with 1 arguments on %v", fn))
}

type I_INVOKE_1_Start interface{ M_1_Start(a_0 M) M }

func F_INVOKE_1_Write(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_Write:
		return x.M_1_Write(a_0)
	case I_GET_Write:
		tmp := x.GET_Write()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Write", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Write' with 1 arguments on %v", fn))
}

type I_INVOKE_1_Write interface{ M_1_Write(a_0 M) M }

func F_INVOKE_1___getattr__(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1___getattr__:
		return x.M_1___getattr__(a_0)
	case I_GET___getattr__:
		tmp := x.GET___getattr__()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("__getattr__", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke '__getattr__' with 1 arguments on %v", fn))
}

type I_INVOKE_1___getattr__ interface{ M_1___getattr__(a_0 M) M }

func F_INVOKE_1_append(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_append:
		return x.M_1_append(a_0)
	case I_GET_append:
		tmp := x.GET_append()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("append", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'append' with 1 arguments on %v", fn))
}

type I_INVOKE_1_append interface{ M_1_append(a_0 M) M }

func F_INVOKE_1_count(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_count:
		return x.M_1_count(a_0)
	case I_GET_count:
		tmp := x.GET_count()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("count", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'count' with 1 arguments on %v", fn))
}

type I_INVOKE_1_count interface{ M_1_count(a_0 M) M }

func F_INVOKE_1_endswith(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_endswith:
		return x.M_1_endswith(a_0)
	case I_GET_endswith:
		tmp := x.GET_endswith()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("endswith", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'endswith' with 1 arguments on %v", fn))
}

type I_INVOKE_1_endswith interface{ M_1_endswith(a_0 M) M }

func F_INVOKE_1_extend(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_extend:
		return x.M_1_extend(a_0)
	case I_GET_extend:
		tmp := x.GET_extend()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("extend", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'extend' with 1 arguments on %v", fn))
}

type I_INVOKE_1_extend interface{ M_1_extend(a_0 M) M }

func F_INVOKE_1_find(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_find:
		return x.M_1_find(a_0)
	case I_GET_find:
		tmp := x.GET_find()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("find", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'find' with 1 arguments on %v", fn))
}

type I_INVOKE_1_find interface{ M_1_find(a_0 M) M }

func F_INVOKE_1_has_key(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_has_key:
		return x.M_1_has_key(a_0)
	case I_GET_has_key:
		tmp := x.GET_has_key()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("has_key", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'has_key' with 1 arguments on %v", fn))
}

type I_INVOKE_1_has_key interface{ M_1_has_key(a_0 M) M }

func F_INVOKE_1_index(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_index:
		return x.M_1_index(a_0)
	case I_GET_index:
		tmp := x.GET_index()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("index", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'index' with 1 arguments on %v", fn))
}

type I_INVOKE_1_index interface{ M_1_index(a_0 M) M }

func F_INVOKE_1_join(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_join:
		return x.M_1_join(a_0)
	case I_GET_join:
		tmp := x.GET_join()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("join", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'join' with 1 arguments on %v", fn))
}

type I_INVOKE_1_join interface{ M_1_join(a_0 M) M }

func F_INVOKE_1_lstrip(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_lstrip:
		return x.M_1_lstrip(a_0)
	case I_GET_lstrip:
		tmp := x.GET_lstrip()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("lstrip", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'lstrip' with 1 arguments on %v", fn))
}

type I_INVOKE_1_lstrip interface{ M_1_lstrip(a_0 M) M }

func F_INVOKE_1_pop(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_pop:
		return x.M_1_pop(a_0)
	case I_GET_pop:
		tmp := x.GET_pop()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("pop", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'pop' with 1 arguments on %v", fn))
}

type I_INVOKE_1_pop interface{ M_1_pop(a_0 M) M }

func F_INVOKE_1_remove(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_remove:
		return x.M_1_remove(a_0)
	case I_GET_remove:
		tmp := x.GET_remove()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("remove", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'remove' with 1 arguments on %v", fn))
}

type I_INVOKE_1_remove interface{ M_1_remove(a_0 M) M }

func F_INVOKE_1_rfind(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_rfind:
		return x.M_1_rfind(a_0)
	case I_GET_rfind:
		tmp := x.GET_rfind()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("rfind", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'rfind' with 1 arguments on %v", fn))
}

type I_INVOKE_1_rfind interface{ M_1_rfind(a_0 M) M }

func F_INVOKE_1_rindex(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_rindex:
		return x.M_1_rindex(a_0)
	case I_GET_rindex:
		tmp := x.GET_rindex()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("rindex", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'rindex' with 1 arguments on %v", fn))
}

type I_INVOKE_1_rindex interface{ M_1_rindex(a_0 M) M }

func F_INVOKE_1_rstrip(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_rstrip:
		return x.M_1_rstrip(a_0)
	case I_GET_rstrip:
		tmp := x.GET_rstrip()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("rstrip", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'rstrip' with 1 arguments on %v", fn))
}

type I_INVOKE_1_rstrip interface{ M_1_rstrip(a_0 M) M }

func F_INVOKE_1_startswith(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_startswith:
		return x.M_1_startswith(a_0)
	case I_GET_startswith:
		tmp := x.GET_startswith()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("startswith", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'startswith' with 1 arguments on %v", fn))
}

type I_INVOKE_1_startswith interface{ M_1_startswith(a_0 M) M }

func F_INVOKE_1_strip(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_strip:
		return x.M_1_strip(a_0)
	case I_GET_strip:
		tmp := x.GET_strip()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("strip", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'strip' with 1 arguments on %v", fn))
}

type I_INVOKE_1_strip interface{ M_1_strip(a_0 M) M }

func F_INVOKE_1_update(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_update:
		return x.M_1_update(a_0)
	case I_GET_update:
		tmp := x.GET_update()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("update", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'update' with 1 arguments on %v", fn))
}

type I_INVOKE_1_update interface{ M_1_update(a_0 M) M }

func F_INVOKE_1_write(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_1_write:
		return x.M_1_write(a_0)
	case I_GET_write:
		tmp := x.GET_write()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("write", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'write' with 1 arguments on %v", fn))
}

type I_INVOKE_1_write interface{ M_1_write(a_0 M) M }

func F_INVOKE_2___init__(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2___init__:
		return x.M_2___init__(a_0, a_1)
	case I_GET___init__:
		tmp := x.GET___init__()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("__init__", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke '__init__' with 2 arguments on %v", fn))
}

type I_INVOKE_2___init__ interface{ M_2___init__(a_0 M, a_1 M) M }

func F_INVOKE_2___setattr__(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2___setattr__:
		return x.M_2___setattr__(a_0, a_1)
	case I_GET___setattr__:
		tmp := x.GET___setattr__()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("__setattr__", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke '__setattr__' with 2 arguments on %v", fn))
}

type I_INVOKE_2___setattr__ interface{ M_2___setattr__(a_0 M, a_1 M) M }

func F_INVOKE_2_get(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2_get:
		return x.M_2_get(a_0, a_1)
	case I_GET_get:
		tmp := x.GET_get()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("get", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke 'get' with 2 arguments on %v", fn))
}

type I_INVOKE_2_get interface{ M_2_get(a_0 M, a_1 M) M }

func F_INVOKE_2_insert(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2_insert:
		return x.M_2_insert(a_0, a_1)
	case I_GET_insert:
		tmp := x.GET_insert()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("insert", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke 'insert' with 2 arguments on %v", fn))
}

type I_INVOKE_2_insert interface{ M_2_insert(a_0 M, a_1 M) M }

func F_INVOKE_2_setdefault(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2_setdefault:
		return x.M_2_setdefault(a_0, a_1)
	case I_GET_setdefault:
		tmp := x.GET_setdefault()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("setdefault", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke 'setdefault' with 2 arguments on %v", fn))
}

type I_INVOKE_2_setdefault interface{ M_2_setdefault(a_0 M, a_1 M) M }

func F_INVOKE_2_split(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_2_split:
		return x.M_2_split(a_0, a_1)
	case I_GET_split:
		tmp := x.GET_split()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("split", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke 'split' with 2 arguments on %v", fn))
}

type I_INVOKE_2_split interface{ M_2_split(a_0 M, a_1 M) M }

func F_INVOKE_3_replace(fn M, a_0 M, a_1 M, a_2 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_3_replace:
		return x.M_3_replace(a_0, a_1, a_2)
	case I_GET_replace:
		tmp := x.GET_replace()
		return CALL_3(tmp, a_0, a_1, a_2)

	case *PGo:
		return x.Invoke("replace", a_0, a_1, a_2)
	}
	panic(fmt.Sprintf("Cannot invoke 'replace' with 3 arguments on %v", fn))
}

type I_INVOKE_3_replace interface{ M_3_replace(a_0 M, a_1 M, a_2 M) M }

func F_INVOKE_3_sort(fn M, a_0 M, a_1 M, a_2 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case I_INVOKE_3_sort:
		return x.M_3_sort(a_0, a_1, a_2)
	case I_GET_sort:
		tmp := x.GET_sort()
		return CALL_3(tmp, a_0, a_1, a_2)

	case *PGo:
		return x.Invoke("sort", a_0, a_1, a_2)
	}
	panic(fmt.Sprintf("Cannot invoke 'sort' with 3 arguments on %v", fn))
}

type I_INVOKE_3_sort interface{ M_3_sort(a_0 M, a_1 M, a_2 M) M }

type I_GET_Close interface{ GET_Close() M }

func F_GET_Close(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Close:
		return x.GET_Close()
	}
	return h.FetchField("Close")
}

type I_GET_Flush interface{ GET_Flush() M }

func F_GET_Flush(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Flush:
		return x.GET_Flush()
	}
	return h.FetchField("Flush")
}

type I_GET_Raise interface{ GET_Raise() M }

func F_GET_Raise(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Raise:
		return x.GET_Raise()
	}
	return h.FetchField("Raise")
}

type I_GET_Recv interface{ GET_Recv() M }

func F_GET_Recv(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Recv:
		return x.GET_Recv()
	}
	return h.FetchField("Recv")
}

type I_GET_Send interface{ GET_Send() M }

func F_GET_Send(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Send:
		return x.GET_Send()
	}
	return h.FetchField("Send")
}

type I_GET_Start interface{ GET_Start() M }

func F_GET_Start(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Start:
		return x.GET_Start()
	}
	return h.FetchField("Start")
}

type I_GET_Wait interface{ GET_Wait() M }

func F_GET_Wait(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Wait:
		return x.GET_Wait()
	}
	return h.FetchField("Wait")
}

type I_GET_Warm interface{ GET_Warm() M }

func F_GET_Warm(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Warm:
		return x.GET_Warm()
	}
	return h.FetchField("Warm")
}

type I_GET_Write interface{ GET_Write() M }

func F_GET_Write(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_Write:
		return x.GET_Write()
	}
	return h.FetchField("Write")
}

type I_GET___getattr__ interface{ GET___getattr__() M }

func F_GET___getattr__(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET___getattr__:
		return x.GET___getattr__()
	}
	return h.FetchField("__getattr__")
}

type I_GET___init__ interface{ GET___init__() M }

func F_GET___init__(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET___init__:
		return x.GET___init__()
	}
	return h.FetchField("__init__")
}

type I_GET___setattr__ interface{ GET___setattr__() M }

func F_GET___setattr__(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET___setattr__:
		return x.GET___setattr__()
	}
	return h.FetchField("__setattr__")
}

type I_GET_append interface{ GET_append() M }

func F_GET_append(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_append:
		return x.GET_append()
	}
	return h.FetchField("append")
}

type I_GET_clear interface{ GET_clear() M }

func F_GET_clear(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_clear:
		return x.GET_clear()
	}
	return h.FetchField("clear")
}

type I_GET_close interface{ GET_close() M }

func F_GET_close(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_close:
		return x.GET_close()
	}
	return h.FetchField("close")
}

type I_GET_copy interface{ GET_copy() M }

func F_GET_copy(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_copy:
		return x.GET_copy()
	}
	return h.FetchField("copy")
}

type I_GET_count interface{ GET_count() M }

func F_GET_count(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_count:
		return x.GET_count()
	}
	return h.FetchField("count")
}

type I_GET_endswith interface{ GET_endswith() M }

func F_GET_endswith(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_endswith:
		return x.GET_endswith()
	}
	return h.FetchField("endswith")
}

type I_GET_extend interface{ GET_extend() M }

func F_GET_extend(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_extend:
		return x.GET_extend()
	}
	return h.FetchField("extend")
}

type I_GET_find interface{ GET_find() M }

func F_GET_find(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_find:
		return x.GET_find()
	}
	return h.FetchField("find")
}

type I_GET_flush interface{ GET_flush() M }

func F_GET_flush(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_flush:
		return x.GET_flush()
	}
	return h.FetchField("flush")
}

type I_GET_get interface{ GET_get() M }

func F_GET_get(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_get:
		return x.GET_get()
	}
	return h.FetchField("get")
}

type I_GET_has_key interface{ GET_has_key() M }

func F_GET_has_key(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_has_key:
		return x.GET_has_key()
	}
	return h.FetchField("has_key")
}

type I_GET_index interface{ GET_index() M }

func F_GET_index(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_index:
		return x.GET_index()
	}
	return h.FetchField("index")
}

type I_GET_insert interface{ GET_insert() M }

func F_GET_insert(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_insert:
		return x.GET_insert()
	}
	return h.FetchField("insert")
}

type I_GET_isalnum interface{ GET_isalnum() M }

func F_GET_isalnum(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_isalnum:
		return x.GET_isalnum()
	}
	return h.FetchField("isalnum")
}

type I_GET_isalpha interface{ GET_isalpha() M }

func F_GET_isalpha(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_isalpha:
		return x.GET_isalpha()
	}
	return h.FetchField("isalpha")
}

type I_GET_isdigit interface{ GET_isdigit() M }

func F_GET_isdigit(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_isdigit:
		return x.GET_isdigit()
	}
	return h.FetchField("isdigit")
}

type I_GET_islower interface{ GET_islower() M }

func F_GET_islower(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_islower:
		return x.GET_islower()
	}
	return h.FetchField("islower")
}

type I_GET_isspace interface{ GET_isspace() M }

func F_GET_isspace(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_isspace:
		return x.GET_isspace()
	}
	return h.FetchField("isspace")
}

type I_GET_isupper interface{ GET_isupper() M }

func F_GET_isupper(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_isupper:
		return x.GET_isupper()
	}
	return h.FetchField("isupper")
}

type I_GET_items interface{ GET_items() M }

func F_GET_items(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_items:
		return x.GET_items()
	}
	return h.FetchField("items")
}

type I_GET_iter interface{ GET_iter() M }

func F_GET_iter(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_iter:
		return x.GET_iter()
	}
	return h.FetchField("iter")
}

type I_GET_iteritems interface{ GET_iteritems() M }

func F_GET_iteritems(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_iteritems:
		return x.GET_iteritems()
	}
	return h.FetchField("iteritems")
}

type I_GET_iterkeys interface{ GET_iterkeys() M }

func F_GET_iterkeys(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_iterkeys:
		return x.GET_iterkeys()
	}
	return h.FetchField("iterkeys")
}

type I_GET_itervalues interface{ GET_itervalues() M }

func F_GET_itervalues(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_itervalues:
		return x.GET_itervalues()
	}
	return h.FetchField("itervalues")
}

type I_GET_join interface{ GET_join() M }

func F_GET_join(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_join:
		return x.GET_join()
	}
	return h.FetchField("join")
}

type I_GET_keys interface{ GET_keys() M }

func F_GET_keys(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_keys:
		return x.GET_keys()
	}
	return h.FetchField("keys")
}

type I_GET_lower interface{ GET_lower() M }

func F_GET_lower(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_lower:
		return x.GET_lower()
	}
	return h.FetchField("lower")
}

type I_GET_lstrip interface{ GET_lstrip() M }

func F_GET_lstrip(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_lstrip:
		return x.GET_lstrip()
	}
	return h.FetchField("lstrip")
}

type I_GET_pop interface{ GET_pop() M }

func F_GET_pop(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_pop:
		return x.GET_pop()
	}
	return h.FetchField("pop")
}

type I_GET_read interface{ GET_read() M }

func F_GET_read(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_read:
		return x.GET_read()
	}
	return h.FetchField("read")
}

type I_GET_remove interface{ GET_remove() M }

func F_GET_remove(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_remove:
		return x.GET_remove()
	}
	return h.FetchField("remove")
}

type I_GET_replace interface{ GET_replace() M }

func F_GET_replace(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_replace:
		return x.GET_replace()
	}
	return h.FetchField("replace")
}

type I_GET_reverse interface{ GET_reverse() M }

func F_GET_reverse(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_reverse:
		return x.GET_reverse()
	}
	return h.FetchField("reverse")
}

type I_GET_rfind interface{ GET_rfind() M }

func F_GET_rfind(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_rfind:
		return x.GET_rfind()
	}
	return h.FetchField("rfind")
}

type I_GET_rindex interface{ GET_rindex() M }

func F_GET_rindex(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_rindex:
		return x.GET_rindex()
	}
	return h.FetchField("rindex")
}

type I_GET_rstrip interface{ GET_rstrip() M }

func F_GET_rstrip(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_rstrip:
		return x.GET_rstrip()
	}
	return h.FetchField("rstrip")
}

type I_GET_setdefault interface{ GET_setdefault() M }

func F_GET_setdefault(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_setdefault:
		return x.GET_setdefault()
	}
	return h.FetchField("setdefault")
}

type I_GET_sort interface{ GET_sort() M }

func F_GET_sort(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_sort:
		return x.GET_sort()
	}
	return h.FetchField("sort")
}

type I_GET_split interface{ GET_split() M }

func F_GET_split(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_split:
		return x.GET_split()
	}
	return h.FetchField("split")
}

type I_GET_startswith interface{ GET_startswith() M }

func F_GET_startswith(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_startswith:
		return x.GET_startswith()
	}
	return h.FetchField("startswith")
}

type I_GET_strip interface{ GET_strip() M }

func F_GET_strip(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_strip:
		return x.GET_strip()
	}
	return h.FetchField("strip")
}

type I_GET_title interface{ GET_title() M }

func F_GET_title(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_title:
		return x.GET_title()
	}
	return h.FetchField("title")
}

type I_GET_update interface{ GET_update() M }

func F_GET_update(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_update:
		return x.GET_update()
	}
	return h.FetchField("update")
}

type I_GET_upper interface{ GET_upper() M }

func F_GET_upper(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_upper:
		return x.GET_upper()
	}
	return h.FetchField("upper")
}

type I_GET_values interface{ GET_values() M }

func F_GET_values(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_values:
		return x.GET_values()
	}
	return h.FetchField("values")
}

type I_GET_write interface{ GET_write() M }

func F_GET_write(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case I_GET_write:
		return x.GET_write()
	}
	return h.FetchField("write")
}

type I_0 interface{ Call0() M }

func CALL_0(fn M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_0:
		return f.Call0()
	case ICallV:
		return f.CallV([]M{}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_1 interface{ Call1(M) M }

func CALL_1(fn M, a_0 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_1:
		return f.Call1(a_0)
	case ICallV:
		return f.CallV([]M{a_0}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_2 interface{ Call2(M, M) M }

func CALL_2(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_2:
		return f.Call2(a_0, a_1)
	case ICallV:
		return f.CallV([]M{a_0, a_1}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_3 interface{ Call3(M, M, M) M }

func CALL_3(fn M, a_0 M, a_1 M, a_2 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_3:
		return f.Call3(a_0, a_1, a_2)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_4 interface{ Call4(M, M, M, M) M }

func CALL_4(fn M, a_0 M, a_1 M, a_2 M, a_3 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_4:
		return f.Call4(a_0, a_1, a_2, a_3)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_5 interface{ Call5(M, M, M, M, M) M }

func CALL_5(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_5:
		return f.Call5(a_0, a_1, a_2, a_3, a_4)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_6 interface{ Call6(M, M, M, M, M, M) M }

func CALL_6(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M, a_5 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_6:
		return f.Call6(a_0, a_1, a_2, a_3, a_4, a_5)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4, a_5}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_7 interface{ Call7(M, M, M, M, M, M, M) M }

func CALL_7(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M, a_5 M, a_6 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_7:
		return f.Call7(a_0, a_1, a_2, a_3, a_4, a_5, a_6)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4, a_5, a_6}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_8 interface {
	Call8(M, M, M, M, M, M, M, M) M
}

func CALL_8(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M, a_5 M, a_6 M, a_7 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_8:
		return f.Call8(a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_9 interface {
	Call9(M, M, M, M, M, M, M, M, M) M
}

func CALL_9(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M, a_5 M, a_6 M, a_7 M, a_8 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_9:
		return f.Call9(a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7, a_8)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7, a_8}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

type I_10 interface {
	Call10(M, M, M, M, M, M, M, M, M, M) M
}

func CALL_10(fn M, a_0 M, a_1 M, a_2 M, a_3 M, a_4 M, a_5 M, a_6 M, a_7 M, a_8 M, a_9 M) M {
	if fn.X == nil {
		panic("cannot CALL on int or str")
	}
	switch f := fn.X.(type) {
	case I_10:
		return f.Call10(a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7, a_8, a_9)
	case ICallV:
		return f.CallV([]M{a_0, a_1, a_2, a_3, a_4, a_5, a_6, a_7, a_8, a_9}, nil, nil, nil)
	}
	panic(fmt.Sprintf("No way to call: %v", fn))
}

// self.signatures.items: {'signature_Write___5b_5duint8_return_int_also_error': 'Write ([]uint8) (int, error)', 'signature_Flush___return_error': 'Flush () (error)', 'signature_Close___return_error': 'Close () (error)'}
type signature_Close___return_error interface{ Close() error }                                   // self.signatures.items
type signature_Flush___return_error interface{ Flush() error }                                   // self.signatures.items
type signature_Write___5b_5duint8_return_int_also_error interface{ Write([]uint8) (int, error) } // self.signatures.items

//ydefs// Exception => Yfunc:(BUILTINS.Exception) [[ <codegen.Yfunc object at 0x7f963cb10c10> ]]
//
//ydefs// PYE_FileDesc => Yfunc:(BUILTINS.PYE_FileDesc) [[ <codegen.Yfunc object at 0x7f963cb1ee90> ]]
//
//ydefs// _rye__force_generation_of_call_4_ => Yfunc:(BUILTINS._rye__force_generation_of_call_4_) [[ <codegen.Yfunc object at 0x7f963cb1ec10> ]]
//
//ydefs// all => Yfunc:(BUILTINS.all) [[ <codegen.Yfunc object at 0x7f963cb16090> ]]
//
//ydefs// any => Yfunc:(BUILTINS.any) [[ <codegen.Yfunc object at 0x7f963cb16050> ]]
//
//ydefs// bool => Yfunc:(BUILTINS.bool) [[ <codegen.Yfunc object at 0x7f963cb163d0> ]]
//
//ydefs// byt => Yfunc:(BUILTINS.byt) [[ <codegen.Yfunc object at 0x7f963cb16450> ]]
//
//ydefs// callable => Yfunc:(BUILTINS.callable) [[ <codegen.Yfunc object at 0x7f963cb10d10> ]]
//
//ydefs// chr => Yfunc:(BUILTINS.chr) [[ <codegen.Yfunc object at 0x7f963cb10f90> ]]
//
//ydefs// cmp => Yfunc:(BUILTINS.cmp) [[ <codegen.Yfunc object at 0x7f963cb10e10> ]]
//
//ydefs// dict => Yfunc:(BUILTINS.dict) [[ <codegen.Yfunc object at 0x7f963cb16350> ]]
//
//ydefs// float => Yfunc:(BUILTINS.float) [[ <codegen.Yfunc object at 0x7f963cb161d0> ]]
//
//ydefs// getattr => Yfunc:(BUILTINS.getattr) [[ <codegen.Yfunc object at 0x7f963cb10e50> ]]
//
//ydefs// go_deref => Yfunc:(BUILTINS.go_deref) [[ <codegen.Yfunc object at 0x7f963cb10c90> ]]
//
//ydefs// hash => Yfunc:(BUILTINS.hash) [[ <codegen.Yfunc object at 0x7f963cb10dd0> ]]
//
//ydefs// id => Yfunc:(BUILTINS.id) [[ <codegen.Yfunc object at 0x7f963cb10d90> ]]
//
//ydefs// int => Yfunc:(BUILTINS.int) [[ <codegen.Yfunc object at 0x7f963cb16190> ]]
//
//ydefs// isinstance => Yfunc:(BUILTINS.isinstance) [[ <codegen.Yfunc object at 0x7f963cb10ed0> ]]
//
//ydefs// issubclass => Yfunc:(BUILTINS.issubclass) [[ <codegen.Yfunc object at 0x7f963cb10f10> ]]
//
//ydefs// len => Yfunc:(BUILTINS.len) [[ <codegen.Yfunc object at 0x7f963cb160d0> ]]
//
//ydefs// list => Yfunc:(BUILTINS.list) [[ <codegen.Yfunc object at 0x7f963cb162d0> ]]
//
//ydefs// map => Yfunc:(BUILTINS.map) [[ <codegen.Yfunc object at 0x7f963cb166d0> ]]
//
//ydefs// max => Yfunc:(BUILTINS.max) [[ <codegen.Yfunc object at 0x7f963cb165d0> ]]
//
//ydefs// min => Yfunc:(BUILTINS.min) [[ <codegen.Yfunc object at 0x7f963cb16610> ]]
//
//ydefs// mkbyt => Yfunc:(BUILTINS.mkbyt) [[ <codegen.Yfunc object at 0x7f963cb16490> ]]
//
//ydefs// object => Yfunc:(BUILTINS.object) [[ <codegen.Yfunc object at 0x7f963cb16750> ]]
//
//ydefs// open => Yfunc:(BUILTINS.open) [[ <codegen.Yfunc object at 0x7f963cb1e850> ]]
//
//ydefs// ord => Yfunc:(BUILTINS.ord) [[ <codegen.Yfunc object at 0x7f963cb10f50> ]]
//
//ydefs// range => Yfunc:(BUILTINS.range) [[ <codegen.Yfunc object at 0x7f963cb16210> ]]
//
//ydefs// reduce => Yfunc:(BUILTINS.reduce) [[ <codegen.Yfunc object at 0x7f963cb16710> ]]
//
//ydefs// repr => Yfunc:(BUILTINS.repr) [[ <codegen.Yfunc object at 0x7f963cb16110> ]]
//
//ydefs// rye_chan => Yfunc:(BUILTINS.rye_chan) [[ <codegen.Yfunc object at 0x7f963cb1e710> ]]
//
//ydefs// rye_opts => Yfunc:(BUILTINS.rye_opts) [[ <codegen.Yfunc object at 0x7f963cb10c50> ]]
//
//ydefs// rye_pickle => Yfunc:(BUILTINS.rye_pickle) [[ <codegen.Yfunc object at 0x7f963cb16510> ]]
//
//ydefs// rye_stack => Yfunc:(BUILTINS.rye_stack) [[ <codegen.Yfunc object at 0x7f963cb164d0> ]]
//
//ydefs// rye_unpickle => Yfunc:(BUILTINS.rye_unpickle) [[ <codegen.Yfunc object at 0x7f963cb16590> ]]
//
//ydefs// rye_what => Yfunc:(BUILTINS.rye_what) [[ <codegen.Yfunc object at 0x7f963cb10cd0> ]]
//
//ydefs// rye_zip_padding_with_None => Yfunc:(BUILTINS.rye_zip_padding_with_None) [[ <codegen.Yfunc object at 0x7f963cb16690> ]]
//
//ydefs// set => Yfunc:(BUILTINS.set) [[ <codegen.Yfunc object at 0x7f963cb16310> ]]
//
//ydefs// setattr => Yfunc:(BUILTINS.setattr) [[ <codegen.Yfunc object at 0x7f963cb10e90> ]]
//
//ydefs// setattrs => Yfunc:(BUILTINS.setattrs) [[ <codegen.Yfunc object at 0x7f963cb10d50> ]]
//
//ydefs// sorted => Yfunc:(BUILTINS.sorted) [[ <codegen.Yfunc object at 0x7f963cb16290> ]]
//
//ydefs// str => Yfunc:(BUILTINS.str) [[ <codegen.Yfunc object at 0x7f963cb16150> ]]
//
//ydefs// sum => Yfunc:(BUILTINS.sum) [[ <codegen.Yfunc object at 0x7f963cb10fd0> ]]
//
//ydefs// tuple => Yfunc:(BUILTINS.tuple) [[ <codegen.Yfunc object at 0x7f963cb16390> ]]
//
//ydefs// type => Yfunc:(BUILTINS.type) [[ <codegen.Yfunc object at 0x7f963cb16410> ]]
//
//ydefs// xrange => Yfunc:(BUILTINS.xrange) [[ <codegen.Yfunc object at 0x7f963cb16250> ]]
//
//ydefs// zip => Yfunc:(BUILTINS.zip) [[ <codegen.Yfunc object at 0x7f963cb16650> ]]
//
//
//ymeth// C_channel => [ ['Close', 'Raise', 'Recv', 'Send', 'Start', 'Wait', 'Warm'] ]
//ydefs// C_channel => Close => Yfunc:(BUILTINS.C_channel::Close) [[ <codegen.Yfunc object at 0x7f963cb1ebd0> ]]
//ydefs// C_channel => Raise => Yfunc:(BUILTINS.C_channel::Raise) [[ <codegen.Yfunc object at 0x7f963cb1e9d0> ]]
//ydefs// C_channel => Recv => Yfunc:(BUILTINS.C_channel::Recv) [[ <codegen.Yfunc object at 0x7f963cb1ead0> ]]
//ydefs// C_channel => Send => Yfunc:(BUILTINS.C_channel::Send) [[ <codegen.Yfunc object at 0x7f963cb1ea50> ]]
//ydefs// C_channel => Start => Yfunc:(BUILTINS.C_channel::Start) [[ <codegen.Yfunc object at 0x7f963cb1e950> ]]
//ydefs// C_channel => Wait => Yfunc:(BUILTINS.C_channel::Wait) [[ <codegen.Yfunc object at 0x7f963cb1eb50> ]]
//ydefs// C_channel => Warm => Yfunc:(BUILTINS.C_channel::Warm) [[ <codegen.Yfunc object at 0x7f963cb1e8d0> ]]
//
//ymeth// C_object => [ ['__getattr__', '__init__', '__setattr__'] ]
//ydefs// C_object => __getattr__ => Yfunc:(BUILTINS.C_object::__getattr__) [[ <codegen.Yfunc object at 0x7f963cb1e7d0> ]]
//ydefs// C_object => __init__ => Yfunc:(BUILTINS.C_object::__init__) [[ <codegen.Yfunc object at 0x7f963cb1e790> ]]
//ydefs// C_object => __setattr__ => Yfunc:(BUILTINS.C_object::__setattr__) [[ <codegen.Yfunc object at 0x7f963cb1e810> ]]
//
//ymeth// PByt => [ ['endswith', 'find', 'index', 'isalnum', 'isalpha', 'isdigit', 'islower', 'isspace', 'isupper', 'join', 'lower', 'lstrip', 'replace', 'rfind', 'rindex', 'rstrip', 'split', 'startswith', 'strip', 'title', 'upper'] ]
//ydefs// PByt => endswith => Yfunc:(BUILTINS.PByt::endswith) [[ <codegen.Yfunc object at 0x7f963cb1af50> ]]
//ydefs// PByt => find => Yfunc:(BUILTINS.PByt::find) [[ <codegen.Yfunc object at 0x7f963cb1e250> ]]
//ydefs// PByt => index => Yfunc:(BUILTINS.PByt::index) [[ <codegen.Yfunc object at 0x7f963cb1e350> ]]
//ydefs// PByt => isalnum => Yfunc:(BUILTINS.PByt::isalnum) [[ <codegen.Yfunc object at 0x7f963cb1e550> ]]
//ydefs// PByt => isalpha => Yfunc:(BUILTINS.PByt::isalpha) [[ <codegen.Yfunc object at 0x7f963cb1e450> ]]
//ydefs// PByt => isdigit => Yfunc:(BUILTINS.PByt::isdigit) [[ <codegen.Yfunc object at 0x7f963cb1e4d0> ]]
//ydefs// PByt => islower => Yfunc:(BUILTINS.PByt::islower) [[ <codegen.Yfunc object at 0x7f963cb1e5d0> ]]
//ydefs// PByt => isspace => Yfunc:(BUILTINS.PByt::isspace) [[ <codegen.Yfunc object at 0x7f963cb1e6d0> ]]
//ydefs// PByt => isupper => Yfunc:(BUILTINS.PByt::isupper) [[ <codegen.Yfunc object at 0x7f963cb1e650> ]]
//ydefs// PByt => join => Yfunc:(BUILTINS.PByt::join) [[ <codegen.Yfunc object at 0x7f963cb1ad50> ]]
//ydefs// PByt => lower => Yfunc:(BUILTINS.PByt::lower) [[ <codegen.Yfunc object at 0x7f963cb1add0> ]]
//ydefs// PByt => lstrip => Yfunc:(BUILTINS.PByt::lstrip) [[ <codegen.Yfunc object at 0x7f963cb1e110> ]]
//ydefs// PByt => replace => Yfunc:(BUILTINS.PByt::replace) [[ <codegen.Yfunc object at 0x7f963cb1e210> ]]
//ydefs// PByt => rfind => Yfunc:(BUILTINS.PByt::rfind) [[ <codegen.Yfunc object at 0x7f963cb1e2d0> ]]
//ydefs// PByt => rindex => Yfunc:(BUILTINS.PByt::rindex) [[ <codegen.Yfunc object at 0x7f963cb1e3d0> ]]
//ydefs// PByt => rstrip => Yfunc:(BUILTINS.PByt::rstrip) [[ <codegen.Yfunc object at 0x7f963cb1e190> ]]
//ydefs// PByt => split => Yfunc:(BUILTINS.PByt::split) [[ <codegen.Yfunc object at 0x7f963cb1ad10> ]]
//ydefs// PByt => startswith => Yfunc:(BUILTINS.PByt::startswith) [[ <codegen.Yfunc object at 0x7f963cb1afd0> ]]
//ydefs// PByt => strip => Yfunc:(BUILTINS.PByt::strip) [[ <codegen.Yfunc object at 0x7f963cb1e090> ]]
//ydefs// PByt => title => Yfunc:(BUILTINS.PByt::title) [[ <codegen.Yfunc object at 0x7f963cb1ae50> ]]
//ydefs// PByt => upper => Yfunc:(BUILTINS.PByt::upper) [[ <codegen.Yfunc object at 0x7f963cb1aed0> ]]
//
//ymeth// PDict => [ ['clear', 'copy', 'get', 'has_key', 'items', 'iter', 'iteritems', 'iterkeys', 'itervalues', 'keys', 'setdefault', 'update', 'values'] ]
//ydefs// PDict => clear => Yfunc:(BUILTINS.PDict::clear) [[ <codegen.Yfunc object at 0x7f963cb16c90> ]]
//ydefs// PDict => copy => Yfunc:(BUILTINS.PDict::copy) [[ <codegen.Yfunc object at 0x7f963cb16d10> ]]
//ydefs// PDict => get => Yfunc:(BUILTINS.PDict::get) [[ <codegen.Yfunc object at 0x7f963cb1a110> ]]
//ydefs// PDict => has_key => Yfunc:(BUILTINS.PDict::has_key) [[ <codegen.Yfunc object at 0x7f963cb1a190> ]]
//ydefs// PDict => items => Yfunc:(BUILTINS.PDict::items) [[ <codegen.Yfunc object at 0x7f963cb16d90> ]]
//ydefs// PDict => iter => Yfunc:(BUILTINS.PDict::iter) [[ <codegen.Yfunc object at 0x7f963cb16f90> ]]
//ydefs// PDict => iteritems => Yfunc:(BUILTINS.PDict::iteritems) [[ <codegen.Yfunc object at 0x7f963cb16e10> ]]
//ydefs// PDict => iterkeys => Yfunc:(BUILTINS.PDict::iterkeys) [[ <codegen.Yfunc object at 0x7f963cb16f10> ]]
//ydefs// PDict => itervalues => Yfunc:(BUILTINS.PDict::itervalues) [[ <codegen.Yfunc object at 0x7f963cb1a0d0> ]]
//ydefs// PDict => keys => Yfunc:(BUILTINS.PDict::keys) [[ <codegen.Yfunc object at 0x7f963cb16e90> ]]
//ydefs// PDict => setdefault => Yfunc:(BUILTINS.PDict::setdefault) [[ <codegen.Yfunc object at 0x7f963cb1a210> ]]
//ydefs// PDict => update => Yfunc:(BUILTINS.PDict::update) [[ <codegen.Yfunc object at 0x7f963cb1a290> ]]
//ydefs// PDict => values => Yfunc:(BUILTINS.PDict::values) [[ <codegen.Yfunc object at 0x7f963cb1a050> ]]
//
//ymeth// PList => [ ['append', 'copy', 'count', 'extend', 'index', 'insert', 'pop', 'remove', 'reverse', 'sort'] ]
//ydefs// PList => append => Yfunc:(BUILTINS.PList::append) [[ <codegen.Yfunc object at 0x7f963cb167d0> ]]
//ydefs// PList => copy => Yfunc:(BUILTINS.PList::copy) [[ <codegen.Yfunc object at 0x7f963cb16c10> ]]
//ydefs// PList => count => Yfunc:(BUILTINS.PList::count) [[ <codegen.Yfunc object at 0x7f963cb168d0> ]]
//ydefs// PList => extend => Yfunc:(BUILTINS.PList::extend) [[ <codegen.Yfunc object at 0x7f963cb16850> ]]
//ydefs// PList => index => Yfunc:(BUILTINS.PList::index) [[ <codegen.Yfunc object at 0x7f963cb16950> ]]
//ydefs// PList => insert => Yfunc:(BUILTINS.PList::insert) [[ <codegen.Yfunc object at 0x7f963cb16a50> ]]
//ydefs// PList => pop => Yfunc:(BUILTINS.PList::pop) [[ <codegen.Yfunc object at 0x7f963cb16a90> ]]
//ydefs// PList => remove => Yfunc:(BUILTINS.PList::remove) [[ <codegen.Yfunc object at 0x7f963cb169d0> ]]
//ydefs// PList => reverse => Yfunc:(BUILTINS.PList::reverse) [[ <codegen.Yfunc object at 0x7f963cb16b10> ]]
//ydefs// PList => sort => Yfunc:(BUILTINS.PList::sort) [[ <codegen.Yfunc object at 0x7f963cb16b90> ]]
//
//ymeth// PStr => [ ['endswith', 'find', 'index', 'isalnum', 'isalpha', 'isdigit', 'islower', 'isspace', 'isupper', 'join', 'lower', 'lstrip', 'replace', 'rfind', 'rindex', 'rstrip', 'split', 'startswith', 'strip', 'title', 'upper'] ]
//ydefs// PStr => endswith => Yfunc:(BUILTINS.PStr::endswith) [[ <codegen.Yfunc object at 0x7f963cb1a550> ]]
//ydefs// PStr => find => Yfunc:(BUILTINS.PStr::find) [[ <codegen.Yfunc object at 0x7f963cb1a810> ]]
//ydefs// PStr => index => Yfunc:(BUILTINS.PStr::index) [[ <codegen.Yfunc object at 0x7f963cb1a910> ]]
//ydefs// PStr => isalnum => Yfunc:(BUILTINS.PStr::isalnum) [[ <codegen.Yfunc object at 0x7f963cb1ab10> ]]
//ydefs// PStr => isalpha => Yfunc:(BUILTINS.PStr::isalpha) [[ <codegen.Yfunc object at 0x7f963cb1aa10> ]]
//ydefs// PStr => isdigit => Yfunc:(BUILTINS.PStr::isdigit) [[ <codegen.Yfunc object at 0x7f963cb1aa90> ]]
//ydefs// PStr => islower => Yfunc:(BUILTINS.PStr::islower) [[ <codegen.Yfunc object at 0x7f963cb1ab90> ]]
//ydefs// PStr => isspace => Yfunc:(BUILTINS.PStr::isspace) [[ <codegen.Yfunc object at 0x7f963cb1ac90> ]]
//ydefs// PStr => isupper => Yfunc:(BUILTINS.PStr::isupper) [[ <codegen.Yfunc object at 0x7f963cb1ac10> ]]
//ydefs// PStr => join => Yfunc:(BUILTINS.PStr::join) [[ <codegen.Yfunc object at 0x7f963cb1a350> ]]
//ydefs// PStr => lower => Yfunc:(BUILTINS.PStr::lower) [[ <codegen.Yfunc object at 0x7f963cb1a3d0> ]]
//ydefs// PStr => lstrip => Yfunc:(BUILTINS.PStr::lstrip) [[ <codegen.Yfunc object at 0x7f963cb1a6d0> ]]
//ydefs// PStr => replace => Yfunc:(BUILTINS.PStr::replace) [[ <codegen.Yfunc object at 0x7f963cb1a7d0> ]]
//ydefs// PStr => rfind => Yfunc:(BUILTINS.PStr::rfind) [[ <codegen.Yfunc object at 0x7f963cb1a890> ]]
//ydefs// PStr => rindex => Yfunc:(BUILTINS.PStr::rindex) [[ <codegen.Yfunc object at 0x7f963cb1a990> ]]
//ydefs// PStr => rstrip => Yfunc:(BUILTINS.PStr::rstrip) [[ <codegen.Yfunc object at 0x7f963cb1a750> ]]
//ydefs// PStr => split => Yfunc:(BUILTINS.PStr::split) [[ <codegen.Yfunc object at 0x7f963cb1a310> ]]
//ydefs// PStr => startswith => Yfunc:(BUILTINS.PStr::startswith) [[ <codegen.Yfunc object at 0x7f963cb1a5d0> ]]
//ydefs// PStr => strip => Yfunc:(BUILTINS.PStr::strip) [[ <codegen.Yfunc object at 0x7f963cb1a650> ]]
//ydefs// PStr => title => Yfunc:(BUILTINS.PStr::title) [[ <codegen.Yfunc object at 0x7f963cb1a450> ]]
//ydefs// PStr => upper => Yfunc:(BUILTINS.PStr::upper) [[ <codegen.Yfunc object at 0x7f963cb1a4d0> ]]
//
//ymeth// PYE_FileDesc => [ ['Close', 'Flush', '__init__', 'close', 'flush', 'read', 'write'] ]
//ydefs// PYE_FileDesc => Close => Yfunc:(BUILTINS.PYE_FileDesc::Close) [[ <codegen.Yfunc object at 0x7f963cb1ee10> ]]
//ydefs// PYE_FileDesc => Flush => Yfunc:(BUILTINS.PYE_FileDesc::Flush) [[ <codegen.Yfunc object at 0x7f963cb1ed90> ]]
//ydefs// PYE_FileDesc => __init__ => Yfunc:(BUILTINS.PYE_FileDesc::__init__) [[ <codegen.Yfunc object at 0x7f963cb1ec50> ]]
//ydefs// PYE_FileDesc => close => Yfunc:(BUILTINS.PYE_FileDesc::close) [[ <codegen.Yfunc object at 0x7f963cb1edd0> ]]
//ydefs// PYE_FileDesc => flush => Yfunc:(BUILTINS.PYE_FileDesc::flush) [[ <codegen.Yfunc object at 0x7f963cb1ed50> ]]
//ydefs// PYE_FileDesc => read => Yfunc:(BUILTINS.PYE_FileDesc::read) [[ <codegen.Yfunc object at 0x7f963cb1ecd0> ]]
//ydefs// PYE_FileDesc => write => Yfunc:(BUILTINS.PYE_FileDesc::write) [[ <codegen.Yfunc object at 0x7f963cb1ed10> ]]
//
//
