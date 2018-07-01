cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

source script/decho.sh || exit 1

gox -output="dist/${TAG}/graylog-mock-server_${TAG}_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./mockserver/exec || exit 1
gox -output="dist/${TAG}/terraform-provider-graylog_${TAG}_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./terraform/ || exit 1
ls dist/${TAG} | xargs -I {} gzip dist/${TAG}/{}
