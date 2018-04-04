TAG=edge
BIN=terraform-provider-graylog
$(BIN): *.go graylog/*.go
	go build -o $(BIN)
	terraform init
.terraform:
	terraform init
plan: $(BIN) .terraform
	terraform plan
apply: $(BIN) .terraform
	terraform apply
test:
	go test -count=1 -v -covermode=atomic ./...
cover:
	go test -count=1 -v -coverprofile=coverage.txt -covermode=atomic ./graylog
	go tool cover -html=coverage.txt
# https://github.com/mitchellh/gox
# brew install gox
# go get github.com/mitchellh/gox
build:
	gox -output="dist/$(TAG)/$(BIN)_$(TAG)_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" .
# https://github.com/tcnksm/ghr
# brew tap tcnksm/ghr
# brew install ghr
# go get -u github.com/tcnksm/ghr
upload:
	ghr $(TAG) dist/$(TAG)
