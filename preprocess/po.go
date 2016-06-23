/*
Preprocess a .po file into a .go file.

Usage:

   PO=sym1,sym2,sym3...  go run preprocess/main.go < runtime.po > runtime.go

Po Syntax:

  //#if sym1
  ...
  //#endif

*/
package main

import . "fmt"
import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var Match = regexp.MustCompile(`[ \t]*//#([a-z]+)[ \t]*([A-Za-z0-9_]*)[ \t]*$`).FindStringSubmatch

var Stack = []bool{true}
var Vars = make(map[string]bool)

func doLine(lineNum int, s string) {
	m := Match(s)

	if m != nil {
		switch m[1] {
		case "if":
			pred, _ := Vars[m[2]]
			Stack = append(Stack, pred)
		case "endif":
	    n := len(Stack)
      if n < 2 {
			  log.Fatalf("Line %d: Unmatched #endif", lineNum)
      }
			Stack = Stack[:n-1]
		default:
			log.Fatalf("Line %d: Unknown control: %q", lineNum, m[1])
		}
		Println("")

	} else {
    printing := true
    for _, e := range Stack {
      if !e {
        printing = false
      }
    }

		if printing {
			Println(s)
		} else {
			Println("")
		}
	}
}

func main() {
	po := os.Getenv("PO")
	for _, s := range strings.Split(po, ",") {
		Vars[s] = true
	}

	bs := bufio.NewScanner(os.Stdin)
  lineNum := 0
	for bs.Scan() {
    lineNum++
		doLine(lineNum, bs.Text())
	}
}
