package api

//go:generate buf generate
//go:generate protoc-go-inject-tag -input */*.pb.go

//go:generate gqlgen generate
