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
done < <(
  find . \
    -type d -name .git -prune -o \
    -type d -name .terraform -prune -o \
    -type f -name "*.json" -print
)
