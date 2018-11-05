cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source scripts/decho.sh || exit 1

decho mkdir -p coverage/$1 || exit 1
decho go test ./$1 -coverprofile=coverage/$1/coverage.txt -covermode=atomic || exit 1
decho go tool cover -html=coverage/$1/coverage.txt
