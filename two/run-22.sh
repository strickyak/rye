case $1 in
  --noinline ) NOINLINE='--set noinline --noinline' ; FLAG='--noinline' ; shift ;;
  * ) NOINLINE='' ; FLAG='' ;;
esac

case $# in
  1 ) X="$1" ;;
  * ) echo "Needs 1 arg" >&2; exit 3 ;;
esac

set -ex

PYTHONPATH=.:.. python rye2.py build_builtins builtins2.py gen_builtins2.pre.go

sh expand-templates2.sh > template2.pre.go

go run ../../prego/main.go --set noinline --noinline < macros2.pre.go > macros2.go

for x in gen_builtins2 native2 runtime2 template2
do
  rm -f $x.go
  : $x
  go run ../../prego/main.go --source macros2.pre.go --set noinline --noinline < $x.pre.go > $x.go
  chmod -w $x.go
done
