set -ex
for f in *.tr
do
	python check-lex.py $f ${f}out
done
