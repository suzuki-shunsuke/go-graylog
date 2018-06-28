cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source script/decho.sh || exit 1

decho npm run metalint || exit 1
decho npm run golint || exit 1
echo "+ staticcheck (failure is ignored)"
decho staticcheck ./... || echo "staticcheck failure is ignored"
