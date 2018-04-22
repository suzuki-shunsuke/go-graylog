set -e

decho() {
  echo "+ $@"
  eval $@
}

# gofmt
echo "! git ls-files | grep .go | xargs gofmt -s -d | grep '^'"
! git ls-files | grep .go | xargs gofmt -s -d | grep '^'

# golint
decho golint client terraform/... mockserver mockserver/store validator

decho go test ./mockserver/... -covermode=atomic

if [ -f env.sh ]; then
  decho source env.sh
fi

decho go test ./util/... ./validator/... ./client/... . -covermode=atomic
decho go test -v ./terraform/... -covermode=atomic
