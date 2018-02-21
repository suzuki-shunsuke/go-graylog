test:
	go test -covermode=atomic
cover:
	go test -coverprofile=coverage.txt -covermode=atomic
	go tool cover -html=coverage.txt
