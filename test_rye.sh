set -ex
python rye.py build "$1"
B=`basename "$1" .py`
"$B/$B" > _r 2>/dev/null
python "$1" > _p 2>&1

tr \" \' < _r > _r2
tr \" \' < _p > _p2

cat -nv _r2
diff _p2 _r2
