run: main.go
	air

fmt:
	go fmt ./...
	go vet ./...

test:
	go test
