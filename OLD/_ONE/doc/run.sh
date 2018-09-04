set -ex
cd `dirname $0`
exec python ../rye.py run doc_server.py
