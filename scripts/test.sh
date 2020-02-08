#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

if [ -f env.sh ]; then
  # shellcheck disable=SC1091
  source env.sh
fi

for d in $(go list ./... | grep -v terraform); do
  go test -race -covermode=atomic "$d"
done
go test -v ./graylog/terraform/... -covermode=atomic
