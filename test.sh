decho() {
  echo "+ $@"
  eval $@
}

npm run fmt || exit 1
npm run golint

decho go test ./mockserver/... -covermode=atomic || exit 1

if [ -f env.sh ]; then
  decho source env.sh
fi

decho go test ./util/... ./validator/... ./client/... . -covermode=atomic || exit 1
decho go test -v ./terraform/... -covermode=atomic || exit 1
