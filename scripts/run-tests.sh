RYE_CMD=${RYE_CMD:=python $PWD/compiler/rye.py}

case $PWD in
	*/rye ) 
		: ok ;;
	* )
		echo "ERROR: $0: Run this script from the rye directory, not from '$PWD'." >&2
		exit 2 ;;
esac

cd tests

set oldtest*.py p*.py test*.py

if [ $# -lt 25 ]
then
	echo "ERROR: $0: Did not glob enough tests: $*" >&2
	exit 2
fi

set -x
for x
do
	b=$(basename $x .py)
	$RYE_CMD -x -x run $x > $b.got || { echo "ERROR: Failure in $x" >&2; exit 13; }

	if test -r $b.want
	then
		sed -e 's/[@][0-9][0-9]*/@12345/g' < $b.want > $b.want.-
		sed -e 's/[@][0-9][0-9]*/@12345/g' < $b.got > $b.got.-
		(set -x ; diff $b.want.- $b.got.-) || { echo "ERROR: Output did not match golden output" >&2 ; exit 13; }
	fi
done

rm -f *.want.- *.got.- *.got
