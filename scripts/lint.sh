cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

npm run metalint || exit 1
npm run golint || exit 1
