set -e

decho() {
  echo "+ $@"
  eval $@
}

decho go test ./mockserver/... -covermode=atomic

if [ -f env.sh ]; then
  decho source env.sh
fi

decho go test ./util/... ./validator/... ./client/... . -covermode=atomic
decho go test -v ./terraform/... -covermode=atomic
