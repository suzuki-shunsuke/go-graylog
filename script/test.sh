decho() {
  echo "+ $@"
  eval $@
}

npm run fmt || exit 1
npm run golint

decho go test ./mockserver/... -covermode=atomic || exit 1

if [ -f script/env.sh ]; then
  decho source script/env.sh
fi

decho go test ./testutil/... ./util/... ./validator/... ./client/... . -covermode=atomic || exit 1
decho go test -v ./terraform/... -covermode=atomic || exit 1
