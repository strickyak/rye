set -e
trap 'cat zzz.tmp' 0

S="$1"

python compile.py "$S" >zzz.tmp 2>&1
trap '' 0

cat zzz.tmp | grep ^@@ | sed 's/^@@//'  > zzz.go

if gofmt < zzz.go > zzz.fmt
then
	mv zzz.go zzz.goo
	mv zzz.fmt zzz.go
fi

set -x
cat -n "$S"
cat -n zzz.go

time go build zzz.go

time -o zzz.time ./zzz > zzz.out 2>&1 || {
	cat -nev zzz.out
	echo %%%%%%%% EXECUTION FAILED -- $S >&2
	exit 1
}
cat -nev zzz.out
cat zzz.time
sed '/^##/d' zzz.out > zzz.got

case $2 in 
  "")
	# Run with python for comparison.
	python $S > zzz.want
	diff -u zzz.want zzz.got  &&  echo OKAY. >&2
	;;
  *)
	diff -u $2 zzz.got && echo OKAY. >&2
	;;
esac
