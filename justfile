default: lint build test

# install development dependencies
install-dev-deps:
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.9
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.10
  go install github.com/envoyproxy/protoc-gen-validate@v0.6.12

lint:
  golangci-lint run

gen:
	go generate ./...

gen-api:
	go generate ./api/generate.go

gen-wire:
	go generate ./internal/app/server/wire_gen.go

build: gen
	go build -v -o build/ .

test: build
  go test ./...

# install just from crates.io
install:
  go install .