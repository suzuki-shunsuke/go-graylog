#!/usr/bin/env bash

set -eu
set -o pipefail

if [ -n "$CODECOV_TOKEN" ]; then
  bash <(curl -s https://codecov.io/bash)
fi
