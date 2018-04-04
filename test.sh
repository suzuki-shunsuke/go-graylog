set -e

decho() {
  echo "+ $@"
  eval $@
}

decho go test ./client/... ./mockserver/... . -covermode=atomic
decho go test -v ./terraform/... -covermode=atomic
