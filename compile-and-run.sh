set -e
trap 'cat /tmp/zzz.go' 0

S="$1"

python compile.py "$S" >/tmp/zzz.go 2>&1
trap '' 0

if gofmt < /tmp/zzz.go > /tmp/zzz.fmt
then
	mv /tmp/zzz.go /tmp/zzz.goo
	mv /tmp/zzz.fmt /tmp/zzz.go
fi

set -x
cat -n "$S"
cat -n /tmp/zzz.go

time go build -o /tmp/zzz /tmp/zzz.go

time -o /tmp/zzz.time /tmp/zzz > /tmp/zzz.out 2>/dev/null || {
	cat -nev /tmp/zzz.out
	echo %%%%%%%% EXECUTION FAILED -- $S >&2
	exit 1
}
cat -nev /tmp/zzz.out
cat /tmp/zzz.time
< /tmp/zzz.out tr \" \' > /tmp/zzz.got

case $2 in 
  ".")
	;;
  "")
	# Run with python for comparison.
	python $S | tr \" \' > /tmp/zzz.want
	diff -u /tmp/zzz.want /tmp/zzz.got  &&  echo OKAY. >&2
	;;
  *)
	diff -u $2 /tmp/zzz.got && echo OKAY. >&2
	;;
esac
