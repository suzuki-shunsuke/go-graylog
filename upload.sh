if [ -n "$TRAVIS_TAG" ]; then
  make upload-dep build upload TAG=$TRAVIS_TAG
fi
