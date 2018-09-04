// +build main

package main

import (
	"bufio"
	"flag"
	"log"
	"regexp"
)

var src = flag.String("src", "", "Source filename")
var dest = flag.String("dest", "", "Destination filename")

// example: // $ 929 $ 53 $
// example: // @ 976 @ 54 @
var LINEMARK = regexp.MustCompile(`^\s*[/][/] @ ([0-9]+) @ ([0-9]+) @\s*$`).FindStringSubmatch

var n int64

func ParseInt(s string) {
	z, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot Parse Int: %q: %v", s, err)
	}
	return z
}

func Process(t string) {
	if m := LINEMARK(t); m != nil {
		
	}
}

func main() {
	flag.Parse()

	if *src == "" {
		log.Fatal("lineate: Flag --src required")
	}

	sc := new bufio.Scanner(os.Stdin)

	for sc.Scan() {
		++n
		t := sc.Text()
		Process(t)
	}
}
