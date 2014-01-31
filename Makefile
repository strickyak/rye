
all:  test run

test:
	set -ex; for x in *.test.sh ; do sh $$x ; done

run:
	sh compile-and-run.sh  test301.py 

clean:
	-rm *.pyc zzz.go
