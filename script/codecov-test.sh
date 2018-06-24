#!/usr/bin/env bash
# https://github.com/codecov/example-go#caveat-multiple-files

echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor | grep -v dummy); do
  go test -race -coverprofile=profile.out -covermode=atomic $d
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
done
