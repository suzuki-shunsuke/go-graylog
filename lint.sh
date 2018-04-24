set -e

echo "+ gofmt"
! git ls-files | grep .go | xargs gofmt -s -d | grep '^'
echo "+ go vet"
go vet $(go list ./... | grep -v /vendor/)
# https://github.com/golang/lint/issues/397
if which golint &> /dev/null; then
  echo "+ golint"
  golint ./mockserver ./client ./terraform/... ./mockserver/logic ./mockserver/store ./mockserver/store/plain ./mockserver/handler ./validator
fi
echo "+ gosimple"
! gosimple ./... | grep '^'
echo "+ unused"
unused ./...
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
