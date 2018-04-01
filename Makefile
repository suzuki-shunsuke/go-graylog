TAG=edge

coverage:
	mkdir coverage
t:
	go test ./... -covermode=atomic
cover-client: coverage
	go test -coverprofile=coverage/client.txt -covermode=atomic
	go tool cover -html=coverage/client.txt
cover-mockserver: coverage *.go mockserver/*.go
	go test ./mockserver -coverprofile=coverage/mockserver.txt -covermode=atomic
	go tool cover -html=coverage/mockserver.txt
cover-logic: coverage
	go test ./mockserver/logic -coverprofile=coverage/logic.txt -covermode=atomic
	go tool cover -html=coverage/logic.txt
cover-inmemory: coverage
	go test ./mockserver/store/inmemory -coverprofile=coverage/inmemory.txt -covermode=atomic
	go tool cover -html=coverage/inmemory.txt
graylog-mock-server: *.go mockserver/*.go mockserver/exec/main.go
	go build -o graylog-mock-server mockserver/exec/main.go
# https://github.com/mitchellh/gox
# brew install gox
# go get github.com/mitchellh/gox
build:
	gox -output="dist/$(TAG)/graylog-mock-server_$(TAG)_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./mockserver/exec
# https://github.com/tcnksm/ghr
# brew tap tcnksm/ghr
# brew install ghr
# go get -u github.com/tcnksm/ghr
upload:
	ghr $(TAG) dist/$(TAG)
