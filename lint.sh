echo "+ gometalinter"
npm run metalint || exit 1
echo "+ golint"
golint ./mockserver ./client/... ./terraform/... ./mockserver/logic ./mockserver/store ./mockserver/store/plain ./mockserver/handler ./validator || exit 1
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
