RYEO=  # Rye compile time options.
RYEC=python2 ../compiler/rye.py

all: _FORCE generated_.go
	rm -rf dbg_ opt_ rye_
	mkdir dbg_ opt_ rye_
	for x in *.go ; do \
		go run ../scripts/prego.go -i $$x -o rye_/generated_$$x --opts="$$RYECOPTS" \
		; done
	#for x in *.go ; do \
	#	go run ../scripts/prego.go -i $$x -o dbg_/generated_$$x --opts="$$RYECOPTS" \
	#	; done
	#for x in *.go ; do \
	#	go run ../scripts/prego.go -i $$x -o opt_/generated_$$x --opts="$$RYECOPTS" \
	#	; done
	cd rye_ && go install -x
	#cd dbg_ && go install -x
	#cd opt_ && go install -x
	: OKAY


generated_.go: _FORCE
	rm -f generated_.go
	$(RYEC) build_builtins ../builtins/builtins.py generated_builtins.go ../compiler/generated_builtins.tmp
	mv ../compiler/generated_builtins.tmp ../compiler/generated_builtins.py

clean:
	rm -f generated_.go
	rm -rf dbg_ opt_ rye_
fresh:

tidy:

_FORCE:
