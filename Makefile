TAG=edge

coverage:
	mkdir coverage
t:
	go test ./... -covermode=atomic
cover: coverage
	go test . -coverprofile=coverage/main.txt -covermode=atomic
	go tool cover -html=coverage/main.txt
cover-client: coverage
	go test ./client -coverprofile=coverage/client.txt -covermode=atomic
	go tool cover -html=coverage/client.txt
cover-endpoint: coverage
	go test ./client/endpoint -coverprofile=coverage/endpoint.txt -covermode=atomic
	go tool cover -html=coverage/endpoint.txt
cover-mockserver: coverage *.go mockserver/*.go
	go test ./mockserver -coverprofile=coverage/mockserver.txt -covermode=atomic
	go tool cover -html=coverage/mockserver.txt
cover-handler: coverage
	go test ./mockserver/handler -coverprofile=coverage/handler.txt -covermode=atomic
	go tool cover -html=coverage/handler.txt
cover-logic: coverage
	go test ./mockserver/logic -coverprofile=coverage/logic.txt -covermode=atomic
	go tool cover -html=coverage/logic.txt
cover-store: coverage
	go test ./mockserver/store -coverprofile=coverage/store.txt -covermode=atomic
	go tool cover -html=coverage/store.txt
cover-plain: coverage
	go test ./mockserver/store/plain -coverprofile=coverage/plain.txt -covermode=atomic
	go tool cover -html=coverage/plain.txt
build:
	gox -output="dist/$(TAG)/graylog-mock-server_$(TAG)_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./mockserver/exec
	gox -output="dist/$(TAG)/terraform-provider-graylog_$(TAG)_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./terraform
upload:
	# GITHUB_TOKEN envrionment variable
	ghr -u suzuki-shunsuke $(TAG) dist/$(TAG)
upload-dep:
	go get -u github.com/tcnksm/ghr
	go get github.com/mitchellh/gox
