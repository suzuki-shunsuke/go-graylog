#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

git ls-files | grep -E ".*\.sh$" | xargs shellcheck env.sh.tmpl
