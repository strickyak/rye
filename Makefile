

test:
	set -ex; for x in *.test.sh ; do sh $$x ; done

clean:
	-rm *.pyc
