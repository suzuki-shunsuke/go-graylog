set -e

echo "! git ls-files | grep .go | xargs gofmt -s -d | grep '^'"
! git ls-files | grep .go | xargs gofmt -s -d | grep '^'
