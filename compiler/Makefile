RYEC=python2 rye.py

all: rye

rye: _FORCE_
	$(RYEC) build rye.py

clean: fresh
	rm -rf rye_dbg_ rye_opt_
fresh:
	rm -f *.tmp *.pyc rye

_FORCE_:
