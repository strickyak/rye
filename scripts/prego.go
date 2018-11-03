// +build main

// Copys the -i input_file to the -o output_file, which should be gofmt'ed go source.
// Recognizes conditional compilation marks like this:
// ```
// if 'x' {
//   stuff()....
// }
// ```
// and only if that letter (in the example, 'x') is in the --opts option (e.g. --opts=lmxyz contains x)
// is the stuff emitted.
//
// The inverse form is
// ```
// if !'x' {
//   stuff()....
// }
// ```
// in which case the stuff is emitted if the letter is not in the --opts.
// Notice that the syntax is syntactially correct go -- it is accepted by gofmt --
// but the compiler will not take it, since the character constant is not a bool.
//
// So you can use prego in places where if statements are not syntatically allowed
// (say, for fields of a struct), you can put two slashes in front of "if" and "}":
//
// type Dict struct {
//   Map    map[string]string
//   //if 'm' {
//   Mu     sync.Mutex
//   //}
// }
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var Opts = flag.String("opts", "", "single-letter options")
var In = flag.String("i", "", "Source filename")
var Out = flag.String("o", "", "Destination filename")

var CONDITIONAL = regexp.MustCompile(`^(\s*)[/]*\s*if\s+([!]?)\s*[']([A-Za-z0-9])[']\s*[{]\s*$`).FindStringSubmatch
var END_CONDITIONAL = regexp.MustCompile(`^\s*[/]*\s*[}]$`).FindString

var lineNum int64

func ConsumeConditionalBlock(ch <-chan string, w io.Writer, tabs string, enabled bool) {
	for {
		line, ok := <-ch
		if !ok {
			log.Panicf("EOF in Conditional Block")
		}
		if END_CONDITIONAL(line) != "" {
			break
		}
		if enabled {
			fmt.Fprintf(w, "%s%s\n", tabs, line)
		} else {
			fmt.Fprintf(w, "%s//-- %s\n", tabs, strings.TrimSpace(line))
		}
	}
}

func Consume(ch <-chan string, w io.Writer, done chan bool) {
	for {
		line, ok := <-ch
		if !ok {
			break
		}

		m := CONDITIONAL(line)
		if m == nil {
			fmt.Fprintf(w, "%s\n", line)
		} else {
			tabs := m[1]
			not := (m[2] == "!")
			opt := m[3][0]
			enabled := Bits[opt]
			if not {
				enabled = !enabled
			}
			fmt.Fprintf(w, "%s//if %s'%s' {\n", tabs, m[2], m[3])
			ConsumeConditionalBlock(ch, w, tabs, enabled)
			fmt.Fprintf(w, "%s//}\n", tabs)
		}
	}
	done <- true
}

var Bits = [256]bool{}

func main() {
	flag.Parse()

	// Set Bits[] based on flags.
	for _, b := range *Opts {
		println(b)
		if b < 256 {
			Bits[b] = true
		}
	}

	ifd, err := os.Open(*In)
	if err != nil {
		log.Fatalf("prego: cannot open input file: %q: %v", *In, err)
	}
	defer ifd.Close()
	sc := bufio.NewScanner(ifd)

	ofd, err := os.Create(*Out)
	if err != nil {
		log.Fatalf("prego: cannot create output file: %q: %v", *In, err)
	}
	defer ofd.Close()
	w := bufio.NewWriter(ofd)
	defer w.Flush()

	ch, done := make(chan string), make(chan bool)
	go Consume(ch, w, done)
	for sc.Scan() {
		lineNum++
		t := sc.Text()
		ch <- t
	}
	close(ch)
	<-done
}
