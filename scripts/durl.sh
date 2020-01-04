#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

find . \
  -type d -name .git -prune -o \
  -type d -name .terraform -prune -o \
  -type d -name dist -prune -o \
  -type f -print |
  durl check
