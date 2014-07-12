set -e
trap 'cat zzz.go' 0

S="$1"

python compile.py "$S" >zzz.go 2>&1
trap '' 0

if gofmt < zzz.go > zzz.fmt
then
	mv zzz.go zzz.goo
	mv zzz.fmt zzz.go
fi

set -x
cat -n "$S"
cat -n zzz.go

time go build zzz.go

time -o zzz.time ./zzz > zzz.out 2>/dev/null || {
	cat -nev zzz.out
	echo %%%%%%%% EXECUTION FAILED -- $S >&2
	exit 1
}
cat -nev zzz.out
cat zzz.time
< zzz.out tr \" \' > zzz.got

case $2 in 
  ".")
	;;
  "")
	# Run with python for comparison.
	python $S | tr \" \' > zzz.want
	diff -u zzz.want zzz.got  &&  echo OKAY. >&2
	;;
  *)
	diff -u $2 zzz.got && echo OKAY. >&2
	;;
esac
