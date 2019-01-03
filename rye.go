package main

import "os"
import "runtime/pprof"
import rye "github.com/strickyak/rye/runtime"
import MY "github.com/strickyak/rye/compiler/rye_/rye"

var _ = os.Args

func main() {
	rye.RememberRyeCompileOptions(``)

	ppfile := os.Getenv("RYE_PPROF")
	if ppfile != "" {
		f, err := os.Create(ppfile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	defer func() {
		// Catch and print FYI for uncaught outer exceptions.
		r := recover()
		if r != nil {
			if rye.DebugExcept < 1 {
				rye.DebugExcept = 1
			}
			rye.PrintStackFYIUnlessEOFBecauseExcept(r)
			panic(r)
		}
	}()

	defer rye.Flushem()
	MY.G___name__ = rye.MkStr("__main__")
	MY.Eval_Module()

	//  // This code was for interp:
	//  glbl := rye.MkDict(make(rye.Scope))
	//  for k, ptr := range MY.ModuleMap {
	//    glbl.SetItem(rye.MkStr(k), *ptr)
	//  }
	//  interp.Eval_Module()
	//  sco := interp.G_0_Scopes().(*interp.C_Scopes)
	//  sco.M_g = glbl
	//  interp.G_1_Repl(sco)

	MY.G_1_main(rye.MkStrs(os.Args[1:]))

	rye.Shutdown()
}
