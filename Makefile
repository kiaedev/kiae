
build: gen
	go build -v -o build/ .

gen-wire:
	go generate ./internal/app/server/wire_gen.go

gen-all:
	go generate ./...