all: a tests

RYEC=python rye.py

a: clean gen_builtins.go
	go install

#interp.bin: interp.py lex.py parse.py rye.py
#	$(RYEC) build interp.py

more: tests test-3

tests:
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test301.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test302.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test303.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test304.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test305.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test306.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test307.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test308.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test309.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test310.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/test311.py
	RYEC="$(RYEC)" sh scripts/test_rye.sh test/legacy/testecho.py
	:
	#$(RYEC) run interp.py --f=test302.py
	#interp/interp --f=test303.py
	#interp/interp --f=test304.py
	:
	$(RYEC) run test/legacy/test_gradtype.py
	:
	$(RYEC) build test/legacy/testbig.py
	#test/legacy/testbig.bin
	test/legacy/testbig.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test/legacy/testbig.want
	:
	$(RYEC) build test/legacy/test401.py
	#test/legacy/test401.bin
	test/legacy/test401.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test/legacy/test401.want
	:
	$(RYEC) build test/legacy/test402.py
	#test/legacy/test402.bin
	test/legacy/test402.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test/legacy/test402.want
	:
	$(RYEC) run test/legacy/testreflect.py
	$(RYEC) run test/legacy/test_6digits.py
	:
	sh scripts/test_rye.sh test/legacy/lisp.py
	echo With RYEC=$(RYEC) : tests ALL OKAY.

gen_builtins.go: builtins.py rye.py lex.py parse.py codegen.py linemap.py
	$(RYEC) build_builtins builtins.py gen_builtins.go
	go install github.com/strickyak/rye

rye.bin: gen_builtins.go rye.py lex.py parse.py codegen.py linemap.py
	$(RYEC) build rye.py

# rye-1 forces rye.bin to be built with python.
rye-1: rye.bin
	make RYEC='python rye.py' clean a rye.bin
	cp rye.bin ./rye-1
test-1: rye-1
	make RYEC='./rye-1' tests

# rye-2 forces rye.bin to be built with rye-1.
rye-2: rye-1
	make RYEC='./rye-1' clean a rye.bin
	cp rye.bin ./rye-2
test-2: test-1 rye-2
	make RYEC='./rye-2' tests

# rye-3 forces rye.bin to be built with rye-2.
rye-3: rye-2
	make RYEC='./rye-2' clean a rye.bin
	cp rye.bin ./rye-3
test-3: test-2 rye-3
	make RYEC='./rye-3' tests

#_ryerye2: rye/rye
#	python rye.py build rye.py
#	cp rye/rye ryerye1
#	rm -r rye lex parse codegen lib/data
#	./ryerye1 build rye.py
#	cp rye/rye ryerye2
#
#_ryerye3: _ryerye2
#	rm -r rye lex parse codegen lib/data
#	./ryerye2 build rye.py
#	cp rye/rye ryerye3
#
#_ryerye4: _ryerye3
#	make a
#	./ryerye3 build rye.py
#	cp rye/rye ryerye4

clean:
	-rm -f *.pyc *.bin test/*/*.bin
	-rm -f gen_builtins.go
	T=`find . -name ryemain.go` ; set -x ; for x in $$T ; do rm -f $$x ; rmdir `dirname $$x` || true ; done
	T=`find . -name ryemodule.go` ; set -x ; for x in $$T ; do rm -f $$x ; D=`dirname $$x` ; B=`basename $$D` ; rmdir $$D ; done
