case $1 in
  --noinline ) NOINLINE='--set noinline --noinline' ; FLAG='--noinline' ; shift ;;
  * ) NOINLINE='' ; FLAG='' ;;
esac

case $# in
  1 ) X="$1" ;;
  * ) echo "Needs 1 arg" >&2; exit 3 ;;
esac


set -ex
make clean
rm -f gen_builtins.go gen_builtins2.go runtime.go runtime2.go macros2.go
rm -rf rye__
make a
echo ========================
echo ========================
echo ========================

: '<' macros2.pre.go '>' macros2.go
go run ../prego/main.go --set noinline --noinline < macros2.pre.go > macros2.go
wc macros2.pre.go  macros2.go

sh expand-templates.sh $FLAG

python rye2.py build_builtins builtins2.py gen_builtins2.pre.go
: '<' gen_builtins2.pre.go '>' gen_builtins2.go 
go run ../prego/main.go $NOINLINE --source macros2.pre.go < gen_builtins2.pre.go > gen_builtins2.go 

#gofmt -w gen_builtins2.go 

python rye2.py $FLAG build $X.py
: '<' rye__/$X/ryemodule.pre.go '>' rye__/$X/ryemodule.go
go run ../prego/main.go $NOINLINE --source macros2.pre.go < rye__/$X/ryemodule.pre.go > rye__/$X/ryemodule.go

go build -x rye__/$X/$X/ryemain.go
time ./ryemain 
echo OKAY '($X)' ./ryemain 
