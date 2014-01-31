python compile.py "$@" | grep ^@@ | sed 's/^@@//'  > ./zzz.go
cat -n ./zzz.go
go run ./zzz.go
