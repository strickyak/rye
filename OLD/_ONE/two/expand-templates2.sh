#!/bin/bash
set -ex

case $1 in
  --noinline ) NOINLINE='--noinline --set noinline' ;;
  * ) NOINLINE='' ;;
esac

echo '// +build prego'
echo ''
echo 'package two'
echo 'import . "github.com/strickyak/rye"'
echo 'import "reflect"'
echo 'import "log"'
echo 'import "math"'
echo 'import "unsafe"'
echo ''
echo 'var _ = unsafe.Sizeof(0)'
echo 'var _ = log.Printf'
echo ''

(
  sed -e 's;OPNAME;Add;g' < binop2.tpl.go
  sed -e 's;OPNAME;Sub;g' < binop2.tpl.go
  sed -e 's;OPNAME;Mul;g' < binop2.tpl.go
  sed -e 's;OPNAME;Div;g' < binop2.tpl.go
  sed -e 's;OPNAME;Mod;g' < binop2.tpl.go
  sed -e 's;OPNAME;Pow;g' < binop2.tpl.go

  sed -e 's;OPNAME;LT;g' < relop2.tpl.go
  sed -e 's;OPNAME;LE;g' < relop2.tpl.go
  sed -e 's;OPNAME;GT;g' < relop2.tpl.go
  sed -e 's;OPNAME;GE;g' < relop2.tpl.go
  sed -e 's;OPNAME;EQ;g' < relop2.tpl.go
  sed -e 's;OPNAME;NE;g' < relop2.tpl.go
) |
sed -e 's;^package\s.*;;'

echo OKAY $0 >&2
