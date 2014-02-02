set -e
trap 'cat zzz.tmp' 0
python compile2.py "$@" >zzz.tmp 2>&1
trap '' 0

cat zzz.tmp | grep ^@@ | sed 's/^@@//'  > ./zzz.go
set -x
cat -n ./zzz.go
go run ./zzz.go
