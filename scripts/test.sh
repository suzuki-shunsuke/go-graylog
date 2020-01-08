#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

if [ -f env.sh ]; then
  # shellcheck disable=SC1091
  source env.sh
fi

ee go test ./util/... ./validator/... ./client/... . -covermode=atomic
ee go test -v ./terraform/... -covermode=atomic
