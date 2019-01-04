all: test

test:
	sh ../scripts/run-tests.sh

clean: fresh

fresh:
	find . -name '*.want.-' -o -name '*.got.-' -o -name '*.got' | xargs rm -f
	sh ../scripts/rm_rye_binaries.sh
