#!/bin/bash
set -ex

REPO=$1

T=/tmp/historical.$$
P=$T/src/github.com/strickyak
R=$T/src/github.com/strickyak/rye

#S=https://github.com/strickyak/rye.git

git log | while read label value
do
  case $label in
    commit)
      hash=$value
      continue
      ;;
    Date:)
      date=$( echo $value | sed 's;[-+].*;;' )
      secs=$( echo "puts [clock scan {$date}]" | tclsh )
      ;;
    *)
      continue
      ;;
  esac

  rm -rf $T
  mkdir -p $P
  git clone -- $PWD $R
  (
    cp benchmark3.py $R
    cd $R
    git checkout -f $hash

    export GOPATH=$T
    go get github.com/strickyak/rye
    make
    python rye.py build benchmark3.py
    find . -name \*benchmark3\* -ls

    BIN=$(test -f benchmark3.bin && echo ./benchmark3.bin || echo benchmark3/benchmark3)

    ls -s $BIN
    wc $BIN
    for M in 1 2 3 4 1 2 3 4 1 2 3 4 1 2 3 4 1 2 3 4
    do
      export GOMAXPROCS=$M 
      /usr/bin/time -f "{'what':'time', 'commit':'$hash', 'cpus':$M, 'real':%e, 'user':%U, 'sys':%S, 'rss':%M, 'exit':%x, 'date':'$date', 'secs':'$secs', }," $BIN >/dev/null 2>./log
      cat -n ./log
      tail ./log | grep "^..what.:.time.," >> /tmp/times.historical
      echo @@@@@@@@@@@@@@@@@@@@@@@@@@@@
      echo @@@@@@@@@@@@@@@@@@@@@@@@@@@@
      tail /tmp/times.historical
      echo @@@@@@@@@@@@@@@@@@@@@@@@@@@@
      echo @@@@@@@@@@@@@@@@@@@@@@@@@@@@
    done
  ) || true
done
