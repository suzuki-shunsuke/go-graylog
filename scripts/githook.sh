#!/usr/bin/env bash

set -eu
set -o pipefail

ee() {
  echo "+ $*"
  eval "$@"
}

cd "$(dirname "$0")/.."
if [ ! -f .git/hooks/pre-commit ]; then
  ee ln -s ../../githooks/pre-commit.sh .git/hooks/pre-commit
fi
ee chmod a+x githooks/*
