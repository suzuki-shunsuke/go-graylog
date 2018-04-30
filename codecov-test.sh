#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

set -e
echo "" > coverage.txt

# ignore testutil from test coverage
go test ./testutil
for d in $(go list ./... | grep -v vendor | grep -v terraform | grep -v testutil); do
  go test -race -coverprofile=profile.out -covermode=atomic $d
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done

go test -v -race -coverprofile=profile.out -covermode=atomic ./terraform/...
if [ -f profile.out ]; then
  cat profile.out >> coverage.txt
  rm profile.out
fi
