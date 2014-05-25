all: test

test: _numbers _lisp
_numbers:
	sh compile-and-run.sh test301.py
	sh compile-and-run.sh test302.py
	sh compile-and-run.sh test303.py
	sh compile-and-run.sh test304.py
	sh compile-and-run.sh test305.py
	sh compile-and-run.sh test306.py
	sh compile-and-run.sh test307.py
	sh compile-and-run.sh test401.py test401.want
_lisp:
	sh compile-and-run.sh lisp.py    lisp.want

clean:
	-rm *.pyc zzz zzz.*
