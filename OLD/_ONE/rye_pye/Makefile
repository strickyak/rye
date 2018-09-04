all:
	python ../rye.py build md5.py
	python ../rye.py build os.py
	python ../rye.py build re.py
	python ../rye.py build sys.py

clean:
	T=`find . -name ryemain.go` ; set -x ; for x in $$T ; do rm $$x ; rmdir `dirname $$x` ; done
	T=`find . -name ryemodule.go` ; set -x ; for x in $$T ; do rm $$x ; D=`dirname $$x` ; B=`basename $$D` ; rm $$D/$$B ; rmdir $$D ; done
