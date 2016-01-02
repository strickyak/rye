all: clean gen_builtins.go test

RYEC=python rye.py

a: clean gen_builtins.go
	$(RYEC) build errfilt.py
b: a
	(cd lib ; python ../rye.py build data.py)
	$(RYEC) build rye.py

interp/interp: interp.py lex.py parse.py rye.py
	$(RYEC) build interp.py

test: a tests rye/rye interp/interp
more: test _ryerye2 _ryerye3 _ryerye4

gen_builtins.go: builtins.ry
	python build_builtins.py
	go install github.com/strickyak/rye

tests:
	RYEC=$(RYEC) sh test_rye.sh test301.py
	RYEC=$(RYEC) sh test_rye.sh test302.py
	RYEC=$(RYEC) sh test_rye.sh test303.py
	RYEC=$(RYEC) sh test_rye.sh test304.py
	RYEC=$(RYEC) sh test_rye.sh test305.py
	RYEC=$(RYEC) sh test_rye.sh test306.py
	RYEC=$(RYEC) sh test_rye.sh test307.py
	RYEC=$(RYEC) sh test_rye.sh test308.py
	RYEC=$(RYEC) sh test_rye.sh test309.py
	RYEC=$(RYEC) sh test_rye.sh test310.py
	RYEC=$(RYEC) sh test_rye.sh test311.py
	RYEC=$(RYEC) sh test_rye.sh testecho.py
	:
	$(RYEC) run interp.py --f=test302.py
	interp/interp --f=test303.py
	interp/interp --f=test304.py
	:
	$(RYEC) run test_gradtype.py
	:
	$(RYEC) build testbig.py
	testbig/testbig
	testbig/testbig | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - testbig.want
	:
	$(RYEC) build test401.py
	test401/test401
	test401/test401 | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test401.want
	:
	$(RYEC) build test402.py
	test402/test402
	test402/test402 | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test402.want
	:
	$(RYEC) run testreflect.py
	$(RYEC) run test_6digits.py
	:
	sh test_rye.sh lisp.py
	echo ALL OKAY.

# rye/rye uses whatever $(RYEC) has been set.
rye/rye: clean a rye.py lex.py parse.py codegen.py linemap.py
	$(RYEC) build rye.py

# rye-1 forces rye/rye to be built with python.
rye-1: rye/rye
	make RYEC='python rye.py' clean a rye/rye
	cp rye/rye ./rye-1
test-1: rye-1
  make RYEC='./rye-1' tests

# rye-2 forces rye/rye to be built with rye-1.
rye-2: rye-1
	make RYEC='./rye-1' clean a rye/rye
	cp rye/rye ./rye-2
test-2: test-1 rye-2
  make RYEC='./rye-2' tests

# rye-3 forces rye/rye to be built with rye-2.
rye-3: rye-2
	make RYEC='./rye-2' clean a rye/rye
	cp rye/rye ./rye-3
test-3: test-2 rye-3
  make RYEC='./rye-3' tests


_ryerye2: rye/rye
	python rye.py build rye.py
	cp rye/rye ryerye1
	rm -r rye lex parse codegen lib/data
	./ryerye1 build rye.py
	cp rye/rye ryerye2

_ryerye3: _ryerye2
	rm -r rye lex parse codegen lib/data
	./ryerye2 build rye.py
	cp rye/rye ryerye3

_ryerye4: _ryerye3
	make a
	./ryerye3 build rye.py
	cp rye/rye ryerye4

clean:
	-rm -f *.pyc
	-rm -f gen_builtins.go
	T=`find . -name ryemain.go` ; set -x ; for x in $$T ; do rm -f $$x ; rmdir `dirname $$x` || true ; done
	T=`find . -name ryemodule.go` ; set -x ; for x in $$T ; do rm -f $$x ; D=`dirname $$x` ; B=`basename $$D` ; rm -f $$D/$$B ; rmdir $$D ; done
