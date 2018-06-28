cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source script/decho.sh || exit 1

decho ghr -u suzuki-shunsuke ${TAG} dist/${TAG}
