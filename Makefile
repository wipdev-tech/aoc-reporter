fmt:
	go fmt ./...

test: fmt
	go test ./...

lint: fmt
	golangci-lint run

