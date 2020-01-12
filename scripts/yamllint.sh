#!/usr/bin/env bash

set -eu
set -o pipefail

cd "$(dirname "$0")/.."

git ls-files | grep -E ".*\.ya?ml$" | xargs yamllint -c .yamllint.yml
