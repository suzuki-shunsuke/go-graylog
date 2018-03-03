TAG=edge

test:
	go test -covermode=atomic
cover:
	go test -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt
graylog-mock-server: *.go mock_server_bin/main.go
	go build -o graylog-mock-server mock_server_bin/main.go
# https://github.com/mitchellh/gox
# brew install gox
# go get github.com/mitchellh/gox
build:
	gox -output="{{.Dir}}/dist/$(TAG)/graylog-mock-server_$(TAG)_{{.OS}}_{{.Arch}}" -osarch="darwin/amd64 linux/amd64 windows/amd64" ./mock_server_bin
# https://github.com/tcnksm/ghr
# brew tap tcnksm/ghr
# brew install ghr
# go get -u github.com/tcnksm/ghr
upload:
	ghr $(TAG) mock_server_bin/dist/$(TAG)
