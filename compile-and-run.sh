set -e
trap 'cat zzz.tmp' 0
python compile.py "$@" >zzz.tmp 2>&1
trap '' 0

cat zzz.tmp | grep ^@@ | sed 's/^@@//'  > ./zzz.go
set -x
cat -n "$@"
cat -n ./zzz.go
time go build ./zzz.go
time ./zzz
