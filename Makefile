test:
	go test -covermode=atomic
cover:
	go test -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt
graylog-mock-server: *.go mock_server_bin/main.go
	go build -o graylog-mock-server mock_server_bin/main.go
