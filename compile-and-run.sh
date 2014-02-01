python compile.py "$@" 2>&1 | tee zzz.tmp

cat zzz.tmp | grep ^@@ | sed 's/^@@//'  > ./zzz.go
set -x
cat -n ./zzz.go
go run ./zzz.go
