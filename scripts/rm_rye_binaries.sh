#!/bin/bash

# Default to finding beneath files & directories starting with lowercase letter.
case $# in
	0) set [a-z]* ;;
esac

# Delete the extensionless binary next to .py files.
(
	for x in $(find "$@" -type f -name '*.py'); do \
		echo $(dirname $x)/$(basename $x .py)
	done
	for x in $(find "$@" -type f -name '*.ry'); do \
		echo $(dirname $x)/$(basename $x .ry)
	done
) | xargs rm -f
