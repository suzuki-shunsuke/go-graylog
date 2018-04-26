set -e

decho() {
  echo "+ $@"
  eval $@
}

# gofmt
echo "! git ls-files | grep \"\\.go$\" | xargs gofmt -s -d | grep '^'"
! git ls-files | grep "\.go$" | xargs gofmt -s -d | grep '^'
# go vet
decho go vet $(go list ./... | grep -v /vendor/)
