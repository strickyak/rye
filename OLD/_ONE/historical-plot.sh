#!/bin/bash

set -ex

gnuplot <<///
set terminal png enhanced size 1024,720 font "arial,10"
set output "historical.png"
plot "historical.data" with errorbars
///

eog historical.png
