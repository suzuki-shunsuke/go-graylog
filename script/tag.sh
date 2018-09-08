# Usage
#   bash script/tag.sh v0.6.1

if [ $# -gt 1 ]; then
  echo "too many arguments" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

if [ $# -lt 1 ]; then
  echo "TAG argument is required" > /dev/stderr
  echo 'Usage tag.sh $TAG' > /dev/stderr
  exit 1
fi

TAG=$1
echo "TAG: $TAG"
VERSION=${TAG#v}

if [ "$TAG" = "$VERSION" ]; then
  echo "TAG must start with 'v'"
  exit 1
fi

echo "cd `dirname $0`/.."
cd `dirname $0`/..

echo "create version.go"
cat << EOS > version.go
package domain

// Version is the go-graylog's version.
const Version = "$VERSION"
EOS

git add version.go
npm run release -- --release-as $TAG
