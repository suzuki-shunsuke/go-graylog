#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."

while read -r f; do
  if ! jsonlint -q "$f" > /dev/null 2>&1; then
    ee jsonlint "$f"
  fi
done < <(git ls-files | grep -E ".*\.json")
