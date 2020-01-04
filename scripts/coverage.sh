#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

if [ $# -eq 0 ]; then
  target="$(go list ./... | fzf)"
  if [ "$target" = "" ]; then
    exit 0
  fi
  target="${target#github.com/suzuki-shunsuke/go-graylog/v8/}"
elif [ $# -eq 1 ]; then
  target="$1"
else
  echo "too many arguments are given: $*" >&2
  exit 1
fi

if [ ! -d "$target" ]; then
  echo "$target is not found" >&2
  exit 1
fi

ee mkdir -p ".coverage/$target"
ee go test "./$target" -coverprofile=".coverage/$target/coverage.txt" -covermode=atomic
ee go tool cover -html=".coverage/$target/coverage.txt"
