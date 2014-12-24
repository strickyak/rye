all: clean gen_builtins.go test

a: clean gen_builtins.go
	python rye.py build errfilt.py

test: clean gen_builtins.go _rye rye/rye _ryerye
more: test _ryerye2 _ryerye3 _ryerye4

gen_builtins.go: builtins.ry
	python build_builtins.py
	go install github.com/strickyak/rye

_rye:
	python rye.py build errfilt.py
	:
	sh test_rye.sh test301.py
	sh test_rye.sh test302.py
	sh test_rye.sh test303.py
	sh test_rye.sh test304.py
	sh test_rye.sh test305.py
	sh test_rye.sh test306.py
	sh test_rye.sh test307.py
	sh test_rye.sh test308.py
	sh test_rye.sh test309.py
	sh test_rye.sh test311.py
	sh test_rye.sh testecho.py
	:
	python rye.py run testbig.py
	:
	python rye.py build testbig.py
	testbig/testbig
	testbig/testbig | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - testbig.want
	:
	python rye.py build test401.py
	test401/test401
	test401/test401 | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - test401.want
	:
	python rye.py build test402.py
	test402/test402
	test402/test402 | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - test402.want
	:
	python rye.py run testreflect.py
	:
	sh test_rye.sh lisp.py
	echo ALL OKAY.

rye/rye: a rye.py tr.py
	python rye.py build rye.py
	:
_ryerye: rye/rye
	rye/rye build testbig.py
	testbig/testbig | cat -n
	testbig/testbig | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - testbig.want
	:
	rye/rye build test401.py
	test401/test401 | cat -n
	test401/test401 | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - test401.want
	:
	rye/rye build test402.py
	test402/test402 | cat -n
	test402/test402 | sed 's/[@][0-9][0-9]*/@@/g' | diff -a - test402.want
	:
	: OKAY rye/rye

_ryerye2: _ryerye
	python rye.py build rye.py
	cp rye/rye ryerye1
	rm -r rye tr Eval
	./ryerye1 build rye.py
	cp rye/rye ryerye2

_ryerye3: _ryerye2
	rm -r rye tr Eval
	./ryerye2 build rye.py
	cp rye/rye ryerye3

_ryerye4: _ryerye3
	make a
	./ryerye3 build rye.py
	cp rye/rye ryerye4

clean:
	-rm -f *.pyc
	-rm -f gen_builtins.go
	T=`find . -name ryemain.go` ; set -x ; for x in $$T ; do rm -f $$x ; rmdir `dirname $$x` ; done
	T=`find . -name ryemodule.go` ; set -x ; for x in $$T ; do rm -f $$x ; D=`dirname $$x` ; B=`basename $$D` ; rm -f $$D/$$B ; rmdir $$D ; done
