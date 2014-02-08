
all:  test run

test:
	set -ex; for x in *.test.sh ; do sh $$x ; done

run:
	# sh compile-and-run.sh  test301.py 
	# sh compile-and-run.sh  test302.py 
	# sh compile-and-run.sh  test303.py 
	# sh compile-and-run.sh  test304.py 
	:
	sh compile-and-run.sh test301.py
	sh compile-and-run.sh test302.py
	sh compile-and-run.sh test303.py
	sh compile-and-run.sh test304.py
	sh compile-and-run.sh test306.py
	sh compile-and-run.sh test305.py

clean:
	-rm *.pyc zzz.*
