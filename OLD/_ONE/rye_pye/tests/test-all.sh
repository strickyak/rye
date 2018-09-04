#!/bin/bash
set -ex

D=$(dirname $0)
RYEC=${RYE:-python $D/../../rye.py}

function die() {
  echo "$*" >&2
  exit 2
}

for x in $D/*_test.py
do
  python "$x" || die "Test $x failed with python."
  python ../../rye.py run "$x" || die "Test $x failed with rye (RYEC=$RYEC)"
done
