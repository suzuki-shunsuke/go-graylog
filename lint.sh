npm run metalint || exit 1
npm run golint || exit 1
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
