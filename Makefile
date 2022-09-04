
build-api:
	cd api && buf generate && gqlgen generate
	protoc-go-inject-tag -input="api/*/*.pb.go"

build: build-api
	go build -v -o build/ .