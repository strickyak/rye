#!/bin/bash -ex
#
# Clone the Rye git repository into a /tmp/ directory and build and run tests there.

S="$(pwd)"
T=/tmp/cloned_test.$$
rm -rf $T
mkdir $T
cd $T
mkdir -p go/src/github.com/strickyak
cd go/src/github.com/strickyak
git clone "$S"
cd rye
if
  GOPATH="$T/go" make ddt clean all test rye-3
then
  echo "$0 OKAY."
  cd /
  rm -rf $T
else
  echo "$0 FAILED.   To clean up:  rm -rf $T"
fi
