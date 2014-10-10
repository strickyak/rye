all: clean _defs test

test: _defs _numbers _lisp _rye

_defs:
	python build_defs.py

_numbers:
	sh compile-and-run.sh test301.py
	sh compile-and-run.sh test302.py
	sh compile-and-run.sh test303.py
	sh compile-and-run.sh test304.py
	sh compile-and-run.sh test305.py
	sh compile-and-run.sh test306.py
	sh compile-and-run.sh test307.py
	sh compile-and-run.sh test308.py
	sh compile-and-run.sh test309.py
	sh compile-and-run.sh test311.py
_lisp:
	sh compile-and-run.sh lisp.py

_rye:
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
	python rye.py build testbig.py twice.py
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
	sh test_rye.sh lisp.py

clean:
	-rm *.pyc zzz zzz.* gen_*.go
	for x in */ryemodule.go ; do rm -r `dirname $$x`/ ; done
