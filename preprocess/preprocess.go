/*
Preprocess a .po file into a .go file.

Usage:

   PO=sym1,sym2,sym3...  go run preprocess/main.go < runtime.po > runtime.go

Po Syntax:

  //#if sym1
  ...
  //#endif

*/
package preprocess

import . "fmt"
import (
	//"bufio"
	"io"
	"log"
	//"os"
	"regexp"
	"strings"
)

var Match = regexp.MustCompile(`[ \t]*//#([a-z]+)[ \t]*([A-Za-z0-9_]*)[ \t]*$`).FindStringSubmatch

var MatchMacroCall = regexp.MustCompile(`\binline[.]([A-Za-z0-9_]+)[(]([^()]*)[)]`)
var MatchIdentifier = regexp.MustCompile(`[A-Za-z0-9_]+`)

type Macro struct {
	Args   []string
	Body   string
	Result string
}

type Po struct {
  Macros map[string]*Macro
  Switches map[string]bool
  Stack []bool
  W io.Writer
}

func Fatalf(s string, args ...interface{}) {
	log.Fatalf("po preprocessor: ERROR: "+s, args...)
}

func (po *Po) replaceFromMap(s string, subs map[string]string) string {
	if z, ok := subs[s]; ok {
		return z
	}
	return s
}

func (po *Po) replaceMacroInvocation(s string) string {
  println("// replaceMacroInvocation:", s)
	m := MatchMacroCall.FindStringSubmatch(s)
	if len(m) != 3 {
		Fatalf("bad len from MatchMacroCall.FindStringSubmatch")
	}
	name := m[1]
	argtext := m[2]
	argwords := strings.Split(argtext, ",")

	macro, ok := po.Macros[name]
	if !ok {
		Fatalf("unknown macro: %q", name)
	}
	if len(argwords) != len(macro.Args) {
		Fatalf("got %d args for macro %q, but wanted %d args", len(argwords), name, len(macro.Args))
	}

	subs := make(map[string]string)
	for i, arg := range macro.Args {
		subs[arg] = argwords[i]
	}
	replacer := func(word string) string { return po.replaceFromMap(word, subs) }

	for _, line := range strings.Split(macro.Body, ",") {
		if len(line) > 0 {
			l2 := MatchIdentifier.ReplaceAllStringFunc(line, replacer)
			l3 := po.SubstitueMacros(l2)
			Fprintln(po.W, l3)
		}
	}

	z := MatchIdentifier.ReplaceAllStringFunc(macro.Result, replacer)
	return po.SubstitueMacros(z)
}

func (po *Po) SubstitueMacros(s string) string {
  println("// SubstitueMacros:", s)
	return MatchMacroCall.ReplaceAllStringFunc(s, po.replaceMacroInvocation)
}

func (po *Po) DoLine(lineNum int, s string) {
	m := Match(s)

	if m != nil {
		switch m[1] {
		case "if":
			pred, _ := po.Switches[m[2]]
			po.Stack = append(po.Stack, pred)
		case "endif":
			n := len(po.Stack)
			if n < 2 {
				Fatalf("Line %d: Unmatched #endif", lineNum)
			}
			po.Stack = po.Stack[:n-1]
		default:
			Fatalf("Line %d: Unknown control: %q", lineNum, m[1])
		}
		Fprintln(po.W, "")

	} else {
		printing := true
		for _, e := range po.Stack {
			if !e {
				printing = false
			}
		}

		if printing {
			Fprintln(po.W, po.SubstitueMacros(s))
		} else {
			Fprintln(po.W, "")
		}
	}
}
