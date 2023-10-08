run: cmd/gx/main.go
	air

deploy: ./fly.toml
	pkgx fly deploy --now -y

build: cmd/gx/main.go
	go build -o ./build/gx ./cmd/gx/main.go

fmt:
	go fmt ./... && \
  bunx prettier -w views ./README.md ./docker-compose.yml

t: ./tests/
	go test -v ./tests/

encrypt: .env
	gpg -c .env

decrypt: .env.gpg
	gpg -d .env.gpg > .env
