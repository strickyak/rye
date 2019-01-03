RYEC=python2 ../compiler/rye.py

all: generated_builtins.go

generated_builtins.go: _FORCE
	rm -f generated_*
	$(RYEC) build_builtins ../builtins/builtins.py generated_builtins.go ../compiler/generated_accessors.tmp
	mv ../compiler/generated_accessors.tmp ../compiler/generated_accessors.py

clean:
	rm -f generated_*
fresh:

tidy:

_FORCE:
