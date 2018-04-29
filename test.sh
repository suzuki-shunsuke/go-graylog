decho() {
  echo "+ $@"
  eval $@
}

# gofmt
npm run fmt || exit 1

# golint
decho golint client/... terraform/... validator mockserver mockserver/store mockserver/handler mockserver/logic mockserver/seed mockserver/exec mockserver/store/plain || exit 1

decho go test ./mockserver/... -covermode=atomic || exit 1

if [ -f env.sh ]; then
  decho source env.sh
fi

decho go test ./util/... ./validator/... ./client/... . -covermode=atomic || exit 1
decho go test -v ./terraform/... -covermode=atomic || exit 1
