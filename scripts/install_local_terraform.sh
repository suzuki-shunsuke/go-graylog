cd `dirname $0`/.. || exit 1
echo "pwd: $PWD" || exit 1

mkdir -p ~/.terraform.d/plugins || exit 1
go build -o ~/.terraform.d/plugins/terraform-provider-graylog ./terraform
