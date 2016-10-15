all: a tests

RYEC=python rye.py
RYE2C=python rye2.py

a: clean gen_builtins.go runtime.go native.go fast.go goapi.py
	cd ../prego && go build main.go

a2: clean gen_builtins.go gen_builtins2.go runtime.go runtime2.go native2.go goapi.py
	cd ../prego && go build main.go

__FORCE:
	:

runtime.go: runtime.pre.go __FORCE
	rm -f runtime.go
	mkdir -p rye__$(OPTS)
	go run ../prego/main.go < runtime.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/runtime.go
	chmod -w rye__$(OPTS)/runtime.go

runtime2.go: runtime2.pre.go __FORCE
	rm -f runtime2.go
	mkdir -p rye__$(OPTS)
	go run ../prego/main.go --source macros2.pre.go < runtime2.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/runtime2.go
	chmod -w rye__$(OPTS)/runtime2.go

native.go: native.pre.go __FORCE
	rm -f native.go
	mkdir -p rye__$(OPTS)
	go run ../prego/main.go --source macros2.pre.go < native.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/native.go
	chmod -w rye__$(OPTS)/native.go

native2.go: native2.pre.go __FORCE
	rm -f native2.go
	mkdir -p rye__$(OPTS)
	go run ../prego/main.go --source macros2.pre.go < native2.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/native2.go
	chmod -w rye__$(OPTS)/native2.go

fast.go: fast.pre.go __FORCE
	rm -f fast.go
	mkdir -p rye__$(OPTS)
	go run ../prego/main.go --source macros2.pre.go < fast.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/fast.go
	chmod -w rye__$(OPTS)/fast.go

goapi.py: grok_goapi.py go1.txt
	rm -f goapi.py
	python grok_goapi.py < go1.txt > goapi.py
	chmod -w goapi.py

#interp.bin: interp.py lex.py parse.py rye.py
#	$(RYEC) build interp.py

more: tests test-3

tests:
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test301.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test302.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test303.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test304.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test305.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test306.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test307.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test308.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test309.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test310.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/test311.py
	RYEC="$(RYEC)" OPTS=$(OPTS) sh scripts/test_rye.sh test-legacy/testecho.py
	:
	#$(RYEC) --opts=$(OPTS) run interp.py --f=test302.py
	#interp/interp --f=test303.py
	#interp/interp --f=test304.py
	:
	$(RYEC) --opts=$(OPTS) run test-legacy/test_gradtype.py
	:
	$(RYEC) --opts=$(OPTS) build test-legacy/testbig.py
	#test-legacy/testbig.bin
	test-legacy/testbig.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test-legacy/testbig.want
	:
	$(RYEC) --opts=$(OPTS) build test-legacy/test401.py
	#test-legacy/test401.bin
	test-legacy/test401.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test-legacy/test401.want
	:
	$(RYEC) --opts=$(OPTS) build test-legacy/test402.py
	#test-legacy/test402.bin
	test-legacy/test402.bin | sed 's/[@][0-9][0-9][0-9][0-9][0-9]*/@99999/g' | diff -a - test-legacy/test402.want
	:
	$(RYEC) --opts=$(OPTS) run test-legacy/testreflect.py
	$(RYEC) --opts=$(OPTS) run test-legacy/test_6digits.py
	:
	sh scripts/test_rye.sh test-legacy/lisp.py
	echo With RYEC=$(RYEC) OPTS=$(OPTS) : tests ALL OKAY.

gen_builtins.go: builtins.py rye.py lex.py parse.py codegen.py linemap.py __FORCE
	rm -f gen_builtins.go
	mkdir -p rye__$(OPTS)/rye__$(OPTS)
	$(RYEC) --opts=$(OPTS) build_builtins builtins.py gen_builtins.pre.go
	go run ../prego/main.go --source macros2.pre.go < gen_builtins.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/gen_builtins.go

gen_builtins2.go: builtins2.py rye.py lex.py parse.py codegen.py linemap.py __FORCE
	rm -f gen_builtins2.pre.go gen_builtins2.go
	mkdir -p rye__$(OPTS)/rye__$(OPTS)
	$(RYE2C) build_builtins builtins2.py gen_builtins2.pre.go
	go run ../prego/main.go --source macros2.pre.go < gen_builtins2.pre.go | sed 's/package rye$$/package rye__$(OPTS)/' > rye__$(OPTS)/gen_builtins2.go

rye.bin: gen_builtins.go rye.py lex.py parse.py codegen.py linemap.py
	$(RYEC) --opts=$(OPTS) build rye.py

# rye-1 forces rye.bin to be built with python.
rye-1: rye.bin
	make RYEC='python rye.py' OPTS=$(OPTS) clean a rye.bin
	cp rye.bin ./rye-1
test-1: rye-1
	make RYEC='./rye-1' OPTS=$(OPTS) tests

# rye-2 forces rye.bin to be built with rye-1.
rye-2: rye-1
	make RYEC='./rye-1' OPTS=$(OPTS) clean a rye.bin
	cp rye.bin ./rye-2
test-2: test-1 rye-2
	make RYEC='./rye-2' OPTS=$(OPTS) tests

# rye-3 forces rye.bin to be built with rye-2.
rye-3: rye-2
	make RYEC='./rye-2' OPTS=$(OPTS) clean a rye.bin
	cp rye.bin ./rye-3
test-3: test-2 rye-3
	make RYEC='./rye-3' OPTS=$(OPTS) tests

#_ryerye2: rye/rye
#	python rye.py build rye.py
#	cp rye/rye ryerye1
#	rm -r rye lex parse codegen lib/data
#	./ryerye1 build rye.py
#	cp rye/rye ryerye2
#
#_ryerye3: _ryerye2
#	rm -r rye lex parse codegen lib/data
#	./ryerye2 build rye.py
#	cp rye/rye ryerye3
#
#_ryerye4: _ryerye3
#	make a
#	./ryerye3 build rye.py
#	cp rye/rye ryerye4

clean:
	-rm -f *.pyc */*.pyc *.bin */*.bin
	-rm -rf  rye__*/  */rye__*/
	-rm -f gen_builtins.go gen_builtins2.go runtime.go runtime2.go macros2.go macros.go native2.go templates.go
	T=`find . -name ryemain.go` ; set -x ; for x in $$T ; do rm -f $$x ; rmdir `dirname $$x` || true ; done
	T=`find . -name ryemodule.go` ; set -x ; for x in $$T ; do rm -f $$x ; D=`dirname $$x` ; B=`basename $$D` ; rmdir $$D ; done
