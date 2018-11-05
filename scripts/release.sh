if [ -n "$TAG" ]; then
  cd `dirname $0`/.. || exit 1
  echo "pwd: $PWD" || exit 1

  bash scripts/build.sh
  bash scripts/upload.sh
fi
