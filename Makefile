run: cmd/gx/main.go
	air

temp: cmd/gx/main.go
	go build -o ./tmp/main ./cmd/gx/main.go

build: cmd/gx/main.go
	go build -o ./bin/gx ./cmd/gx/main.go

fmt:
	go fmt ./...
	go vet ./...

t: ./tests/
	go test -v ./tests/
