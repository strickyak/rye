set -ex
RYEC=${RYEC:-python rye.py}

$RYEC build "$1"
B=`basename "$1" .py`
D=`dirname "$1"`
T=/tmp/tmp.rye.$$
"$D/$B.bin" > ${T}_r 2>/dev/null
python "$1" > ${T}_p

tr \" \' < ${T}_r > ${T}_r2
tr \" \' < ${T}_p > ${T}_p2

cat -nv ${T}_r2
diff ${T}_p2 ${T}_r2
rm -f ${T}_r ${T}_p ${T}_p2 ${T}_r2
