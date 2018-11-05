cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source scripts/decho.sh || exit 1

# gofmt
echo "! git ls-files | grep \"\\.go$\" | xargs gofmt -s -d | grep '^'" || exit 1
! git ls-files | grep "\.go$" | xargs gofmt -s -d | grep '^'  || exit 1
# go vet
decho go vet $(go list ./... | grep -v /vendor/)
