// Adapted from ../../yak-labs/chirp-lang/goapi/default/wrap.go
// Changes:
//    -- Change package to runt
//    -- Delete import chirp-lang
// TODO: automate generating this file.

package runt

// import . `github.com/yak-labs/chirp-lang`
import (
	bufio `bufio`
	bytes `bytes`
	encoding_base64 `encoding/base64`
	fmt `fmt`
	io_ioutil `io/ioutil`
	math `math`
	math_big `math/big`
	net_http `net/http`
	os `os`
	reflect `reflect`
	regexp `regexp`
	strconv `strconv`
	strings `strings`
	time `time`
)

func init() {
	Roots[`/bufio/NewReadWriter`] = FuncRoot{ Func: reflect.ValueOf(bufio.NewReadWriter) }
	Roots[`/bufio/NewReader`] = FuncRoot{ Func: reflect.ValueOf(bufio.NewReader) }
	Roots[`/bufio/NewReaderSize`] = FuncRoot{ Func: reflect.ValueOf(bufio.NewReaderSize) }
	Roots[`/bufio/NewWriter`] = FuncRoot{ Func: reflect.ValueOf(bufio.NewWriter) }
	Roots[`/bufio/NewWriterSize`] = FuncRoot{ Func: reflect.ValueOf(bufio.NewWriterSize) }
	Roots[`/bytes/Compare`] = FuncRoot{ Func: reflect.ValueOf(bytes.Compare) }
	Roots[`/bytes/Contains`] = FuncRoot{ Func: reflect.ValueOf(bytes.Contains) }
	Roots[`/bytes/Count`] = FuncRoot{ Func: reflect.ValueOf(bytes.Count) }
	Roots[`/bytes/Equal`] = FuncRoot{ Func: reflect.ValueOf(bytes.Equal) }
	Roots[`/bytes/EqualFold`] = FuncRoot{ Func: reflect.ValueOf(bytes.EqualFold) }
	Roots[`/bytes/Fields`] = FuncRoot{ Func: reflect.ValueOf(bytes.Fields) }
	Roots[`/bytes/FieldsFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.FieldsFunc) }
	Roots[`/bytes/HasPrefix`] = FuncRoot{ Func: reflect.ValueOf(bytes.HasPrefix) }
	Roots[`/bytes/HasSuffix`] = FuncRoot{ Func: reflect.ValueOf(bytes.HasSuffix) }
	Roots[`/bytes/Index`] = FuncRoot{ Func: reflect.ValueOf(bytes.Index) }
	Roots[`/bytes/IndexAny`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexAny) }
	Roots[`/bytes/IndexByte`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexByte) }
	Roots[`/bytes/IndexFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexFunc) }
	Roots[`/bytes/IndexRune`] = FuncRoot{ Func: reflect.ValueOf(bytes.IndexRune) }
	Roots[`/bytes/Join`] = FuncRoot{ Func: reflect.ValueOf(bytes.Join) }
	Roots[`/bytes/LastIndex`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndex) }
	Roots[`/bytes/LastIndexAny`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndexAny) }
	Roots[`/bytes/LastIndexFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.LastIndexFunc) }
	Roots[`/bytes/Map`] = FuncRoot{ Func: reflect.ValueOf(bytes.Map) }
	Roots[`/bytes/NewBuffer`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewBuffer) }
	Roots[`/bytes/NewBufferString`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewBufferString) }
	Roots[`/bytes/NewReader`] = FuncRoot{ Func: reflect.ValueOf(bytes.NewReader) }
	Roots[`/bytes/Repeat`] = FuncRoot{ Func: reflect.ValueOf(bytes.Repeat) }
	Roots[`/bytes/Replace`] = FuncRoot{ Func: reflect.ValueOf(bytes.Replace) }
	Roots[`/bytes/Runes`] = FuncRoot{ Func: reflect.ValueOf(bytes.Runes) }
	Roots[`/bytes/Split`] = FuncRoot{ Func: reflect.ValueOf(bytes.Split) }
	Roots[`/bytes/SplitAfter`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitAfter) }
	Roots[`/bytes/SplitAfterN`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitAfterN) }
	Roots[`/bytes/SplitN`] = FuncRoot{ Func: reflect.ValueOf(bytes.SplitN) }
	Roots[`/bytes/Title`] = FuncRoot{ Func: reflect.ValueOf(bytes.Title) }
	Roots[`/bytes/ToLower`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToLower) }
	Roots[`/bytes/ToLowerSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToLowerSpecial) }
	Roots[`/bytes/ToTitle`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToTitle) }
	Roots[`/bytes/ToTitleSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToTitleSpecial) }
	Roots[`/bytes/ToUpper`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToUpper) }
	Roots[`/bytes/ToUpperSpecial`] = FuncRoot{ Func: reflect.ValueOf(bytes.ToUpperSpecial) }
	Roots[`/bytes/Trim`] = FuncRoot{ Func: reflect.ValueOf(bytes.Trim) }
	Roots[`/bytes/TrimFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimFunc) }
	Roots[`/bytes/TrimLeft`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimLeft) }
	Roots[`/bytes/TrimLeftFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimLeftFunc) }
	Roots[`/bytes/TrimRight`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimRight) }
	Roots[`/bytes/TrimRightFunc`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimRightFunc) }
	Roots[`/bytes/TrimSpace`] = FuncRoot{ Func: reflect.ValueOf(bytes.TrimSpace) }
	Roots[`/encoding/base64/NewDecoder`] = FuncRoot{ Func: reflect.ValueOf(encoding_base64.NewDecoder) }
	Roots[`/encoding/base64/NewEncoder`] = FuncRoot{ Func: reflect.ValueOf(encoding_base64.NewEncoder) }
	Roots[`/encoding/base64/NewEncoding`] = FuncRoot{ Func: reflect.ValueOf(encoding_base64.NewEncoding) }
	Roots[`/fmt/Errorf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Errorf) }
	Roots[`/fmt/Fprint`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprint) }
	Roots[`/fmt/Fprintf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprintf) }
	Roots[`/fmt/Fprintln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fprintln) }
	Roots[`/fmt/Fscan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscan) }
	Roots[`/fmt/Fscanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscanf) }
	Roots[`/fmt/Fscanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Fscanln) }
	Roots[`/fmt/Print`] = FuncRoot{ Func: reflect.ValueOf(fmt.Print) }
	Roots[`/fmt/Printf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Printf) }
	Roots[`/fmt/Println`] = FuncRoot{ Func: reflect.ValueOf(fmt.Println) }
	Roots[`/fmt/Scan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scan) }
	Roots[`/fmt/Scanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scanf) }
	Roots[`/fmt/Scanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Scanln) }
	Roots[`/fmt/Sprint`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprint) }
	Roots[`/fmt/Sprintf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprintf) }
	Roots[`/fmt/Sprintln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sprintln) }
	Roots[`/fmt/Sscan`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscan) }
	Roots[`/fmt/Sscanf`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscanf) }
	Roots[`/fmt/Sscanln`] = FuncRoot{ Func: reflect.ValueOf(fmt.Sscanln) }
	Roots[`/io/ioutil/NopCloser`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.NopCloser) }
	Roots[`/io/ioutil/ReadAll`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.ReadAll) }
	Roots[`/io/ioutil/ReadDir`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.ReadDir) }
	Roots[`/io/ioutil/ReadFile`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.ReadFile) }
	Roots[`/io/ioutil/TempDir`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.TempDir) }
	Roots[`/io/ioutil/TempFile`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.TempFile) }
	Roots[`/io/ioutil/WriteFile`] = FuncRoot{ Func: reflect.ValueOf(io_ioutil.WriteFile) }
	Roots[`/math/Abs`] = FuncRoot{ Func: reflect.ValueOf(math.Abs) }
	Roots[`/math/Acos`] = FuncRoot{ Func: reflect.ValueOf(math.Acos) }
	Roots[`/math/Acosh`] = FuncRoot{ Func: reflect.ValueOf(math.Acosh) }
	Roots[`/math/Asin`] = FuncRoot{ Func: reflect.ValueOf(math.Asin) }
	Roots[`/math/Asinh`] = FuncRoot{ Func: reflect.ValueOf(math.Asinh) }
	Roots[`/math/Atan`] = FuncRoot{ Func: reflect.ValueOf(math.Atan) }
	Roots[`/math/Atan2`] = FuncRoot{ Func: reflect.ValueOf(math.Atan2) }
	Roots[`/math/Atanh`] = FuncRoot{ Func: reflect.ValueOf(math.Atanh) }
	Roots[`/math/Cbrt`] = FuncRoot{ Func: reflect.ValueOf(math.Cbrt) }
	Roots[`/math/Ceil`] = FuncRoot{ Func: reflect.ValueOf(math.Ceil) }
	Roots[`/math/Copysign`] = FuncRoot{ Func: reflect.ValueOf(math.Copysign) }
	Roots[`/math/Cos`] = FuncRoot{ Func: reflect.ValueOf(math.Cos) }
	Roots[`/math/Cosh`] = FuncRoot{ Func: reflect.ValueOf(math.Cosh) }
	Roots[`/math/Dim`] = FuncRoot{ Func: reflect.ValueOf(math.Dim) }
	Roots[`/math/Erf`] = FuncRoot{ Func: reflect.ValueOf(math.Erf) }
	Roots[`/math/Erfc`] = FuncRoot{ Func: reflect.ValueOf(math.Erfc) }
	Roots[`/math/Exp`] = FuncRoot{ Func: reflect.ValueOf(math.Exp) }
	Roots[`/math/Exp2`] = FuncRoot{ Func: reflect.ValueOf(math.Exp2) }
	Roots[`/math/Expm1`] = FuncRoot{ Func: reflect.ValueOf(math.Expm1) }
	Roots[`/math/Float32bits`] = FuncRoot{ Func: reflect.ValueOf(math.Float32bits) }
	Roots[`/math/Float32frombits`] = FuncRoot{ Func: reflect.ValueOf(math.Float32frombits) }
	Roots[`/math/Float64bits`] = FuncRoot{ Func: reflect.ValueOf(math.Float64bits) }
	Roots[`/math/Float64frombits`] = FuncRoot{ Func: reflect.ValueOf(math.Float64frombits) }
	Roots[`/math/Floor`] = FuncRoot{ Func: reflect.ValueOf(math.Floor) }
	Roots[`/math/Frexp`] = FuncRoot{ Func: reflect.ValueOf(math.Frexp) }
	Roots[`/math/Gamma`] = FuncRoot{ Func: reflect.ValueOf(math.Gamma) }
	Roots[`/math/Hypot`] = FuncRoot{ Func: reflect.ValueOf(math.Hypot) }
	Roots[`/math/Ilogb`] = FuncRoot{ Func: reflect.ValueOf(math.Ilogb) }
	Roots[`/math/Inf`] = FuncRoot{ Func: reflect.ValueOf(math.Inf) }
	Roots[`/math/IsInf`] = FuncRoot{ Func: reflect.ValueOf(math.IsInf) }
	Roots[`/math/IsNaN`] = FuncRoot{ Func: reflect.ValueOf(math.IsNaN) }
	Roots[`/math/J0`] = FuncRoot{ Func: reflect.ValueOf(math.J0) }
	Roots[`/math/J1`] = FuncRoot{ Func: reflect.ValueOf(math.J1) }
	Roots[`/math/Jn`] = FuncRoot{ Func: reflect.ValueOf(math.Jn) }
	Roots[`/math/Ldexp`] = FuncRoot{ Func: reflect.ValueOf(math.Ldexp) }
	Roots[`/math/Lgamma`] = FuncRoot{ Func: reflect.ValueOf(math.Lgamma) }
	Roots[`/math/Log`] = FuncRoot{ Func: reflect.ValueOf(math.Log) }
	Roots[`/math/Log10`] = FuncRoot{ Func: reflect.ValueOf(math.Log10) }
	Roots[`/math/Log1p`] = FuncRoot{ Func: reflect.ValueOf(math.Log1p) }
	Roots[`/math/Log2`] = FuncRoot{ Func: reflect.ValueOf(math.Log2) }
	Roots[`/math/Logb`] = FuncRoot{ Func: reflect.ValueOf(math.Logb) }
	Roots[`/math/Max`] = FuncRoot{ Func: reflect.ValueOf(math.Max) }
	Roots[`/math/Min`] = FuncRoot{ Func: reflect.ValueOf(math.Min) }
	Roots[`/math/Mod`] = FuncRoot{ Func: reflect.ValueOf(math.Mod) }
	Roots[`/math/Modf`] = FuncRoot{ Func: reflect.ValueOf(math.Modf) }
	Roots[`/math/NaN`] = FuncRoot{ Func: reflect.ValueOf(math.NaN) }
	Roots[`/math/Nextafter`] = FuncRoot{ Func: reflect.ValueOf(math.Nextafter) }
	Roots[`/math/Pow`] = FuncRoot{ Func: reflect.ValueOf(math.Pow) }
	Roots[`/math/Pow10`] = FuncRoot{ Func: reflect.ValueOf(math.Pow10) }
	Roots[`/math/Remainder`] = FuncRoot{ Func: reflect.ValueOf(math.Remainder) }
	Roots[`/math/Signbit`] = FuncRoot{ Func: reflect.ValueOf(math.Signbit) }
	Roots[`/math/Sin`] = FuncRoot{ Func: reflect.ValueOf(math.Sin) }
	Roots[`/math/Sincos`] = FuncRoot{ Func: reflect.ValueOf(math.Sincos) }
	Roots[`/math/Sinh`] = FuncRoot{ Func: reflect.ValueOf(math.Sinh) }
	Roots[`/math/Sqrt`] = FuncRoot{ Func: reflect.ValueOf(math.Sqrt) }
	Roots[`/math/Tan`] = FuncRoot{ Func: reflect.ValueOf(math.Tan) }
	Roots[`/math/Tanh`] = FuncRoot{ Func: reflect.ValueOf(math.Tanh) }
	Roots[`/math/Trunc`] = FuncRoot{ Func: reflect.ValueOf(math.Trunc) }
	Roots[`/math/Y0`] = FuncRoot{ Func: reflect.ValueOf(math.Y0) }
	Roots[`/math/Y1`] = FuncRoot{ Func: reflect.ValueOf(math.Y1) }
	Roots[`/math/Yn`] = FuncRoot{ Func: reflect.ValueOf(math.Yn) }
	Roots[`/math/big/NewInt`] = FuncRoot{ Func: reflect.ValueOf(math_big.NewInt) }
	Roots[`/math/big/NewRat`] = FuncRoot{ Func: reflect.ValueOf(math_big.NewRat) }
	Roots[`/net/http/CanonicalHeaderKey`] = FuncRoot{ Func: reflect.ValueOf(net_http.CanonicalHeaderKey) }
	Roots[`/net/http/DetectContentType`] = FuncRoot{ Func: reflect.ValueOf(net_http.DetectContentType) }
	Roots[`/net/http/Error`] = FuncRoot{ Func: reflect.ValueOf(net_http.Error) }
	Roots[`/net/http/FileServer`] = FuncRoot{ Func: reflect.ValueOf(net_http.FileServer) }
	Roots[`/net/http/Get`] = FuncRoot{ Func: reflect.ValueOf(net_http.Get) }
	Roots[`/net/http/Handle`] = FuncRoot{ Func: reflect.ValueOf(net_http.Handle) }
	Roots[`/net/http/HandleFunc`] = FuncRoot{ Func: reflect.ValueOf(net_http.HandleFunc) }
	Roots[`/net/http/Head`] = FuncRoot{ Func: reflect.ValueOf(net_http.Head) }
	Roots[`/net/http/ListenAndServe`] = FuncRoot{ Func: reflect.ValueOf(net_http.ListenAndServe) }
	Roots[`/net/http/ListenAndServeTLS`] = FuncRoot{ Func: reflect.ValueOf(net_http.ListenAndServeTLS) }
	Roots[`/net/http/MaxBytesReader`] = FuncRoot{ Func: reflect.ValueOf(net_http.MaxBytesReader) }
	Roots[`/net/http/NewFileTransport`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewFileTransport) }
	Roots[`/net/http/NewRequest`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewRequest) }
	Roots[`/net/http/NewServeMux`] = FuncRoot{ Func: reflect.ValueOf(net_http.NewServeMux) }
	Roots[`/net/http/NotFound`] = FuncRoot{ Func: reflect.ValueOf(net_http.NotFound) }
	Roots[`/net/http/NotFoundHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.NotFoundHandler) }
	Roots[`/net/http/ParseHTTPVersion`] = FuncRoot{ Func: reflect.ValueOf(net_http.ParseHTTPVersion) }
	Roots[`/net/http/Post`] = FuncRoot{ Func: reflect.ValueOf(net_http.Post) }
	Roots[`/net/http/PostForm`] = FuncRoot{ Func: reflect.ValueOf(net_http.PostForm) }
	Roots[`/net/http/ProxyFromEnvironment`] = FuncRoot{ Func: reflect.ValueOf(net_http.ProxyFromEnvironment) }
	Roots[`/net/http/ProxyURL`] = FuncRoot{ Func: reflect.ValueOf(net_http.ProxyURL) }
	Roots[`/net/http/ReadRequest`] = FuncRoot{ Func: reflect.ValueOf(net_http.ReadRequest) }
	Roots[`/net/http/ReadResponse`] = FuncRoot{ Func: reflect.ValueOf(net_http.ReadResponse) }
	Roots[`/net/http/Redirect`] = FuncRoot{ Func: reflect.ValueOf(net_http.Redirect) }
	Roots[`/net/http/RedirectHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.RedirectHandler) }
	Roots[`/net/http/Serve`] = FuncRoot{ Func: reflect.ValueOf(net_http.Serve) }
	Roots[`/net/http/ServeContent`] = FuncRoot{ Func: reflect.ValueOf(net_http.ServeContent) }
	Roots[`/net/http/ServeFile`] = FuncRoot{ Func: reflect.ValueOf(net_http.ServeFile) }
	Roots[`/net/http/SetCookie`] = FuncRoot{ Func: reflect.ValueOf(net_http.SetCookie) }
	Roots[`/net/http/StatusText`] = FuncRoot{ Func: reflect.ValueOf(net_http.StatusText) }
	Roots[`/net/http/StripPrefix`] = FuncRoot{ Func: reflect.ValueOf(net_http.StripPrefix) }
	Roots[`/net/http/TimeoutHandler`] = FuncRoot{ Func: reflect.ValueOf(net_http.TimeoutHandler) }
	Roots[`/os/Chdir`] = FuncRoot{ Func: reflect.ValueOf(os.Chdir) }
	Roots[`/os/Chmod`] = FuncRoot{ Func: reflect.ValueOf(os.Chmod) }
	Roots[`/os/Chown`] = FuncRoot{ Func: reflect.ValueOf(os.Chown) }
	Roots[`/os/Chtimes`] = FuncRoot{ Func: reflect.ValueOf(os.Chtimes) }
	Roots[`/os/Clearenv`] = FuncRoot{ Func: reflect.ValueOf(os.Clearenv) }
	Roots[`/os/Create`] = FuncRoot{ Func: reflect.ValueOf(os.Create) }
	Roots[`/os/Environ`] = FuncRoot{ Func: reflect.ValueOf(os.Environ) }
	Roots[`/os/Exit`] = FuncRoot{ Func: reflect.ValueOf(os.Exit) }
	Roots[`/os/Expand`] = FuncRoot{ Func: reflect.ValueOf(os.Expand) }
	Roots[`/os/ExpandEnv`] = FuncRoot{ Func: reflect.ValueOf(os.ExpandEnv) }
	Roots[`/os/FindProcess`] = FuncRoot{ Func: reflect.ValueOf(os.FindProcess) }
	Roots[`/os/Getegid`] = FuncRoot{ Func: reflect.ValueOf(os.Getegid) }
	Roots[`/os/Getenv`] = FuncRoot{ Func: reflect.ValueOf(os.Getenv) }
	Roots[`/os/Geteuid`] = FuncRoot{ Func: reflect.ValueOf(os.Geteuid) }
	Roots[`/os/Getgid`] = FuncRoot{ Func: reflect.ValueOf(os.Getgid) }
	Roots[`/os/Getgroups`] = FuncRoot{ Func: reflect.ValueOf(os.Getgroups) }
	Roots[`/os/Getpagesize`] = FuncRoot{ Func: reflect.ValueOf(os.Getpagesize) }
	Roots[`/os/Getpid`] = FuncRoot{ Func: reflect.ValueOf(os.Getpid) }
	Roots[`/os/Getppid`] = FuncRoot{ Func: reflect.ValueOf(os.Getppid) }
	Roots[`/os/Getuid`] = FuncRoot{ Func: reflect.ValueOf(os.Getuid) }
	Roots[`/os/Getwd`] = FuncRoot{ Func: reflect.ValueOf(os.Getwd) }
	Roots[`/os/Hostname`] = FuncRoot{ Func: reflect.ValueOf(os.Hostname) }
	Roots[`/os/IsExist`] = FuncRoot{ Func: reflect.ValueOf(os.IsExist) }
	Roots[`/os/IsNotExist`] = FuncRoot{ Func: reflect.ValueOf(os.IsNotExist) }
	Roots[`/os/IsPathSeparator`] = FuncRoot{ Func: reflect.ValueOf(os.IsPathSeparator) }
	Roots[`/os/IsPermission`] = FuncRoot{ Func: reflect.ValueOf(os.IsPermission) }
	Roots[`/os/Lchown`] = FuncRoot{ Func: reflect.ValueOf(os.Lchown) }
	Roots[`/os/Link`] = FuncRoot{ Func: reflect.ValueOf(os.Link) }
	Roots[`/os/Lstat`] = FuncRoot{ Func: reflect.ValueOf(os.Lstat) }
	Roots[`/os/Mkdir`] = FuncRoot{ Func: reflect.ValueOf(os.Mkdir) }
	Roots[`/os/MkdirAll`] = FuncRoot{ Func: reflect.ValueOf(os.MkdirAll) }
	Roots[`/os/NewFile`] = FuncRoot{ Func: reflect.ValueOf(os.NewFile) }
	Roots[`/os/NewSyscallError`] = FuncRoot{ Func: reflect.ValueOf(os.NewSyscallError) }
	Roots[`/os/Open`] = FuncRoot{ Func: reflect.ValueOf(os.Open) }
	Roots[`/os/OpenFile`] = FuncRoot{ Func: reflect.ValueOf(os.OpenFile) }
	Roots[`/os/Pipe`] = FuncRoot{ Func: reflect.ValueOf(os.Pipe) }
	Roots[`/os/Readlink`] = FuncRoot{ Func: reflect.ValueOf(os.Readlink) }
	Roots[`/os/Remove`] = FuncRoot{ Func: reflect.ValueOf(os.Remove) }
	Roots[`/os/RemoveAll`] = FuncRoot{ Func: reflect.ValueOf(os.RemoveAll) }
	Roots[`/os/Rename`] = FuncRoot{ Func: reflect.ValueOf(os.Rename) }
	Roots[`/os/SameFile`] = FuncRoot{ Func: reflect.ValueOf(os.SameFile) }
	Roots[`/os/Setenv`] = FuncRoot{ Func: reflect.ValueOf(os.Setenv) }
	Roots[`/os/StartProcess`] = FuncRoot{ Func: reflect.ValueOf(os.StartProcess) }
	Roots[`/os/Stat`] = FuncRoot{ Func: reflect.ValueOf(os.Stat) }
	Roots[`/os/Symlink`] = FuncRoot{ Func: reflect.ValueOf(os.Symlink) }
	Roots[`/os/TempDir`] = FuncRoot{ Func: reflect.ValueOf(os.TempDir) }
	Roots[`/os/Truncate`] = FuncRoot{ Func: reflect.ValueOf(os.Truncate) }
	Roots[`/reflect/Append`] = FuncRoot{ Func: reflect.ValueOf(reflect.Append) }
	Roots[`/reflect/AppendSlice`] = FuncRoot{ Func: reflect.ValueOf(reflect.AppendSlice) }
	Roots[`/reflect/Copy`] = FuncRoot{ Func: reflect.ValueOf(reflect.Copy) }
	Roots[`/reflect/DeepEqual`] = FuncRoot{ Func: reflect.ValueOf(reflect.DeepEqual) }
	Roots[`/reflect/Indirect`] = FuncRoot{ Func: reflect.ValueOf(reflect.Indirect) }
	Roots[`/reflect/MakeChan`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeChan) }
	Roots[`/reflect/MakeMap`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeMap) }
	Roots[`/reflect/MakeSlice`] = FuncRoot{ Func: reflect.ValueOf(reflect.MakeSlice) }
	Roots[`/reflect/New`] = FuncRoot{ Func: reflect.ValueOf(reflect.New) }
	Roots[`/reflect/NewAt`] = FuncRoot{ Func: reflect.ValueOf(reflect.NewAt) }
	Roots[`/reflect/PtrTo`] = FuncRoot{ Func: reflect.ValueOf(reflect.PtrTo) }
	Roots[`/reflect/TypeOf`] = FuncRoot{ Func: reflect.ValueOf(reflect.TypeOf) }
	Roots[`/reflect/ValueOf`] = FuncRoot{ Func: reflect.ValueOf(reflect.ValueOf) }
	Roots[`/reflect/Zero`] = FuncRoot{ Func: reflect.ValueOf(reflect.Zero) }
	Roots[`/regexp/Compile`] = FuncRoot{ Func: reflect.ValueOf(regexp.Compile) }
	Roots[`/regexp/CompilePOSIX`] = FuncRoot{ Func: reflect.ValueOf(regexp.CompilePOSIX) }
	Roots[`/regexp/Match`] = FuncRoot{ Func: reflect.ValueOf(regexp.Match) }
	Roots[`/regexp/MatchReader`] = FuncRoot{ Func: reflect.ValueOf(regexp.MatchReader) }
	Roots[`/regexp/MatchString`] = FuncRoot{ Func: reflect.ValueOf(regexp.MatchString) }
	Roots[`/regexp/MustCompile`] = FuncRoot{ Func: reflect.ValueOf(regexp.MustCompile) }
	Roots[`/regexp/MustCompilePOSIX`] = FuncRoot{ Func: reflect.ValueOf(regexp.MustCompilePOSIX) }
	Roots[`/regexp/QuoteMeta`] = FuncRoot{ Func: reflect.ValueOf(regexp.QuoteMeta) }
	Roots[`/strconv/AppendBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendBool) }
	Roots[`/strconv/AppendFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendFloat) }
	Roots[`/strconv/AppendInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendInt) }
	Roots[`/strconv/AppendQuote`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuote) }
	Roots[`/strconv/AppendQuoteRune`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteRune) }
	Roots[`/strconv/AppendQuoteRuneToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteRuneToASCII) }
	Roots[`/strconv/AppendQuoteToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendQuoteToASCII) }
	Roots[`/strconv/AppendUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.AppendUint) }
	Roots[`/strconv/Atoi`] = FuncRoot{ Func: reflect.ValueOf(strconv.Atoi) }
	Roots[`/strconv/CanBackquote`] = FuncRoot{ Func: reflect.ValueOf(strconv.CanBackquote) }
	Roots[`/strconv/FormatBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatBool) }
	Roots[`/strconv/FormatFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatFloat) }
	Roots[`/strconv/FormatInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatInt) }
	Roots[`/strconv/FormatUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.FormatUint) }
	Roots[`/strconv/IsPrint`] = FuncRoot{ Func: reflect.ValueOf(strconv.IsPrint) }
	Roots[`/strconv/Itoa`] = FuncRoot{ Func: reflect.ValueOf(strconv.Itoa) }
	Roots[`/strconv/ParseBool`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseBool) }
	Roots[`/strconv/ParseFloat`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseFloat) }
	Roots[`/strconv/ParseInt`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseInt) }
	Roots[`/strconv/ParseUint`] = FuncRoot{ Func: reflect.ValueOf(strconv.ParseUint) }
	Roots[`/strconv/Quote`] = FuncRoot{ Func: reflect.ValueOf(strconv.Quote) }
	Roots[`/strconv/QuoteRune`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteRune) }
	Roots[`/strconv/QuoteRuneToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteRuneToASCII) }
	Roots[`/strconv/QuoteToASCII`] = FuncRoot{ Func: reflect.ValueOf(strconv.QuoteToASCII) }
	Roots[`/strconv/Unquote`] = FuncRoot{ Func: reflect.ValueOf(strconv.Unquote) }
	Roots[`/strconv/UnquoteChar`] = FuncRoot{ Func: reflect.ValueOf(strconv.UnquoteChar) }
	Roots[`/strings/Contains`] = FuncRoot{ Func: reflect.ValueOf(strings.Contains) }
	Roots[`/strings/ContainsAny`] = FuncRoot{ Func: reflect.ValueOf(strings.ContainsAny) }
	Roots[`/strings/ContainsRune`] = FuncRoot{ Func: reflect.ValueOf(strings.ContainsRune) }
	Roots[`/strings/Count`] = FuncRoot{ Func: reflect.ValueOf(strings.Count) }
	Roots[`/strings/EqualFold`] = FuncRoot{ Func: reflect.ValueOf(strings.EqualFold) }
	Roots[`/strings/Fields`] = FuncRoot{ Func: reflect.ValueOf(strings.Fields) }
	Roots[`/strings/FieldsFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.FieldsFunc) }
	Roots[`/strings/HasPrefix`] = FuncRoot{ Func: reflect.ValueOf(strings.HasPrefix) }
	Roots[`/strings/HasSuffix`] = FuncRoot{ Func: reflect.ValueOf(strings.HasSuffix) }
	Roots[`/strings/Index`] = FuncRoot{ Func: reflect.ValueOf(strings.Index) }
	Roots[`/strings/IndexAny`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexAny) }
	Roots[`/strings/IndexFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexFunc) }
	Roots[`/strings/IndexRune`] = FuncRoot{ Func: reflect.ValueOf(strings.IndexRune) }
	Roots[`/strings/Join`] = FuncRoot{ Func: reflect.ValueOf(strings.Join) }
	Roots[`/strings/LastIndex`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndex) }
	Roots[`/strings/LastIndexAny`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndexAny) }
	Roots[`/strings/LastIndexFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.LastIndexFunc) }
	Roots[`/strings/Map`] = FuncRoot{ Func: reflect.ValueOf(strings.Map) }
	Roots[`/strings/NewReader`] = FuncRoot{ Func: reflect.ValueOf(strings.NewReader) }
	Roots[`/strings/NewReplacer`] = FuncRoot{ Func: reflect.ValueOf(strings.NewReplacer) }
	Roots[`/strings/Repeat`] = FuncRoot{ Func: reflect.ValueOf(strings.Repeat) }
	Roots[`/strings/Replace`] = FuncRoot{ Func: reflect.ValueOf(strings.Replace) }
	Roots[`/strings/Split`] = FuncRoot{ Func: reflect.ValueOf(strings.Split) }
	Roots[`/strings/SplitAfter`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitAfter) }
	Roots[`/strings/SplitAfterN`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitAfterN) }
	Roots[`/strings/SplitN`] = FuncRoot{ Func: reflect.ValueOf(strings.SplitN) }
	Roots[`/strings/Title`] = FuncRoot{ Func: reflect.ValueOf(strings.Title) }
	Roots[`/strings/ToLower`] = FuncRoot{ Func: reflect.ValueOf(strings.ToLower) }
	Roots[`/strings/ToLowerSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToLowerSpecial) }
	Roots[`/strings/ToTitle`] = FuncRoot{ Func: reflect.ValueOf(strings.ToTitle) }
	Roots[`/strings/ToTitleSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToTitleSpecial) }
	Roots[`/strings/ToUpper`] = FuncRoot{ Func: reflect.ValueOf(strings.ToUpper) }
	Roots[`/strings/ToUpperSpecial`] = FuncRoot{ Func: reflect.ValueOf(strings.ToUpperSpecial) }
	Roots[`/strings/Trim`] = FuncRoot{ Func: reflect.ValueOf(strings.Trim) }
	Roots[`/strings/TrimFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimFunc) }
	Roots[`/strings/TrimLeft`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimLeft) }
	Roots[`/strings/TrimLeftFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimLeftFunc) }
	Roots[`/strings/TrimRight`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimRight) }
	Roots[`/strings/TrimRightFunc`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimRightFunc) }
	Roots[`/strings/TrimSpace`] = FuncRoot{ Func: reflect.ValueOf(strings.TrimSpace) }
	Roots[`/time/After`] = FuncRoot{ Func: reflect.ValueOf(time.After) }
	Roots[`/time/AfterFunc`] = FuncRoot{ Func: reflect.ValueOf(time.AfterFunc) }
	Roots[`/time/Date`] = FuncRoot{ Func: reflect.ValueOf(time.Date) }
	Roots[`/time/FixedZone`] = FuncRoot{ Func: reflect.ValueOf(time.FixedZone) }
	Roots[`/time/LoadLocation`] = FuncRoot{ Func: reflect.ValueOf(time.LoadLocation) }
	Roots[`/time/NewTicker`] = FuncRoot{ Func: reflect.ValueOf(time.NewTicker) }
	Roots[`/time/NewTimer`] = FuncRoot{ Func: reflect.ValueOf(time.NewTimer) }
	Roots[`/time/Now`] = FuncRoot{ Func: reflect.ValueOf(time.Now) }
	Roots[`/time/Parse`] = FuncRoot{ Func: reflect.ValueOf(time.Parse) }
	Roots[`/time/ParseDuration`] = FuncRoot{ Func: reflect.ValueOf(time.ParseDuration) }
	Roots[`/time/Since`] = FuncRoot{ Func: reflect.ValueOf(time.Since) }
	Roots[`/time/Sleep`] = FuncRoot{ Func: reflect.ValueOf(time.Sleep) }
	Roots[`/time/Tick`] = FuncRoot{ Func: reflect.ValueOf(time.Tick) }
	Roots[`/time/Unix`] = FuncRoot{ Func: reflect.ValueOf(time.Unix) }
	Roots[`/bufio/ErrBufferFull`] = VarRoot{ Var: reflect.ValueOf(&bufio.ErrBufferFull) }
	Roots[`/bufio/ErrInvalidUnreadByte`] = VarRoot{ Var: reflect.ValueOf(&bufio.ErrInvalidUnreadByte) }
	Roots[`/bufio/ErrInvalidUnreadRune`] = VarRoot{ Var: reflect.ValueOf(&bufio.ErrInvalidUnreadRune) }
	Roots[`/bufio/ErrNegativeCount`] = VarRoot{ Var: reflect.ValueOf(&bufio.ErrNegativeCount) }
	Roots[`/bytes/ErrTooLarge`] = VarRoot{ Var: reflect.ValueOf(&bytes.ErrTooLarge) }
	Roots[`/encoding/base64/StdEncoding`] = VarRoot{ Var: reflect.ValueOf(&encoding_base64.StdEncoding) }
	Roots[`/encoding/base64/URLEncoding`] = VarRoot{ Var: reflect.ValueOf(&encoding_base64.URLEncoding) }
	Roots[`/io/ioutil/Discard`] = VarRoot{ Var: reflect.ValueOf(&io_ioutil.Discard) }
	Roots[`/net/http/DefaultClient`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultClient) }
	Roots[`/net/http/DefaultServeMux`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultServeMux) }
	Roots[`/net/http/DefaultTransport`] = VarRoot{ Var: reflect.ValueOf(&net_http.DefaultTransport) }
	Roots[`/net/http/ErrBodyNotAllowed`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrBodyNotAllowed) }
	Roots[`/net/http/ErrBodyReadAfterClose`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrBodyReadAfterClose) }
	Roots[`/net/http/ErrContentLength`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrContentLength) }
	Roots[`/net/http/ErrHandlerTimeout`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHandlerTimeout) }
	Roots[`/net/http/ErrHeaderTooLong`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHeaderTooLong) }
	Roots[`/net/http/ErrHijacked`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrHijacked) }
	Roots[`/net/http/ErrLineTooLong`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrLineTooLong) }
	Roots[`/net/http/ErrMissingBoundary`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingBoundary) }
	Roots[`/net/http/ErrMissingContentLength`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingContentLength) }
	Roots[`/net/http/ErrMissingFile`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrMissingFile) }
	Roots[`/net/http/ErrNoCookie`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNoCookie) }
	Roots[`/net/http/ErrNoLocation`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNoLocation) }
	Roots[`/net/http/ErrNotMultipart`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNotMultipart) }
	Roots[`/net/http/ErrNotSupported`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrNotSupported) }
	Roots[`/net/http/ErrShortBody`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrShortBody) }
	Roots[`/net/http/ErrUnexpectedTrailer`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrUnexpectedTrailer) }
	Roots[`/net/http/ErrWriteAfterFlush`] = VarRoot{ Var: reflect.ValueOf(&net_http.ErrWriteAfterFlush) }
	Roots[`/os/Args`] = VarRoot{ Var: reflect.ValueOf(&os.Args) }
	Roots[`/os/ErrExist`] = VarRoot{ Var: reflect.ValueOf(&os.ErrExist) }
	Roots[`/os/ErrInvalid`] = VarRoot{ Var: reflect.ValueOf(&os.ErrInvalid) }
	Roots[`/os/ErrNotExist`] = VarRoot{ Var: reflect.ValueOf(&os.ErrNotExist) }
	Roots[`/os/ErrPermission`] = VarRoot{ Var: reflect.ValueOf(&os.ErrPermission) }
	Roots[`/os/Interrupt`] = VarRoot{ Var: reflect.ValueOf(&os.Interrupt) }
	Roots[`/os/Kill`] = VarRoot{ Var: reflect.ValueOf(&os.Kill) }
	Roots[`/os/Stderr`] = VarRoot{ Var: reflect.ValueOf(&os.Stderr) }
	Roots[`/os/Stdin`] = VarRoot{ Var: reflect.ValueOf(&os.Stdin) }
	Roots[`/os/Stdout`] = VarRoot{ Var: reflect.ValueOf(&os.Stdout) }
	Roots[`/strconv/ErrRange`] = VarRoot{ Var: reflect.ValueOf(&strconv.ErrRange) }
	Roots[`/strconv/ErrSyntax`] = VarRoot{ Var: reflect.ValueOf(&strconv.ErrSyntax) }
	Roots[`/time/Local`] = VarRoot{ Var: reflect.ValueOf(&time.Local) }
	Roots[`/time/UTC`] = VarRoot{ Var: reflect.ValueOf(&time.UTC) }
	{
	var tmp *bufio.ReadWriter
	Roots[`/bufio/ReadWriter`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bufio.ReadWriter
	Roots[`/bufio/ReadWriter`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bufio.ReadWriter
	Roots[`/bufio/ReadWriter`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bufio.Reader
	Roots[`/bufio/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bufio.Writer
	Roots[`/bufio/Writer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bytes.Buffer
	Roots[`/bytes/Buffer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *bytes.Reader
	Roots[`/bytes/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *encoding_base64.Encoding
	Roots[`/encoding/base64/Encoding`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *math_big.Int
	Roots[`/math/big/Int`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *math_big.Rat
	Roots[`/math/big/Rat`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Client
	Roots[`/net/http/Client`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Cookie
	Roots[`/net/http/Cookie`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ProtocolError
	Roots[`/net/http/ProtocolError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ProtocolError
	Roots[`/net/http/ProtocolError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Request
	Roots[`/net/http/Request`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Response
	Roots[`/net/http/Response`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.ServeMux
	Roots[`/net/http/ServeMux`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Server
	Roots[`/net/http/Server`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *net_http.Transport
	Roots[`/net/http/Transport`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.File
	Roots[`/os/File`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.LinkError
	Roots[`/os/LinkError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.PathError
	Roots[`/os/PathError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcAttr
	Roots[`/os/ProcAttr`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.Process
	Roots[`/os/Process`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.Process
	Roots[`/os/Process`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.ProcessState
	Roots[`/os/ProcessState`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *os.SyscallError
	Roots[`/os/SyscallError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Method
	Roots[`/reflect/Method`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.SliceHeader
	Roots[`/reflect/SliceHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StringHeader
	Roots[`/reflect/StringHeader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.StructField
	Roots[`/reflect/StructField`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.Value
	Roots[`/reflect/Value`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *reflect.ValueError
	Roots[`/reflect/ValueError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *regexp.Regexp
	Roots[`/regexp/Regexp`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strconv.NumError
	Roots[`/strconv/NumError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strings.Reader
	Roots[`/strings/Reader`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *strings.Replacer
	Roots[`/strings/Replacer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Location
	Roots[`/time/Location`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.ParseError
	Roots[`/time/ParseError`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Ticker
	Roots[`/time/Ticker`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Ticker
	Roots[`/time/Ticker`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Time
	Roots[`/time/Time`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Timer
	Roots[`/time/Timer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	{
	var tmp *time.Timer
	Roots[`/time/Timer`] = TypeRoot{ Type: reflect.ValueOf(tmp).Type().Elem() }
	}
	Roots[`/bytes/MinRead`] = ConstRoot{ Const: int64(bytes.MinRead) }
	Roots[`/math/E`] = ConstRoot{ Const: math.E }
	Roots[`/math/Ln10`] = ConstRoot{ Const: math.Ln10 }
	Roots[`/math/Ln2`] = ConstRoot{ Const: math.Ln2 }
	Roots[`/math/Log10E`] = ConstRoot{ Const: math.Log10E }
	Roots[`/math/Log2E`] = ConstRoot{ Const: math.Log2E }
	Roots[`/math/MaxFloat32`] = ConstRoot{ Const: math.MaxFloat32 }
	Roots[`/math/MaxFloat64`] = ConstRoot{ Const: math.MaxFloat64 }
	Roots[`/math/MaxInt16`] = ConstRoot{ Const: int64(math.MaxInt16) }
	Roots[`/math/MaxInt32`] = ConstRoot{ Const: int64(math.MaxInt32) }
	Roots[`/math/MaxInt64`] = ConstRoot{ Const: int64(math.MaxInt64) }
	Roots[`/math/MaxInt8`] = ConstRoot{ Const: int64(math.MaxInt8) }
	Roots[`/math/MaxUint16`] = ConstRoot{ Const: int64(math.MaxUint16) }
	Roots[`/math/MaxUint32`] = ConstRoot{ Const: int64(math.MaxUint32) }
	Roots[`/math/MaxUint64`] = ConstRoot{ Const: uint64(math.MaxUint64) }
	Roots[`/math/MaxUint8`] = ConstRoot{ Const: int64(math.MaxUint8) }
	Roots[`/math/MinInt16`] = ConstRoot{ Const: int64(math.MinInt16) }
	Roots[`/math/MinInt32`] = ConstRoot{ Const: int64(math.MinInt32) }
	Roots[`/math/MinInt64`] = ConstRoot{ Const: int64(math.MinInt64) }
	Roots[`/math/MinInt8`] = ConstRoot{ Const: int64(math.MinInt8) }
	Roots[`/math/Phi`] = ConstRoot{ Const: math.Phi }
	Roots[`/math/Pi`] = ConstRoot{ Const: math.Pi }
	Roots[`/math/SmallestNonzeroFloat32`] = ConstRoot{ Const: math.SmallestNonzeroFloat32 }
	Roots[`/math/SmallestNonzeroFloat64`] = ConstRoot{ Const: math.SmallestNonzeroFloat64 }
	Roots[`/math/Sqrt2`] = ConstRoot{ Const: math.Sqrt2 }
	Roots[`/math/SqrtE`] = ConstRoot{ Const: math.SqrtE }
	Roots[`/math/SqrtPhi`] = ConstRoot{ Const: math.SqrtPhi }
	Roots[`/math/SqrtPi`] = ConstRoot{ Const: math.SqrtPi }
	Roots[`/math/big/MaxBase`] = ConstRoot{ Const: math_big.MaxBase }
	Roots[`/net/http/DefaultMaxHeaderBytes`] = ConstRoot{ Const: int64(net_http.DefaultMaxHeaderBytes) }
	Roots[`/net/http/DefaultMaxIdleConnsPerHost`] = ConstRoot{ Const: int64(net_http.DefaultMaxIdleConnsPerHost) }
	Roots[`/net/http/StatusAccepted`] = ConstRoot{ Const: int64(net_http.StatusAccepted) }
	Roots[`/net/http/StatusBadGateway`] = ConstRoot{ Const: int64(net_http.StatusBadGateway) }
	Roots[`/net/http/StatusBadRequest`] = ConstRoot{ Const: int64(net_http.StatusBadRequest) }
	Roots[`/net/http/StatusConflict`] = ConstRoot{ Const: int64(net_http.StatusConflict) }
	Roots[`/net/http/StatusContinue`] = ConstRoot{ Const: int64(net_http.StatusContinue) }
	Roots[`/net/http/StatusCreated`] = ConstRoot{ Const: int64(net_http.StatusCreated) }
	Roots[`/net/http/StatusExpectationFailed`] = ConstRoot{ Const: int64(net_http.StatusExpectationFailed) }
	Roots[`/net/http/StatusForbidden`] = ConstRoot{ Const: int64(net_http.StatusForbidden) }
	Roots[`/net/http/StatusFound`] = ConstRoot{ Const: int64(net_http.StatusFound) }
	Roots[`/net/http/StatusGatewayTimeout`] = ConstRoot{ Const: int64(net_http.StatusGatewayTimeout) }
	Roots[`/net/http/StatusGone`] = ConstRoot{ Const: int64(net_http.StatusGone) }
	Roots[`/net/http/StatusHTTPVersionNotSupported`] = ConstRoot{ Const: int64(net_http.StatusHTTPVersionNotSupported) }
	Roots[`/net/http/StatusInternalServerError`] = ConstRoot{ Const: int64(net_http.StatusInternalServerError) }
	Roots[`/net/http/StatusLengthRequired`] = ConstRoot{ Const: int64(net_http.StatusLengthRequired) }
	Roots[`/net/http/StatusMethodNotAllowed`] = ConstRoot{ Const: int64(net_http.StatusMethodNotAllowed) }
	Roots[`/net/http/StatusMovedPermanently`] = ConstRoot{ Const: int64(net_http.StatusMovedPermanently) }
	Roots[`/net/http/StatusMultipleChoices`] = ConstRoot{ Const: int64(net_http.StatusMultipleChoices) }
	Roots[`/net/http/StatusNoContent`] = ConstRoot{ Const: int64(net_http.StatusNoContent) }
	Roots[`/net/http/StatusNonAuthoritativeInfo`] = ConstRoot{ Const: int64(net_http.StatusNonAuthoritativeInfo) }
	Roots[`/net/http/StatusNotAcceptable`] = ConstRoot{ Const: int64(net_http.StatusNotAcceptable) }
	Roots[`/net/http/StatusNotFound`] = ConstRoot{ Const: int64(net_http.StatusNotFound) }
	Roots[`/net/http/StatusNotImplemented`] = ConstRoot{ Const: int64(net_http.StatusNotImplemented) }
	Roots[`/net/http/StatusNotModified`] = ConstRoot{ Const: int64(net_http.StatusNotModified) }
	Roots[`/net/http/StatusOK`] = ConstRoot{ Const: int64(net_http.StatusOK) }
	Roots[`/net/http/StatusPartialContent`] = ConstRoot{ Const: int64(net_http.StatusPartialContent) }
	Roots[`/net/http/StatusPaymentRequired`] = ConstRoot{ Const: int64(net_http.StatusPaymentRequired) }
	Roots[`/net/http/StatusPreconditionFailed`] = ConstRoot{ Const: int64(net_http.StatusPreconditionFailed) }
	Roots[`/net/http/StatusProxyAuthRequired`] = ConstRoot{ Const: int64(net_http.StatusProxyAuthRequired) }
	Roots[`/net/http/StatusRequestEntityTooLarge`] = ConstRoot{ Const: int64(net_http.StatusRequestEntityTooLarge) }
	Roots[`/net/http/StatusRequestTimeout`] = ConstRoot{ Const: int64(net_http.StatusRequestTimeout) }
	Roots[`/net/http/StatusRequestURITooLong`] = ConstRoot{ Const: int64(net_http.StatusRequestURITooLong) }
	Roots[`/net/http/StatusRequestedRangeNotSatisfiable`] = ConstRoot{ Const: int64(net_http.StatusRequestedRangeNotSatisfiable) }
	Roots[`/net/http/StatusResetContent`] = ConstRoot{ Const: int64(net_http.StatusResetContent) }
	Roots[`/net/http/StatusSeeOther`] = ConstRoot{ Const: int64(net_http.StatusSeeOther) }
	Roots[`/net/http/StatusServiceUnavailable`] = ConstRoot{ Const: int64(net_http.StatusServiceUnavailable) }
	Roots[`/net/http/StatusSwitchingProtocols`] = ConstRoot{ Const: int64(net_http.StatusSwitchingProtocols) }
	Roots[`/net/http/StatusTeapot`] = ConstRoot{ Const: int64(net_http.StatusTeapot) }
	Roots[`/net/http/StatusTemporaryRedirect`] = ConstRoot{ Const: int64(net_http.StatusTemporaryRedirect) }
	Roots[`/net/http/StatusUnauthorized`] = ConstRoot{ Const: int64(net_http.StatusUnauthorized) }
	Roots[`/net/http/StatusUnsupportedMediaType`] = ConstRoot{ Const: int64(net_http.StatusUnsupportedMediaType) }
	Roots[`/net/http/StatusUseProxy`] = ConstRoot{ Const: int64(net_http.StatusUseProxy) }
	Roots[`/net/http/TimeFormat`] = ConstRoot{ Const: net_http.TimeFormat }
	Roots[`/os/DevNull`] = ConstRoot{ Const: os.DevNull }
	Roots[`/os/ModeAppend`] = ConstRoot{ Const: os.ModeAppend }
	Roots[`/os/ModeCharDevice`] = ConstRoot{ Const: os.ModeCharDevice }
	Roots[`/os/ModeDevice`] = ConstRoot{ Const: os.ModeDevice }
	Roots[`/os/ModeDir`] = ConstRoot{ Const: os.ModeDir }
	Roots[`/os/ModeExclusive`] = ConstRoot{ Const: os.ModeExclusive }
	Roots[`/os/ModeNamedPipe`] = ConstRoot{ Const: os.ModeNamedPipe }
	Roots[`/os/ModePerm`] = ConstRoot{ Const: os.ModePerm }
	Roots[`/os/ModeSetgid`] = ConstRoot{ Const: os.ModeSetgid }
	Roots[`/os/ModeSetuid`] = ConstRoot{ Const: os.ModeSetuid }
	Roots[`/os/ModeSocket`] = ConstRoot{ Const: os.ModeSocket }
	Roots[`/os/ModeSticky`] = ConstRoot{ Const: os.ModeSticky }
	Roots[`/os/ModeSymlink`] = ConstRoot{ Const: os.ModeSymlink }
	Roots[`/os/ModeTemporary`] = ConstRoot{ Const: os.ModeTemporary }
	Roots[`/os/ModeType`] = ConstRoot{ Const: os.ModeType }
	Roots[`/os/O_APPEND`] = ConstRoot{ Const: os.O_APPEND }
	Roots[`/os/O_CREATE`] = ConstRoot{ Const: os.O_CREATE }
	Roots[`/os/O_EXCL`] = ConstRoot{ Const: os.O_EXCL }
	Roots[`/os/O_RDONLY`] = ConstRoot{ Const: os.O_RDONLY }
	Roots[`/os/O_RDWR`] = ConstRoot{ Const: os.O_RDWR }
	Roots[`/os/O_SYNC`] = ConstRoot{ Const: os.O_SYNC }
	Roots[`/os/O_TRUNC`] = ConstRoot{ Const: os.O_TRUNC }
	Roots[`/os/O_WRONLY`] = ConstRoot{ Const: os.O_WRONLY }
	Roots[`/os/PathListSeparator`] = ConstRoot{ Const: os.PathListSeparator }
	Roots[`/os/PathSeparator`] = ConstRoot{ Const: os.PathSeparator }
	Roots[`/os/SEEK_CUR`] = ConstRoot{ Const: os.SEEK_CUR }
	Roots[`/os/SEEK_END`] = ConstRoot{ Const: os.SEEK_END }
	Roots[`/os/SEEK_SET`] = ConstRoot{ Const: os.SEEK_SET }
	Roots[`/reflect/Array`] = ConstRoot{ Const: reflect.Array }
	Roots[`/reflect/Bool`] = ConstRoot{ Const: reflect.Bool }
	Roots[`/reflect/BothDir`] = ConstRoot{ Const: reflect.BothDir }
	Roots[`/reflect/Chan`] = ConstRoot{ Const: reflect.Chan }
	Roots[`/reflect/Complex128`] = ConstRoot{ Const: reflect.Complex128 }
	Roots[`/reflect/Complex64`] = ConstRoot{ Const: reflect.Complex64 }
	Roots[`/reflect/Float32`] = ConstRoot{ Const: reflect.Float32 }
	Roots[`/reflect/Float64`] = ConstRoot{ Const: reflect.Float64 }
	Roots[`/reflect/Func`] = ConstRoot{ Const: reflect.Func }
	Roots[`/reflect/Int`] = ConstRoot{ Const: reflect.Int }
	Roots[`/reflect/Int16`] = ConstRoot{ Const: reflect.Int16 }
	Roots[`/reflect/Int32`] = ConstRoot{ Const: reflect.Int32 }
	Roots[`/reflect/Int64`] = ConstRoot{ Const: reflect.Int64 }
	Roots[`/reflect/Int8`] = ConstRoot{ Const: reflect.Int8 }
	Roots[`/reflect/Interface`] = ConstRoot{ Const: reflect.Interface }
	Roots[`/reflect/Invalid`] = ConstRoot{ Const: reflect.Invalid }
	Roots[`/reflect/Map`] = ConstRoot{ Const: reflect.Map }
	Roots[`/reflect/Ptr`] = ConstRoot{ Const: reflect.Ptr }
	Roots[`/reflect/RecvDir`] = ConstRoot{ Const: reflect.RecvDir }
	Roots[`/reflect/SendDir`] = ConstRoot{ Const: reflect.SendDir }
	Roots[`/reflect/Slice`] = ConstRoot{ Const: reflect.Slice }
	Roots[`/reflect/String`] = ConstRoot{ Const: reflect.String }
	Roots[`/reflect/Struct`] = ConstRoot{ Const: reflect.Struct }
	Roots[`/reflect/Uint`] = ConstRoot{ Const: reflect.Uint }
	Roots[`/reflect/Uint16`] = ConstRoot{ Const: reflect.Uint16 }
	Roots[`/reflect/Uint32`] = ConstRoot{ Const: reflect.Uint32 }
	Roots[`/reflect/Uint64`] = ConstRoot{ Const: reflect.Uint64 }
	Roots[`/reflect/Uint8`] = ConstRoot{ Const: reflect.Uint8 }
	Roots[`/reflect/Uintptr`] = ConstRoot{ Const: reflect.Uintptr }
	Roots[`/reflect/UnsafePointer`] = ConstRoot{ Const: reflect.UnsafePointer }
	Roots[`/strconv/IntSize`] = ConstRoot{ Const: int64(strconv.IntSize) }
	Roots[`/time/ANSIC`] = ConstRoot{ Const: time.ANSIC }
	Roots[`/time/April`] = ConstRoot{ Const: time.April }
	Roots[`/time/August`] = ConstRoot{ Const: time.August }
	Roots[`/time/December`] = ConstRoot{ Const: time.December }
	Roots[`/time/February`] = ConstRoot{ Const: time.February }
	Roots[`/time/Friday`] = ConstRoot{ Const: time.Friday }
	Roots[`/time/Hour`] = ConstRoot{ Const: time.Hour }
	Roots[`/time/January`] = ConstRoot{ Const: time.January }
	Roots[`/time/July`] = ConstRoot{ Const: time.July }
	Roots[`/time/June`] = ConstRoot{ Const: time.June }
	Roots[`/time/Kitchen`] = ConstRoot{ Const: time.Kitchen }
	Roots[`/time/March`] = ConstRoot{ Const: time.March }
	Roots[`/time/May`] = ConstRoot{ Const: time.May }
	Roots[`/time/Microsecond`] = ConstRoot{ Const: time.Microsecond }
	Roots[`/time/Millisecond`] = ConstRoot{ Const: time.Millisecond }
	Roots[`/time/Minute`] = ConstRoot{ Const: time.Minute }
	Roots[`/time/Monday`] = ConstRoot{ Const: time.Monday }
	Roots[`/time/Nanosecond`] = ConstRoot{ Const: time.Nanosecond }
	Roots[`/time/November`] = ConstRoot{ Const: time.November }
	Roots[`/time/October`] = ConstRoot{ Const: time.October }
	Roots[`/time/RFC1123`] = ConstRoot{ Const: time.RFC1123 }
	Roots[`/time/RFC1123Z`] = ConstRoot{ Const: time.RFC1123Z }
	Roots[`/time/RFC3339`] = ConstRoot{ Const: time.RFC3339 }
	Roots[`/time/RFC3339Nano`] = ConstRoot{ Const: time.RFC3339Nano }
	Roots[`/time/RFC822`] = ConstRoot{ Const: time.RFC822 }
	Roots[`/time/RFC822Z`] = ConstRoot{ Const: time.RFC822Z }
	Roots[`/time/RFC850`] = ConstRoot{ Const: time.RFC850 }
	Roots[`/time/RubyDate`] = ConstRoot{ Const: time.RubyDate }
	Roots[`/time/Saturday`] = ConstRoot{ Const: time.Saturday }
	Roots[`/time/Second`] = ConstRoot{ Const: time.Second }
	Roots[`/time/September`] = ConstRoot{ Const: time.September }
	Roots[`/time/Stamp`] = ConstRoot{ Const: time.Stamp }
	Roots[`/time/StampMicro`] = ConstRoot{ Const: time.StampMicro }
	Roots[`/time/StampMilli`] = ConstRoot{ Const: time.StampMilli }
	Roots[`/time/StampNano`] = ConstRoot{ Const: time.StampNano }
	Roots[`/time/Sunday`] = ConstRoot{ Const: time.Sunday }
	Roots[`/time/Thursday`] = ConstRoot{ Const: time.Thursday }
	Roots[`/time/Tuesday`] = ConstRoot{ Const: time.Tuesday }
	Roots[`/time/UnixDate`] = ConstRoot{ Const: time.UnixDate }
	Roots[`/time/Wednesday`] = ConstRoot{ Const: time.Wednesday }
}
