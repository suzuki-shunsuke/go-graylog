set -e

decho() {
  echo "+ $@"
  eval $@
}

# go fmt
echo "! git ls-files | grep .go | xargs gofmt -s -d | grep '^'"
! git ls-files | grep .go | xargs gofmt -s -d | grep '^'

decho go test ./mockserver/... -covermode=atomic

if [ -f env.sh ]; then
  decho source env.sh
fi

decho go test ./util/... ./validator/... ./client/... . -covermode=atomic
decho go test -v ./terraform/... -covermode=atomic
