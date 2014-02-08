set -e
trap 'cat zzz.tmp' 0

S="$1"
W="$(basename $S .py).want"

python compile.py "$S" >zzz.tmp 2>&1
trap '' 0

cat zzz.tmp | grep ^@@ | sed 's/^@@//'  > zzz.go
set -x
cat -n "$S"
cat -n zzz.go
time go build zzz.go
time ./zzz > zzz.out 2>&1 || {
	cat -nev zzz.out
	echo %%%%%%%% EXECUTION FAILED -- $S >&2
	exit 1
}
cat -nev zzz.out

sed '/^##/d' zzz.out > zzz.got
if test -f $W
then
	diff $W zzz.got || {
		echo %%%%%%%% OUTPUT FAILS DIFF -- $S >&2
		exit 1
	}
else
	echo %%%%%%%% SAVING OUTPUT AS WANT -- $S >&2
	cp zzz.got $W
fi
echo OKAY $S.
