all: test

test: _numbers _lisp _rye

_numbers:
	sh compile-and-run.sh test301.py
	sh compile-and-run.sh test302.py
	sh compile-and-run.sh test303.py
	sh compile-and-run.sh test304.py
	sh compile-and-run.sh test305.py
	sh compile-and-run.sh test306.py
	sh compile-and-run.sh test307.py
	#sh compile-and-run.sh test401.py test401.want
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
	python rye.py build test401.py twice.py
	test401/test401
	test401/test401 | diff - test401.want
	sh test_rye.sh lisp.py
clean:
	-rm *.pyc zzz zzz.*
	-rm -r test[0-9][0-9][0-9] lisp
