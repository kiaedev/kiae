package api

//go:generate buf mod update
//go:generate buf generate
//go:generate go run github.com/favadi/protoc-go-inject-tag -input */*.pb.go

//go:generate go run github.com/99designs/gqlgen generate
