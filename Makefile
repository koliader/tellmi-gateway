.PHONY: air protoc test dev server

air:
	go install github.com/air-verse/air@latest

protoc:
	rm -f ./internal/pb/*.go
	mkdir -p ./internal/pb
	protoc -I ./proto \
	--go_out ./internal/pb --go_opt paths=source_relative \
	--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
	proto/*.proto

test:
	go test -v -cover ./...

dev:
	air -c .air.toml

server:
	go run cmd/main.go
