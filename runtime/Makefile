RYEC=python2 ../compiler/rye.py

all: _FORCE
	rm -f generated*
	sed -e '1s/.*//' -e 's/\<T\>/PDict/g' -e 's/_T/_PDict/g' xdict.template.go | go run ../scripts/prego.go -i /dev/stdin -o generated_dict.go --opts=""
	sed -e '1s/.*//' -e 's/\<T\>/PSyncDict/g' -e 's/_T/_PSyncDict/g' xdict.template.go | go run ../scripts/prego.go -i /dev/stdin -o generated_syncdict.go --opts="m"
	$(RYEC) build_builtins builtins.py generated_builtins.go generated_accessors.tmp
	mv generated_accessors.tmp ../compiler/generated_accessors.py

clean:
	rm -f generated*
fresh:

tidy:

_FORCE:
