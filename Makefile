
gen:
	go generate ./...

gen-api:
	go generate ./api/generate.go

gen-wire:
	go generate ./internal/app/server/wire_gen.go

build: gen
	go build -v -o build/ .