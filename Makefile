all:
	rm -vrf $$(find . -name 'rye_' -type d -print)
	cd runtime && make
	cd compiler && make
	sed -e 's;^[/][/] [+]build rye_main;;' < compiler/rye_/rye_main.go | gofmt > ./rye.go
	(mkdir -p bin && cd bin && go build -x ../rye.go)

test: all
	rm -rf tests/rye_ emulation/tests/rye_
	sh emulation/tests/test-all.sh
	sh scripts/run-tests.sh

clean: fresh
	rm -vrf $$(find * -type d -name 'rye_')
	cd runtime && make clean
	cd compiler && make clean
	rm -vf compiler/rye ./rye
binclean: clean
	rm -vf bin/rye-[0-9]

# Make fresh leaves behind the .go files needed to compile rye.
fresh:
	cd runtime && make fresh
	cd compiler && make fresh
	# Delete the extensionless binary next to .py files.
	for x in $$(find [a-z]* -type f -name '*.py'); do \
		b=$$(dirname $$x)/$$(basename $$x .py) ; \
		rm -f $$b ; \
	done
	rm -vf $$(find [a-z]* -type f -name '*.pyc')

rye-1:
	cd runtime && make clean all
	cd compiler; make clean rye
	mkdir -p bin
	cp -v compiler/rye bin/rye-1
	rm -rf tests/rye_/
	RYE_CMD=$$(pwd)/bin/rye-1 sh scripts/run-tests.sh

rye-2: rye-1
	cd runtime && make RYEC=../bin/rye-1 clean all
	cd compiler; make RYEC=../bin/rye-1 clean rye
	mkdir -p bin
	cp -v compiler/rye bin/rye-2
	rm -rf tests/rye_/
	RYE_CMD=$$(pwd)/bin/rye-2 sh scripts/run-tests.sh

rye-3: rye-2
	cd runtime && make RYEC=../bin/rye-2 clean all
	cd compiler; make RYEC=../bin/rye-2 clean rye
	mkdir -p bin
	cp -v bin/rye bin/rye-3
	rm -rf tests/rye_/
	RYE_CMD=$$(pwd)/bin/rye-3 sh scripts/run-tests.sh

ci:
	make clean
	find * -type f -writable -print0 | xargs -0 ci -l -m/dev/null -t/dev/null -q
