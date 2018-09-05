#!/bin/bash
set -ex

D=$(dirname $0)
RYEC=${RYE:-python $D/../../compiler/rye.py}

function die() {
  echo "$*" >&2
  exit 2
}

rm -rf tmp1

for x in $D/*_test.py
do
  python "$x" || die "Test $x failed with python."
  $RYEC run "$x" || die "Test $x failed with rye (RYEC=$RYEC)"
done
