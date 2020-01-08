#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor | grep -v terraform); do
  echo "$d"
  ee go test -race -coverprofile=profile.out -covermode=atomic "$d"
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done

ee go test -v -race -coverprofile=profile.out -covermode=atomic ./terraform/...
if [ -f profile.out ]; then
  cat profile.out >> coverage.txt
  rm profile.out
fi
