cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source script/decho.sh || exit 1

npm run fmt || exit 1
npm run golint || exit 1

decho go test ./mockserver/... -covermode=atomic || exit 1

if [ -f env.sh ]; then
  source env.sh
fi

decho go test ./testutil/... ./util/... ./validator/... ./client/... . -covermode=atomic || exit 1
decho go test -v ./terraform/... -covermode=atomic || exit 1
