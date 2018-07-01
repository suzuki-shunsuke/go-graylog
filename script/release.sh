if [ -n "$TRAVIS_TAG" ]; then
  cd `dirname $0`/.. || exit 1
  echo "pwd: $PWD" || exit 1

  bash script/build.sh
  bash script/upload.sh
fi
