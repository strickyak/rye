all: clean gen_builtins.go test

test: gen_builtins.go _rye

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
	:
	python rye.py build testbig.py
	testbig/testbig
	testbig/testbig | sed 's/[@][0-9][0-9]*/@@/g' | diff - testbig.want
	:
	python rye.py build test401.py twice.py
	test401/test401
	test401/test401 | sed 's/[@][0-9][0-9]*/@@/g' | diff - test401.want
	:
	python rye.py build test402.py twice.py
	test402/test402
	test402/test402 | sed 's/[@][0-9][0-9]*/@@/g' | diff - test402.want
	:
	python rye.py run testreflect.py
	:
	sh test_rye.sh lisp.py
	echo ALL OKAY.

clean:
	-rm -f *.pyc zzz zzz.* gen_*.go
	for x in `find */ryemodule.go -type f` ; do rm -r `dirname $$x`/ ; done
	-rm -f gen_builtins.go
