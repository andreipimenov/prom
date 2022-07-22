.PHONY: build-server
build-server:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/server cmd/server/main.go

.PHONY: build-client
build-client:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o dist/client cmd/client/main.go

.PHONY: build
build: build-server build-client

.PHONY: up
up: build
	@cd deploy && docker-compose up -d

.PHONY: down
down:
	@cd deploy && docker-compose down --volumes
